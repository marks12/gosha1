
	<template>
    <WorkSpace>
        <template #header>
            <slot name="pageHeader">
                <VHead level="h1">PageTemplateFilter</VHead>
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
                            v-for="pageTemplateFilterItem in pageTemplateFilterList"
                            :key="pageTemplateFilterItem.Id"
                            @click="selectPageTemplateFilterItem(pageTemplateFilterItem)"
                            class="sw-table__row_can-select"
                            :class="{'sw-table__row_is-selected': pageTemplateFilterItem.Id === currentPageTemplateFilterItem.item.Id}"
                        >
                            <td v-for="(value, key) in fields" :key="key + '-fields'">
                                <VCheckbox v-if="isCheckbox(pageTemplateFilterItem[key])" :checked="pageTemplateFilterItem[key]" disabled></VCheckbox>
                                <VText v-else>{{ pageTemplateFilterItem[key] }}</VText>
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
                                        :for="`currentPageTemplateFilterItem${key}`"
                                    >{{ filed }}</VLabel>
                                    <VInput
										v-if="isInput(currentPageTemplateFilterItem.item[key])"
                                        v-model="currentPageTemplateFilterItem.item[key]"
                                        width="dyn"
                                        :id="`currentPageTemplateFilterItem${key}`"
                                        @input="changeCurrentPageTemplateFilterItem"
                                    />
									<VCheckbox
										v-if="isCheckbox(currentPageTemplateFilterItem.item[key])"
                                        v-model="currentPageTemplateFilterItem.item[key]"
                                        :id="`currentPageTemplateFilterItem${key}`"
										@input="changeCurrentPageTemplateFilterItem"
									/>
									
                                </VSet>
                            </VSet>
                            <button type="submit" :disabled="!currentPageTemplateFilterItem.hasChange" hidden></button>
                        </form>
                    </template>

                    <template #footer>
                        <VSet>
                            <VButton
                                @click="saveChangesSubmit"
                                accent
                                :text="panelSubmitButtonText"
                                :disabled="!currentPageTemplateFilterItem.hasChange"
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
                    v-if="currentPageTemplateFilterItem.showDeleteConfirmation"
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
                        :disabled="!currentPageTemplateFilterItem.isSelected"
                        @click="deletePageTemplateFilterItemHandler"
                    />
                </VSet>
            </slot>
        </template>
    </WorkSpace>
</template>

<script>
    import pageTemplateFilterData from "../data/PageTemplateFilterData";
    import { PageTemplateFilter } from '../apiModel';
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
        name: 'PageTemplateFilterGen',

        components: {VSelectSimple, VSign, VIcon, VButton, VPanel, VText, VInput, VLabel, VSet, VHead, WorkSpace, VCheckbox},

        props: {
            fields: {
                type: Object,
                default() {
                    const pageTemplateFilterItem = new PageTemplateFilter();
                    const fieldsObj = {};

                    for (let prop in pageTemplateFilterItem) {

                        if (pageTemplateFilterItem.hasOwnProperty(prop)) {
                            fieldsObj[prop] = prop;
                        }

                    }

                    return fieldsObj;
                }
            },
            editFields: {
                type: Object,
                default() {
                    const pageTemplateFilterItem = new PageTemplateFilter();
                    const fieldsObj = {};

                    for (let prop in pageTemplateFilterItem) {

                        if (pageTemplateFilterItem.hasOwnProperty(prop)) {
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
            return pageTemplateFilterData;
        },

        created() {
            this.fillPageTemplateFilterFilter();
            this.fetchPageTemplateFilterData();
        },

        computed: {
            ...mapGetters({
                pageTemplateFilterList: 'getListPageTemplateFilter'
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
                'findPageTemplateFilter',
                'updatePageTemplateFilter',
                'deletePageTemplateFilter',
                'createPageTemplateFilter',
            ]),

            ...mapMutations([
                'addPageTemplateFilterItemToList',
                'deletePageTemplateFilterFromList',
                'updatePageTemplateFilterById',
            ]),

            fillPageTemplateFilterFilter() {
                this.pageTemplateFilterFilter.CurrentPage = 1;
                this.pageTemplateFilterFilter.PerPage = 1000;
            },

            fetchPageTemplateFilterData() {
                return this.findPageTemplateFilter({
                    filter: this.pageTemplateFilterFilter
                });
            },

            /**
             *
             * @param type
             */
            showPanel(type) {
                if (type === this.panel.create) {
                    this.panel.type = this.panel.create;
                    this.clearPanelPageTemplateFilterItem();
                } else if (type === this.panel.edit) {
                    this.panel.type = this.panel.edit;
                    this.currentPageTemplateFilterItem.isSelected = true;
                }

                this.panel.show = true;
            },

            closePanel() {
                this.panel.show = false;
                this.currentPageTemplateFilterItem.isSelected = false;
                this.clearPanelPageTemplateFilterItem();
            },

            selectPageTemplateFilterItem(pageTemplateFilterItem) {
                this.showPanel(this.panel.edit);
                this.currentPageTemplateFilterItem.isSelected = true;
                Object.assign(this.currentPageTemplateFilterItem.item, pageTemplateFilterItem);
            },

            changeCurrentPageTemplateFilterItem() {
                this.currentPageTemplateFilterItem.hasChange = true;
            },

            cancelChanges() {
                this.clearPanelPageTemplateFilterItem();
                this.closePanel();
            },

            clearPanelPageTemplateFilterItem() {
                this.currentPageTemplateFilterItem.item = new PageTemplateFilter();
                this.currentPageTemplateFilterItem.hasChange = false;
            },

            saveChangesSubmit() {
                if (this.isPanelCreate) {
                    this.createPageTemplateFilterItemSubmit();
                    return;
                }

                if (this.isPanelEdit) {
                    this.editPageTemplateFilterItemSubmit();
                }
            },

            createPageTemplateFilterItemSubmit() {
                return this.createPageTemplateFilter({
					data: this.currentPageTemplateFilterItem.item,
                }).then((response) => {

                    if (response.Model) {
                        this.addPageTemplateFilterItemToList(response.Model);
                        this.clearPanelPageTemplateFilterItem();
                    } else {
                        console.error('Ошибка создания записи: ', response.Error);
                    }

                }).catch(error => {
                    console.error('Ошибка создания записи: ', error);
                });
            },

            editPageTemplateFilterItemSubmit() {

                if (this.currentPageTemplateFilterItem.hasChange) {
                    return this.updatePageTemplateFilter({
                        id: this.currentPageTemplateFilterItem.item.Id,
                        data: this.currentPageTemplateFilterItem.item,
                    }).then((response) => {

                        if (response.Model) {
                            this.updatePageTemplateFilterById(response.Model);
                            this.currentPageTemplateFilterItem.hasChange = false;
                            this.clearPanelPageTemplateFilterItem();
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

            deletePageTemplateFilterItemHandler() {
                let deletedItemId = this.currentPageTemplateFilterItem.item.Id;

                if (!this.currentPageTemplateFilterItem.canDelete) {
                    this.currentPageTemplateFilterItem.showDeleteConfirmation = true;
                    return;
                }

                this.deletePageTemplateFilter({
                    id: deletedItemId
                }).then(response => {

                    if (response.IsSuccess) {
                        this.deletePageTemplateFilterFromList(deletedItemId);
                        this.clearPanelPageTemplateFilterItem();
                        this.currentPageTemplateFilterItem.canDelete = false;
                        this.currentPageTemplateFilterItem.isSelected = false;
                        this.panel.show = false;
                    } else {
                        console.error('Ошибка удаления элемента: ', response.Error);
                    }

                }).catch(error => {
                    console.error('Ошибка удаления элемента: ', error);
                });
            },

            confirmDeleteHandler() {
                this.currentPageTemplateFilterItem.showDeleteConfirmation = false;
                this.currentPageTemplateFilterItem.canDelete = true;
                this.deletePageTemplateFilterItemHandler();
            },

            closeConfirmationPanel() {
                this.currentPageTemplateFilterItem.showDeleteConfirmation = false;
            },
        },
    }
</script>

<style lang="scss">

</style>
