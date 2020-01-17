
	<template>
    <WorkSpace>
        <template #header>
            <slot name="pageHeader">
                <VHead level="h1">PageContentFilter</VHead>
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
                            v-for="pageContentFilterItem in pageContentFilterList"
                            :key="pageContentFilterItem.Id"
                            @click="selectPageContentFilterItem(pageContentFilterItem)"
                            class="sw-table__row_can-select"
                            :class="{'sw-table__row_is-selected': pageContentFilterItem.Id === currentPageContentFilterItem.item.Id}"
                        >
                            <td v-for="(value, key) in fields" :key="key + '-fields'">
                                <VCheckbox v-if="isCheckbox(pageContentFilterItem[key])" :checked="pageContentFilterItem[key]" disabled></VCheckbox>
                                <VText v-else>{{ pageContentFilterItem[key] }}</VText>
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
                                        :for="`currentPageContentFilterItem${key}`"
                                    >{{ filed }}</VLabel>
                                    <VInput
										v-if="isInput(currentPageContentFilterItem.item[key])"
                                        v-model="currentPageContentFilterItem.item[key]"
                                        width="dyn"
                                        :id="`currentPageContentFilterItem${key}`"
                                        @input="changeCurrentPageContentFilterItem"
                                    />
									<VCheckbox
										v-if="isCheckbox(currentPageContentFilterItem.item[key])"
                                        v-model="currentPageContentFilterItem.item[key]"
                                        :id="`currentPageContentFilterItem${key}`"
										@input="changeCurrentPageContentFilterItem"
									/>
									
                                </VSet>
                            </VSet>
                            <button type="submit" :disabled="!currentPageContentFilterItem.hasChange" hidden></button>
                        </form>
                    </template>

                    <template #footer>
                        <VSet>
                            <VButton
                                @click="saveChangesSubmit"
                                accent
                                :text="panelSubmitButtonText"
                                :disabled="!currentPageContentFilterItem.hasChange"
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
                    v-if="currentPageContentFilterItem.showDeleteConfirmation"
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
                        :disabled="!currentPageContentFilterItem.isSelected"
                        @click="deletePageContentFilterItemHandler"
                    />
                </VSet>
            </slot>
        </template>
    </WorkSpace>
</template>

<script>
    import pageContentFilterData from "../data/PageContentFilterData";
    import { PageContentFilter } from '../apiModel';
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

    export default {
        name: 'PageContentFilterGen',

        components: {VButton, VPanel, VText, VInput, VLabel, VSet, VHead, WorkSpace, VCheckbox},

        props: {
            fields: {
                type: Object,
                default() {
                    const pageContentFilterItem = new PageContentFilter();
                    const fieldsObj = {};

                    for (let prop in pageContentFilterItem) {

                        if (pageContentFilterItem.hasOwnProperty(prop)) {
                            fieldsObj[prop] = prop;
                        }

                    }

                    return fieldsObj;
                }
            },
            editFields: {
                type: Object,
                default() {
                    const pageContentFilterItem = new PageContentFilter();
                    const fieldsObj = {};

                    for (let prop in pageContentFilterItem) {

                        if (pageContentFilterItem.hasOwnProperty(prop)) {
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
            return pageContentFilterData;
        },

        created() {
			this.onCreated();
        },

        computed: {
            ...mapGetters({
                pageContentFilterList: 'getListPageContentFilter'
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
                'findPageContentFilter',
                'updatePageContentFilter',
                'deletePageContentFilter',
                'createPageContentFilter',
            ]),

            ...mapMutations([
                'addPageContentFilterItemToList',
                'deletePageContentFilterFromList',
                'updatePageContentFilterById',
            ]),

			onCreated() {
				this.fillPageContentFilterFilter();
	            this.fetchPageContentFilterData();
			},

            fillPageContentFilterFilter() {
                this.pageContentFilterFilter.CurrentPage = 1;
                this.pageContentFilterFilter.PerPage = 1000;
            },

            fetchPageContentFilterData() {
                return this.findPageContentFilter({
                    filter: this.pageContentFilterFilter
                });
            },

            /**
             *
             * @param type
             */
            showPanel(type) {
                if (type === this.panel.create) {
                    this.panel.type = this.panel.create;
                    this.clearPanelPageContentFilterItem();
                } else if (type === this.panel.edit) {
                    this.panel.type = this.panel.edit;
                    this.currentPageContentFilterItem.isSelected = true;
                }

                this.panel.show = true;
            },

            closePanel() {
                this.panel.show = false;
                this.currentPageContentFilterItem.isSelected = false;
                this.clearPanelPageContentFilterItem();
            },

            selectPageContentFilterItem(pageContentFilterItem) {
                this.showPanel(this.panel.edit);
                this.currentPageContentFilterItem.isSelected = true;
                Object.assign(this.currentPageContentFilterItem.item, pageContentFilterItem);
            },

            changeCurrentPageContentFilterItem() {
                this.currentPageContentFilterItem.hasChange = true;
            },

            cancelChanges() {
                this.clearPanelPageContentFilterItem();
                this.closePanel();
            },

            clearPanelPageContentFilterItem() {
                this.currentPageContentFilterItem.item = new PageContentFilter();
                this.currentPageContentFilterItem.hasChange = false;
            },

            saveChangesSubmit() {
                if (this.isPanelCreate) {
                    this.createPageContentFilterItemSubmit();
                    return;
                }

                if (this.isPanelEdit) {
                    this.editPageContentFilterItemSubmit();
                }
            },

            createPageContentFilterItemSubmit() {
                return this.createPageContentFilter({
					data: this.currentPageContentFilterItem.item,
                }).then((response) => {

                    if (response.Model) {
                        this.addPageContentFilterItemToList(response.Model);
                        this.clearPanelPageContentFilterItem();
                    } else {
                        console.error('Ошибка создания записи: ', response.Error);
                    }

                }).catch(error => {
                    console.error('Ошибка создания записи: ', error);
                });
            },

            editPageContentFilterItemSubmit() {

                if (this.currentPageContentFilterItem.hasChange) {
                    return this.updatePageContentFilter({
                        id: this.currentPageContentFilterItem.item.Id,
                        data: this.currentPageContentFilterItem.item,
                    }).then((response) => {

                        if (response.Model) {
                            this.updatePageContentFilterById(response.Model);
                            this.currentPageContentFilterItem.hasChange = false;
                            this.clearPanelPageContentFilterItem();
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

            deletePageContentFilterItemHandler() {
                let deletedItemId = this.currentPageContentFilterItem.item.Id;

                if (!this.currentPageContentFilterItem.canDelete) {
                    this.currentPageContentFilterItem.showDeleteConfirmation = true;
                    return;
                }

                this.deletePageContentFilter({
                    id: deletedItemId
                }).then(response => {

                    if (response.IsSuccess) {
                        this.deletePageContentFilterFromList(deletedItemId);
                        this.clearPanelPageContentFilterItem();
                        this.currentPageContentFilterItem.canDelete = false;
                        this.currentPageContentFilterItem.isSelected = false;
                        this.panel.show = false;
                    } else {
                        console.error('Ошибка удаления элемента: ', response.Error);
                    }

                }).catch(error => {
                    console.error('Ошибка удаления элемента: ', error);
                });
            },

            confirmDeleteHandler() {
                this.currentPageContentFilterItem.showDeleteConfirmation = false;
                this.currentPageContentFilterItem.canDelete = true;
                this.deletePageContentFilterItemHandler();
            },

            closeConfirmationPanel() {
                this.currentPageContentFilterItem.showDeleteConfirmation = false;
            },
        },
    }
</script>

<style lang="scss">

</style>
