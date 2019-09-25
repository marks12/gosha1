
	<template>
    <WorkSpace>
        <template #header>
            <slot name="pageHeader">
                <VHead level="h1">SettingFilter</VHead>
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
                            v-for="settingFilterItem in settingFilterList"
                            :key="settingFilterItem.Id"
                            @click="selectSettingFilterItem(settingFilterItem)"
                            class="sw-table__row_can-select"
                            :class="{'sw-table__row_is-selected': settingFilterItem.Id === currentSettingFilterItem.item.Id}"
                        >
                            <td v-for="(value, key) in fields" :key="key + '-fields'">
                                <VCheckbox v-if="isCheckbox(settingFilterItem[key])" :checked="settingFilterItem[key]" disabled></VCheckbox>
                                <VText v-else>{{ settingFilterItem[key] }}</VText>
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
                                        :for="`currentSettingFilterItem${key}`"
                                    >{{ filed }}</VLabel>
                                    <VInput
										v-if="isInput(currentSettingFilterItem.item[key])"
                                        v-model="currentSettingFilterItem.item[key]"
                                        width="dyn"
                                        :id="`currentSettingFilterItem${key}`"
                                        @input="changeCurrentSettingFilterItem"
                                    />
									<VCheckbox
										v-if="isCheckbox(currentSettingFilterItem.item[key])"
                                        v-model="currentSettingFilterItem.item[key]"
                                        :id="`currentSettingFilterItem${key}`"
										@input="changeCurrentSettingFilterItem"
									/>
									
                                </VSet>
                            </VSet>
                            <button type="submit" :disabled="!currentSettingFilterItem.hasChange" hidden></button>
                        </form>
                    </template>

                    <template #footer>
                        <VSet>
                            <VButton
                                @click="saveChangesSubmit"
                                accent
                                :text="panelSubmitButtonText"
                                :disabled="!currentSettingFilterItem.hasChange"
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
                    v-if="currentSettingFilterItem.showDeleteConfirmation"
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
                        :disabled="!currentSettingFilterItem.isSelected"
                        @click="deleteSettingFilterItemHandler"
                    />
                </VSet>
            </slot>
        </template>
    </WorkSpace>
</template>

<script>
    import settingFilterData from "../data/SettingFilterData";
    import { SettingFilter } from '../apiModel';
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
        name: 'SettingFilterGen',

        components: {VSelectSimple, VSign, VIcon, VButton, VPanel, VText, VInput, VLabel, VSet, VHead, WorkSpace, VCheckbox},

        props: {
            fields: {
                type: Object,
                default() {
                    const settingFilterItem = new SettingFilter();
                    const fieldsObj = {};

                    for (let prop in settingFilterItem) {

                        if (settingFilterItem.hasOwnProperty(prop)) {
                            fieldsObj[prop] = prop;
                        }

                    }

                    return fieldsObj;
                }
            },
            editFields: {
                type: Object,
                default() {
                    const settingFilterItem = new SettingFilter();
                    const fieldsObj = {};

                    for (let prop in settingFilterItem) {

                        if (settingFilterItem.hasOwnProperty(prop)) {
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
            return settingFilterData;
        },

        created() {
            this.fillSettingFilterFilter();
            this.fetchSettingFilterData();
        },

        computed: {
            ...mapGetters({
                settingFilterList: 'getListSettingFilter'
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
                'findSettingFilter',
                'updateSettingFilter',
                'deleteSettingFilter',
                'createSettingFilter',
            ]),

            ...mapMutations([
                'addSettingFilterItemToList',
                'deleteSettingFilterFromList',
                'updateSettingFilterById',
            ]),

            fillSettingFilterFilter() {
                this.settingFilterFilter.CurrentPage = 1;
                this.settingFilterFilter.PerPage = 1000;
            },

            fetchSettingFilterData() {
                return this.findSettingFilter({
                    filter: this.settingFilterFilter
                });
            },

            /**
             *
             * @param type
             */
            showPanel(type) {
                if (type === this.panel.create) {
                    this.panel.type = this.panel.create;
                    this.clearPanelSettingFilterItem();
                } else if (type === this.panel.edit) {
                    this.panel.type = this.panel.edit;
                    this.currentSettingFilterItem.isSelected = true;
                }

                this.panel.show = true;
            },

            closePanel() {
                this.panel.show = false;
                this.currentSettingFilterItem.isSelected = false;
                this.clearPanelSettingFilterItem();
            },

            selectSettingFilterItem(settingFilterItem) {
                this.showPanel(this.panel.edit);
                this.currentSettingFilterItem.isSelected = true;
                Object.assign(this.currentSettingFilterItem.item, settingFilterItem);
            },

            changeCurrentSettingFilterItem() {
                this.currentSettingFilterItem.hasChange = true;
            },

            cancelChanges() {
                this.clearPanelSettingFilterItem();
                this.closePanel();
            },

            clearPanelSettingFilterItem() {
                this.currentSettingFilterItem.item = new SettingFilter();
                this.currentSettingFilterItem.hasChange = false;
            },

            saveChangesSubmit() {
                if (this.isPanelCreate) {
                    this.createSettingFilterItemSubmit();
                    return;
                }

                if (this.isPanelEdit) {
                    this.editSettingFilterItemSubmit();
                }
            },

            createSettingFilterItemSubmit() {
                this.createSettingFilter({
					data: this.currentSettingFilterItem.item,
                }).then((response) => {

                    if (response.Model) {
                        this.addSettingFilterItemToList(response.Model);
                        this.clearPanelSettingFilterItem();
                    } else {
                        console.error('Ошибка создания записи: ', response.Error);
                    }

                }).catch(error => {
                    console.error('Ошибка создания записи: ', error);
                });
            },

            editSettingFilterItemSubmit() {
                if (this.currentSettingFilterItem.hasChange) {
                    this.updateSettingFilter({
                        id: this.currentSettingFilterItem.item.Id,
                        data: this.currentSettingFilterItem.item,
                    }).then((response) => {

                        if (response.Model) {
                            this.updateSettingFilterById(response.Model);
                            this.currentSettingFilterItem.hasChange = false;
                            this.clearPanelSettingFilterItem();
                            this.closePanel();
                        } else {
                            console.error('Ошибка изменения записи: ', response.Error);
                        }

                    }).catch(error => {
                        console.error('Ошибка изменения записи: ', error);
                    });
                }
            },

            deleteSettingFilterItemHandler() {
                let deletedItemId = this.currentSettingFilterItem.item.Id;

                if (!this.currentSettingFilterItem.canDelete) {
                    this.currentSettingFilterItem.showDeleteConfirmation = true;
                    return;
                }

                this.deleteSettingFilter({
                    id: deletedItemId
                }).then(response => {

                    if (response.IsSuccess) {
                        this.deleteSettingFilterFromList(deletedItemId);
                        this.clearPanelSettingFilterItem();
                        this.currentSettingFilterItem.canDelete = false;
                        this.currentSettingFilterItem.isSelected = false;
                        this.panel.show = false;
                    } else {
                        console.error('Ошибка удаления элемента: ', response.Error);
                    }

                }).catch(error => {
                    console.error('Ошибка удаления элемента: ', error);
                });
            },

            confirmDeleteHandler() {
                this.currentSettingFilterItem.showDeleteConfirmation = false;
                this.currentSettingFilterItem.canDelete = true;
                this.deleteSettingFilterItemHandler();
            },

            closeConfirmationPanel() {
                this.currentSettingFilterItem.showDeleteConfirmation = false;
            },
        },
    }
</script>

<style lang="scss">

</style>
