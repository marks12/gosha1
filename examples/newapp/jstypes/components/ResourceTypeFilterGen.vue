
	<template>
    <WorkSpace>
        <template #header>
            <slot name="pageHeader">
                <VHead level="h1">ResourceTypeFilter</VHead>
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
                            v-for="resourceTypeFilterItem in resourceTypeFilterList"
                            :key="resourceTypeFilterItem.Id"
                            @click="selectResourceTypeFilterItem(resourceTypeFilterItem)"
                            class="sw-table__row_can-select"
                            :class="{'sw-table__row_is-selected': resourceTypeFilterItem.Id === currentResourceTypeFilterItem.item.Id}"
                        >
                            <td v-for="(value, key) in fields" :key="key + '-fields'">
                                <VCheckbox v-if="isCheckbox(resourceTypeFilterItem[key])" :checked="resourceTypeFilterItem[key]" disabled></VCheckbox>
                                <VText v-else>{{ resourceTypeFilterItem[key] }}</VText>
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
                                        :for="`currentResourceTypeFilterItem${key}`"
                                    >{{ filed }}</VLabel>
                                    <VInput
										v-if="isInput(currentResourceTypeFilterItem.item[key])"
                                        v-model="currentResourceTypeFilterItem.item[key]"
                                        width="dyn"
                                        :id="`currentResourceTypeFilterItem${key}`"
                                        @input="changeCurrentResourceTypeFilterItem"
                                    />
									<VCheckbox
										v-if="isCheckbox(currentResourceTypeFilterItem.item[key])"
                                        v-model="currentResourceTypeFilterItem.item[key]"
                                        :id="`currentResourceTypeFilterItem${key}`"
										@input="changeCurrentResourceTypeFilterItem"
									/>
									
                                </VSet>
                            </VSet>
                            <button type="submit" :disabled="!currentResourceTypeFilterItem.hasChange" hidden></button>
                        </form>
                    </template>

                    <template #footer>
                        <VSet>
                            <VButton
                                @click="saveChangesSubmit"
                                accent
                                :text="panelSubmitButtonText"
                                :disabled="!currentResourceTypeFilterItem.hasChange"
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
                    v-if="currentResourceTypeFilterItem.showDeleteConfirmation"
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
                        :disabled="!currentResourceTypeFilterItem.isSelected"
                        @click="deleteResourceTypeFilterItemHandler"
                    />
                </VSet>
            </slot>
        </template>
    </WorkSpace>
</template>

<script>
    import resourceTypeFilterData from "../data/ResourceTypeFilterData";
    import { ResourceTypeFilter } from '../apiModel';
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
        name: 'ResourceTypeFilterGen',

        components: {VButton, VPanel, VText, VInput, VLabel, VSet, VHead, WorkSpace, VCheckbox},

        props: {
            fields: {
                type: Object,
                default() {
                    const resourceTypeFilterItem = new ResourceTypeFilter();
                    const fieldsObj = {};

                    for (let prop in resourceTypeFilterItem) {

                        if (resourceTypeFilterItem.hasOwnProperty(prop)) {
                            fieldsObj[prop] = prop;
                        }

                    }

                    return fieldsObj;
                }
            },
            editFields: {
                type: Object,
                default() {
                    const resourceTypeFilterItem = new ResourceTypeFilter();
                    const fieldsObj = {};

                    for (let prop in resourceTypeFilterItem) {

                        if (resourceTypeFilterItem.hasOwnProperty(prop)) {
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
            return resourceTypeFilterData;
        },

        created() {
			this.onCreated();
        },

        computed: {
            ...mapGetters({
                resourceTypeFilterList: 'getListResourceTypeFilter'
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
                'findResourceTypeFilter',
                'updateResourceTypeFilter',
                'deleteResourceTypeFilter',
                'createResourceTypeFilter',
            ]),

            ...mapMutations([
                'addResourceTypeFilterItemToList',
                'deleteResourceTypeFilterFromList',
                'updateResourceTypeFilterById',
            ]),

			onCreated() {
				this.fillResourceTypeFilterFilter();
	            this.fetchResourceTypeFilterData();
			},

            fillResourceTypeFilterFilter() {
                this.resourceTypeFilterFilter.CurrentPage = 1;
                this.resourceTypeFilterFilter.PerPage = 1000;
            },

            fetchResourceTypeFilterData() {
                return this.findResourceTypeFilter({
                    filter: this.resourceTypeFilterFilter
                });
            },

            /**
             *
             * @param type
             */
            showPanel(type) {
                if (type === this.panel.create) {
                    this.panel.type = this.panel.create;
                    this.clearPanelResourceTypeFilterItem();
                } else if (type === this.panel.edit) {
                    this.panel.type = this.panel.edit;
                    this.currentResourceTypeFilterItem.isSelected = true;
                }

                this.panel.show = true;
            },

            closePanel() {
                this.panel.show = false;
                this.currentResourceTypeFilterItem.isSelected = false;
                this.clearPanelResourceTypeFilterItem();
            },

            selectResourceTypeFilterItem(resourceTypeFilterItem) {
                this.showPanel(this.panel.edit);
                this.currentResourceTypeFilterItem.isSelected = true;
                Object.assign(this.currentResourceTypeFilterItem.item, resourceTypeFilterItem);
            },

            changeCurrentResourceTypeFilterItem() {
                this.currentResourceTypeFilterItem.hasChange = true;
            },

            cancelChanges() {
                this.clearPanelResourceTypeFilterItem();
                this.closePanel();
            },

            clearPanelResourceTypeFilterItem() {
                this.currentResourceTypeFilterItem.item = new ResourceTypeFilter();
                this.currentResourceTypeFilterItem.hasChange = false;
            },

            saveChangesSubmit() {
                if (this.isPanelCreate) {
                    this.createResourceTypeFilterItemSubmit();
                    return;
                }

                if (this.isPanelEdit) {
                    this.editResourceTypeFilterItemSubmit();
                }
            },

            createResourceTypeFilterItemSubmit() {
                return this.createResourceTypeFilter({
					data: this.currentResourceTypeFilterItem.item,
                }).then((response) => {

                    if (response.Model) {
                        this.addResourceTypeFilterItemToList(response.Model);
                        this.clearPanelResourceTypeFilterItem();
                    } else {
                        console.error('Ошибка создания записи: ', response.Error);
                    }

                }).catch(error => {
                    console.error('Ошибка создания записи: ', error);
                });
            },

            editResourceTypeFilterItemSubmit() {

                if (this.currentResourceTypeFilterItem.hasChange) {
                    return this.updateResourceTypeFilter({
                        id: this.currentResourceTypeFilterItem.item.Id,
                        data: this.currentResourceTypeFilterItem.item,
                    }).then((response) => {

                        if (response.Model) {
                            this.updateResourceTypeFilterById(response.Model);
                            this.currentResourceTypeFilterItem.hasChange = false;
                            this.clearPanelResourceTypeFilterItem();
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

            deleteResourceTypeFilterItemHandler() {
                let deletedItemId = this.currentResourceTypeFilterItem.item.Id;

                if (!this.currentResourceTypeFilterItem.canDelete) {
                    this.currentResourceTypeFilterItem.showDeleteConfirmation = true;
                    return;
                }

                this.deleteResourceTypeFilter({
                    id: deletedItemId
                }).then(response => {

                    if (response.IsSuccess) {
                        this.deleteResourceTypeFilterFromList(deletedItemId);
                        this.clearPanelResourceTypeFilterItem();
                        this.currentResourceTypeFilterItem.canDelete = false;
                        this.currentResourceTypeFilterItem.isSelected = false;
                        this.panel.show = false;
                    } else {
                        console.error('Ошибка удаления элемента: ', response.Error);
                    }

                }).catch(error => {
                    console.error('Ошибка удаления элемента: ', error);
                });
            },

            confirmDeleteHandler() {
                this.currentResourceTypeFilterItem.showDeleteConfirmation = false;
                this.currentResourceTypeFilterItem.canDelete = true;
                this.deleteResourceTypeFilterItemHandler();
            },

            closeConfirmationPanel() {
                this.currentResourceTypeFilterItem.showDeleteConfirmation = false;
            },
        },
    }
</script>

<style lang="scss">

</style>
