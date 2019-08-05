
	<template>
    <WorkSpace>
        <template #header>
            <slot name="pageHeader">
                <VHead level="h1">EntityFieldFilter</VHead>
            </slot>
        </template>

        <template #content>
            <slot name="data">
                <table>
                    <thead>
                        <tr>
                            <th v-for="header in fields">{{ header }}</th>
                        </tr>
                    </thead>
            
                    <tbody>
                        <tr
                            v-for="entityFieldFilterItem in entityFieldFilterList"
                            :key="entityFieldFilterItem.Id"
                            @click="selectEntityFieldFilterItem(entityFieldFilterItem)"
                            class="sw-table__row_can-select"
                            :class="{'sw-table__row_is-selected': entityFieldFilterItem.Id === currentEntityFieldFilterItem.item.Id}"
                        >
                            <td v-for="(value, key) in fields">
                                <VCheckbox v-if="isCheckbox(applicationItem[key])" :checked="applicationItem[key]" disabled></VCheckbox>
                                <VText v-else>{{ applicationItem[key] }}</VText>
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
                                    v-for="(filed, key) in editFields"
                                    vertical-align="center"
                                >
                                    <VLabel
                                        width="col4"
                                        :for="`currentEntityFieldFilterItem${key}`"
                                    >{{ filed }}</VLabel>
                                    <VInput
										v-if="isInput(currentEntityFieldFilterItem.item[key])"
                                        v-model="currentEntityFieldFilterItem.item[key]"
                                        width="dyn"
                                        :id="`currentEntityFieldFilterItem${key}`"
                                        @input="changeCurrentEntityFieldFilterItem"
                                    />
									<VCheckbox
										v-if="isCheckbox(currentEntityFieldFilterItem.item[key])"
                                        v-model="currentEntityFieldFilterItem.item[key]"
                                        :id="`currentEntityFieldFilterItem${key}`"
										@input="changeCurrentApplicationItem"
									/>
									
                                </VSet>
                            </VSet>
                            <button type="submit" :disabled="!currentEntityFieldFilterItem.hasChange" hidden></button>
                        </form>
                    </template>

                    <template #footer>
                        <VSet>
                            <VButton
                                @click="saveChangesSubmit"
                                accent
                                :text="panelSubmitButtonText"
                                :disabled="!currentEntityFieldFilterItem.hasChange"
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
                    v-if="currentEntityFieldFilterItem.showDeleteConfirmation"
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
                        text="Добавить"
                        accent
                        @click="showPanel(panel.create)"
                    />
                    <VButton
                        text="Удалить"
                        :disabled="!currentEntityFieldFilterItem.isSelected"
                        @click="deleteEntityFieldFilterItemHandler"
                    />
                </VSet>
            </slot>
        </template>
    </WorkSpace>
</template>

<script>
    import entityFieldFilterData from "../data/EntityFieldFilterData";
    import { EntityFieldFilter } from '../apiModel';
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
    import VSelect from "swui/src/components/VSelect";

    export default {
        name: 'EntityFieldFilterGen',

        components: {VSelect, VSign, VIcon, VButton, VPanel, VText, VInput, VLabel, VSet, VHead, WorkSpace, VCheckbox},

        props: {
            fields: {
                type: Object,
                default() {
                    const entityFieldFilterItem = new EntityFieldFilter();
                    const fieldsObj = {};

                    for (let prop in entityFieldFilterItem) {

                        if (entityFieldFilterItem.hasOwnProperty(prop)) {
                            fieldsObj[prop] = prop;
                        }

                    }

                    return fieldsObj;
                }
            },
            editFields: {
                type: Object,
                default() {
                    const entityFieldFilterItem = new EntityFieldFilter();
                    const fieldsObj = {};

                    for (let prop in entityFieldFilterItem) {

                        if (entityFieldFilterItem.hasOwnProperty(prop)) {
                            fieldsObj[prop] = prop;
                        }

                    }

                    return fieldsObj;
                }
            }
        },

        data() {
            return entityFieldFilterData;
        },

        created() {
            this.fillEntityFieldFilterFilter();
            this.fetchEntityFieldFilterData();
        },

        computed: {
            ...mapGetters({
                entityFieldFilterList: 'getListEntityFieldFilter'
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
                'findEntityFieldFilter',
                'updateEntityFieldFilter',
                'deleteEntityFieldFilter',
                'createEntityFieldFilter',
            ]),

            ...mapMutations([
                'addEntityFieldFilterItemToList',
                'deleteEntityFieldFilterFromList',
                'updateEntityFieldFilterById',
            ]),

            fillEntityFieldFilterFilter() {
                this.entityFieldFilterFilter.CurrentPage = 1;
                this.entityFieldFilterFilter.PerPage = 1000;
            },

            fetchEntityFieldFilterData() {
                return this.findEntityFieldFilter({
                    filter: this.entityFieldFilterFilter
                });
            },

            /**
             *
             * @param type
             */
            showPanel(type) {
                if (type === this.panel.create) {
                    this.panel.type = this.panel.create;
                    this.clearPanelEntityFieldFilterItem();
                } else if (type === this.panel.edit) {
                    this.panel.type = this.panel.edit;
                    this.currentEntityFieldFilterItem.isSelected = true;
                }

                this.panel.show = true;
            },

            closePanel() {
                this.panel.show = false;
                this.currentEntityFieldFilterItem.isSelected = false;
                this.clearPanelEntityFieldFilterItem();
            },

            selectEntityFieldFilterItem(entityFieldFilterItem) {
                this.showPanel(this.panel.edit);
                this.currentEntityFieldFilterItem.isSelected = true;
                Object.assign(this.currentEntityFieldFilterItem.item, entityFieldFilterItem);
            },

            changeCurrentEntityFieldFilterItem() {
                this.currentEntityFieldFilterItem.hasChange = true;
            },

            cancelChanges() {
                this.clearPanelEntityFieldFilterItem();
                this.closePanel();
            },

            clearPanelEntityFieldFilterItem() {
                this.currentEntityFieldFilterItem.item = new EntityFieldFilter();
                this.currentEntityFieldFilterItem.hasChange = false;
            },

            saveChangesSubmit() {
                if (this.isPanelCreate) {
                    this.createEntityFieldFilterItemSubmit();
                    return;
                }

                if (this.isPanelEdit) {
                    this.editEntityFieldFilterItemSubmit();
                }
            },

            createEntityFieldFilterItemSubmit() {
                this.createEntityFieldFilter({
                    data: {
                        Name: this.currentEntityFieldFilterItem.item.Name,
                        Value: this.currentEntityFieldFilterItem.item.Value,
                        Description: this.currentEntityFieldFilterItem.item.Description,
                    }
                }).then((response) => {

                    if (response.Model) {
                        this.addEntityFieldFilterItemToList(response.Model);
                        this.clearPanelEntityFieldFilterItem();
                    } else {
                        console.error('Ошибка создания записи: ', response.Error);
                    }

                }).catch(error => {
                    console.error('Ошибка создания записи: ', error);
                });
            },

            editEntityFieldFilterItemSubmit() {
                if (this.currentEntityFieldFilterItem.hasChange) {
                    this.updateEntityFieldFilter({
                        id: this.currentEntityFieldFilterItem.item.Id,
                        data: this.currentEntityFieldFilterItem.item,
                    }).then((response) => {

                        if (response.Model) {
                            this.updateEntityFieldFilterById(response.Model);
                            this.currentEntityFieldFilterItem.hasChange = false;
                            this.clearPanelEntityFieldFilterItem();
                            this.closePanel();
                        } else {
                            console.error('Ошибка изменения записи: ', response.Error);
                        }

                    }).catch(error => {
                        console.error('Ошибка изменения записи: ', error);
                    });
                }
            },

            deleteEntityFieldFilterItemHandler() {
                let deletedItemId = this.currentEntityFieldFilterItem.item.Id;

                if (!this.currentEntityFieldFilterItem.canDelete) {
                    this.currentEntityFieldFilterItem.showDeleteConfirmation = true;
                    return;
                }

                this.deleteEntityFieldFilter({
                    id: deletedItemId
                }).then(response => {

                    if (response.IsSuccess) {
                        this.deleteEntityFieldFilterFromList(deletedItemId);
                        this.clearPanelEntityFieldFilterItem();
                        this.currentEntityFieldFilterItem.canDelete = false;
                        this.currentEntityFieldFilterItem.isSelected = false;
                        this.panel.show = false;
                    } else {
                        console.error('Ошибка удаления элемента: ', response.Error);
                    }

                }).catch(error => {
                    console.error('Ошибка удаления элемента: ', error);
                });
            },

            confirmDeleteHandler() {
                this.currentEntityFieldFilterItem.showDeleteConfirmation = false;
                this.currentEntityFieldFilterItem.canDelete = true;
                this.deleteEntityFieldFilterItemHandler();
            },

            closeConfirmationPanel() {
                this.currentEntityFieldFilterItem.showDeleteConfirmation = false;
            },
        },
    }
</script>

<style lang="scss">

</style>
