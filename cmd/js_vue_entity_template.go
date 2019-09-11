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
                            v-for="{entity}Item in {entity}List"
                            :key="{entity}Item.Id"
                            @click="select{Entity}Item({entity}Item)"
                            class="sw-table__row_can-select"
                            :class="{'sw-table__row_is-selected': {entity}Item.Id === current{Entity}Item.item.Id}"
                        >
                            <td v-for="(value, key) in fields">
                                <VCheckbox v-if="isCheckbox({entity}Item[key])" :checked="{entity}Item[key]" disabled></VCheckbox>
                                <VText v-else>{{ {entity}Item[key] }}</VText>
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
                                    v-for="(filed, key) in editFields"
                                    vertical-align="center"
                                >
                                    <VLabel
                                        width="col4"
                                        :for="` + "`" + `current{Entity}Item${key}` + "`" + `"
                                    >{{ filed }}</VLabel>
                                    <VInput
										v-if="isInput(current{Entity}Item.item[key])"
                                        v-model="current{Entity}Item.item[key]"
                                        width="dyn"
                                        :id="` + "`" + `current{Entity}Item${key}` + "`" + `"
                                        @input="changeCurrent{Entity}Item"
                                    />
									<VCheckbox
										v-if="isCheckbox(current{Entity}Item.item[key])"
                                        v-model="current{Entity}Item.item[key]"
                                        :id="` + "`" + `current{Entity}Item${key}` + "`" + `"
										@input="changeCurrent{Entity}Item"
									/>
									
                                </VSet>
                            </VSet>
                            <button type="submit" :disabled="!current{Entity}Item.hasChange" hidden></button>
                        </form>
                    </template>

                    <template #footer>
                        <VSet>
                            <VButton
                                @click="saveChangesSubmit"
                                accent
                                :text="panelSubmitButtonText"
                                :disabled="!current{Entity}Item.hasChange"
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
                    v-if="current{Entity}Item.showDeleteConfirmation"
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
                        @click="showPanel(panel.create)"
                    />
                    <VButton
                        text="Удалить"
                        :disabled="!current{Entity}Item.isSelected"
                        @click="delete{Entity}ItemHandler"
                    />
                </VSet>
            </slot>
        </template>
    </WorkSpace>
</template>

<script>
    import {entity}Data from "../data/{Entity}Data";
    import { {Entity} } from '../apiModel';
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
        name: '{Entity}Gen',

        components: {VSelectSimple, VSign, VIcon, VButton, VPanel, VText, VInput, VLabel, VSet, VHead, WorkSpace, VCheckbox},

        props: {
            fields: {
                type: Object,
                default() {
                    const {entity}Item = new {Entity}();
                    const fieldsObj = {};

                    for (let prop in {entity}Item) {

                        if ({entity}Item.hasOwnProperty(prop)) {
                            fieldsObj[prop] = prop;
                        }

                    }

                    return fieldsObj;
                }
            },
            editFields: {
                type: Object,
                default() {
                    const {entity}Item = new {Entity}();
                    const fieldsObj = {};

                    for (let prop in {entity}Item) {

                        if ({entity}Item.hasOwnProperty(prop)) {
                            fieldsObj[prop] = prop;
                        }

                    }

                    return fieldsObj;
                }
            }
        },

        data() {
            return {entity}Data;
        },

        created() {
            this.fill{Entity}Filter();
            this.fetch{Entity}Data();
        },

        computed: {
            ...mapGetters({
                {entity}List: 'getList{Entity}'
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
                'find{Entity}',
                'update{Entity}',
                'delete{Entity}',
                'create{Entity}',
            ]),

            ...mapMutations([
                'add{Entity}ItemToList',
                'delete{Entity}FromList',
                'update{Entity}ById',
            ]),

            fill{Entity}Filter() {
                this.{entity}Filter.CurrentPage = 1;
                this.{entity}Filter.PerPage = 1000;
            },

            fetch{Entity}Data() {
                return this.find{Entity}({
                    filter: this.{entity}Filter
                });
            },

            /**
             *
             * @param type
             */
            showPanel(type) {
                if (type === this.panel.create) {
                    this.panel.type = this.panel.create;
                    this.clearPanel{Entity}Item();
                } else if (type === this.panel.edit) {
                    this.panel.type = this.panel.edit;
                    this.current{Entity}Item.isSelected = true;
                }

                this.panel.show = true;
            },

            closePanel() {
                this.panel.show = false;
                this.current{Entity}Item.isSelected = false;
                this.clearPanel{Entity}Item();
            },

            select{Entity}Item({entity}Item) {
                this.showPanel(this.panel.edit);
                this.current{Entity}Item.isSelected = true;
                Object.assign(this.current{Entity}Item.item, {entity}Item);
            },

            changeCurrent{Entity}Item() {
                this.current{Entity}Item.hasChange = true;
            },

            cancelChanges() {
                this.clearPanel{Entity}Item();
                this.closePanel();
            },

            clearPanel{Entity}Item() {
                this.current{Entity}Item.item = new {Entity}();
                this.current{Entity}Item.hasChange = false;
            },

            saveChangesSubmit() {
                if (this.isPanelCreate) {
                    this.create{Entity}ItemSubmit();
                    return;
                }

                if (this.isPanelEdit) {
                    this.edit{Entity}ItemSubmit();
                }
            },

            create{Entity}ItemSubmit() {
                this.create{Entity}({
                    data: {
                        Name: this.current{Entity}Item.item.Name,
                        Value: this.current{Entity}Item.item.Value,
                        Description: this.current{Entity}Item.item.Description,
                    }
                }).then((response) => {

                    if (response.Model) {
                        this.add{Entity}ItemToList(response.Model);
                        this.clearPanel{Entity}Item();
                    } else {
                        console.error('Ошибка создания записи: ', response.Error);
                    }

                }).catch(error => {
                    console.error('Ошибка создания записи: ', error);
                });
            },

            edit{Entity}ItemSubmit() {
                if (this.current{Entity}Item.hasChange) {
                    this.update{Entity}({
                        id: this.current{Entity}Item.item.Id,
                        data: this.current{Entity}Item.item,
                    }).then((response) => {

                        if (response.Model) {
                            this.update{Entity}ById(response.Model);
                            this.current{Entity}Item.hasChange = false;
                            this.clearPanel{Entity}Item();
                            this.closePanel();
                        } else {
                            console.error('Ошибка изменения записи: ', response.Error);
                        }

                    }).catch(error => {
                        console.error('Ошибка изменения записи: ', error);
                    });
                }
            },

            delete{Entity}ItemHandler() {
                let deletedItemId = this.current{Entity}Item.item.Id;

                if (!this.current{Entity}Item.canDelete) {
                    this.current{Entity}Item.showDeleteConfirmation = true;
                    return;
                }

                this.delete{Entity}({
                    id: deletedItemId
                }).then(response => {

                    if (response.IsSuccess) {
                        this.delete{Entity}FromList(deletedItemId);
                        this.clearPanel{Entity}Item();
                        this.current{Entity}Item.canDelete = false;
                        this.current{Entity}Item.isSelected = false;
                        this.panel.show = false;
                    } else {
                        console.error('Ошибка удаления элемента: ', response.Error);
                    }

                }).catch(error => {
                    console.error('Ошибка удаления элемента: ', error);
                });
            },

            confirmDeleteHandler() {
                this.current{Entity}Item.showDeleteConfirmation = false;
                this.current{Entity}Item.canDelete = true;
                this.delete{Entity}ItemHandler();
            },

            closeConfirmationPanel() {
                this.current{Entity}Item.showDeleteConfirmation = false;
            },
        },
    }
</script>

<style lang="scss">

</style>
`