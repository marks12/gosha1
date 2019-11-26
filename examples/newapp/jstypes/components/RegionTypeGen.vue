
	<template>
    <WorkSpace>
        <template #header>
            <slot name="pageHeader">
                <VHead level="h1">RegionType</VHead>
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
                            v-for="regionTypeItem in regionTypeList"
                            :key="regionTypeItem.Id"
                            @click="selectRegionTypeItem(regionTypeItem)"
                            class="sw-table__row_can-select"
                            :class="{'sw-table__row_is-selected': regionTypeItem.Id === currentRegionTypeItem.item.Id}"
                        >
                            <td v-for="(value, key) in fields" :key="key + '-fields'">
                                <VCheckbox v-if="isCheckbox(regionTypeItem[key])" :checked="regionTypeItem[key]" disabled></VCheckbox>
                                <VText v-else>{{ regionTypeItem[key] }}</VText>
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
                                        :for="`currentRegionTypeItem${key}`"
                                    >{{ filed }}</VLabel>
                                    <VInput
										v-if="isInput(currentRegionTypeItem.item[key])"
                                        v-model="currentRegionTypeItem.item[key]"
                                        width="dyn"
                                        :id="`currentRegionTypeItem${key}`"
                                        @input="changeCurrentRegionTypeItem"
                                    />
									<VCheckbox
										v-if="isCheckbox(currentRegionTypeItem.item[key])"
                                        v-model="currentRegionTypeItem.item[key]"
                                        :id="`currentRegionTypeItem${key}`"
										@input="changeCurrentRegionTypeItem"
									/>
									
                                </VSet>
                            </VSet>
                            <button type="submit" :disabled="!currentRegionTypeItem.hasChange" hidden></button>
                        </form>
                    </template>

                    <template #footer>
                        <VSet>
                            <VButton
                                @click="saveChangesSubmit"
                                accent
                                :text="panelSubmitButtonText"
                                :disabled="!currentRegionTypeItem.hasChange"
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
                    v-if="currentRegionTypeItem.showDeleteConfirmation"
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
                        :disabled="!currentRegionTypeItem.isSelected"
                        @click="deleteRegionTypeItemHandler"
                    />
                </VSet>
            </slot>
        </template>
    </WorkSpace>
</template>

<script>
    import regionTypeData from "../data/RegionTypeData";
    import { RegionType } from '../apiModel';
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
        name: 'RegionTypeGen',

        components: {VSelectSimple, VSign, VIcon, VButton, VPanel, VText, VInput, VLabel, VSet, VHead, WorkSpace, VCheckbox},

        props: {
            fields: {
                type: Object,
                default() {
                    const regionTypeItem = new RegionType();
                    const fieldsObj = {};

                    for (let prop in regionTypeItem) {

                        if (regionTypeItem.hasOwnProperty(prop)) {
                            fieldsObj[prop] = prop;
                        }

                    }

                    return fieldsObj;
                }
            },
            editFields: {
                type: Object,
                default() {
                    const regionTypeItem = new RegionType();
                    const fieldsObj = {};

                    for (let prop in regionTypeItem) {

                        if (regionTypeItem.hasOwnProperty(prop)) {
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
            return regionTypeData;
        },

        created() {
            this.fillRegionTypeFilter();
            this.fetchRegionTypeData();
        },

        computed: {
            ...mapGetters({
                regionTypeList: 'getListRegionType'
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
                'findRegionType',
                'updateRegionType',
                'deleteRegionType',
                'createRegionType',
            ]),

            ...mapMutations([
                'addRegionTypeItemToList',
                'deleteRegionTypeFromList',
                'updateRegionTypeById',
            ]),

            fillRegionTypeFilter() {
                this.regionTypeFilter.CurrentPage = 1;
                this.regionTypeFilter.PerPage = 1000;
            },

            fetchRegionTypeData() {
                return this.findRegionType({
                    filter: this.regionTypeFilter
                });
            },

            /**
             *
             * @param type
             */
            showPanel(type) {
                if (type === this.panel.create) {
                    this.panel.type = this.panel.create;
                    this.clearPanelRegionTypeItem();
                } else if (type === this.panel.edit) {
                    this.panel.type = this.panel.edit;
                    this.currentRegionTypeItem.isSelected = true;
                }

                this.panel.show = true;
            },

            closePanel() {
                this.panel.show = false;
                this.currentRegionTypeItem.isSelected = false;
                this.clearPanelRegionTypeItem();
            },

            selectRegionTypeItem(regionTypeItem) {
                this.showPanel(this.panel.edit);
                this.currentRegionTypeItem.isSelected = true;
                Object.assign(this.currentRegionTypeItem.item, regionTypeItem);
            },

            changeCurrentRegionTypeItem() {
                this.currentRegionTypeItem.hasChange = true;
            },

            cancelChanges() {
                this.clearPanelRegionTypeItem();
                this.closePanel();
            },

            clearPanelRegionTypeItem() {
                this.currentRegionTypeItem.item = new RegionType();
                this.currentRegionTypeItem.hasChange = false;
            },

            saveChangesSubmit() {
                if (this.isPanelCreate) {
                    this.createRegionTypeItemSubmit();
                    return;
                }

                if (this.isPanelEdit) {
                    this.editRegionTypeItemSubmit();
                }
            },

            createRegionTypeItemSubmit() {
                return this.createRegionType({
					data: this.currentRegionTypeItem.item,
                }).then((response) => {

                    if (response.Model) {
                        this.addRegionTypeItemToList(response.Model);
                        this.clearPanelRegionTypeItem();
                    } else {
                        console.error('Ошибка создания записи: ', response.Error);
                    }

                }).catch(error => {
                    console.error('Ошибка создания записи: ', error);
                });
            },

            editRegionTypeItemSubmit() {

                if (this.currentRegionTypeItem.hasChange) {
                    return this.updateRegionType({
                        id: this.currentRegionTypeItem.item.Id,
                        data: this.currentRegionTypeItem.item,
                    }).then((response) => {

                        if (response.Model) {
                            this.updateRegionTypeById(response.Model);
                            this.currentRegionTypeItem.hasChange = false;
                            this.clearPanelRegionTypeItem();
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

            deleteRegionTypeItemHandler() {
                let deletedItemId = this.currentRegionTypeItem.item.Id;

                if (!this.currentRegionTypeItem.canDelete) {
                    this.currentRegionTypeItem.showDeleteConfirmation = true;
                    return;
                }

                this.deleteRegionType({
                    id: deletedItemId
                }).then(response => {

                    if (response.IsSuccess) {
                        this.deleteRegionTypeFromList(deletedItemId);
                        this.clearPanelRegionTypeItem();
                        this.currentRegionTypeItem.canDelete = false;
                        this.currentRegionTypeItem.isSelected = false;
                        this.panel.show = false;
                    } else {
                        console.error('Ошибка удаления элемента: ', response.Error);
                    }

                }).catch(error => {
                    console.error('Ошибка удаления элемента: ', error);
                });
            },

            confirmDeleteHandler() {
                this.currentRegionTypeItem.showDeleteConfirmation = false;
                this.currentRegionTypeItem.canDelete = true;
                this.deleteRegionTypeItemHandler();
            },

            closeConfirmationPanel() {
                this.currentRegionTypeItem.showDeleteConfirmation = false;
            },
        },
    }
</script>

<style lang="scss">

</style>
