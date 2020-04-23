<template>
    <WorkSpace width="fit" height="fit" :class="{'bg-filter' : entityItem.IsFilter}">
        <VSet divider vertical height="fit">
            <VSet>
                <VText :title="entityItem.Name">{{entityItem.Name}}</VText>
                <VSign width="M">Db</VSign>
                <VSign width="M">Types</VSign>
            </VSet>
            <VSet>
                <VButton :disabled="! entityItem.HttpMethods.IsFind"  small :title="'Send Find Request'" :class="entityItem.HttpMethods.IsFind ? 'small-button find' : 'small-button weak'" text="F"></VButton>
                <VButton :disabled="! entityItem.HttpMethods.IsCreate" small :title="'Send Create Request'" :class="entityItem.HttpMethods.IsCreate ? 'create' : 'small-button weak'" text="C">C</VButton>
                <VButton :disabled="! entityItem.HttpMethods.IsRead" small :title="'Send Read Request'" :class="entityItem.HttpMethods.IsRead ? 'read' : 'small-button weak'" text="R">R</VButton>
                <VButton :disabled="! entityItem.HttpMethods.IsUpdate" small :title="'Send Update Request'" :class="entityItem.HttpMethods.IsUpdate ? 'update' : 'small-button weak'" text="U">U</VButton>
                <VButton :disabled="! entityItem.HttpMethods.IsDelete" small :title="'Send Delete Request'" :class="entityItem.HttpMethods.IsDelete ? 'delete' : 'small-button weak'" text="D">D</VButton>
                <VButton :disabled="! entityItem.HttpMethods.IsFindOrCreate" small :title="'Send Find or create Request'" :class="entityItem.HttpMethods.IsFindOrCreate ? 'foc' : 'small-button weak'" text="FoC" title="Find or create">FindOrCreate</VButton>
                <VButton :disabled="! entityItem.HttpMethods.IsUpdateOrCreate" small :title="'Send Update or create Request'" :class="entityItem.HttpMethods.IsUpdateOrCreate ? 'uoc' : 'small-button weak'" text="UoC" title="Update or create">UoC</VButton>
<!--                <VBadge :color="entityItem.HttpMethods.IsCreate ? 'action' : 'weak'">C</VBadge>-->
<!--                <VBadge :color="entityItem.HttpMethods.IsRead ? 'action' : 'weak'">R</VBadge>-->
<!--                <VBadge :color="entityItem.HttpMethods.IsUpdate ? 'action' : 'weak'">U</VBadge>-->
<!--                <VBadge :color="entityItem.HttpMethods.IsDelete ? 'action' : 'weak'">D</VBadge>-->
<!--                <VBadge :color="entityItem.HttpMethods.IsFindOrCreate ? 'action' : 'weak'" title="Find or create">FindOrCreate</VBadge>-->
<!--                <VBadge :color="entityItem.HttpMethods.IsUpdateOrCreate ? 'action' : 'weak'" title="Update or create">UoC</VBadge>-->
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
        },
        methods: {
            onRequest(item) {
                item.Action = "find";
                this.$emit("onRequest", item);
            },
        },
    }
</script>

<style scoped lang="scss">

    @import "~swtui/src/scss/variables";

    .bg-filter {
        background: #daffc9;
    }
    .sw-icon {
        background: transparent;
    }
    .sw-button_small.find { background: $color-action; color: $color-bg-light;}
    .sw-button_small.create { background: $color-attention-secondary; color: $color-bg-light;}
    .sw-button_small.read { background: $color-action; color: $color-bg-light;}
    .sw-button_small.update { background: $color-selection-light; color: $color-bg-light;}
    .sw-button_small.delete { background: $color-attention; color: $color-bg-light;}
    .small-button {
        padding: 0 12px;
        margin: 0 5px;
    }
</style>