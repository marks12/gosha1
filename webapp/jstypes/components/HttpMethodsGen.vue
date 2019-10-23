
	<template>
    <WorkSpace>
        <template #header>
            <slot name="pageHeader">
                <VHead level="h1">HttpMethods</VHead>
            </slot>
        </template>

        <template #content>
            <slot name="data">
                <table>
                    <thead>
                        <tr>
                            <th v-for="(header, index) in fields" :key="index">{{ header }}</th>
                        </tr>
                    </thead>
            
                    <tbody>
                        <tr
                            v-for="httpMethodsItem in httpMethodsList"
                            :key="httpMethodsItem.Id"
                            @click="selectHttpMethodsItem(httpMethodsItem)"
                            class="sw-table__row_can-select"
                            :class="{'sw-table__row_is-selected': httpMethodsItem.Id === currentHttpMethodsItem.item.Id}"
                        >
                            <td v-for="(value, key) in fields" :key="key + '-fields'">
                                <VCheckbox v-if="isCheckbox(httpMethodsItem[key])" :checked="httpMethodsItem[key]" disabled></VCheckbox>
                                <VText v-else>{{ httpMethodsItem[key] }}</VText>
                            </td>
                        </tr>
                    </tbody>
                </table>
            </slot>
            
            <slot name="panel">
                <VPanel
                    v-if="panel.show"
                    width="col3"
                    @close="closePanel"
                >
                    <template #header>
                        <VHead level="h3">{{ panelHeader }}</VHead>
                    </template>
        
                    <template #content>
                        <form @submit.prevent="saveChangesSubmit">
                            <VSet direction="vertical">
                                <VSet
                                    v-for="(filed, key) in editFields" :key="key + '-editFields'"
                                    vertical-align="center"
                                >
                                    <VLabel
                                        width="col4"
                                        :for="`currentHttpMethodsItem${key}`"
                                    >{{ filed }}</VLabel>
                                    <VInput
										v-if="isInput(currentHttpMethodsItem.item[key])"
                                        v-model="currentHttpMethodsItem.item[key]"
                                        width="dyn"
                                        :id="`currentHttpMethodsItem${key}`"
                                        @input="changeCurrentHttpMethodsItem"
                                    />
									<VCheckbox
										v-if="isCheckbox(currentHttpMethodsItem.item[key])"
                                        v-model="currentHttpMethodsItem.item[key]"
                                        :id="`currentHttpMethodsItem${key}`"
										@input="changeCurrentHttpMethodsItem"
									/>
									
                                </VSet>
                            </VSet>
                            <button type="submit" :disabled="!currentHttpMethodsItem.hasChange" hidden></button>
                        </form>
                    </template>

                    <template #footer>
                        <VSet>
                            <VButton
                                @click="saveChangesSubmit"
                                accent
                                :text="panelSubmitButtonText"
                                :disabled="!currentHttpMethodsItem.hasChange"
                            />
                            <VButton
                                @click="cancelChanges"
                                text="Отменить"
                            />
                        </VSet>
                    </template>
                </VPanel>
            </slot>

            <slot name="confirmationPanel">
                <VPanel
                    v-if="currentHttpMethodsItem.showDeleteConfirmation"
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
            <slot name="pageFooter">
                <VSet>
                    <VButton
                        v-if="canCreate"
                        text="Добавить"
                        accent
                        @click="showPanel(panel.create)"
                    />
                    <VButton
                        v-if="canDelete"
                        text="Удалить"
                        :disabled="!currentHttpMethodsItem.isSelected"
                        @click="deleteHttpMethodsItemHandler"
                    />
                </VSet>
            </slot>
        </template>
    </WorkSpace>
</template>

<script>
    import httpMethodsData from "../data/HttpMethodsData";
    import { HttpMethods } from '../apiModel';
    import { mapGetters, mapMutations, mapActions } from 'vuex';
    import WorkSpace from "swui/src/components/WorkSpace";
    import VHead from "swui/src/components/VHead";
    import VSet from "swui/src/components/VSet";
    import VLabel from "swui/src/components/VLabel";
    import VInput from "swui/src/components/VInput";
    import VCheckbox from "swui/src/components/VCheckbox";
    import VText from "swui/src/components/VText";
    import VPanel from "swui/src/components/VPanel";
    import VButton from "swui/src/components/VButton";
    import VIcon from "swui/src/components/VIcon";
    import VSign from "swui/src/components/VSign";
    import VSelectSimple from "swui/src/components/VSelectSimple";

    export default {
        name: 'HttpMethodsGen',

        components: {VSelectSimple, VSign, VIcon, VButton, VPanel, VText, VInput, VLabel, VSet, VHead, WorkSpace, VCheckbox},

        props: {
            fields: {
                type: Object,
                default() {
                    const httpMethodsItem = new HttpMethods();
                    const fieldsObj = {};

                    for (let prop in httpMethodsItem) {

                        if (httpMethodsItem.hasOwnProperty(prop)) {
                            fieldsObj[prop] = prop;
                        }

                    }

                    return fieldsObj;
                }
            },
            editFields: {
                type: Object,
                default() {
                    const httpMethodsItem = new HttpMethods();
                    const fieldsObj = {};

                    for (let prop in httpMethodsItem) {

                        if (httpMethodsItem.hasOwnProperty(prop)) {
                            fieldsObj[prop] = prop;
                        }

                    }

                    return fieldsObj;
                }
            },
            canDelete: {
                type: Boolean,
                default: true,
            },
            canCreate: {
                type: Boolean,
                default: true,
            },
        },

        data() {
            return httpMethodsData;
        },

        created() {
            this.fillHttpMethodsFilter();
            this.fetchHttpMethodsData();
        },

        computed: {
            ...mapGetters({
                httpMethodsList: 'getListHttpMethods'
            }),
            isPanelCreate() {
                return this.panel.type === this.panel.create;
            },
            isPanelEdit() {
                return this.panel.type === this.panel.edit;
            },
            panelHeader() {
                if (this.isPanelCreate) {
                    return this.panelHeaderCreate;
                }

                if (this.isPanelEdit) {
                    return this.panelHeaderEdit;
                }

                return  '';
            },
            panelSubmitButtonText() {
                if (this.isPanelCreate) {
                    return this.panelSubmitButtonTextCreate;
                }

                if (this.isPanelEdit) {
                    return this.panelSubmitButtonTextEdit;
                }

                return  '';
            },
            isCheckbox() {
                return data => {
                    return typeof data === "boolean";
                }
            },
            isInput() {
                return data => {
                    return ! this.isCheckbox(data);
                }
            },
        },

        methods: {
            ...mapActions([
                'findHttpMethods',
                'updateHttpMethods',
                'deleteHttpMethods',
                'createHttpMethods',
            ]),

            ...mapMutations([
                'addHttpMethodsItemToList',
                'deleteHttpMethodsFromList',
                'updateHttpMethodsById',
            ]),

            fillHttpMethodsFilter() {
                this.httpMethodsFilter.CurrentPage = 1;
                this.httpMethodsFilter.PerPage = 1000;
            },

            fetchHttpMethodsData() {
                return this.findHttpMethods({
                    filter: this.httpMethodsFilter
                });
            },

            /**
             *
             * @param type
             */
            showPanel(type) {
                if (type === this.panel.create) {
                    this.panel.type = this.panel.create;
                    this.clearPanelHttpMethodsItem();
                } else if (type === this.panel.edit) {
                    this.panel.type = this.panel.edit;
                    this.currentHttpMethodsItem.isSelected = true;
                }

                this.panel.show = true;
            },

            closePanel() {
                this.panel.show = false;
                this.currentHttpMethodsItem.isSelected = false;
                this.clearPanelHttpMethodsItem();
            },

            selectHttpMethodsItem(httpMethodsItem) {
                this.showPanel(this.panel.edit);
                this.currentHttpMethodsItem.isSelected = true;
                Object.assign(this.currentHttpMethodsItem.item, httpMethodsItem);
            },

            changeCurrentHttpMethodsItem() {
                this.currentHttpMethodsItem.hasChange = true;
            },

            cancelChanges() {
                this.clearPanelHttpMethodsItem();
                this.closePanel();
            },

            clearPanelHttpMethodsItem() {
                this.currentHttpMethodsItem.item = new HttpMethods();
                this.currentHttpMethodsItem.hasChange = false;
            },

            saveChangesSubmit() {
                if (this.isPanelCreate) {
                    this.createHttpMethodsItemSubmit();
                    return;
                }

                if (this.isPanelEdit) {
                    this.editHttpMethodsItemSubmit();
                }
            },

            createHttpMethodsItemSubmit() {
                return this.createHttpMethods({
					data: this.currentHttpMethodsItem.item,
                }).then((response) => {

                    if (response.Model) {
                        this.addHttpMethodsItemToList(response.Model);
                        this.clearPanelHttpMethodsItem();
                    } else {
                        console.error('Ошибка создания записи: ', response.Error);
                    }

                }).catch(error => {
                    console.error('Ошибка создания записи: ', error);
                });
            },

            editHttpMethodsItemSubmit() {

                if (this.currentHttpMethodsItem.hasChange) {
                    return this.updateHttpMethods({
                        id: this.currentHttpMethodsItem.item.Id,
                        data: this.currentHttpMethodsItem.item,
                    }).then((response) => {

                        if (response.Model) {
                            this.updateHttpMethodsById(response.Model);
                            this.currentHttpMethodsItem.hasChange = false;
                            this.clearPanelHttpMethodsItem();
                            this.closePanel();
                        } else {
                            console.error('Ошибка изменения записи: ', response.Error);
                        }

                    }).catch(error => {
                        console.error('Ошибка изменения записи: ', error);
                    });

                } else {
					return new Promise(function(resolve, reject) {reject("Item has no changes. Nothing to save");})
				}
            },

            deleteHttpMethodsItemHandler() {
                let deletedItemId = this.currentHttpMethodsItem.item.Id;

                if (!this.currentHttpMethodsItem.canDelete) {
                    this.currentHttpMethodsItem.showDeleteConfirmation = true;
                    return;
                }

                this.deleteHttpMethods({
                    id: deletedItemId
                }).then(response => {

                    if (response.IsSuccess) {
                        this.deleteHttpMethodsFromList(deletedItemId);
                        this.clearPanelHttpMethodsItem();
                        this.currentHttpMethodsItem.canDelete = false;
                        this.currentHttpMethodsItem.isSelected = false;
                        this.panel.show = false;
                    } else {
                        console.error('Ошибка удаления элемента: ', response.Error);
                    }

                }).catch(error => {
                    console.error('Ошибка удаления элемента: ', error);
                });
            },

            confirmDeleteHandler() {
                this.currentHttpMethodsItem.showDeleteConfirmation = false;
                this.currentHttpMethodsItem.canDelete = true;
                this.deleteHttpMethodsItemHandler();
            },

            closeConfirmationPanel() {
                this.currentHttpMethodsItem.showDeleteConfirmation = false;
            },
        },
    }
</script>

<style lang="scss">

</style>
