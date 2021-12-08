<template>
  <TheLayout id="app">
    <template #navigation>
      <VHead style="width: 200px;">
        {{ projectName }}
        <VBadge v-if="IsReadonly" :color="'attention'">readonly</VBadge>
      </VHead>
      <MainMenu></MainMenu>
    </template>
    <template slot="content">
      <router-view #content/>
    </template>
  </TheLayout>
</template>

<style lang="scss">

</style>

<script>
import TheLayout from "swtui/src/components/TheLayout";
import MainMenu from "./components/MainMenu";
import {mapGetters, mapActions} from 'vuex';
import {ProjectInfoFilter} from "../../webapp/jstypes/apiModel";
import VButton from "swtui/src/components/VButton";
import VHead from "swtui/src/components/VHead";
import VBadge from "swtui/src/components/VBadge";
import Vue from 'vue'
import VueSimpleMarkdown from 'vue-simple-markdown'
// You need a specific loader for CSS files like https://github.com/webpack/css-loader
import 'vue-simple-markdown/dist/vue-simple-markdown.css'

Vue.use(VueSimpleMarkdown)

export default {
  components: {VHead, VButton, MainMenu, TheLayout, VBadge},
  methods: {
    ...mapGetters('gosha', [
      'getListProjectInfo',
    ]),
    ...mapActions('gosha', [
      'findProjectInfo',
    ])
  },
  data() {
    return {
      piFileter: new ProjectInfoFilter(),
    };
  },
  created() {

    this.piFileter.CurrentPage = 1;
    this.piFileter.PerPage = 1000;

    this.findProjectInfo({filter: this.piFileter});
  },
  computed: {
    ...mapGetters('gosha', {
      panelMaxWidth: "getPanelMaxWidth",
      currentApp: "getCurrentApp",
    }),

    IsReadonly() {
      let current = this.currentApp;
      return current.IsReadonlyMode === true;
    },

    projectName() {

      let infoList = this.getListProjectInfo();

      let name = infoList.find(item => item.Name === "ProjectName");

      return name ? name.Value : '';
    },
  },

}
</script>
