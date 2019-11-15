
	<template>
    <WorkSpace>
        <template #header>
            <slot name="pageHeader">
                <VHead level="h1">Authenticator</VHead>
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
                            v-for="authenticatorItem in authenticatorList"
                            :key="authenticatorItem.Id"
                            @click="selectAuthenticatorItem(authenticatorItem)"
                            class="sw-table__row_can-select"
                            :class="{'sw-table__row_is-selected': authenticatorItem.Id === currentAuthenticatorItem.item.Id}"
                        >
                            <td v-for="(value, key) in fields" :key="key + '-fields'">
                                <VCheckbox v-if="isCheckbox(authenticatorItem[key])" :checked="authenticatorItem[key]" disabled></VCheckbox>
                                <VText v-else>{{ authenticatorItem[key] }}</VText>
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
                                        :for="`currentAuthenticatorItem${key}`"
                                    >{{ filed }}</VLabel>
                                    <VInput
										v-if="isInput(currentAuthenticatorItem.item[key])"
                                        v-model="currentAuthenticatorItem.item[key]"
                                        width="dyn"
                                        :id="`currentAuthenticatorItem${key}`"
                                        @input="changeCurrentAuthenticatorItem"
                                    />
									<VCheckbox
										v-if="isCheckbox(currentAuthenticatorItem.item[key])"
                                        v-model="currentAuthenticatorItem.item[key]"
                                        :id="`currentAuthenticatorItem${key}`"
										@input="changeCurrentAuthenticatorItem"
									/>
									
                                </VSet>
                            </VSet>
                            <button type="submit" :disabled="!currentAuthenticatorItem.hasChange" hidden></button>
                        </form>
                    </template>

                    <template #footer>
                        <VSet>
                            <VButton
                                @click="saveChangesSubmit"
                                accent
                                :text="panelSubmitButtonText"
                                :disabled="!currentAuthenticatorItem.hasChange"
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
                    v-if="currentAuthenticatorItem.showDeleteConfirmation"
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
                        :disabled="!currentAuthenticatorItem.isSelected"
                        @click="deleteAuthenticatorItemHandler"
                    />
                </VSet>
            </slot>
        </template>
    </WorkSpace>
</template>

<script>
    import authenticatorData from "../data/AuthenticatorData";
    import { Authenticator } from '../apiModel';
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
        name: 'AuthenticatorGen',

        components: {VSelectSimple, VSign, VIcon, VButton, VPanel, VText, VInput, VLabel, VSet, VHead, WorkSpace, VCheckbox},

        props: {
            fields: {
                type: Object,
                default() {
                    const authenticatorItem = new Authenticator();
                    const fieldsObj = {};

                    for (let prop in authenticatorItem) {

                        if (authenticatorItem.hasOwnProperty(prop)) {
                            fieldsObj[prop] = prop;
                        }

                    }

                    return fieldsObj;
                }
            },
            editFields: {
                type: Object,
                default() {
                    const authenticatorItem = new Authenticator();
                    const fieldsObj = {};

                    for (let prop in authenticatorItem) {

                        if (authenticatorItem.hasOwnProperty(prop)) {
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
            return authenticatorData;
        },

        created() {
            this.fillAuthenticatorFilter();
            this.fetchAuthenticatorData();
        },

        computed: {
            ...mapGetters({
                authenticatorList: 'getListAuthenticator'
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
                'findAuthenticator',
                'updateAuthenticator',
                'deleteAuthenticator',
                'createAuthenticator',
            ]),

            ...mapMutations([
                'addAuthenticatorItemToList',
                'deleteAuthenticatorFromList',
                'updateAuthenticatorById',
            ]),

            fillAuthenticatorFilter() {
                this.authenticatorFilter.CurrentPage = 1;
                this.authenticatorFilter.PerPage = 1000;
            },

            fetchAuthenticatorData() {
                return this.findAuthenticator({
                    filter: this.authenticatorFilter
                });
            },

            /**
             *
             * @param type
             */
            showPanel(type) {
                if (type === this.panel.create) {
                    this.panel.type = this.panel.create;
                    this.clearPanelAuthenticatorItem();
                } else if (type === this.panel.edit) {
                    this.panel.type = this.panel.edit;
                    this.currentAuthenticatorItem.isSelected = true;
                }

                this.panel.show = true;
            },

            closePanel() {
                this.panel.show = false;
                this.currentAuthenticatorItem.isSelected = false;
                this.clearPanelAuthenticatorItem();
            },

            selectAuthenticatorItem(authenticatorItem) {
                this.showPanel(this.panel.edit);
                this.currentAuthenticatorItem.isSelected = true;
                Object.assign(this.currentAuthenticatorItem.item, authenticatorItem);
            },

            changeCurrentAuthenticatorItem() {
                this.currentAuthenticatorItem.hasChange = true;
            },

            cancelChanges() {
                this.clearPanelAuthenticatorItem();
                this.closePanel();
            },

            clearPanelAuthenticatorItem() {
                this.currentAuthenticatorItem.item = new Authenticator();
                this.currentAuthenticatorItem.hasChange = false;
            },

            saveChangesSubmit() {
                if (this.isPanelCreate) {
                    this.createAuthenticatorItemSubmit();
                    return;
                }

                if (this.isPanelEdit) {
                    this.editAuthenticatorItemSubmit();
                }
            },

            createAuthenticatorItemSubmit() {
                return this.createAuthenticator({
					data: this.currentAuthenticatorItem.item,
                }).then((response) => {

                    if (response.Model) {
                        this.addAuthenticatorItemToList(response.Model);
                        this.clearPanelAuthenticatorItem();
                    } else {
                        console.error('Ошибка создания записи: ', response.Error);
                    }

                }).catch(error => {
                    console.error('Ошибка создания записи: ', error);
                });
            },

            editAuthenticatorItemSubmit() {

                if (this.currentAuthenticatorItem.hasChange) {
                    return this.updateAuthenticator({
                        id: this.currentAuthenticatorItem.item.Id,
                        data: this.currentAuthenticatorItem.item,
                    }).then((response) => {

                        if (response.Model) {
                            this.updateAuthenticatorById(response.Model);
                            this.currentAuthenticatorItem.hasChange = false;
                            this.clearPanelAuthenticatorItem();
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

            deleteAuthenticatorItemHandler() {
                let deletedItemId = this.currentAuthenticatorItem.item.Id;

                if (!this.currentAuthenticatorItem.canDelete) {
                    this.currentAuthenticatorItem.showDeleteConfirmation = true;
                    return;
                }

                this.deleteAuthenticator({
                    id: deletedItemId
                }).then(response => {

                    if (response.IsSuccess) {
                        this.deleteAuthenticatorFromList(deletedItemId);
                        this.clearPanelAuthenticatorItem();
                        this.currentAuthenticatorItem.canDelete = false;
                        this.currentAuthenticatorItem.isSelected = false;
                        this.panel.show = false;
                    } else {
                        console.error('Ошибка удаления элемента: ', response.Error);
                    }

                }).catch(error => {
                    console.error('Ошибка удаления элемента: ', error);
                });
            },

            confirmDeleteHandler() {
                this.currentAuthenticatorItem.showDeleteConfirmation = false;
                this.currentAuthenticatorItem.canDelete = true;
                this.deleteAuthenticatorItemHandler();
            },

            closeConfirmationPanel() {
                this.currentAuthenticatorItem.showDeleteConfirmation = false;
            },
        },
    }
</script>

<style lang="scss">

</style>
