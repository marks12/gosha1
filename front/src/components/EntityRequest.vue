<template>
    <VSet width="dyn" vertical>
        <VText>
            <VBadge :color="colors[action.toLowerCase()]">{{Method}}</VBadge>
            Request entity
        </VText>
        <HttpRequest :entity="entity.Name" :id="id" :route="entity.HttpRoutes.Find" v-if="IsFind" action="find"></HttpRequest>
        <HttpRequest :entity="entity.Name" :id="id" :route="entity.HttpRoutes.Create" v-if="IsCreate" action="create"></HttpRequest>
        <HttpRequest :entity="entity.Name" :id="id" :route="entity.HttpRoutes.Read" v-if="IsRead" action="read"></HttpRequest>
        <HttpRequest :entity="entity.Name" :id="id" :route="entity.HttpRoutes.Update" v-if="IsUpdate" action="update"></HttpRequest>
        <HttpRequest :entity="entity.Name" :id="id" :route="entity.HttpRoutes.Delete" v-if="IsDelete" action="delete"></HttpRequest>
        <HttpRequest :entity="entity.Name" :id="id" :route="entity.HttpRoutes.FindOrCreate" v-if="IsFindOrCreate" action="findOrCreate"></HttpRequest>
        <HttpRequest :entity="entity.Name" :id="id" :route="entity.HttpRoutes.UpdateOrCreate" v-if="IsFindOrCreate" action="updateOrCreate"></HttpRequest>
    </VSet>
</template>

<script>
    import HttpRequest from "./HttpRequest";
    import VSet from "swtui/src/components/VSet";
    import VText from "swtui/src/components/VText";
    export default {
        name: "EntityRequest",
        components: {HttpRequest, VSet, VText},
        props: {
            entity: {
                type: Object,
                default: {
                    HttpMethods: [],
                    HttpRoutes: [],
                },
            },
            action: String,
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
                    findOrCreate:'bg-light'
                },
            };
        },
        computed: {
            IsFind() {
                return this.action.toLowerCase() === "find" && entity.HttpMethods.IsFind;
            },
            IsCreate() {
                return this.action.toLowerCase() === "create" && entity.HttpMethods.IsCreate;
            },
            IsRead() {
                return this.action.toLowerCase() === "read" && entity.HttpMethods.IsRead;
            },
            IsUpdate() {
                return this.action.toLowerCase() === "update" && entity.HttpMethods.IsUpdate;
            },
            IsDelete() {
                return this.action.toLowerCase() === "delete" && entity.HttpMethods.IsDelete;
            },
            IsFindOrCreate() {
                return this.action.toLowerCase() === "findorcreate" && entity.HttpMethods.IsFindOrCreate;
            },
            IsUpdateOrCreate() {
                return this.action.toLowerCase() === "updateorcreate" && entity.HttpMethods.IsUpdateOrCreate;
            },
        },
        methods: {
        },
    }
</script>

<style scoped>

</style>