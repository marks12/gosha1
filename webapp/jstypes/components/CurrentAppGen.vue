
	<template>
    <WorkSpace>
        <template #header>
            <slot name="pageHeader">
                <VHead level="h1">CurrentApp</VHead>
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
                            v-for="currentAppItem in currentAppList"
                            :key="currentAppItem.Id"
                            @click="selectCurrentAppItem(currentAppItem)"
                            class="sw-table__row_can-select"
                            :class="{'sw-table__row_is-selected': currentAppItem.Id === currentCurrentAppItem.item.Id}"
                        >
                            <td v-for="(value, key) in fields" :key="key + '-fields'">
                                <VCheckbox v-if="isCheckbox(currentAppItem[key])" :checked="currentAppItem[key]" disabled></VCheckbox>
                                <VText v-else>{{ currentAppItem[key] }}</VText>
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
                                        :for="`currentCurrentAppItem${key}`"
                                    >{{ filed }}</VLabel>
                                    <VInput
										v-if="isInput(currentCurrentAppItem.item[key])"
                                        v-model="currentCurrentAppItem.item[key]"
                                        width="dyn"
                                        :id="`currentCurrentAppItem${key}`"
                                        @input="changeCurrentCurrentAppItem"
                                    />
									<VCheckbox
										v-if="isCheckbox(currentCurrentAppItem.item[key])"
                                        v-model="currentCurrentAppItem.item[key]"
                                        :id="`currentCurrentAppItem${key}`"
										@input="changeCurrentCurrentAppItem"
									/>
									
                                </VSet>
                            </VSet>
                            <button type="submit" :disabled="!currentCurrentAppItem.hasChange" hidden></button>
                        </form>
                    </template>

                    <template #footer>
                        <VSet>
                            <VButton
                                @click="saveChangesSubmit"
                                accent
                                :text="panelSubmitButtonText"
                                :disabled="!currentCurrentAppItem.hasChange"
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
                    v-if="currentCurrentAppItem.showDeleteConfirmation"
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
                        :disabled="!currentCurrentAppItem.isSelected"
                        @click="deleteCurrentAppItemHandler"
                    />
                </VSet>
            </slot>
        </template>
    </WorkSpace>
</template>

<script>
    import currentAppData from "../data/CurrentAppData";
    import { CurrentApp } from '../apiModel';
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
        name: 'CurrentAppGen',

        components: {VButton, VPanel, VText, VInput, VLabel, VSet, VHead, WorkSpace, VCheckbox},

        props: {
            fields: {
                type: Object,
                default() {
                    const currentAppItem = new CurrentApp();
                    const fieldsObj = {};

                    for (let prop in currentAppItem) {

                        if (currentAppItem.hasOwnProperty(prop)) {
                            fieldsObj[prop] = prop;
                        }

                    }

                    return fieldsObj;
                }
            },
            editFields: {
                type: Object,
                default() {
                    const currentAppItem = new CurrentApp();
                    const fieldsObj = {};

                    for (let prop in currentAppItem) {

                        if (currentAppItem.hasOwnProperty(prop)) {
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
            return currentAppData;
        },

        created() {
			this.onCreated();
        },

        computed: {
            ...mapGetters('gosha', {
                currentAppList: 'getListCurrentApp'
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
                'findCurrentApp',
                'updateCurrentApp',
                'deleteCurrentApp',
                'createCurrentApp',
            ]),

            ...mapMutations('gosha', [
                'addCurrentAppItemToList',
                'deleteCurrentAppFromList',
                'updateCurrentAppById',
            ]),

			onCreated() {
				this.fillCurrentAppFilter();
	            this.fetchCurrentAppData();
			},

            fillCurrentAppFilter() {
                this.currentAppFilter.CurrentPage = 1;
                this.currentAppFilter.PerPage = 1000;
            },

            fetchCurrentAppData() {
                return this.findCurrentApp({
                    filter: this.currentAppFilter
                });
            },

            /**
             *
             * @param type
             */
            showPanel(type) {
                if (type === this.panel.create) {
                    this.panel.type = this.panel.create;
                    this.clearPanelCurrentAppItem();
                } else if (type === this.panel.edit) {
                    this.panel.type = this.panel.edit;
                    this.currentCurrentAppItem.isSelected = true;
                }

                this.panel.show = true;
            },

            closePanel() {
                this.panel.show = false;
                this.currentCurrentAppItem.isSelected = false;
                this.clearPanelCurrentAppItem();
            },

            selectCurrentAppItem(currentAppItem) {
                this.showPanel(this.panel.edit);
                this.currentCurrentAppItem.isSelected = true;
                Object.assign(this.currentCurrentAppItem.item, currentAppItem);
            },

            changeCurrentCurrentAppItem() {
                this.currentCurrentAppItem.hasChange = true;
            },

            cancelChanges() {
                this.clearPanelCurrentAppItem();
                this.closePanel();
            },

            clearPanelCurrentAppItem() {
                this.currentCurrentAppItem.item = new CurrentApp();
                this.currentCurrentAppItem.hasChange = false;
            },

            saveChangesSubmit() {
                if (this.isPanelCreate) {
                    this.createCurrentAppItemSubmit();
                    return;
                }

                if (this.isPanelEdit) {
                    this.editCurrentAppItemSubmit();
                }
            },

            createCurrentAppItemSubmit() {
                return this.createCurrentApp({
					data: this.currentCurrentAppItem.item,
                }).then((response) => {

                    if (response.Model) {
                        this.addCurrentAppItemToList(response.Model);
                        this.clearPanelCurrentAppItem();
                    } else {
                        console.error('Ошибка создания записи: ', response.Error);
                    }

                }).catch(error => {
                    console.error('Ошибка создания записи: ', error);
                });
            },

            editCurrentAppItemSubmit() {

                if (this.currentCurrentAppItem.hasChange) {
                    return this.updateCurrentApp({
                        id: this.currentCurrentAppItem.item.Id,
                        data: this.currentCurrentAppItem.item,
                    }).then((response) => {

                        if (response.Model) {
                            this.updateCurrentAppById(response.Model);
                            this.currentCurrentAppItem.hasChange = false;
                            this.clearPanelCurrentAppItem();
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

            deleteCurrentAppItemHandler() {
                let deletedItemId = this.currentCurrentAppItem.item.Id;

                if (!this.currentCurrentAppItem.canDelete) {
                    this.currentCurrentAppItem.showDeleteConfirmation = true;
                    return;
                }

                this.deleteCurrentApp({
                    id: deletedItemId
                }).then(response => {

                    if (response.IsSuccess) {
                        this.deleteCurrentAppFromList(deletedItemId);
                        this.clearPanelCurrentAppItem();
                        this.currentCurrentAppItem.canDelete = false;
                        this.currentCurrentAppItem.isSelected = false;
                        this.panel.show = false;
                    } else {
                        console.error('Ошибка удаления элемента: ', response.Error);
                    }

                }).catch(error => {
                    console.error('Ошибка удаления элемента: ', error);
                });
            },

            confirmDeleteHandler() {
                this.currentCurrentAppItem.showDeleteConfirmation = false;
                this.currentCurrentAppItem.canDelete = true;
                this.deleteCurrentAppItemHandler();
            },

            closeConfirmationPanel() {
                this.currentCurrentAppItem.showDeleteConfirmation = false;
            },
        },
    }
</script>

<style lang="scss">

</style>
