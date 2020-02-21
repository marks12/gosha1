
	<template>
    <WorkSpace>
        <template #header>
            <slot name="pageHeader">
                <VHead level="h1">FieldType</VHead>
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
                            v-for="fieldTypeItem in fieldTypeList"
                            :key="fieldTypeItem.Id"
                            @click="selectFieldTypeItem(fieldTypeItem)"
                            class="sw-table__row_can-select"
                            :class="{'sw-table__row_is-selected': fieldTypeItem.Id === currentFieldTypeItem.item.Id}"
                        >
                            <td v-for="(value, key) in fields" :key="key + '-fields'">
                                <VCheckbox v-if="isCheckbox(fieldTypeItem[key])" :checked="fieldTypeItem[key]" disabled></VCheckbox>
                                <VText v-else>{{ fieldTypeItem[key] }}</VText>
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
                                        :for="`currentFieldTypeItem${key}`"
                                    >{{ filed }}</VLabel>
                                    <VInput
										v-if="isInput(currentFieldTypeItem.item[key])"
                                        v-model="currentFieldTypeItem.item[key]"
                                        width="dyn"
                                        :id="`currentFieldTypeItem${key}`"
                                        @input="changeCurrentFieldTypeItem"
                                    />
									<VCheckbox
										v-if="isCheckbox(currentFieldTypeItem.item[key])"
                                        v-model="currentFieldTypeItem.item[key]"
                                        :id="`currentFieldTypeItem${key}`"
										@input="changeCurrentFieldTypeItem"
									/>
									
                                </VSet>
                            </VSet>
                            <button type="submit" :disabled="!currentFieldTypeItem.hasChange" hidden></button>
                        </form>
                    </template>

                    <template #footer>
                        <VSet>
                            <VButton
                                @click="saveChangesSubmit"
                                accent
                                :text="panelSubmitButtonText"
                                :disabled="!currentFieldTypeItem.hasChange"
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
                    v-if="currentFieldTypeItem.showDeleteConfirmation"
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
                        :disabled="!currentFieldTypeItem.isSelected"
                        @click="deleteFieldTypeItemHandler"
                    />
                </VSet>
            </slot>
        </template>
    </WorkSpace>
</template>

<script>
    import fieldTypeData from "../data/FieldTypeData";
    import { FieldType } from '../apiModel';
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
        name: 'FieldTypeGen',

        components: {VButton, VPanel, VText, VInput, VLabel, VSet, VHead, WorkSpace, VCheckbox},

        props: {
            fields: {
                type: Object,
                default() {
                    const fieldTypeItem = new FieldType();
                    const fieldsObj = {};

                    for (let prop in fieldTypeItem) {

                        if (fieldTypeItem.hasOwnProperty(prop)) {
                            fieldsObj[prop] = prop;
                        }

                    }

                    return fieldsObj;
                }
            },
            editFields: {
                type: Object,
                default() {
                    const fieldTypeItem = new FieldType();
                    const fieldsObj = {};

                    for (let prop in fieldTypeItem) {

                        if (fieldTypeItem.hasOwnProperty(prop)) {
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
            return fieldTypeData;
        },

        created() {
			this.onCreated();
        },

        computed: {
            ...mapGetters({
                fieldTypeList: 'getListFieldType'
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
                'findFieldType',
                'updateFieldType',
                'deleteFieldType',
                'createFieldType',
            ]),

            ...mapMutations([
                'addFieldTypeItemToList',
                'deleteFieldTypeFromList',
                'updateFieldTypeById',
            ]),

			onCreated() {
				this.fillFieldTypeFilter();
	            this.fetchFieldTypeData();
			},

            fillFieldTypeFilter() {
                this.fieldTypeFilter.CurrentPage = 1;
                this.fieldTypeFilter.PerPage = 1000;
            },

            fetchFieldTypeData() {
                return this.findFieldType({
                    filter: this.fieldTypeFilter
                });
            },

            /**
             *
             * @param type
             */
            showPanel(type) {
                if (type === this.panel.create) {
                    this.panel.type = this.panel.create;
                    this.clearPanelFieldTypeItem();
                } else if (type === this.panel.edit) {
                    this.panel.type = this.panel.edit;
                    this.currentFieldTypeItem.isSelected = true;
                }

                this.panel.show = true;
            },

            closePanel() {
                this.panel.show = false;
                this.currentFieldTypeItem.isSelected = false;
                this.clearPanelFieldTypeItem();
            },

            selectFieldTypeItem(fieldTypeItem) {
                this.showPanel(this.panel.edit);
                this.currentFieldTypeItem.isSelected = true;
                Object.assign(this.currentFieldTypeItem.item, fieldTypeItem);
            },

            changeCurrentFieldTypeItem() {
                this.currentFieldTypeItem.hasChange = true;
            },

            cancelChanges() {
                this.clearPanelFieldTypeItem();
                this.closePanel();
            },

            clearPanelFieldTypeItem() {
                this.currentFieldTypeItem.item = new FieldType();
                this.currentFieldTypeItem.hasChange = false;
            },

            saveChangesSubmit() {
                if (this.isPanelCreate) {
                    this.createFieldTypeItemSubmit();
                    return;
                }

                if (this.isPanelEdit) {
                    this.editFieldTypeItemSubmit();
                }
            },

            createFieldTypeItemSubmit() {
                return this.createFieldType({
					data: this.currentFieldTypeItem.item,
                }).then((response) => {

                    if (response.Model) {
                        this.addFieldTypeItemToList(response.Model);
                        this.clearPanelFieldTypeItem();
                    } else {
                        console.error('Ошибка создания записи: ', response.Error);
                    }

                }).catch(error => {
                    console.error('Ошибка создания записи: ', error);
                });
            },

            editFieldTypeItemSubmit() {

                if (this.currentFieldTypeItem.hasChange) {
                    return this.updateFieldType({
                        id: this.currentFieldTypeItem.item.Id,
                        data: this.currentFieldTypeItem.item,
                    }).then((response) => {

                        if (response.Model) {
                            this.updateFieldTypeById(response.Model);
                            this.currentFieldTypeItem.hasChange = false;
                            this.clearPanelFieldTypeItem();
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

            deleteFieldTypeItemHandler() {
                let deletedItemId = this.currentFieldTypeItem.item.Id;

                if (!this.currentFieldTypeItem.canDelete) {
                    this.currentFieldTypeItem.showDeleteConfirmation = true;
                    return;
                }

                this.deleteFieldType({
                    id: deletedItemId
                }).then(response => {

                    if (response.IsSuccess) {
                        this.deleteFieldTypeFromList(deletedItemId);
                        this.clearPanelFieldTypeItem();
                        this.currentFieldTypeItem.canDelete = false;
                        this.currentFieldTypeItem.isSelected = false;
                        this.panel.show = false;
                    } else {
                        console.error('Ошибка удаления элемента: ', response.Error);
                    }

                }).catch(error => {
                    console.error('Ошибка удаления элемента: ', error);
                });
            },

            confirmDeleteHandler() {
                this.currentFieldTypeItem.showDeleteConfirmation = false;
                this.currentFieldTypeItem.canDelete = true;
                this.deleteFieldTypeItemHandler();
            },

            closeConfirmationPanel() {
                this.currentFieldTypeItem.showDeleteConfirmation = false;
            },
        },
    }
</script>

<style lang="scss">

</style>
