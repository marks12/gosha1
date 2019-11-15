
	<template>
    <WorkSpace>
        <template #header>
            <slot name="pageHeader">
                <VHead level="h1">UserRoleFilter</VHead>
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
                            v-for="userRoleFilterItem in userRoleFilterList"
                            :key="userRoleFilterItem.Id"
                            @click="selectUserRoleFilterItem(userRoleFilterItem)"
                            class="sw-table__row_can-select"
                            :class="{'sw-table__row_is-selected': userRoleFilterItem.Id === currentUserRoleFilterItem.item.Id}"
                        >
                            <td v-for="(value, key) in fields" :key="key + '-fields'">
                                <VCheckbox v-if="isCheckbox(userRoleFilterItem[key])" :checked="userRoleFilterItem[key]" disabled></VCheckbox>
                                <VText v-else>{{ userRoleFilterItem[key] }}</VText>
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
                                        :for="`currentUserRoleFilterItem${key}`"
                                    >{{ filed }}</VLabel>
                                    <VInput
										v-if="isInput(currentUserRoleFilterItem.item[key])"
                                        v-model="currentUserRoleFilterItem.item[key]"
                                        width="dyn"
                                        :id="`currentUserRoleFilterItem${key}`"
                                        @input="changeCurrentUserRoleFilterItem"
                                    />
									<VCheckbox
										v-if="isCheckbox(currentUserRoleFilterItem.item[key])"
                                        v-model="currentUserRoleFilterItem.item[key]"
                                        :id="`currentUserRoleFilterItem${key}`"
										@input="changeCurrentUserRoleFilterItem"
									/>
									
                                </VSet>
                            </VSet>
                            <button type="submit" :disabled="!currentUserRoleFilterItem.hasChange" hidden></button>
                        </form>
                    </template>

                    <template #footer>
                        <VSet>
                            <VButton
                                @click="saveChangesSubmit"
                                accent
                                :text="panelSubmitButtonText"
                                :disabled="!currentUserRoleFilterItem.hasChange"
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
                    v-if="currentUserRoleFilterItem.showDeleteConfirmation"
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
                        :disabled="!currentUserRoleFilterItem.isSelected"
                        @click="deleteUserRoleFilterItemHandler"
                    />
                </VSet>
            </slot>
        </template>
    </WorkSpace>
</template>

<script>
    import userRoleFilterData from "../data/UserRoleFilterData";
    import { UserRoleFilter } from '../apiModel';
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
        name: 'UserRoleFilterGen',

        components: {VSelectSimple, VSign, VIcon, VButton, VPanel, VText, VInput, VLabel, VSet, VHead, WorkSpace, VCheckbox},

        props: {
            fields: {
                type: Object,
                default() {
                    const userRoleFilterItem = new UserRoleFilter();
                    const fieldsObj = {};

                    for (let prop in userRoleFilterItem) {

                        if (userRoleFilterItem.hasOwnProperty(prop)) {
                            fieldsObj[prop] = prop;
                        }

                    }

                    return fieldsObj;
                }
            },
            editFields: {
                type: Object,
                default() {
                    const userRoleFilterItem = new UserRoleFilter();
                    const fieldsObj = {};

                    for (let prop in userRoleFilterItem) {

                        if (userRoleFilterItem.hasOwnProperty(prop)) {
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
            return userRoleFilterData;
        },

        created() {
            this.fillUserRoleFilterFilter();
            this.fetchUserRoleFilterData();
        },

        computed: {
            ...mapGetters({
                userRoleFilterList: 'getListUserRoleFilter'
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
                'findUserRoleFilter',
                'updateUserRoleFilter',
                'deleteUserRoleFilter',
                'createUserRoleFilter',
            ]),

            ...mapMutations([
                'addUserRoleFilterItemToList',
                'deleteUserRoleFilterFromList',
                'updateUserRoleFilterById',
            ]),

            fillUserRoleFilterFilter() {
                this.userRoleFilterFilter.CurrentPage = 1;
                this.userRoleFilterFilter.PerPage = 1000;
            },

            fetchUserRoleFilterData() {
                return this.findUserRoleFilter({
                    filter: this.userRoleFilterFilter
                });
            },

            /**
             *
             * @param type
             */
            showPanel(type) {
                if (type === this.panel.create) {
                    this.panel.type = this.panel.create;
                    this.clearPanelUserRoleFilterItem();
                } else if (type === this.panel.edit) {
                    this.panel.type = this.panel.edit;
                    this.currentUserRoleFilterItem.isSelected = true;
                }

                this.panel.show = true;
            },

            closePanel() {
                this.panel.show = false;
                this.currentUserRoleFilterItem.isSelected = false;
                this.clearPanelUserRoleFilterItem();
            },

            selectUserRoleFilterItem(userRoleFilterItem) {
                this.showPanel(this.panel.edit);
                this.currentUserRoleFilterItem.isSelected = true;
                Object.assign(this.currentUserRoleFilterItem.item, userRoleFilterItem);
            },

            changeCurrentUserRoleFilterItem() {
                this.currentUserRoleFilterItem.hasChange = true;
            },

            cancelChanges() {
                this.clearPanelUserRoleFilterItem();
                this.closePanel();
            },

            clearPanelUserRoleFilterItem() {
                this.currentUserRoleFilterItem.item = new UserRoleFilter();
                this.currentUserRoleFilterItem.hasChange = false;
            },

            saveChangesSubmit() {
                if (this.isPanelCreate) {
                    this.createUserRoleFilterItemSubmit();
                    return;
                }

                if (this.isPanelEdit) {
                    this.editUserRoleFilterItemSubmit();
                }
            },

            createUserRoleFilterItemSubmit() {
                return this.createUserRoleFilter({
					data: this.currentUserRoleFilterItem.item,
                }).then((response) => {

                    if (response.Model) {
                        this.addUserRoleFilterItemToList(response.Model);
                        this.clearPanelUserRoleFilterItem();
                    } else {
                        console.error('Ошибка создания записи: ', response.Error);
                    }

                }).catch(error => {
                    console.error('Ошибка создания записи: ', error);
                });
            },

            editUserRoleFilterItemSubmit() {

                if (this.currentUserRoleFilterItem.hasChange) {
                    return this.updateUserRoleFilter({
                        id: this.currentUserRoleFilterItem.item.Id,
                        data: this.currentUserRoleFilterItem.item,
                    }).then((response) => {

                        if (response.Model) {
                            this.updateUserRoleFilterById(response.Model);
                            this.currentUserRoleFilterItem.hasChange = false;
                            this.clearPanelUserRoleFilterItem();
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

            deleteUserRoleFilterItemHandler() {
                let deletedItemId = this.currentUserRoleFilterItem.item.Id;

                if (!this.currentUserRoleFilterItem.canDelete) {
                    this.currentUserRoleFilterItem.showDeleteConfirmation = true;
                    return;
                }

                this.deleteUserRoleFilter({
                    id: deletedItemId
                }).then(response => {

                    if (response.IsSuccess) {
                        this.deleteUserRoleFilterFromList(deletedItemId);
                        this.clearPanelUserRoleFilterItem();
                        this.currentUserRoleFilterItem.canDelete = false;
                        this.currentUserRoleFilterItem.isSelected = false;
                        this.panel.show = false;
                    } else {
                        console.error('Ошибка удаления элемента: ', response.Error);
                    }

                }).catch(error => {
                    console.error('Ошибка удаления элемента: ', error);
                });
            },

            confirmDeleteHandler() {
                this.currentUserRoleFilterItem.showDeleteConfirmation = false;
                this.currentUserRoleFilterItem.canDelete = true;
                this.deleteUserRoleFilterItemHandler();
            },

            closeConfirmationPanel() {
                this.currentUserRoleFilterItem.showDeleteConfirmation = false;
            },
        },
    }
</script>

<style lang="scss">

</style>
