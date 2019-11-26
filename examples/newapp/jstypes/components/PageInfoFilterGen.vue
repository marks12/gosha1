
	<template>
    <WorkSpace>
        <template #header>
            <slot name="pageHeader">
                <VHead level="h1">PageInfoFilter</VHead>
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
                            v-for="pageInfoFilterItem in pageInfoFilterList"
                            :key="pageInfoFilterItem.Id"
                            @click="selectPageInfoFilterItem(pageInfoFilterItem)"
                            class="sw-table__row_can-select"
                            :class="{'sw-table__row_is-selected': pageInfoFilterItem.Id === currentPageInfoFilterItem.item.Id}"
                        >
                            <td v-for="(value, key) in fields" :key="key + '-fields'">
                                <VCheckbox v-if="isCheckbox(pageInfoFilterItem[key])" :checked="pageInfoFilterItem[key]" disabled></VCheckbox>
                                <VText v-else>{{ pageInfoFilterItem[key] }}</VText>
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
                                        :for="`currentPageInfoFilterItem${key}`"
                                    >{{ filed }}</VLabel>
                                    <VInput
										v-if="isInput(currentPageInfoFilterItem.item[key])"
                                        v-model="currentPageInfoFilterItem.item[key]"
                                        width="dyn"
                                        :id="`currentPageInfoFilterItem${key}`"
                                        @input="changeCurrentPageInfoFilterItem"
                                    />
									<VCheckbox
										v-if="isCheckbox(currentPageInfoFilterItem.item[key])"
                                        v-model="currentPageInfoFilterItem.item[key]"
                                        :id="`currentPageInfoFilterItem${key}`"
										@input="changeCurrentPageInfoFilterItem"
									/>
									
                                </VSet>
                            </VSet>
                            <button type="submit" :disabled="!currentPageInfoFilterItem.hasChange" hidden></button>
                        </form>
                    </template>

                    <template #footer>
                        <VSet>
                            <VButton
                                @click="saveChangesSubmit"
                                accent
                                :text="panelSubmitButtonText"
                                :disabled="!currentPageInfoFilterItem.hasChange"
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
                    v-if="currentPageInfoFilterItem.showDeleteConfirmation"
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
                        :disabled="!currentPageInfoFilterItem.isSelected"
                        @click="deletePageInfoFilterItemHandler"
                    />
                </VSet>
            </slot>
        </template>
    </WorkSpace>
</template>

<script>
    import pageInfoFilterData from "../data/PageInfoFilterData";
    import { PageInfoFilter } from '../apiModel';
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
        name: 'PageInfoFilterGen',

        components: {VSelectSimple, VSign, VIcon, VButton, VPanel, VText, VInput, VLabel, VSet, VHead, WorkSpace, VCheckbox},

        props: {
            fields: {
                type: Object,
                default() {
                    const pageInfoFilterItem = new PageInfoFilter();
                    const fieldsObj = {};

                    for (let prop in pageInfoFilterItem) {

                        if (pageInfoFilterItem.hasOwnProperty(prop)) {
                            fieldsObj[prop] = prop;
                        }

                    }

                    return fieldsObj;
                }
            },
            editFields: {
                type: Object,
                default() {
                    const pageInfoFilterItem = new PageInfoFilter();
                    const fieldsObj = {};

                    for (let prop in pageInfoFilterItem) {

                        if (pageInfoFilterItem.hasOwnProperty(prop)) {
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
            return pageInfoFilterData;
        },

        created() {
            this.fillPageInfoFilterFilter();
            this.fetchPageInfoFilterData();
        },

        computed: {
            ...mapGetters({
                pageInfoFilterList: 'getListPageInfoFilter'
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
                'findPageInfoFilter',
                'updatePageInfoFilter',
                'deletePageInfoFilter',
                'createPageInfoFilter',
            ]),

            ...mapMutations([
                'addPageInfoFilterItemToList',
                'deletePageInfoFilterFromList',
                'updatePageInfoFilterById',
            ]),

            fillPageInfoFilterFilter() {
                this.pageInfoFilterFilter.CurrentPage = 1;
                this.pageInfoFilterFilter.PerPage = 1000;
            },

            fetchPageInfoFilterData() {
                return this.findPageInfoFilter({
                    filter: this.pageInfoFilterFilter
                });
            },

            /**
             *
             * @param type
             */
            showPanel(type) {
                if (type === this.panel.create) {
                    this.panel.type = this.panel.create;
                    this.clearPanelPageInfoFilterItem();
                } else if (type === this.panel.edit) {
                    this.panel.type = this.panel.edit;
                    this.currentPageInfoFilterItem.isSelected = true;
                }

                this.panel.show = true;
            },

            closePanel() {
                this.panel.show = false;
                this.currentPageInfoFilterItem.isSelected = false;
                this.clearPanelPageInfoFilterItem();
            },

            selectPageInfoFilterItem(pageInfoFilterItem) {
                this.showPanel(this.panel.edit);
                this.currentPageInfoFilterItem.isSelected = true;
                Object.assign(this.currentPageInfoFilterItem.item, pageInfoFilterItem);
            },

            changeCurrentPageInfoFilterItem() {
                this.currentPageInfoFilterItem.hasChange = true;
            },

            cancelChanges() {
                this.clearPanelPageInfoFilterItem();
                this.closePanel();
            },

            clearPanelPageInfoFilterItem() {
                this.currentPageInfoFilterItem.item = new PageInfoFilter();
                this.currentPageInfoFilterItem.hasChange = false;
            },

            saveChangesSubmit() {
                if (this.isPanelCreate) {
                    this.createPageInfoFilterItemSubmit();
                    return;
                }

                if (this.isPanelEdit) {
                    this.editPageInfoFilterItemSubmit();
                }
            },

            createPageInfoFilterItemSubmit() {
                return this.createPageInfoFilter({
					data: this.currentPageInfoFilterItem.item,
                }).then((response) => {

                    if (response.Model) {
                        this.addPageInfoFilterItemToList(response.Model);
                        this.clearPanelPageInfoFilterItem();
                    } else {
                        console.error('Ошибка создания записи: ', response.Error);
                    }

                }).catch(error => {
                    console.error('Ошибка создания записи: ', error);
                });
            },

            editPageInfoFilterItemSubmit() {

                if (this.currentPageInfoFilterItem.hasChange) {
                    return this.updatePageInfoFilter({
                        id: this.currentPageInfoFilterItem.item.Id,
                        data: this.currentPageInfoFilterItem.item,
                    }).then((response) => {

                        if (response.Model) {
                            this.updatePageInfoFilterById(response.Model);
                            this.currentPageInfoFilterItem.hasChange = false;
                            this.clearPanelPageInfoFilterItem();
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

            deletePageInfoFilterItemHandler() {
                let deletedItemId = this.currentPageInfoFilterItem.item.Id;

                if (!this.currentPageInfoFilterItem.canDelete) {
                    this.currentPageInfoFilterItem.showDeleteConfirmation = true;
                    return;
                }

                this.deletePageInfoFilter({
                    id: deletedItemId
                }).then(response => {

                    if (response.IsSuccess) {
                        this.deletePageInfoFilterFromList(deletedItemId);
                        this.clearPanelPageInfoFilterItem();
                        this.currentPageInfoFilterItem.canDelete = false;
                        this.currentPageInfoFilterItem.isSelected = false;
                        this.panel.show = false;
                    } else {
                        console.error('Ошибка удаления элемента: ', response.Error);
                    }

                }).catch(error => {
                    console.error('Ошибка удаления элемента: ', error);
                });
            },

            confirmDeleteHandler() {
                this.currentPageInfoFilterItem.showDeleteConfirmation = false;
                this.currentPageInfoFilterItem.canDelete = true;
                this.deletePageInfoFilterItemHandler();
            },

            closeConfirmationPanel() {
                this.currentPageInfoFilterItem.showDeleteConfirmation = false;
            },
        },
    }
</script>

<style lang="scss">

</style>
