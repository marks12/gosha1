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
            action: String,
            entity: String,
            route: String,
        },
        data() {
            return {
                routeUrl: this.route,
            };
        },
        computed: {
            /**
             * @return {string}
             */
            Method() {
                let method = "";
                switch (this.action.toLowerCase()) {

                    case "find":
                    case "read":
                        method = "GET";
                        break;

                    case "update":
                        method = "PUT";
                        break;
                    case "create":
                    case "findOrCreate":
                        method = "POST";
                        break;
                    case "delete":
                        method = "DELETE";
                        break;

                    default:
                        method = "UNSUPPORTED METHOD";
                        break;
                }
                return method;
            },
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