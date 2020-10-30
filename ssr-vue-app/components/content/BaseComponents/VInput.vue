<template>
  <input
    :type="type"
    class="form__input"
    :placeholder="placeholder"
    :value="value"
  >
</template>

<script>
  export default {
    name: 'VInput',
    props: {
      type: {
        type: String,
        default: 'text'
      },

      value: {
        default: ''
      },

      placeholder: {
        type: String,
        default: ''
      }
    }
  }
</script>

<style lang="scss">
  .form__input {
    width: 100%;
  }
</style>
