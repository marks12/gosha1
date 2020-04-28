<template>
    <VSet>
        <VSet>
            <VSet vertical indent-size="XL" divider v-if="! error">
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
            <VText v-else>{{error}}</VText>
        </VSet>
    </VSet>
</template>

<script>

    import VSet from "swtui/src/components/VSet";
    import VBadge from "swtui/src/components/VBadge";
    import VButton from "swtui/src/components/VButton";
    import VInput from "swtui/src/components/VInput";
    import VLabel from "swtui/src/components/VLabel";
    import VText from "swtui/src/components/VText";
    import VInputAutocomplete from "swtui/src/components/VInputAutocomplete";
    import {mapGetters, mapActions} from "vuex";
    import {EntityFilter} from "../../../webapp/jstypes/apiModel";

    export default {
        components: {VButton, VSet, VBadge, VInput, VInputAutocomplete, VText, VLabel},
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
            };
        },
        created() {

            this.$nextTick(()=>{
                this.urlChanged();
            });

            this.fillServers();
            this.findFilter();

        },
        methods: {

            ...mapGetters('gosha', {
                entityList: 'getListEntity'
            }),

            ...mapActions('gosha', {
                findEntity: 'findEntity',
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