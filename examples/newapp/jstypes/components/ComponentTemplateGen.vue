
	<template>
    <WorkSpace>
        <template #header>
            <slot name="pageHeader">
                <VHead level="h1">ComponentTemplate</VHead>
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
                            v-for="componentTemplateItem in componentTemplateList"
                            :key="componentTemplateItem.Id"
                            @click="selectComponentTemplateItem(componentTemplateItem)"
                            class="sw-table__row_can-select"
                            :class="{'sw-table__row_is-selected': componentTemplateItem.Id === currentComponentTemplateItem.item.Id}"
                        >
                            <td v-for="(value, key) in fields" :key="key + '-fields'">
                                <VCheckbox v-if="isCheckbox(componentTemplateItem[key])" :checked="componentTemplateItem[key]" disabled></VCheckbox>
                                <VText v-else>{{ componentTemplateItem[key] }}</VText>
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
                                        :for="`currentComponentTemplateItem${key}`"
                                    >{{ filed }}</VLabel>
                                    <VInput
										v-if="isInput(currentComponentTemplateItem.item[key])"
                                        v-model="currentComponentTemplateItem.item[key]"
                                        width="dyn"
                                        :id="`currentComponentTemplateItem${key}`"
                                        @input="changeCurrentComponentTemplateItem"
                                    />
									<VCheckbox
										v-if="isCheckbox(currentComponentTemplateItem.item[key])"
                                        v-model="currentComponentTemplateItem.item[key]"
                                        :id="`currentComponentTemplateItem${key}`"
										@input="changeCurrentComponentTemplateItem"
									/>
									
                                </VSet>
                            </VSet>
                            <button type="submit" :disabled="!currentComponentTemplateItem.hasChange" hidden></button>
                        </form>
                    </template>

                    <template #footer>
                        <VSet>
                            <VButton
                                @click="saveChangesSubmit"
                                accent
                                :text="panelSubmitButtonText"
                                :disabled="!currentComponentTemplateItem.hasChange"
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
                    v-if="currentComponentTemplateItem.showDeleteConfirmation"
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
                        :disabled="!currentComponentTemplateItem.isSelected"
                        @click="deleteComponentTemplateItemHandler"
                    />
                </VSet>
            </slot>
        </template>
    </WorkSpace>
</template>

<script>
    import componentTemplateData from "../data/ComponentTemplateData";
    import { ComponentTemplate } from '../apiModel';
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
        name: 'ComponentTemplateGen',

        components: {VSelectSimple, VSign, VIcon, VButton, VPanel, VText, VInput, VLabel, VSet, VHead, WorkSpace, VCheckbox},

        props: {
            fields: {
                type: Object,
                default() {
                    const componentTemplateItem = new ComponentTemplate();
                    const fieldsObj = {};

                    for (let prop in componentTemplateItem) {

                        if (componentTemplateItem.hasOwnProperty(prop)) {
                            fieldsObj[prop] = prop;
                        }

                    }

                    return fieldsObj;
                }
            },
            editFields: {
                type: Object,
                default() {
                    const componentTemplateItem = new ComponentTemplate();
                    const fieldsObj = {};

                    for (let prop in componentTemplateItem) {

                        if (componentTemplateItem.hasOwnProperty(prop)) {
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
            return componentTemplateData;
        },

        created() {
            this.fillComponentTemplateFilter();
            this.fetchComponentTemplateData();
        },

        computed: {
            ...mapGetters({
                componentTemplateList: 'getListComponentTemplate'
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
                'findComponentTemplate',
                'updateComponentTemplate',
                'deleteComponentTemplate',
                'createComponentTemplate',
            ]),

            ...mapMutations([
                'addComponentTemplateItemToList',
                'deleteComponentTemplateFromList',
                'updateComponentTemplateById',
            ]),

            fillComponentTemplateFilter() {
                this.componentTemplateFilter.CurrentPage = 1;
                this.componentTemplateFilter.PerPage = 1000;
            },

            fetchComponentTemplateData() {
                return this.findComponentTemplate({
                    filter: this.componentTemplateFilter
                });
            },

            /**
             *
             * @param type
             */
            showPanel(type) {
                if (type === this.panel.create) {
                    this.panel.type = this.panel.create;
                    this.clearPanelComponentTemplateItem();
                } else if (type === this.panel.edit) {
                    this.panel.type = this.panel.edit;
                    this.currentComponentTemplateItem.isSelected = true;
                }

                this.panel.show = true;
            },

            closePanel() {
                this.panel.show = false;
                this.currentComponentTemplateItem.isSelected = false;
                this.clearPanelComponentTemplateItem();
            },

            selectComponentTemplateItem(componentTemplateItem) {
                this.showPanel(this.panel.edit);
                this.currentComponentTemplateItem.isSelected = true;
                Object.assign(this.currentComponentTemplateItem.item, componentTemplateItem);
            },

            changeCurrentComponentTemplateItem() {
                this.currentComponentTemplateItem.hasChange = true;
            },

            cancelChanges() {
                this.clearPanelComponentTemplateItem();
                this.closePanel();
            },

            clearPanelComponentTemplateItem() {
                this.currentComponentTemplateItem.item = new ComponentTemplate();
                this.currentComponentTemplateItem.hasChange = false;
            },

            saveChangesSubmit() {
                if (this.isPanelCreate) {
                    this.createComponentTemplateItemSubmit();
                    return;
                }

                if (this.isPanelEdit) {
                    this.editComponentTemplateItemSubmit();
                }
            },

            createComponentTemplateItemSubmit() {
                return this.createComponentTemplate({
					data: this.currentComponentTemplateItem.item,
                }).then((response) => {

                    if (response.Model) {
                        this.addComponentTemplateItemToList(response.Model);
                        this.clearPanelComponentTemplateItem();
                    } else {
                        console.error('Ошибка создания записи: ', response.Error);
                    }

                }).catch(error => {
                    console.error('Ошибка создания записи: ', error);
                });
            },

            editComponentTemplateItemSubmit() {

                if (this.currentComponentTemplateItem.hasChange) {
                    return this.updateComponentTemplate({
                        id: this.currentComponentTemplateItem.item.Id,
                        data: this.currentComponentTemplateItem.item,
                    }).then((response) => {

                        if (response.Model) {
                            this.updateComponentTemplateById(response.Model);
                            this.currentComponentTemplateItem.hasChange = false;
                            this.clearPanelComponentTemplateItem();
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

            deleteComponentTemplateItemHandler() {
                let deletedItemId = this.currentComponentTemplateItem.item.Id;

                if (!this.currentComponentTemplateItem.canDelete) {
                    this.currentComponentTemplateItem.showDeleteConfirmation = true;
                    return;
                }

                this.deleteComponentTemplate({
                    id: deletedItemId
                }).then(response => {

                    if (response.IsSuccess) {
                        this.deleteComponentTemplateFromList(deletedItemId);
                        this.clearPanelComponentTemplateItem();
                        this.currentComponentTemplateItem.canDelete = false;
                        this.currentComponentTemplateItem.isSelected = false;
                        this.panel.show = false;
                    } else {
                        console.error('Ошибка удаления элемента: ', response.Error);
                    }

                }).catch(error => {
                    console.error('Ошибка удаления элемента: ', error);
                });
            },

            confirmDeleteHandler() {
                this.currentComponentTemplateItem.showDeleteConfirmation = false;
                this.currentComponentTemplateItem.canDelete = true;
                this.deleteComponentTemplateItemHandler();
            },

            closeConfirmationPanel() {
                this.currentComponentTemplateItem.showDeleteConfirmation = false;
            },
        },
    }
</script>

<style lang="scss">

</style>
