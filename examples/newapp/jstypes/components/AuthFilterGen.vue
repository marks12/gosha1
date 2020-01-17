
	<template>
    <WorkSpace>
        <template #header>
            <slot name="pageHeader">
                <VHead level="h1">AuthFilter</VHead>
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
                            v-for="authFilterItem in authFilterList"
                            :key="authFilterItem.Id"
                            @click="selectAuthFilterItem(authFilterItem)"
                            class="sw-table__row_can-select"
                            :class="{'sw-table__row_is-selected': authFilterItem.Id === currentAuthFilterItem.item.Id}"
                        >
                            <td v-for="(value, key) in fields" :key="key + '-fields'">
                                <VCheckbox v-if="isCheckbox(authFilterItem[key])" :checked="authFilterItem[key]" disabled></VCheckbox>
                                <VText v-else>{{ authFilterItem[key] }}</VText>
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
                                        :for="`currentAuthFilterItem${key}`"
                                    >{{ filed }}</VLabel>
                                    <VInput
										v-if="isInput(currentAuthFilterItem.item[key])"
                                        v-model="currentAuthFilterItem.item[key]"
                                        width="dyn"
                                        :id="`currentAuthFilterItem${key}`"
                                        @input="changeCurrentAuthFilterItem"
                                    />
									<VCheckbox
										v-if="isCheckbox(currentAuthFilterItem.item[key])"
                                        v-model="currentAuthFilterItem.item[key]"
                                        :id="`currentAuthFilterItem${key}`"
										@input="changeCurrentAuthFilterItem"
									/>
									
                                </VSet>
                            </VSet>
                            <button type="submit" :disabled="!currentAuthFilterItem.hasChange" hidden></button>
                        </form>
                    </template>

                    <template #footer>
                        <VSet>
                            <VButton
                                @click="saveChangesSubmit"
                                accent
                                :text="panelSubmitButtonText"
                                :disabled="!currentAuthFilterItem.hasChange"
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
                    v-if="currentAuthFilterItem.showDeleteConfirmation"
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
                        :disabled="!currentAuthFilterItem.isSelected"
                        @click="deleteAuthFilterItemHandler"
                    />
                </VSet>
            </slot>
        </template>
    </WorkSpace>
</template>

<script>
    import authFilterData from "../data/AuthFilterData";
    import { AuthFilter } from '../apiModel';
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
        name: 'AuthFilterGen',

        components: {VButton, VPanel, VText, VInput, VLabel, VSet, VHead, WorkSpace, VCheckbox},

        props: {
            fields: {
                type: Object,
                default() {
                    const authFilterItem = new AuthFilter();
                    const fieldsObj = {};

                    for (let prop in authFilterItem) {

                        if (authFilterItem.hasOwnProperty(prop)) {
                            fieldsObj[prop] = prop;
                        }

                    }

                    return fieldsObj;
                }
            },
            editFields: {
                type: Object,
                default() {
                    const authFilterItem = new AuthFilter();
                    const fieldsObj = {};

                    for (let prop in authFilterItem) {

                        if (authFilterItem.hasOwnProperty(prop)) {
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
            return authFilterData;
        },

        created() {
			this.onCreated();
        },

        computed: {
            ...mapGetters({
                authFilterList: 'getListAuthFilter'
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
                'findAuthFilter',
                'updateAuthFilter',
                'deleteAuthFilter',
                'createAuthFilter',
            ]),

            ...mapMutations([
                'addAuthFilterItemToList',
                'deleteAuthFilterFromList',
                'updateAuthFilterById',
            ]),

			onCreated() {
				this.fillAuthFilterFilter();
	            this.fetchAuthFilterData();
			},

            fillAuthFilterFilter() {
                this.authFilterFilter.CurrentPage = 1;
                this.authFilterFilter.PerPage = 1000;
            },

            fetchAuthFilterData() {
                return this.findAuthFilter({
                    filter: this.authFilterFilter
                });
            },

            /**
             *
             * @param type
             */
            showPanel(type) {
                if (type === this.panel.create) {
                    this.panel.type = this.panel.create;
                    this.clearPanelAuthFilterItem();
                } else if (type === this.panel.edit) {
                    this.panel.type = this.panel.edit;
                    this.currentAuthFilterItem.isSelected = true;
                }

                this.panel.show = true;
            },

            closePanel() {
                this.panel.show = false;
                this.currentAuthFilterItem.isSelected = false;
                this.clearPanelAuthFilterItem();
            },

            selectAuthFilterItem(authFilterItem) {
                this.showPanel(this.panel.edit);
                this.currentAuthFilterItem.isSelected = true;
                Object.assign(this.currentAuthFilterItem.item, authFilterItem);
            },

            changeCurrentAuthFilterItem() {
                this.currentAuthFilterItem.hasChange = true;
            },

            cancelChanges() {
                this.clearPanelAuthFilterItem();
                this.closePanel();
            },

            clearPanelAuthFilterItem() {
                this.currentAuthFilterItem.item = new AuthFilter();
                this.currentAuthFilterItem.hasChange = false;
            },

            saveChangesSubmit() {
                if (this.isPanelCreate) {
                    this.createAuthFilterItemSubmit();
                    return;
                }

                if (this.isPanelEdit) {
                    this.editAuthFilterItemSubmit();
                }
            },

            createAuthFilterItemSubmit() {
                return this.createAuthFilter({
					data: this.currentAuthFilterItem.item,
                }).then((response) => {

                    if (response.Model) {
                        this.addAuthFilterItemToList(response.Model);
                        this.clearPanelAuthFilterItem();
                    } else {
                        console.error('Ошибка создания записи: ', response.Error);
                    }

                }).catch(error => {
                    console.error('Ошибка создания записи: ', error);
                });
            },

            editAuthFilterItemSubmit() {

                if (this.currentAuthFilterItem.hasChange) {
                    return this.updateAuthFilter({
                        id: this.currentAuthFilterItem.item.Id,
                        data: this.currentAuthFilterItem.item,
                    }).then((response) => {

                        if (response.Model) {
                            this.updateAuthFilterById(response.Model);
                            this.currentAuthFilterItem.hasChange = false;
                            this.clearPanelAuthFilterItem();
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

            deleteAuthFilterItemHandler() {
                let deletedItemId = this.currentAuthFilterItem.item.Id;

                if (!this.currentAuthFilterItem.canDelete) {
                    this.currentAuthFilterItem.showDeleteConfirmation = true;
                    return;
                }

                this.deleteAuthFilter({
                    id: deletedItemId
                }).then(response => {

                    if (response.IsSuccess) {
                        this.deleteAuthFilterFromList(deletedItemId);
                        this.clearPanelAuthFilterItem();
                        this.currentAuthFilterItem.canDelete = false;
                        this.currentAuthFilterItem.isSelected = false;
                        this.panel.show = false;
                    } else {
                        console.error('Ошибка удаления элемента: ', response.Error);
                    }

                }).catch(error => {
                    console.error('Ошибка удаления элемента: ', error);
                });
            },

            confirmDeleteHandler() {
                this.currentAuthFilterItem.showDeleteConfirmation = false;
                this.currentAuthFilterItem.canDelete = true;
                this.deleteAuthFilterItemHandler();
            },

            closeConfirmationPanel() {
                this.currentAuthFilterItem.showDeleteConfirmation = false;
            },
        },
    }
</script>

<style lang="scss">

</style>
