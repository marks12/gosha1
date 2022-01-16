<template>
  <VSet vertical>
    <Documents @rename="rename" :document="document"></Documents>
    <VSet height="dyn">
      <WorkSpace width="fit">
        <AuthGosha v-if="isShowAuth" :callback="callbackAuth"></AuthGosha>
        <VSet vertical>
          <VHead>Элементы</VHead>
          <div id="BubuToolbox"></div>
          <VButton v-if="isAuthorized" style="padding: 0 15px;" @click="saveDoc" text="Save" small></VButton>
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
import documentor from "../store/bpm";

export default {
  name: "Bpm",
  components: {Documents, AuthGosha, VHead, VImage, VText, WorkSpace, VSet, VButton, VLink},
  data() {
    return {
      text: "Some asd",
      bubu: null,
      src: {
        task: "",
        condition: "",
      },
      document: {
        Name: "Новый документ",
        Id: 0,
        ProjectId: 0,
        Scale: 0,
        ZeroPointX: 0,
        ZeroPointY: 0,
        Description: "",
        Items: [],
      },
      isAuthorized: false,
      isShowAuth: false,
    };
  },
  created() {
    this.$nextTick(() => {
      this.bubu = documentor.GetInstance('SomeCanvas', 'BubuToolbox');
    });
    this.checkAuth();
  },
  methods: {
    rename(name) {
      this.document.Name = name;
    },
    allowDrop: function (ev) {
      ev.preventDefault();
    },
    drop: function (event) {
      this.bubu.DropElement(event)
    },
    saveDoc: function () {
      let data = this.bubu.GetData();
      this.document.Scale = data.Scale;
      this.document.ZeroPointX = data.ZeroCoords.X;
      this.document.ZeroPointY = data.ZeroCoords.Y;
      this.document.ZeroPointY = data.ZeroCoords.Y;
      this.document.DocId = data.DocId;
      this.document.Items = data.Items;
      console.log('doc', this.document);
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
