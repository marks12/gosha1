
	<template>
    <WorkSpace>
        <template #header>
            <slot name="pageHeader">
                <VHead level="h1">CurrentUserFilter</VHead>
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
                            v-for="currentUserFilterItem in currentUserFilterList"
                            :key="currentUserFilterItem.Id"
                            @click="selectCurrentUserFilterItem(currentUserFilterItem)"
                            class="sw-table__row_can-select"
                            :class="{'sw-table__row_is-selected': currentUserFilterItem.Id === currentCurrentUserFilterItem.item.Id}"
                        >
                            <td v-for="(value, key) in fields" :key="key + '-fields'">
                                <VCheckbox v-if="isCheckbox(currentUserFilterItem[key])" :checked="currentUserFilterItem[key]" disabled></VCheckbox>
                                <VText v-else>{{ currentUserFilterItem[key] }}</VText>
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
                                        :for="`currentCurrentUserFilterItem${key}`"
                                    >{{ filed }}</VLabel>
                                    <VInput
										v-if="isInput(currentCurrentUserFilterItem.item[key])"
                                        v-model="currentCurrentUserFilterItem.item[key]"
                                        width="dyn"
                                        :id="`currentCurrentUserFilterItem${key}`"
                                        @input="changeCurrentCurrentUserFilterItem"
                                    />
									<VCheckbox
										v-if="isCheckbox(currentCurrentUserFilterItem.item[key])"
                                        v-model="currentCurrentUserFilterItem.item[key]"
                                        :id="`currentCurrentUserFilterItem${key}`"
										@input="changeCurrentCurrentUserFilterItem"
									/>
									
                                </VSet>
                            </VSet>
                            <button type="submit" :disabled="!currentCurrentUserFilterItem.hasChange" hidden></button>
                        </form>
                    </template>

                    <template #footer>
                        <VSet>
                            <VButton
                                @click="saveChangesSubmit"
                                accent
                                :text="panelSubmitButtonText"
                                :disabled="!currentCurrentUserFilterItem.hasChange"
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
                    v-if="currentCurrentUserFilterItem.showDeleteConfirmation"
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
                        :disabled="!currentCurrentUserFilterItem.isSelected"
                        @click="deleteCurrentUserFilterItemHandler"
                    />
                </VSet>
            </slot>
        </template>
    </WorkSpace>
</template>

<script>
    import currentUserFilterData from "../data/CurrentUserFilterData";
    import { CurrentUserFilter } from '../apiModel';
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
        name: 'CurrentUserFilterGen',

        components: {VSelectSimple, VSign, VIcon, VButton, VPanel, VText, VInput, VLabel, VSet, VHead, WorkSpace, VCheckbox},

        props: {
            fields: {
                type: Object,
                default() {
                    const currentUserFilterItem = new CurrentUserFilter();
                    const fieldsObj = {};

                    for (let prop in currentUserFilterItem) {

                        if (currentUserFilterItem.hasOwnProperty(prop)) {
                            fieldsObj[prop] = prop;
                        }

                    }

                    return fieldsObj;
                }
            },
            editFields: {
                type: Object,
                default() {
                    const currentUserFilterItem = new CurrentUserFilter();
                    const fieldsObj = {};

                    for (let prop in currentUserFilterItem) {

                        if (currentUserFilterItem.hasOwnProperty(prop)) {
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
            return currentUserFilterData;
        },

        created() {
            this.fillCurrentUserFilterFilter();
            this.fetchCurrentUserFilterData();
        },

        computed: {
            ...mapGetters({
                currentUserFilterList: 'getListCurrentUserFilter'
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
                'findCurrentUserFilter',
                'updateCurrentUserFilter',
                'deleteCurrentUserFilter',
                'createCurrentUserFilter',
            ]),

            ...mapMutations([
                'addCurrentUserFilterItemToList',
                'deleteCurrentUserFilterFromList',
                'updateCurrentUserFilterById',
            ]),

            fillCurrentUserFilterFilter() {
                this.currentUserFilterFilter.CurrentPage = 1;
                this.currentUserFilterFilter.PerPage = 1000;
            },

            fetchCurrentUserFilterData() {
                return this.findCurrentUserFilter({
                    filter: this.currentUserFilterFilter
                });
            },

            /**
             *
             * @param type
             */
            showPanel(type) {
                if (type === this.panel.create) {
                    this.panel.type = this.panel.create;
                    this.clearPanelCurrentUserFilterItem();
                } else if (type === this.panel.edit) {
                    this.panel.type = this.panel.edit;
                    this.currentCurrentUserFilterItem.isSelected = true;
                }

                this.panel.show = true;
            },

            closePanel() {
                this.panel.show = false;
                this.currentCurrentUserFilterItem.isSelected = false;
                this.clearPanelCurrentUserFilterItem();
            },

            selectCurrentUserFilterItem(currentUserFilterItem) {
                this.showPanel(this.panel.edit);
                this.currentCurrentUserFilterItem.isSelected = true;
                Object.assign(this.currentCurrentUserFilterItem.item, currentUserFilterItem);
            },

            changeCurrentCurrentUserFilterItem() {
                this.currentCurrentUserFilterItem.hasChange = true;
            },

            cancelChanges() {
                this.clearPanelCurrentUserFilterItem();
                this.closePanel();
            },

            clearPanelCurrentUserFilterItem() {
                this.currentCurrentUserFilterItem.item = new CurrentUserFilter();
                this.currentCurrentUserFilterItem.hasChange = false;
            },

            saveChangesSubmit() {
                if (this.isPanelCreate) {
                    this.createCurrentUserFilterItemSubmit();
                    return;
                }

                if (this.isPanelEdit) {
                    this.editCurrentUserFilterItemSubmit();
                }
            },

            createCurrentUserFilterItemSubmit() {
                return this.createCurrentUserFilter({
					data: this.currentCurrentUserFilterItem.item,
                }).then((response) => {

                    if (response.Model) {
                        this.addCurrentUserFilterItemToList(response.Model);
                        this.clearPanelCurrentUserFilterItem();
                    } else {
                        console.error('Ошибка создания записи: ', response.Error);
                    }

                }).catch(error => {
                    console.error('Ошибка создания записи: ', error);
                });
            },

            editCurrentUserFilterItemSubmit() {

                if (this.currentCurrentUserFilterItem.hasChange) {
                    return this.updateCurrentUserFilter({
                        id: this.currentCurrentUserFilterItem.item.Id,
                        data: this.currentCurrentUserFilterItem.item,
                    }).then((response) => {

                        if (response.Model) {
                            this.updateCurrentUserFilterById(response.Model);
                            this.currentCurrentUserFilterItem.hasChange = false;
                            this.clearPanelCurrentUserFilterItem();
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

            deleteCurrentUserFilterItemHandler() {
                let deletedItemId = this.currentCurrentUserFilterItem.item.Id;

                if (!this.currentCurrentUserFilterItem.canDelete) {
                    this.currentCurrentUserFilterItem.showDeleteConfirmation = true;
                    return;
                }

                this.deleteCurrentUserFilter({
                    id: deletedItemId
                }).then(response => {

                    if (response.IsSuccess) {
                        this.deleteCurrentUserFilterFromList(deletedItemId);
                        this.clearPanelCurrentUserFilterItem();
                        this.currentCurrentUserFilterItem.canDelete = false;
                        this.currentCurrentUserFilterItem.isSelected = false;
                        this.panel.show = false;
                    } else {
                        console.error('Ошибка удаления элемента: ', response.Error);
                    }

                }).catch(error => {
                    console.error('Ошибка удаления элемента: ', error);
                });
            },

            confirmDeleteHandler() {
                this.currentCurrentUserFilterItem.showDeleteConfirmation = false;
                this.currentCurrentUserFilterItem.canDelete = true;
                this.deleteCurrentUserFilterItemHandler();
            },

            closeConfirmationPanel() {
                this.currentCurrentUserFilterItem.showDeleteConfirmation = false;
            },
        },
    }
</script>

<style lang="scss">

</style>
