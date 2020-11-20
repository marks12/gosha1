<template>
  <VSet vertical-align="top">
    <VSet>
      <VSet vertical width="dyn">
        <VSet vertical v-if="action === 'create' || action === 'update'">
          <VSet>
            <VLabel>Request</VLabel>
            <VButton :disabled="! isValidJson" @click="beautifyRequest" small text="Beatify (CTRL+B)"></VButton>
            <VButton @click="loadModel" :disabled="isLoadingModel" small text="Get model (CTRL+L)"></VButton>
            <VBadge color="attention" v-if="! isValidJson">Invalid json</VBadge>
          </VSet>
          <VInput multiline rows="17" width="dyn" ref="requestModelField" placeholder="request"
                  v-model="requestModel"
                  style="white-space:  nowrap !important;"
                  @keydown.native="handleReqHotKey"></VInput>
          <VSet>
            <div><VBadge>CTRL+X</VBadge> - cut line</div>
            <div><VBadge>CTRL+D</VBadge> - duplicate line</div>
          </VSet>

          </VSet>
        <VSet indent-size="XS">
          <VLabel>Response</VLabel>
          <VSet v-if="isValidCode" indent-size="XS">
            <VBadge :color="getResponseCode() > 299 ? 'attention' : 'action'">{{ getResponseCode() }}</VBadge>
            <VLabel>{{ getResponseStatusText() }}</VLabel>
          </VSet>
        </VSet>
        <VInput multiline wrap="off" :rows="action !== 'find' ? '16' : '34'" width="dyn" :value="response"
                placeholder="response"></VInput>
      </VSet>
    </VSet>
    <VSet vertical divider>
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
          <VSet vertical v-if="entityFilter.Fields && entityFilter.Fields.length"
                v-for="(item, index) in entityFilter.Fields" :key="'k-' + item.Name + '-0' + index" width="col6"
                indent-size="XS">
            <VLabel>{{ item.Name }}</VLabel>
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
      requestModel: "",
      isValidJson: true,
      isLoadingModel: false,
    };
  },
  created() {

    this.$nextTick(() => {
      this.resetResponse();
      this.urlChanged();
    });

    this.fillServers();
    this.findFilter();

    this.setPanelMaxWidth("col12");

    this.requestModel = "";

    if (["create", "update"].indexOf(this.action) !== -1) {

      this.requestModel = localStorage.getItem(this.action + "-" + this.entity.Name) || "";

      if (! this.requestModel) {
        this.findEntityModel();
      } else {
        this.beautifyRequest();
      }
    }

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
      localStorage.setItem("routeUrl", newval);
      this.urlChanged();
    },
    requestFilter: {
      deep: true,
      handler(newVal) {
        console.log('requestFilter changed!', newVal);
        this.setFilters(newVal);
      }
    },
    authToken(newVal) {
      localStorage.setItem("authToken", newVal);
    },
    requestModel(newVal) {
      this.isValidJson = this.isValidJsonString(newVal);
      this.setBodyModel(newVal);
      localStorage.setItem(this.action + "-" + this.entity.Name, newVal)
    },
  },
  methods: {

    ...mapGetters('gosha', {
      entityList: 'getListEntity',
      getResponse: 'getResponse',
      getResponseCode: 'getResponseCode',
      getResponseStatusText: 'getResponseStatusText',
      getRequestUrl: 'getRequestUrl',
    }),

    ...mapActions('gosha', {
      resetResponse: 'resetResponse',
      findEntity: 'findEntity',
      setPanelMaxWidth: "setPanelMaxWidth",
      setFilters: "setFilters",
      setBodyModel: "setBodyModel",
      setRequestUrl: "setRequestUrl",
      setRouteUrl: "setRouteUrl",
    }),

    isValidJsonString(jsonString) {
      if (!(jsonString && typeof jsonString === "string")) {
        return false;
      }
      try {
        JSON.parse(jsonString);
        return true;
      } catch (error) {
        return false;
      }
    },
    tryAuth() {
      this.$refs.auth.tryAuth();
    },

    setToken(token) {
      this.authToken = token;
    },

    beautifyRequest() {
      if (this.isValidJson) {
        let r = JSON.parse(this.requestModel);
        this.requestModel = JSON.stringify(r, null, 4);
      }
    },
    loadModel() {
      this.findEntityModel();
    },
    handleReqHotKey(event) {
      if (event) {
        if (event.ctrlKey && event.key && (event.key.toLowerCase() === "b" || event.key.toLowerCase() === "и")) {
          this.beautifyRequest();
          event.preventDefault();
        }
        if (event.ctrlKey && event.key && (event.key.toLowerCase() === "l" || event.key.toLowerCase() === "д")) {
          this.findEntityModel();
          event.preventDefault();
        }
        if (event.ctrlKey && event.key && (event.key.toLowerCase() === "x" || event.key.toLowerCase() === "ч")) {
          let start = event.target.selectionStart;
          let end = event.target.selectionEnd;

          if (start !== end) {
            return ;
          }

          let lines = this.requestModel.replace(/\r\n/g,"\n").split("\n").filter(line => line);

          let len = 0;
          for (let i=0;i<lines.length;i++){
            len += lines[i].length + 1;
            if (len>start) {
              setTimeout(()=>{
                  let pos = lines.slice(0, i+1).join("\n").length
                  event.target.selectionStart = event.target.selectionEnd = pos;
              })
              lines.splice(i, 1);
              this.requestModel = lines.join("\n");
              break;
            }
          }
          event.preventDefault();
        }
        if (event.ctrlKey && event.key && (event.key.toLowerCase() === "d" || event.key.toLowerCase() === "в")) {
          let start = event.target.selectionStart;
          let lines = this.requestModel.replace(/\r\n/g,"\n").split("\n").filter(line => line);

          let len = 0;
          for (let i=0;i<lines.length;i++){
            len += lines[i].length + 1;
            if (len>start) {
              lines[i] = lines[i] + "\n" + lines[i];
              setTimeout(()=>{
                let pos = lines.slice(0, i+1).join("\n").length
                event.target.selectionStart = event.target.selectionEnd = pos;
              })
              this.requestModel = lines.join("\n");
              break;
            }
          }
          event.preventDefault();
        }
        if (event.key === "Tab") {
          let start = event.target.selectionStart;
          let val = event.target.value;
          event.target.value = val.substr(0, start) + "    " + val.substr(event.target.selectionEnd);
          event.target.selectionStart = event.target.selectionEnd = start + 4;
          event.preventDefault();
        }
      }
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
      }).then((res) => {

        if (res && res.List && res.List.length) {

          for (let k in res.List) {

            if (res.List[k].Name === this.entity.Name + "Filter" && res.List[k].IsFilter === true) {
              this.entityFilter = res.List[k];
            }
          }

          if (this.requestFilter && !(this.requestFilter.CurrentPage)) {
            this.requestFilter.CurrentPage = 1;
            this.setFilters(this.requestFilter);
          }

          if (this.requestFilter && !(this.requestFilter.PerPage)) {
            this.requestFilter.PerPage = 10;
            this.setFilters(this.requestFilter);
          }

          return;
        }

        this.error = "Filter not found";

        return false;
      });

    },

    findEmptyEntityModel(entityName){

      this.isLoadingModel = true;

      let f = new EntityFilter();
      f.PerPage = 1000;
      f.CurrentPage = 1;

      if (entityName) {
        f.Search = entityName
      } else {
        f.Search = this.entity.Name;
      }

      return this.findEntity({
        filter: f,
        noMutation: true,
      }).then((res) => {

        this.isLoadingModel = false;

        if (res && res.List && res.List.length) {

          for (let k in res.List) {
            if (res.List[k].Name === f.Search && res.List[k].IsFilter === false) {
              this.requestModel = this.genModel(res.List[k].Fields);
              this.beautifyRequest();
              return true;
            }
          }
          return;
        }

        this.error = "Model not found";

        return false;
      });
    },

    findEntityModel(entityName) {

      let url = this.getRequestUrl();
      if (url.match("/[0-9]+$")) {
        return this.readEntityModel(entityName);
      } else {
        return this.findEmptyEntityModel(entityName);
      }
    },

    readEntityModel(entityName) {

      this.isLoadingModel = true;

      return fetch(this.getRequestUrl(), {
        headers: {
          'Content-Type': 'application/json',
          'Token': this.authToken,
        },
      })
      .then((res)=>{
        return res.json();
      }).catch((error)=>{
        console.error('Cant get model from url:', this.getRequestUrl(), error);
        return this.findEmptyEntityModel(entityName);
      }).
      then((res) => {

        this.isLoadingModel = false;

        if (res && res.Model) {
            this.requestModel = JSON.stringify(res.Model);
            this.beautifyRequest();
            return true;
        }

        this.error = "Model not found";

        return false;
      });
    },

    genModel(fieldsList) {
      let reqObject = {};
      for (let item in fieldsList) {
        if (["CreatedAt", "UpdatedAt", "DeletedAt"].indexOf(fieldsList[item].Name) !== -1) {
          continue;
        }

        if (["Int"].indexOf(fieldsList[item].Type) !== -1) {
          reqObject[fieldsList[item].Name] = 0;
          continue;
        }
        if (["*Int"].indexOf(fieldsList[item].Type) !== -1) {
          reqObject[fieldsList[item].Name] = null;
          continue;
        }

        if (["String"].indexOf(fieldsList[item].Type) !== -1) {
          reqObject[fieldsList[item].Name] = "";
          continue;
        }
      }

      return JSON.stringify(reqObject);
    },

    urlChanged() {
      this.$emit("changeUrl", this.serverUrl + "" + this.routeUrl);
      this.setRequestUrl(this.serverUrl + "" + this.routeUrl);
      this.setRouteUrl(this.routeUrl);
    },
    fillServers() {
      let stored = localStorage.getItem("servers");
      this.servers = stored ? JSON.parse(store) : [];
    },
  },
  filters: {
    pretty: function (value) {
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