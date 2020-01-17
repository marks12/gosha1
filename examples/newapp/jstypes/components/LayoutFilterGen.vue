
	<template>
    <WorkSpace>
        <template #header>
            <slot name="pageHeader">
                <VHead level="h1">LayoutFilter</VHead>
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
                            v-for="layoutFilterItem in layoutFilterList"
                            :key="layoutFilterItem.Id"
                            @click="selectLayoutFilterItem(layoutFilterItem)"
                            class="sw-table__row_can-select"
                            :class="{'sw-table__row_is-selected': layoutFilterItem.Id === currentLayoutFilterItem.item.Id}"
                        >
                            <td v-for="(value, key) in fields" :key="key + '-fields'">
                                <VCheckbox v-if="isCheckbox(layoutFilterItem[key])" :checked="layoutFilterItem[key]" disabled></VCheckbox>
                                <VText v-else>{{ layoutFilterItem[key] }}</VText>
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
                                        :for="`currentLayoutFilterItem${key}`"
                                    >{{ filed }}</VLabel>
                                    <VInput
										v-if="isInput(currentLayoutFilterItem.item[key])"
                                        v-model="currentLayoutFilterItem.item[key]"
                                        width="dyn"
                                        :id="`currentLayoutFilterItem${key}`"
                                        @input="changeCurrentLayoutFilterItem"
                                    />
									<VCheckbox
										v-if="isCheckbox(currentLayoutFilterItem.item[key])"
                                        v-model="currentLayoutFilterItem.item[key]"
                                        :id="`currentLayoutFilterItem${key}`"
										@input="changeCurrentLayoutFilterItem"
									/>
									
                                </VSet>
                            </VSet>
                            <button type="submit" :disabled="!currentLayoutFilterItem.hasChange" hidden></button>
                        </form>
                    </template>

                    <template #footer>
                        <VSet>
                            <VButton
                                @click="saveChangesSubmit"
                                accent
                                :text="panelSubmitButtonText"
                                :disabled="!currentLayoutFilterItem.hasChange"
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
                    v-if="currentLayoutFilterItem.showDeleteConfirmation"
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
                        :disabled="!currentLayoutFilterItem.isSelected"
                        @click="deleteLayoutFilterItemHandler"
                    />
                </VSet>
            </slot>
        </template>
    </WorkSpace>
</template>

<script>
    import layoutFilterData from "../data/LayoutFilterData";
    import { LayoutFilter } from '../apiModel';
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
        name: 'LayoutFilterGen',

        components: {VButton, VPanel, VText, VInput, VLabel, VSet, VHead, WorkSpace, VCheckbox},

        props: {
            fields: {
                type: Object,
                default() {
                    const layoutFilterItem = new LayoutFilter();
                    const fieldsObj = {};

                    for (let prop in layoutFilterItem) {

                        if (layoutFilterItem.hasOwnProperty(prop)) {
                            fieldsObj[prop] = prop;
                        }

                    }

                    return fieldsObj;
                }
            },
            editFields: {
                type: Object,
                default() {
                    const layoutFilterItem = new LayoutFilter();
                    const fieldsObj = {};

                    for (let prop in layoutFilterItem) {

                        if (layoutFilterItem.hasOwnProperty(prop)) {
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
            return layoutFilterData;
        },

        created() {
			this.onCreated();
        },

        computed: {
            ...mapGetters({
                layoutFilterList: 'getListLayoutFilter'
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
                'findLayoutFilter',
                'updateLayoutFilter',
                'deleteLayoutFilter',
                'createLayoutFilter',
            ]),

            ...mapMutations([
                'addLayoutFilterItemToList',
                'deleteLayoutFilterFromList',
                'updateLayoutFilterById',
            ]),

			onCreated() {
				this.fillLayoutFilterFilter();
	            this.fetchLayoutFilterData();
			},

            fillLayoutFilterFilter() {
                this.layoutFilterFilter.CurrentPage = 1;
                this.layoutFilterFilter.PerPage = 1000;
            },

            fetchLayoutFilterData() {
                return this.findLayoutFilter({
                    filter: this.layoutFilterFilter
                });
            },

            /**
             *
             * @param type
             */
            showPanel(type) {
                if (type === this.panel.create) {
                    this.panel.type = this.panel.create;
                    this.clearPanelLayoutFilterItem();
                } else if (type === this.panel.edit) {
                    this.panel.type = this.panel.edit;
                    this.currentLayoutFilterItem.isSelected = true;
                }

                this.panel.show = true;
            },

            closePanel() {
                this.panel.show = false;
                this.currentLayoutFilterItem.isSelected = false;
                this.clearPanelLayoutFilterItem();
            },

            selectLayoutFilterItem(layoutFilterItem) {
                this.showPanel(this.panel.edit);
                this.currentLayoutFilterItem.isSelected = true;
                Object.assign(this.currentLayoutFilterItem.item, layoutFilterItem);
            },

            changeCurrentLayoutFilterItem() {
                this.currentLayoutFilterItem.hasChange = true;
            },

            cancelChanges() {
                this.clearPanelLayoutFilterItem();
                this.closePanel();
            },

            clearPanelLayoutFilterItem() {
                this.currentLayoutFilterItem.item = new LayoutFilter();
                this.currentLayoutFilterItem.hasChange = false;
            },

            saveChangesSubmit() {
                if (this.isPanelCreate) {
                    this.createLayoutFilterItemSubmit();
                    return;
                }

                if (this.isPanelEdit) {
                    this.editLayoutFilterItemSubmit();
                }
            },

            createLayoutFilterItemSubmit() {
                return this.createLayoutFilter({
					data: this.currentLayoutFilterItem.item,
                }).then((response) => {

                    if (response.Model) {
                        this.addLayoutFilterItemToList(response.Model);
                        this.clearPanelLayoutFilterItem();
                    } else {
                        console.error('Ошибка создания записи: ', response.Error);
                    }

                }).catch(error => {
                    console.error('Ошибка создания записи: ', error);
                });
            },

            editLayoutFilterItemSubmit() {

                if (this.currentLayoutFilterItem.hasChange) {
                    return this.updateLayoutFilter({
                        id: this.currentLayoutFilterItem.item.Id,
                        data: this.currentLayoutFilterItem.item,
                    }).then((response) => {

                        if (response.Model) {
                            this.updateLayoutFilterById(response.Model);
                            this.currentLayoutFilterItem.hasChange = false;
                            this.clearPanelLayoutFilterItem();
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

            deleteLayoutFilterItemHandler() {
                let deletedItemId = this.currentLayoutFilterItem.item.Id;

                if (!this.currentLayoutFilterItem.canDelete) {
                    this.currentLayoutFilterItem.showDeleteConfirmation = true;
                    return;
                }

                this.deleteLayoutFilter({
                    id: deletedItemId
                }).then(response => {

                    if (response.IsSuccess) {
                        this.deleteLayoutFilterFromList(deletedItemId);
                        this.clearPanelLayoutFilterItem();
                        this.currentLayoutFilterItem.canDelete = false;
                        this.currentLayoutFilterItem.isSelected = false;
                        this.panel.show = false;
                    } else {
                        console.error('Ошибка удаления элемента: ', response.Error);
                    }

                }).catch(error => {
                    console.error('Ошибка удаления элемента: ', error);
                });
            },

            confirmDeleteHandler() {
                this.currentLayoutFilterItem.showDeleteConfirmation = false;
                this.currentLayoutFilterItem.canDelete = true;
                this.deleteLayoutFilterItemHandler();
            },

            closeConfirmationPanel() {
                this.currentLayoutFilterItem.showDeleteConfirmation = false;
            },
        },
    }
</script>

<style lang="scss">

</style>
