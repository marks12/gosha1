
	<template>
    <WorkSpace>
        <template #header>
            <slot name="pageHeader">
                <VHead level="h1">Pagination</VHead>
            </slot>
        </template>

        <template #content>
            <slot name="data">
                <table>
                    <thead>
                        <tr>
                            <th v-for="header in fields">{{ header }}</th>
                        </tr>
                    </thead>
            
                    <tbody>
                        <tr
                            v-for="paginationItem in paginationList"
                            :key="paginationItem.Id"
                            @click="selectPaginationItem(paginationItem)"
                            class="sw-table__row_can-select"
                            :class="{'sw-table__row_is-selected': paginationItem.Id === currentPaginationItem.item.Id}"
                        >
                            <td v-for="(value, key) in fields">
                                <VCheckbox v-if="isCheckbox(applicationItem[key])" :checked="applicationItem[key]" disabled></VCheckbox>
                                <VText v-else>{{ applicationItem[key] }}</VText>
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
                                    v-for="(filed, key) in editFields"
                                    vertical-align="center"
                                >
                                    <VLabel
                                        width="col4"
                                        :for="`currentPaginationItem${key}`"
                                    >{{ filed }}</VLabel>
                                    <VInput
										v-if="isInput(currentPaginationItem.item[key])"
                                        v-model="currentPaginationItem.item[key]"
                                        width="dyn"
                                        :id="`currentPaginationItem${key}`"
                                        @input="changeCurrentPaginationItem"
                                    />
									<VCheckbox
										v-if="isCheckbox(currentPaginationItem.item[key])"
                                        v-model="currentPaginationItem.item[key]"
                                        :id="`currentPaginationItem${key}`"
										@input="changeCurrentApplicationItem"
									/>
									
                                </VSet>
                            </VSet>
                            <button type="submit" :disabled="!currentPaginationItem.hasChange" hidden></button>
                        </form>
                    </template>

                    <template #footer>
                        <VSet>
                            <VButton
                                @click="saveChangesSubmit"
                                accent
                                :text="panelSubmitButtonText"
                                :disabled="!currentPaginationItem.hasChange"
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
                    v-if="currentPaginationItem.showDeleteConfirmation"
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
                        text="Добавить"
                        accent
                        @click="showPanel(panel.create)"
                    />
                    <VButton
                        text="Удалить"
                        :disabled="!currentPaginationItem.isSelected"
                        @click="deletePaginationItemHandler"
                    />
                </VSet>
            </slot>
        </template>
    </WorkSpace>
</template>

<script>
    import paginationData from "../data/PaginationData";
    import { Pagination } from '../apiModel';
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
    import VSelect from "swui/src/components/VSelect";

    export default {
        name: 'PaginationGen',

        components: {VSelect, VSign, VIcon, VButton, VPanel, VText, VInput, VLabel, VSet, VHead, WorkSpace, VCheckbox},

        props: {
            fields: {
                type: Object,
                default() {
                    const paginationItem = new Pagination();
                    const fieldsObj = {};

                    for (let prop in paginationItem) {

                        if (paginationItem.hasOwnProperty(prop)) {
                            fieldsObj[prop] = prop;
                        }

                    }

                    return fieldsObj;
                }
            },
            editFields: {
                type: Object,
                default() {
                    const paginationItem = new Pagination();
                    const fieldsObj = {};

                    for (let prop in paginationItem) {

                        if (paginationItem.hasOwnProperty(prop)) {
                            fieldsObj[prop] = prop;
                        }

                    }

                    return fieldsObj;
                }
            }
        },

        data() {
            return paginationData;
        },

        created() {
            this.fillPaginationFilter();
            this.fetchPaginationData();
        },

        computed: {
            ...mapGetters({
                paginationList: 'getListPagination'
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
                'findPagination',
                'updatePagination',
                'deletePagination',
                'createPagination',
            ]),

            ...mapMutations([
                'addPaginationItemToList',
                'deletePaginationFromList',
                'updatePaginationById',
            ]),

            fillPaginationFilter() {
                this.paginationFilter.CurrentPage = 1;
                this.paginationFilter.PerPage = 1000;
            },

            fetchPaginationData() {
                return this.findPagination({
                    filter: this.paginationFilter
                });
            },

            /**
             *
             * @param type
             */
            showPanel(type) {
                if (type === this.panel.create) {
                    this.panel.type = this.panel.create;
                    this.clearPanelPaginationItem();
                } else if (type === this.panel.edit) {
                    this.panel.type = this.panel.edit;
                    this.currentPaginationItem.isSelected = true;
                }

                this.panel.show = true;
            },

            closePanel() {
                this.panel.show = false;
                this.currentPaginationItem.isSelected = false;
                this.clearPanelPaginationItem();
            },

            selectPaginationItem(paginationItem) {
                this.showPanel(this.panel.edit);
                this.currentPaginationItem.isSelected = true;
                Object.assign(this.currentPaginationItem.item, paginationItem);
            },

            changeCurrentPaginationItem() {
                this.currentPaginationItem.hasChange = true;
            },

            cancelChanges() {
                this.clearPanelPaginationItem();
                this.closePanel();
            },

            clearPanelPaginationItem() {
                this.currentPaginationItem.item = new Pagination();
                this.currentPaginationItem.hasChange = false;
            },

            saveChangesSubmit() {
                if (this.isPanelCreate) {
                    this.createPaginationItemSubmit();
                    return;
                }

                if (this.isPanelEdit) {
                    this.editPaginationItemSubmit();
                }
            },

            createPaginationItemSubmit() {
                this.createPagination({
                    data: {
                        Name: this.currentPaginationItem.item.Name,
                        Value: this.currentPaginationItem.item.Value,
                        Description: this.currentPaginationItem.item.Description,
                    }
                }).then((response) => {

                    if (response.Model) {
                        this.addPaginationItemToList(response.Model);
                        this.clearPanelPaginationItem();
                    } else {
                        console.error('Ошибка создания записи: ', response.Error);
                    }

                }).catch(error => {
                    console.error('Ошибка создания записи: ', error);
                });
            },

            editPaginationItemSubmit() {
                if (this.currentPaginationItem.hasChange) {
                    this.updatePagination({
                        id: this.currentPaginationItem.item.Id,
                        data: this.currentPaginationItem.item,
                    }).then((response) => {

                        if (response.Model) {
                            this.updatePaginationById(response.Model);
                            this.currentPaginationItem.hasChange = false;
                            this.clearPanelPaginationItem();
                            this.closePanel();
                        } else {
                            console.error('Ошибка изменения записи: ', response.Error);
                        }

                    }).catch(error => {
                        console.error('Ошибка изменения записи: ', error);
                    });
                }
            },

            deletePaginationItemHandler() {
                let deletedItemId = this.currentPaginationItem.item.Id;

                if (!this.currentPaginationItem.canDelete) {
                    this.currentPaginationItem.showDeleteConfirmation = true;
                    return;
                }

                this.deletePagination({
                    id: deletedItemId
                }).then(response => {

                    if (response.IsSuccess) {
                        this.deletePaginationFromList(deletedItemId);
                        this.clearPanelPaginationItem();
                        this.currentPaginationItem.canDelete = false;
                        this.currentPaginationItem.isSelected = false;
                        this.panel.show = false;
                    } else {
                        console.error('Ошибка удаления элемента: ', response.Error);
                    }

                }).catch(error => {
                    console.error('Ошибка удаления элемента: ', error);
                });
            },

            confirmDeleteHandler() {
                this.currentPaginationItem.showDeleteConfirmation = false;
                this.currentPaginationItem.canDelete = true;
                this.deletePaginationItemHandler();
            },

            closeConfirmationPanel() {
                this.currentPaginationItem.showDeleteConfirmation = false;
            },
        },
    }
</script>

<style lang="scss">

</style>
