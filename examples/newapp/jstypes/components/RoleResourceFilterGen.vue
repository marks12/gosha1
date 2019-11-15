
	<template>
    <WorkSpace>
        <template #header>
            <slot name="pageHeader">
                <VHead level="h1">RoleResourceFilter</VHead>
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
                            v-for="roleResourceFilterItem in roleResourceFilterList"
                            :key="roleResourceFilterItem.Id"
                            @click="selectRoleResourceFilterItem(roleResourceFilterItem)"
                            class="sw-table__row_can-select"
                            :class="{'sw-table__row_is-selected': roleResourceFilterItem.Id === currentRoleResourceFilterItem.item.Id}"
                        >
                            <td v-for="(value, key) in fields" :key="key + '-fields'">
                                <VCheckbox v-if="isCheckbox(roleResourceFilterItem[key])" :checked="roleResourceFilterItem[key]" disabled></VCheckbox>
                                <VText v-else>{{ roleResourceFilterItem[key] }}</VText>
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
                                        :for="`currentRoleResourceFilterItem${key}`"
                                    >{{ filed }}</VLabel>
                                    <VInput
										v-if="isInput(currentRoleResourceFilterItem.item[key])"
                                        v-model="currentRoleResourceFilterItem.item[key]"
                                        width="dyn"
                                        :id="`currentRoleResourceFilterItem${key}`"
                                        @input="changeCurrentRoleResourceFilterItem"
                                    />
									<VCheckbox
										v-if="isCheckbox(currentRoleResourceFilterItem.item[key])"
                                        v-model="currentRoleResourceFilterItem.item[key]"
                                        :id="`currentRoleResourceFilterItem${key}`"
										@input="changeCurrentRoleResourceFilterItem"
									/>
									
                                </VSet>
                            </VSet>
                            <button type="submit" :disabled="!currentRoleResourceFilterItem.hasChange" hidden></button>
                        </form>
                    </template>

                    <template #footer>
                        <VSet>
                            <VButton
                                @click="saveChangesSubmit"
                                accent
                                :text="panelSubmitButtonText"
                                :disabled="!currentRoleResourceFilterItem.hasChange"
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
                    v-if="currentRoleResourceFilterItem.showDeleteConfirmation"
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
                        :disabled="!currentRoleResourceFilterItem.isSelected"
                        @click="deleteRoleResourceFilterItemHandler"
                    />
                </VSet>
            </slot>
        </template>
    </WorkSpace>
</template>

<script>
    import roleResourceFilterData from "../data/RoleResourceFilterData";
    import { RoleResourceFilter } from '../apiModel';
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
        name: 'RoleResourceFilterGen',

        components: {VSelectSimple, VSign, VIcon, VButton, VPanel, VText, VInput, VLabel, VSet, VHead, WorkSpace, VCheckbox},

        props: {
            fields: {
                type: Object,
                default() {
                    const roleResourceFilterItem = new RoleResourceFilter();
                    const fieldsObj = {};

                    for (let prop in roleResourceFilterItem) {

                        if (roleResourceFilterItem.hasOwnProperty(prop)) {
                            fieldsObj[prop] = prop;
                        }

                    }

                    return fieldsObj;
                }
            },
            editFields: {
                type: Object,
                default() {
                    const roleResourceFilterItem = new RoleResourceFilter();
                    const fieldsObj = {};

                    for (let prop in roleResourceFilterItem) {

                        if (roleResourceFilterItem.hasOwnProperty(prop)) {
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
            return roleResourceFilterData;
        },

        created() {
            this.fillRoleResourceFilterFilter();
            this.fetchRoleResourceFilterData();
        },

        computed: {
            ...mapGetters({
                roleResourceFilterList: 'getListRoleResourceFilter'
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
                'findRoleResourceFilter',
                'updateRoleResourceFilter',
                'deleteRoleResourceFilter',
                'createRoleResourceFilter',
            ]),

            ...mapMutations([
                'addRoleResourceFilterItemToList',
                'deleteRoleResourceFilterFromList',
                'updateRoleResourceFilterById',
            ]),

            fillRoleResourceFilterFilter() {
                this.roleResourceFilterFilter.CurrentPage = 1;
                this.roleResourceFilterFilter.PerPage = 1000;
            },

            fetchRoleResourceFilterData() {
                return this.findRoleResourceFilter({
                    filter: this.roleResourceFilterFilter
                });
            },

            /**
             *
             * @param type
             */
            showPanel(type) {
                if (type === this.panel.create) {
                    this.panel.type = this.panel.create;
                    this.clearPanelRoleResourceFilterItem();
                } else if (type === this.panel.edit) {
                    this.panel.type = this.panel.edit;
                    this.currentRoleResourceFilterItem.isSelected = true;
                }

                this.panel.show = true;
            },

            closePanel() {
                this.panel.show = false;
                this.currentRoleResourceFilterItem.isSelected = false;
                this.clearPanelRoleResourceFilterItem();
            },

            selectRoleResourceFilterItem(roleResourceFilterItem) {
                this.showPanel(this.panel.edit);
                this.currentRoleResourceFilterItem.isSelected = true;
                Object.assign(this.currentRoleResourceFilterItem.item, roleResourceFilterItem);
            },

            changeCurrentRoleResourceFilterItem() {
                this.currentRoleResourceFilterItem.hasChange = true;
            },

            cancelChanges() {
                this.clearPanelRoleResourceFilterItem();
                this.closePanel();
            },

            clearPanelRoleResourceFilterItem() {
                this.currentRoleResourceFilterItem.item = new RoleResourceFilter();
                this.currentRoleResourceFilterItem.hasChange = false;
            },

            saveChangesSubmit() {
                if (this.isPanelCreate) {
                    this.createRoleResourceFilterItemSubmit();
                    return;
                }

                if (this.isPanelEdit) {
                    this.editRoleResourceFilterItemSubmit();
                }
            },

            createRoleResourceFilterItemSubmit() {
                return this.createRoleResourceFilter({
					data: this.currentRoleResourceFilterItem.item,
                }).then((response) => {

                    if (response.Model) {
                        this.addRoleResourceFilterItemToList(response.Model);
                        this.clearPanelRoleResourceFilterItem();
                    } else {
                        console.error('Ошибка создания записи: ', response.Error);
                    }

                }).catch(error => {
                    console.error('Ошибка создания записи: ', error);
                });
            },

            editRoleResourceFilterItemSubmit() {

                if (this.currentRoleResourceFilterItem.hasChange) {
                    return this.updateRoleResourceFilter({
                        id: this.currentRoleResourceFilterItem.item.Id,
                        data: this.currentRoleResourceFilterItem.item,
                    }).then((response) => {

                        if (response.Model) {
                            this.updateRoleResourceFilterById(response.Model);
                            this.currentRoleResourceFilterItem.hasChange = false;
                            this.clearPanelRoleResourceFilterItem();
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

            deleteRoleResourceFilterItemHandler() {
                let deletedItemId = this.currentRoleResourceFilterItem.item.Id;

                if (!this.currentRoleResourceFilterItem.canDelete) {
                    this.currentRoleResourceFilterItem.showDeleteConfirmation = true;
                    return;
                }

                this.deleteRoleResourceFilter({
                    id: deletedItemId
                }).then(response => {

                    if (response.IsSuccess) {
                        this.deleteRoleResourceFilterFromList(deletedItemId);
                        this.clearPanelRoleResourceFilterItem();
                        this.currentRoleResourceFilterItem.canDelete = false;
                        this.currentRoleResourceFilterItem.isSelected = false;
                        this.panel.show = false;
                    } else {
                        console.error('Ошибка удаления элемента: ', response.Error);
                    }

                }).catch(error => {
                    console.error('Ошибка удаления элемента: ', error);
                });
            },

            confirmDeleteHandler() {
                this.currentRoleResourceFilterItem.showDeleteConfirmation = false;
                this.currentRoleResourceFilterItem.canDelete = true;
                this.deleteRoleResourceFilterItemHandler();
            },

            closeConfirmationPanel() {
                this.currentRoleResourceFilterItem.showDeleteConfirmation = false;
            },
        },
    }
</script>

<style lang="scss">

</style>
