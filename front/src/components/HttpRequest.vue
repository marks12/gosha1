<template>
    <VSet vertical indent-size="XL" width="dyn" divider>
        <VSet>
            <VInputAutocomplete
                    width="col5"
                    v-model="serverUrl"
                    :items="servers"
                    not-found-text="not used before"
                    placeholder="Server: http://server.com"
            >
            </VInputAutocomplete>
            <VInput width="dyn" v-model="routeUrl"></VInput>
        </VSet>
        <VSet>
            <VSet vertical>
                <VSet>
                    <VText>Token</VText>
                    <VText>
                        <VInput></VInput>
                    </VText>
                </VSet>
            </VSet>
            <VSet vertical>
                <VSet>
                    <VText>CurrentPage</VText>
                    <VText>
                        <VInput value="1"></VInput>
                    </VText>
                </VSet>
                <VSet>
                    <VText>PerPage</VText>
                    <VText>
                        <VInput value="10"></VInput>
                    </VText>
                </VSet>
            </VSet>
        </VSet>
        <VSet vertical>
            <VText>{{getFilter}}</VText>
        </VSet>
    </VSet>
</template>

<script>

    import VSet from "swtui/src/components/VSet";
    import VBadge from "swtui/src/components/VBadge";
    import VInput from "swtui/src/components/VInput";
    import VText from "swtui/src/components/VText";
    import VInputAutocomplete from "swtui/src/components/VInputAutocomplete";
    import {mapGetters} from "vuex";

    export default {
        components: {VSet, VBadge, VInput, VInputAutocomplete, VText},
        name: "HttpRequest",
        props: {
            action: {
                type: String,
                default: "",
            },
            entity: Object,
            route: String,
        },
        data() {
            return {
                routeUrl: this.route,
                serverUrl: localStorage.getItem("serverUrl") || "",
                servers: [],
                url: "",
                filter: {},
            };
        },
        created() {

            this.$nextTick(()=>{
                this.urlChanged();
            })

            this.fillServers();
            this.findFilter();

        },
        methods: {

            ...mapGetters('gosha', {
                entityList: 'getListEntity'
            }),

            findFilter() {

            },

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