
	<template>
    <WorkSpace>
        <template #header>
            <slot name="pageHeader">
                <VHead level="h1">PageTypeFilter</VHead>
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
                            v-for="pageTypeFilterItem in pageTypeFilterList"
                            :key="pageTypeFilterItem.Id"
                            @click="selectPageTypeFilterItem(pageTypeFilterItem)"
                            class="sw-table__row_can-select"
                            :class="{'sw-table__row_is-selected': pageTypeFilterItem.Id === currentPageTypeFilterItem.item.Id}"
                        >
                            <td v-for="(value, key) in fields" :key="key + '-fields'">
                                <VCheckbox v-if="isCheckbox(pageTypeFilterItem[key])" :checked="pageTypeFilterItem[key]" disabled></VCheckbox>
                                <VText v-else>{{ pageTypeFilterItem[key] }}</VText>
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
                                        :for="`currentPageTypeFilterItem${key}`"
                                    >{{ filed }}</VLabel>
                                    <VInput
										v-if="isInput(currentPageTypeFilterItem.item[key])"
                                        v-model="currentPageTypeFilterItem.item[key]"
                                        width="dyn"
                                        :id="`currentPageTypeFilterItem${key}`"
                                        @input="changeCurrentPageTypeFilterItem"
                                    />
									<VCheckbox
										v-if="isCheckbox(currentPageTypeFilterItem.item[key])"
                                        v-model="currentPageTypeFilterItem.item[key]"
                                        :id="`currentPageTypeFilterItem${key}`"
										@input="changeCurrentPageTypeFilterItem"
									/>
									
                                </VSet>
                            </VSet>
                            <button type="submit" :disabled="!currentPageTypeFilterItem.hasChange" hidden></button>
                        </form>
                    </template>

                    <template #footer>
                        <VSet>
                            <VButton
                                @click="saveChangesSubmit"
                                accent
                                :text="panelSubmitButtonText"
                                :disabled="!currentPageTypeFilterItem.hasChange"
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
                    v-if="currentPageTypeFilterItem.showDeleteConfirmation"
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
                        :disabled="!currentPageTypeFilterItem.isSelected"
                        @click="deletePageTypeFilterItemHandler"
                    />
                </VSet>
            </slot>
        </template>
    </WorkSpace>
</template>

<script>
    import pageTypeFilterData from "../data/PageTypeFilterData";
    import { PageTypeFilter } from '../apiModel';
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
        name: 'PageTypeFilterGen',

        components: {VSelectSimple, VSign, VIcon, VButton, VPanel, VText, VInput, VLabel, VSet, VHead, WorkSpace, VCheckbox},

        props: {
            fields: {
                type: Object,
                default() {
                    const pageTypeFilterItem = new PageTypeFilter();
                    const fieldsObj = {};

                    for (let prop in pageTypeFilterItem) {

                        if (pageTypeFilterItem.hasOwnProperty(prop)) {
                            fieldsObj[prop] = prop;
                        }

                    }

                    return fieldsObj;
                }
            },
            editFields: {
                type: Object,
                default() {
                    const pageTypeFilterItem = new PageTypeFilter();
                    const fieldsObj = {};

                    for (let prop in pageTypeFilterItem) {

                        if (pageTypeFilterItem.hasOwnProperty(prop)) {
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
            return pageTypeFilterData;
        },

        created() {
            this.fillPageTypeFilterFilter();
            this.fetchPageTypeFilterData();
        },

        computed: {
            ...mapGetters({
                pageTypeFilterList: 'getListPageTypeFilter'
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
                'findPageTypeFilter',
                'updatePageTypeFilter',
                'deletePageTypeFilter',
                'createPageTypeFilter',
            ]),

            ...mapMutations([
                'addPageTypeFilterItemToList',
                'deletePageTypeFilterFromList',
                'updatePageTypeFilterById',
            ]),

            fillPageTypeFilterFilter() {
                this.pageTypeFilterFilter.CurrentPage = 1;
                this.pageTypeFilterFilter.PerPage = 1000;
            },

            fetchPageTypeFilterData() {
                return this.findPageTypeFilter({
                    filter: this.pageTypeFilterFilter
                });
            },

            /**
             *
             * @param type
             */
            showPanel(type) {
                if (type === this.panel.create) {
                    this.panel.type = this.panel.create;
                    this.clearPanelPageTypeFilterItem();
                } else if (type === this.panel.edit) {
                    this.panel.type = this.panel.edit;
                    this.currentPageTypeFilterItem.isSelected = true;
                }

                this.panel.show = true;
            },

            closePanel() {
                this.panel.show = false;
                this.currentPageTypeFilterItem.isSelected = false;
                this.clearPanelPageTypeFilterItem();
            },

            selectPageTypeFilterItem(pageTypeFilterItem) {
                this.showPanel(this.panel.edit);
                this.currentPageTypeFilterItem.isSelected = true;
                Object.assign(this.currentPageTypeFilterItem.item, pageTypeFilterItem);
            },

            changeCurrentPageTypeFilterItem() {
                this.currentPageTypeFilterItem.hasChange = true;
            },

            cancelChanges() {
                this.clearPanelPageTypeFilterItem();
                this.closePanel();
            },

            clearPanelPageTypeFilterItem() {
                this.currentPageTypeFilterItem.item = new PageTypeFilter();
                this.currentPageTypeFilterItem.hasChange = false;
            },

            saveChangesSubmit() {
                if (this.isPanelCreate) {
                    this.createPageTypeFilterItemSubmit();
                    return;
                }

                if (this.isPanelEdit) {
                    this.editPageTypeFilterItemSubmit();
                }
            },

            createPageTypeFilterItemSubmit() {
                return this.createPageTypeFilter({
					data: this.currentPageTypeFilterItem.item,
                }).then((response) => {

                    if (response.Model) {
                        this.addPageTypeFilterItemToList(response.Model);
                        this.clearPanelPageTypeFilterItem();
                    } else {
                        console.error('Ошибка создания записи: ', response.Error);
                    }

                }).catch(error => {
                    console.error('Ошибка создания записи: ', error);
                });
            },

            editPageTypeFilterItemSubmit() {

                if (this.currentPageTypeFilterItem.hasChange) {
                    return this.updatePageTypeFilter({
                        id: this.currentPageTypeFilterItem.item.Id,
                        data: this.currentPageTypeFilterItem.item,
                    }).then((response) => {

                        if (response.Model) {
                            this.updatePageTypeFilterById(response.Model);
                            this.currentPageTypeFilterItem.hasChange = false;
                            this.clearPanelPageTypeFilterItem();
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

            deletePageTypeFilterItemHandler() {
                let deletedItemId = this.currentPageTypeFilterItem.item.Id;

                if (!this.currentPageTypeFilterItem.canDelete) {
                    this.currentPageTypeFilterItem.showDeleteConfirmation = true;
                    return;
                }

                this.deletePageTypeFilter({
                    id: deletedItemId
                }).then(response => {

                    if (response.IsSuccess) {
                        this.deletePageTypeFilterFromList(deletedItemId);
                        this.clearPanelPageTypeFilterItem();
                        this.currentPageTypeFilterItem.canDelete = false;
                        this.currentPageTypeFilterItem.isSelected = false;
                        this.panel.show = false;
                    } else {
                        console.error('Ошибка удаления элемента: ', response.Error);
                    }

                }).catch(error => {
                    console.error('Ошибка удаления элемента: ', error);
                });
            },

            confirmDeleteHandler() {
                this.currentPageTypeFilterItem.showDeleteConfirmation = false;
                this.currentPageTypeFilterItem.canDelete = true;
                this.deletePageTypeFilterItemHandler();
            },

            closeConfirmationPanel() {
                this.currentPageTypeFilterItem.showDeleteConfirmation = false;
            },
        },
    }
</script>

<style lang="scss">

</style>
