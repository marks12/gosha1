<template>
    <VSet vertical indent-size="XL" width="dyn">
        <VSet>
            <VInputAutocomplete
                    width="col5"
                    v-model="serverUrl"
                    :items="servers"
                    not-found-text="not found"
                    placeholder="введите название фрукта"
            >
                <template #header>
                    header
                </template>
            </VInputAutocomplete>
            <VInput width="dyn" v-model="routeUrl"></VInput>
        </VSet>
        <VText>{{url}}</VText>
    </VSet>
</template>

<script>

    import VSet from "swtui/src/components/VSet";
    import VBadge from "swtui/src/components/VBadge";
    import VInput from "swtui/src/components/VInput";
    import VText from "swtui/src/components/VText";
    import VInputAutocomplete from "swtui/src/components/VInputAutocomplete";

    export default {
        components: {VSet, VBadge, VInput, VInputAutocomplete, VText},
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
                serverUrl: localStorage.getItem("serverUrl") || "",
                servers: [],
                url: "",
            };
        },
        created() {

            this.$nextTick(()=>{
                this.urlChanged();
            })

            this.fillServers();

        },
        methods: {
            urlChanged() {
                this.$emit("changeUrl", this.serverUrl + "" + this.routeUrl);
            },
            fillServers() {
                let stored = localStorage.getItem("servers");
                this.servers = stored ? JSON.parse(store) : [];
            },
        },
        watch: {
            route(newval) {
                this.routeUrl = newval;
            },
            serverUrl(newval) {
                localStorage.setItem("serverUrl", newval);
                this.urlChanged();
            },
            routeUrl(newval) {
                localStorage.setItem("serverUrl", newval);
                this.urlChanged();
            },
        },
    }
</script>

<style scoped>

</style>