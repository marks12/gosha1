
	<template>
    <WorkSpace>
        <template #header>
            <slot name="pageHeader">
                <VHead level="h1">OrderFilter</VHead>
            </slot>
        </template>

        <template #content>
            <slot name="data">
                <table>
                    <thead>
                        <tr>
                            <th v-for="header in fields">{{ header }}</th>
                        </tr>
                    </thead>
            
                    <tbody>
                        <tr
                            v-for="orderFilterItem in orderFilterList"
                            :key="orderFilterItem.Id"
                            @click="selectOrderFilterItem(orderFilterItem)"
                            class="sw-table__row_can-select"
                            :class="{'sw-table__row_is-selected': orderFilterItem.Id === currentOrderFilterItem.item.Id}"
                        >
                            <td v-for="(value, key) in fields">
                                <VCheckbox v-if="isCheckbox(orderFilterItem[key])" :checked="orderFilterItem[key]" disabled></VCheckbox>
                                <VText v-else>{{ orderFilterItem[key] }}</VText>
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
                                    v-for="(filed, key) in editFields"
                                    vertical-align="center"
                                >
                                    <VLabel
                                        width="col4"
                                        :for="`currentOrderFilterItem${key}`"
                                    >{{ filed }}</VLabel>
                                    <VInput
										v-if="isInput(currentOrderFilterItem.item[key])"
                                        v-model="currentOrderFilterItem.item[key]"
                                        width="dyn"
                                        :id="`currentOrderFilterItem${key}`"
                                        @input="changeCurrentOrderFilterItem"
                                    />
									<VCheckbox
										v-if="isCheckbox(currentOrderFilterItem.item[key])"
                                        v-model="currentOrderFilterItem.item[key]"
                                        :id="`currentOrderFilterItem${key}`"
										@input="changeCurrentApplicationItem"
									/>
									
                                </VSet>
                            </VSet>
                            <button type="submit" :disabled="!currentOrderFilterItem.hasChange" hidden></button>
                        </form>
                    </template>

                    <template #footer>
                        <VSet>
                            <VButton
                                @click="saveChangesSubmit"
                                accent
                                :text="panelSubmitButtonText"
                                :disabled="!currentOrderFilterItem.hasChange"
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
                    v-if="currentOrderFilterItem.showDeleteConfirmation"
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
                        text="Добавить"
                        accent
                        @click="showPanel(panel.create)"
                    />
                    <VButton
                        text="Удалить"
                        :disabled="!currentOrderFilterItem.isSelected"
                        @click="deleteOrderFilterItemHandler"
                    />
                </VSet>
            </slot>
        </template>
    </WorkSpace>
</template>

<script>
    import orderFilterData from "../data/OrderFilterData";
    import { OrderFilter } from '../apiModel';
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
        name: 'OrderFilterGen',

        components: {VSelectSimple, VSign, VIcon, VButton, VPanel, VText, VInput, VLabel, VSet, VHead, WorkSpace, VCheckbox},

        props: {
            fields: {
                type: Object,
                default() {
                    const orderFilterItem = new OrderFilter();
                    const fieldsObj = {};

                    for (let prop in orderFilterItem) {

                        if (orderFilterItem.hasOwnProperty(prop)) {
                            fieldsObj[prop] = prop;
                        }

                    }

                    return fieldsObj;
                }
            },
            editFields: {
                type: Object,
                default() {
                    const orderFilterItem = new OrderFilter();
                    const fieldsObj = {};

                    for (let prop in orderFilterItem) {

                        if (orderFilterItem.hasOwnProperty(prop)) {
                            fieldsObj[prop] = prop;
                        }

                    }

                    return fieldsObj;
                }
            }
        },

        data() {
            return orderFilterData;
        },

        created() {
            this.fillOrderFilterFilter();
            this.fetchOrderFilterData();
        },

        computed: {
            ...mapGetters({
                orderFilterList: 'getListOrderFilter'
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
                'findOrderFilter',
                'updateOrderFilter',
                'deleteOrderFilter',
                'createOrderFilter',
            ]),

            ...mapMutations([
                'addOrderFilterItemToList',
                'deleteOrderFilterFromList',
                'updateOrderFilterById',
            ]),

            fillOrderFilterFilter() {
                this.orderFilterFilter.CurrentPage = 1;
                this.orderFilterFilter.PerPage = 1000;
            },

            fetchOrderFilterData() {
                return this.findOrderFilter({
                    filter: this.orderFilterFilter
                });
            },

            /**
             *
             * @param type
             */
            showPanel(type) {
                if (type === this.panel.create) {
                    this.panel.type = this.panel.create;
                    this.clearPanelOrderFilterItem();
                } else if (type === this.panel.edit) {
                    this.panel.type = this.panel.edit;
                    this.currentOrderFilterItem.isSelected = true;
                }

                this.panel.show = true;
            },

            closePanel() {
                this.panel.show = false;
                this.currentOrderFilterItem.isSelected = false;
                this.clearPanelOrderFilterItem();
            },

            selectOrderFilterItem(orderFilterItem) {
                this.showPanel(this.panel.edit);
                this.currentOrderFilterItem.isSelected = true;
                Object.assign(this.currentOrderFilterItem.item, orderFilterItem);
            },

            changeCurrentOrderFilterItem() {
                this.currentOrderFilterItem.hasChange = true;
            },

            cancelChanges() {
                this.clearPanelOrderFilterItem();
                this.closePanel();
            },

            clearPanelOrderFilterItem() {
                this.currentOrderFilterItem.item = new OrderFilter();
                this.currentOrderFilterItem.hasChange = false;
            },

            saveChangesSubmit() {
                if (this.isPanelCreate) {
                    this.createOrderFilterItemSubmit();
                    return;
                }

                if (this.isPanelEdit) {
                    this.editOrderFilterItemSubmit();
                }
            },

            createOrderFilterItemSubmit() {
                this.createOrderFilter({
                    data: {
                        Name: this.currentOrderFilterItem.item.Name,
                        Value: this.currentOrderFilterItem.item.Value,
                        Description: this.currentOrderFilterItem.item.Description,
                    }
                }).then((response) => {

                    if (response.Model) {
                        this.addOrderFilterItemToList(response.Model);
                        this.clearPanelOrderFilterItem();
                    } else {
                        console.error('Ошибка создания записи: ', response.Error);
                    }

                }).catch(error => {
                    console.error('Ошибка создания записи: ', error);
                });
            },

            editOrderFilterItemSubmit() {
                if (this.currentOrderFilterItem.hasChange) {
                    this.updateOrderFilter({
                        id: this.currentOrderFilterItem.item.Id,
                        data: this.currentOrderFilterItem.item,
                    }).then((response) => {

                        if (response.Model) {
                            this.updateOrderFilterById(response.Model);
                            this.currentOrderFilterItem.hasChange = false;
                            this.clearPanelOrderFilterItem();
                            this.closePanel();
                        } else {
                            console.error('Ошибка изменения записи: ', response.Error);
                        }

                    }).catch(error => {
                        console.error('Ошибка изменения записи: ', error);
                    });
                }
            },

            deleteOrderFilterItemHandler() {
                let deletedItemId = this.currentOrderFilterItem.item.Id;

                if (!this.currentOrderFilterItem.canDelete) {
                    this.currentOrderFilterItem.showDeleteConfirmation = true;
                    return;
                }

                this.deleteOrderFilter({
                    id: deletedItemId
                }).then(response => {

                    if (response.IsSuccess) {
                        this.deleteOrderFilterFromList(deletedItemId);
                        this.clearPanelOrderFilterItem();
                        this.currentOrderFilterItem.canDelete = false;
                        this.currentOrderFilterItem.isSelected = false;
                        this.panel.show = false;
                    } else {
                        console.error('Ошибка удаления элемента: ', response.Error);
                    }

                }).catch(error => {
                    console.error('Ошибка удаления элемента: ', error);
                });
            },

            confirmDeleteHandler() {
                this.currentOrderFilterItem.showDeleteConfirmation = false;
                this.currentOrderFilterItem.canDelete = true;
                this.deleteOrderFilterItemHandler();
            },

            closeConfirmationPanel() {
                this.currentOrderFilterItem.showDeleteConfirmation = false;
            },
        },
    }
</script>

<style lang="scss">

</style>
