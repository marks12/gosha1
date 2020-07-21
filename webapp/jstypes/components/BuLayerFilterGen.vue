
	<template>
    <WorkSpace>
        <template #header>
            <slot name="pageHeader">
                <VHead level="h1">BuLayerFilter</VHead>
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
                            v-for="buLayerFilterItem in buLayerFilterList"
                            :key="buLayerFilterItem.Id"
                            @click="selectBuLayerFilterItem(buLayerFilterItem)"
                            class="sw-table__row_can-select"
                            :class="{'sw-table__row_is-selected': buLayerFilterItem.Id === currentBuLayerFilterItem.item.Id}"
                        >
                            <td v-for="(value, key) in fields" :key="key + '-fields'">
                                <VCheckbox v-if="isCheckbox(buLayerFilterItem[key])" :checked="buLayerFilterItem[key]" disabled></VCheckbox>
                                <VText v-else>{{ buLayerFilterItem[key] }}</VText>
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
                                        :for="`currentBuLayerFilterItem${key}`"
                                    >{{ filed }}</VLabel>
                                    <VInput
										v-if="isInput(currentBuLayerFilterItem.item[key])"
                                        v-model="currentBuLayerFilterItem.item[key]"
                                        width="dyn"
                                        :id="`currentBuLayerFilterItem${key}`"
                                        @input="changeCurrentBuLayerFilterItem"
                                    />
									<VCheckbox
										v-if="isCheckbox(currentBuLayerFilterItem.item[key])"
                                        v-model="currentBuLayerFilterItem.item[key]"
                                        :id="`currentBuLayerFilterItem${key}`"
										@input="changeCurrentBuLayerFilterItem"
									/>
									
                                </VSet>
                            </VSet>
                            <button type="submit" :disabled="!currentBuLayerFilterItem.hasChange" hidden></button>
                        </form>
                    </template>

                    <template #footer>
                        <VSet>
                            <VButton
                                @click="saveChangesSubmit"
                                accent
                                :text="panelSubmitButtonText"
                                :disabled="!currentBuLayerFilterItem.hasChange"
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
                    v-if="currentBuLayerFilterItem.showDeleteConfirmation"
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
                        :disabled="!currentBuLayerFilterItem.isSelected"
                        @click="deleteBuLayerFilterItemHandler"
                    />
                </VSet>
            </slot>
        </template>
    </WorkSpace>
</template>

<script>
    import buLayerFilterData from "../data/BuLayerFilterData";
    import { BuLayerFilter } from '../apiModel';
    import { mapGetters, mapMutations, mapActions } from 'vuex';
    import WorkSpace from "swtui/src/components/WorkSpace";
    import VHead from "swtui/src/components/VHead";
    import VSet from "swtui/src/components/VSet";
    import VLabel from "swtui/src/components/VLabel";
    import VInput from "swtui/src/components/VInput";
    import VCheckbox from "swtui/src/components/VCheckbox";
    import VText from "swtui/src/components/VText";
    import VPanel from "swtui/src/components/VPanel";
    import VButton from "swtui/src/components/VButton";

    export default {
        name: 'BuLayerFilterGen',

        components: {VButton, VPanel, VText, VInput, VLabel, VSet, VHead, WorkSpace, VCheckbox},

        props: {
            fields: {
                type: Object,
                default() {
                    const buLayerFilterItem = new BuLayerFilter();
                    const fieldsObj = {};

                    for (let prop in buLayerFilterItem) {

                        if (buLayerFilterItem.hasOwnProperty(prop)) {
                            fieldsObj[prop] = prop;
                        }

                    }

                    return fieldsObj;
                }
            },
            editFields: {
                type: Object,
                default() {
                    const buLayerFilterItem = new BuLayerFilter();
                    const fieldsObj = {};

                    for (let prop in buLayerFilterItem) {

                        if (buLayerFilterItem.hasOwnProperty(prop)) {
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
            return buLayerFilterData;
        },

        created() {
			this.onCreated();
        },

        computed: {
            ...mapGetters('gosha', {
                buLayerFilterList: 'getListBuLayerFilter'
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
                'findBuLayerFilter',
                'updateBuLayerFilter',
                'deleteBuLayerFilter',
                'createBuLayerFilter',
            ]),

            ...mapMutations('gosha', [
                'addBuLayerFilterItemToList',
                'deleteBuLayerFilterFromList',
                'updateBuLayerFilterById',
            ]),

			onCreated() {
				this.fillBuLayerFilterFilter();
	            this.fetchBuLayerFilterData();
			},

            fillBuLayerFilterFilter() {
                this.buLayerFilterFilter.CurrentPage = 1;
                this.buLayerFilterFilter.PerPage = 1000;
            },

            fetchBuLayerFilterData() {
                return this.findBuLayerFilter({
                    filter: this.buLayerFilterFilter
                });
            },

            /**
             *
             * @param type
             */
            showPanel(type) {
                if (type === this.panel.create) {
                    this.panel.type = this.panel.create;
                    this.clearPanelBuLayerFilterItem();
                } else if (type === this.panel.edit) {
                    this.panel.type = this.panel.edit;
                    this.currentBuLayerFilterItem.isSelected = true;
                }

                this.panel.show = true;
            },

            closePanel() {
                this.panel.show = false;
                this.currentBuLayerFilterItem.isSelected = false;
                this.clearPanelBuLayerFilterItem();
            },

            selectBuLayerFilterItem(buLayerFilterItem) {
                this.showPanel(this.panel.edit);
                this.currentBuLayerFilterItem.isSelected = true;
                Object.assign(this.currentBuLayerFilterItem.item, buLayerFilterItem);
            },

            changeCurrentBuLayerFilterItem() {
                this.currentBuLayerFilterItem.hasChange = true;
            },

            cancelChanges() {
                this.clearPanelBuLayerFilterItem();
                this.closePanel();
            },

            clearPanelBuLayerFilterItem() {
                this.currentBuLayerFilterItem.item = new BuLayerFilter();
                this.currentBuLayerFilterItem.hasChange = false;
            },

            saveChangesSubmit() {
                if (this.isPanelCreate) {
                    this.createBuLayerFilterItemSubmit();
                    return;
                }

                if (this.isPanelEdit) {
                    this.editBuLayerFilterItemSubmit();
                }
            },

            createBuLayerFilterItemSubmit() {
                return this.createBuLayerFilter({
					data: this.currentBuLayerFilterItem.item,
                }).then((response) => {

                    if (response.Model) {
                        this.addBuLayerFilterItemToList(response.Model);
                        this.clearPanelBuLayerFilterItem();
                    } else {
                        console.error('Ошибка создания записи: ', response.Error);
                    }

                }).catch(error => {
                    console.error('Ошибка создания записи: ', error);
                });
            },

            editBuLayerFilterItemSubmit() {

                if (this.currentBuLayerFilterItem.hasChange) {
                    return this.updateBuLayerFilter({
                        id: this.currentBuLayerFilterItem.item.Id,
                        data: this.currentBuLayerFilterItem.item,
                    }).then((response) => {

                        if (response.Model) {
                            this.updateBuLayerFilterById(response.Model);
                            this.currentBuLayerFilterItem.hasChange = false;
                            this.clearPanelBuLayerFilterItem();
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

            deleteBuLayerFilterItemHandler() {
                let deletedItemId = this.currentBuLayerFilterItem.item.Id;

                if (!this.currentBuLayerFilterItem.canDelete) {
                    this.currentBuLayerFilterItem.showDeleteConfirmation = true;
                    return;
                }

                this.deleteBuLayerFilter({
                    id: deletedItemId
                }).then(response => {

                    if (response.IsSuccess) {
                        this.deleteBuLayerFilterFromList(deletedItemId);
                        this.clearPanelBuLayerFilterItem();
                        this.currentBuLayerFilterItem.canDelete = false;
                        this.currentBuLayerFilterItem.isSelected = false;
                        this.panel.show = false;
                    } else {
                        console.error('Ошибка удаления элемента: ', response.Error);
                    }

                }).catch(error => {
                    console.error('Ошибка удаления элемента: ', error);
                });
            },

            confirmDeleteHandler() {
                this.currentBuLayerFilterItem.showDeleteConfirmation = false;
                this.currentBuLayerFilterItem.canDelete = true;
                this.deleteBuLayerFilterItemHandler();
            },

            closeConfirmationPanel() {
                this.currentBuLayerFilterItem.showDeleteConfirmation = false;
            },
        },
    }
</script>

<style lang="scss">

</style>
