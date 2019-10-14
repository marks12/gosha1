
	<template>
    <WorkSpace>
        <template #header>
            <slot name="pageHeader">
                <VHead level="h1">FieldTypeFilter</VHead>
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
                            v-for="fieldTypeFilterItem in fieldTypeFilterList"
                            :key="fieldTypeFilterItem.Id"
                            @click="selectFieldTypeFilterItem(fieldTypeFilterItem)"
                            class="sw-table__row_can-select"
                            :class="{'sw-table__row_is-selected': fieldTypeFilterItem.Id === currentFieldTypeFilterItem.item.Id}"
                        >
                            <td v-for="(value, key) in fields" :key="key + '-fields'">
                                <VCheckbox v-if="isCheckbox(fieldTypeFilterItem[key])" :checked="fieldTypeFilterItem[key]" disabled></VCheckbox>
                                <VText v-else>{{ fieldTypeFilterItem[key] }}</VText>
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
                                        :for="`currentFieldTypeFilterItem${key}`"
                                    >{{ filed }}</VLabel>
                                    <VInput
										v-if="isInput(currentFieldTypeFilterItem.item[key])"
                                        v-model="currentFieldTypeFilterItem.item[key]"
                                        width="dyn"
                                        :id="`currentFieldTypeFilterItem${key}`"
                                        @input="changeCurrentFieldTypeFilterItem"
                                    />
									<VCheckbox
										v-if="isCheckbox(currentFieldTypeFilterItem.item[key])"
                                        v-model="currentFieldTypeFilterItem.item[key]"
                                        :id="`currentFieldTypeFilterItem${key}`"
										@input="changeCurrentFieldTypeFilterItem"
									/>
									
                                </VSet>
                            </VSet>
                            <button type="submit" :disabled="!currentFieldTypeFilterItem.hasChange" hidden></button>
                        </form>
                    </template>

                    <template #footer>
                        <VSet>
                            <VButton
                                @click="saveChangesSubmit"
                                accent
                                :text="panelSubmitButtonText"
                                :disabled="!currentFieldTypeFilterItem.hasChange"
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
                    v-if="currentFieldTypeFilterItem.showDeleteConfirmation"
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
                        :disabled="!currentFieldTypeFilterItem.isSelected"
                        @click="deleteFieldTypeFilterItemHandler"
                    />
                </VSet>
            </slot>
        </template>
    </WorkSpace>
</template>

<script>
    import fieldTypeFilterData from "../data/FieldTypeFilterData";
    import { FieldTypeFilter } from '../apiModel';
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
        name: 'FieldTypeFilterGen',

        components: {VSelectSimple, VSign, VIcon, VButton, VPanel, VText, VInput, VLabel, VSet, VHead, WorkSpace, VCheckbox},

        props: {
            fields: {
                type: Object,
                default() {
                    const fieldTypeFilterItem = new FieldTypeFilter();
                    const fieldsObj = {};

                    for (let prop in fieldTypeFilterItem) {

                        if (fieldTypeFilterItem.hasOwnProperty(prop)) {
                            fieldsObj[prop] = prop;
                        }

                    }

                    return fieldsObj;
                }
            },
            editFields: {
                type: Object,
                default() {
                    const fieldTypeFilterItem = new FieldTypeFilter();
                    const fieldsObj = {};

                    for (let prop in fieldTypeFilterItem) {

                        if (fieldTypeFilterItem.hasOwnProperty(prop)) {
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
            return fieldTypeFilterData;
        },

        created() {
            this.fillFieldTypeFilterFilter();
            this.fetchFieldTypeFilterData();
        },

        computed: {
            ...mapGetters({
                fieldTypeFilterList: 'getListFieldTypeFilter'
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
                'findFieldTypeFilter',
                'updateFieldTypeFilter',
                'deleteFieldTypeFilter',
                'createFieldTypeFilter',
            ]),

            ...mapMutations([
                'addFieldTypeFilterItemToList',
                'deleteFieldTypeFilterFromList',
                'updateFieldTypeFilterById',
            ]),

            fillFieldTypeFilterFilter() {
                this.fieldTypeFilterFilter.CurrentPage = 1;
                this.fieldTypeFilterFilter.PerPage = 1000;
            },

            fetchFieldTypeFilterData() {
                return this.findFieldTypeFilter({
                    filter: this.fieldTypeFilterFilter
                });
            },

            /**
             *
             * @param type
             */
            showPanel(type) {
                if (type === this.panel.create) {
                    this.panel.type = this.panel.create;
                    this.clearPanelFieldTypeFilterItem();
                } else if (type === this.panel.edit) {
                    this.panel.type = this.panel.edit;
                    this.currentFieldTypeFilterItem.isSelected = true;
                }

                this.panel.show = true;
            },

            closePanel() {
                this.panel.show = false;
                this.currentFieldTypeFilterItem.isSelected = false;
                this.clearPanelFieldTypeFilterItem();
            },

            selectFieldTypeFilterItem(fieldTypeFilterItem) {
                this.showPanel(this.panel.edit);
                this.currentFieldTypeFilterItem.isSelected = true;
                Object.assign(this.currentFieldTypeFilterItem.item, fieldTypeFilterItem);
            },

            changeCurrentFieldTypeFilterItem() {
                this.currentFieldTypeFilterItem.hasChange = true;
            },

            cancelChanges() {
                this.clearPanelFieldTypeFilterItem();
                this.closePanel();
            },

            clearPanelFieldTypeFilterItem() {
                this.currentFieldTypeFilterItem.item = new FieldTypeFilter();
                this.currentFieldTypeFilterItem.hasChange = false;
            },

            saveChangesSubmit() {
                if (this.isPanelCreate) {
                    this.createFieldTypeFilterItemSubmit();
                    return;
                }

                if (this.isPanelEdit) {
                    this.editFieldTypeFilterItemSubmit();
                }
            },

            createFieldTypeFilterItemSubmit() {
                this.createFieldTypeFilter({
					data: this.currentFieldTypeFilterItem.item,
                }).then((response) => {

                    if (response.Model) {
                        this.addFieldTypeFilterItemToList(response.Model);
                        this.clearPanelFieldTypeFilterItem();
                    } else {
                        console.error('Ошибка создания записи: ', response.Error);
                    }

                }).catch(error => {
                    console.error('Ошибка создания записи: ', error);
                });
            },

            editFieldTypeFilterItemSubmit() {
                if (this.currentFieldTypeFilterItem.hasChange) {
                    this.updateFieldTypeFilter({
                        id: this.currentFieldTypeFilterItem.item.Id,
                        data: this.currentFieldTypeFilterItem.item,
                    }).then((response) => {

                        if (response.Model) {
                            this.updateFieldTypeFilterById(response.Model);
                            this.currentFieldTypeFilterItem.hasChange = false;
                            this.clearPanelFieldTypeFilterItem();
                            this.closePanel();
                        } else {
                            console.error('Ошибка изменения записи: ', response.Error);
                        }

                    }).catch(error => {
                        console.error('Ошибка изменения записи: ', error);
                    });
                }
            },

            deleteFieldTypeFilterItemHandler() {
                let deletedItemId = this.currentFieldTypeFilterItem.item.Id;

                if (!this.currentFieldTypeFilterItem.canDelete) {
                    this.currentFieldTypeFilterItem.showDeleteConfirmation = true;
                    return;
                }

                this.deleteFieldTypeFilter({
                    id: deletedItemId
                }).then(response => {

                    if (response.IsSuccess) {
                        this.deleteFieldTypeFilterFromList(deletedItemId);
                        this.clearPanelFieldTypeFilterItem();
                        this.currentFieldTypeFilterItem.canDelete = false;
                        this.currentFieldTypeFilterItem.isSelected = false;
                        this.panel.show = false;
                    } else {
                        console.error('Ошибка удаления элемента: ', response.Error);
                    }

                }).catch(error => {
                    console.error('Ошибка удаления элемента: ', error);
                });
            },

            confirmDeleteHandler() {
                this.currentFieldTypeFilterItem.showDeleteConfirmation = false;
                this.currentFieldTypeFilterItem.canDelete = true;
                this.deleteFieldTypeFilterItemHandler();
            },

            closeConfirmationPanel() {
                this.currentFieldTypeFilterItem.showDeleteConfirmation = false;
            },
        },
    }
</script>

<style lang="scss">

</style>
