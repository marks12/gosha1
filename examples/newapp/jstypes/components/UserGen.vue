
	<template>
    <WorkSpace>
        <template #header>
            <slot name="pageHeader">
                <VHead level="h1">User</VHead>
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
                            v-for="userItem in userList"
                            :key="userItem.Id"
                            @click="selectUserItem(userItem)"
                            class="sw-table__row_can-select"
                            :class="{'sw-table__row_is-selected': userItem.Id === currentUserItem.item.Id}"
                        >
                            <td v-for="(value, key) in fields" :key="key + '-fields'">
                                <VCheckbox v-if="isCheckbox(userItem[key])" :checked="userItem[key]" disabled></VCheckbox>
                                <VText v-else>{{ userItem[key] }}</VText>
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
                                        :for="`currentUserItem${key}`"
                                    >{{ filed }}</VLabel>
                                    <VInput
										v-if="isInput(currentUserItem.item[key])"
                                        v-model="currentUserItem.item[key]"
                                        width="dyn"
                                        :id="`currentUserItem${key}`"
                                        @input="changeCurrentUserItem"
                                    />
									<VCheckbox
										v-if="isCheckbox(currentUserItem.item[key])"
                                        v-model="currentUserItem.item[key]"
                                        :id="`currentUserItem${key}`"
										@input="changeCurrentUserItem"
									/>
									
                                </VSet>
                            </VSet>
                            <button type="submit" :disabled="!currentUserItem.hasChange" hidden></button>
                        </form>
                    </template>

                    <template #footer>
                        <VSet>
                            <VButton
                                @click="saveChangesSubmit"
                                accent
                                :text="panelSubmitButtonText"
                                :disabled="!currentUserItem.hasChange"
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
                    v-if="currentUserItem.showDeleteConfirmation"
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
                        :disabled="!currentUserItem.isSelected"
                        @click="deleteUserItemHandler"
                    />
                </VSet>
            </slot>
        </template>
    </WorkSpace>
</template>

<script>
    import userData from "../data/UserData";
    import { User } from '../apiModel';
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
        name: 'UserGen',

        components: {VButton, VPanel, VText, VInput, VLabel, VSet, VHead, WorkSpace, VCheckbox},

        props: {
            fields: {
                type: Object,
                default() {
                    const userItem = new User();
                    const fieldsObj = {};

                    for (let prop in userItem) {

                        if (userItem.hasOwnProperty(prop)) {
                            fieldsObj[prop] = prop;
                        }

                    }

                    return fieldsObj;
                }
            },
            editFields: {
                type: Object,
                default() {
                    const userItem = new User();
                    const fieldsObj = {};

                    for (let prop in userItem) {

                        if (userItem.hasOwnProperty(prop)) {
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
            return userData;
        },

        created() {
			this.onCreated();
        },

        computed: {
            ...mapGetters({
                userList: 'getListUser'
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
                'findUser',
                'updateUser',
                'deleteUser',
                'createUser',
            ]),

            ...mapMutations([
                'addUserItemToList',
                'deleteUserFromList',
                'updateUserById',
            ]),

			onCreated() {
				this.fillUserFilter();
	            this.fetchUserData();
			},

            fillUserFilter() {
                this.userFilter.CurrentPage = 1;
                this.userFilter.PerPage = 1000;
            },

            fetchUserData() {
                return this.findUser({
                    filter: this.userFilter
                });
            },

            /**
             *
             * @param type
             */
            showPanel(type) {
                if (type === this.panel.create) {
                    this.panel.type = this.panel.create;
                    this.clearPanelUserItem();
                } else if (type === this.panel.edit) {
                    this.panel.type = this.panel.edit;
                    this.currentUserItem.isSelected = true;
                }

                this.panel.show = true;
            },

            closePanel() {
                this.panel.show = false;
                this.currentUserItem.isSelected = false;
                this.clearPanelUserItem();
            },

            selectUserItem(userItem) {
                this.showPanel(this.panel.edit);
                this.currentUserItem.isSelected = true;
                Object.assign(this.currentUserItem.item, userItem);
            },

            changeCurrentUserItem() {
                this.currentUserItem.hasChange = true;
            },

            cancelChanges() {
                this.clearPanelUserItem();
                this.closePanel();
            },

            clearPanelUserItem() {
                this.currentUserItem.item = new User();
                this.currentUserItem.hasChange = false;
            },

            saveChangesSubmit() {
                if (this.isPanelCreate) {
                    this.createUserItemSubmit();
                    return;
                }

                if (this.isPanelEdit) {
                    this.editUserItemSubmit();
                }
            },

            createUserItemSubmit() {
                return this.createUser({
					data: this.currentUserItem.item,
                }).then((response) => {

                    if (response.Model) {
                        this.addUserItemToList(response.Model);
                        this.clearPanelUserItem();
                    } else {
                        console.error('Ошибка создания записи: ', response.Error);
                    }

                }).catch(error => {
                    console.error('Ошибка создания записи: ', error);
                });
            },

            editUserItemSubmit() {

                if (this.currentUserItem.hasChange) {
                    return this.updateUser({
                        id: this.currentUserItem.item.Id,
                        data: this.currentUserItem.item,
                    }).then((response) => {

                        if (response.Model) {
                            this.updateUserById(response.Model);
                            this.currentUserItem.hasChange = false;
                            this.clearPanelUserItem();
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

            deleteUserItemHandler() {
                let deletedItemId = this.currentUserItem.item.Id;

                if (!this.currentUserItem.canDelete) {
                    this.currentUserItem.showDeleteConfirmation = true;
                    return;
                }

                this.deleteUser({
                    id: deletedItemId
                }).then(response => {

                    if (response.IsSuccess) {
                        this.deleteUserFromList(deletedItemId);
                        this.clearPanelUserItem();
                        this.currentUserItem.canDelete = false;
                        this.currentUserItem.isSelected = false;
                        this.panel.show = false;
                    } else {
                        console.error('Ошибка удаления элемента: ', response.Error);
                    }

                }).catch(error => {
                    console.error('Ошибка удаления элемента: ', error);
                });
            },

            confirmDeleteHandler() {
                this.currentUserItem.showDeleteConfirmation = false;
                this.currentUserItem.canDelete = true;
                this.deleteUserItemHandler();
            },

            closeConfirmationPanel() {
                this.currentUserItem.showDeleteConfirmation = false;
            },
        },
    }
</script>

<style lang="scss">

</style>
