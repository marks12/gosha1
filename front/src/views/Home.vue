<template>
    <WorkSpace footer-position="bottom">
        <template #content>
            <template v-if="isLoaded">
                <VSet height="dyn" vertical vertical-align="center" horizontal-align="center">
                    <VSet width="fit">
                        <VHead level="h1" width="fit">Application dashboard</VHead>
                    </VSet>
                    <VText width="fit">
                        Please create application in current folder
                    </VText>
                    <template v-if="currentApp.IsValidStructure">
                        Your application successfully created!
                    </template>
                    <template v-else>
                        <VButton text="Create application" @click="createApp" accent></VButton>
                    </template>
                </VSet>

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
    import VText from "swtui/src/components/VText";

    export default {
        name: 'home',
        data: function () {
            return {
                data: "Some data2",
                isLoaded: false,
            }
        },
        components: {VText, VSet, VHead, WorkSpace, VButton},
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
