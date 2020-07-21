
	<template>
    <WorkSpace>
        <template #header>
            <slot name="pageHeader">
                <VHead level="h1">ProjectInfoFilter</VHead>
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
                            v-for="projectInfoFilterItem in projectInfoFilterList"
                            :key="projectInfoFilterItem.Id"
                            @click="selectProjectInfoFilterItem(projectInfoFilterItem)"
                            class="sw-table__row_can-select"
                            :class="{'sw-table__row_is-selected': projectInfoFilterItem.Id === currentProjectInfoFilterItem.item.Id}"
                        >
                            <td v-for="(value, key) in fields" :key="key + '-fields'">
                                <VCheckbox v-if="isCheckbox(projectInfoFilterItem[key])" :checked="projectInfoFilterItem[key]" disabled></VCheckbox>
                                <VText v-else>{{ projectInfoFilterItem[key] }}</VText>
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
                            <VSet vertical>
                                <VSet
                                    v-for="(filed, key) in editFields" :key="key + '-editFields'"
                                    vertical-align="center"
                                >
                                    <VLabel
                                        width="col4"
                                        :for="`currentProjectInfoFilterItem${key}`"
                                    >{{ filed }}</VLabel>
                                    <VInput
										v-if="isInput(currentProjectInfoFilterItem.item[key])"
                                        v-model="currentProjectInfoFilterItem.item[key]"
                                        width="dyn"
                                        :id="`currentProjectInfoFilterItem${key}`"
                                        @input="changeCurrentProjectInfoFilterItem"
                                    />
									<VCheckbox
										v-if="isCheckbox(currentProjectInfoFilterItem.item[key])"
                                        v-model="currentProjectInfoFilterItem.item[key]"
                                        :id="`currentProjectInfoFilterItem${key}`"
										@input="changeCurrentProjectInfoFilterItem"
									/>
									
                                </VSet>
                            </VSet>
                            <button type="submit" :disabled="!currentProjectInfoFilterItem.hasChange" hidden></button>
                        </form>
                    </template>

                    <template #footer>
                        <VSet>
                            <VButton
                                @click="saveChangesSubmit"
                                accent
                                :text="panelSubmitButtonText"
                                :disabled="!currentProjectInfoFilterItem.hasChange"
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
                    v-if="currentProjectInfoFilterItem.showDeleteConfirmation"
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
                        :disabled="!currentProjectInfoFilterItem.isSelected"
                        @click="deleteProjectInfoFilterItemHandler"
                    />
                </VSet>
            </slot>
        </template>
    </WorkSpace>
</template>

<script>
    import projectInfoFilterData from "../data/ProjectInfoFilterData";
    import { ProjectInfoFilter } from '../apiModel';
    import { mapGetters, mapMutations, mapActions } from 'vuex';
    import WorkSpace from "swtui/src/components/WorkSpace";
    import VHead from "swtui/src/components/VHead";
    import VSet from "swtui/src/components/VSet";
    import VLabel from "swtui/src/components/VLabel";
    import VInput from "swtui/src/components/VInput";
    import VCheckbox from "swtui/src/components/VCheckbox";
    import VText from "swtui/src/components/VText";
    import VPanel from "swtui/src/components/VPanel";
    import VButton from "swtui/src/components/VButton";

    export default {
        name: 'ProjectInfoFilterGen',

        components: {VButton, VPanel, VText, VInput, VLabel, VSet, VHead, WorkSpace, VCheckbox},

        props: {
            fields: {
                type: Object,
                default() {
                    const projectInfoFilterItem = new ProjectInfoFilter();
                    const fieldsObj = {};

                    for (let prop in projectInfoFilterItem) {

                        if (projectInfoFilterItem.hasOwnProperty(prop)) {
                            fieldsObj[prop] = prop;
                        }

                    }

                    return fieldsObj;
                }
            },
            editFields: {
                type: Object,
                default() {
                    const projectInfoFilterItem = new ProjectInfoFilter();
                    const fieldsObj = {};

                    for (let prop in projectInfoFilterItem) {

                        if (projectInfoFilterItem.hasOwnProperty(prop)) {
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
            return projectInfoFilterData;
        },

        created() {
			this.onCreated();
        },

        computed: {
            ...mapGetters('gosha', {
                projectInfoFilterList: 'getListProjectInfoFilter'
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
            ...mapActions('gosha', [
                'findProjectInfoFilter',
                'updateProjectInfoFilter',
                'deleteProjectInfoFilter',
                'createProjectInfoFilter',
            ]),

            ...mapMutations('gosha', [
                'addProjectInfoFilterItemToList',
                'deleteProjectInfoFilterFromList',
                'updateProjectInfoFilterById',
            ]),

			onCreated() {
				this.fillProjectInfoFilterFilter();
	            this.fetchProjectInfoFilterData();
			},

            fillProjectInfoFilterFilter() {
                this.projectInfoFilterFilter.CurrentPage = 1;
                this.projectInfoFilterFilter.PerPage = 1000;
            },

            fetchProjectInfoFilterData() {
                return this.findProjectInfoFilter({
                    filter: this.projectInfoFilterFilter
                });
            },

            /**
             *
             * @param type
             */
            showPanel(type) {
                if (type === this.panel.create) {
                    this.panel.type = this.panel.create;
                    this.clearPanelProjectInfoFilterItem();
                } else if (type === this.panel.edit) {
                    this.panel.type = this.panel.edit;
                    this.currentProjectInfoFilterItem.isSelected = true;
                }

                this.panel.show = true;
            },

            closePanel() {
                this.panel.show = false;
                this.currentProjectInfoFilterItem.isSelected = false;
                this.clearPanelProjectInfoFilterItem();
            },

            selectProjectInfoFilterItem(projectInfoFilterItem) {
                this.showPanel(this.panel.edit);
                this.currentProjectInfoFilterItem.isSelected = true;
                Object.assign(this.currentProjectInfoFilterItem.item, projectInfoFilterItem);
            },

            changeCurrentProjectInfoFilterItem() {
                this.currentProjectInfoFilterItem.hasChange = true;
            },

            cancelChanges() {
                this.clearPanelProjectInfoFilterItem();
                this.closePanel();
            },

            clearPanelProjectInfoFilterItem() {
                this.currentProjectInfoFilterItem.item = new ProjectInfoFilter();
                this.currentProjectInfoFilterItem.hasChange = false;
            },

            saveChangesSubmit() {
                if (this.isPanelCreate) {
                    this.createProjectInfoFilterItemSubmit();
                    return;
                }

                if (this.isPanelEdit) {
                    this.editProjectInfoFilterItemSubmit();
                }
            },

            createProjectInfoFilterItemSubmit() {
                return this.createProjectInfoFilter({
					data: this.currentProjectInfoFilterItem.item,
                }).then((response) => {

                    if (response.Model) {
                        this.addProjectInfoFilterItemToList(response.Model);
                        this.clearPanelProjectInfoFilterItem();
                    } else {
                        console.error('Ошибка создания записи: ', response.Error);
                    }

                }).catch(error => {
                    console.error('Ошибка создания записи: ', error);
                });
            },

            editProjectInfoFilterItemSubmit() {

                if (this.currentProjectInfoFilterItem.hasChange) {
                    return this.updateProjectInfoFilter({
                        id: this.currentProjectInfoFilterItem.item.Id,
                        data: this.currentProjectInfoFilterItem.item,
                    }).then((response) => {

                        if (response.Model) {
                            this.updateProjectInfoFilterById(response.Model);
                            this.currentProjectInfoFilterItem.hasChange = false;
                            this.clearPanelProjectInfoFilterItem();
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

            deleteProjectInfoFilterItemHandler() {
                let deletedItemId = this.currentProjectInfoFilterItem.item.Id;

                if (!this.currentProjectInfoFilterItem.canDelete) {
                    this.currentProjectInfoFilterItem.showDeleteConfirmation = true;
                    return;
                }

                this.deleteProjectInfoFilter({
                    id: deletedItemId
                }).then(response => {

                    if (response.IsSuccess) {
                        this.deleteProjectInfoFilterFromList(deletedItemId);
                        this.clearPanelProjectInfoFilterItem();
                        this.currentProjectInfoFilterItem.canDelete = false;
                        this.currentProjectInfoFilterItem.isSelected = false;
                        this.panel.show = false;
                    } else {
                        console.error('Ошибка удаления элемента: ', response.Error);
                    }

                }).catch(error => {
                    console.error('Ошибка удаления элемента: ', error);
                });
            },

            confirmDeleteHandler() {
                this.currentProjectInfoFilterItem.showDeleteConfirmation = false;
                this.currentProjectInfoFilterItem.canDelete = true;
                this.deleteProjectInfoFilterItemHandler();
            },

            closeConfirmationPanel() {
                this.currentProjectInfoFilterItem.showDeleteConfirmation = false;
            },
        },
    }
</script>

<style lang="scss">

</style>
