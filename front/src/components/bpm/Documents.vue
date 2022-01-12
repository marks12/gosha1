<template>
  <VSet>
    <WorkSpace no-bg>
      <VSet width="dyn">
        <VSet width="dyn">
          <Breadcrumbs></Breadcrumbs>
        </VSet>
        <VButton small text="Документы в cloud"></VButton>
      </VSet>
    </WorkSpace>
  </VSet>
</template>

<script>
import Breadcrumbs from "../../components/bpm/Breadcrumbs";
import VSet from "swtui/src/components/VSet";
import VButton from "swtui/src/components/VButton";
import WorkSpace from "swtui/src/components/WorkSpace";

export default {
  name: "Documents",
  components: {Breadcrumbs, VSet, WorkSpace, VButton},
}
</script>

<style scoped>

</style>
