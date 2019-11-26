
	<template>
    <WorkSpace>
        <template #header>
            <slot name="pageHeader">
                <VHead level="h1">PageType</VHead>
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
                            v-for="pageTypeItem in pageTypeList"
                            :key="pageTypeItem.Id"
                            @click="selectPageTypeItem(pageTypeItem)"
                            class="sw-table__row_can-select"
                            :class="{'sw-table__row_is-selected': pageTypeItem.Id === currentPageTypeItem.item.Id}"
                        >
                            <td v-for="(value, key) in fields" :key="key + '-fields'">
                                <VCheckbox v-if="isCheckbox(pageTypeItem[key])" :checked="pageTypeItem[key]" disabled></VCheckbox>
                                <VText v-else>{{ pageTypeItem[key] }}</VText>
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
                                        :for="`currentPageTypeItem${key}`"
                                    >{{ filed }}</VLabel>
                                    <VInput
										v-if="isInput(currentPageTypeItem.item[key])"
                                        v-model="currentPageTypeItem.item[key]"
                                        width="dyn"
                                        :id="`currentPageTypeItem${key}`"
                                        @input="changeCurrentPageTypeItem"
                                    />
									<VCheckbox
										v-if="isCheckbox(currentPageTypeItem.item[key])"
                                        v-model="currentPageTypeItem.item[key]"
                                        :id="`currentPageTypeItem${key}`"
										@input="changeCurrentPageTypeItem"
									/>
									
                                </VSet>
                            </VSet>
                            <button type="submit" :disabled="!currentPageTypeItem.hasChange" hidden></button>
                        </form>
                    </template>

                    <template #footer>
                        <VSet>
                            <VButton
                                @click="saveChangesSubmit"
                                accent
                                :text="panelSubmitButtonText"
                                :disabled="!currentPageTypeItem.hasChange"
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
                    v-if="currentPageTypeItem.showDeleteConfirmation"
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
                        :disabled="!currentPageTypeItem.isSelected"
                        @click="deletePageTypeItemHandler"
                    />
                </VSet>
            </slot>
        </template>
    </WorkSpace>
</template>

<script>
    import pageTypeData from "../data/PageTypeData";
    import { PageType } from '../apiModel';
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
        name: 'PageTypeGen',

        components: {VSelectSimple, VSign, VIcon, VButton, VPanel, VText, VInput, VLabel, VSet, VHead, WorkSpace, VCheckbox},

        props: {
            fields: {
                type: Object,
                default() {
                    const pageTypeItem = new PageType();
                    const fieldsObj = {};

                    for (let prop in pageTypeItem) {

                        if (pageTypeItem.hasOwnProperty(prop)) {
                            fieldsObj[prop] = prop;
                        }

                    }

                    return fieldsObj;
                }
            },
            editFields: {
                type: Object,
                default() {
                    const pageTypeItem = new PageType();
                    const fieldsObj = {};

                    for (let prop in pageTypeItem) {

                        if (pageTypeItem.hasOwnProperty(prop)) {
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
            return pageTypeData;
        },

        created() {
            this.fillPageTypeFilter();
            this.fetchPageTypeData();
        },

        computed: {
            ...mapGetters({
                pageTypeList: 'getListPageType'
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
                'findPageType',
                'updatePageType',
                'deletePageType',
                'createPageType',
            ]),

            ...mapMutations([
                'addPageTypeItemToList',
                'deletePageTypeFromList',
                'updatePageTypeById',
            ]),

            fillPageTypeFilter() {
                this.pageTypeFilter.CurrentPage = 1;
                this.pageTypeFilter.PerPage = 1000;
            },

            fetchPageTypeData() {
                return this.findPageType({
                    filter: this.pageTypeFilter
                });
            },

            /**
             *
             * @param type
             */
            showPanel(type) {
                if (type === this.panel.create) {
                    this.panel.type = this.panel.create;
                    this.clearPanelPageTypeItem();
                } else if (type === this.panel.edit) {
                    this.panel.type = this.panel.edit;
                    this.currentPageTypeItem.isSelected = true;
                }

                this.panel.show = true;
            },

            closePanel() {
                this.panel.show = false;
                this.currentPageTypeItem.isSelected = false;
                this.clearPanelPageTypeItem();
            },

            selectPageTypeItem(pageTypeItem) {
                this.showPanel(this.panel.edit);
                this.currentPageTypeItem.isSelected = true;
                Object.assign(this.currentPageTypeItem.item, pageTypeItem);
            },

            changeCurrentPageTypeItem() {
                this.currentPageTypeItem.hasChange = true;
            },

            cancelChanges() {
                this.clearPanelPageTypeItem();
                this.closePanel();
            },

            clearPanelPageTypeItem() {
                this.currentPageTypeItem.item = new PageType();
                this.currentPageTypeItem.hasChange = false;
            },

            saveChangesSubmit() {
                if (this.isPanelCreate) {
                    this.createPageTypeItemSubmit();
                    return;
                }

                if (this.isPanelEdit) {
                    this.editPageTypeItemSubmit();
                }
            },

            createPageTypeItemSubmit() {
                return this.createPageType({
					data: this.currentPageTypeItem.item,
                }).then((response) => {

                    if (response.Model) {
                        this.addPageTypeItemToList(response.Model);
                        this.clearPanelPageTypeItem();
                    } else {
                        console.error('Ошибка создания записи: ', response.Error);
                    }

                }).catch(error => {
                    console.error('Ошибка создания записи: ', error);
                });
            },

            editPageTypeItemSubmit() {

                if (this.currentPageTypeItem.hasChange) {
                    return this.updatePageType({
                        id: this.currentPageTypeItem.item.Id,
                        data: this.currentPageTypeItem.item,
                    }).then((response) => {

                        if (response.Model) {
                            this.updatePageTypeById(response.Model);
                            this.currentPageTypeItem.hasChange = false;
                            this.clearPanelPageTypeItem();
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

            deletePageTypeItemHandler() {
                let deletedItemId = this.currentPageTypeItem.item.Id;

                if (!this.currentPageTypeItem.canDelete) {
                    this.currentPageTypeItem.showDeleteConfirmation = true;
                    return;
                }

                this.deletePageType({
                    id: deletedItemId
                }).then(response => {

                    if (response.IsSuccess) {
                        this.deletePageTypeFromList(deletedItemId);
                        this.clearPanelPageTypeItem();
                        this.currentPageTypeItem.canDelete = false;
                        this.currentPageTypeItem.isSelected = false;
                        this.panel.show = false;
                    } else {
                        console.error('Ошибка удаления элемента: ', response.Error);
                    }

                }).catch(error => {
                    console.error('Ошибка удаления элемента: ', error);
                });
            },

            confirmDeleteHandler() {
                this.currentPageTypeItem.showDeleteConfirmation = false;
                this.currentPageTypeItem.canDelete = true;
                this.deletePageTypeItemHandler();
            },

            closeConfirmationPanel() {
                this.currentPageTypeItem.showDeleteConfirmation = false;
            },
        },
    }
</script>

<style lang="scss">

</style>
