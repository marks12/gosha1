<template>
    <span
      class="icon"
      :class="iconClass"
    >
    </span>
</template>

<script>
  export default {
    name: 'VIcon',
    props: {
      name: {
        type: String,
        required: true
      }
    },

    computed: {
      iconClass() {
        return `icon_${this.name}`;
      }
    }
  }
</script>

<style lang="scss">
  .icon {
    display: inline-block;
    vertical-align: middle;
  }
</style>
