<template>

    <WorkSpace footerPosition="bottom" noBg>
        <template #header>
            <VSet>
                <VSet vertical  indent-size="XS">
                    <VHead level="h1">Entity</VHead>
                    <VSet>
                        <VInput  placeholder="Поиск" v-model="searchModel" @input="search"></VInput>
                        <VCheckbox v-model="entityFilter.WithFilter">
                            <VText font-size="txs">Отображать фильтры</VText>
                        </VCheckbox>
                    </VSet>
                </VSet>
            </VSet>
        </template>

        <template slot="content">
            <template>
                <VSet v-if="entityList && entityList.length" wrap>
                    <template v-for="entityItem in entityList">
                        <EntityItem :entityItem="entityItem" :onEdit="editItem"></EntityItem>
                    </template>
                </VSet>
                <VText class="loading"></VText>
            </template>

            <template>
                <VPanel
                        v-if="panel.show"
                        width="col3"
                        @close="closePanel"
                >
                    <VSet slot="header">
                        <VHead level="h3" width="fit">{{ panelHeader }} {{ currentEntityItem.Name }}</VHead>
                    </VSet>

                    <template slot="content">
                        <VSet vertical @submit.prevent="saveChangesSubmit">

                            <VSet vertical indent-size="XS">
                                <VSign>Entity name</VSign>
                                <VInput v-model="currentEntityItem.Name" :disabled="currentEntityItem.Id > 0" width="dyn"></VInput>
                            </VSet>

                            <VSet vertical indent-size="XS">
                                <VSet  indent-size="XS" direction="vertical" v-for="field in currentEntityItem.Fields" :key="'id-'+field.Name">
                                    <VSet>
                                        <template>
                                            <VSign width="col3">{{field.Type}}</VSign>
                                            <VText>{{field.Name}}</VText>
                                        </template>
                                    </VSet>
                                </VSet>
                            </VSet>

                            <VHead level="h3" style="margin: 20px 0;">Add new field</VHead>
                            <VSet vertical>
                                <VSet v-if="isPanelCreate" width="dyn" :key="'newf-id'">
                                    <VInput width="dyn" value="Id" disabled></VInput>
                                    <VInput value="int" disabled></VInput>
                                </VSet>
                                <VSet v-for="(f, index) in newFields" width="dyn" :key="'newf-' + index">
                                    <VInput v-model="f.Name" @input="updateNewFieldsList" width="dyn"></VInput>
                                    <VSelect v-model="f.Type" :items="getTypes"></VSelect>
                                </VSet>
                            </VSet>

                        </VSet>
                    </template>

                    <template #footer>

                        <VSet vertical>
                            <VSet vertical>
                                <VGroup v-for="(e, index) in errors" color="attention-light" :key="'err-' + index" width="dyn">{{e}}</VGroup>
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
            </template>

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
    import VSpoiler from "swtui/src/components/VSpoiler";
    import VGroup from "swtui/src/components/VGroup";
    import VSelect from "swtui/src/components/VSelect";
    import EntityItem from "../components/EntityItem";
    import { mapGetters, mapMutations, mapActions } from 'vuex';
    import {Entity, EntityField, EntityFilter, FieldTypeFilter} from "../../../webapp/jstypes/apiModel";

    export default {
        name: "Entity",
        components: {EntityItem, VSpoiler, VBadge, EntityGen, VGroup, VSelect},
        mixins: [
            EntityGen,
        ],
        data() {
            return {
                searchModel: "",
                entityFilter: new EntityFilter(),
                currentEntityItem: new Entity(),
                newFields: [
                    new EntityField(),
                ],
                fieldTypeFilter: new FieldTypeFilter(),
                errors: [],
                isLoading: true,
            };
        },
        created: function () {

            this.fieldTypeFilter.CurrentPage = 1;
            this.fieldTypeFilter.PerPage = 100;

            this.findFieldType({
                filter: this.fieldTypeFilter,
            }).then(()=>{
                this.isLoading = false;
            });

        },
        methods: {
            ...mapActions([
                "findFieldType",
                "createEntity",
            ]),
            ...mapGetters([
                "getListFieldType",
            ]),
            search() {

                this.entityFilter.Search = this.searchModel;

                this.findEntity({
                    filter: this.entityFilter
                });
            },
            editItem(item) {

                this.currentEntityItem = item;
                this.showPanel(this.panel.edit)
            },
            clearNewFields() {
                this.newFields  = [
                    new EntityField(),
                ];
            },
            updateNewFieldsList() {

                for (let i = this.newFields.length - 1; i >= 0; i--) {
                    if (this.newFields[i].Name.trim().length < 1 && this.newFields.length > 1) {
                        this.newFields.splice(i , 1);
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
            saveChangesSubmit() {

                if (! this.IsValidEntity()) {
                    return;
                }

                if (this.isPanelCreate) {
                    this.currentEntityItem.item.Name = this.currentEntityItem.Name;
                    this.currentEntityItem.item.Fields = this.newFields.slice(0, -1);
                    this.createEntityItemSubmit().then(()=>{
                        this.fetchEntityData();
                        this.closePanel();
                        this.clearNewFields();
                    });
                    return;
                }

                if (this.isPanelEdit) {

                    let item = new Entity();

                    item.Name = this.currentEntityItem.Name;
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
            showPanel(type) {

                this.errors = [];

                if (type === this.panel.create) {
                    this.panel.type = this.panel.create;
                    this.clearPanelEntityItem();
                } else if (type === this.panel.edit) {
                    this.panel.type = this.panel.edit;
                    this.currentEntityItem.isSelected = true;
                }

                this.panel.show = true;
            },
            showCreateForm() {
                this.currentEntityItem = new Entity();
                this.showPanel(this.panel.create)
            },
        },
        computed: {
            hasFields(entityItem) {
                return (entityItem) => {
                    return  entityItem &&
                            entityItem.TypeFields &&
                            entityItem.TypeFields.length;
                }
            },
            getTypes() {
                return this.getListFieldType().map( (item)=> {return item.Name});
            },
        },
        watch: {
            'entityFilter.WithFilter': function (newFilter) {

                this.findEntity({
                    filter: this.entityFilter
                })
            },
        },
    }
</script>

<style scoped>

</style>