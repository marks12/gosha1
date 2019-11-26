
	<template>
    <WorkSpace>
        <template #header>
            <slot name="pageHeader">
                <VHead level="h1">PageInfo</VHead>
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
                            v-for="pageInfoItem in pageInfoList"
                            :key="pageInfoItem.Id"
                            @click="selectPageInfoItem(pageInfoItem)"
                            class="sw-table__row_can-select"
                            :class="{'sw-table__row_is-selected': pageInfoItem.Id === currentPageInfoItem.item.Id}"
                        >
                            <td v-for="(value, key) in fields" :key="key + '-fields'">
                                <VCheckbox v-if="isCheckbox(pageInfoItem[key])" :checked="pageInfoItem[key]" disabled></VCheckbox>
                                <VText v-else>{{ pageInfoItem[key] }}</VText>
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
                                        :for="`currentPageInfoItem${key}`"
                                    >{{ filed }}</VLabel>
                                    <VInput
										v-if="isInput(currentPageInfoItem.item[key])"
                                        v-model="currentPageInfoItem.item[key]"
                                        width="dyn"
                                        :id="`currentPageInfoItem${key}`"
                                        @input="changeCurrentPageInfoItem"
                                    />
									<VCheckbox
										v-if="isCheckbox(currentPageInfoItem.item[key])"
                                        v-model="currentPageInfoItem.item[key]"
                                        :id="`currentPageInfoItem${key}`"
										@input="changeCurrentPageInfoItem"
									/>
									
                                </VSet>
                            </VSet>
                            <button type="submit" :disabled="!currentPageInfoItem.hasChange" hidden></button>
                        </form>
                    </template>

                    <template #footer>
                        <VSet>
                            <VButton
                                @click="saveChangesSubmit"
                                accent
                                :text="panelSubmitButtonText"
                                :disabled="!currentPageInfoItem.hasChange"
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
                    v-if="currentPageInfoItem.showDeleteConfirmation"
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
                        :disabled="!currentPageInfoItem.isSelected"
                        @click="deletePageInfoItemHandler"
                    />
                </VSet>
            </slot>
        </template>
    </WorkSpace>
</template>

<script>
    import pageInfoData from "../data/PageInfoData";
    import { PageInfo } from '../apiModel';
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
        name: 'PageInfoGen',

        components: {VSelectSimple, VSign, VIcon, VButton, VPanel, VText, VInput, VLabel, VSet, VHead, WorkSpace, VCheckbox},

        props: {
            fields: {
                type: Object,
                default() {
                    const pageInfoItem = new PageInfo();
                    const fieldsObj = {};

                    for (let prop in pageInfoItem) {

                        if (pageInfoItem.hasOwnProperty(prop)) {
                            fieldsObj[prop] = prop;
                        }

                    }

                    return fieldsObj;
                }
            },
            editFields: {
                type: Object,
                default() {
                    const pageInfoItem = new PageInfo();
                    const fieldsObj = {};

                    for (let prop in pageInfoItem) {

                        if (pageInfoItem.hasOwnProperty(prop)) {
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
            return pageInfoData;
        },

        created() {
            this.fillPageInfoFilter();
            this.fetchPageInfoData();
        },

        computed: {
            ...mapGetters({
                pageInfoList: 'getListPageInfo'
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
                'findPageInfo',
                'updatePageInfo',
                'deletePageInfo',
                'createPageInfo',
            ]),

            ...mapMutations([
                'addPageInfoItemToList',
                'deletePageInfoFromList',
                'updatePageInfoById',
            ]),

            fillPageInfoFilter() {
                this.pageInfoFilter.CurrentPage = 1;
                this.pageInfoFilter.PerPage = 1000;
            },

            fetchPageInfoData() {
                return this.findPageInfo({
                    filter: this.pageInfoFilter
                });
            },

            /**
             *
             * @param type
             */
            showPanel(type) {
                if (type === this.panel.create) {
                    this.panel.type = this.panel.create;
                    this.clearPanelPageInfoItem();
                } else if (type === this.panel.edit) {
                    this.panel.type = this.panel.edit;
                    this.currentPageInfoItem.isSelected = true;
                }

                this.panel.show = true;
            },

            closePanel() {
                this.panel.show = false;
                this.currentPageInfoItem.isSelected = false;
                this.clearPanelPageInfoItem();
            },

            selectPageInfoItem(pageInfoItem) {
                this.showPanel(this.panel.edit);
                this.currentPageInfoItem.isSelected = true;
                Object.assign(this.currentPageInfoItem.item, pageInfoItem);
            },

            changeCurrentPageInfoItem() {
                this.currentPageInfoItem.hasChange = true;
            },

            cancelChanges() {
                this.clearPanelPageInfoItem();
                this.closePanel();
            },

            clearPanelPageInfoItem() {
                this.currentPageInfoItem.item = new PageInfo();
                this.currentPageInfoItem.hasChange = false;
            },

            saveChangesSubmit() {
                if (this.isPanelCreate) {
                    this.createPageInfoItemSubmit();
                    return;
                }

                if (this.isPanelEdit) {
                    this.editPageInfoItemSubmit();
                }
            },

            createPageInfoItemSubmit() {
                return this.createPageInfo({
					data: this.currentPageInfoItem.item,
                }).then((response) => {

                    if (response.Model) {
                        this.addPageInfoItemToList(response.Model);
                        this.clearPanelPageInfoItem();
                    } else {
                        console.error('Ошибка создания записи: ', response.Error);
                    }

                }).catch(error => {
                    console.error('Ошибка создания записи: ', error);
                });
            },

            editPageInfoItemSubmit() {

                if (this.currentPageInfoItem.hasChange) {
                    return this.updatePageInfo({
                        id: this.currentPageInfoItem.item.Id,
                        data: this.currentPageInfoItem.item,
                    }).then((response) => {

                        if (response.Model) {
                            this.updatePageInfoById(response.Model);
                            this.currentPageInfoItem.hasChange = false;
                            this.clearPanelPageInfoItem();
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

            deletePageInfoItemHandler() {
                let deletedItemId = this.currentPageInfoItem.item.Id;

                if (!this.currentPageInfoItem.canDelete) {
                    this.currentPageInfoItem.showDeleteConfirmation = true;
                    return;
                }

                this.deletePageInfo({
                    id: deletedItemId
                }).then(response => {

                    if (response.IsSuccess) {
                        this.deletePageInfoFromList(deletedItemId);
                        this.clearPanelPageInfoItem();
                        this.currentPageInfoItem.canDelete = false;
                        this.currentPageInfoItem.isSelected = false;
                        this.panel.show = false;
                    } else {
                        console.error('Ошибка удаления элемента: ', response.Error);
                    }

                }).catch(error => {
                    console.error('Ошибка удаления элемента: ', error);
                });
            },

            confirmDeleteHandler() {
                this.currentPageInfoItem.showDeleteConfirmation = false;
                this.currentPageInfoItem.canDelete = true;
                this.deletePageInfoItemHandler();
            },

            closeConfirmationPanel() {
                this.currentPageInfoItem.showDeleteConfirmation = false;
            },
        },
    }
</script>

<style lang="scss">

</style>
