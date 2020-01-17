
	<template>
    <WorkSpace>
        <template #header>
            <slot name="pageHeader">
                <VHead level="h1">ComponentGroupFilter</VHead>
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
                            v-for="componentGroupFilterItem in componentGroupFilterList"
                            :key="componentGroupFilterItem.Id"
                            @click="selectComponentGroupFilterItem(componentGroupFilterItem)"
                            class="sw-table__row_can-select"
                            :class="{'sw-table__row_is-selected': componentGroupFilterItem.Id === currentComponentGroupFilterItem.item.Id}"
                        >
                            <td v-for="(value, key) in fields" :key="key + '-fields'">
                                <VCheckbox v-if="isCheckbox(componentGroupFilterItem[key])" :checked="componentGroupFilterItem[key]" disabled></VCheckbox>
                                <VText v-else>{{ componentGroupFilterItem[key] }}</VText>
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
                                        :for="`currentComponentGroupFilterItem${key}`"
                                    >{{ filed }}</VLabel>
                                    <VInput
										v-if="isInput(currentComponentGroupFilterItem.item[key])"
                                        v-model="currentComponentGroupFilterItem.item[key]"
                                        width="dyn"
                                        :id="`currentComponentGroupFilterItem${key}`"
                                        @input="changeCurrentComponentGroupFilterItem"
                                    />
									<VCheckbox
										v-if="isCheckbox(currentComponentGroupFilterItem.item[key])"
                                        v-model="currentComponentGroupFilterItem.item[key]"
                                        :id="`currentComponentGroupFilterItem${key}`"
										@input="changeCurrentComponentGroupFilterItem"
									/>
									
                                </VSet>
                            </VSet>
                            <button type="submit" :disabled="!currentComponentGroupFilterItem.hasChange" hidden></button>
                        </form>
                    </template>

                    <template #footer>
                        <VSet>
                            <VButton
                                @click="saveChangesSubmit"
                                accent
                                :text="panelSubmitButtonText"
                                :disabled="!currentComponentGroupFilterItem.hasChange"
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
                    v-if="currentComponentGroupFilterItem.showDeleteConfirmation"
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
                        :disabled="!currentComponentGroupFilterItem.isSelected"
                        @click="deleteComponentGroupFilterItemHandler"
                    />
                </VSet>
            </slot>
        </template>
    </WorkSpace>
</template>

<script>
    import componentGroupFilterData from "../data/ComponentGroupFilterData";
    import { ComponentGroupFilter } from '../apiModel';
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
        name: 'ComponentGroupFilterGen',

        components: {VButton, VPanel, VText, VInput, VLabel, VSet, VHead, WorkSpace, VCheckbox},

        props: {
            fields: {
                type: Object,
                default() {
                    const componentGroupFilterItem = new ComponentGroupFilter();
                    const fieldsObj = {};

                    for (let prop in componentGroupFilterItem) {

                        if (componentGroupFilterItem.hasOwnProperty(prop)) {
                            fieldsObj[prop] = prop;
                        }

                    }

                    return fieldsObj;
                }
            },
            editFields: {
                type: Object,
                default() {
                    const componentGroupFilterItem = new ComponentGroupFilter();
                    const fieldsObj = {};

                    for (let prop in componentGroupFilterItem) {

                        if (componentGroupFilterItem.hasOwnProperty(prop)) {
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
            return componentGroupFilterData;
        },

        created() {
			this.onCreated();
        },

        computed: {
            ...mapGetters({
                componentGroupFilterList: 'getListComponentGroupFilter'
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
                'findComponentGroupFilter',
                'updateComponentGroupFilter',
                'deleteComponentGroupFilter',
                'createComponentGroupFilter',
            ]),

            ...mapMutations([
                'addComponentGroupFilterItemToList',
                'deleteComponentGroupFilterFromList',
                'updateComponentGroupFilterById',
            ]),

			onCreated() {
				this.fillComponentGroupFilterFilter();
	            this.fetchComponentGroupFilterData();
			},

            fillComponentGroupFilterFilter() {
                this.componentGroupFilterFilter.CurrentPage = 1;
                this.componentGroupFilterFilter.PerPage = 1000;
            },

            fetchComponentGroupFilterData() {
                return this.findComponentGroupFilter({
                    filter: this.componentGroupFilterFilter
                });
            },

            /**
             *
             * @param type
             */
            showPanel(type) {
                if (type === this.panel.create) {
                    this.panel.type = this.panel.create;
                    this.clearPanelComponentGroupFilterItem();
                } else if (type === this.panel.edit) {
                    this.panel.type = this.panel.edit;
                    this.currentComponentGroupFilterItem.isSelected = true;
                }

                this.panel.show = true;
            },

            closePanel() {
                this.panel.show = false;
                this.currentComponentGroupFilterItem.isSelected = false;
                this.clearPanelComponentGroupFilterItem();
            },

            selectComponentGroupFilterItem(componentGroupFilterItem) {
                this.showPanel(this.panel.edit);
                this.currentComponentGroupFilterItem.isSelected = true;
                Object.assign(this.currentComponentGroupFilterItem.item, componentGroupFilterItem);
            },

            changeCurrentComponentGroupFilterItem() {
                this.currentComponentGroupFilterItem.hasChange = true;
            },

            cancelChanges() {
                this.clearPanelComponentGroupFilterItem();
                this.closePanel();
            },

            clearPanelComponentGroupFilterItem() {
                this.currentComponentGroupFilterItem.item = new ComponentGroupFilter();
                this.currentComponentGroupFilterItem.hasChange = false;
            },

            saveChangesSubmit() {
                if (this.isPanelCreate) {
                    this.createComponentGroupFilterItemSubmit();
                    return;
                }

                if (this.isPanelEdit) {
                    this.editComponentGroupFilterItemSubmit();
                }
            },

            createComponentGroupFilterItemSubmit() {
                return this.createComponentGroupFilter({
					data: this.currentComponentGroupFilterItem.item,
                }).then((response) => {

                    if (response.Model) {
                        this.addComponentGroupFilterItemToList(response.Model);
                        this.clearPanelComponentGroupFilterItem();
                    } else {
                        console.error('Ошибка создания записи: ', response.Error);
                    }

                }).catch(error => {
                    console.error('Ошибка создания записи: ', error);
                });
            },

            editComponentGroupFilterItemSubmit() {

                if (this.currentComponentGroupFilterItem.hasChange) {
                    return this.updateComponentGroupFilter({
                        id: this.currentComponentGroupFilterItem.item.Id,
                        data: this.currentComponentGroupFilterItem.item,
                    }).then((response) => {

                        if (response.Model) {
                            this.updateComponentGroupFilterById(response.Model);
                            this.currentComponentGroupFilterItem.hasChange = false;
                            this.clearPanelComponentGroupFilterItem();
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

            deleteComponentGroupFilterItemHandler() {
                let deletedItemId = this.currentComponentGroupFilterItem.item.Id;

                if (!this.currentComponentGroupFilterItem.canDelete) {
                    this.currentComponentGroupFilterItem.showDeleteConfirmation = true;
                    return;
                }

                this.deleteComponentGroupFilter({
                    id: deletedItemId
                }).then(response => {

                    if (response.IsSuccess) {
                        this.deleteComponentGroupFilterFromList(deletedItemId);
                        this.clearPanelComponentGroupFilterItem();
                        this.currentComponentGroupFilterItem.canDelete = false;
                        this.currentComponentGroupFilterItem.isSelected = false;
                        this.panel.show = false;
                    } else {
                        console.error('Ошибка удаления элемента: ', response.Error);
                    }

                }).catch(error => {
                    console.error('Ошибка удаления элемента: ', error);
                });
            },

            confirmDeleteHandler() {
                this.currentComponentGroupFilterItem.showDeleteConfirmation = false;
                this.currentComponentGroupFilterItem.canDelete = true;
                this.deleteComponentGroupFilterItemHandler();
            },

            closeConfirmationPanel() {
                this.currentComponentGroupFilterItem.showDeleteConfirmation = false;
            },
        },
    }
</script>

<style lang="scss">

</style>
