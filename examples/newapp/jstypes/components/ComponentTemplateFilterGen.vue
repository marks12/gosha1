
	<template>
    <WorkSpace>
        <template #header>
            <slot name="pageHeader">
                <VHead level="h1">ComponentTemplateFilter</VHead>
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
                            v-for="componentTemplateFilterItem in componentTemplateFilterList"
                            :key="componentTemplateFilterItem.Id"
                            @click="selectComponentTemplateFilterItem(componentTemplateFilterItem)"
                            class="sw-table__row_can-select"
                            :class="{'sw-table__row_is-selected': componentTemplateFilterItem.Id === currentComponentTemplateFilterItem.item.Id}"
                        >
                            <td v-for="(value, key) in fields" :key="key + '-fields'">
                                <VCheckbox v-if="isCheckbox(componentTemplateFilterItem[key])" :checked="componentTemplateFilterItem[key]" disabled></VCheckbox>
                                <VText v-else>{{ componentTemplateFilterItem[key] }}</VText>
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
                                        :for="`currentComponentTemplateFilterItem${key}`"
                                    >{{ filed }}</VLabel>
                                    <VInput
										v-if="isInput(currentComponentTemplateFilterItem.item[key])"
                                        v-model="currentComponentTemplateFilterItem.item[key]"
                                        width="dyn"
                                        :id="`currentComponentTemplateFilterItem${key}`"
                                        @input="changeCurrentComponentTemplateFilterItem"
                                    />
									<VCheckbox
										v-if="isCheckbox(currentComponentTemplateFilterItem.item[key])"
                                        v-model="currentComponentTemplateFilterItem.item[key]"
                                        :id="`currentComponentTemplateFilterItem${key}`"
										@input="changeCurrentComponentTemplateFilterItem"
									/>
									
                                </VSet>
                            </VSet>
                            <button type="submit" :disabled="!currentComponentTemplateFilterItem.hasChange" hidden></button>
                        </form>
                    </template>

                    <template #footer>
                        <VSet>
                            <VButton
                                @click="saveChangesSubmit"
                                accent
                                :text="panelSubmitButtonText"
                                :disabled="!currentComponentTemplateFilterItem.hasChange"
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
                    v-if="currentComponentTemplateFilterItem.showDeleteConfirmation"
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
                        :disabled="!currentComponentTemplateFilterItem.isSelected"
                        @click="deleteComponentTemplateFilterItemHandler"
                    />
                </VSet>
            </slot>
        </template>
    </WorkSpace>
</template>

<script>
    import componentTemplateFilterData from "../data/ComponentTemplateFilterData";
    import { ComponentTemplateFilter } from '../apiModel';
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
        name: 'ComponentTemplateFilterGen',

        components: {VButton, VPanel, VText, VInput, VLabel, VSet, VHead, WorkSpace, VCheckbox},

        props: {
            fields: {
                type: Object,
                default() {
                    const componentTemplateFilterItem = new ComponentTemplateFilter();
                    const fieldsObj = {};

                    for (let prop in componentTemplateFilterItem) {

                        if (componentTemplateFilterItem.hasOwnProperty(prop)) {
                            fieldsObj[prop] = prop;
                        }

                    }

                    return fieldsObj;
                }
            },
            editFields: {
                type: Object,
                default() {
                    const componentTemplateFilterItem = new ComponentTemplateFilter();
                    const fieldsObj = {};

                    for (let prop in componentTemplateFilterItem) {

                        if (componentTemplateFilterItem.hasOwnProperty(prop)) {
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
            return componentTemplateFilterData;
        },

        created() {
			this.onCreated();
        },

        computed: {
            ...mapGetters({
                componentTemplateFilterList: 'getListComponentTemplateFilter'
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
                'findComponentTemplateFilter',
                'updateComponentTemplateFilter',
                'deleteComponentTemplateFilter',
                'createComponentTemplateFilter',
            ]),

            ...mapMutations([
                'addComponentTemplateFilterItemToList',
                'deleteComponentTemplateFilterFromList',
                'updateComponentTemplateFilterById',
            ]),

			onCreated() {
				this.fillComponentTemplateFilterFilter();
	            this.fetchComponentTemplateFilterData();
			},

            fillComponentTemplateFilterFilter() {
                this.componentTemplateFilterFilter.CurrentPage = 1;
                this.componentTemplateFilterFilter.PerPage = 1000;
            },

            fetchComponentTemplateFilterData() {
                return this.findComponentTemplateFilter({
                    filter: this.componentTemplateFilterFilter
                });
            },

            /**
             *
             * @param type
             */
            showPanel(type) {
                if (type === this.panel.create) {
                    this.panel.type = this.panel.create;
                    this.clearPanelComponentTemplateFilterItem();
                } else if (type === this.panel.edit) {
                    this.panel.type = this.panel.edit;
                    this.currentComponentTemplateFilterItem.isSelected = true;
                }

                this.panel.show = true;
            },

            closePanel() {
                this.panel.show = false;
                this.currentComponentTemplateFilterItem.isSelected = false;
                this.clearPanelComponentTemplateFilterItem();
            },

            selectComponentTemplateFilterItem(componentTemplateFilterItem) {
                this.showPanel(this.panel.edit);
                this.currentComponentTemplateFilterItem.isSelected = true;
                Object.assign(this.currentComponentTemplateFilterItem.item, componentTemplateFilterItem);
            },

            changeCurrentComponentTemplateFilterItem() {
                this.currentComponentTemplateFilterItem.hasChange = true;
            },

            cancelChanges() {
                this.clearPanelComponentTemplateFilterItem();
                this.closePanel();
            },

            clearPanelComponentTemplateFilterItem() {
                this.currentComponentTemplateFilterItem.item = new ComponentTemplateFilter();
                this.currentComponentTemplateFilterItem.hasChange = false;
            },

            saveChangesSubmit() {
                if (this.isPanelCreate) {
                    this.createComponentTemplateFilterItemSubmit();
                    return;
                }

                if (this.isPanelEdit) {
                    this.editComponentTemplateFilterItemSubmit();
                }
            },

            createComponentTemplateFilterItemSubmit() {
                return this.createComponentTemplateFilter({
					data: this.currentComponentTemplateFilterItem.item,
                }).then((response) => {

                    if (response.Model) {
                        this.addComponentTemplateFilterItemToList(response.Model);
                        this.clearPanelComponentTemplateFilterItem();
                    } else {
                        console.error('Ошибка создания записи: ', response.Error);
                    }

                }).catch(error => {
                    console.error('Ошибка создания записи: ', error);
                });
            },

            editComponentTemplateFilterItemSubmit() {

                if (this.currentComponentTemplateFilterItem.hasChange) {
                    return this.updateComponentTemplateFilter({
                        id: this.currentComponentTemplateFilterItem.item.Id,
                        data: this.currentComponentTemplateFilterItem.item,
                    }).then((response) => {

                        if (response.Model) {
                            this.updateComponentTemplateFilterById(response.Model);
                            this.currentComponentTemplateFilterItem.hasChange = false;
                            this.clearPanelComponentTemplateFilterItem();
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

            deleteComponentTemplateFilterItemHandler() {
                let deletedItemId = this.currentComponentTemplateFilterItem.item.Id;

                if (!this.currentComponentTemplateFilterItem.canDelete) {
                    this.currentComponentTemplateFilterItem.showDeleteConfirmation = true;
                    return;
                }

                this.deleteComponentTemplateFilter({
                    id: deletedItemId
                }).then(response => {

                    if (response.IsSuccess) {
                        this.deleteComponentTemplateFilterFromList(deletedItemId);
                        this.clearPanelComponentTemplateFilterItem();
                        this.currentComponentTemplateFilterItem.canDelete = false;
                        this.currentComponentTemplateFilterItem.isSelected = false;
                        this.panel.show = false;
                    } else {
                        console.error('Ошибка удаления элемента: ', response.Error);
                    }

                }).catch(error => {
                    console.error('Ошибка удаления элемента: ', error);
                });
            },

            confirmDeleteHandler() {
                this.currentComponentTemplateFilterItem.showDeleteConfirmation = false;
                this.currentComponentTemplateFilterItem.canDelete = true;
                this.deleteComponentTemplateFilterItemHandler();
            },

            closeConfirmationPanel() {
                this.currentComponentTemplateFilterItem.showDeleteConfirmation = false;
            },
        },
    }
</script>

<style lang="scss">

</style>
