package cmd

const usualEntityVueComponent = `
	<template>
    <WorkSpace>
        <template #header>
            <slot name="pageHeader">
                <VHead level="h1">{Entity}</VHead>
            </slot>
        </template>

        <template #content>
            <slot name="data">
                <table>
                    <thead>
                        <tr>
                            <th v-for="header in fields">{{ header }}</th>
                        </tr>
                    </thead>
            
                    <tbody>
                        <tr
                            v-for="setting in {entity}List"
                            :key="setting.Id"
                            @click="selectSetting(setting)"
                            class="sw-table__row_can-select"
                            :class="{'sw-table__row_is-selected': setting.Id === {entity}Data.currentSetting.item.Id}"
                        >
                            <td v-for="(value, key) in fields">
                                <VText>{{ setting[key] }}</VText>
                            </td>
                        </tr>
                    </tbody>
                </table>
            </slot>
            
            <slot name="panel">
                <VPanel
                    v-if="{entity}Data.panel.show"
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
                                    v-for="(filed, key) in fields"
                                    vertical-align="center"
                                >
                                    <VLabel
                                        width="col4"
                                        :for="` + "`" + `currentSetting${key}` + "`" + `"
                                    >{{ filed }}</VLabel>
                                    <VInput
                                        v-model="{entity}Data.currentSetting.item[key]"
                                        width="dyn"
                                        :id="` + "`" + `currentSetting${key}` + "`" + `"
                                        @input="changeCurrentSetting"
                                    />
                                </VSet>
                            </VSet>
                            <button type="submit" :disabled="!{entity}Data.currentSetting.hasChange" hidden></button>
                        </form>
                    </template>

                    <template #footer>
                        <VSet>
                            <VButton
                                @click="saveChangesSubmit"
                                accent
                                :text="panelSubmitButtonText"
                                :disabled="!{entity}Data.currentSetting.hasChange"
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
                    v-if="{entity}Data.currentSetting.showDeleteConfirmation"
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
                        text="Добавить"
                        accent
                        @click="showPanel({entity}Data.panel.create)"
                    />
                    <VButton
                        text="Удалить"
                        :disabled="!{entity}Data.currentSetting.isSelected"
                        @click="deleteSettingHandler"
                    />
                </VSet>
            </slot>
        </template>
    </WorkSpace>
</template>

<script>
    import {entity}Data from "../data/{entity}Data";
    import { {Entity} } from '../apiModel';
    import { mapGetters, mapMutations, mapActions } from 'vuex';
    import WorkSpace from "swui/src/components/WorkSpace";
    import VHead from "swui/src/components/VHead";
    import VSet from "swui/src/components/VSet";
    import VLabel from "swui/src/components/VLabel";
    import VInput from "swui/src/components/VInput";
    import VText from "swui/src/components/VText";
    import VPanel from "swui/src/components/VPanel";
    import VButton from "swui/src/components/VButton";
    import VIcon from "swui/src/components/VIcon";
    import VSign from "swui/src/components/VSign";
    import VSelect from "swui/src/components/VSelect";

    export default {
        name: '{Entity}Gen',

        components: {VSelect, VSign, VIcon, VButton, VPanel, VText, VInput, VLabel, VSet, VHead, WorkSpace},

        props: {
            fields: {
                type: Object,
                default() {
                    const setting = new {Entity}();
                    const fieldsObj = {};

                    for (let prop in setting) {

                        if (setting.hasOwnProperty(prop)) {
                            fieldsObj[prop] = prop;
                        }

                    }

                    return fieldsObj;
                }
            }
        },

        data() {
            return {
                {entity}Data: {}
            }
        },

        created() {
            this.{entity}Data = {entity}Data;
            this.fill{Entity}Filter();
            this.fetch{Entity}Data();
        },

        computed: {
            ...mapGetters({
                {entity}List: 'getList{Entity}'
            }),
            isPanelCreate() {
                return {entity}Data.panel.type === {entity}Data.panel.create;
            },
            isPanelEdit() {
                return {entity}Data.panel.type === {entity}Data.panel.edit;
            },
            panelHeader() {
                if (this.isPanelCreate) {
                    return {entity}Data.panelHeaderCreate;
                }

                if (this.isPanelEdit) {
                    return {entity}Data.panelHeaderEdit;
                }

                return  '';
            },
            panelSubmitButtonText() {
                if (this.isPanelCreate) {
                    return {entity}Data.panelSubmitButtonTextCreate;
                }

                if (this.isPanelEdit) {
                    return {entity}Data.panelSubmitButtonTextEdit;
                }

                return  '';
            }
        },

        methods: {
            ...mapActions([
                'find{Entity}',
                'update{Entity}',
                'delete{Entity}',
                'create{Entity}',
            ]),

            ...mapMutations([
                'addSettingToList',
                'delete{Entity}FromList',
                'update{Entity}ById',
            ]),

            fill{Entity}Filter() {
                this.{entity}Data.{entity}Filter.CurrentPage = 1;
                this.{entity}Data.{entity}Filter.PerPage = 1000;
            },

            fetch{Entity}Data() {
                return this.find{Entity}({
                    filter: this.{entity}Data.{entity}Filter
                });
            },

            /**
             *
             * @param type
             */
            showPanel(type) {
                if (type === {entity}Data.panel.create) {
                    {entity}Data.panel.type = {entity}Data.panel.create;
                    this.clearPanelSetting();
                } else if (type === {entity}Data.panel.edit) {
                    {entity}Data.panel.type = {entity}Data.panel.edit;
                    {entity}Data.currentSetting.isSelected = true;
                }

                {entity}Data.panel.show = true;
            },

            closePanel() {
                this.{entity}Data.panel.show = false;
                this.{entity}Data.currentSetting.isSelected = false;
                this.clearPanelSetting();
            },

            selectSetting(setting) {
                this.showPanel({entity}Data.panel.edit);
                this.{entity}Data.currentSetting.isSelected = true;
                Object.assign(this.{entity}Data.currentSetting.item, setting);
            },

            changeCurrentSetting() {
                this.{entity}Data.currentSetting.hasChange = true;
            },

            cancelChanges() {
                this.clearPanelSetting();
                this.closePanel();
            },

            clearPanelSetting() {
                this.{entity}Data.currentSetting.item = new {Entity}();
                this.{entity}Data.currentSetting.hasChange = false;
            },

            saveChangesSubmit() {
                if (this.isPanelCreate) {
                    this.createSettingSubmit();
                    return;
                }

                if (this.isPanelEdit) {
                    this.editSettingSubmit();
                }
            },

            createSettingSubmit() {
                this.create{Entity}({
                    data: {
                        Name: this.{entity}Data.currentSetting.item.Name,
                        Value: this.{entity}Data.currentSetting.item.Value,
                        Description: this.{entity}Data.currentSetting.item.Description,
                    }
                }).then((response) => {

                    if (response.Model) {
                        this.addSettingToList(response.Model);
                        this.clearPanelSetting();
                    } else {
                        console.error('Ошибка создания записи: ', response.Error);
                    }

                }).catch(error => {
                    console.error('Ошибка создания записи: ', error);
                });
            },

            editSettingSubmit() {
                if (this.{entity}Data.currentSetting.hasChange) {
                    this.update{Entity}({
                        id: this.{entity}Data.currentSetting.item.Id,
                        data: this.{entity}Data.currentSetting.item,
                    }).then((response) => {

                        if (response.Model) {
                            this.update{Entity}ById(response.Model);
                            this.{entity}Data.currentSetting.hasChange = false;
                            this.clearPanelSetting();
                            this.closePanel();
                        } else {
                            console.error('Ошибка изменения записи: ', response.Error);
                        }

                    }).catch(error => {
                        console.error('Ошибка изменения записи: ', error);
                    });
                }
            },

            deleteSettingHandler() {
                let deletedItemId = this.{entity}Data.currentSetting.item.Id;

                if (!this.{entity}Data.currentSetting.canDelete) {
                    this.{entity}Data.currentSetting.showDeleteConfirmation = true;
                    return;
                }

                this.delete{Entity}({
                    id: deletedItemId
                }).then(response => {

                    if (response.IsSuccess) {
                        this.delete{Entity}FromList(deletedItemId);
                        this.clearPanelSetting();
                        this.{entity}Data.currentSetting.canDelete = false;
                        this.{entity}Data.currentSetting.isSelected = false;
                        this.{entity}Data.panel.show = false;
                    } else {
                        console.error('Ошибка удаления элемента: ', response.Error);
                    }

                }).catch(error => {
                    console.error('Ошибка удаления элемента: ', error);
                });
            },

            confirmDeleteHandler() {
                this.{entity}Data.currentSetting.showDeleteConfirmation = false;
                this.{entity}Data.currentSetting.canDelete = true;
                this.deleteSettingHandler();
            },

            closeConfirmationPanel() {
                {entity}Data.currentSetting.showDeleteConfirmation = false;
            },
        },
    }
</script>

<style lang="scss"></style>
`