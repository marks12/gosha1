
	<template>
    <WorkSpace>
        <template #header>
            <slot name="pageHeader">
                <VHead level="h1">PageTemplate</VHead>
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
                            v-for="pageTemplateItem in pageTemplateList"
                            :key="pageTemplateItem.Id"
                            @click="selectPageTemplateItem(pageTemplateItem)"
                            class="sw-table__row_can-select"
                            :class="{'sw-table__row_is-selected': pageTemplateItem.Id === currentPageTemplateItem.item.Id}"
                        >
                            <td v-for="(value, key) in fields" :key="key + '-fields'">
                                <VCheckbox v-if="isCheckbox(pageTemplateItem[key])" :checked="pageTemplateItem[key]" disabled></VCheckbox>
                                <VText v-else>{{ pageTemplateItem[key] }}</VText>
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
                                        :for="`currentPageTemplateItem${key}`"
                                    >{{ filed }}</VLabel>
                                    <VInput
										v-if="isInput(currentPageTemplateItem.item[key])"
                                        v-model="currentPageTemplateItem.item[key]"
                                        width="dyn"
                                        :id="`currentPageTemplateItem${key}`"
                                        @input="changeCurrentPageTemplateItem"
                                    />
									<VCheckbox
										v-if="isCheckbox(currentPageTemplateItem.item[key])"
                                        v-model="currentPageTemplateItem.item[key]"
                                        :id="`currentPageTemplateItem${key}`"
										@input="changeCurrentPageTemplateItem"
									/>
									
                                </VSet>
                            </VSet>
                            <button type="submit" :disabled="!currentPageTemplateItem.hasChange" hidden></button>
                        </form>
                    </template>

                    <template #footer>
                        <VSet>
                            <VButton
                                @click="saveChangesSubmit"
                                accent
                                :text="panelSubmitButtonText"
                                :disabled="!currentPageTemplateItem.hasChange"
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
                    v-if="currentPageTemplateItem.showDeleteConfirmation"
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
                        :disabled="!currentPageTemplateItem.isSelected"
                        @click="deletePageTemplateItemHandler"
                    />
                </VSet>
            </slot>
        </template>
    </WorkSpace>
</template>

<script>
    import pageTemplateData from "../data/PageTemplateData";
    import { PageTemplate } from '../apiModel';
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
        name: 'PageTemplateGen',

        components: {VSelectSimple, VSign, VIcon, VButton, VPanel, VText, VInput, VLabel, VSet, VHead, WorkSpace, VCheckbox},

        props: {
            fields: {
                type: Object,
                default() {
                    const pageTemplateItem = new PageTemplate();
                    const fieldsObj = {};

                    for (let prop in pageTemplateItem) {

                        if (pageTemplateItem.hasOwnProperty(prop)) {
                            fieldsObj[prop] = prop;
                        }

                    }

                    return fieldsObj;
                }
            },
            editFields: {
                type: Object,
                default() {
                    const pageTemplateItem = new PageTemplate();
                    const fieldsObj = {};

                    for (let prop in pageTemplateItem) {

                        if (pageTemplateItem.hasOwnProperty(prop)) {
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
            return pageTemplateData;
        },

        created() {
            this.fillPageTemplateFilter();
            this.fetchPageTemplateData();
        },

        computed: {
            ...mapGetters({
                pageTemplateList: 'getListPageTemplate'
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
                'findPageTemplate',
                'updatePageTemplate',
                'deletePageTemplate',
                'createPageTemplate',
            ]),

            ...mapMutations([
                'addPageTemplateItemToList',
                'deletePageTemplateFromList',
                'updatePageTemplateById',
            ]),

            fillPageTemplateFilter() {
                this.pageTemplateFilter.CurrentPage = 1;
                this.pageTemplateFilter.PerPage = 1000;
            },

            fetchPageTemplateData() {
                return this.findPageTemplate({
                    filter: this.pageTemplateFilter
                });
            },

            /**
             *
             * @param type
             */
            showPanel(type) {
                if (type === this.panel.create) {
                    this.panel.type = this.panel.create;
                    this.clearPanelPageTemplateItem();
                } else if (type === this.panel.edit) {
                    this.panel.type = this.panel.edit;
                    this.currentPageTemplateItem.isSelected = true;
                }

                this.panel.show = true;
            },

            closePanel() {
                this.panel.show = false;
                this.currentPageTemplateItem.isSelected = false;
                this.clearPanelPageTemplateItem();
            },

            selectPageTemplateItem(pageTemplateItem) {
                this.showPanel(this.panel.edit);
                this.currentPageTemplateItem.isSelected = true;
                Object.assign(this.currentPageTemplateItem.item, pageTemplateItem);
            },

            changeCurrentPageTemplateItem() {
                this.currentPageTemplateItem.hasChange = true;
            },

            cancelChanges() {
                this.clearPanelPageTemplateItem();
                this.closePanel();
            },

            clearPanelPageTemplateItem() {
                this.currentPageTemplateItem.item = new PageTemplate();
                this.currentPageTemplateItem.hasChange = false;
            },

            saveChangesSubmit() {
                if (this.isPanelCreate) {
                    this.createPageTemplateItemSubmit();
                    return;
                }

                if (this.isPanelEdit) {
                    this.editPageTemplateItemSubmit();
                }
            },

            createPageTemplateItemSubmit() {
                return this.createPageTemplate({
					data: this.currentPageTemplateItem.item,
                }).then((response) => {

                    if (response.Model) {
                        this.addPageTemplateItemToList(response.Model);
                        this.clearPanelPageTemplateItem();
                    } else {
                        console.error('Ошибка создания записи: ', response.Error);
                    }

                }).catch(error => {
                    console.error('Ошибка создания записи: ', error);
                });
            },

            editPageTemplateItemSubmit() {

                if (this.currentPageTemplateItem.hasChange) {
                    return this.updatePageTemplate({
                        id: this.currentPageTemplateItem.item.Id,
                        data: this.currentPageTemplateItem.item,
                    }).then((response) => {

                        if (response.Model) {
                            this.updatePageTemplateById(response.Model);
                            this.currentPageTemplateItem.hasChange = false;
                            this.clearPanelPageTemplateItem();
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

            deletePageTemplateItemHandler() {
                let deletedItemId = this.currentPageTemplateItem.item.Id;

                if (!this.currentPageTemplateItem.canDelete) {
                    this.currentPageTemplateItem.showDeleteConfirmation = true;
                    return;
                }

                this.deletePageTemplate({
                    id: deletedItemId
                }).then(response => {

                    if (response.IsSuccess) {
                        this.deletePageTemplateFromList(deletedItemId);
                        this.clearPanelPageTemplateItem();
                        this.currentPageTemplateItem.canDelete = false;
                        this.currentPageTemplateItem.isSelected = false;
                        this.panel.show = false;
                    } else {
                        console.error('Ошибка удаления элемента: ', response.Error);
                    }

                }).catch(error => {
                    console.error('Ошибка удаления элемента: ', error);
                });
            },

            confirmDeleteHandler() {
                this.currentPageTemplateItem.showDeleteConfirmation = false;
                this.currentPageTemplateItem.canDelete = true;
                this.deletePageTemplateItemHandler();
            },

            closeConfirmationPanel() {
                this.currentPageTemplateItem.showDeleteConfirmation = false;
            },
        },
    }
</script>

<style lang="scss">

</style>
