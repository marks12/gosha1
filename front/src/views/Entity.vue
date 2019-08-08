<template>
    <EntityGen>

        <template #pageHeader>
            <VSet vertical  indent-size="XS">
                <VHead level="h1">Entity</VHead>
                <VInput  placeholder="Поиск" v-model="searchModel" @input="search"></VInput>
            </VSet>
        </template>

        <template #data>

            <VSet vertical v-if="entityList && entityList.length">

                <VGroup width="dyn" control v-for="entityItem in entityList" :key="entityItem.Id">

                    <VSet vertical>

                        <VSet indent-size="L" wrap style="margin-bottom: 30px;">

                            <VHead>{{ entityItem.Name }}</VHead>

                            <table style="width: auto;">
                                <tr v-if="hasFields(entityItem)">
                                    <td>Types.Entity</td>
                                    <td  v-for="f in entityItem.TypeFields" :key="f.Id">
                                        <VSet vertical indent-size="XS">
                                            <VText width="fit">{{f.Name}}</VText>
                                            <VBadge>{{f.Type}}</VBadge>
                                        </VSet>
                                    </td>
                                </tr>
                                <tr v-if="hasFields(entityItem)">
                                    <td>Dbmodels.Entity</td>
                                    <td  v-for="f in entityItem.ModelFields" :key="f.Id">
                                        <VSet vertical indent-size="XS">
                                            <VText width="fit">{{f.Name}}</VText>
                                            <VBadge>{{f.Type}}</VBadge>
                                        </VSet>
                                    </td>
                                </tr>
                                <tr v-if="! hasFields(entityItem)">
                                    <td>
                                        <VText>No  fields</VText>
                                    </td>
                                </tr>
                            </table>

                        </VSet>
                    </VSet>
                </VGroup>
            </VSet>
            <VText v-else>Models not loaded</VText>

        </template>

        <template #pageFooter>
            <VSet>
                <VButton
                        text="Добавить"
                        accent
                        @click="showPanel(panel.create)"
                />
            </VSet>
        </template>

        <template #panel>
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
                            <VSet vertical-align="center">
                                <VLabel width="col4">{{ filed }}</VLabel>
                                <VInput v-model="currentEntityItem.item.Name" width="fit"/>
                                {{currentEntityItem.item.Name}}
                            </VSet>
                        </VSet>
                        <button type="submit" :disabled="!currentEntityItem.hasChange" hidden></button>
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
                                text="Отменить"
                        />
                    </VSet>
                </template>
            </VPanel>
        </template>

    </EntityGen>


</template>

<script>
    import EntityGen from "../../../webapp/jstypes/components/EntityGen";
    import VBadge from "swtui/src/components/VBadge";
    import VSpoiler from "swtui/src/components/VSpoiler";
    import VGroup from "swtui/src/components/VGroup";

    export default {
        name: "Entity",
        components: {VSpoiler, VBadge, EntityGen, VGroup},
        mixins: [
            EntityGen,
        ],
        data() {
            return {
                searchModel: "",
            };
        },
        methods: {
            search() {

                this.entityFilter.Search = this.searchModel;

                this.findEntity({
                    filter: this.entityFilter
                })
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
    }
</script>

<style scoped>
    .h100 {
        height: 100vh;
    }
</style>