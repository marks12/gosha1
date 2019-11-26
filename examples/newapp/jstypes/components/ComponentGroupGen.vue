
	<template>
    <WorkSpace>
        <template #header>
            <slot name="pageHeader">
                <VHead level="h1">ComponentGroup</VHead>
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
                            v-for="componentGroupItem in componentGroupList"
                            :key="componentGroupItem.Id"
                            @click="selectComponentGroupItem(componentGroupItem)"
                            class="sw-table__row_can-select"
                            :class="{'sw-table__row_is-selected': componentGroupItem.Id === currentComponentGroupItem.item.Id}"
                        >
                            <td v-for="(value, key) in fields" :key="key + '-fields'">
                                <VCheckbox v-if="isCheckbox(componentGroupItem[key])" :checked="componentGroupItem[key]" disabled></VCheckbox>
                                <VText v-else>{{ componentGroupItem[key] }}</VText>
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
                                        :for="`currentComponentGroupItem${key}`"
                                    >{{ filed }}</VLabel>
                                    <VInput
										v-if="isInput(currentComponentGroupItem.item[key])"
                                        v-model="currentComponentGroupItem.item[key]"
                                        width="dyn"
                                        :id="`currentComponentGroupItem${key}`"
                                        @input="changeCurrentComponentGroupItem"
                                    />
									<VCheckbox
										v-if="isCheckbox(currentComponentGroupItem.item[key])"
                                        v-model="currentComponentGroupItem.item[key]"
                                        :id="`currentComponentGroupItem${key}`"
										@input="changeCurrentComponentGroupItem"
									/>
									
                                </VSet>
                            </VSet>
                            <button type="submit" :disabled="!currentComponentGroupItem.hasChange" hidden></button>
                        </form>
                    </template>

                    <template #footer>
                        <VSet>
                            <VButton
                                @click="saveChangesSubmit"
                                accent
                                :text="panelSubmitButtonText"
                                :disabled="!currentComponentGroupItem.hasChange"
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
                    v-if="currentComponentGroupItem.showDeleteConfirmation"
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
                        :disabled="!currentComponentGroupItem.isSelected"
                        @click="deleteComponentGroupItemHandler"
                    />
                </VSet>
            </slot>
        </template>
    </WorkSpace>
</template>

<script>
    import componentGroupData from "../data/ComponentGroupData";
    import { ComponentGroup } from '../apiModel';
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
        name: 'ComponentGroupGen',

        components: {VSelectSimple, VSign, VIcon, VButton, VPanel, VText, VInput, VLabel, VSet, VHead, WorkSpace, VCheckbox},

        props: {
            fields: {
                type: Object,
                default() {
                    const componentGroupItem = new ComponentGroup();
                    const fieldsObj = {};

                    for (let prop in componentGroupItem) {

                        if (componentGroupItem.hasOwnProperty(prop)) {
                            fieldsObj[prop] = prop;
                        }

                    }

                    return fieldsObj;
                }
            },
            editFields: {
                type: Object,
                default() {
                    const componentGroupItem = new ComponentGroup();
                    const fieldsObj = {};

                    for (let prop in componentGroupItem) {

                        if (componentGroupItem.hasOwnProperty(prop)) {
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
            return componentGroupData;
        },

        created() {
            this.fillComponentGroupFilter();
            this.fetchComponentGroupData();
        },

        computed: {
            ...mapGetters({
                componentGroupList: 'getListComponentGroup'
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
                'findComponentGroup',
                'updateComponentGroup',
                'deleteComponentGroup',
                'createComponentGroup',
            ]),

            ...mapMutations([
                'addComponentGroupItemToList',
                'deleteComponentGroupFromList',
                'updateComponentGroupById',
            ]),

            fillComponentGroupFilter() {
                this.componentGroupFilter.CurrentPage = 1;
                this.componentGroupFilter.PerPage = 1000;
            },

            fetchComponentGroupData() {
                return this.findComponentGroup({
                    filter: this.componentGroupFilter
                });
            },

            /**
             *
             * @param type
             */
            showPanel(type) {
                if (type === this.panel.create) {
                    this.panel.type = this.panel.create;
                    this.clearPanelComponentGroupItem();
                } else if (type === this.panel.edit) {
                    this.panel.type = this.panel.edit;
                    this.currentComponentGroupItem.isSelected = true;
                }

                this.panel.show = true;
            },

            closePanel() {
                this.panel.show = false;
                this.currentComponentGroupItem.isSelected = false;
                this.clearPanelComponentGroupItem();
            },

            selectComponentGroupItem(componentGroupItem) {
                this.showPanel(this.panel.edit);
                this.currentComponentGroupItem.isSelected = true;
                Object.assign(this.currentComponentGroupItem.item, componentGroupItem);
            },

            changeCurrentComponentGroupItem() {
                this.currentComponentGroupItem.hasChange = true;
            },

            cancelChanges() {
                this.clearPanelComponentGroupItem();
                this.closePanel();
            },

            clearPanelComponentGroupItem() {
                this.currentComponentGroupItem.item = new ComponentGroup();
                this.currentComponentGroupItem.hasChange = false;
            },

            saveChangesSubmit() {
                if (this.isPanelCreate) {
                    this.createComponentGroupItemSubmit();
                    return;
                }

                if (this.isPanelEdit) {
                    this.editComponentGroupItemSubmit();
                }
            },

            createComponentGroupItemSubmit() {
                return this.createComponentGroup({
					data: this.currentComponentGroupItem.item,
                }).then((response) => {

                    if (response.Model) {
                        this.addComponentGroupItemToList(response.Model);
                        this.clearPanelComponentGroupItem();
                    } else {
                        console.error('Ошибка создания записи: ', response.Error);
                    }

                }).catch(error => {
                    console.error('Ошибка создания записи: ', error);
                });
            },

            editComponentGroupItemSubmit() {

                if (this.currentComponentGroupItem.hasChange) {
                    return this.updateComponentGroup({
                        id: this.currentComponentGroupItem.item.Id,
                        data: this.currentComponentGroupItem.item,
                    }).then((response) => {

                        if (response.Model) {
                            this.updateComponentGroupById(response.Model);
                            this.currentComponentGroupItem.hasChange = false;
                            this.clearPanelComponentGroupItem();
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

            deleteComponentGroupItemHandler() {
                let deletedItemId = this.currentComponentGroupItem.item.Id;

                if (!this.currentComponentGroupItem.canDelete) {
                    this.currentComponentGroupItem.showDeleteConfirmation = true;
                    return;
                }

                this.deleteComponentGroup({
                    id: deletedItemId
                }).then(response => {

                    if (response.IsSuccess) {
                        this.deleteComponentGroupFromList(deletedItemId);
                        this.clearPanelComponentGroupItem();
                        this.currentComponentGroupItem.canDelete = false;
                        this.currentComponentGroupItem.isSelected = false;
                        this.panel.show = false;
                    } else {
                        console.error('Ошибка удаления элемента: ', response.Error);
                    }

                }).catch(error => {
                    console.error('Ошибка удаления элемента: ', error);
                });
            },

            confirmDeleteHandler() {
                this.currentComponentGroupItem.showDeleteConfirmation = false;
                this.currentComponentGroupItem.canDelete = true;
                this.deleteComponentGroupItemHandler();
            },

            closeConfirmationPanel() {
                this.currentComponentGroupItem.showDeleteConfirmation = false;
            },
        },
    }
</script>

<style lang="scss">

</style>
