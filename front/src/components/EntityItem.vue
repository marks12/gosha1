<template>
    <WorkSpace width="fit" height="fit" :class="{'bg-filter' : entityItem.IsFilter}">
        <VSet divider vertical height="fit">
            <VSet>
                <VText :title="entityItem.Name">{{entityItem.Name}}</VText>
                <VSign width="M">Db</VSign>
                <VSign width="M">Types</VSign>
            </VSet>
            <VSet>
                <VBadge :color="entityItem.HttpMethods.IsFind ? 'action' : 'weak'">F</VBadge>
                <VBadge :color="entityItem.HttpMethods.IsCreate ? 'action' : 'weak'">C</VBadge>
                <VBadge :color="entityItem.HttpMethods.IsRead ? 'action' : 'weak'">R</VBadge>
                <VBadge :color="entityItem.HttpMethods.IsUpdate ? 'action' : 'weak'">U</VBadge>
                <VBadge :color="entityItem.HttpMethods.IsDelete ? 'action' : 'weak'">D</VBadge>
                <VBadge :color="entityItem.HttpMethods.IsFindOrCreate ? 'action' : 'weak'" title="Find or create">FoC</VBadge>
            </VSet>
            <VSet vertical hasNoIndent>
                <template v-for="(field, i) in entityItem.Fields">
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
            </VSet>
        </VSet>
        <template slot="footer">
            <VSet>
                <VButton text="Edit" small @click="onEdit(entityItem)"></VButton>
                <VButton text="Request" small accent @click="onRequest(entityItem)"></VButton>
            </VSet>
        </template>
    </WorkSpace>
</template>

<script>
    import VBadge from "swtui/src/components/VBadge";
    import VButton from "swtui/src/components/VButton";
    import VSpoiler from "swtui/src/components/VSpoiler";
    import VGroup from "swtui/src/components/VGroup";
    import VSet from "swtui/src/components/VSet";
    import VText from "swtui/src/components/VText";
    import VIcon from "swtui/src/components/VIcon";
    import VSign from "swtui/src/components/VSign";
    import WorkSpace from "swtui/src/components/WorkSpace";

    export default {
        name: "EntityItem",
        components: {WorkSpace, VSpoiler, VBadge, VGroup, VSet, VText, VIcon, VSign, VButton},

        props: {
            entityItem: Object,
            onEdit: Function,
            onRequest: Function,
        },
    }
</script>

<style scoped lang="scss">
    .bg-filter {
        background: #daffc9;
    }
    .sw-icon {
        background: transparent;
    }
</style>