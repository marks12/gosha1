
	<template>
    <WorkSpace>
        <template #header>
            <slot name="pageHeader">
                <VHead level="h1">Auth</VHead>
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
                            v-for="authItem in authList"
                            :key="authItem.Id"
                            @click="selectAuthItem(authItem)"
                            class="sw-table__row_can-select"
                            :class="{'sw-table__row_is-selected': authItem.Id === currentAuthItem.item.Id}"
                        >
                            <td v-for="(value, key) in fields">
                                <VCheckbox v-if="isCheckbox(authItem[key])" :checked="authItem[key]" disabled></VCheckbox>
                                <VText v-else>{{ authItem[key] }}</VText>
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
                                        :for="`currentAuthItem${key}`"
                                    >{{ filed }}</VLabel>
                                    <VInput
										v-if="isInput(currentAuthItem.item[key])"
                                        v-model="currentAuthItem.item[key]"
                                        width="dyn"
                                        :id="`currentAuthItem${key}`"
                                        @input="changeCurrentAuthItem"
                                    />
									<VCheckbox
										v-if="isCheckbox(currentAuthItem.item[key])"
                                        v-model="currentAuthItem.item[key]"
                                        :id="`currentAuthItem${key}`"
										@input="changeCurrentApplicationItem"
									/>
									
                                </VSet>
                            </VSet>
                            <button type="submit" :disabled="!currentAuthItem.hasChange" hidden></button>
                        </form>
                    </template>

                    <template #footer>
                        <VSet>
                            <VButton
                                @click="saveChangesSubmit"
                                accent
                                :text="panelSubmitButtonText"
                                :disabled="!currentAuthItem.hasChange"
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
                    v-if="currentAuthItem.showDeleteConfirmation"
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
                        :disabled="!currentAuthItem.isSelected"
                        @click="deleteAuthItemHandler"
                    />
                </VSet>
            </slot>
        </template>
    </WorkSpace>
</template>

<script>
    import authData from "../data/AuthData";
    import { Auth } from '../apiModel';
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
    import VSelect from "swui/src/components/VSelect";

    export default {
        name: 'AuthGen',

        components: {VSelect, VSign, VIcon, VButton, VPanel, VText, VInput, VLabel, VSet, VHead, WorkSpace, VCheckbox},

        props: {
            fields: {
                type: Object,
                default() {
                    const authItem = new Auth();
                    const fieldsObj = {};

                    for (let prop in authItem) {

                        if (authItem.hasOwnProperty(prop)) {
                            fieldsObj[prop] = prop;
                        }

                    }

                    return fieldsObj;
                }
            },
            editFields: {
                type: Object,
                default() {
                    const authItem = new Auth();
                    const fieldsObj = {};

                    for (let prop in authItem) {

                        if (authItem.hasOwnProperty(prop)) {
                            fieldsObj[prop] = prop;
                        }

                    }

                    return fieldsObj;
                }
            }
        },

        data() {
            return authData;
        },

        created() {
            this.fillAuthFilter();
            this.fetchAuthData();
        },

        computed: {
            ...mapGetters({
                authList: 'getListAuth'
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
                'findAuth',
                'updateAuth',
                'deleteAuth',
                'createAuth',
            ]),

            ...mapMutations([
                'addAuthItemToList',
                'deleteAuthFromList',
                'updateAuthById',
            ]),

            fillAuthFilter() {
                this.authFilter.CurrentPage = 1;
                this.authFilter.PerPage = 1000;
            },

            fetchAuthData() {
                return this.findAuth({
                    filter: this.authFilter
                });
            },

            /**
             *
             * @param type
             */
            showPanel(type) {
                if (type === this.panel.create) {
                    this.panel.type = this.panel.create;
                    this.clearPanelAuthItem();
                } else if (type === this.panel.edit) {
                    this.panel.type = this.panel.edit;
                    this.currentAuthItem.isSelected = true;
                }

                this.panel.show = true;
            },

            closePanel() {
                this.panel.show = false;
                this.currentAuthItem.isSelected = false;
                this.clearPanelAuthItem();
            },

            selectAuthItem(authItem) {
                this.showPanel(this.panel.edit);
                this.currentAuthItem.isSelected = true;
                Object.assign(this.currentAuthItem.item, authItem);
            },

            changeCurrentAuthItem() {
                this.currentAuthItem.hasChange = true;
            },

            cancelChanges() {
                this.clearPanelAuthItem();
                this.closePanel();
            },

            clearPanelAuthItem() {
                this.currentAuthItem.item = new Auth();
                this.currentAuthItem.hasChange = false;
            },

            saveChangesSubmit() {
                if (this.isPanelCreate) {
                    this.createAuthItemSubmit();
                    return;
                }

                if (this.isPanelEdit) {
                    this.editAuthItemSubmit();
                }
            },

            createAuthItemSubmit() {
                this.createAuth({
                    data: {
                        Name: this.currentAuthItem.item.Name,
                        Value: this.currentAuthItem.item.Value,
                        Description: this.currentAuthItem.item.Description,
                    }
                }).then((response) => {

                    if (response.Model) {
                        this.addAuthItemToList(response.Model);
                        this.clearPanelAuthItem();
                    } else {
                        console.error('Ошибка создания записи: ', response.Error);
                    }

                }).catch(error => {
                    console.error('Ошибка создания записи: ', error);
                });
            },

            editAuthItemSubmit() {
                if (this.currentAuthItem.hasChange) {
                    this.updateAuth({
                        id: this.currentAuthItem.item.Id,
                        data: this.currentAuthItem.item,
                    }).then((response) => {

                        if (response.Model) {
                            this.updateAuthById(response.Model);
                            this.currentAuthItem.hasChange = false;
                            this.clearPanelAuthItem();
                            this.closePanel();
                        } else {
                            console.error('Ошибка изменения записи: ', response.Error);
                        }

                    }).catch(error => {
                        console.error('Ошибка изменения записи: ', error);
                    });
                }
            },

            deleteAuthItemHandler() {
                let deletedItemId = this.currentAuthItem.item.Id;

                if (!this.currentAuthItem.canDelete) {
                    this.currentAuthItem.showDeleteConfirmation = true;
                    return;
                }

                this.deleteAuth({
                    id: deletedItemId
                }).then(response => {

                    if (response.IsSuccess) {
                        this.deleteAuthFromList(deletedItemId);
                        this.clearPanelAuthItem();
                        this.currentAuthItem.canDelete = false;
                        this.currentAuthItem.isSelected = false;
                        this.panel.show = false;
                    } else {
                        console.error('Ошибка удаления элемента: ', response.Error);
                    }

                }).catch(error => {
                    console.error('Ошибка удаления элемента: ', error);
                });
            },

            confirmDeleteHandler() {
                this.currentAuthItem.showDeleteConfirmation = false;
                this.currentAuthItem.canDelete = true;
                this.deleteAuthItemHandler();
            },

            closeConfirmationPanel() {
                this.currentAuthItem.showDeleteConfirmation = false;
            },
        },
    }
</script>

<style lang="scss">

</style>
