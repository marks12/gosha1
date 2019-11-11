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
                        <VSet vertical width="fit">
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

    export default {
        name: 'home',
        data: function () {
            return {
                data: "Some data2",
                isLoaded: false,
                adminMail: "",
                adminPass: "",
                adminRePass: "",
                error: "",
            }
        },
        components: {VText, VSet, VHead, WorkSpace, VButton, VInput},
        created: function () {
            this.loadCurrentAppData();
        },
        methods: {
            ...mapActions([
                "loadCurrentApp",
                "createCurrentApp"
            ]),
            createApp() {

                this.validate();

                if (this.error.length) {
                    return
                }

                let app = new CurrentApp();
                app.Email = this.adminMail;
                app.Password = this.adminPass;

                this.createCurrentApp(new CurrentApp()).then(()=>{
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
            ...mapGetters({
                currentApp: "getCurrentApp"
            }),
        },
    }
</script>
