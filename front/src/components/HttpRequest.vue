<template>
    <VSet vertical-align="top">
        <VSet>
            <VSet vertical width="dyn">
                <VSet vertical v-if="action !== 'find'">
                    <VLabel>Request</VLabel>
                    <VInput multiline rows="17" width="dyn" placeholder="request"></VInput>
                </VSet>
                <VSet indent-size="XS">
                    <VLabel>Response</VLabel>
                    <VSet v-if="isValidCode" indent-size="XS">
                        <VBadge :color="getResponseCode() > 299 ? 'attention' : 'action'">{{getResponseCode()}}</VBadge>
                        <VLabel>{{getResponseStatusText()}}</VLabel>
                    </VSet>
                </VSet>
                <VInput multiline :rows="action !== 'find' ? '16' : '34'" width="dyn" :value="response" placeholder="response"></VInput>
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
                        <VInput width="dyn" v-model="authToken"></VInput>
                        <VButton text="Auth" @click="tryAuth"></VButton>
                        <AuthSettings ref="auth" @setToken="setToken"></AuthSettings>
                    </VSet>
                </VSet>
            </VSet>
            <VSet vertical>
                <VSet wrap>
                    <VSet vertical v-if="entityFilter.Fields && entityFilter.Fields.length" v-for="item in entityFilter.Fields" :key="'k-' + item.Name" width="col6" indent-size="XS">
                        <VLabel>{{item.Name}}</VLabel>
                        <VInput v-model="requestFilter[item.Name]" width="dyn"></VInput>
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
    import VIcon from "swtui/src/components/VIcon";
    import VInputAutocomplete from "swtui/src/components/VInputAutocomplete";
    import {mapGetters, mapActions} from "vuex";
    import {EntityFilter} from "../../../webapp/jstypes/apiModel";
    import AuthSettings from "./AuthSettings";

    export default {
        components: {AuthSettings, VButton, VSet, VBadge, VInput, VInputAutocomplete, VText, VLabel, VGroup, VIcon},
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
                authToken: localStorage.getItem("authToken") || "",
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
                this.resetResponse();
                this.urlChanged();
            });

            this.fillServers();
            this.findFilter();

            this.setPanelMaxWidth("col12");

        },
        computed: {
            response() {

                let r = this.getResponse();

                try {
                    r = JSON.parse(r);
                    r = JSON.stringify(r, null, 4)
                } catch (e) {

                }

                return r;
            },
            isValidCode() {
                return 1 * this.getResponseCode() >= 200;
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
                getResponse: 'getResponse',
                getResponseCode: 'getResponseCode',
                getResponseStatusText: 'getResponseStatusText',
            }),

            ...mapActions('gosha', {
                resetResponse: 'resetResponse',
                findEntity: 'findEntity',
                setPanelMaxWidth: "setPanelMaxWidth",
                setFilters: "setFilters",
                setRequestUrl: "setRequestUrl",
            }),

            tryAuth() {
                this.$refs.auth.tryAuth();
            },

            setToken(token) {
                this.authToken = token;
            },

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

<style>
    .sw-set_has-indent.sw-set_indent-M > .sw-width_col6 {
        width: calc(((100% - 11 * 16px) * 6 / 12) + ((6 - 1) * 14px)) !important;
    }
</style>