
	<template>
    <WorkSpace>
        <template #header>
            <slot name="pageHeader">
                <VHead level="h1">ProjectInfo</VHead>
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
                            v-for="projectInfoItem in projectInfoList"
                            :key="projectInfoItem.Id"
                            @click="selectProjectInfoItem(projectInfoItem)"
                            class="sw-table__row_can-select"
                            :class="{'sw-table__row_is-selected': projectInfoItem.Id === currentProjectInfoItem.item.Id}"
                        >
                            <td v-for="(value, key) in fields" :key="key + '-fields'">
                                <VCheckbox v-if="isCheckbox(projectInfoItem[key])" :checked="projectInfoItem[key]" disabled></VCheckbox>
                                <VText v-else>{{ projectInfoItem[key] }}</VText>
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
                                        :for="`currentProjectInfoItem${key}`"
                                    >{{ filed }}</VLabel>
                                    <VInput
										v-if="isInput(currentProjectInfoItem.item[key])"
                                        v-model="currentProjectInfoItem.item[key]"
                                        width="dyn"
                                        :id="`currentProjectInfoItem${key}`"
                                        @input="changeCurrentProjectInfoItem"
                                    />
									<VCheckbox
										v-if="isCheckbox(currentProjectInfoItem.item[key])"
                                        v-model="currentProjectInfoItem.item[key]"
                                        :id="`currentProjectInfoItem${key}`"
										@input="changeCurrentProjectInfoItem"
									/>
									
                                </VSet>
                            </VSet>
                            <button type="submit" :disabled="!currentProjectInfoItem.hasChange" hidden></button>
                        </form>
                    </template>

                    <template #footer>
                        <VSet>
                            <VButton
                                @click="saveChangesSubmit"
                                accent
                                :text="panelSubmitButtonText"
                                :disabled="!currentProjectInfoItem.hasChange"
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
                    v-if="currentProjectInfoItem.showDeleteConfirmation"
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
                        :disabled="!currentProjectInfoItem.isSelected"
                        @click="deleteProjectInfoItemHandler"
                    />
                </VSet>
            </slot>
        </template>
    </WorkSpace>
</template>

<script>
    import projectInfoData from "../data/ProjectInfoData";
    import { ProjectInfo } from '../apiModel';
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
        name: 'ProjectInfoGen',

        components: {VButton, VPanel, VText, VInput, VLabel, VSet, VHead, WorkSpace, VCheckbox},

        props: {
            fields: {
                type: Object,
                default() {
                    const projectInfoItem = new ProjectInfo();
                    const fieldsObj = {};

                    for (let prop in projectInfoItem) {

                        if (projectInfoItem.hasOwnProperty(prop)) {
                            fieldsObj[prop] = prop;
                        }

                    }

                    return fieldsObj;
                }
            },
            editFields: {
                type: Object,
                default() {
                    const projectInfoItem = new ProjectInfo();
                    const fieldsObj = {};

                    for (let prop in projectInfoItem) {

                        if (projectInfoItem.hasOwnProperty(prop)) {
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
            return projectInfoData;
        },

        created() {
			this.onCreated();
        },

        computed: {
            ...mapGetters('gosha', {
                projectInfoList: 'getListProjectInfo'
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
                'findProjectInfo',
                'updateProjectInfo',
                'deleteProjectInfo',
                'createProjectInfo',
            ]),

            ...mapMutations('gosha', [
                'addProjectInfoItemToList',
                'deleteProjectInfoFromList',
                'updateProjectInfoById',
            ]),

			onCreated() {
				this.fillProjectInfoFilter();
	            this.fetchProjectInfoData();
			},

            fillProjectInfoFilter() {
                this.projectInfoFilter.CurrentPage = 1;
                this.projectInfoFilter.PerPage = 1000;
            },

            fetchProjectInfoData() {
                return this.findProjectInfo({
                    filter: this.projectInfoFilter
                });
            },

            /**
             *
             * @param type
             */
            showPanel(type) {
                if (type === this.panel.create) {
                    this.panel.type = this.panel.create;
                    this.clearPanelProjectInfoItem();
                } else if (type === this.panel.edit) {
                    this.panel.type = this.panel.edit;
                    this.currentProjectInfoItem.isSelected = true;
                }

                this.panel.show = true;
            },

            closePanel() {
                this.panel.show = false;
                this.currentProjectInfoItem.isSelected = false;
                this.clearPanelProjectInfoItem();
            },

            selectProjectInfoItem(projectInfoItem) {
                this.showPanel(this.panel.edit);
                this.currentProjectInfoItem.isSelected = true;
                Object.assign(this.currentProjectInfoItem.item, projectInfoItem);
            },

            changeCurrentProjectInfoItem() {
                this.currentProjectInfoItem.hasChange = true;
            },

            cancelChanges() {
                this.clearPanelProjectInfoItem();
                this.closePanel();
            },

            clearPanelProjectInfoItem() {
                this.currentProjectInfoItem.item = new ProjectInfo();
                this.currentProjectInfoItem.hasChange = false;
            },

            saveChangesSubmit() {
                if (this.isPanelCreate) {
                    this.createProjectInfoItemSubmit();
                    return;
                }

                if (this.isPanelEdit) {
                    this.editProjectInfoItemSubmit();
                }
            },

            createProjectInfoItemSubmit() {
                return this.createProjectInfo({
					data: this.currentProjectInfoItem.item,
                }).then((response) => {

                    if (response.Model) {
                        this.addProjectInfoItemToList(response.Model);
                        this.clearPanelProjectInfoItem();
                    } else {
                        console.error('Ошибка создания записи: ', response.Error);
                    }

                }).catch(error => {
                    console.error('Ошибка создания записи: ', error);
                });
            },

            editProjectInfoItemSubmit() {

                if (this.currentProjectInfoItem.hasChange) {
                    return this.updateProjectInfo({
                        id: this.currentProjectInfoItem.item.Id,
                        data: this.currentProjectInfoItem.item,
                    }).then((response) => {

                        if (response.Model) {
                            this.updateProjectInfoById(response.Model);
                            this.currentProjectInfoItem.hasChange = false;
                            this.clearPanelProjectInfoItem();
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

            deleteProjectInfoItemHandler() {
                let deletedItemId = this.currentProjectInfoItem.item.Id;

                if (!this.currentProjectInfoItem.canDelete) {
                    this.currentProjectInfoItem.showDeleteConfirmation = true;
                    return;
                }

                this.deleteProjectInfo({
                    id: deletedItemId
                }).then(response => {

                    if (response.IsSuccess) {
                        this.deleteProjectInfoFromList(deletedItemId);
                        this.clearPanelProjectInfoItem();
                        this.currentProjectInfoItem.canDelete = false;
                        this.currentProjectInfoItem.isSelected = false;
                        this.panel.show = false;
                    } else {
                        console.error('Ошибка удаления элемента: ', response.Error);
                    }

                }).catch(error => {
                    console.error('Ошибка удаления элемента: ', error);
                });
            },

            confirmDeleteHandler() {
                this.currentProjectInfoItem.showDeleteConfirmation = false;
                this.currentProjectInfoItem.canDelete = true;
                this.deleteProjectInfoItemHandler();
            },

            closeConfirmationPanel() {
                this.currentProjectInfoItem.showDeleteConfirmation = false;
            },
        },
    }
</script>

<style lang="scss">

</style>
