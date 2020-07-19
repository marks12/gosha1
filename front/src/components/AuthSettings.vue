<template>
    <div>
        <VIcon name="settings" control size="M" title="Auth config" @click="showPanelModal = true"></VIcon>
        <VPanel maxWidth="col-12" v-if="showPanelModal" modal @close="showPanelModal = false">
            <template #header>
                Auth request configuration
            </template>
            <template #content>

                <VSet vertical divider>
                    <VSet vertical>
                        <VSet>
                            <VInputAutocomplete
                                    v-model="serverUrl"
                                    :items="servers"
                                    not-found-text="Never used"
                                    placeholder="Server: http://server.com"
                            >
                            </VInputAutocomplete>
                            <VInput width="dyn" id="url" placeholder="Url" v-model="authRoute">
                                <template #icon>
                                    <VIcon name="question" ref="buttonShowPopover" control @mouseenter="showPopoverUrl = !showPopoverUrl"/>
                                    <VPopover
                                            v-if="showPopoverUrl"
                                            :attachedTo="$refs.buttonShowPopover"
                                            trigger="mouseenter"
                                            @close="showPopoverUrl = false"
                                    >
                                        Popover Popover Popover Popover 1<br>
                                        Popover Popover Popover Popover 2<br>
                                        Popover Popover Popover Popover 3<br>
                                    </VPopover>
                                </template>
                            </VInput>
                        </VSet>

                        <VInput width="dyn" v-model="authEmail" id="user" placeholder="User"></VInput>
                        <VInput width="dyn" v-model="authPass" id="password" placeholder="Password" type="password"></VInput>


                    </VSet>

                    <VSet vertical>
                        <VLabel>Response</VLabel>
                        <VInput width="dyn" multiline rows="10" v-model="authResponse"/>
                    </VSet>
                </VSet>

            </template>
            <template #footer>
                <VSet>
                    <VButton text="Try" accent @click="tryAuth" />
                </VSet>
            </template>
        </VPanel>
    </div>
</template>

<script>
    import WorkSpace from "swtui/src/components/WorkSpace";
    import VButton from "swtui/src/components/VButton";
    import VSet from "swtui/src/components/VSet";
    import VPanel from "swtui/src/components/VPanel";
    import VLabel from "swtui/src/components/VLabel";
    import VInput from "swtui/src/components/VInput";
    import VIcon from "swtui/src/components/VIcon";
    import VPopover from "swtui/src/components/VPopover";
    import VInputAutocomplete from "swtui/src/components/VInputAutocomplete";

    export default {
        components: {WorkSpace, VButton, VSet, VPanel, VLabel, VInput, VIcon, VPopover, VInputAutocomplete},
        name: "AuthSettings",
        data: function () {
            return {
                authRoute: localStorage.getItem("authRoute") || "",
                showPanelModal: false,
                showPopoverUrl: false,
                response: "",
                authResponse: null,
                serverUrl: localStorage.getItem("serverUrl") || "",
                authEmail: localStorage.getItem("authEmail") || "",
                authPass: localStorage.getItem("authPass") || "",
                authToken: localStorage.getItem("authToken") || "",
                servers: [],
            };
        },
        created() {
            this.fillServers();
        },
        methods: {
            fillServers() {
                let stored = localStorage.getItem("servers");
                this.servers = stored ? JSON.parse(store) : [];
            },
            tryAuth() {

                let data = {
                    Email: this.authEmail,
                    Password: this.authPass,
                };

                try {

                    let url = this.serverUrl + this.authRoute;

                    const response = fetch(url, {
                        method: 'POST',
                        body: JSON.stringify(data),
                        headers: {
                            'Content-Type': 'application/json'
                        }
                    }).then((res)=>{

                        return res.json();

                    }).catch((error)=>{
                        console.error('Errors:', error);
                    }).then((data)=>{
                        this.authToken = data.Model.Token;
                        this.authResponse = JSON.stringify(data.Model, null, 4);
                        this.$emit('setToken', this.authToken);
                    });

                    return response;

                } catch (error) {
                    console.error('error:', error);
                }

                return new Response()

            },
        },
        watch: {
            authRoute(newVal) {
                localStorage.setItem("authRoute", newVal);
            },
            authEmail(newVal) {
                localStorage.setItem("authEmail", newVal);
            },
            authPass(newVal) {
                localStorage.setItem("authPass", newVal);
            },
            authToken(newVal) {
                localStorage.setItem("authToken", newVal);
            },
        },
    }
</script>

<style scoped>

</style>