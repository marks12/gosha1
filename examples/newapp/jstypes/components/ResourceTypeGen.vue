
	<template>
    <WorkSpace>
        <template #header>
            <slot name="pageHeader">
                <VHead level="h1">ResourceType</VHead>
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
                            v-for="resourceTypeItem in resourceTypeList"
                            :key="resourceTypeItem.Id"
                            @click="selectResourceTypeItem(resourceTypeItem)"
                            class="sw-table__row_can-select"
                            :class="{'sw-table__row_is-selected': resourceTypeItem.Id === currentResourceTypeItem.item.Id}"
                        >
                            <td v-for="(value, key) in fields" :key="key + '-fields'">
                                <VCheckbox v-if="isCheckbox(resourceTypeItem[key])" :checked="resourceTypeItem[key]" disabled></VCheckbox>
                                <VText v-else>{{ resourceTypeItem[key] }}</VText>
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
                                        :for="`currentResourceTypeItem${key}`"
                                    >{{ filed }}</VLabel>
                                    <VInput
										v-if="isInput(currentResourceTypeItem.item[key])"
                                        v-model="currentResourceTypeItem.item[key]"
                                        width="dyn"
                                        :id="`currentResourceTypeItem${key}`"
                                        @input="changeCurrentResourceTypeItem"
                                    />
									<VCheckbox
										v-if="isCheckbox(currentResourceTypeItem.item[key])"
                                        v-model="currentResourceTypeItem.item[key]"
                                        :id="`currentResourceTypeItem${key}`"
										@input="changeCurrentResourceTypeItem"
									/>
									
                                </VSet>
                            </VSet>
                            <button type="submit" :disabled="!currentResourceTypeItem.hasChange" hidden></button>
                        </form>
                    </template>

                    <template #footer>
                        <VSet>
                            <VButton
                                @click="saveChangesSubmit"
                                accent
                                :text="panelSubmitButtonText"
                                :disabled="!currentResourceTypeItem.hasChange"
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
                    v-if="currentResourceTypeItem.showDeleteConfirmation"
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
                        :disabled="!currentResourceTypeItem.isSelected"
                        @click="deleteResourceTypeItemHandler"
                    />
                </VSet>
            </slot>
        </template>
    </WorkSpace>
</template>

<script>
    import resourceTypeData from "../data/ResourceTypeData";
    import { ResourceType } from '../apiModel';
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
        name: 'ResourceTypeGen',

        components: {VSelectSimple, VSign, VIcon, VButton, VPanel, VText, VInput, VLabel, VSet, VHead, WorkSpace, VCheckbox},

        props: {
            fields: {
                type: Object,
                default() {
                    const resourceTypeItem = new ResourceType();
                    const fieldsObj = {};

                    for (let prop in resourceTypeItem) {

                        if (resourceTypeItem.hasOwnProperty(prop)) {
                            fieldsObj[prop] = prop;
                        }

                    }

                    return fieldsObj;
                }
            },
            editFields: {
                type: Object,
                default() {
                    const resourceTypeItem = new ResourceType();
                    const fieldsObj = {};

                    for (let prop in resourceTypeItem) {

                        if (resourceTypeItem.hasOwnProperty(prop)) {
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
            return resourceTypeData;
        },

        created() {
            this.fillResourceTypeFilter();
            this.fetchResourceTypeData();
        },

        computed: {
            ...mapGetters({
                resourceTypeList: 'getListResourceType'
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
                'findResourceType',
                'updateResourceType',
                'deleteResourceType',
                'createResourceType',
            ]),

            ...mapMutations([
                'addResourceTypeItemToList',
                'deleteResourceTypeFromList',
                'updateResourceTypeById',
            ]),

            fillResourceTypeFilter() {
                this.resourceTypeFilter.CurrentPage = 1;
                this.resourceTypeFilter.PerPage = 1000;
            },

            fetchResourceTypeData() {
                return this.findResourceType({
                    filter: this.resourceTypeFilter
                });
            },

            /**
             *
             * @param type
             */
            showPanel(type) {
                if (type === this.panel.create) {
                    this.panel.type = this.panel.create;
                    this.clearPanelResourceTypeItem();
                } else if (type === this.panel.edit) {
                    this.panel.type = this.panel.edit;
                    this.currentResourceTypeItem.isSelected = true;
                }

                this.panel.show = true;
            },

            closePanel() {
                this.panel.show = false;
                this.currentResourceTypeItem.isSelected = false;
                this.clearPanelResourceTypeItem();
            },

            selectResourceTypeItem(resourceTypeItem) {
                this.showPanel(this.panel.edit);
                this.currentResourceTypeItem.isSelected = true;
                Object.assign(this.currentResourceTypeItem.item, resourceTypeItem);
            },

            changeCurrentResourceTypeItem() {
                this.currentResourceTypeItem.hasChange = true;
            },

            cancelChanges() {
                this.clearPanelResourceTypeItem();
                this.closePanel();
            },

            clearPanelResourceTypeItem() {
                this.currentResourceTypeItem.item = new ResourceType();
                this.currentResourceTypeItem.hasChange = false;
            },

            saveChangesSubmit() {
                if (this.isPanelCreate) {
                    this.createResourceTypeItemSubmit();
                    return;
                }

                if (this.isPanelEdit) {
                    this.editResourceTypeItemSubmit();
                }
            },

            createResourceTypeItemSubmit() {
                return this.createResourceType({
					data: this.currentResourceTypeItem.item,
                }).then((response) => {

                    if (response.Model) {
                        this.addResourceTypeItemToList(response.Model);
                        this.clearPanelResourceTypeItem();
                    } else {
                        console.error('Ошибка создания записи: ', response.Error);
                    }

                }).catch(error => {
                    console.error('Ошибка создания записи: ', error);
                });
            },

            editResourceTypeItemSubmit() {

                if (this.currentResourceTypeItem.hasChange) {
                    return this.updateResourceType({
                        id: this.currentResourceTypeItem.item.Id,
                        data: this.currentResourceTypeItem.item,
                    }).then((response) => {

                        if (response.Model) {
                            this.updateResourceTypeById(response.Model);
                            this.currentResourceTypeItem.hasChange = false;
                            this.clearPanelResourceTypeItem();
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

            deleteResourceTypeItemHandler() {
                let deletedItemId = this.currentResourceTypeItem.item.Id;

                if (!this.currentResourceTypeItem.canDelete) {
                    this.currentResourceTypeItem.showDeleteConfirmation = true;
                    return;
                }

                this.deleteResourceType({
                    id: deletedItemId
                }).then(response => {

                    if (response.IsSuccess) {
                        this.deleteResourceTypeFromList(deletedItemId);
                        this.clearPanelResourceTypeItem();
                        this.currentResourceTypeItem.canDelete = false;
                        this.currentResourceTypeItem.isSelected = false;
                        this.panel.show = false;
                    } else {
                        console.error('Ошибка удаления элемента: ', response.Error);
                    }

                }).catch(error => {
                    console.error('Ошибка удаления элемента: ', error);
                });
            },

            confirmDeleteHandler() {
                this.currentResourceTypeItem.showDeleteConfirmation = false;
                this.currentResourceTypeItem.canDelete = true;
                this.deleteResourceTypeItemHandler();
            },

            closeConfirmationPanel() {
                this.currentResourceTypeItem.showDeleteConfirmation = false;
            },
        },
    }
</script>

<style lang="scss">

</style>
