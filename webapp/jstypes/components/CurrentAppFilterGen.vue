
	<template>
    <WorkSpace>
        <template #header>
            <slot name="pageHeader">
                <VHead level="h1">CurrentAppFilter</VHead>
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
                            v-for="currentAppFilterItem in currentAppFilterList"
                            :key="currentAppFilterItem.Id"
                            @click="selectCurrentAppFilterItem(currentAppFilterItem)"
                            class="sw-table__row_can-select"
                            :class="{'sw-table__row_is-selected': currentAppFilterItem.Id === currentCurrentAppFilterItem.item.Id}"
                        >
                            <td v-for="(value, key) in fields" :key="key + '-fields'">
                                <VCheckbox v-if="isCheckbox(currentAppFilterItem[key])" :checked="currentAppFilterItem[key]" disabled></VCheckbox>
                                <VText v-else>{{ currentAppFilterItem[key] }}</VText>
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
                                        :for="`currentCurrentAppFilterItem${key}`"
                                    >{{ filed }}</VLabel>
                                    <VInput
										v-if="isInput(currentCurrentAppFilterItem.item[key])"
                                        v-model="currentCurrentAppFilterItem.item[key]"
                                        width="dyn"
                                        :id="`currentCurrentAppFilterItem${key}`"
                                        @input="changeCurrentCurrentAppFilterItem"
                                    />
									<VCheckbox
										v-if="isCheckbox(currentCurrentAppFilterItem.item[key])"
                                        v-model="currentCurrentAppFilterItem.item[key]"
                                        :id="`currentCurrentAppFilterItem${key}`"
										@input="changeCurrentCurrentAppFilterItem"
									/>
									
                                </VSet>
                            </VSet>
                            <button type="submit" :disabled="!currentCurrentAppFilterItem.hasChange" hidden></button>
                        </form>
                    </template>

                    <template #footer>
                        <VSet>
                            <VButton
                                @click="saveChangesSubmit"
                                accent
                                :text="panelSubmitButtonText"
                                :disabled="!currentCurrentAppFilterItem.hasChange"
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
                    v-if="currentCurrentAppFilterItem.showDeleteConfirmation"
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
                        :disabled="!currentCurrentAppFilterItem.isSelected"
                        @click="deleteCurrentAppFilterItemHandler"
                    />
                </VSet>
            </slot>
        </template>
    </WorkSpace>
</template>

<script>
    import currentAppFilterData from "../data/CurrentAppFilterData";
    import { CurrentAppFilter } from '../apiModel';
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
        name: 'CurrentAppFilterGen',

        components: {VSelectSimple, VSign, VIcon, VButton, VPanel, VText, VInput, VLabel, VSet, VHead, WorkSpace, VCheckbox},

        props: {
            fields: {
                type: Object,
                default() {
                    const currentAppFilterItem = new CurrentAppFilter();
                    const fieldsObj = {};

                    for (let prop in currentAppFilterItem) {

                        if (currentAppFilterItem.hasOwnProperty(prop)) {
                            fieldsObj[prop] = prop;
                        }

                    }

                    return fieldsObj;
                }
            },
            editFields: {
                type: Object,
                default() {
                    const currentAppFilterItem = new CurrentAppFilter();
                    const fieldsObj = {};

                    for (let prop in currentAppFilterItem) {

                        if (currentAppFilterItem.hasOwnProperty(prop)) {
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
            return currentAppFilterData;
        },

        created() {
            this.fillCurrentAppFilterFilter();
            this.fetchCurrentAppFilterData();
        },

        computed: {
            ...mapGetters({
                currentAppFilterList: 'getListCurrentAppFilter'
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
                'findCurrentAppFilter',
                'updateCurrentAppFilter',
                'deleteCurrentAppFilter',
                'createCurrentAppFilter',
            ]),

            ...mapMutations([
                'addCurrentAppFilterItemToList',
                'deleteCurrentAppFilterFromList',
                'updateCurrentAppFilterById',
            ]),

            fillCurrentAppFilterFilter() {
                this.currentAppFilterFilter.CurrentPage = 1;
                this.currentAppFilterFilter.PerPage = 1000;
            },

            fetchCurrentAppFilterData() {
                return this.findCurrentAppFilter({
                    filter: this.currentAppFilterFilter
                });
            },

            /**
             *
             * @param type
             */
            showPanel(type) {
                if (type === this.panel.create) {
                    this.panel.type = this.panel.create;
                    this.clearPanelCurrentAppFilterItem();
                } else if (type === this.panel.edit) {
                    this.panel.type = this.panel.edit;
                    this.currentCurrentAppFilterItem.isSelected = true;
                }

                this.panel.show = true;
            },

            closePanel() {
                this.panel.show = false;
                this.currentCurrentAppFilterItem.isSelected = false;
                this.clearPanelCurrentAppFilterItem();
            },

            selectCurrentAppFilterItem(currentAppFilterItem) {
                this.showPanel(this.panel.edit);
                this.currentCurrentAppFilterItem.isSelected = true;
                Object.assign(this.currentCurrentAppFilterItem.item, currentAppFilterItem);
            },

            changeCurrentCurrentAppFilterItem() {
                this.currentCurrentAppFilterItem.hasChange = true;
            },

            cancelChanges() {
                this.clearPanelCurrentAppFilterItem();
                this.closePanel();
            },

            clearPanelCurrentAppFilterItem() {
                this.currentCurrentAppFilterItem.item = new CurrentAppFilter();
                this.currentCurrentAppFilterItem.hasChange = false;
            },

            saveChangesSubmit() {
                if (this.isPanelCreate) {
                    this.createCurrentAppFilterItemSubmit();
                    return;
                }

                if (this.isPanelEdit) {
                    this.editCurrentAppFilterItemSubmit();
                }
            },

            createCurrentAppFilterItemSubmit() {
                this.createCurrentAppFilter({
					data: this.currentCurrentAppFilterItem.item,
                }).then((response) => {

                    if (response.Model) {
                        this.addCurrentAppFilterItemToList(response.Model);
                        this.clearPanelCurrentAppFilterItem();
                    } else {
                        console.error('Ошибка создания записи: ', response.Error);
                    }

                }).catch(error => {
                    console.error('Ошибка создания записи: ', error);
                });
            },

            editCurrentAppFilterItemSubmit() {
                if (this.currentCurrentAppFilterItem.hasChange) {
                    this.updateCurrentAppFilter({
                        id: this.currentCurrentAppFilterItem.item.Id,
                        data: this.currentCurrentAppFilterItem.item,
                    }).then((response) => {

                        if (response.Model) {
                            this.updateCurrentAppFilterById(response.Model);
                            this.currentCurrentAppFilterItem.hasChange = false;
                            this.clearPanelCurrentAppFilterItem();
                            this.closePanel();
                        } else {
                            console.error('Ошибка изменения записи: ', response.Error);
                        }

                    }).catch(error => {
                        console.error('Ошибка изменения записи: ', error);
                    });
                }
            },

            deleteCurrentAppFilterItemHandler() {
                let deletedItemId = this.currentCurrentAppFilterItem.item.Id;

                if (!this.currentCurrentAppFilterItem.canDelete) {
                    this.currentCurrentAppFilterItem.showDeleteConfirmation = true;
                    return;
                }

                this.deleteCurrentAppFilter({
                    id: deletedItemId
                }).then(response => {

                    if (response.IsSuccess) {
                        this.deleteCurrentAppFilterFromList(deletedItemId);
                        this.clearPanelCurrentAppFilterItem();
                        this.currentCurrentAppFilterItem.canDelete = false;
                        this.currentCurrentAppFilterItem.isSelected = false;
                        this.panel.show = false;
                    } else {
                        console.error('Ошибка удаления элемента: ', response.Error);
                    }

                }).catch(error => {
                    console.error('Ошибка удаления элемента: ', error);
                });
            },

            confirmDeleteHandler() {
                this.currentCurrentAppFilterItem.showDeleteConfirmation = false;
                this.currentCurrentAppFilterItem.canDelete = true;
                this.deleteCurrentAppFilterItemHandler();
            },

            closeConfirmationPanel() {
                this.currentCurrentAppFilterItem.showDeleteConfirmation = false;
            },
        },
    }
</script>

<style lang="scss">

</style>
