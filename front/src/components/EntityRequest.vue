<template>
    <VSet vertical>
        <VText>
            {{capitalizeFirstLetter(entity.Action)}} <VBadge :color="colors[this.action]">{{Method}}</VBadge>
            Request to <strong>{{entity.Name}}</strong>
        </VText>
        <HttpRequest @changeUrl="changeUrl" :entity="entity" :id="id" :route="entity.HttpRoutes.Find" v-if="IsFind" action="find"></HttpRequest>
        <HttpRequest @changeUrl="changeUrl" :entity="entity" :id="id" :route="entity.HttpRoutes.Create" v-if="IsCreate" action="create"></HttpRequest>
        <HttpRequest @changeUrl="changeUrl" :entity="entity" :id="id" :route="entity.HttpRoutes.Read" v-if="IsRead" action="read"></HttpRequest>
        <HttpRequest @changeUrl="changeUrl" :entity="entity" :id="id" :route="entity.HttpRoutes.Update" v-if="IsUpdate" action="update"></HttpRequest>
        <HttpRequest @changeUrl="changeUrl" :entity="entity" :id="id" :route="entity.HttpRoutes.Delete" v-if="IsDelete" action="delete"></HttpRequest>
        <HttpRequest @changeUrl="changeUrl" :entity="entity" :id="id" :route="entity.HttpRoutes.FindOrCreate" v-if="IsFindOrCreate" action="findOrCreate"></HttpRequest>
        <HttpRequest @changeUrl="changeUrl" :entity="entity" :id="id" :route="entity.HttpRoutes.UpdateOrCreate" v-if="IsUpdateOrCreate" action="updateOrCreate"></HttpRequest>
    </VSet>
</template>

<script>
    import HttpRequest from "./HttpRequest";
    import VSet from "swtui/src/components/VSet";
    import VText from "swtui/src/components/VText";
    import VBadge from "swtui/src/components/VBadge";
    import VLabel from "swtui/src/components/VLabel";
    export default {
        name: "EntityRequest",
        components: {HttpRequest, VSet, VText, VBadge, VLabel},
        props: {
            entity: {
                type: Object,
                default: {
                    HttpMethods: [],
                    HttpRoutes: [],
                    Action: "find",
                },
            },
        },
        data() {
            return {
                id: 0,
                colors: {
                    get: 'selection',
                    delete: 'attention',
                    create:'attention-secondary',
                    find: 'action',
                    read: 'weak',
                    // read: 'we',
                    findorcreate:'bg-light',
                    updateorcreate:'bg-light'
                },
                url: "",
            };
        },
        computed: {
            IsFind() {
                return this.action === "find" && this.entity.HttpMethods.IsFind;
            },
            IsCreate() {
                return this.action === "create" && this.entity.HttpMethods.IsCreate;
            },
            IsRead() {
                return this.action === "read" && this.entity.HttpMethods.IsRead;
            },
            IsUpdate() {
                return this.action === "update" && this.entity.HttpMethods.IsUpdate;
            },
            IsDelete() {
                return this.action === "delete" && this.entity.HttpMethods.IsDelete;
            },
            IsFindOrCreate() {
                return this.action === "findorcreate" && this.entity.HttpMethods.IsFindOrCreate;
            },
            IsUpdateOrCreate() {
                return this.action === "updateorcreate" && this.entity.HttpMethods.IsUpdateOrCreate;
            },
            /**
             * @return {string}
             */
            Method() {
                let action = this.entity ? this.entity.Action ? this.entity.Action.toLowerCase() : "" : "";
                let method = "";
                switch (action) {

                    case "find":
                    case "read":
                        method = "GET";
                        break;

                    case "update":
                    case "findorcreate":
                        method = "PUT";
                        break;
                    case "create":
                        method = "POST";
                        break;
                    case "delete":
                        method = "DELETE";
                        break;
                    case "updateorcreate":
                        method = "PATCH";
                        break;

                    default:
                        method = "UNSUPPORTED METHOD";
                        break;
                }
                return method;
            },
            action() {
                return this.entity ? this.entity.Action ? this.entity.Action : "" : "";
            },
        },
        methods: {
            changeUrl(newUrl) {
                this.url = newUrl;
            },
            capitalizeFirstLetter(string) {
                if (! string) {
                    return "";
                }

                return string.charAt(0).toUpperCase() + string.slice(1);
            }
        },
    }
</script>

<style scoped>

</style>