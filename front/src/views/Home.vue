<template>
    <WorkSpace>
        <template #content>
            <VHead level="h1">Start page</VHead>

            <template v-if="isLoaded">
                <template v-if="currentApp.IsValidStructure">
                    Your application successfully created!
                </template>
                <template v-else>
                    <VSet>
                        <VButton text="Create application" @click="createApp"></VButton>
                    </VSet>
                </template>
            </template>

        </template>
    </WorkSpace>
</template>

<script>
    // @ is an alias to /src

    import WorkSpace from "swtui/src/components/WorkSpace";
    import VHead from "swtui/src/components/VHead";
    import VButton from "swtui/src/components/VButton";
    import {mapGetters, mapMutations, mapActions} from 'vuex';
    import {CurrentApp, CurrentAppFilter} from "../../../webapp/jstypes/apiModel";
    import currentApp from "../../../webapp/jstypes/store/CurrentApp";
    import VSet from "swtui/src/components/VSet";

    export default {
        name: 'home',
        data: function () {
            return {
                data: "Some data2",
                isLoaded: false,
            }
        },
        components: {VSet, VHead, WorkSpace, VButton},
        created: function () {
            this.loadCurrentAppData();
        },
        methods: {
            ...mapActions([
                "loadCurrentApp",
                "createCurrentApp"
            ]),
            createApp() {
                this.createCurrentApp(new CurrentApp()).then(()=>{
                    this.loadCurrentAppData();
                });
            },
            loadCurrentAppData() {
                this.loadCurrentApp({id: 'current'}).then(()=>{
                    this.isLoaded = true;
                });
            },
        },
        computed: {
            ...mapGetters({
                currentApp: "getCurrentApp"
            }),
        },
    }
</script>
