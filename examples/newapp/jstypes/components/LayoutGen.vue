
	<template>
    <WorkSpace>
        <template #header>
            <slot name="pageHeader">
                <VHead level="h1">Layout</VHead>
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
                            v-for="layoutItem in layoutList"
                            :key="layoutItem.Id"
                            @click="selectLayoutItem(layoutItem)"
                            class="sw-table__row_can-select"
                            :class="{'sw-table__row_is-selected': layoutItem.Id === currentLayoutItem.item.Id}"
                        >
                            <td v-for="(value, key) in fields" :key="key + '-fields'">
                                <VCheckbox v-if="isCheckbox(layoutItem[key])" :checked="layoutItem[key]" disabled></VCheckbox>
                                <VText v-else>{{ layoutItem[key] }}</VText>
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
                                        :for="`currentLayoutItem${key}`"
                                    >{{ filed }}</VLabel>
                                    <VInput
										v-if="isInput(currentLayoutItem.item[key])"
                                        v-model="currentLayoutItem.item[key]"
                                        width="dyn"
                                        :id="`currentLayoutItem${key}`"
                                        @input="changeCurrentLayoutItem"
                                    />
									<VCheckbox
										v-if="isCheckbox(currentLayoutItem.item[key])"
                                        v-model="currentLayoutItem.item[key]"
                                        :id="`currentLayoutItem${key}`"
										@input="changeCurrentLayoutItem"
									/>
									
                                </VSet>
                            </VSet>
                            <button type="submit" :disabled="!currentLayoutItem.hasChange" hidden></button>
                        </form>
                    </template>

                    <template #footer>
                        <VSet>
                            <VButton
                                @click="saveChangesSubmit"
                                accent
                                :text="panelSubmitButtonText"
                                :disabled="!currentLayoutItem.hasChange"
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
                    v-if="currentLayoutItem.showDeleteConfirmation"
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
                        :disabled="!currentLayoutItem.isSelected"
                        @click="deleteLayoutItemHandler"
                    />
                </VSet>
            </slot>
        </template>
    </WorkSpace>
</template>

<script>
    import layoutData from "../data/LayoutData";
    import { Layout } from '../apiModel';
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
        name: 'LayoutGen',

        components: {VSelectSimple, VSign, VIcon, VButton, VPanel, VText, VInput, VLabel, VSet, VHead, WorkSpace, VCheckbox},

        props: {
            fields: {
                type: Object,
                default() {
                    const layoutItem = new Layout();
                    const fieldsObj = {};

                    for (let prop in layoutItem) {

                        if (layoutItem.hasOwnProperty(prop)) {
                            fieldsObj[prop] = prop;
                        }

                    }

                    return fieldsObj;
                }
            },
            editFields: {
                type: Object,
                default() {
                    const layoutItem = new Layout();
                    const fieldsObj = {};

                    for (let prop in layoutItem) {

                        if (layoutItem.hasOwnProperty(prop)) {
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
            return layoutData;
        },

        created() {
            this.fillLayoutFilter();
            this.fetchLayoutData();
        },

        computed: {
            ...mapGetters({
                layoutList: 'getListLayout'
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
                'findLayout',
                'updateLayout',
                'deleteLayout',
                'createLayout',
            ]),

            ...mapMutations([
                'addLayoutItemToList',
                'deleteLayoutFromList',
                'updateLayoutById',
            ]),

            fillLayoutFilter() {
                this.layoutFilter.CurrentPage = 1;
                this.layoutFilter.PerPage = 1000;
            },

            fetchLayoutData() {
                return this.findLayout({
                    filter: this.layoutFilter
                });
            },

            /**
             *
             * @param type
             */
            showPanel(type) {
                if (type === this.panel.create) {
                    this.panel.type = this.panel.create;
                    this.clearPanelLayoutItem();
                } else if (type === this.panel.edit) {
                    this.panel.type = this.panel.edit;
                    this.currentLayoutItem.isSelected = true;
                }

                this.panel.show = true;
            },

            closePanel() {
                this.panel.show = false;
                this.currentLayoutItem.isSelected = false;
                this.clearPanelLayoutItem();
            },

            selectLayoutItem(layoutItem) {
                this.showPanel(this.panel.edit);
                this.currentLayoutItem.isSelected = true;
                Object.assign(this.currentLayoutItem.item, layoutItem);
            },

            changeCurrentLayoutItem() {
                this.currentLayoutItem.hasChange = true;
            },

            cancelChanges() {
                this.clearPanelLayoutItem();
                this.closePanel();
            },

            clearPanelLayoutItem() {
                this.currentLayoutItem.item = new Layout();
                this.currentLayoutItem.hasChange = false;
            },

            saveChangesSubmit() {
                if (this.isPanelCreate) {
                    this.createLayoutItemSubmit();
                    return;
                }

                if (this.isPanelEdit) {
                    this.editLayoutItemSubmit();
                }
            },

            createLayoutItemSubmit() {
                return this.createLayout({
					data: this.currentLayoutItem.item,
                }).then((response) => {

                    if (response.Model) {
                        this.addLayoutItemToList(response.Model);
                        this.clearPanelLayoutItem();
                    } else {
                        console.error('Ошибка создания записи: ', response.Error);
                    }

                }).catch(error => {
                    console.error('Ошибка создания записи: ', error);
                });
            },

            editLayoutItemSubmit() {

                if (this.currentLayoutItem.hasChange) {
                    return this.updateLayout({
                        id: this.currentLayoutItem.item.Id,
                        data: this.currentLayoutItem.item,
                    }).then((response) => {

                        if (response.Model) {
                            this.updateLayoutById(response.Model);
                            this.currentLayoutItem.hasChange = false;
                            this.clearPanelLayoutItem();
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

            deleteLayoutItemHandler() {
                let deletedItemId = this.currentLayoutItem.item.Id;

                if (!this.currentLayoutItem.canDelete) {
                    this.currentLayoutItem.showDeleteConfirmation = true;
                    return;
                }

                this.deleteLayout({
                    id: deletedItemId
                }).then(response => {

                    if (response.IsSuccess) {
                        this.deleteLayoutFromList(deletedItemId);
                        this.clearPanelLayoutItem();
                        this.currentLayoutItem.canDelete = false;
                        this.currentLayoutItem.isSelected = false;
                        this.panel.show = false;
                    } else {
                        console.error('Ошибка удаления элемента: ', response.Error);
                    }

                }).catch(error => {
                    console.error('Ошибка удаления элемента: ', error);
                });
            },

            confirmDeleteHandler() {
                this.currentLayoutItem.showDeleteConfirmation = false;
                this.currentLayoutItem.canDelete = true;
                this.deleteLayoutItemHandler();
            },

            closeConfirmationPanel() {
                this.currentLayoutItem.showDeleteConfirmation = false;
            },
        },
    }
</script>

<style lang="scss">

</style>
