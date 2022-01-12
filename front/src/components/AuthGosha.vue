<template>
  <div>
    <VPanel maxWidth="col-12" modal @close="callback()">
      <template #header>
        Авторизация в Gosha Cloud
      </template>
      <template #content>

        <VSet vertical divider>
          <VSet vertical>
            <VSet>
              <VInputAutocomplete
                  v-model="serverUrl"
                  :items="servers"
                  not-found-text="Never used"
                  placeholder="Server: https://gosha.info"
              >
              </VInputAutocomplete>
            </VSet>

            <VInput width="dyn" v-model="authEmail" id="user" placeholder="Email"></VInput>
            <VInput width="dyn" v-model="authPass" id="password" placeholder="Password" type="password"></VInput>
            <VText color="attention" v-if="authError">{{ authError }}</VText>

          </VSet>
        </VSet>

      </template>
      <template #footer>
        <VSet>
          <VButton text="Вход" accent @click="tryGoshaAuth"/>
          <a target="_blank" :href="getRegisterLink()">Регистрация / восстановление</a>
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
import VText from "swtui/src/components/VText";
import VInputAutocomplete from "swtui/src/components/VInputAutocomplete";
import cloud from "../store/cloud";

export default {
  props: {
    callback: {
      type: Function,
      default: (token) => {
        console.error("Close function not set", "token", token);
      },
    },
  },
  components: {WorkSpace, VButton, VSet, VPanel, VLabel, VInput, VIcon, VPopover, VInputAutocomplete, VText},
  name: "AuthGosha",
  data: function () {
    return {
      authRoute: localStorage.getItem("authRoute") || "",
      showPanelModal: false,
      showPopoverUrl: false,
      response: "",
      authError: "",
      authResponse: null,
      serverUrl: localStorage.getItem("serverUrl") || "",
      authEmail: "",
      authPass: "",
      authGoshaToken: localStorage.getItem("authGoshaToken") || "",
      servers: [],
    };
  },
  created() {
    this.fillServers();
  },
  methods: {
    fillServers() {
      let stored = localStorage.getItem("gosha-servers");
      this.servers = stored ? JSON.parse(stored) : [];

      this.servers.push(cloud.getServerUrl());

      if (this.servers.length < 2) {
        this.serverUrl = cloud.getServerUrl();
      }
    },
    tryGoshaAuth() {

      this.authError = "";

      let data = {
        Email: this.authEmail,
        Password: this.authPass,
      };

      localStorage.setItem("serverUrl", this.serverUrl);

      try {

        return cloud.auth(
            {
              Email: this.authEmail,
              Password: this.authPass,
            },
            (token) => {

              this.$root.$emit('setGoshaToken', token);
              this.callback(token);

            },
            (e) => {
              this.authError = e;
            })

      } catch (error) {
        this.authError = "Ошибка сети";
      }

      return new Response()
    },
    getRegisterLink() {
      return cloud.getServerUrl() + '/auth/register';
    },
  },
  watch: {
    authRoute() {
      localStorage.setItem("gosha-servers", JSON.stringify(this.servers));
    },
    authToken(newVal) {
      localStorage.setItem("authGoshaToken", newVal);
    },
  },
}
</script>

<style scoped>

</style>
