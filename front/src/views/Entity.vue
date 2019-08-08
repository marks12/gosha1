<template>
    <EntityGen>
        <template #data>

            <VSet vertical>



                <VGroup width="dyn" control v-for="entityItem in entityList" :key="entityItem.Id">

                    <VSet vertical>

                        <VSet indent-size="L" wrap style="margin-bottom: 30px;">

                            <VHead>{{ entityItem.Name }}</VHead>

                            <template  v-if="entityItem.TypeFields && entityItem.TypeFields.length">

                                <VSet>
                                    <VLabel :width="'col2'">Types.Entity</VLabel>
                                    <VSet vertical indent-size="XS" width="fit" v-for="f in entityItem.TypeFields" :key="f.Id">
                                        <VText width="fit">{{f.Name}}</VText>
                                        <VBadge>{{f.Type}}</VBadge>
                                    </VSet>
                                </VSet>

                            </template>

                            <template  v-if="entityItem.TypeFields && entityItem.TypeFields.length">

                                <VSet>
                                    <VLabel :width="'col2'">Dbmodels.Entity</VLabel>
                                    <VSet vertical indent-size="XS" width="fit" v-for="f in entityItem.TypeFields" :key="f.Id">
                                        <VText width="fit">{{f.Name}}</VText>
                                        <VBadge>{{f.Type}}</VBadge>
                                    </VSet>
                                </VSet>

                            </template>

                            <VText v-else>No  fields</VText>

                        </VSet>
                    </VSet>
                </VGroup>
            </VSet>

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
    }
</script>

<style scoped>
    .h100 {
        height: 100vh;
    }
</style>