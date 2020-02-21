
	<template>
    <WorkSpace>
        <template #header>
            <slot name="pageHeader">
                <VHead level="h1">APIStatus</VHead>
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
                            v-for="aPIStatusItem in aPIStatusList"
                            :key="aPIStatusItem.Id"
                            @click="selectAPIStatusItem(aPIStatusItem)"
                            class="sw-table__row_can-select"
                            :class="{'sw-table__row_is-selected': aPIStatusItem.Id === currentAPIStatusItem.item.Id}"
                        >
                            <td v-for="(value, key) in fields" :key="key + '-fields'">
                                <VCheckbox v-if="isCheckbox(aPIStatusItem[key])" :checked="aPIStatusItem[key]" disabled></VCheckbox>
                                <VText v-else>{{ aPIStatusItem[key] }}</VText>
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
                            <VSet vertical>
                                <VSet
                                    v-for="(filed, key) in editFields" :key="key + '-editFields'"
                                    vertical-align="center"
                                >
                                    <VLabel
                                        width="col4"
                                        :for="`currentAPIStatusItem${key}`"
                                    >{{ filed }}</VLabel>
                                    <VInput
										v-if="isInput(currentAPIStatusItem.item[key])"
                                        v-model="currentAPIStatusItem.item[key]"
                                        width="dyn"
                                        :id="`currentAPIStatusItem${key}`"
                                        @input="changeCurrentAPIStatusItem"
                                    />
									<VCheckbox
										v-if="isCheckbox(currentAPIStatusItem.item[key])"
                                        v-model="currentAPIStatusItem.item[key]"
                                        :id="`currentAPIStatusItem${key}`"
										@input="changeCurrentAPIStatusItem"
									/>
									
                                </VSet>
                            </VSet>
                            <button type="submit" :disabled="!currentAPIStatusItem.hasChange" hidden></button>
                        </form>
                    </template>

                    <template #footer>
                        <VSet>
                            <VButton
                                @click="saveChangesSubmit"
                                accent
                                :text="panelSubmitButtonText"
                                :disabled="!currentAPIStatusItem.hasChange"
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
                    v-if="currentAPIStatusItem.showDeleteConfirmation"
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
                        :disabled="!currentAPIStatusItem.isSelected"
                        @click="deleteAPIStatusItemHandler"
                    />
                </VSet>
            </slot>
        </template>
    </WorkSpace>
</template>

<script>
    import aPIStatusData from "../data/APIStatusData";
    import { APIStatus } from '../apiModel';
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
        name: 'APIStatusGen',

        components: {VButton, VPanel, VText, VInput, VLabel, VSet, VHead, WorkSpace, VCheckbox},

        props: {
            fields: {
                type: Object,
                default() {
                    const aPIStatusItem = new APIStatus();
                    const fieldsObj = {};

                    for (let prop in aPIStatusItem) {

                        if (aPIStatusItem.hasOwnProperty(prop)) {
                            fieldsObj[prop] = prop;
                        }

                    }

                    return fieldsObj;
                }
            },
            editFields: {
                type: Object,
                default() {
                    const aPIStatusItem = new APIStatus();
                    const fieldsObj = {};

                    for (let prop in aPIStatusItem) {

                        if (aPIStatusItem.hasOwnProperty(prop)) {
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
            return aPIStatusData;
        },

        created() {
			this.onCreated();
        },

        computed: {
            ...mapGetters({
                aPIStatusList: 'getListAPIStatus'
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
                'findAPIStatus',
                'updateAPIStatus',
                'deleteAPIStatus',
                'createAPIStatus',
            ]),

            ...mapMutations([
                'addAPIStatusItemToList',
                'deleteAPIStatusFromList',
                'updateAPIStatusById',
            ]),

			onCreated() {
				this.fillAPIStatusFilter();
	            this.fetchAPIStatusData();
			},

            fillAPIStatusFilter() {
                this.aPIStatusFilter.CurrentPage = 1;
                this.aPIStatusFilter.PerPage = 1000;
            },

            fetchAPIStatusData() {
                return this.findAPIStatus({
                    filter: this.aPIStatusFilter
                });
            },

            /**
             *
             * @param type
             */
            showPanel(type) {
                if (type === this.panel.create) {
                    this.panel.type = this.panel.create;
                    this.clearPanelAPIStatusItem();
                } else if (type === this.panel.edit) {
                    this.panel.type = this.panel.edit;
                    this.currentAPIStatusItem.isSelected = true;
                }

                this.panel.show = true;
            },

            closePanel() {
                this.panel.show = false;
                this.currentAPIStatusItem.isSelected = false;
                this.clearPanelAPIStatusItem();
            },

            selectAPIStatusItem(aPIStatusItem) {
                this.showPanel(this.panel.edit);
                this.currentAPIStatusItem.isSelected = true;
                Object.assign(this.currentAPIStatusItem.item, aPIStatusItem);
            },

            changeCurrentAPIStatusItem() {
                this.currentAPIStatusItem.hasChange = true;
            },

            cancelChanges() {
                this.clearPanelAPIStatusItem();
                this.closePanel();
            },

            clearPanelAPIStatusItem() {
                this.currentAPIStatusItem.item = new APIStatus();
                this.currentAPIStatusItem.hasChange = false;
            },

            saveChangesSubmit() {
                if (this.isPanelCreate) {
                    this.createAPIStatusItemSubmit();
                    return;
                }

                if (this.isPanelEdit) {
                    this.editAPIStatusItemSubmit();
                }
            },

            createAPIStatusItemSubmit() {
                return this.createAPIStatus({
					data: this.currentAPIStatusItem.item,
                }).then((response) => {

                    if (response.Model) {
                        this.addAPIStatusItemToList(response.Model);
                        this.clearPanelAPIStatusItem();
                    } else {
                        console.error('Ошибка создания записи: ', response.Error);
                    }

                }).catch(error => {
                    console.error('Ошибка создания записи: ', error);
                });
            },

            editAPIStatusItemSubmit() {

                if (this.currentAPIStatusItem.hasChange) {
                    return this.updateAPIStatus({
                        id: this.currentAPIStatusItem.item.Id,
                        data: this.currentAPIStatusItem.item,
                    }).then((response) => {

                        if (response.Model) {
                            this.updateAPIStatusById(response.Model);
                            this.currentAPIStatusItem.hasChange = false;
                            this.clearPanelAPIStatusItem();
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

            deleteAPIStatusItemHandler() {
                let deletedItemId = this.currentAPIStatusItem.item.Id;

                if (!this.currentAPIStatusItem.canDelete) {
                    this.currentAPIStatusItem.showDeleteConfirmation = true;
                    return;
                }

                this.deleteAPIStatus({
                    id: deletedItemId
                }).then(response => {

                    if (response.IsSuccess) {
                        this.deleteAPIStatusFromList(deletedItemId);
                        this.clearPanelAPIStatusItem();
                        this.currentAPIStatusItem.canDelete = false;
                        this.currentAPIStatusItem.isSelected = false;
                        this.panel.show = false;
                    } else {
                        console.error('Ошибка удаления элемента: ', response.Error);
                    }

                }).catch(error => {
                    console.error('Ошибка удаления элемента: ', error);
                });
            },

            confirmDeleteHandler() {
                this.currentAPIStatusItem.showDeleteConfirmation = false;
                this.currentAPIStatusItem.canDelete = true;
                this.deleteAPIStatusItemHandler();
            },

            closeConfirmationPanel() {
                this.currentAPIStatusItem.showDeleteConfirmation = false;
            },
        },
    }
</script>

<style lang="scss">

</style>
