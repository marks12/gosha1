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
                        <VHead level="h3" width="fit">{{ panelHeader }} {{ currentEntityItem.Name }}</VHead>
                    </VSet>

                    <template slot="content">
                        <VSet vertical @submit.prevent="saveChangesSubmit">

                            <VSet vertical indent-size="XS">
                                <VSign>Entity name</VSign>
                                <VInput v-model="currentEntityItem.Name" :disabled="currentEntityItem.Id > 0" width="dyn"></VInput>
                            </VSet>

                            <VSet vertical indent-size="XS">
                                <VSet  indent-size="XS" direction="vertical" v-for="field in currentEntityItem.Fields" :key="'id-'+field.Name">
                                    <VSet>
                                        <template>
                                            <VSign width="col3">{{field.Type}}</VSign>
                                            <VText>{{field.Name}}</VText>
                                        </template>
                                    </VSet>
                                </VSet>
                            </VSet>

                            <VHead level="h3" style="margin: 20px 0;">Add new field</VHead>
                            <VSet vertical>
                                <VSet v-for="f in newFields" width="dyn" :key="'newf-' + f.Id">
                                    <VInput v-model="f.Name" @input="updateNewFieldsList" width="dyn"></VInput>
                                    <VSelect v-model="f.Type" :items="getTypes"></VSelect>
                                </VSet>
                            </VSet>
                        </VSet>
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
    import VSelect from "swtui/src/components/VSelect";
    import EntityItem from "../components/EntityItem";
    import { mapGetters, mapMutations, mapActions } from 'vuex';
    import {Entity, EntityField, EntityFilter, FieldTypeFilter} from "../../../webapp/jstypes/apiModel";

    export default {
        name: "Entity",
        components: {EntityItem, VSpoiler, VBadge, EntityGen, VGroup, VSelect},
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
                fieldTypeFilter: new FieldTypeFilter(),
                isLoading: true,
            };
        },
        created: function () {

            this.fieldTypeFilter.CurrentPage = 1;
            this.fieldTypeFilter.PerPage = 100;

            this.findFieldType({
                filter: this.fieldTypeFilter,
            }).then(()=>{
                this.isLoading = false;
            });

        },
        methods: {
            ...mapActions([
                "findFieldType",
            ]),
            ...mapGetters([
                "getListFieldType",
            ]),
            search() {

                this.entityFilter.Search = this.searchModel;

                this.findEntity({
                    filter: this.entityFilter
                });
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
            getTypes() {
                return this.getListFieldType().map( (item)=> {return item.Name});
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