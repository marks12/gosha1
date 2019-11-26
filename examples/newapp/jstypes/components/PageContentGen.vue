
	<template>
    <WorkSpace>
        <template #header>
            <slot name="pageHeader">
                <VHead level="h1">PageContent</VHead>
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
                            v-for="pageContentItem in pageContentList"
                            :key="pageContentItem.Id"
                            @click="selectPageContentItem(pageContentItem)"
                            class="sw-table__row_can-select"
                            :class="{'sw-table__row_is-selected': pageContentItem.Id === currentPageContentItem.item.Id}"
                        >
                            <td v-for="(value, key) in fields" :key="key + '-fields'">
                                <VCheckbox v-if="isCheckbox(pageContentItem[key])" :checked="pageContentItem[key]" disabled></VCheckbox>
                                <VText v-else>{{ pageContentItem[key] }}</VText>
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
                                        :for="`currentPageContentItem${key}`"
                                    >{{ filed }}</VLabel>
                                    <VInput
										v-if="isInput(currentPageContentItem.item[key])"
                                        v-model="currentPageContentItem.item[key]"
                                        width="dyn"
                                        :id="`currentPageContentItem${key}`"
                                        @input="changeCurrentPageContentItem"
                                    />
									<VCheckbox
										v-if="isCheckbox(currentPageContentItem.item[key])"
                                        v-model="currentPageContentItem.item[key]"
                                        :id="`currentPageContentItem${key}`"
										@input="changeCurrentPageContentItem"
									/>
									
                                </VSet>
                            </VSet>
                            <button type="submit" :disabled="!currentPageContentItem.hasChange" hidden></button>
                        </form>
                    </template>

                    <template #footer>
                        <VSet>
                            <VButton
                                @click="saveChangesSubmit"
                                accent
                                :text="panelSubmitButtonText"
                                :disabled="!currentPageContentItem.hasChange"
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
                    v-if="currentPageContentItem.showDeleteConfirmation"
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
                        :disabled="!currentPageContentItem.isSelected"
                        @click="deletePageContentItemHandler"
                    />
                </VSet>
            </slot>
        </template>
    </WorkSpace>
</template>

<script>
    import pageContentData from "../data/PageContentData";
    import { PageContent } from '../apiModel';
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
        name: 'PageContentGen',

        components: {VSelectSimple, VSign, VIcon, VButton, VPanel, VText, VInput, VLabel, VSet, VHead, WorkSpace, VCheckbox},

        props: {
            fields: {
                type: Object,
                default() {
                    const pageContentItem = new PageContent();
                    const fieldsObj = {};

                    for (let prop in pageContentItem) {

                        if (pageContentItem.hasOwnProperty(prop)) {
                            fieldsObj[prop] = prop;
                        }

                    }

                    return fieldsObj;
                }
            },
            editFields: {
                type: Object,
                default() {
                    const pageContentItem = new PageContent();
                    const fieldsObj = {};

                    for (let prop in pageContentItem) {

                        if (pageContentItem.hasOwnProperty(prop)) {
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
            return pageContentData;
        },

        created() {
            this.fillPageContentFilter();
            this.fetchPageContentData();
        },

        computed: {
            ...mapGetters({
                pageContentList: 'getListPageContent'
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
                'findPageContent',
                'updatePageContent',
                'deletePageContent',
                'createPageContent',
            ]),

            ...mapMutations([
                'addPageContentItemToList',
                'deletePageContentFromList',
                'updatePageContentById',
            ]),

            fillPageContentFilter() {
                this.pageContentFilter.CurrentPage = 1;
                this.pageContentFilter.PerPage = 1000;
            },

            fetchPageContentData() {
                return this.findPageContent({
                    filter: this.pageContentFilter
                });
            },

            /**
             *
             * @param type
             */
            showPanel(type) {
                if (type === this.panel.create) {
                    this.panel.type = this.panel.create;
                    this.clearPanelPageContentItem();
                } else if (type === this.panel.edit) {
                    this.panel.type = this.panel.edit;
                    this.currentPageContentItem.isSelected = true;
                }

                this.panel.show = true;
            },

            closePanel() {
                this.panel.show = false;
                this.currentPageContentItem.isSelected = false;
                this.clearPanelPageContentItem();
            },

            selectPageContentItem(pageContentItem) {
                this.showPanel(this.panel.edit);
                this.currentPageContentItem.isSelected = true;
                Object.assign(this.currentPageContentItem.item, pageContentItem);
            },

            changeCurrentPageContentItem() {
                this.currentPageContentItem.hasChange = true;
            },

            cancelChanges() {
                this.clearPanelPageContentItem();
                this.closePanel();
            },

            clearPanelPageContentItem() {
                this.currentPageContentItem.item = new PageContent();
                this.currentPageContentItem.hasChange = false;
            },

            saveChangesSubmit() {
                if (this.isPanelCreate) {
                    this.createPageContentItemSubmit();
                    return;
                }

                if (this.isPanelEdit) {
                    this.editPageContentItemSubmit();
                }
            },

            createPageContentItemSubmit() {
                return this.createPageContent({
					data: this.currentPageContentItem.item,
                }).then((response) => {

                    if (response.Model) {
                        this.addPageContentItemToList(response.Model);
                        this.clearPanelPageContentItem();
                    } else {
                        console.error('Ошибка создания записи: ', response.Error);
                    }

                }).catch(error => {
                    console.error('Ошибка создания записи: ', error);
                });
            },

            editPageContentItemSubmit() {

                if (this.currentPageContentItem.hasChange) {
                    return this.updatePageContent({
                        id: this.currentPageContentItem.item.Id,
                        data: this.currentPageContentItem.item,
                    }).then((response) => {

                        if (response.Model) {
                            this.updatePageContentById(response.Model);
                            this.currentPageContentItem.hasChange = false;
                            this.clearPanelPageContentItem();
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

            deletePageContentItemHandler() {
                let deletedItemId = this.currentPageContentItem.item.Id;

                if (!this.currentPageContentItem.canDelete) {
                    this.currentPageContentItem.showDeleteConfirmation = true;
                    return;
                }

                this.deletePageContent({
                    id: deletedItemId
                }).then(response => {

                    if (response.IsSuccess) {
                        this.deletePageContentFromList(deletedItemId);
                        this.clearPanelPageContentItem();
                        this.currentPageContentItem.canDelete = false;
                        this.currentPageContentItem.isSelected = false;
                        this.panel.show = false;
                    } else {
                        console.error('Ошибка удаления элемента: ', response.Error);
                    }

                }).catch(error => {
                    console.error('Ошибка удаления элемента: ', error);
                });
            },

            confirmDeleteHandler() {
                this.currentPageContentItem.showDeleteConfirmation = false;
                this.currentPageContentItem.canDelete = true;
                this.deletePageContentItemHandler();
            },

            closeConfirmationPanel() {
                this.currentPageContentItem.showDeleteConfirmation = false;
            },
        },
    }
</script>

<style lang="scss">

</style>
