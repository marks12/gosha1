
	<template>
    <WorkSpace>
        <template #header>
            <slot name="pageHeader">
                <VHead level="h1">CurrentUser</VHead>
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
                            v-for="currentUserItem in currentUserList"
                            :key="currentUserItem.Id"
                            @click="selectCurrentUserItem(currentUserItem)"
                            class="sw-table__row_can-select"
                            :class="{'sw-table__row_is-selected': currentUserItem.Id === currentCurrentUserItem.item.Id}"
                        >
                            <td v-for="(value, key) in fields" :key="key + '-fields'">
                                <VCheckbox v-if="isCheckbox(currentUserItem[key])" :checked="currentUserItem[key]" disabled></VCheckbox>
                                <VText v-else>{{ currentUserItem[key] }}</VText>
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
                                        :for="`currentCurrentUserItem${key}`"
                                    >{{ filed }}</VLabel>
                                    <VInput
										v-if="isInput(currentCurrentUserItem.item[key])"
                                        v-model="currentCurrentUserItem.item[key]"
                                        width="dyn"
                                        :id="`currentCurrentUserItem${key}`"
                                        @input="changeCurrentCurrentUserItem"
                                    />
									<VCheckbox
										v-if="isCheckbox(currentCurrentUserItem.item[key])"
                                        v-model="currentCurrentUserItem.item[key]"
                                        :id="`currentCurrentUserItem${key}`"
										@input="changeCurrentCurrentUserItem"
									/>
									
                                </VSet>
                            </VSet>
                            <button type="submit" :disabled="!currentCurrentUserItem.hasChange" hidden></button>
                        </form>
                    </template>

                    <template #footer>
                        <VSet>
                            <VButton
                                @click="saveChangesSubmit"
                                accent
                                :text="panelSubmitButtonText"
                                :disabled="!currentCurrentUserItem.hasChange"
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
                    v-if="currentCurrentUserItem.showDeleteConfirmation"
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
                        :disabled="!currentCurrentUserItem.isSelected"
                        @click="deleteCurrentUserItemHandler"
                    />
                </VSet>
            </slot>
        </template>
    </WorkSpace>
</template>

<script>
    import currentUserData from "../data/CurrentUserData";
    import { CurrentUser } from '../apiModel';
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
        name: 'CurrentUserGen',

        components: {VButton, VPanel, VText, VInput, VLabel, VSet, VHead, WorkSpace, VCheckbox},

        props: {
            fields: {
                type: Object,
                default() {
                    const currentUserItem = new CurrentUser();
                    const fieldsObj = {};

                    for (let prop in currentUserItem) {

                        if (currentUserItem.hasOwnProperty(prop)) {
                            fieldsObj[prop] = prop;
                        }

                    }

                    return fieldsObj;
                }
            },
            editFields: {
                type: Object,
                default() {
                    const currentUserItem = new CurrentUser();
                    const fieldsObj = {};

                    for (let prop in currentUserItem) {

                        if (currentUserItem.hasOwnProperty(prop)) {
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
            return currentUserData;
        },

        created() {
			this.onCreated();
        },

        computed: {
            ...mapGetters({
                currentUserList: 'getListCurrentUser'
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
                'findCurrentUser',
                'updateCurrentUser',
                'deleteCurrentUser',
                'createCurrentUser',
            ]),

            ...mapMutations([
                'addCurrentUserItemToList',
                'deleteCurrentUserFromList',
                'updateCurrentUserById',
            ]),

			onCreated() {
				this.fillCurrentUserFilter();
	            this.fetchCurrentUserData();
			},

            fillCurrentUserFilter() {
                this.currentUserFilter.CurrentPage = 1;
                this.currentUserFilter.PerPage = 1000;
            },

            fetchCurrentUserData() {
                return this.findCurrentUser({
                    filter: this.currentUserFilter
                });
            },

            /**
             *
             * @param type
             */
            showPanel(type) {
                if (type === this.panel.create) {
                    this.panel.type = this.panel.create;
                    this.clearPanelCurrentUserItem();
                } else if (type === this.panel.edit) {
                    this.panel.type = this.panel.edit;
                    this.currentCurrentUserItem.isSelected = true;
                }

                this.panel.show = true;
            },

            closePanel() {
                this.panel.show = false;
                this.currentCurrentUserItem.isSelected = false;
                this.clearPanelCurrentUserItem();
            },

            selectCurrentUserItem(currentUserItem) {
                this.showPanel(this.panel.edit);
                this.currentCurrentUserItem.isSelected = true;
                Object.assign(this.currentCurrentUserItem.item, currentUserItem);
            },

            changeCurrentCurrentUserItem() {
                this.currentCurrentUserItem.hasChange = true;
            },

            cancelChanges() {
                this.clearPanelCurrentUserItem();
                this.closePanel();
            },

            clearPanelCurrentUserItem() {
                this.currentCurrentUserItem.item = new CurrentUser();
                this.currentCurrentUserItem.hasChange = false;
            },

            saveChangesSubmit() {
                if (this.isPanelCreate) {
                    this.createCurrentUserItemSubmit();
                    return;
                }

                if (this.isPanelEdit) {
                    this.editCurrentUserItemSubmit();
                }
            },

            createCurrentUserItemSubmit() {
                return this.createCurrentUser({
					data: this.currentCurrentUserItem.item,
                }).then((response) => {

                    if (response.Model) {
                        this.addCurrentUserItemToList(response.Model);
                        this.clearPanelCurrentUserItem();
                    } else {
                        console.error('Ошибка создания записи: ', response.Error);
                    }

                }).catch(error => {
                    console.error('Ошибка создания записи: ', error);
                });
            },

            editCurrentUserItemSubmit() {

                if (this.currentCurrentUserItem.hasChange) {
                    return this.updateCurrentUser({
                        id: this.currentCurrentUserItem.item.Id,
                        data: this.currentCurrentUserItem.item,
                    }).then((response) => {

                        if (response.Model) {
                            this.updateCurrentUserById(response.Model);
                            this.currentCurrentUserItem.hasChange = false;
                            this.clearPanelCurrentUserItem();
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

            deleteCurrentUserItemHandler() {
                let deletedItemId = this.currentCurrentUserItem.item.Id;

                if (!this.currentCurrentUserItem.canDelete) {
                    this.currentCurrentUserItem.showDeleteConfirmation = true;
                    return;
                }

                this.deleteCurrentUser({
                    id: deletedItemId
                }).then(response => {

                    if (response.IsSuccess) {
                        this.deleteCurrentUserFromList(deletedItemId);
                        this.clearPanelCurrentUserItem();
                        this.currentCurrentUserItem.canDelete = false;
                        this.currentCurrentUserItem.isSelected = false;
                        this.panel.show = false;
                    } else {
                        console.error('Ошибка удаления элемента: ', response.Error);
                    }

                }).catch(error => {
                    console.error('Ошибка удаления элемента: ', error);
                });
            },

            confirmDeleteHandler() {
                this.currentCurrentUserItem.showDeleteConfirmation = false;
                this.currentCurrentUserItem.canDelete = true;
                this.deleteCurrentUserItemHandler();
            },

            closeConfirmationPanel() {
                this.currentCurrentUserItem.showDeleteConfirmation = false;
            },
        },
    }
</script>

<style lang="scss">

</style>
