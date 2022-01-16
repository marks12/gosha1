<template>
  <VSet>
    <VSet v-if="isShowRename" indent-size="S">
      <VInput v-model="newName"></VInput>
      <VButton text="Ok" @click="renameDocument" accent></VButton>
      <VButton text="Отмена" @click="hideRename"></VButton>
    </VSet>
    <a v-else @click="showRename">{{doc.Name}}</a>
  </VSet>
</template>

<script>
import VSet from "swtui/src/components/VSet";
import VInput from "swtui/src/components/VInput";
import VButton from "swtui/src/components/VButton";

export default {
  props: {
    document: {
      type: Object,
    }
  },
  name: "Breadcrumbs",
  components: {VSet, VInput, VButton},
  data() {
    return {
      isShowRename: false,
      newName: "",
    }
  },
  methods: {
    showRename() {
      this.isShowRename = true;
      this.newName = this.document.Name;
    },
    hideRename() {
      this.isShowRename = false;
    },
    renameDocument() {
      this.$emit("rename", this.newName)
      this.isShowRename = false;
    },
  },
  computed: {
    doc() {
      return this.document;
    },
  }
}
</script>

<style scoped>

</style>
