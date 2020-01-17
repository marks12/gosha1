
	<template>
    <WorkSpace>
        <template #header>
            <slot name="pageHeader">
                <VHead level="h1">UserRole</VHead>
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
                            v-for="userRoleItem in userRoleList"
                            :key="userRoleItem.Id"
                            @click="selectUserRoleItem(userRoleItem)"
                            class="sw-table__row_can-select"
                            :class="{'sw-table__row_is-selected': userRoleItem.Id === currentUserRoleItem.item.Id}"
                        >
                            <td v-for="(value, key) in fields" :key="key + '-fields'">
                                <VCheckbox v-if="isCheckbox(userRoleItem[key])" :checked="userRoleItem[key]" disabled></VCheckbox>
                                <VText v-else>{{ userRoleItem[key] }}</VText>
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
                                        :for="`currentUserRoleItem${key}`"
                                    >{{ filed }}</VLabel>
                                    <VInput
										v-if="isInput(currentUserRoleItem.item[key])"
                                        v-model="currentUserRoleItem.item[key]"
                                        width="dyn"
                                        :id="`currentUserRoleItem${key}`"
                                        @input="changeCurrentUserRoleItem"
                                    />
									<VCheckbox
										v-if="isCheckbox(currentUserRoleItem.item[key])"
                                        v-model="currentUserRoleItem.item[key]"
                                        :id="`currentUserRoleItem${key}`"
										@input="changeCurrentUserRoleItem"
									/>
									
                                </VSet>
                            </VSet>
                            <button type="submit" :disabled="!currentUserRoleItem.hasChange" hidden></button>
                        </form>
                    </template>

                    <template #footer>
                        <VSet>
                            <VButton
                                @click="saveChangesSubmit"
                                accent
                                :text="panelSubmitButtonText"
                                :disabled="!currentUserRoleItem.hasChange"
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
                    v-if="currentUserRoleItem.showDeleteConfirmation"
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
                        :disabled="!currentUserRoleItem.isSelected"
                        @click="deleteUserRoleItemHandler"
                    />
                </VSet>
            </slot>
        </template>
    </WorkSpace>
</template>

<script>
    import userRoleData from "../data/UserRoleData";
    import { UserRole } from '../apiModel';
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
        name: 'UserRoleGen',

        components: {VButton, VPanel, VText, VInput, VLabel, VSet, VHead, WorkSpace, VCheckbox},

        props: {
            fields: {
                type: Object,
                default() {
                    const userRoleItem = new UserRole();
                    const fieldsObj = {};

                    for (let prop in userRoleItem) {

                        if (userRoleItem.hasOwnProperty(prop)) {
                            fieldsObj[prop] = prop;
                        }

                    }

                    return fieldsObj;
                }
            },
            editFields: {
                type: Object,
                default() {
                    const userRoleItem = new UserRole();
                    const fieldsObj = {};

                    for (let prop in userRoleItem) {

                        if (userRoleItem.hasOwnProperty(prop)) {
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
            return userRoleData;
        },

        created() {
			this.onCreated();
        },

        computed: {
            ...mapGetters({
                userRoleList: 'getListUserRole'
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
                'findUserRole',
                'updateUserRole',
                'deleteUserRole',
                'createUserRole',
            ]),

            ...mapMutations([
                'addUserRoleItemToList',
                'deleteUserRoleFromList',
                'updateUserRoleById',
            ]),

			onCreated() {
				this.fillUserRoleFilter();
	            this.fetchUserRoleData();
			},

            fillUserRoleFilter() {
                this.userRoleFilter.CurrentPage = 1;
                this.userRoleFilter.PerPage = 1000;
            },

            fetchUserRoleData() {
                return this.findUserRole({
                    filter: this.userRoleFilter
                });
            },

            /**
             *
             * @param type
             */
            showPanel(type) {
                if (type === this.panel.create) {
                    this.panel.type = this.panel.create;
                    this.clearPanelUserRoleItem();
                } else if (type === this.panel.edit) {
                    this.panel.type = this.panel.edit;
                    this.currentUserRoleItem.isSelected = true;
                }

                this.panel.show = true;
            },

            closePanel() {
                this.panel.show = false;
                this.currentUserRoleItem.isSelected = false;
                this.clearPanelUserRoleItem();
            },

            selectUserRoleItem(userRoleItem) {
                this.showPanel(this.panel.edit);
                this.currentUserRoleItem.isSelected = true;
                Object.assign(this.currentUserRoleItem.item, userRoleItem);
            },

            changeCurrentUserRoleItem() {
                this.currentUserRoleItem.hasChange = true;
            },

            cancelChanges() {
                this.clearPanelUserRoleItem();
                this.closePanel();
            },

            clearPanelUserRoleItem() {
                this.currentUserRoleItem.item = new UserRole();
                this.currentUserRoleItem.hasChange = false;
            },

            saveChangesSubmit() {
                if (this.isPanelCreate) {
                    this.createUserRoleItemSubmit();
                    return;
                }

                if (this.isPanelEdit) {
                    this.editUserRoleItemSubmit();
                }
            },

            createUserRoleItemSubmit() {
                return this.createUserRole({
					data: this.currentUserRoleItem.item,
                }).then((response) => {

                    if (response.Model) {
                        this.addUserRoleItemToList(response.Model);
                        this.clearPanelUserRoleItem();
                    } else {
                        console.error('Ошибка создания записи: ', response.Error);
                    }

                }).catch(error => {
                    console.error('Ошибка создания записи: ', error);
                });
            },

            editUserRoleItemSubmit() {

                if (this.currentUserRoleItem.hasChange) {
                    return this.updateUserRole({
                        id: this.currentUserRoleItem.item.Id,
                        data: this.currentUserRoleItem.item,
                    }).then((response) => {

                        if (response.Model) {
                            this.updateUserRoleById(response.Model);
                            this.currentUserRoleItem.hasChange = false;
                            this.clearPanelUserRoleItem();
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

            deleteUserRoleItemHandler() {
                let deletedItemId = this.currentUserRoleItem.item.Id;

                if (!this.currentUserRoleItem.canDelete) {
                    this.currentUserRoleItem.showDeleteConfirmation = true;
                    return;
                }

                this.deleteUserRole({
                    id: deletedItemId
                }).then(response => {

                    if (response.IsSuccess) {
                        this.deleteUserRoleFromList(deletedItemId);
                        this.clearPanelUserRoleItem();
                        this.currentUserRoleItem.canDelete = false;
                        this.currentUserRoleItem.isSelected = false;
                        this.panel.show = false;
                    } else {
                        console.error('Ошибка удаления элемента: ', response.Error);
                    }

                }).catch(error => {
                    console.error('Ошибка удаления элемента: ', error);
                });
            },

            confirmDeleteHandler() {
                this.currentUserRoleItem.showDeleteConfirmation = false;
                this.currentUserRoleItem.canDelete = true;
                this.deleteUserRoleItemHandler();
            },

            closeConfirmationPanel() {
                this.currentUserRoleItem.showDeleteConfirmation = false;
            },
        },
    }
</script>

<style lang="scss">

</style>
