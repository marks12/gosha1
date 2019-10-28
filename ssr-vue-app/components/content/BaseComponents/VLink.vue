<template>
    <a @click="click()" class="link">
      {{value}}
    </a>
</template>

<script>
  export default {
    name: 'VLink',
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
  @import "../../../assets/scss/_variables";

  .link {
    &:hover {
      text-decoration: none;
    }
  }

</style>
