<template>
  <WorkSpace>
    <template #content>
      {{data}}
    </template>
  </WorkSpace>
</template>

<script>
// @ is an alias to /src

import WorkSpace from "swtui/src/components/WorkSpace";
export default {
  name: 'home',
  data: function () {
      return {
          data: "Some data2",
      }
  },
  components: {
      WorkSpace,
  }
}
</script>
