
	<template>
    <WorkSpace>
        <template #header>
            <slot name="pageHeader">
                <VHead level="h1">Resource</VHead>
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
                            v-for="resourceItem in resourceList"
                            :key="resourceItem.Id"
                            @click="selectResourceItem(resourceItem)"
                            class="sw-table__row_can-select"
                            :class="{'sw-table__row_is-selected': resourceItem.Id === currentResourceItem.item.Id}"
                        >
                            <td v-for="(value, key) in fields" :key="key + '-fields'">
                                <VCheckbox v-if="isCheckbox(resourceItem[key])" :checked="resourceItem[key]" disabled></VCheckbox>
                                <VText v-else>{{ resourceItem[key] }}</VText>
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
                                        :for="`currentResourceItem${key}`"
                                    >{{ filed }}</VLabel>
                                    <VInput
										v-if="isInput(currentResourceItem.item[key])"
                                        v-model="currentResourceItem.item[key]"
                                        width="dyn"
                                        :id="`currentResourceItem${key}`"
                                        @input="changeCurrentResourceItem"
                                    />
									<VCheckbox
										v-if="isCheckbox(currentResourceItem.item[key])"
                                        v-model="currentResourceItem.item[key]"
                                        :id="`currentResourceItem${key}`"
										@input="changeCurrentResourceItem"
									/>
									
                                </VSet>
                            </VSet>
                            <button type="submit" :disabled="!currentResourceItem.hasChange" hidden></button>
                        </form>
                    </template>

                    <template #footer>
                        <VSet>
                            <VButton
                                @click="saveChangesSubmit"
                                accent
                                :text="panelSubmitButtonText"
                                :disabled="!currentResourceItem.hasChange"
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
                    v-if="currentResourceItem.showDeleteConfirmation"
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
                        :disabled="!currentResourceItem.isSelected"
                        @click="deleteResourceItemHandler"
                    />
                </VSet>
            </slot>
        </template>
    </WorkSpace>
</template>

<script>
    import resourceData from "../data/ResourceData";
    import { Resource } from '../apiModel';
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
        name: 'ResourceGen',

        components: {VButton, VPanel, VText, VInput, VLabel, VSet, VHead, WorkSpace, VCheckbox},

        props: {
            fields: {
                type: Object,
                default() {
                    const resourceItem = new Resource();
                    const fieldsObj = {};

                    for (let prop in resourceItem) {

                        if (resourceItem.hasOwnProperty(prop)) {
                            fieldsObj[prop] = prop;
                        }

                    }

                    return fieldsObj;
                }
            },
            editFields: {
                type: Object,
                default() {
                    const resourceItem = new Resource();
                    const fieldsObj = {};

                    for (let prop in resourceItem) {

                        if (resourceItem.hasOwnProperty(prop)) {
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
            return resourceData;
        },

        created() {
			this.onCreated();
        },

        computed: {
            ...mapGetters({
                resourceList: 'getListResource'
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
                'findResource',
                'updateResource',
                'deleteResource',
                'createResource',
            ]),

            ...mapMutations([
                'addResourceItemToList',
                'deleteResourceFromList',
                'updateResourceById',
            ]),

			onCreated() {
				this.fillResourceFilter();
	            this.fetchResourceData();
			},

            fillResourceFilter() {
                this.resourceFilter.CurrentPage = 1;
                this.resourceFilter.PerPage = 1000;
            },

            fetchResourceData() {
                return this.findResource({
                    filter: this.resourceFilter
                });
            },

            /**
             *
             * @param type
             */
            showPanel(type) {
                if (type === this.panel.create) {
                    this.panel.type = this.panel.create;
                    this.clearPanelResourceItem();
                } else if (type === this.panel.edit) {
                    this.panel.type = this.panel.edit;
                    this.currentResourceItem.isSelected = true;
                }

                this.panel.show = true;
            },

            closePanel() {
                this.panel.show = false;
                this.currentResourceItem.isSelected = false;
                this.clearPanelResourceItem();
            },

            selectResourceItem(resourceItem) {
                this.showPanel(this.panel.edit);
                this.currentResourceItem.isSelected = true;
                Object.assign(this.currentResourceItem.item, resourceItem);
            },

            changeCurrentResourceItem() {
                this.currentResourceItem.hasChange = true;
            },

            cancelChanges() {
                this.clearPanelResourceItem();
                this.closePanel();
            },

            clearPanelResourceItem() {
                this.currentResourceItem.item = new Resource();
                this.currentResourceItem.hasChange = false;
            },

            saveChangesSubmit() {
                if (this.isPanelCreate) {
                    this.createResourceItemSubmit();
                    return;
                }

                if (this.isPanelEdit) {
                    this.editResourceItemSubmit();
                }
            },

            createResourceItemSubmit() {
                return this.createResource({
					data: this.currentResourceItem.item,
                }).then((response) => {

                    if (response.Model) {
                        this.addResourceItemToList(response.Model);
                        this.clearPanelResourceItem();
                    } else {
                        console.error('Ошибка создания записи: ', response.Error);
                    }

                }).catch(error => {
                    console.error('Ошибка создания записи: ', error);
                });
            },

            editResourceItemSubmit() {

                if (this.currentResourceItem.hasChange) {
                    return this.updateResource({
                        id: this.currentResourceItem.item.Id,
                        data: this.currentResourceItem.item,
                    }).then((response) => {

                        if (response.Model) {
                            this.updateResourceById(response.Model);
                            this.currentResourceItem.hasChange = false;
                            this.clearPanelResourceItem();
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

            deleteResourceItemHandler() {
                let deletedItemId = this.currentResourceItem.item.Id;

                if (!this.currentResourceItem.canDelete) {
                    this.currentResourceItem.showDeleteConfirmation = true;
                    return;
                }

                this.deleteResource({
                    id: deletedItemId
                }).then(response => {

                    if (response.IsSuccess) {
                        this.deleteResourceFromList(deletedItemId);
                        this.clearPanelResourceItem();
                        this.currentResourceItem.canDelete = false;
                        this.currentResourceItem.isSelected = false;
                        this.panel.show = false;
                    } else {
                        console.error('Ошибка удаления элемента: ', response.Error);
                    }

                }).catch(error => {
                    console.error('Ошибка удаления элемента: ', error);
                });
            },

            confirmDeleteHandler() {
                this.currentResourceItem.showDeleteConfirmation = false;
                this.currentResourceItem.canDelete = true;
                this.deleteResourceItemHandler();
            },

            closeConfirmationPanel() {
                this.currentResourceItem.showDeleteConfirmation = false;
            },
        },
    }
</script>

<style lang="scss">

</style>
