<template>
    <EntityGen v-if="entityList && entityList.length">

        <template #pageHeader>
            <VSet vertical  indent-size="XS">
                <VHead level="h1">Entity</VHead>
                <VInput  placeholder="Поиск" v-model="searchModel" @input="search"></VInput>
            </VSet>
        </template>

        <template #data>

            <VSet width="fit" :isControl="false">
                <div id="masonry">
                    <VGroup height="dyn" class="entity" v-for="entityItem in entityList" :key="entityItem.Id">
                        <VSet divider vertical>
                            <VSet>
                                <VText width="col9">{{entityItem.Name}}</VText>
                                <VSign width="col1">Db</VSign>
                                <VSign width="col2">Types</VSign>
                            </VSet>
                            <VSet vertical hasNoIndent>
                                <VSet v-for="(field, i) in entityItem.Fields" :key="entityItem.Id + i+ field.Name">
                                    <VText width="col9">
                                        {{ field.Name }}
                                    </VText>
                                    <VText width="col1">
                                        <VIcon name="check" v-if="field.IsDb"></VIcon>
                                        <VIcon name="minus" v-else></VIcon>
                                    </VText>
                                    <VText width="col2" align="center">
                                        <VIcon name="check" v-if="field.IsType"></VIcon>
                                        <VIcon name="minus" v-else></VIcon>
                                    </VText>
                                </VSet>
                            </VSet>
                        </VSet>
                    </VGroup>
                </div>
            </VSet>
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

    div#masonry {
        display: flex;
        flex-direction: row;
        flex-wrap: wrap;
        font-size: 0;
    }
    div#masonry img {
        width: 33.3%;
        transition: .8s opacity;
    }

    div#masonry:hover img { opacity: 0.3; }
    div#masonry:hover img:hover { opacity: 1; }

    /* fallback for earlier versions of Firefox */

    @supports not (flex-wrap: wrap) {
        div#masonry { display: block; }
        div#masonry img {
            display: inline-block;
            vertical-align: top;
        }
    }

    .entity {
        width:  270px;
        margin: 50px;
    }

</style>