<template>
  <WorkSpace>
    <template #content>
      <VHead level="h1">Start page</VHead>



    </template>
  </WorkSpace>
</template>

<script>
// @ is an alias to /src

import WorkSpace from "swtui/src/components/WorkSpace";
import VHead from "swtui/src/components/VHead";
export default {
  name: 'home',
  data: function () {
      return {
          data: "Some data2",
      }
  },
  components: {
      VHead,
      WorkSpace,
  }
}
</script>
