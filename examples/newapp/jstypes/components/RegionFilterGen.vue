
	<template>
    <WorkSpace>
        <template #header>
            <slot name="pageHeader">
                <VHead level="h1">RegionFilter</VHead>
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
                            v-for="regionFilterItem in regionFilterList"
                            :key="regionFilterItem.Id"
                            @click="selectRegionFilterItem(regionFilterItem)"
                            class="sw-table__row_can-select"
                            :class="{'sw-table__row_is-selected': regionFilterItem.Id === currentRegionFilterItem.item.Id}"
                        >
                            <td v-for="(value, key) in fields" :key="key + '-fields'">
                                <VCheckbox v-if="isCheckbox(regionFilterItem[key])" :checked="regionFilterItem[key]" disabled></VCheckbox>
                                <VText v-else>{{ regionFilterItem[key] }}</VText>
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
                                        :for="`currentRegionFilterItem${key}`"
                                    >{{ filed }}</VLabel>
                                    <VInput
										v-if="isInput(currentRegionFilterItem.item[key])"
                                        v-model="currentRegionFilterItem.item[key]"
                                        width="dyn"
                                        :id="`currentRegionFilterItem${key}`"
                                        @input="changeCurrentRegionFilterItem"
                                    />
									<VCheckbox
										v-if="isCheckbox(currentRegionFilterItem.item[key])"
                                        v-model="currentRegionFilterItem.item[key]"
                                        :id="`currentRegionFilterItem${key}`"
										@input="changeCurrentRegionFilterItem"
									/>
									
                                </VSet>
                            </VSet>
                            <button type="submit" :disabled="!currentRegionFilterItem.hasChange" hidden></button>
                        </form>
                    </template>

                    <template #footer>
                        <VSet>
                            <VButton
                                @click="saveChangesSubmit"
                                accent
                                :text="panelSubmitButtonText"
                                :disabled="!currentRegionFilterItem.hasChange"
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
                    v-if="currentRegionFilterItem.showDeleteConfirmation"
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
                        :disabled="!currentRegionFilterItem.isSelected"
                        @click="deleteRegionFilterItemHandler"
                    />
                </VSet>
            </slot>
        </template>
    </WorkSpace>
</template>

<script>
    import regionFilterData from "../data/RegionFilterData";
    import { RegionFilter } from '../apiModel';
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
        name: 'RegionFilterGen',

        components: {VSelectSimple, VSign, VIcon, VButton, VPanel, VText, VInput, VLabel, VSet, VHead, WorkSpace, VCheckbox},

        props: {
            fields: {
                type: Object,
                default() {
                    const regionFilterItem = new RegionFilter();
                    const fieldsObj = {};

                    for (let prop in regionFilterItem) {

                        if (regionFilterItem.hasOwnProperty(prop)) {
                            fieldsObj[prop] = prop;
                        }

                    }

                    return fieldsObj;
                }
            },
            editFields: {
                type: Object,
                default() {
                    const regionFilterItem = new RegionFilter();
                    const fieldsObj = {};

                    for (let prop in regionFilterItem) {

                        if (regionFilterItem.hasOwnProperty(prop)) {
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
            return regionFilterData;
        },

        created() {
            this.fillRegionFilterFilter();
            this.fetchRegionFilterData();
        },

        computed: {
            ...mapGetters({
                regionFilterList: 'getListRegionFilter'
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
                'findRegionFilter',
                'updateRegionFilter',
                'deleteRegionFilter',
                'createRegionFilter',
            ]),

            ...mapMutations([
                'addRegionFilterItemToList',
                'deleteRegionFilterFromList',
                'updateRegionFilterById',
            ]),

            fillRegionFilterFilter() {
                this.regionFilterFilter.CurrentPage = 1;
                this.regionFilterFilter.PerPage = 1000;
            },

            fetchRegionFilterData() {
                return this.findRegionFilter({
                    filter: this.regionFilterFilter
                });
            },

            /**
             *
             * @param type
             */
            showPanel(type) {
                if (type === this.panel.create) {
                    this.panel.type = this.panel.create;
                    this.clearPanelRegionFilterItem();
                } else if (type === this.panel.edit) {
                    this.panel.type = this.panel.edit;
                    this.currentRegionFilterItem.isSelected = true;
                }

                this.panel.show = true;
            },

            closePanel() {
                this.panel.show = false;
                this.currentRegionFilterItem.isSelected = false;
                this.clearPanelRegionFilterItem();
            },

            selectRegionFilterItem(regionFilterItem) {
                this.showPanel(this.panel.edit);
                this.currentRegionFilterItem.isSelected = true;
                Object.assign(this.currentRegionFilterItem.item, regionFilterItem);
            },

            changeCurrentRegionFilterItem() {
                this.currentRegionFilterItem.hasChange = true;
            },

            cancelChanges() {
                this.clearPanelRegionFilterItem();
                this.closePanel();
            },

            clearPanelRegionFilterItem() {
                this.currentRegionFilterItem.item = new RegionFilter();
                this.currentRegionFilterItem.hasChange = false;
            },

            saveChangesSubmit() {
                if (this.isPanelCreate) {
                    this.createRegionFilterItemSubmit();
                    return;
                }

                if (this.isPanelEdit) {
                    this.editRegionFilterItemSubmit();
                }
            },

            createRegionFilterItemSubmit() {
                return this.createRegionFilter({
					data: this.currentRegionFilterItem.item,
                }).then((response) => {

                    if (response.Model) {
                        this.addRegionFilterItemToList(response.Model);
                        this.clearPanelRegionFilterItem();
                    } else {
                        console.error('Ошибка создания записи: ', response.Error);
                    }

                }).catch(error => {
                    console.error('Ошибка создания записи: ', error);
                });
            },

            editRegionFilterItemSubmit() {

                if (this.currentRegionFilterItem.hasChange) {
                    return this.updateRegionFilter({
                        id: this.currentRegionFilterItem.item.Id,
                        data: this.currentRegionFilterItem.item,
                    }).then((response) => {

                        if (response.Model) {
                            this.updateRegionFilterById(response.Model);
                            this.currentRegionFilterItem.hasChange = false;
                            this.clearPanelRegionFilterItem();
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

            deleteRegionFilterItemHandler() {
                let deletedItemId = this.currentRegionFilterItem.item.Id;

                if (!this.currentRegionFilterItem.canDelete) {
                    this.currentRegionFilterItem.showDeleteConfirmation = true;
                    return;
                }

                this.deleteRegionFilter({
                    id: deletedItemId
                }).then(response => {

                    if (response.IsSuccess) {
                        this.deleteRegionFilterFromList(deletedItemId);
                        this.clearPanelRegionFilterItem();
                        this.currentRegionFilterItem.canDelete = false;
                        this.currentRegionFilterItem.isSelected = false;
                        this.panel.show = false;
                    } else {
                        console.error('Ошибка удаления элемента: ', response.Error);
                    }

                }).catch(error => {
                    console.error('Ошибка удаления элемента: ', error);
                });
            },

            confirmDeleteHandler() {
                this.currentRegionFilterItem.showDeleteConfirmation = false;
                this.currentRegionFilterItem.canDelete = true;
                this.deleteRegionFilterItemHandler();
            },

            closeConfirmationPanel() {
                this.currentRegionFilterItem.showDeleteConfirmation = false;
            },
        },
    }
</script>

<style lang="scss">

</style>
