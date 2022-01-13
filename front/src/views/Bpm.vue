<template>
  <VSet vertical>
    <Documents></Documents>
    <VSet height="dyn">
      <WorkSpace width="fit">
        <AuthGosha v-if="isShowAuth" :callback="callbackAuth"></AuthGosha>
        <VSet vertical>
          <VHead>Элементы</VHead>
          <div id="BubuToolbox"></div>
          <VButton v-if="isAuthorized" style="padding: 0 15px;" @click="getData" text="Save" small></VButton>
          <VButton v-else style="padding: 0 15px;" @click="showAuth" text="Auth" small></VButton>
        </VSet>
      </WorkSpace>
      <WorkSpace noIndent>
        <canvas id="SomeCanvas"></canvas>
      </WorkSpace>
    </VSet>
  </VSet>
</template>

<style>
#BubuToolbox {
  display: flex;
  flex-direction: column;
  width: 100%;
}

#BubuToolbox img {
  margin-bottom: 15px;
}
</style>

<script>

import BuBu from '../bubu/main';
import {TYPES} from '../bubu/constants';
import VLink from "swtui/src/components/VLink";
import VButton from "swtui/src/components/VButton";
import VSet from "swtui/src/components/VSet";
import WorkSpace from "swtui/src/components/WorkSpace";
import VText from "swtui/src/components/VText";
import VImage from "swtui/src/components/VImage";
import VHead from "swtui/src/components/VHead";
import AuthGosha from "../components/AuthGosha";
import Documents from "../components/bpm/Documents";
import cloud from "../store/cloud";

export default {
  name: "Bpm",
  components: {Documents, AuthGosha, VHead, VImage, VText, WorkSpace, VSet, VButton, VLink},
  data() {
    return {
      text: "Some asd",
      bubu: {},
      src: {
        task: "",
        condition: "",
      },
      isAuthorized: false,
      isShowAuth: false,
    };
  },
  updated() {
    this.bubu.UpdateCanvas();
  },
  created() {
    this.$nextTick(() => {
      this.bubu = new BuBu('SomeCanvas', 'BubuToolbox');
      this.src.task = this.bubu.GetSrcImageTask();
      this.src.condition = this.bubu.GetSrcImageCondition();
      this.bubu.UpdateCanvas();
      this.bubu.Render();
    });
    this.checkAuth();
  },
  methods: {
    allowDrop: function (ev) {
      ev.preventDefault();
    },
    drop: function (event) {
      this.bubu.DropElement(event)
    },
    getData: function () {
      let data = this.bubu.GetData();

      console.log('data', data);
    },
    showAuth: function () {
      this.isShowAuth = true;
    },
    callbackAuth: function (token) {
      if (token) {
        this.checkAuth();
      } else {
        console.log("auth not changed", token);
      }
      this.isShowAuth = false;
    },
    checkAuth() {

      cloud.checkAuthorized((isAuthorized)=>{
        this.isAuthorized = isAuthorized;
      }, ()=>{})
    },
  },
}
</script>

<style scoped>

</style>
