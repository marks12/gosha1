
	<template>
    <WorkSpace>
        <template #header>
            <slot name="pageHeader">
                <VHead level="h1">BuLayer</VHead>
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
                            v-for="buLayerItem in buLayerList"
                            :key="buLayerItem.Id"
                            @click="selectBuLayerItem(buLayerItem)"
                            class="sw-table__row_can-select"
                            :class="{'sw-table__row_is-selected': buLayerItem.Id === currentBuLayerItem.item.Id}"
                        >
                            <td v-for="(value, key) in fields" :key="key + '-fields'">
                                <VCheckbox v-if="isCheckbox(buLayerItem[key])" :checked="buLayerItem[key]" disabled></VCheckbox>
                                <VText v-else>{{ buLayerItem[key] }}</VText>
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
                                        :for="`currentBuLayerItem${key}`"
                                    >{{ filed }}</VLabel>
                                    <VInput
										v-if="isInput(currentBuLayerItem.item[key])"
                                        v-model="currentBuLayerItem.item[key]"
                                        width="dyn"
                                        :id="`currentBuLayerItem${key}`"
                                        @input="changeCurrentBuLayerItem"
                                    />
									<VCheckbox
										v-if="isCheckbox(currentBuLayerItem.item[key])"
                                        v-model="currentBuLayerItem.item[key]"
                                        :id="`currentBuLayerItem${key}`"
										@input="changeCurrentBuLayerItem"
									/>
									
                                </VSet>
                            </VSet>
                            <button type="submit" :disabled="!currentBuLayerItem.hasChange" hidden></button>
                        </form>
                    </template>

                    <template #footer>
                        <VSet>
                            <VButton
                                @click="saveChangesSubmit"
                                accent
                                :text="panelSubmitButtonText"
                                :disabled="!currentBuLayerItem.hasChange"
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
                    v-if="currentBuLayerItem.showDeleteConfirmation"
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
                        :disabled="!currentBuLayerItem.isSelected"
                        @click="deleteBuLayerItemHandler"
                    />
                </VSet>
            </slot>
        </template>
    </WorkSpace>
</template>

<script>
    import buLayerData from "../data/BuLayerData";
    import { BuLayer } from '../apiModel';
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
        name: 'BuLayerGen',

        components: {VButton, VPanel, VText, VInput, VLabel, VSet, VHead, WorkSpace, VCheckbox},

        props: {
            fields: {
                type: Object,
                default() {
                    const buLayerItem = new BuLayer();
                    const fieldsObj = {};

                    for (let prop in buLayerItem) {

                        if (buLayerItem.hasOwnProperty(prop)) {
                            fieldsObj[prop] = prop;
                        }

                    }

                    return fieldsObj;
                }
            },
            editFields: {
                type: Object,
                default() {
                    const buLayerItem = new BuLayer();
                    const fieldsObj = {};

                    for (let prop in buLayerItem) {

                        if (buLayerItem.hasOwnProperty(prop)) {
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
            return buLayerData;
        },

        created() {
			this.onCreated();
        },

        computed: {
            ...mapGetters({
                buLayerList: 'getListBuLayer'
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
                'findBuLayer',
                'updateBuLayer',
                'deleteBuLayer',
                'createBuLayer',
            ]),

            ...mapMutations([
                'addBuLayerItemToList',
                'deleteBuLayerFromList',
                'updateBuLayerById',
            ]),

			onCreated() {
				this.fillBuLayerFilter();
	            this.fetchBuLayerData();
			},

            fillBuLayerFilter() {
                this.buLayerFilter.CurrentPage = 1;
                this.buLayerFilter.PerPage = 1000;
            },

            fetchBuLayerData() {
                return this.findBuLayer({
                    filter: this.buLayerFilter
                });
            },

            /**
             *
             * @param type
             */
            showPanel(type) {
                if (type === this.panel.create) {
                    this.panel.type = this.panel.create;
                    this.clearPanelBuLayerItem();
                } else if (type === this.panel.edit) {
                    this.panel.type = this.panel.edit;
                    this.currentBuLayerItem.isSelected = true;
                }

                this.panel.show = true;
            },

            closePanel() {
                this.panel.show = false;
                this.currentBuLayerItem.isSelected = false;
                this.clearPanelBuLayerItem();
            },

            selectBuLayerItem(buLayerItem) {
                this.showPanel(this.panel.edit);
                this.currentBuLayerItem.isSelected = true;
                Object.assign(this.currentBuLayerItem.item, buLayerItem);
            },

            changeCurrentBuLayerItem() {
                this.currentBuLayerItem.hasChange = true;
            },

            cancelChanges() {
                this.clearPanelBuLayerItem();
                this.closePanel();
            },

            clearPanelBuLayerItem() {
                this.currentBuLayerItem.item = new BuLayer();
                this.currentBuLayerItem.hasChange = false;
            },

            saveChangesSubmit() {
                if (this.isPanelCreate) {
                    this.createBuLayerItemSubmit();
                    return;
                }

                if (this.isPanelEdit) {
                    this.editBuLayerItemSubmit();
                }
            },

            createBuLayerItemSubmit() {
                return this.createBuLayer({
					data: this.currentBuLayerItem.item,
                }).then((response) => {

                    if (response.Model) {
                        this.addBuLayerItemToList(response.Model);
                        this.clearPanelBuLayerItem();
                    } else {
                        console.error('Ошибка создания записи: ', response.Error);
                    }

                }).catch(error => {
                    console.error('Ошибка создания записи: ', error);
                });
            },

            editBuLayerItemSubmit() {

                if (this.currentBuLayerItem.hasChange) {
                    return this.updateBuLayer({
                        id: this.currentBuLayerItem.item.Id,
                        data: this.currentBuLayerItem.item,
                    }).then((response) => {

                        if (response.Model) {
                            this.updateBuLayerById(response.Model);
                            this.currentBuLayerItem.hasChange = false;
                            this.clearPanelBuLayerItem();
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

            deleteBuLayerItemHandler() {
                let deletedItemId = this.currentBuLayerItem.item.Id;

                if (!this.currentBuLayerItem.canDelete) {
                    this.currentBuLayerItem.showDeleteConfirmation = true;
                    return;
                }

                this.deleteBuLayer({
                    id: deletedItemId
                }).then(response => {

                    if (response.IsSuccess) {
                        this.deleteBuLayerFromList(deletedItemId);
                        this.clearPanelBuLayerItem();
                        this.currentBuLayerItem.canDelete = false;
                        this.currentBuLayerItem.isSelected = false;
                        this.panel.show = false;
                    } else {
                        console.error('Ошибка удаления элемента: ', response.Error);
                    }

                }).catch(error => {
                    console.error('Ошибка удаления элемента: ', error);
                });
            },

            confirmDeleteHandler() {
                this.currentBuLayerItem.showDeleteConfirmation = false;
                this.currentBuLayerItem.canDelete = true;
                this.deleteBuLayerItemHandler();
            },

            closeConfirmationPanel() {
                this.currentBuLayerItem.showDeleteConfirmation = false;
            },
        },
    }
</script>

<style lang="scss">

</style>
