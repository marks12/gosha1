
	<template>
    <WorkSpace>
        <template #header>
            <slot name="pageHeader">
                <VHead level="h1">Entity</VHead>
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
                            v-for="entityItem in entityList"
                            :key="entityItem.Id"
                            @click="selectEntityItem(entityItem)"
                            class="sw-table__row_can-select"
                            :class="{'sw-table__row_is-selected': entityItem.Id === currentEntityItem.item.Id}"
                        >
                            <td v-for="(value, key) in fields">
                                <VCheckbox v-if="isCheckbox(entityItem[key])" :checked="entityItem[key]" disabled></VCheckbox>
                                <VText v-else>{{ entityItem[key] }}</VText>
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
                                        :for="`currentEntityItem${key}`"
                                    >{{ filed }}</VLabel>
                                    <VInput
										v-if="isInput(currentEntityItem.item[key])"
                                        v-model="currentEntityItem.item[key]"
                                        width="dyn"
                                        :id="`currentEntityItem${key}`"
                                        @input="changeCurrentEntityItem"
                                    />
									<VCheckbox
										v-if="isCheckbox(currentEntityItem.item[key])"
                                        v-model="currentEntityItem.item[key]"
                                        :id="`currentEntityItem${key}`"
										@input="changeCurrentEntityItem"
									/>
									
                                </VSet>
                            </VSet>
                            <button type="submit" :disabled="!currentEntityItem.hasChange" hidden></button>
                        </form>
                    </template>

                    <template #footer>
                        <VSet>
                            <VButton
                                @click="saveChangesSubmit"
                                accent
                                :text="panelSubmitButtonText"
                                :disabled="!currentEntityItem.hasChange"
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
                    v-if="currentEntityItem.showDeleteConfirmation"
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
                        :disabled="!currentEntityItem.isSelected"
                        @click="deleteEntityItemHandler"
                    />
                </VSet>
            </slot>
        </template>
    </WorkSpace>
</template>

<script>
    import entityData from "../data/EntityData";
    import { Entity } from '../apiModel';
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
        name: 'EntityGen',

        components: {VSelectSimple, VSign, VIcon, VButton, VPanel, VText, VInput, VLabel, VSet, VHead, WorkSpace, VCheckbox},

        props: {
            fields: {
                type: Object,
                default() {
                    const entityItem = new Entity();
                    const fieldsObj = {};

                    for (let prop in entityItem) {

                        if (entityItem.hasOwnProperty(prop)) {
                            fieldsObj[prop] = prop;
                        }

                    }

                    return fieldsObj;
                }
            },
            editFields: {
                type: Object,
                default() {
                    const entityItem = new Entity();
                    const fieldsObj = {};

                    for (let prop in entityItem) {

                        if (entityItem.hasOwnProperty(prop)) {
                            fieldsObj[prop] = prop;
                        }

                    }

                    return fieldsObj;
                }
            }
        },

        data() {
            return entityData;
        },

        created() {
            this.fillEntityFilter();
            this.fetchEntityData();
        },

        computed: {
            ...mapGetters('im3', {
                entityList: 'getListEntity'
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
            ...mapActions('im3', [
                'findEntity',
                'updateEntity',
                'deleteEntity',
                'createEntity',
            ]),

            ...mapMutations('im3', [
                'addEntityItemToList',
                'deleteEntityFromList',
                'updateEntityById',
            ]),

            fillEntityFilter() {
                this.entityFilter.CurrentPage = 1;
                this.entityFilter.PerPage = 1000;
            },

            fetchEntityData() {
                return this.findEntity({
                    filter: this.entityFilter
                });
            },

            /**
             *
             * @param type
             */
            showPanel(type) {
                if (type === this.panel.create) {
                    this.panel.type = this.panel.create;
                    this.clearPanelEntityItem();
                } else if (type === this.panel.edit) {
                    this.panel.type = this.panel.edit;
                    this.currentEntityItem.isSelected = true;
                }

                this.panel.show = true;
            },

            closePanel() {
                this.panel.show = false;
                this.currentEntityItem.isSelected = false;
                this.clearPanelEntityItem();
            },

            selectEntityItem(entityItem) {
                this.showPanel(this.panel.edit);
                this.currentEntityItem.isSelected = true;
                Object.assign(this.currentEntityItem.item, entityItem);
            },

            changeCurrentEntityItem() {
                this.currentEntityItem.hasChange = true;
            },

            cancelChanges() {
                this.clearPanelEntityItem();
                this.closePanel();
            },

            clearPanelEntityItem() {
                this.currentEntityItem.item = new Entity();
                this.currentEntityItem.hasChange = false;
            },

            saveChangesSubmit() {
                if (this.isPanelCreate) {
                    this.createEntityItemSubmit();
                    return;
                }

                if (this.isPanelEdit) {
                    this.editEntityItemSubmit();
                }
            },

            createEntityItemSubmit() {
                this.createEntity({
                    data: {
                        Name: this.currentEntityItem.item.Name,
                        Value: this.currentEntityItem.item.Value,
                        Description: this.currentEntityItem.item.Description,
                    }
                }).then((response) => {

                    if (response.Model) {
                        this.addEntityItemToList(response.Model);
                        this.clearPanelEntityItem();
                    } else {
                        console.error('Ошибка создания записи: ', response.Error);
                    }

                }).catch(error => {
                    console.error('Ошибка создания записи: ', error);
                });
            },

            editEntityItemSubmit() {
                if (this.currentEntityItem.hasChange) {
                    this.updateEntity({
                        id: this.currentEntityItem.item.Id,
                        data: this.currentEntityItem.item,
                    }).then((response) => {

                        if (response.Model) {
                            this.updateEntityById(response.Model);
                            this.currentEntityItem.hasChange = false;
                            this.clearPanelEntityItem();
                            this.closePanel();
                        } else {
                            console.error('Ошибка изменения записи: ', response.Error);
                        }

                    }).catch(error => {
                        console.error('Ошибка изменения записи: ', error);
                    });
                }
            },

            deleteEntityItemHandler() {
                let deletedItemId = this.currentEntityItem.item.Id;

                if (!this.currentEntityItem.canDelete) {
                    this.currentEntityItem.showDeleteConfirmation = true;
                    return;
                }

                this.deleteEntity({
                    id: deletedItemId
                }).then(response => {

                    if (response.IsSuccess) {
                        this.deleteEntityFromList(deletedItemId);
                        this.clearPanelEntityItem();
                        this.currentEntityItem.canDelete = false;
                        this.currentEntityItem.isSelected = false;
                        this.panel.show = false;
                    } else {
                        console.error('Ошибка удаления элемента: ', response.Error);
                    }

                }).catch(error => {
                    console.error('Ошибка удаления элемента: ', error);
                });
            },

            confirmDeleteHandler() {
                this.currentEntityItem.showDeleteConfirmation = false;
                this.currentEntityItem.canDelete = true;
                this.deleteEntityItemHandler();
            },

            closeConfirmationPanel() {
                this.currentEntityItem.showDeleteConfirmation = false;
            },
        },
    }
</script>

<style lang="scss">

</style>
