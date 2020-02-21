
	<template>
    <WorkSpace>
        <template #header>
            <slot name="pageHeader">
                <VHead level="h1">Field</VHead>
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
                            v-for="fieldItem in fieldList"
                            :key="fieldItem.Id"
                            @click="selectFieldItem(fieldItem)"
                            class="sw-table__row_can-select"
                            :class="{'sw-table__row_is-selected': fieldItem.Id === currentFieldItem.item.Id}"
                        >
                            <td v-for="(value, key) in fields" :key="key + '-fields'">
                                <VCheckbox v-if="isCheckbox(fieldItem[key])" :checked="fieldItem[key]" disabled></VCheckbox>
                                <VText v-else>{{ fieldItem[key] }}</VText>
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
                                        :for="`currentFieldItem${key}`"
                                    >{{ filed }}</VLabel>
                                    <VInput
										v-if="isInput(currentFieldItem.item[key])"
                                        v-model="currentFieldItem.item[key]"
                                        width="dyn"
                                        :id="`currentFieldItem${key}`"
                                        @input="changeCurrentFieldItem"
                                    />
									<VCheckbox
										v-if="isCheckbox(currentFieldItem.item[key])"
                                        v-model="currentFieldItem.item[key]"
                                        :id="`currentFieldItem${key}`"
										@input="changeCurrentFieldItem"
									/>
									
                                </VSet>
                            </VSet>
                            <button type="submit" :disabled="!currentFieldItem.hasChange" hidden></button>
                        </form>
                    </template>

                    <template #footer>
                        <VSet>
                            <VButton
                                @click="saveChangesSubmit"
                                accent
                                :text="panelSubmitButtonText"
                                :disabled="!currentFieldItem.hasChange"
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
                    v-if="currentFieldItem.showDeleteConfirmation"
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
                        :disabled="!currentFieldItem.isSelected"
                        @click="deleteFieldItemHandler"
                    />
                </VSet>
            </slot>
        </template>
    </WorkSpace>
</template>

<script>
    import fieldData from "../data/FieldData";
    import { Field } from '../apiModel';
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
        name: 'FieldGen',

        components: {VButton, VPanel, VText, VInput, VLabel, VSet, VHead, WorkSpace, VCheckbox},

        props: {
            fields: {
                type: Object,
                default() {
                    const fieldItem = new Field();
                    const fieldsObj = {};

                    for (let prop in fieldItem) {

                        if (fieldItem.hasOwnProperty(prop)) {
                            fieldsObj[prop] = prop;
                        }

                    }

                    return fieldsObj;
                }
            },
            editFields: {
                type: Object,
                default() {
                    const fieldItem = new Field();
                    const fieldsObj = {};

                    for (let prop in fieldItem) {

                        if (fieldItem.hasOwnProperty(prop)) {
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
            return fieldData;
        },

        created() {
			this.onCreated();
        },

        computed: {
            ...mapGetters({
                fieldList: 'getListField'
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
                'findField',
                'updateField',
                'deleteField',
                'createField',
            ]),

            ...mapMutations([
                'addFieldItemToList',
                'deleteFieldFromList',
                'updateFieldById',
            ]),

			onCreated() {
				this.fillFieldFilter();
	            this.fetchFieldData();
			},

            fillFieldFilter() {
                this.fieldFilter.CurrentPage = 1;
                this.fieldFilter.PerPage = 1000;
            },

            fetchFieldData() {
                return this.findField({
                    filter: this.fieldFilter
                });
            },

            /**
             *
             * @param type
             */
            showPanel(type) {
                if (type === this.panel.create) {
                    this.panel.type = this.panel.create;
                    this.clearPanelFieldItem();
                } else if (type === this.panel.edit) {
                    this.panel.type = this.panel.edit;
                    this.currentFieldItem.isSelected = true;
                }

                this.panel.show = true;
            },

            closePanel() {
                this.panel.show = false;
                this.currentFieldItem.isSelected = false;
                this.clearPanelFieldItem();
            },

            selectFieldItem(fieldItem) {
                this.showPanel(this.panel.edit);
                this.currentFieldItem.isSelected = true;
                Object.assign(this.currentFieldItem.item, fieldItem);
            },

            changeCurrentFieldItem() {
                this.currentFieldItem.hasChange = true;
            },

            cancelChanges() {
                this.clearPanelFieldItem();
                this.closePanel();
            },

            clearPanelFieldItem() {
                this.currentFieldItem.item = new Field();
                this.currentFieldItem.hasChange = false;
            },

            saveChangesSubmit() {
                if (this.isPanelCreate) {
                    this.createFieldItemSubmit();
                    return;
                }

                if (this.isPanelEdit) {
                    this.editFieldItemSubmit();
                }
            },

            createFieldItemSubmit() {
                return this.createField({
					data: this.currentFieldItem.item,
                }).then((response) => {

                    if (response.Model) {
                        this.addFieldItemToList(response.Model);
                        this.clearPanelFieldItem();
                    } else {
                        console.error('Ошибка создания записи: ', response.Error);
                    }

                }).catch(error => {
                    console.error('Ошибка создания записи: ', error);
                });
            },

            editFieldItemSubmit() {

                if (this.currentFieldItem.hasChange) {
                    return this.updateField({
                        id: this.currentFieldItem.item.Id,
                        data: this.currentFieldItem.item,
                    }).then((response) => {

                        if (response.Model) {
                            this.updateFieldById(response.Model);
                            this.currentFieldItem.hasChange = false;
                            this.clearPanelFieldItem();
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

            deleteFieldItemHandler() {
                let deletedItemId = this.currentFieldItem.item.Id;

                if (!this.currentFieldItem.canDelete) {
                    this.currentFieldItem.showDeleteConfirmation = true;
                    return;
                }

                this.deleteField({
                    id: deletedItemId
                }).then(response => {

                    if (response.IsSuccess) {
                        this.deleteFieldFromList(deletedItemId);
                        this.clearPanelFieldItem();
                        this.currentFieldItem.canDelete = false;
                        this.currentFieldItem.isSelected = false;
                        this.panel.show = false;
                    } else {
                        console.error('Ошибка удаления элемента: ', response.Error);
                    }

                }).catch(error => {
                    console.error('Ошибка удаления элемента: ', error);
                });
            },

            confirmDeleteHandler() {
                this.currentFieldItem.showDeleteConfirmation = false;
                this.currentFieldItem.canDelete = true;
                this.deleteFieldItemHandler();
            },

            closeConfirmationPanel() {
                this.currentFieldItem.showDeleteConfirmation = false;
            },
        },
    }
</script>

<style lang="scss">

</style>
