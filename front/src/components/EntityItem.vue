<template>
    <VGroup width="dyn">
        <VSet divider vertical>
            <VSet>
                <VText :title="entityItem.Name">{{entityItem.Name}}</VText>
                <VSign width="M">Db</VSign>
                <VSign width="M">Types</VSign>
            </VSet>
            <VSet vertical hasNoIndent>
                <VSpoiler v-for="(field, i) in entityItem.Fields" :key="entityItem.Id + i+ field.Name">
                    <template #head>
                        <VSet>
                            <VText>
                                {{ field.Name }}
                            </VText>
                            <VText width="M">
                                <VIcon name="check" v-if="field.IsDb"></VIcon>
                                <VIcon name="minus" v-else></VIcon>
                            </VText>
                            <VText width="M">
                                <VIcon name="check" v-if="field.IsType"></VIcon>
                                <VIcon name="minus" v-else></VIcon>
                            </VText>
                        </VSet>
                    </template>
                    <template #content>
                        <EntityItem :entityItem="entityItem"></EntityItem>
                    </template>
                </VSpoiler>
            </VSet>
        </VSet>
    </VGroup>
</template>

<script>
    import VBadge from "swtui/src/components/VBadge";
    import VSpoiler from "swtui/src/components/VSpoiler";
    import VGroup from "swtui/src/components/VGroup";
    import VSet from "swtui/src/components/VSet";
    import VText from "swtui/src/components/VText";
    import VIcon from "swtui/src/components/VIcon";
    import VSign from "swtui/src/components/VSign";

    export default {
        name: "EntityItem",
        components: {VSpoiler, VBadge, VGroup, VSet, VText, VIcon, VSign},

        props: {
            entityItem: Object,
        },
    }
</script>

<style scoped>

</style>