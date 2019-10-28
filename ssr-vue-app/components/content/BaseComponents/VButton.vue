<template>
  <button
    @click="click()"
    :disabled="disabled"
    class="button">{{value}}</button>
</template>

<script>
  export default {
    name: 'VButton',
    props: {
      value: {
        type: String
      },
      disabled: {
        type: Boolean
      }
    },
    methods: {
      click() {
        this.$emit('click');
      }
    }
  }
</script>

<style lang="scss">

</style>
