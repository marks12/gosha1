
	<template>
    <WorkSpace>
        <template #header>
            <slot name="pageHeader">
                <VHead level="h1">UserFilter</VHead>
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
                            v-for="userFilterItem in userFilterList"
                            :key="userFilterItem.Id"
                            @click="selectUserFilterItem(userFilterItem)"
                            class="sw-table__row_can-select"
                            :class="{'sw-table__row_is-selected': userFilterItem.Id === currentUserFilterItem.item.Id}"
                        >
                            <td v-for="(value, key) in fields" :key="key + '-fields'">
                                <VCheckbox v-if="isCheckbox(userFilterItem[key])" :checked="userFilterItem[key]" disabled></VCheckbox>
                                <VText v-else>{{ userFilterItem[key] }}</VText>
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
                                        :for="`currentUserFilterItem${key}`"
                                    >{{ filed }}</VLabel>
                                    <VInput
										v-if="isInput(currentUserFilterItem.item[key])"
                                        v-model="currentUserFilterItem.item[key]"
                                        width="dyn"
                                        :id="`currentUserFilterItem${key}`"
                                        @input="changeCurrentUserFilterItem"
                                    />
									<VCheckbox
										v-if="isCheckbox(currentUserFilterItem.item[key])"
                                        v-model="currentUserFilterItem.item[key]"
                                        :id="`currentUserFilterItem${key}`"
										@input="changeCurrentUserFilterItem"
									/>
									
                                </VSet>
                            </VSet>
                            <button type="submit" :disabled="!currentUserFilterItem.hasChange" hidden></button>
                        </form>
                    </template>

                    <template #footer>
                        <VSet>
                            <VButton
                                @click="saveChangesSubmit"
                                accent
                                :text="panelSubmitButtonText"
                                :disabled="!currentUserFilterItem.hasChange"
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
                    v-if="currentUserFilterItem.showDeleteConfirmation"
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
                        :disabled="!currentUserFilterItem.isSelected"
                        @click="deleteUserFilterItemHandler"
                    />
                </VSet>
            </slot>
        </template>
    </WorkSpace>
</template>

<script>
    import userFilterData from "../data/UserFilterData";
    import { UserFilter } from '../apiModel';
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
        name: 'UserFilterGen',

        components: {VButton, VPanel, VText, VInput, VLabel, VSet, VHead, WorkSpace, VCheckbox},

        props: {
            fields: {
                type: Object,
                default() {
                    const userFilterItem = new UserFilter();
                    const fieldsObj = {};

                    for (let prop in userFilterItem) {

                        if (userFilterItem.hasOwnProperty(prop)) {
                            fieldsObj[prop] = prop;
                        }

                    }

                    return fieldsObj;
                }
            },
            editFields: {
                type: Object,
                default() {
                    const userFilterItem = new UserFilter();
                    const fieldsObj = {};

                    for (let prop in userFilterItem) {

                        if (userFilterItem.hasOwnProperty(prop)) {
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
            return userFilterData;
        },

        created() {
			this.onCreated();
        },

        computed: {
            ...mapGetters({
                userFilterList: 'getListUserFilter'
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
                'findUserFilter',
                'updateUserFilter',
                'deleteUserFilter',
                'createUserFilter',
            ]),

            ...mapMutations([
                'addUserFilterItemToList',
                'deleteUserFilterFromList',
                'updateUserFilterById',
            ]),

			onCreated() {
				this.fillUserFilterFilter();
	            this.fetchUserFilterData();
			},

            fillUserFilterFilter() {
                this.userFilterFilter.CurrentPage = 1;
                this.userFilterFilter.PerPage = 1000;
            },

            fetchUserFilterData() {
                return this.findUserFilter({
                    filter: this.userFilterFilter
                });
            },

            /**
             *
             * @param type
             */
            showPanel(type) {
                if (type === this.panel.create) {
                    this.panel.type = this.panel.create;
                    this.clearPanelUserFilterItem();
                } else if (type === this.panel.edit) {
                    this.panel.type = this.panel.edit;
                    this.currentUserFilterItem.isSelected = true;
                }

                this.panel.show = true;
            },

            closePanel() {
                this.panel.show = false;
                this.currentUserFilterItem.isSelected = false;
                this.clearPanelUserFilterItem();
            },

            selectUserFilterItem(userFilterItem) {
                this.showPanel(this.panel.edit);
                this.currentUserFilterItem.isSelected = true;
                Object.assign(this.currentUserFilterItem.item, userFilterItem);
            },

            changeCurrentUserFilterItem() {
                this.currentUserFilterItem.hasChange = true;
            },

            cancelChanges() {
                this.clearPanelUserFilterItem();
                this.closePanel();
            },

            clearPanelUserFilterItem() {
                this.currentUserFilterItem.item = new UserFilter();
                this.currentUserFilterItem.hasChange = false;
            },

            saveChangesSubmit() {
                if (this.isPanelCreate) {
                    this.createUserFilterItemSubmit();
                    return;
                }

                if (this.isPanelEdit) {
                    this.editUserFilterItemSubmit();
                }
            },

            createUserFilterItemSubmit() {
                return this.createUserFilter({
					data: this.currentUserFilterItem.item,
                }).then((response) => {

                    if (response.Model) {
                        this.addUserFilterItemToList(response.Model);
                        this.clearPanelUserFilterItem();
                    } else {
                        console.error('Ошибка создания записи: ', response.Error);
                    }

                }).catch(error => {
                    console.error('Ошибка создания записи: ', error);
                });
            },

            editUserFilterItemSubmit() {

                if (this.currentUserFilterItem.hasChange) {
                    return this.updateUserFilter({
                        id: this.currentUserFilterItem.item.Id,
                        data: this.currentUserFilterItem.item,
                    }).then((response) => {

                        if (response.Model) {
                            this.updateUserFilterById(response.Model);
                            this.currentUserFilterItem.hasChange = false;
                            this.clearPanelUserFilterItem();
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

            deleteUserFilterItemHandler() {
                let deletedItemId = this.currentUserFilterItem.item.Id;

                if (!this.currentUserFilterItem.canDelete) {
                    this.currentUserFilterItem.showDeleteConfirmation = true;
                    return;
                }

                this.deleteUserFilter({
                    id: deletedItemId
                }).then(response => {

                    if (response.IsSuccess) {
                        this.deleteUserFilterFromList(deletedItemId);
                        this.clearPanelUserFilterItem();
                        this.currentUserFilterItem.canDelete = false;
                        this.currentUserFilterItem.isSelected = false;
                        this.panel.show = false;
                    } else {
                        console.error('Ошибка удаления элемента: ', response.Error);
                    }

                }).catch(error => {
                    console.error('Ошибка удаления элемента: ', error);
                });
            },

            confirmDeleteHandler() {
                this.currentUserFilterItem.showDeleteConfirmation = false;
                this.currentUserFilterItem.canDelete = true;
                this.deleteUserFilterItemHandler();
            },

            closeConfirmationPanel() {
                this.currentUserFilterItem.showDeleteConfirmation = false;
            },
        },
    }
</script>

<style lang="scss">

</style>
