
	<template>
    <WorkSpace>
        <template #header>
            <slot name="pageHeader">
                <VHead level="h1">LayoutContentFilter</VHead>
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
                            v-for="layoutContentFilterItem in layoutContentFilterList"
                            :key="layoutContentFilterItem.Id"
                            @click="selectLayoutContentFilterItem(layoutContentFilterItem)"
                            class="sw-table__row_can-select"
                            :class="{'sw-table__row_is-selected': layoutContentFilterItem.Id === currentLayoutContentFilterItem.item.Id}"
                        >
                            <td v-for="(value, key) in fields" :key="key + '-fields'">
                                <VCheckbox v-if="isCheckbox(layoutContentFilterItem[key])" :checked="layoutContentFilterItem[key]" disabled></VCheckbox>
                                <VText v-else>{{ layoutContentFilterItem[key] }}</VText>
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
                                        :for="`currentLayoutContentFilterItem${key}`"
                                    >{{ filed }}</VLabel>
                                    <VInput
										v-if="isInput(currentLayoutContentFilterItem.item[key])"
                                        v-model="currentLayoutContentFilterItem.item[key]"
                                        width="dyn"
                                        :id="`currentLayoutContentFilterItem${key}`"
                                        @input="changeCurrentLayoutContentFilterItem"
                                    />
									<VCheckbox
										v-if="isCheckbox(currentLayoutContentFilterItem.item[key])"
                                        v-model="currentLayoutContentFilterItem.item[key]"
                                        :id="`currentLayoutContentFilterItem${key}`"
										@input="changeCurrentLayoutContentFilterItem"
									/>
									
                                </VSet>
                            </VSet>
                            <button type="submit" :disabled="!currentLayoutContentFilterItem.hasChange" hidden></button>
                        </form>
                    </template>

                    <template #footer>
                        <VSet>
                            <VButton
                                @click="saveChangesSubmit"
                                accent
                                :text="panelSubmitButtonText"
                                :disabled="!currentLayoutContentFilterItem.hasChange"
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
                    v-if="currentLayoutContentFilterItem.showDeleteConfirmation"
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
                        :disabled="!currentLayoutContentFilterItem.isSelected"
                        @click="deleteLayoutContentFilterItemHandler"
                    />
                </VSet>
            </slot>
        </template>
    </WorkSpace>
</template>

<script>
    import layoutContentFilterData from "../data/LayoutContentFilterData";
    import { LayoutContentFilter } from '../apiModel';
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
        name: 'LayoutContentFilterGen',

        components: {VButton, VPanel, VText, VInput, VLabel, VSet, VHead, WorkSpace, VCheckbox},

        props: {
            fields: {
                type: Object,
                default() {
                    const layoutContentFilterItem = new LayoutContentFilter();
                    const fieldsObj = {};

                    for (let prop in layoutContentFilterItem) {

                        if (layoutContentFilterItem.hasOwnProperty(prop)) {
                            fieldsObj[prop] = prop;
                        }

                    }

                    return fieldsObj;
                }
            },
            editFields: {
                type: Object,
                default() {
                    const layoutContentFilterItem = new LayoutContentFilter();
                    const fieldsObj = {};

                    for (let prop in layoutContentFilterItem) {

                        if (layoutContentFilterItem.hasOwnProperty(prop)) {
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
            return layoutContentFilterData;
        },

        created() {
			this.onCreated();
        },

        computed: {
            ...mapGetters({
                layoutContentFilterList: 'getListLayoutContentFilter'
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
                'findLayoutContentFilter',
                'updateLayoutContentFilter',
                'deleteLayoutContentFilter',
                'createLayoutContentFilter',
            ]),

            ...mapMutations([
                'addLayoutContentFilterItemToList',
                'deleteLayoutContentFilterFromList',
                'updateLayoutContentFilterById',
            ]),

			onCreated() {
				this.fillLayoutContentFilterFilter();
	            this.fetchLayoutContentFilterData();
			},

            fillLayoutContentFilterFilter() {
                this.layoutContentFilterFilter.CurrentPage = 1;
                this.layoutContentFilterFilter.PerPage = 1000;
            },

            fetchLayoutContentFilterData() {
                return this.findLayoutContentFilter({
                    filter: this.layoutContentFilterFilter
                });
            },

            /**
             *
             * @param type
             */
            showPanel(type) {
                if (type === this.panel.create) {
                    this.panel.type = this.panel.create;
                    this.clearPanelLayoutContentFilterItem();
                } else if (type === this.panel.edit) {
                    this.panel.type = this.panel.edit;
                    this.currentLayoutContentFilterItem.isSelected = true;
                }

                this.panel.show = true;
            },

            closePanel() {
                this.panel.show = false;
                this.currentLayoutContentFilterItem.isSelected = false;
                this.clearPanelLayoutContentFilterItem();
            },

            selectLayoutContentFilterItem(layoutContentFilterItem) {
                this.showPanel(this.panel.edit);
                this.currentLayoutContentFilterItem.isSelected = true;
                Object.assign(this.currentLayoutContentFilterItem.item, layoutContentFilterItem);
            },

            changeCurrentLayoutContentFilterItem() {
                this.currentLayoutContentFilterItem.hasChange = true;
            },

            cancelChanges() {
                this.clearPanelLayoutContentFilterItem();
                this.closePanel();
            },

            clearPanelLayoutContentFilterItem() {
                this.currentLayoutContentFilterItem.item = new LayoutContentFilter();
                this.currentLayoutContentFilterItem.hasChange = false;
            },

            saveChangesSubmit() {
                if (this.isPanelCreate) {
                    this.createLayoutContentFilterItemSubmit();
                    return;
                }

                if (this.isPanelEdit) {
                    this.editLayoutContentFilterItemSubmit();
                }
            },

            createLayoutContentFilterItemSubmit() {
                return this.createLayoutContentFilter({
					data: this.currentLayoutContentFilterItem.item,
                }).then((response) => {

                    if (response.Model) {
                        this.addLayoutContentFilterItemToList(response.Model);
                        this.clearPanelLayoutContentFilterItem();
                    } else {
                        console.error('Ошибка создания записи: ', response.Error);
                    }

                }).catch(error => {
                    console.error('Ошибка создания записи: ', error);
                });
            },

            editLayoutContentFilterItemSubmit() {

                if (this.currentLayoutContentFilterItem.hasChange) {
                    return this.updateLayoutContentFilter({
                        id: this.currentLayoutContentFilterItem.item.Id,
                        data: this.currentLayoutContentFilterItem.item,
                    }).then((response) => {

                        if (response.Model) {
                            this.updateLayoutContentFilterById(response.Model);
                            this.currentLayoutContentFilterItem.hasChange = false;
                            this.clearPanelLayoutContentFilterItem();
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

            deleteLayoutContentFilterItemHandler() {
                let deletedItemId = this.currentLayoutContentFilterItem.item.Id;

                if (!this.currentLayoutContentFilterItem.canDelete) {
                    this.currentLayoutContentFilterItem.showDeleteConfirmation = true;
                    return;
                }

                this.deleteLayoutContentFilter({
                    id: deletedItemId
                }).then(response => {

                    if (response.IsSuccess) {
                        this.deleteLayoutContentFilterFromList(deletedItemId);
                        this.clearPanelLayoutContentFilterItem();
                        this.currentLayoutContentFilterItem.canDelete = false;
                        this.currentLayoutContentFilterItem.isSelected = false;
                        this.panel.show = false;
                    } else {
                        console.error('Ошибка удаления элемента: ', response.Error);
                    }

                }).catch(error => {
                    console.error('Ошибка удаления элемента: ', error);
                });
            },

            confirmDeleteHandler() {
                this.currentLayoutContentFilterItem.showDeleteConfirmation = false;
                this.currentLayoutContentFilterItem.canDelete = true;
                this.deleteLayoutContentFilterItemHandler();
            },

            closeConfirmationPanel() {
                this.currentLayoutContentFilterItem.showDeleteConfirmation = false;
            },
        },
    }
</script>

<style lang="scss">

</style>
