
	<template>
    <WorkSpace>
        <template #header>
            <slot name="pageHeader">
                <VHead level="h1">APIError</VHead>
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
                            v-for="aPIErrorItem in aPIErrorList"
                            :key="aPIErrorItem.Id"
                            @click="selectAPIErrorItem(aPIErrorItem)"
                            class="sw-table__row_can-select"
                            :class="{'sw-table__row_is-selected': aPIErrorItem.Id === currentAPIErrorItem.item.Id}"
                        >
                            <td v-for="(value, key) in fields" :key="key + '-fields'">
                                <VCheckbox v-if="isCheckbox(aPIErrorItem[key])" :checked="aPIErrorItem[key]" disabled></VCheckbox>
                                <VText v-else>{{ aPIErrorItem[key] }}</VText>
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
                                        :for="`currentAPIErrorItem${key}`"
                                    >{{ filed }}</VLabel>
                                    <VInput
										v-if="isInput(currentAPIErrorItem.item[key])"
                                        v-model="currentAPIErrorItem.item[key]"
                                        width="dyn"
                                        :id="`currentAPIErrorItem${key}`"
                                        @input="changeCurrentAPIErrorItem"
                                    />
									<VCheckbox
										v-if="isCheckbox(currentAPIErrorItem.item[key])"
                                        v-model="currentAPIErrorItem.item[key]"
                                        :id="`currentAPIErrorItem${key}`"
										@input="changeCurrentAPIErrorItem"
									/>
									
                                </VSet>
                            </VSet>
                            <button type="submit" :disabled="!currentAPIErrorItem.hasChange" hidden></button>
                        </form>
                    </template>

                    <template #footer>
                        <VSet>
                            <VButton
                                @click="saveChangesSubmit"
                                accent
                                :text="panelSubmitButtonText"
                                :disabled="!currentAPIErrorItem.hasChange"
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
                    v-if="currentAPIErrorItem.showDeleteConfirmation"
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
                        :disabled="!currentAPIErrorItem.isSelected"
                        @click="deleteAPIErrorItemHandler"
                    />
                </VSet>
            </slot>
        </template>
    </WorkSpace>
</template>

<script>
    import aPIErrorData from "../data/APIErrorData";
    import { APIError } from '../apiModel';
    import { mapGetters, mapMutations, mapActions } from 'vuex';
    import WorkSpace from "swtui/src/components/WorkSpace";
    import VHead from "swtui/src/components/VHead";
    import VSet from "swtui/src/components/VSet";
    import VLabel from "swtui/src/components/VLabel";
    import VInput from "swtui/src/components/VInput";
    import VCheckbox from "swtui/src/components/VCheckbox";
    import VText from "swtui/src/components/VText";
    import VPanel from "swtui/src/components/VPanel";
    import VButton from "swtui/src/components/VButton";

    export default {
        name: 'APIErrorGen',

        components: {VButton, VPanel, VText, VInput, VLabel, VSet, VHead, WorkSpace, VCheckbox},

        props: {
            fields: {
                type: Object,
                default() {
                    const aPIErrorItem = new APIError();
                    const fieldsObj = {};

                    for (let prop in aPIErrorItem) {

                        if (aPIErrorItem.hasOwnProperty(prop)) {
                            fieldsObj[prop] = prop;
                        }

                    }

                    return fieldsObj;
                }
            },
            editFields: {
                type: Object,
                default() {
                    const aPIErrorItem = new APIError();
                    const fieldsObj = {};

                    for (let prop in aPIErrorItem) {

                        if (aPIErrorItem.hasOwnProperty(prop)) {
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
            return aPIErrorData;
        },

        created() {
			this.onCreated();
        },

        computed: {
            ...mapGetters('gosha', {
                aPIErrorList: 'getListAPIError'
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
            ...mapActions('gosha', [
                'findAPIError',
                'updateAPIError',
                'deleteAPIError',
                'createAPIError',
            ]),

            ...mapMutations('gosha', [
                'addAPIErrorItemToList',
                'deleteAPIErrorFromList',
                'updateAPIErrorById',
            ]),

			onCreated() {
				this.fillAPIErrorFilter();
	            this.fetchAPIErrorData();
			},

            fillAPIErrorFilter() {
                this.aPIErrorFilter.CurrentPage = 1;
                this.aPIErrorFilter.PerPage = 1000;
            },

            fetchAPIErrorData() {
                return this.findAPIError({
                    filter: this.aPIErrorFilter
                });
            },

            /**
             *
             * @param type
             */
            showPanel(type) {
                if (type === this.panel.create) {
                    this.panel.type = this.panel.create;
                    this.clearPanelAPIErrorItem();
                } else if (type === this.panel.edit) {
                    this.panel.type = this.panel.edit;
                    this.currentAPIErrorItem.isSelected = true;
                }

                this.panel.show = true;
            },

            closePanel() {
                this.panel.show = false;
                this.currentAPIErrorItem.isSelected = false;
                this.clearPanelAPIErrorItem();
            },

            selectAPIErrorItem(aPIErrorItem) {
                this.showPanel(this.panel.edit);
                this.currentAPIErrorItem.isSelected = true;
                Object.assign(this.currentAPIErrorItem.item, aPIErrorItem);
            },

            changeCurrentAPIErrorItem() {
                this.currentAPIErrorItem.hasChange = true;
            },

            cancelChanges() {
                this.clearPanelAPIErrorItem();
                this.closePanel();
            },

            clearPanelAPIErrorItem() {
                this.currentAPIErrorItem.item = new APIError();
                this.currentAPIErrorItem.hasChange = false;
            },

            saveChangesSubmit() {
                if (this.isPanelCreate) {
                    this.createAPIErrorItemSubmit();
                    return;
                }

                if (this.isPanelEdit) {
                    this.editAPIErrorItemSubmit();
                }
            },

            createAPIErrorItemSubmit() {
                return this.createAPIError({
					data: this.currentAPIErrorItem.item,
                }).then((response) => {

                    if (response.Model) {
                        this.addAPIErrorItemToList(response.Model);
                        this.clearPanelAPIErrorItem();
                    } else {
                        console.error('Ошибка создания записи: ', response.Error);
                    }

                }).catch(error => {
                    console.error('Ошибка создания записи: ', error);
                });
            },

            editAPIErrorItemSubmit() {

                if (this.currentAPIErrorItem.hasChange) {
                    return this.updateAPIError({
                        id: this.currentAPIErrorItem.item.Id,
                        data: this.currentAPIErrorItem.item,
                    }).then((response) => {

                        if (response.Model) {
                            this.updateAPIErrorById(response.Model);
                            this.currentAPIErrorItem.hasChange = false;
                            this.clearPanelAPIErrorItem();
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

            deleteAPIErrorItemHandler() {
                let deletedItemId = this.currentAPIErrorItem.item.Id;

                if (!this.currentAPIErrorItem.canDelete) {
                    this.currentAPIErrorItem.showDeleteConfirmation = true;
                    return;
                }

                this.deleteAPIError({
                    id: deletedItemId
                }).then(response => {

                    if (response.IsSuccess) {
                        this.deleteAPIErrorFromList(deletedItemId);
                        this.clearPanelAPIErrorItem();
                        this.currentAPIErrorItem.canDelete = false;
                        this.currentAPIErrorItem.isSelected = false;
                        this.panel.show = false;
                    } else {
                        console.error('Ошибка удаления элемента: ', response.Error);
                    }

                }).catch(error => {
                    console.error('Ошибка удаления элемента: ', error);
                });
            },

            confirmDeleteHandler() {
                this.currentAPIErrorItem.showDeleteConfirmation = false;
                this.currentAPIErrorItem.canDelete = true;
                this.deleteAPIErrorItemHandler();
            },

            closeConfirmationPanel() {
                this.currentAPIErrorItem.showDeleteConfirmation = false;
            },
        },
    }
</script>

<style lang="scss">

</style>
