<template>
  <WorkSpace>
    <template #content>
      Some content
    </template>
  </WorkSpace>
</template>

<script>
// @ is an alias to /src

import WorkSpace from "swtui/src/components/WorkSpace";
export default {
  name: 'home',
  components: {
      WorkSpace,
  }
}
</script>
