
	<template>
    <WorkSpace>
        <template #header>
            <slot name="pageHeader">
                <VHead level="h1">ResourceFilter</VHead>
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
                            v-for="resourceFilterItem in resourceFilterList"
                            :key="resourceFilterItem.Id"
                            @click="selectResourceFilterItem(resourceFilterItem)"
                            class="sw-table__row_can-select"
                            :class="{'sw-table__row_is-selected': resourceFilterItem.Id === currentResourceFilterItem.item.Id}"
                        >
                            <td v-for="(value, key) in fields" :key="key + '-fields'">
                                <VCheckbox v-if="isCheckbox(resourceFilterItem[key])" :checked="resourceFilterItem[key]" disabled></VCheckbox>
                                <VText v-else>{{ resourceFilterItem[key] }}</VText>
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
                                        :for="`currentResourceFilterItem${key}`"
                                    >{{ filed }}</VLabel>
                                    <VInput
										v-if="isInput(currentResourceFilterItem.item[key])"
                                        v-model="currentResourceFilterItem.item[key]"
                                        width="dyn"
                                        :id="`currentResourceFilterItem${key}`"
                                        @input="changeCurrentResourceFilterItem"
                                    />
									<VCheckbox
										v-if="isCheckbox(currentResourceFilterItem.item[key])"
                                        v-model="currentResourceFilterItem.item[key]"
                                        :id="`currentResourceFilterItem${key}`"
										@input="changeCurrentResourceFilterItem"
									/>
									
                                </VSet>
                            </VSet>
                            <button type="submit" :disabled="!currentResourceFilterItem.hasChange" hidden></button>
                        </form>
                    </template>

                    <template #footer>
                        <VSet>
                            <VButton
                                @click="saveChangesSubmit"
                                accent
                                :text="panelSubmitButtonText"
                                :disabled="!currentResourceFilterItem.hasChange"
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
                    v-if="currentResourceFilterItem.showDeleteConfirmation"
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
                        :disabled="!currentResourceFilterItem.isSelected"
                        @click="deleteResourceFilterItemHandler"
                    />
                </VSet>
            </slot>
        </template>
    </WorkSpace>
</template>

<script>
    import resourceFilterData from "../data/ResourceFilterData";
    import { ResourceFilter } from '../apiModel';
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
        name: 'ResourceFilterGen',

        components: {VButton, VPanel, VText, VInput, VLabel, VSet, VHead, WorkSpace, VCheckbox},

        props: {
            fields: {
                type: Object,
                default() {
                    const resourceFilterItem = new ResourceFilter();
                    const fieldsObj = {};

                    for (let prop in resourceFilterItem) {

                        if (resourceFilterItem.hasOwnProperty(prop)) {
                            fieldsObj[prop] = prop;
                        }

                    }

                    return fieldsObj;
                }
            },
            editFields: {
                type: Object,
                default() {
                    const resourceFilterItem = new ResourceFilter();
                    const fieldsObj = {};

                    for (let prop in resourceFilterItem) {

                        if (resourceFilterItem.hasOwnProperty(prop)) {
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
            return resourceFilterData;
        },

        created() {
			this.onCreated();
        },

        computed: {
            ...mapGetters({
                resourceFilterList: 'getListResourceFilter'
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
                'findResourceFilter',
                'updateResourceFilter',
                'deleteResourceFilter',
                'createResourceFilter',
            ]),

            ...mapMutations([
                'addResourceFilterItemToList',
                'deleteResourceFilterFromList',
                'updateResourceFilterById',
            ]),

			onCreated() {
				this.fillResourceFilterFilter();
	            this.fetchResourceFilterData();
			},

            fillResourceFilterFilter() {
                this.resourceFilterFilter.CurrentPage = 1;
                this.resourceFilterFilter.PerPage = 1000;
            },

            fetchResourceFilterData() {
                return this.findResourceFilter({
                    filter: this.resourceFilterFilter
                });
            },

            /**
             *
             * @param type
             */
            showPanel(type) {
                if (type === this.panel.create) {
                    this.panel.type = this.panel.create;
                    this.clearPanelResourceFilterItem();
                } else if (type === this.panel.edit) {
                    this.panel.type = this.panel.edit;
                    this.currentResourceFilterItem.isSelected = true;
                }

                this.panel.show = true;
            },

            closePanel() {
                this.panel.show = false;
                this.currentResourceFilterItem.isSelected = false;
                this.clearPanelResourceFilterItem();
            },

            selectResourceFilterItem(resourceFilterItem) {
                this.showPanel(this.panel.edit);
                this.currentResourceFilterItem.isSelected = true;
                Object.assign(this.currentResourceFilterItem.item, resourceFilterItem);
            },

            changeCurrentResourceFilterItem() {
                this.currentResourceFilterItem.hasChange = true;
            },

            cancelChanges() {
                this.clearPanelResourceFilterItem();
                this.closePanel();
            },

            clearPanelResourceFilterItem() {
                this.currentResourceFilterItem.item = new ResourceFilter();
                this.currentResourceFilterItem.hasChange = false;
            },

            saveChangesSubmit() {
                if (this.isPanelCreate) {
                    this.createResourceFilterItemSubmit();
                    return;
                }

                if (this.isPanelEdit) {
                    this.editResourceFilterItemSubmit();
                }
            },

            createResourceFilterItemSubmit() {
                return this.createResourceFilter({
					data: this.currentResourceFilterItem.item,
                }).then((response) => {

                    if (response.Model) {
                        this.addResourceFilterItemToList(response.Model);
                        this.clearPanelResourceFilterItem();
                    } else {
                        console.error('Ошибка создания записи: ', response.Error);
                    }

                }).catch(error => {
                    console.error('Ошибка создания записи: ', error);
                });
            },

            editResourceFilterItemSubmit() {

                if (this.currentResourceFilterItem.hasChange) {
                    return this.updateResourceFilter({
                        id: this.currentResourceFilterItem.item.Id,
                        data: this.currentResourceFilterItem.item,
                    }).then((response) => {

                        if (response.Model) {
                            this.updateResourceFilterById(response.Model);
                            this.currentResourceFilterItem.hasChange = false;
                            this.clearPanelResourceFilterItem();
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

            deleteResourceFilterItemHandler() {
                let deletedItemId = this.currentResourceFilterItem.item.Id;

                if (!this.currentResourceFilterItem.canDelete) {
                    this.currentResourceFilterItem.showDeleteConfirmation = true;
                    return;
                }

                this.deleteResourceFilter({
                    id: deletedItemId
                }).then(response => {

                    if (response.IsSuccess) {
                        this.deleteResourceFilterFromList(deletedItemId);
                        this.clearPanelResourceFilterItem();
                        this.currentResourceFilterItem.canDelete = false;
                        this.currentResourceFilterItem.isSelected = false;
                        this.panel.show = false;
                    } else {
                        console.error('Ошибка удаления элемента: ', response.Error);
                    }

                }).catch(error => {
                    console.error('Ошибка удаления элемента: ', error);
                });
            },

            confirmDeleteHandler() {
                this.currentResourceFilterItem.showDeleteConfirmation = false;
                this.currentResourceFilterItem.canDelete = true;
                this.deleteResourceFilterItemHandler();
            },

            closeConfirmationPanel() {
                this.currentResourceFilterItem.showDeleteConfirmation = false;
            },
        },
    }
</script>

<style lang="scss">

</style>
