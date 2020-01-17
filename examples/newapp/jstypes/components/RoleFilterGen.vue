
	<template>
    <WorkSpace>
        <template #header>
            <slot name="pageHeader">
                <VHead level="h1">RoleFilter</VHead>
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
                            v-for="roleFilterItem in roleFilterList"
                            :key="roleFilterItem.Id"
                            @click="selectRoleFilterItem(roleFilterItem)"
                            class="sw-table__row_can-select"
                            :class="{'sw-table__row_is-selected': roleFilterItem.Id === currentRoleFilterItem.item.Id}"
                        >
                            <td v-for="(value, key) in fields" :key="key + '-fields'">
                                <VCheckbox v-if="isCheckbox(roleFilterItem[key])" :checked="roleFilterItem[key]" disabled></VCheckbox>
                                <VText v-else>{{ roleFilterItem[key] }}</VText>
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
                                        :for="`currentRoleFilterItem${key}`"
                                    >{{ filed }}</VLabel>
                                    <VInput
										v-if="isInput(currentRoleFilterItem.item[key])"
                                        v-model="currentRoleFilterItem.item[key]"
                                        width="dyn"
                                        :id="`currentRoleFilterItem${key}`"
                                        @input="changeCurrentRoleFilterItem"
                                    />
									<VCheckbox
										v-if="isCheckbox(currentRoleFilterItem.item[key])"
                                        v-model="currentRoleFilterItem.item[key]"
                                        :id="`currentRoleFilterItem${key}`"
										@input="changeCurrentRoleFilterItem"
									/>
									
                                </VSet>
                            </VSet>
                            <button type="submit" :disabled="!currentRoleFilterItem.hasChange" hidden></button>
                        </form>
                    </template>

                    <template #footer>
                        <VSet>
                            <VButton
                                @click="saveChangesSubmit"
                                accent
                                :text="panelSubmitButtonText"
                                :disabled="!currentRoleFilterItem.hasChange"
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
                    v-if="currentRoleFilterItem.showDeleteConfirmation"
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
                        :disabled="!currentRoleFilterItem.isSelected"
                        @click="deleteRoleFilterItemHandler"
                    />
                </VSet>
            </slot>
        </template>
    </WorkSpace>
</template>

<script>
    import roleFilterData from "../data/RoleFilterData";
    import { RoleFilter } from '../apiModel';
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
        name: 'RoleFilterGen',

        components: {VButton, VPanel, VText, VInput, VLabel, VSet, VHead, WorkSpace, VCheckbox},

        props: {
            fields: {
                type: Object,
                default() {
                    const roleFilterItem = new RoleFilter();
                    const fieldsObj = {};

                    for (let prop in roleFilterItem) {

                        if (roleFilterItem.hasOwnProperty(prop)) {
                            fieldsObj[prop] = prop;
                        }

                    }

                    return fieldsObj;
                }
            },
            editFields: {
                type: Object,
                default() {
                    const roleFilterItem = new RoleFilter();
                    const fieldsObj = {};

                    for (let prop in roleFilterItem) {

                        if (roleFilterItem.hasOwnProperty(prop)) {
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
            return roleFilterData;
        },

        created() {
			this.onCreated();
        },

        computed: {
            ...mapGetters({
                roleFilterList: 'getListRoleFilter'
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
                'findRoleFilter',
                'updateRoleFilter',
                'deleteRoleFilter',
                'createRoleFilter',
            ]),

            ...mapMutations([
                'addRoleFilterItemToList',
                'deleteRoleFilterFromList',
                'updateRoleFilterById',
            ]),

			onCreated() {
				this.fillRoleFilterFilter();
	            this.fetchRoleFilterData();
			},

            fillRoleFilterFilter() {
                this.roleFilterFilter.CurrentPage = 1;
                this.roleFilterFilter.PerPage = 1000;
            },

            fetchRoleFilterData() {
                return this.findRoleFilter({
                    filter: this.roleFilterFilter
                });
            },

            /**
             *
             * @param type
             */
            showPanel(type) {
                if (type === this.panel.create) {
                    this.panel.type = this.panel.create;
                    this.clearPanelRoleFilterItem();
                } else if (type === this.panel.edit) {
                    this.panel.type = this.panel.edit;
                    this.currentRoleFilterItem.isSelected = true;
                }

                this.panel.show = true;
            },

            closePanel() {
                this.panel.show = false;
                this.currentRoleFilterItem.isSelected = false;
                this.clearPanelRoleFilterItem();
            },

            selectRoleFilterItem(roleFilterItem) {
                this.showPanel(this.panel.edit);
                this.currentRoleFilterItem.isSelected = true;
                Object.assign(this.currentRoleFilterItem.item, roleFilterItem);
            },

            changeCurrentRoleFilterItem() {
                this.currentRoleFilterItem.hasChange = true;
            },

            cancelChanges() {
                this.clearPanelRoleFilterItem();
                this.closePanel();
            },

            clearPanelRoleFilterItem() {
                this.currentRoleFilterItem.item = new RoleFilter();
                this.currentRoleFilterItem.hasChange = false;
            },

            saveChangesSubmit() {
                if (this.isPanelCreate) {
                    this.createRoleFilterItemSubmit();
                    return;
                }

                if (this.isPanelEdit) {
                    this.editRoleFilterItemSubmit();
                }
            },

            createRoleFilterItemSubmit() {
                return this.createRoleFilter({
					data: this.currentRoleFilterItem.item,
                }).then((response) => {

                    if (response.Model) {
                        this.addRoleFilterItemToList(response.Model);
                        this.clearPanelRoleFilterItem();
                    } else {
                        console.error('Ошибка создания записи: ', response.Error);
                    }

                }).catch(error => {
                    console.error('Ошибка создания записи: ', error);
                });
            },

            editRoleFilterItemSubmit() {

                if (this.currentRoleFilterItem.hasChange) {
                    return this.updateRoleFilter({
                        id: this.currentRoleFilterItem.item.Id,
                        data: this.currentRoleFilterItem.item,
                    }).then((response) => {

                        if (response.Model) {
                            this.updateRoleFilterById(response.Model);
                            this.currentRoleFilterItem.hasChange = false;
                            this.clearPanelRoleFilterItem();
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

            deleteRoleFilterItemHandler() {
                let deletedItemId = this.currentRoleFilterItem.item.Id;

                if (!this.currentRoleFilterItem.canDelete) {
                    this.currentRoleFilterItem.showDeleteConfirmation = true;
                    return;
                }

                this.deleteRoleFilter({
                    id: deletedItemId
                }).then(response => {

                    if (response.IsSuccess) {
                        this.deleteRoleFilterFromList(deletedItemId);
                        this.clearPanelRoleFilterItem();
                        this.currentRoleFilterItem.canDelete = false;
                        this.currentRoleFilterItem.isSelected = false;
                        this.panel.show = false;
                    } else {
                        console.error('Ошибка удаления элемента: ', response.Error);
                    }

                }).catch(error => {
                    console.error('Ошибка удаления элемента: ', error);
                });
            },

            confirmDeleteHandler() {
                this.currentRoleFilterItem.showDeleteConfirmation = false;
                this.currentRoleFilterItem.canDelete = true;
                this.deleteRoleFilterItemHandler();
            },

            closeConfirmationPanel() {
                this.currentRoleFilterItem.showDeleteConfirmation = false;
            },
        },
    }
</script>

<style lang="scss">

</style>
