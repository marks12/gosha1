<template>
  <VSet>
    <WorkSpace no-bg>
      <VSet width="dyn">
        <VSet width="dyn">
          <Breadcrumbs></Breadcrumbs>
        </VSet>
        <VButton :disabled="! isAuthorized" :title="isAuthorized ? 'Документы в cloud' : 'Не авторизован' " small text="Cloud документы"></VButton>
      </VSet>
    </WorkSpace>
  </VSet>
</template>

<script>
import Breadcrumbs from "../../components/bpm/Breadcrumbs";
import VSet from "swtui/src/components/VSet";
import VButton from "swtui/src/components/VButton";
import WorkSpace from "swtui/src/components/WorkSpace";
import cloud from "../../store/cloud";

export default {
  name: "Documents",
  components: {Breadcrumbs, VSet, WorkSpace, VButton},
  data(){
    return {
      isAuthorized: false,
    };
  },
  mounted() {
    cloud.onAuthorize((isAuth)=>{this.isAuthorized = isAuth});
  }
}
</script>

<style scoped>

</style>
