
	<template>
    <WorkSpace>
        <template #header>
            <slot name="pageHeader">
                <VHead level="h1">LayoutContent</VHead>
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
                            v-for="layoutContentItem in layoutContentList"
                            :key="layoutContentItem.Id"
                            @click="selectLayoutContentItem(layoutContentItem)"
                            class="sw-table__row_can-select"
                            :class="{'sw-table__row_is-selected': layoutContentItem.Id === currentLayoutContentItem.item.Id}"
                        >
                            <td v-for="(value, key) in fields" :key="key + '-fields'">
                                <VCheckbox v-if="isCheckbox(layoutContentItem[key])" :checked="layoutContentItem[key]" disabled></VCheckbox>
                                <VText v-else>{{ layoutContentItem[key] }}</VText>
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
                                        :for="`currentLayoutContentItem${key}`"
                                    >{{ filed }}</VLabel>
                                    <VInput
										v-if="isInput(currentLayoutContentItem.item[key])"
                                        v-model="currentLayoutContentItem.item[key]"
                                        width="dyn"
                                        :id="`currentLayoutContentItem${key}`"
                                        @input="changeCurrentLayoutContentItem"
                                    />
									<VCheckbox
										v-if="isCheckbox(currentLayoutContentItem.item[key])"
                                        v-model="currentLayoutContentItem.item[key]"
                                        :id="`currentLayoutContentItem${key}`"
										@input="changeCurrentLayoutContentItem"
									/>
									
                                </VSet>
                            </VSet>
                            <button type="submit" :disabled="!currentLayoutContentItem.hasChange" hidden></button>
                        </form>
                    </template>

                    <template #footer>
                        <VSet>
                            <VButton
                                @click="saveChangesSubmit"
                                accent
                                :text="panelSubmitButtonText"
                                :disabled="!currentLayoutContentItem.hasChange"
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
                    v-if="currentLayoutContentItem.showDeleteConfirmation"
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
                        :disabled="!currentLayoutContentItem.isSelected"
                        @click="deleteLayoutContentItemHandler"
                    />
                </VSet>
            </slot>
        </template>
    </WorkSpace>
</template>

<script>
    import layoutContentData from "../data/LayoutContentData";
    import { LayoutContent } from '../apiModel';
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
        name: 'LayoutContentGen',

        components: {VSelectSimple, VSign, VIcon, VButton, VPanel, VText, VInput, VLabel, VSet, VHead, WorkSpace, VCheckbox},

        props: {
            fields: {
                type: Object,
                default() {
                    const layoutContentItem = new LayoutContent();
                    const fieldsObj = {};

                    for (let prop in layoutContentItem) {

                        if (layoutContentItem.hasOwnProperty(prop)) {
                            fieldsObj[prop] = prop;
                        }

                    }

                    return fieldsObj;
                }
            },
            editFields: {
                type: Object,
                default() {
                    const layoutContentItem = new LayoutContent();
                    const fieldsObj = {};

                    for (let prop in layoutContentItem) {

                        if (layoutContentItem.hasOwnProperty(prop)) {
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
            return layoutContentData;
        },

        created() {
            this.fillLayoutContentFilter();
            this.fetchLayoutContentData();
        },

        computed: {
            ...mapGetters({
                layoutContentList: 'getListLayoutContent'
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
                'findLayoutContent',
                'updateLayoutContent',
                'deleteLayoutContent',
                'createLayoutContent',
            ]),

            ...mapMutations([
                'addLayoutContentItemToList',
                'deleteLayoutContentFromList',
                'updateLayoutContentById',
            ]),

            fillLayoutContentFilter() {
                this.layoutContentFilter.CurrentPage = 1;
                this.layoutContentFilter.PerPage = 1000;
            },

            fetchLayoutContentData() {
                return this.findLayoutContent({
                    filter: this.layoutContentFilter
                });
            },

            /**
             *
             * @param type
             */
            showPanel(type) {
                if (type === this.panel.create) {
                    this.panel.type = this.panel.create;
                    this.clearPanelLayoutContentItem();
                } else if (type === this.panel.edit) {
                    this.panel.type = this.panel.edit;
                    this.currentLayoutContentItem.isSelected = true;
                }

                this.panel.show = true;
            },

            closePanel() {
                this.panel.show = false;
                this.currentLayoutContentItem.isSelected = false;
                this.clearPanelLayoutContentItem();
            },

            selectLayoutContentItem(layoutContentItem) {
                this.showPanel(this.panel.edit);
                this.currentLayoutContentItem.isSelected = true;
                Object.assign(this.currentLayoutContentItem.item, layoutContentItem);
            },

            changeCurrentLayoutContentItem() {
                this.currentLayoutContentItem.hasChange = true;
            },

            cancelChanges() {
                this.clearPanelLayoutContentItem();
                this.closePanel();
            },

            clearPanelLayoutContentItem() {
                this.currentLayoutContentItem.item = new LayoutContent();
                this.currentLayoutContentItem.hasChange = false;
            },

            saveChangesSubmit() {
                if (this.isPanelCreate) {
                    this.createLayoutContentItemSubmit();
                    return;
                }

                if (this.isPanelEdit) {
                    this.editLayoutContentItemSubmit();
                }
            },

            createLayoutContentItemSubmit() {
                return this.createLayoutContent({
					data: this.currentLayoutContentItem.item,
                }).then((response) => {

                    if (response.Model) {
                        this.addLayoutContentItemToList(response.Model);
                        this.clearPanelLayoutContentItem();
                    } else {
                        console.error('Ошибка создания записи: ', response.Error);
                    }

                }).catch(error => {
                    console.error('Ошибка создания записи: ', error);
                });
            },

            editLayoutContentItemSubmit() {

                if (this.currentLayoutContentItem.hasChange) {
                    return this.updateLayoutContent({
                        id: this.currentLayoutContentItem.item.Id,
                        data: this.currentLayoutContentItem.item,
                    }).then((response) => {

                        if (response.Model) {
                            this.updateLayoutContentById(response.Model);
                            this.currentLayoutContentItem.hasChange = false;
                            this.clearPanelLayoutContentItem();
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

            deleteLayoutContentItemHandler() {
                let deletedItemId = this.currentLayoutContentItem.item.Id;

                if (!this.currentLayoutContentItem.canDelete) {
                    this.currentLayoutContentItem.showDeleteConfirmation = true;
                    return;
                }

                this.deleteLayoutContent({
                    id: deletedItemId
                }).then(response => {

                    if (response.IsSuccess) {
                        this.deleteLayoutContentFromList(deletedItemId);
                        this.clearPanelLayoutContentItem();
                        this.currentLayoutContentItem.canDelete = false;
                        this.currentLayoutContentItem.isSelected = false;
                        this.panel.show = false;
                    } else {
                        console.error('Ошибка удаления элемента: ', response.Error);
                    }

                }).catch(error => {
                    console.error('Ошибка удаления элемента: ', error);
                });
            },

            confirmDeleteHandler() {
                this.currentLayoutContentItem.showDeleteConfirmation = false;
                this.currentLayoutContentItem.canDelete = true;
                this.deleteLayoutContentItemHandler();
            },

            closeConfirmationPanel() {
                this.currentLayoutContentItem.showDeleteConfirmation = false;
            },
        },
    }
</script>

<style lang="scss">

</style>
