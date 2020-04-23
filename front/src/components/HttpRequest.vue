<template>
    <VSet vertical indent-size="XL" width="dyn">
        <VSet>
            <VInput width="dyn" v-model="routeUrl"></VInput>
        </VSet>
    </VSet>
</template>

<script>

    import VSet from "swtui/src/components/VSet";
    import VBadge from "swtui/src/components/VBadge";
    import VInput from "swtui/src/components/VInput";

    export default {
        components: {VSet, VBadge, VInput},
        name: "HttpRequest",
        props: {
            action: {
                type: String,
                default: "",
            },
            entity: String,
            route: String,
        },
        data() {
            return {
                routeUrl: this.route,
            };
        },
        computed: {
        },
        watch: {
            route(newval) {
                this.routeUrl = newval;
            },
        },
    }
</script>

<style scoped>

</style>