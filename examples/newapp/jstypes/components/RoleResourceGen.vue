
	<template>
    <WorkSpace>
        <template #header>
            <slot name="pageHeader">
                <VHead level="h1">RoleResource</VHead>
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
                            v-for="roleResourceItem in roleResourceList"
                            :key="roleResourceItem.Id"
                            @click="selectRoleResourceItem(roleResourceItem)"
                            class="sw-table__row_can-select"
                            :class="{'sw-table__row_is-selected': roleResourceItem.Id === currentRoleResourceItem.item.Id}"
                        >
                            <td v-for="(value, key) in fields" :key="key + '-fields'">
                                <VCheckbox v-if="isCheckbox(roleResourceItem[key])" :checked="roleResourceItem[key]" disabled></VCheckbox>
                                <VText v-else>{{ roleResourceItem[key] }}</VText>
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
                                        :for="`currentRoleResourceItem${key}`"
                                    >{{ filed }}</VLabel>
                                    <VInput
										v-if="isInput(currentRoleResourceItem.item[key])"
                                        v-model="currentRoleResourceItem.item[key]"
                                        width="dyn"
                                        :id="`currentRoleResourceItem${key}`"
                                        @input="changeCurrentRoleResourceItem"
                                    />
									<VCheckbox
										v-if="isCheckbox(currentRoleResourceItem.item[key])"
                                        v-model="currentRoleResourceItem.item[key]"
                                        :id="`currentRoleResourceItem${key}`"
										@input="changeCurrentRoleResourceItem"
									/>
									
                                </VSet>
                            </VSet>
                            <button type="submit" :disabled="!currentRoleResourceItem.hasChange" hidden></button>
                        </form>
                    </template>

                    <template #footer>
                        <VSet>
                            <VButton
                                @click="saveChangesSubmit"
                                accent
                                :text="panelSubmitButtonText"
                                :disabled="!currentRoleResourceItem.hasChange"
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
                    v-if="currentRoleResourceItem.showDeleteConfirmation"
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
                        :disabled="!currentRoleResourceItem.isSelected"
                        @click="deleteRoleResourceItemHandler"
                    />
                </VSet>
            </slot>
        </template>
    </WorkSpace>
</template>

<script>
    import roleResourceData from "../data/RoleResourceData";
    import { RoleResource } from '../apiModel';
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
        name: 'RoleResourceGen',

        components: {VSelectSimple, VSign, VIcon, VButton, VPanel, VText, VInput, VLabel, VSet, VHead, WorkSpace, VCheckbox},

        props: {
            fields: {
                type: Object,
                default() {
                    const roleResourceItem = new RoleResource();
                    const fieldsObj = {};

                    for (let prop in roleResourceItem) {

                        if (roleResourceItem.hasOwnProperty(prop)) {
                            fieldsObj[prop] = prop;
                        }

                    }

                    return fieldsObj;
                }
            },
            editFields: {
                type: Object,
                default() {
                    const roleResourceItem = new RoleResource();
                    const fieldsObj = {};

                    for (let prop in roleResourceItem) {

                        if (roleResourceItem.hasOwnProperty(prop)) {
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
            return roleResourceData;
        },

        created() {
            this.fillRoleResourceFilter();
            this.fetchRoleResourceData();
        },

        computed: {
            ...mapGetters({
                roleResourceList: 'getListRoleResource'
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
                'findRoleResource',
                'updateRoleResource',
                'deleteRoleResource',
                'createRoleResource',
            ]),

            ...mapMutations([
                'addRoleResourceItemToList',
                'deleteRoleResourceFromList',
                'updateRoleResourceById',
            ]),

            fillRoleResourceFilter() {
                this.roleResourceFilter.CurrentPage = 1;
                this.roleResourceFilter.PerPage = 1000;
            },

            fetchRoleResourceData() {
                return this.findRoleResource({
                    filter: this.roleResourceFilter
                });
            },

            /**
             *
             * @param type
             */
            showPanel(type) {
                if (type === this.panel.create) {
                    this.panel.type = this.panel.create;
                    this.clearPanelRoleResourceItem();
                } else if (type === this.panel.edit) {
                    this.panel.type = this.panel.edit;
                    this.currentRoleResourceItem.isSelected = true;
                }

                this.panel.show = true;
            },

            closePanel() {
                this.panel.show = false;
                this.currentRoleResourceItem.isSelected = false;
                this.clearPanelRoleResourceItem();
            },

            selectRoleResourceItem(roleResourceItem) {
                this.showPanel(this.panel.edit);
                this.currentRoleResourceItem.isSelected = true;
                Object.assign(this.currentRoleResourceItem.item, roleResourceItem);
            },

            changeCurrentRoleResourceItem() {
                this.currentRoleResourceItem.hasChange = true;
            },

            cancelChanges() {
                this.clearPanelRoleResourceItem();
                this.closePanel();
            },

            clearPanelRoleResourceItem() {
                this.currentRoleResourceItem.item = new RoleResource();
                this.currentRoleResourceItem.hasChange = false;
            },

            saveChangesSubmit() {
                if (this.isPanelCreate) {
                    this.createRoleResourceItemSubmit();
                    return;
                }

                if (this.isPanelEdit) {
                    this.editRoleResourceItemSubmit();
                }
            },

            createRoleResourceItemSubmit() {
                return this.createRoleResource({
					data: this.currentRoleResourceItem.item,
                }).then((response) => {

                    if (response.Model) {
                        this.addRoleResourceItemToList(response.Model);
                        this.clearPanelRoleResourceItem();
                    } else {
                        console.error('Ошибка создания записи: ', response.Error);
                    }

                }).catch(error => {
                    console.error('Ошибка создания записи: ', error);
                });
            },

            editRoleResourceItemSubmit() {

                if (this.currentRoleResourceItem.hasChange) {
                    return this.updateRoleResource({
                        id: this.currentRoleResourceItem.item.Id,
                        data: this.currentRoleResourceItem.item,
                    }).then((response) => {

                        if (response.Model) {
                            this.updateRoleResourceById(response.Model);
                            this.currentRoleResourceItem.hasChange = false;
                            this.clearPanelRoleResourceItem();
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

            deleteRoleResourceItemHandler() {
                let deletedItemId = this.currentRoleResourceItem.item.Id;

                if (!this.currentRoleResourceItem.canDelete) {
                    this.currentRoleResourceItem.showDeleteConfirmation = true;
                    return;
                }

                this.deleteRoleResource({
                    id: deletedItemId
                }).then(response => {

                    if (response.IsSuccess) {
                        this.deleteRoleResourceFromList(deletedItemId);
                        this.clearPanelRoleResourceItem();
                        this.currentRoleResourceItem.canDelete = false;
                        this.currentRoleResourceItem.isSelected = false;
                        this.panel.show = false;
                    } else {
                        console.error('Ошибка удаления элемента: ', response.Error);
                    }

                }).catch(error => {
                    console.error('Ошибка удаления элемента: ', error);
                });
            },

            confirmDeleteHandler() {
                this.currentRoleResourceItem.showDeleteConfirmation = false;
                this.currentRoleResourceItem.canDelete = true;
                this.deleteRoleResourceItemHandler();
            },

            closeConfirmationPanel() {
                this.currentRoleResourceItem.showDeleteConfirmation = false;
            },
        },
    }
</script>

<style lang="scss">

</style>
