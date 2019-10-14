<template>

    <WorkSpace footerPosition="bottom" noBg>
        <template #header>
            <VSet>
                <VSet vertical  indent-size="XS">
                    <VHead level="h1">Entity</VHead>
                    <VSet>
                        <VInput  placeholder="Поиск" v-model="searchModel" @input="search"></VInput>
                        <VCheckbox v-model="entityFilter.WithFilter">
                            <VText font-size="txs">Отображать фильтры</VText>
                        </VCheckbox>
                    </VSet>
                </VSet>
            </VSet>
        </template>

        <template slot="content">
            <template>
                <VSet v-if="entityList && entityList.length" wrap>
                    <template v-for="entityItem in entityList">
                        <EntityItem :entityItem="entityItem" :onEdit="editItem"></EntityItem>
                    </template>
                </VSet>
                <VText class="loading"></VText>
            </template>

            <template>
                <VPanel
                        v-if="panel.show"
                        width="col3"
                        @close="closePanel"
                >
                    <VSet slot="header">
                        <VHead level="h3" width="fit">{{ panelHeader }}</VHead>
                        <VText>{{ currentEntityItem.Name }}</VText>
                    </VSet>

                    <template slot="content">
                        <form @submit.prevent="saveChangesSubmit">
                            <VSet direction="vertical" v-for="field in currentEntityItem.Fields" :key="'id-'+field.id">
                                <VSet>
                                    <template>
                                        <VSign width="col3">{{field.Type}}</VSign>
                                        <VText>{{field.Name}}</VText>
                                    </template>
                                </VSet>
                            </VSet>

                            <VHead level="h3" style="margin: 20px 0;">Add new field</VHead>
                            <VSet vertical>
                                <VSet v-for="f in newFields">
                                    <VInput v-model="f.Name" @input="updateNewFieldsList"></VInput>
                                    <VSelect v-model="f.Type" :items="filedTypes"></VSelect>
                                </VSet>
                            </VSet>
                        </form>
                    </template>

                    <template #footer>
                        <VSet>
                            <VButton
                                    @click="saveChangesSubmit"
                                    accent
                                    :text="panelSubmitButtonText"
                                    :disabled="!currentEntityItem.hasChange"
                            />
                            <VButton
                                    @click="cancelChanges"
                                    text="Cancel"
                            />
                        </VSet>
                    </template>
                </VPanel>
            </template>

            <slot name="confirmationPanel">
                <VPanel
                        v-if="currentEntityItem.showDeleteConfirmation"
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
            <VSet>
                <VButton
                        text="Добавить"
                        accent
                        @click="showPanel(panel.create)"
                />
            </VSet>
        </template>
    </WorkSpace>


</template>

<script>
    import EntityGen from "../../../webapp/jstypes/components/EntityGen";
    import VBadge from "swtui/src/components/VBadge";
    import VSpoiler from "swtui/src/components/VSpoiler";
    import VGroup from "swtui/src/components/VGroup";
    import EntityItem from "../components/EntityItem";
    import {Entity, EntityField, EntityFilter} from "../../../webapp/jstypes/apiModel";

    export default {
        name: "Entity",
        components: {EntityItem, VSpoiler, VBadge, EntityGen, VGroup},
        mixins: [
            EntityGen,
        ],
        data() {
            return {
                searchModel: "",
                entityFilter: new EntityFilter(),
                currentEntityItem: new Entity(),
                newFields: [
                    new EntityField(),
                ],
                filedTypes: [
                    "String",
                    "Bool",
                    "Integer",
                ],
            };
        },
        methods: {
            search() {

                this.entityFilter.Search = this.searchModel;

                this.findEntity({
                    filter: this.entityFilter
                })
            },
            editItem(item) {

                this.currentEntityItem = item;
                this.showPanel(this.panel.edit)
            },
            updateNewFieldsList() {

                for (let i = this.newFields.length - 1; i >= 0; i--) {
                    if (this.newFields[i].Name.trim().length < 1) {
                        this.newFields.splice(i , 1);
                    }
                }

                if (this.newFields[this.newFields.length - 1].Name.trim().length > 0) {
                    this.newFields.push(new EntityField());
                }
            },
        },
        computed: {
            hasFields(entityItem) {
                return (entityItem) => {
                    return  entityItem &&
                            entityItem.TypeFields &&
                            entityItem.TypeFields.length;
                }
            },
        },
        watch: {
            'entityFilter.WithFilter': function (newFilter) {

                this.findEntity({
                    filter: this.entityFilter
                })
            },
        },
    }
</script>

<style scoped>
</style>