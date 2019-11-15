
	<template>
    <WorkSpace>
        <template #header>
            <slot name="pageHeader">
                <VHead level="h1">Role</VHead>
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
                            v-for="roleItem in roleList"
                            :key="roleItem.Id"
                            @click="selectRoleItem(roleItem)"
                            class="sw-table__row_can-select"
                            :class="{'sw-table__row_is-selected': roleItem.Id === currentRoleItem.item.Id}"
                        >
                            <td v-for="(value, key) in fields" :key="key + '-fields'">
                                <VCheckbox v-if="isCheckbox(roleItem[key])" :checked="roleItem[key]" disabled></VCheckbox>
                                <VText v-else>{{ roleItem[key] }}</VText>
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
                                        :for="`currentRoleItem${key}`"
                                    >{{ filed }}</VLabel>
                                    <VInput
										v-if="isInput(currentRoleItem.item[key])"
                                        v-model="currentRoleItem.item[key]"
                                        width="dyn"
                                        :id="`currentRoleItem${key}`"
                                        @input="changeCurrentRoleItem"
                                    />
									<VCheckbox
										v-if="isCheckbox(currentRoleItem.item[key])"
                                        v-model="currentRoleItem.item[key]"
                                        :id="`currentRoleItem${key}`"
										@input="changeCurrentRoleItem"
									/>
									
                                </VSet>
                            </VSet>
                            <button type="submit" :disabled="!currentRoleItem.hasChange" hidden></button>
                        </form>
                    </template>

                    <template #footer>
                        <VSet>
                            <VButton
                                @click="saveChangesSubmit"
                                accent
                                :text="panelSubmitButtonText"
                                :disabled="!currentRoleItem.hasChange"
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
                    v-if="currentRoleItem.showDeleteConfirmation"
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
                        :disabled="!currentRoleItem.isSelected"
                        @click="deleteRoleItemHandler"
                    />
                </VSet>
            </slot>
        </template>
    </WorkSpace>
</template>

<script>
    import roleData from "../data/RoleData";
    import { Role } from '../apiModel';
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
        name: 'RoleGen',

        components: {VSelectSimple, VSign, VIcon, VButton, VPanel, VText, VInput, VLabel, VSet, VHead, WorkSpace, VCheckbox},

        props: {
            fields: {
                type: Object,
                default() {
                    const roleItem = new Role();
                    const fieldsObj = {};

                    for (let prop in roleItem) {

                        if (roleItem.hasOwnProperty(prop)) {
                            fieldsObj[prop] = prop;
                        }

                    }

                    return fieldsObj;
                }
            },
            editFields: {
                type: Object,
                default() {
                    const roleItem = new Role();
                    const fieldsObj = {};

                    for (let prop in roleItem) {

                        if (roleItem.hasOwnProperty(prop)) {
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
            return roleData;
        },

        created() {
            this.fillRoleFilter();
            this.fetchRoleData();
        },

        computed: {
            ...mapGetters({
                roleList: 'getListRole'
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
                'findRole',
                'updateRole',
                'deleteRole',
                'createRole',
            ]),

            ...mapMutations([
                'addRoleItemToList',
                'deleteRoleFromList',
                'updateRoleById',
            ]),

            fillRoleFilter() {
                this.roleFilter.CurrentPage = 1;
                this.roleFilter.PerPage = 1000;
            },

            fetchRoleData() {
                return this.findRole({
                    filter: this.roleFilter
                });
            },

            /**
             *
             * @param type
             */
            showPanel(type) {
                if (type === this.panel.create) {
                    this.panel.type = this.panel.create;
                    this.clearPanelRoleItem();
                } else if (type === this.panel.edit) {
                    this.panel.type = this.panel.edit;
                    this.currentRoleItem.isSelected = true;
                }

                this.panel.show = true;
            },

            closePanel() {
                this.panel.show = false;
                this.currentRoleItem.isSelected = false;
                this.clearPanelRoleItem();
            },

            selectRoleItem(roleItem) {
                this.showPanel(this.panel.edit);
                this.currentRoleItem.isSelected = true;
                Object.assign(this.currentRoleItem.item, roleItem);
            },

            changeCurrentRoleItem() {
                this.currentRoleItem.hasChange = true;
            },

            cancelChanges() {
                this.clearPanelRoleItem();
                this.closePanel();
            },

            clearPanelRoleItem() {
                this.currentRoleItem.item = new Role();
                this.currentRoleItem.hasChange = false;
            },

            saveChangesSubmit() {
                if (this.isPanelCreate) {
                    this.createRoleItemSubmit();
                    return;
                }

                if (this.isPanelEdit) {
                    this.editRoleItemSubmit();
                }
            },

            createRoleItemSubmit() {
                return this.createRole({
					data: this.currentRoleItem.item,
                }).then((response) => {

                    if (response.Model) {
                        this.addRoleItemToList(response.Model);
                        this.clearPanelRoleItem();
                    } else {
                        console.error('Ошибка создания записи: ', response.Error);
                    }

                }).catch(error => {
                    console.error('Ошибка создания записи: ', error);
                });
            },

            editRoleItemSubmit() {

                if (this.currentRoleItem.hasChange) {
                    return this.updateRole({
                        id: this.currentRoleItem.item.Id,
                        data: this.currentRoleItem.item,
                    }).then((response) => {

                        if (response.Model) {
                            this.updateRoleById(response.Model);
                            this.currentRoleItem.hasChange = false;
                            this.clearPanelRoleItem();
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

            deleteRoleItemHandler() {
                let deletedItemId = this.currentRoleItem.item.Id;

                if (!this.currentRoleItem.canDelete) {
                    this.currentRoleItem.showDeleteConfirmation = true;
                    return;
                }

                this.deleteRole({
                    id: deletedItemId
                }).then(response => {

                    if (response.IsSuccess) {
                        this.deleteRoleFromList(deletedItemId);
                        this.clearPanelRoleItem();
                        this.currentRoleItem.canDelete = false;
                        this.currentRoleItem.isSelected = false;
                        this.panel.show = false;
                    } else {
                        console.error('Ошибка удаления элемента: ', response.Error);
                    }

                }).catch(error => {
                    console.error('Ошибка удаления элемента: ', error);
                });
            },

            confirmDeleteHandler() {
                this.currentRoleItem.showDeleteConfirmation = false;
                this.currentRoleItem.canDelete = true;
                this.deleteRoleItemHandler();
            },

            closeConfirmationPanel() {
                this.currentRoleItem.showDeleteConfirmation = false;
            },
        },
    }
</script>

<style lang="scss">

</style>
