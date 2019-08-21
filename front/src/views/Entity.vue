<template>
    <EntityGen>

        <template #pageHeader>
            <VSet vertical  indent-size="XS">
                <VHead level="h1">Entity</VHead>
                <VInput  placeholder="Поиск" v-model="searchModel" @input="search"></VInput>
            </VSet>
        </template>

        <template #data>

            <VSet vertical :isControl="false" v-if="entityList && entityList.length">
                <template v-for="entityItem in entityList">
                    <EntityItem :entityItem="entityItem"></EntityItem>
                </template>
            </VSet>
            <VText class="loading"></VText>

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
    import EntityItem from "../components/EntityItem";

    export default {
        name: "Entity",
        components: {EntityItem, VSpoiler, VBadge, EntityGen, VGroup},
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

</style>