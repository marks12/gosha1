<template>
    <WorkSpace footer-position="bottom">
        <template #content>
            <template v-if="isLoaded">
                <VSet height="dyn" vertical vertical-align="center" horizontal-align="center">
                    <VSet width="fit">
                        <VHead level="h1" width="fit">Application dashboard</VHead>
                    </VSet>
                    <template v-if="currentApp.IsValidStructure">
                        <VText width="fit">
                            Your application successfully created!
                        </VText>
                    </template>
                    <template v-else>
                        <VText width="fit">
                            Please create application in current folder
                        </VText>
                        <VSet vertical width="fit" style="min-width: 360px;">
                            <VSet>
                                <VText>Email</VText>
                                <VInput v-model="adminMail" type="email"></VInput>
                            </VSet>
                            <VSet>
                                <VText>Password</VText>
                                <VInput v-model="adminPass" type="password"></VInput>
                            </VSet>
                            <VSet>
                                <VText>Re-Password</VText>
                                <VInput v-model="adminRePass" type="password"></VInput>
                            </VSet>
                            <VSet>
                                <VText width="dyn">Database</VText>
                                <VRadioGroup :items="segmentControls" v-model="appDb"/>
                            </VSet>
                            <VSet>
                                <VCheckbox v-model="IsUuidMode" width="dyn">
                                    <VText>Uuid as primary key</VText>
                                </VCheckbox>
                            </VSet>
                            <VSet>
                                <VText>Primary key</VText>
                                <VBadge  v-if="IsUuidMode" :color="IsUuidMode ? 'selection' : 'attention-secondary'">UUID</VBadge>
                                <VBadge  v-else :color="IsUuidMode ? 'selection' : 'attention-secondary'">Integer</VBadge>
                            </VSet>
                        </VSet>
                        <VSet width="fit">
                            <VText color="attention">{{ error }}</VText>
                        </VSet>
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
    import VInput from "swtui/src/components/VInput";
    import VRadioGroup from "swtui/src/components/VRadioGroup";
    import VCheckbox from "swtui/src/components/VCheckbox";
    import VBadge from "swtui/src/components/VBadge";

    export default {
        name: 'home',
        data: function () {
            return {
                data: "Some data2",
                isLoaded: false,
                adminMail: "",
                adminPass: "",
                adminRePass: "",
                appDb: "postgres",
                error: "",
                IsUuidMode: localStorage.getItem("IsUuidMode") === 'true',
                segmentControls: [
                    {
                        label: 'MySql',
                        value: "mysql"
                    },
                    {
                        label: 'Postgres',
                        value: "postgres"
                    }
                ],
            }
        },
        components: {VText, VSet, VHead, WorkSpace, VButton, VInput, VRadioGroup, VCheckbox, VBadge},
        created: function () {
            this.loadCurrentAppData();
        },
        methods: {
            ...mapActions('gosha', [
                "loadCurrentApp",
                "createCurrentApp"
            ]),
            createApp() {

                this.validate();

                if (this.error.length) {
                    return
                }

                let app = new CurrentApp();
                app.AdminEmail = this.adminMail;
                app.AdminPassword = this.adminPass;
                app.DbType = this.appDb;
                app.IsUuidMode = this.IsUuidMode;

                this.createCurrentApp({
                    data: app
                }).then(()=>{
                    this.loadCurrentAppData();
                });
            },
            loadCurrentAppData() {
                this.loadCurrentApp({id: 'current'}).then(()=>{
                    this.isLoaded = true;
                });
            },
            validate() {

                this.error = "";

                let re = /^(([^<>()\[\]\\.,;:\s@"]+(\.[^<>()\[\]\\.,;:\s@"]+)*)|(".+"))@((\[[0-9]{1,3}\.[0-9]{1,3}\.[0-9]{1,3}\.[0-9]{1,3}\])|(([a-zA-Z\-0-9]+\.)+[a-zA-Z]{2,}))$/;
                if (! re.test(String(this.adminMail).toLowerCase())) {
                    this.error += " Wrong email address. ";
                }

                if (this.adminPass.trim().length < 4) {
                    this.error += "Admin password must be more then 3 chars. ";
                }

                if (this.adminPass.trim() !== this.adminRePass) {
                    this.error += "Admin password and repeat password is not same. ";
                }
            },
        },
        computed: {
            ...mapGetters('gosha', {
                currentApp: "getCurrentApp"
            }),
        },
        watch: {
            IsUuidMode(newVal) {
                localStorage.setItem("IsUuidMode", newVal);
            },
        },
    }
</script>
