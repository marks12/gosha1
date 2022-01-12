<template>
    <VSet horizontal>
        <VSet horizontal width="dyn">
            <VLink url="/" routerLink>Start</VLink>
            <VLink url="/entity" routerLink>Entity</VLink>
            <VLink url="/bpm" routerLink>Modeller</VLink>
        </VSet>
        <VSet horizontal width="fit">
            <VLink url="/about" routerLink nowrap>О проекте</VLink>
        </VSet>
    </Vset>
</template>

<script>
    import VSet from "swtui/src/components/VSet";
    import VLink from "swtui/src/components/VLink";
    export default {
        name: "MainMenu",
        components: {VLink, VSet}
    }
</script>

<style scoped>

</style>
