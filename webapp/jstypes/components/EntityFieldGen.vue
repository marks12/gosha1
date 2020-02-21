
	<template>
    <WorkSpace>
        <template #header>
            <slot name="pageHeader">
                <VHead level="h1">EntityField</VHead>
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
                            v-for="entityFieldItem in entityFieldList"
                            :key="entityFieldItem.Id"
                            @click="selectEntityFieldItem(entityFieldItem)"
                            class="sw-table__row_can-select"
                            :class="{'sw-table__row_is-selected': entityFieldItem.Id === currentEntityFieldItem.item.Id}"
                        >
                            <td v-for="(value, key) in fields" :key="key + '-fields'">
                                <VCheckbox v-if="isCheckbox(entityFieldItem[key])" :checked="entityFieldItem[key]" disabled></VCheckbox>
                                <VText v-else>{{ entityFieldItem[key] }}</VText>
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
                                        :for="`currentEntityFieldItem${key}`"
                                    >{{ filed }}</VLabel>
                                    <VInput
										v-if="isInput(currentEntityFieldItem.item[key])"
                                        v-model="currentEntityFieldItem.item[key]"
                                        width="dyn"
                                        :id="`currentEntityFieldItem${key}`"
                                        @input="changeCurrentEntityFieldItem"
                                    />
									<VCheckbox
										v-if="isCheckbox(currentEntityFieldItem.item[key])"
                                        v-model="currentEntityFieldItem.item[key]"
                                        :id="`currentEntityFieldItem${key}`"
										@input="changeCurrentEntityFieldItem"
									/>
									
                                </VSet>
                            </VSet>
                            <button type="submit" :disabled="!currentEntityFieldItem.hasChange" hidden></button>
                        </form>
                    </template>

                    <template #footer>
                        <VSet>
                            <VButton
                                @click="saveChangesSubmit"
                                accent
                                :text="panelSubmitButtonText"
                                :disabled="!currentEntityFieldItem.hasChange"
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
                    v-if="currentEntityFieldItem.showDeleteConfirmation"
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
                        :disabled="!currentEntityFieldItem.isSelected"
                        @click="deleteEntityFieldItemHandler"
                    />
                </VSet>
            </slot>
        </template>
    </WorkSpace>
</template>

<script>
    import entityFieldData from "../data/EntityFieldData";
    import { EntityField } from '../apiModel';
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
        name: 'EntityFieldGen',

        components: {VButton, VPanel, VText, VInput, VLabel, VSet, VHead, WorkSpace, VCheckbox},

        props: {
            fields: {
                type: Object,
                default() {
                    const entityFieldItem = new EntityField();
                    const fieldsObj = {};

                    for (let prop in entityFieldItem) {

                        if (entityFieldItem.hasOwnProperty(prop)) {
                            fieldsObj[prop] = prop;
                        }

                    }

                    return fieldsObj;
                }
            },
            editFields: {
                type: Object,
                default() {
                    const entityFieldItem = new EntityField();
                    const fieldsObj = {};

                    for (let prop in entityFieldItem) {

                        if (entityFieldItem.hasOwnProperty(prop)) {
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
            return entityFieldData;
        },

        created() {
			this.onCreated();
        },

        computed: {
            ...mapGetters({
                entityFieldList: 'getListEntityField'
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
                'findEntityField',
                'updateEntityField',
                'deleteEntityField',
                'createEntityField',
            ]),

            ...mapMutations([
                'addEntityFieldItemToList',
                'deleteEntityFieldFromList',
                'updateEntityFieldById',
            ]),

			onCreated() {
				this.fillEntityFieldFilter();
	            this.fetchEntityFieldData();
			},

            fillEntityFieldFilter() {
                this.entityFieldFilter.CurrentPage = 1;
                this.entityFieldFilter.PerPage = 1000;
            },

            fetchEntityFieldData() {
                return this.findEntityField({
                    filter: this.entityFieldFilter
                });
            },

            /**
             *
             * @param type
             */
            showPanel(type) {
                if (type === this.panel.create) {
                    this.panel.type = this.panel.create;
                    this.clearPanelEntityFieldItem();
                } else if (type === this.panel.edit) {
                    this.panel.type = this.panel.edit;
                    this.currentEntityFieldItem.isSelected = true;
                }

                this.panel.show = true;
            },

            closePanel() {
                this.panel.show = false;
                this.currentEntityFieldItem.isSelected = false;
                this.clearPanelEntityFieldItem();
            },

            selectEntityFieldItem(entityFieldItem) {
                this.showPanel(this.panel.edit);
                this.currentEntityFieldItem.isSelected = true;
                Object.assign(this.currentEntityFieldItem.item, entityFieldItem);
            },

            changeCurrentEntityFieldItem() {
                this.currentEntityFieldItem.hasChange = true;
            },

            cancelChanges() {
                this.clearPanelEntityFieldItem();
                this.closePanel();
            },

            clearPanelEntityFieldItem() {
                this.currentEntityFieldItem.item = new EntityField();
                this.currentEntityFieldItem.hasChange = false;
            },

            saveChangesSubmit() {
                if (this.isPanelCreate) {
                    this.createEntityFieldItemSubmit();
                    return;
                }

                if (this.isPanelEdit) {
                    this.editEntityFieldItemSubmit();
                }
            },

            createEntityFieldItemSubmit() {
                return this.createEntityField({
					data: this.currentEntityFieldItem.item,
                }).then((response) => {

                    if (response.Model) {
                        this.addEntityFieldItemToList(response.Model);
                        this.clearPanelEntityFieldItem();
                    } else {
                        console.error('Ошибка создания записи: ', response.Error);
                    }

                }).catch(error => {
                    console.error('Ошибка создания записи: ', error);
                });
            },

            editEntityFieldItemSubmit() {

                if (this.currentEntityFieldItem.hasChange) {
                    return this.updateEntityField({
                        id: this.currentEntityFieldItem.item.Id,
                        data: this.currentEntityFieldItem.item,
                    }).then((response) => {

                        if (response.Model) {
                            this.updateEntityFieldById(response.Model);
                            this.currentEntityFieldItem.hasChange = false;
                            this.clearPanelEntityFieldItem();
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

            deleteEntityFieldItemHandler() {
                let deletedItemId = this.currentEntityFieldItem.item.Id;

                if (!this.currentEntityFieldItem.canDelete) {
                    this.currentEntityFieldItem.showDeleteConfirmation = true;
                    return;
                }

                this.deleteEntityField({
                    id: deletedItemId
                }).then(response => {

                    if (response.IsSuccess) {
                        this.deleteEntityFieldFromList(deletedItemId);
                        this.clearPanelEntityFieldItem();
                        this.currentEntityFieldItem.canDelete = false;
                        this.currentEntityFieldItem.isSelected = false;
                        this.panel.show = false;
                    } else {
                        console.error('Ошибка удаления элемента: ', response.Error);
                    }

                }).catch(error => {
                    console.error('Ошибка удаления элемента: ', error);
                });
            },

            confirmDeleteHandler() {
                this.currentEntityFieldItem.showDeleteConfirmation = false;
                this.currentEntityFieldItem.canDelete = true;
                this.deleteEntityFieldItemHandler();
            },

            closeConfirmationPanel() {
                this.currentEntityFieldItem.showDeleteConfirmation = false;
            },
        },
    }
</script>

<style lang="scss">

</style>
