<template>
    <VSet vertical-align="top">
        <VSet>
            <VSet vertical width="dyn">
                <VSet vertical>
                    <VLabel>Request</VLabel>
                    <VInput multiline rows="17" width="dyn" placeholder="request"></VInput>
                </VSet>
                <VLabel>Response</VLabel>
                <VInput multiline rows="16" width="dyn" :value="response" placeholder="response"></VInput>
            </VSet>
        </VSet>
        <VSet vertical divider v-if="! error">
            <VSet>
                <VInputAutocomplete
                        v-model="serverUrl"
                        :items="servers"
                        not-found-text="not used before"
                        placeholder="Server: http://server.com"
                >
                </VInputAutocomplete>
                <VInput width="dyn" v-model="routeUrl"></VInput>
            </VSet>
            <VSet>
                <VSet vertical indent-size="XS">
                    <VLabel>Token</VLabel>
                    <VSet>
                        <VInput width="dyn"></VInput>
                        <VButton text="Auth"></VButton>
                    </VSet>
                </VSet>
            </VSet>
            <VSet vertical>
                <VSet wrap>
                    <VSet vertical v-if="entityFilter.Fields && entityFilter.Fields.length" v-for="item in entityFilter.Fields" :key="'k-' + item.Name" width="fit" indent-size="XS">
                        <VLabel>{{item.Name}}</VLabel>
                        <VText>
                            <VInput v-model="requestFilter[item.Name]"></VInput>
                        </VText>
                    </VSet>
                </VSet>
            </VSet>
        </VSet>
    </VSet>
</template>

<script>

    import VSet from "swtui/src/components/VSet";
    import VGroup from "swtui/src/components/VGroup";
    import VBadge from "swtui/src/components/VBadge";
    import VButton from "swtui/src/components/VButton";
    import VInput from "swtui/src/components/VInput";
    import VLabel from "swtui/src/components/VLabel";
    import VText from "swtui/src/components/VText";
    import VInputAutocomplete from "swtui/src/components/VInputAutocomplete";
    import {mapGetters, mapActions} from "vuex";
    import {EntityFilter} from "../../../webapp/jstypes/apiModel";

    export default {
        components: {VButton, VSet, VBadge, VInput, VInputAutocomplete, VText, VLabel, VGroup},
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
                entityFilter: {},
                requestFilter: {},
                error: "",
                responseData: null,
            };
        },
        created() {

            this.$nextTick(()=>{
                this.urlChanged();
            });

            this.fillServers();
            this.findFilter();

            this.setPanelMaxWidth("col8");

        },
        computed: {
            response() {
                return JSON.stringify(this.getResponse(), null, 4);
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
            requestFilter: {
                deep: true,
                handler(newVal) {
                    console.log('requestFilter changed!', newVal);
                    this.setFilters(newVal);
                }
            },
        },
        methods: {

            ...mapGetters('gosha', {
                entityList: 'getListEntity',
                getResponse: 'getResponse'
            }),

            ...mapActions('gosha', {
                findEntity: 'findEntity',
                setPanelMaxWidth: "setPanelMaxWidth",
                setFilters: "setFilters",
                setRequestUrl: "setRequestUrl",
            }),

            findFilter() {
                let f = new EntityFilter();
                f.PerPage = 1000;
                f.CurrentPage = 1;
                f.WithFilter = true;
                f.Search = this.entity.Name + "Filter";

                this.findEntity({
                    filter: f,
                    noMutation: true,
                }).then((res)=>{

                    if (res && res.List && res.List.length) {

                        for (let k in res.List) {

                            if (res.List[k].Name === this.entity.Name + "Filter" && res.List[k].IsFilter === true) {
                                this.entityFilter = res.List[k];
                            }
                        }

                        return;
                    }

                    this.error = "Filter not found";

                    return false;
                });

            },

            urlChanged() {
                this.$emit("changeUrl", this.serverUrl + "" + this.routeUrl);
                this.setRequestUrl(this.serverUrl + "" + this.routeUrl);
            },
            fillServers() {
                let stored = localStorage.getItem("servers");
                this.servers = stored ? JSON.parse(store) : [];
            },
        },
        filters: {
            pretty: function(value) {
                return JSON.stringify(JSON.parse(value), null, 2);
            }
        }

    }
</script>

<style scoped>

</style>