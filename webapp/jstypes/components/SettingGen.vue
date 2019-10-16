
	<template>
    <WorkSpace>
        <template #header>
            <slot name="pageHeader">
                <VHead level="h1">Setting</VHead>
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
                            v-for="settingItem in settingList"
                            :key="settingItem.Id"
                            @click="selectSettingItem(settingItem)"
                            class="sw-table__row_can-select"
                            :class="{'sw-table__row_is-selected': settingItem.Id === currentSettingItem.item.Id}"
                        >
                            <td v-for="(value, key) in fields" :key="key + '-fields'">
                                <VCheckbox v-if="isCheckbox(settingItem[key])" :checked="settingItem[key]" disabled></VCheckbox>
                                <VText v-else>{{ settingItem[key] }}</VText>
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
                                        :for="`currentSettingItem${key}`"
                                    >{{ filed }}</VLabel>
                                    <VInput
										v-if="isInput(currentSettingItem.item[key])"
                                        v-model="currentSettingItem.item[key]"
                                        width="dyn"
                                        :id="`currentSettingItem${key}`"
                                        @input="changeCurrentSettingItem"
                                    />
									<VCheckbox
										v-if="isCheckbox(currentSettingItem.item[key])"
                                        v-model="currentSettingItem.item[key]"
                                        :id="`currentSettingItem${key}`"
										@input="changeCurrentSettingItem"
									/>
									
                                </VSet>
                            </VSet>
                            <button type="submit" :disabled="!currentSettingItem.hasChange" hidden></button>
                        </form>
                    </template>

                    <template #footer>
                        <VSet>
                            <VButton
                                @click="saveChangesSubmit"
                                accent
                                :text="panelSubmitButtonText"
                                :disabled="!currentSettingItem.hasChange"
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
                    v-if="currentSettingItem.showDeleteConfirmation"
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
                        :disabled="!currentSettingItem.isSelected"
                        @click="deleteSettingItemHandler"
                    />
                </VSet>
            </slot>
        </template>
    </WorkSpace>
</template>

<script>
    import settingData from "../data/SettingData";
    import { Setting } from '../apiModel';
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
        name: 'SettingGen',

        components: {VSelectSimple, VSign, VIcon, VButton, VPanel, VText, VInput, VLabel, VSet, VHead, WorkSpace, VCheckbox},

        props: {
            fields: {
                type: Object,
                default() {
                    const settingItem = new Setting();
                    const fieldsObj = {};

                    for (let prop in settingItem) {

                        if (settingItem.hasOwnProperty(prop)) {
                            fieldsObj[prop] = prop;
                        }

                    }

                    return fieldsObj;
                }
            },
            editFields: {
                type: Object,
                default() {
                    const settingItem = new Setting();
                    const fieldsObj = {};

                    for (let prop in settingItem) {

                        if (settingItem.hasOwnProperty(prop)) {
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
            return settingData;
        },

        created() {
            this.fillSettingFilter();
            this.fetchSettingData();
        },

        computed: {
            ...mapGetters({
                settingList: 'getListSetting'
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
                'findSetting',
                'updateSetting',
                'deleteSetting',
                'createSetting',
            ]),

            ...mapMutations([
                'addSettingItemToList',
                'deleteSettingFromList',
                'updateSettingById',
            ]),

            fillSettingFilter() {
                this.settingFilter.CurrentPage = 1;
                this.settingFilter.PerPage = 1000;
            },

            fetchSettingData() {
                return this.findSetting({
                    filter: this.settingFilter
                });
            },

            /**
             *
             * @param type
             */
            showPanel(type) {
                if (type === this.panel.create) {
                    this.panel.type = this.panel.create;
                    this.clearPanelSettingItem();
                } else if (type === this.panel.edit) {
                    this.panel.type = this.panel.edit;
                    this.currentSettingItem.isSelected = true;
                }

                this.panel.show = true;
            },

            closePanel() {
                this.panel.show = false;
                this.currentSettingItem.isSelected = false;
                this.clearPanelSettingItem();
            },

            selectSettingItem(settingItem) {
                this.showPanel(this.panel.edit);
                this.currentSettingItem.isSelected = true;
                Object.assign(this.currentSettingItem.item, settingItem);
            },

            changeCurrentSettingItem() {
                this.currentSettingItem.hasChange = true;
            },

            cancelChanges() {
                this.clearPanelSettingItem();
                this.closePanel();
            },

            clearPanelSettingItem() {
                this.currentSettingItem.item = new Setting();
                this.currentSettingItem.hasChange = false;
            },

            saveChangesSubmit() {
                if (this.isPanelCreate) {
                    this.createSettingItemSubmit();
                    return;
                }

                if (this.isPanelEdit) {
                    this.editSettingItemSubmit();
                }
            },

            createSettingItemSubmit() {
                return this.createSetting({
					data: this.currentSettingItem.item,
                }).then((response) => {

                    if (response.Model) {
                        this.addSettingItemToList(response.Model);
                        this.clearPanelSettingItem();
                    } else {
                        console.error('Ошибка создания записи: ', response.Error);
                    }

                }).catch(error => {
                    console.error('Ошибка создания записи: ', error);
                });
            },

            editSettingItemSubmit() {

                if (this.currentSettingItem.hasChange) {
                    return this.updateSetting({
                        id: this.currentSettingItem.item.Id,
                        data: this.currentSettingItem.item,
                    }).then((response) => {

                        if (response.Model) {
                            this.updateSettingById(response.Model);
                            this.currentSettingItem.hasChange = false;
                            this.clearPanelSettingItem();
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

            deleteSettingItemHandler() {
                let deletedItemId = this.currentSettingItem.item.Id;

                if (!this.currentSettingItem.canDelete) {
                    this.currentSettingItem.showDeleteConfirmation = true;
                    return;
                }

                this.deleteSetting({
                    id: deletedItemId
                }).then(response => {

                    if (response.IsSuccess) {
                        this.deleteSettingFromList(deletedItemId);
                        this.clearPanelSettingItem();
                        this.currentSettingItem.canDelete = false;
                        this.currentSettingItem.isSelected = false;
                        this.panel.show = false;
                    } else {
                        console.error('Ошибка удаления элемента: ', response.Error);
                    }

                }).catch(error => {
                    console.error('Ошибка удаления элемента: ', error);
                });
            },

            confirmDeleteHandler() {
                this.currentSettingItem.showDeleteConfirmation = false;
                this.currentSettingItem.canDelete = true;
                this.deleteSettingItemHandler();
            },

            closeConfirmationPanel() {
                this.currentSettingItem.showDeleteConfirmation = false;
            },
        },
    }
</script>

<style lang="scss">

</style>
