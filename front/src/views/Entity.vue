<template>

    <WorkSpace footerPosition="bottom" noBg v-on:scrolled-bottom="onScrolledBottom" ref="workspace">
        <template #header>
            <VSet>
                <VSet vertical indent-size="XS">
                    <VHead level="h1">Entity</VHead>
                    <VSet>
                        <VInput placeholder="Поиск" v-model="searchModel" @input="search"></VInput>
                        <VCheckbox v-model="entityFilter.WithFilter">
                            <VText font-size="txs">Show filters</VText>
                        </VCheckbox>
                    </VSet>
                </VSet>
            </VSet>
        </template>

        <template slot="content">
            <template>
                <VSet v-if="entityList && entityList.length" wrap class="pt14">
                    <template v-for="entityItem in partialList">
                        <EntityItem
                                :key="entityItem.Id"
                                :entityItem="entityItem"
                                :onEdit="editItem"
                                @onRequest="requestItem"
                        ></EntityItem>
                    </template>
                </VSet>
                <VText class="loading"></VText>
            </template>

            <VPanel v-if="panel.show" @close="closePanel" :maxWidth="panelMaxWidth">

                <VSet slot="header">
                    <VHead level="h3" width="fit">{{ panelHeader }} {{ currentEntityItem.Name }}</VHead>
                </VSet>

                <template slot="content">

                    <EntityRequest v-if="panel.type === panel.request" :entity="currentEntityItem"></EntityRequest>

                    <template v-else>
                        <VSet vertical @submit.prevent="saveChangesSubmit">

                            <VSet vertical indent-size="XS">
                                <VSign v-if="currentEntityItem.IsFilter">Filter name</VSign>
                                <VSign v-else>Entity name</VSign>
                                <VInput v-model="currentEntityItem.Name" :disabled="currentEntityItem.Id > 0"
                                        width="dyn"></VInput>

                                <template v-if="! currentEntityItem.IsFilter">
                                    <VSign>Http methods</VSign>
                                    <VSet>
                                        <VCheckbox v-if="isHttpMethods" v-model="currentEntityItem.HttpMethods.IsFind"
                                                   :disabled="currentEntityItem.Id > 0">
                                            <VText>Find</VText>
                                        </VCheckbox>
                                        <VCheckbox v-if="isHttpMethods" v-model="currentEntityItem.HttpMethods.IsCreate"
                                                   :disabled="currentEntityItem.Id > 0">
                                            <VText>Create</VText>
                                        </VCheckbox>
                                        <VCheckbox v-if="isHttpMethods" v-model="currentEntityItem.HttpMethods.IsRead"
                                                   :disabled="currentEntityItem.Id > 0">
                                            <VText>Read</VText>
                                        </VCheckbox>
                                        <VCheckbox v-if="isHttpMethods" v-model="currentEntityItem.HttpMethods.IsUpdate"
                                                   :disabled="currentEntityItem.Id > 0">
                                            <VText>Update</VText>
                                        </VCheckbox>
                                        <VCheckbox v-if="isHttpMethods" v-model="currentEntityItem.HttpMethods.IsDelete"
                                                   :disabled="currentEntityItem.Id > 0">
                                            <VText>Delete</VText>
                                        </VCheckbox>
                                    </VSet>
                                    <VSet>
                                        <VCheckbox v-if="isHttpMethods"
                                                   v-model="currentEntityItem.HttpMethods.IsFindOrCreate"
                                                   :disabled="currentEntityItem.Id > 0">
                                            <VText>FindOrCreate</VText>
                                        </VCheckbox>
                                        <VCheckbox v-if="isHttpMethods"
                                                   v-model="currentEntityItem.HttpMethods.IsUpdateOrCreate"
                                                   :disabled="currentEntityItem.Id > 0">
                                            <VText>IsUpdateOrCreate</VText>
                                        </VCheckbox>
                                    </VSet>
                                </template>

                                <template v-if="! currentEntityItem.IsFilter">
                                    <VSign>Backend structures</VSign>
                                    <VSet>
                                        <VCheckbox v-if="currentEntityItem && currentEntityItem.Structures"
                                                   v-model="currentEntityItem.Structures.WithoutDbModel"
                                                   :disabled="currentEntityItem.Id > 0">
                                            <VText>Without Db Model</VText>
                                        </VCheckbox>
                                    </VSet>
                                </template>

                            </VSet>

                            <VSet vertical indent-size="XS">
                                <VSet indent-size="XS" direction="vertical" v-for="field in currentEntityItem.Fields"
                                      :key="'id-'+field.Name">
                                    <VSet>
                                        <template>
                                            <VSign width="col3">{{field.Type}}</VSign>
                                            <VText>{{field.Name}}</VText>
                                        </template>
                                    </VSet>
                                </VSet>
                            </VSet>

                            <VHead level="h3" style="margin: 20px 0;" v-if="currentEntityItem.IsFilter">Add new filter
                            </VHead>
                            <VHead level="h3" style="margin: 20px 0;" v-else>Add new field</VHead>
                            <VSet vertical width="dyn">
                                <VSet v-if="isPanelCreate" width="dyn" :key="'newf-id'">
                                    <VSet width="fit">
                                        <VText><strong>Id</strong></VText>
                                    </VSet>
                                    <VSet width="dyn">
                                        <VSign v-if="IsUuidMode">UUID</VSign>
                                        <VSign v-else>Integer</VSign>

                                        <VCheckbox v-model="IsUuidMode">
                                            <VText>Uuid as primary key</VText>
                                        </VCheckbox>
                                    </VSet>

                                    <VSet width="dyn">

                                      <VSign v-if="IsViewMode">With view mode</VSign>
                                      <VSign v-else>Without view</VSign>

                                        <VCheckbox v-model="IsViewMode">
                                            <VText>ViewMode</VText>
                                        </VCheckbox>
                                    </VSet>

                                </VSet>
                                <VSet v-for="(f, index) in newFields" width="dyn" :key="'newf-' + index">
                                    <VInput v-model="f.Name" @input="updateNewFieldsList"></VInput>
                                    <VSelect v-model="f.Type" :items="getTypes"></VSelect>
                                </VSet>
                            </VSet>

                        </VSet>
                    </template>

                </template>

                <template #footer>

                    <VSet vertical>
                        <VSet vertical>
                            <VGroup v-for="(e, index) in errors" color="attention-light" :key="'err-' + index"
                                    width="dyn">{{e}}
                            </VGroup>
                        </VSet>

                        <VSet v-if="panel.type !== panel.request">
                            <VCheckbox v-model="IsRegenerateJsTypes">
                                <VText>Regenerate jstypes</VText>
                            </VCheckbox>
                        </VSet>

                        <VSet>
                            <VButton
                                    @click="saveChangesSubmit"
                                    accent
                                    :text="panelSubmitButtonText"
                            />
                            <VButton
                                    @click="cancelChanges"
                                    text="Cancel"
                            />
                        </VSet>
                    </VSet>

                </template>
            </VPanel>

            <slot name="confirmationPanel">
                <VPanel
                        v-if="currentEntityItem.showDeleteConfirmation"
                        modal
                        @close="closeConfirmationPanel"
                >
                    <template #content>
                        <VHead level="h3">Удалить элемент?</VHead>
                    </template>

                    <template #footer>
                        <VSet>
                            <VButton
                                    text="Подтвердить"
                                    accent
                                    @click="confirmDeleteHandler"
                            />
                            <VButton
                                    text="Отмена"
                                    @click="closeConfirmationPanel"
                            />
                        </VSet>
                    </template>
                </VPanel>
            </slot>
        </template>

        <template #footer>
            <VSet>
                <VButton
                        text="Добавить"
                        accent
                        @click="showCreateForm"
                />
            </VSet>
        </template>
    </WorkSpace>


</template>

<script>

    import EntityGen from "../../../webapp/jstypes/components/EntityGen";
    import VBadge from "swtui/src/components/VBadge";
    import VSet from "swtui/src/components/VSet";
    import VSpoiler from "swtui/src/components/VSpoiler";
    import VGroup from "swtui/src/components/VGroup";
    import VSelect from "swtui/src/components/VSelect";
    import VSign from "swtui/src/components/VSign";
    import EntityItem from "../components/EntityItem";
    import {mapGetters, mapActions} from 'vuex';
    import {Entity, EntityField, EntityFilter, FieldTypeFilter} from "../../../webapp/jstypes/apiModel";
    import EntityRequest from "../components/EntityRequest";

    function appendParams(u, params) {
        let uparams = "";
        switch (typeof params) {
            case "object":
                if (Object.keys(params).length < 1) {
                    return u;
                }
                u = u + (u.includes("?") ? "" : "?");
                for (const f of Object.keys(params)) {

                    if (uparams !== "") {
                        uparams += "&";
                    }

                    let list = "";

                    switch (typeof params[f]) {
                        case "object":

                            if (params && params[f]) {
                                for (let j = 0; j < params[f].length; j++) {
                                    if (list.length) {
                                        list += "&";
                                    }
                                    list += f + "[]=" + encodeURIComponent(params[f][j]);
                                }

                                uparams += list;
                            }
                            break;
                        default:
                            uparams += f + "=" + encodeURIComponent(params[f]);
                            break;
                    }
                }
                break;
        }
        return (u + uparams).replace(/[&]+/, '&');
    }

    export default {
        name: "Entity",
        components: {EntityRequest, EntityItem, VSpoiler, VBadge, EntityGen, VGroup, VSelect, VSign, VSet},
        mixins: [
            EntityGen,
        ],
        data() {
            return {
                searchModel: "",
                entityFilter: new EntityFilter(),
                currentEntityItem: this.getNewEntity(),
                newFields: [
                    new EntityField(),
                ],
                fieldTypeFilter: new FieldTypeFilter(),
                errors: [],
                isLoading: true,
                IsRegenerateJsTypes: localStorage.getItem("IsRegenerateJsTypes") === 'true',
                IsUuidMode: localStorage.getItem("IsUuidMode") === 'true',
                IsViewMode: localStorage.getItem("IsViewMode") === 'true',
                parts: 1,
            };
        },
        created: function () {

            // this.panel.request = "request";

            this.fieldTypeFilter.CurrentPage = 1;
            this.fieldTypeFilter.PerPage = 100;

            this.findFieldType({
                filter: this.fieldTypeFilter,
            }).then(() => {
                this.isLoading = false;
            });

        },
        computed: {
            ...mapGetters('gosha', {
                panelMaxWidth: "getPanelMaxWidth",
            }),

            partialList() {
                if (this.entityList && this.entityList.length > 15 * this.parts) {
                    return this.entityList.slice(0, 15 * this.parts)
                }
                return this.entityList
            },
            isHttpMethods() {
                return this.currentEntityItem && this.currentEntityItem.HttpMethods;
            },

            hasFields(entityItem) {
                return (entityItem) => {
                    return entityItem &&
                        entityItem.TypeFields &&
                        entityItem.TypeFields.length;
                }
            },
            getTypes() {
                return this.getListFieldType().filter((item) => {
                    return (this.currentEntityItem.IsFilter && item.Type === "filter") ||
                        (!this.currentEntityItem.IsFilter && item.Type !== "filter")
                }).map((item) => {
                    return item.Name
                });
            },
            entity() {
                return this.getEntityById();
            },
            isPanelRequest() {
                return this.panel.type === this.panel.request;
            },
            panelSubmitButtonText() {

                if (this.isPanelCreate) {
                    return this.panelSubmitButtonTextCreate;
                }

                if (this.isPanelEdit) {
                    return this.panelSubmitButtonTextEdit;
                }

                if (this.isPanelRequest) {
                    return 'Request';
                }

                return '';
            },
        },
        methods: {
            ...mapActions('gosha', [
                "findFieldType",
                "createEntity",
                "setResponse",
                "setResponseCode",
                "setResponseStatusText",
                "resetResponse",
            ]),
            ...mapGetters('gosha', [
                "getListFieldType",
                "getEntityById",
                "getPanelMaxWidth",
                "getFilters",
                "getHeaders",
                "getBodyModel",
                "getRequestUrl",
            ]),

            onScrolledBottom(){
                this.parts = this.parts + 1
            },

            onRequest(item) {
                console.log("request action", item);
            },

            getNewEntity() {

                let ne = new Entity();

                ne.IsFind = true;
                ne.IsCreate = true;
                ne.IsRead = true;
                ne.IsUpdate = true;
                ne.IsDelete = true;
                ne.IsFindOrCreate = false;
                ne.IsUpdateOrCreate = false;
                ne.Action = "find";

                ne.HttpMethods = {
                    IsFind: true,
                    IsCreate: true,
                    IsRead: true,
                    IsUpdate: true,
                    IsDelete: true,
                    IsFindOrCreate: true,
                    IsUpdateOrCreate: true,
                };

                ne.Structures = {
                    WithoutDbModel: false,
                };

                return ne;
            },
            search() {

                this.entityFilter.Search = this.searchModel;

                this.findEntity({
                    filter: this.entityFilter
                });
                this.parts = 1

            },
            editItem(item) {

                this.currentEntityItem = item;
                this.showPanel(this.panel.edit)
            },
            requestItem(item) {
                this.currentEntityItem = {};
                this.$nextTick(() => {
                    this.currentEntityItem = item;
                    this.showPanel(this.panel.request)
                })
            },
            clearNewFields() {
                this.newFields = [
                    new EntityField(),
                ];
            },
            updateNewFieldsList() {

                for (let i = this.newFields.length - 1; i >= 0; i--) {
                    if (this.newFields[i].Name.trim().length < 1 && this.newFields.length > 1) {
                        this.newFields.splice(i, 1);
                    }
                }

                if (this.newFields[this.newFields.length - 1].Name.trim().length > 0) {
                    this.newFields.push(new EntityField());
                }
            },
            /**
             * @return {boolean}
             */
            IsValidEntity() {

                this.errors = [];

                if (this.currentEntityItem.Name.trim().length < 2) {
                    this.errors.push('Invalid name length');
                }

                if (this.newFields.length < 2) {
                    this.errors.push('Invalid new fields count');
                }

                for (let i in this.newFields) {

                    if (i * 1 === this.newFields.length - 1) {
                        continue;
                    }

                    if (this.getTypes.indexOf(this.newFields[i].Type) === -1) {
                        this.errors.push('Invalid type this.newFields[i].Type');
                    }

                    if (this.newFields[i].Name.trim().length < 1) {
                        this.errors.push('Invalid type this.newFields[i].Name');
                    }
                }

                return this.errors.length === 0;
            },

            getClearFilters() {

                let f = JSON.parse(JSON.stringify(this.getFilters()));

                if (f.Ids && f.Ids.length > 0) {

                    let idsStr = f.Ids.replace("[", "").replace("]", "");
                    let items = idsStr.split(",",);

                    f.Ids = items;
                    console.log('items', items);
                }

                return f;
            },

            async sendRequest() {

                console.log('send request');
                console.log("this.getFilters", JSON.stringify(this.getFilters()));
                console.log(JSON.stringify(this.getFilters()));
                console.log("body model", JSON.stringify(this.getBodyModel()));
                console.log(this.currentEntityItem.Action);

                this.resetResponse();

                let result = null;
                let t1 = this.microtime(true);

                let filters = this.getClearFilters();

                switch (this.currentEntityItem.Action) {
                    case "find":

                        result = fetch(appendParams(this.getRequestUrl(), filters), {
                            method: 'GET',
                            headers: {Token: localStorage.getItem("authToken")},
                        });

                        break;
                    case "read":

                        result = fetch(appendParams(this.getRequestUrl(), filters), {
                            method: 'GET',
                            headers: {Token: localStorage.getItem("authToken")},
                        });

                        break;
                    case "post":

                        result = fetch(appendParams(this.getRequestUrl(), filters), {
                            method: 'POST',
                            headers: {Token: localStorage.getItem("authToken")},
                            body: JSON.stringify(data),
                        });

                        break;
                }


                result.then((response) => {
                    console.log('!!!res', response);
                    this.setResponseCode(response.status);
                    this.setResponseStatusText(response.statusText);

                    return response.text();
                })
                    .then((r) => {

                        console.log('result', r);
                        console.log('time', (this.microtime(true) - t1).toFixed(3));
                        this.setResponse(r);
                    })
            },

            microtime(get_as_float) {
                let now = new Date().getTime() / 1000;
                let s = parseInt(now);
                return (get_as_float) ? now : (Math.round((now - s) * 1000) / 1000) + ' ' + s;
            },

            saveChangesSubmit() {

                if (this.isPanelRequest) {
                    this.sendRequest();
                    return;
                }

                if (!this.IsValidEntity()) {
                    return;
                }

                if (this.isPanelCreate) {
                    this.currentEntityItem.item.Name = this.currentEntityItem.Name;
                    this.currentEntityItem.item.Fields = this.newFields.slice(0, -1);
                    this.currentEntityItem.item.Structures = {
                        WithoutDbModel: this.currentEntityItem.Structures.WithoutDbModel || false,
                    };

                    this.currentEntityItem.item.HttpMethods = {
                        IsFind: this.currentEntityItem.HttpMethods.IsFind || false,
                        IsCreate: this.currentEntityItem.HttpMethods.IsCreate || false,
                        IsRead: this.currentEntityItem.HttpMethods.IsRead || false,
                        IsUpdate: this.currentEntityItem.HttpMethods.IsUpdate || false,
                        IsDelete: this.currentEntityItem.HttpMethods.IsDelete || false,
                        IsFindOrCreate: this.currentEntityItem.HttpMethods.IsFindOrCreate || false,
                        IsUpdateOrCreate: this.currentEntityItem.HttpMethods.IsUpdateOrCreate || false,
                    };

                    this.createEntityItemSend().then(() => {
                        this.fetchEntityData();
                        this.closePanel();
                        this.clearNewFields();
                    });
                    return;
                }

                if (this.isPanelEdit) {

                    let item = this.currentEntityItem;
                    item.Fields = this.newFields.slice(0, -1);

                    return this.updateEntity({
                        id: item.Name,
                        data: item,
                    }).then((response) => {

                        if (response.Model) {
                            this.updateEntityById(response.Model);
                            this.fetchEntityData();
                            this.closePanel();
                            this.clearNewFields();
                        } else {
                            console.error('Ошибка изменения записи: ', response.Error);
                        }

                    }).catch(error => {
                        console.error('Ошибка изменения записи: ', error);
                    });
                }
            },

            createEntityItemSend() {
                return this.createEntity({
                    filter: {
                        IsRegenerateJsTypes: this.IsRegenerateJsTypes,
                        IsUuidMode: this.IsUuidMode,
                        IsViewMode: this.IsViewMode,
                    },
                    data: this.currentEntityItem.item,
                }).then((response) => {

                    if (response.Model) {
                        this.addEntityItemToList(response.Model);
                        this.clearPanelEntityItem();
                    } else {
                        console.error('Ошибка создания записи: ', response.Error);
                    }

                }).catch(error => {
                    console.error('Ошибка создания записи: ', error);
                });
            },

            showPanel(type) {

                this.errors = [];

                if (type === this.panel.create) {
                    this.panel.type = this.panel.create;
                    this.clearPanelEntityItem();
                } else if (type === this.panel.edit) {
                    this.panel.type = this.panel.edit;
                    this.currentEntityItem.isSelected = true;
                } else if (type === this.panel.request) {
                    this.panel.type = this.panel.request;
                    this.currentEntityItem.isSelected = true;
                }

                this.panel.show = true;
            },
            showCreateForm() {
                this.currentEntityItem = this.getNewEntity();
                this.showPanel(this.panel.create)
            },

        },
        watch: {
            'entityFilter.WithFilter': function (newFilter) {
                this.findEntity({
                    filter: this.entityFilter
                })
                this.parts = 1

            },
            IsRegenerateJsTypes(newVal) {
                localStorage.setItem("IsRegenerateJsTypes", newVal);
            },
            IsUuidMode(newVal) {
                localStorage.setItem("IsUuidMode", newVal);
            },
            IsViewMode(newVal) {
                localStorage.setItem("IsViewMode", newVal);
            },
        },
    }
</script>

<style>
    #second .sw-panel__inner {
        left: -10px !important;
    }

    .pt14 {
        padding-top: 14px !important;
    }

</style>