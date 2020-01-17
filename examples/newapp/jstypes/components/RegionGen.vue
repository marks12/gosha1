
	<template>
    <WorkSpace>
        <template #header>
            <slot name="pageHeader">
                <VHead level="h1">Region</VHead>
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
                            v-for="regionItem in regionList"
                            :key="regionItem.Id"
                            @click="selectRegionItem(regionItem)"
                            class="sw-table__row_can-select"
                            :class="{'sw-table__row_is-selected': regionItem.Id === currentRegionItem.item.Id}"
                        >
                            <td v-for="(value, key) in fields" :key="key + '-fields'">
                                <VCheckbox v-if="isCheckbox(regionItem[key])" :checked="regionItem[key]" disabled></VCheckbox>
                                <VText v-else>{{ regionItem[key] }}</VText>
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
                                        :for="`currentRegionItem${key}`"
                                    >{{ filed }}</VLabel>
                                    <VInput
										v-if="isInput(currentRegionItem.item[key])"
                                        v-model="currentRegionItem.item[key]"
                                        width="dyn"
                                        :id="`currentRegionItem${key}`"
                                        @input="changeCurrentRegionItem"
                                    />
									<VCheckbox
										v-if="isCheckbox(currentRegionItem.item[key])"
                                        v-model="currentRegionItem.item[key]"
                                        :id="`currentRegionItem${key}`"
										@input="changeCurrentRegionItem"
									/>
									
                                </VSet>
                            </VSet>
                            <button type="submit" :disabled="!currentRegionItem.hasChange" hidden></button>
                        </form>
                    </template>

                    <template #footer>
                        <VSet>
                            <VButton
                                @click="saveChangesSubmit"
                                accent
                                :text="panelSubmitButtonText"
                                :disabled="!currentRegionItem.hasChange"
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
                    v-if="currentRegionItem.showDeleteConfirmation"
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
                        :disabled="!currentRegionItem.isSelected"
                        @click="deleteRegionItemHandler"
                    />
                </VSet>
            </slot>
        </template>
    </WorkSpace>
</template>

<script>
    import regionData from "../data/RegionData";
    import { Region } from '../apiModel';
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
        name: 'RegionGen',

        components: {VButton, VPanel, VText, VInput, VLabel, VSet, VHead, WorkSpace, VCheckbox},

        props: {
            fields: {
                type: Object,
                default() {
                    const regionItem = new Region();
                    const fieldsObj = {};

                    for (let prop in regionItem) {

                        if (regionItem.hasOwnProperty(prop)) {
                            fieldsObj[prop] = prop;
                        }

                    }

                    return fieldsObj;
                }
            },
            editFields: {
                type: Object,
                default() {
                    const regionItem = new Region();
                    const fieldsObj = {};

                    for (let prop in regionItem) {

                        if (regionItem.hasOwnProperty(prop)) {
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
            return regionData;
        },

        created() {
			this.onCreated();
        },

        computed: {
            ...mapGetters({
                regionList: 'getListRegion'
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
                'findRegion',
                'updateRegion',
                'deleteRegion',
                'createRegion',
            ]),

            ...mapMutations([
                'addRegionItemToList',
                'deleteRegionFromList',
                'updateRegionById',
            ]),

			onCreated() {
				this.fillRegionFilter();
	            this.fetchRegionData();
			},

            fillRegionFilter() {
                this.regionFilter.CurrentPage = 1;
                this.regionFilter.PerPage = 1000;
            },

            fetchRegionData() {
                return this.findRegion({
                    filter: this.regionFilter
                });
            },

            /**
             *
             * @param type
             */
            showPanel(type) {
                if (type === this.panel.create) {
                    this.panel.type = this.panel.create;
                    this.clearPanelRegionItem();
                } else if (type === this.panel.edit) {
                    this.panel.type = this.panel.edit;
                    this.currentRegionItem.isSelected = true;
                }

                this.panel.show = true;
            },

            closePanel() {
                this.panel.show = false;
                this.currentRegionItem.isSelected = false;
                this.clearPanelRegionItem();
            },

            selectRegionItem(regionItem) {
                this.showPanel(this.panel.edit);
                this.currentRegionItem.isSelected = true;
                Object.assign(this.currentRegionItem.item, regionItem);
            },

            changeCurrentRegionItem() {
                this.currentRegionItem.hasChange = true;
            },

            cancelChanges() {
                this.clearPanelRegionItem();
                this.closePanel();
            },

            clearPanelRegionItem() {
                this.currentRegionItem.item = new Region();
                this.currentRegionItem.hasChange = false;
            },

            saveChangesSubmit() {
                if (this.isPanelCreate) {
                    this.createRegionItemSubmit();
                    return;
                }

                if (this.isPanelEdit) {
                    this.editRegionItemSubmit();
                }
            },

            createRegionItemSubmit() {
                return this.createRegion({
					data: this.currentRegionItem.item,
                }).then((response) => {

                    if (response.Model) {
                        this.addRegionItemToList(response.Model);
                        this.clearPanelRegionItem();
                    } else {
                        console.error('Ошибка создания записи: ', response.Error);
                    }

                }).catch(error => {
                    console.error('Ошибка создания записи: ', error);
                });
            },

            editRegionItemSubmit() {

                if (this.currentRegionItem.hasChange) {
                    return this.updateRegion({
                        id: this.currentRegionItem.item.Id,
                        data: this.currentRegionItem.item,
                    }).then((response) => {

                        if (response.Model) {
                            this.updateRegionById(response.Model);
                            this.currentRegionItem.hasChange = false;
                            this.clearPanelRegionItem();
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

            deleteRegionItemHandler() {
                let deletedItemId = this.currentRegionItem.item.Id;

                if (!this.currentRegionItem.canDelete) {
                    this.currentRegionItem.showDeleteConfirmation = true;
                    return;
                }

                this.deleteRegion({
                    id: deletedItemId
                }).then(response => {

                    if (response.IsSuccess) {
                        this.deleteRegionFromList(deletedItemId);
                        this.clearPanelRegionItem();
                        this.currentRegionItem.canDelete = false;
                        this.currentRegionItem.isSelected = false;
                        this.panel.show = false;
                    } else {
                        console.error('Ошибка удаления элемента: ', response.Error);
                    }

                }).catch(error => {
                    console.error('Ошибка удаления элемента: ', error);
                });
            },

            confirmDeleteHandler() {
                this.currentRegionItem.showDeleteConfirmation = false;
                this.currentRegionItem.canDelete = true;
                this.deleteRegionItemHandler();
            },

            closeConfirmationPanel() {
                this.currentRegionItem.showDeleteConfirmation = false;
            },
        },
    }
</script>

<style lang="scss">

</style>
