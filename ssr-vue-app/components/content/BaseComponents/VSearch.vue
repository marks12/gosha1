<template>
  <form action="#" class="form-search">
    <VInput
      :placeholder="placeholder"
    />
    <button type="submit" class="form-search__btn">Найти</button>
  </form>
</template>

<script>
  import VInput from './VInput';

  export default {
    name: 'VSearch',
    components: {VInput},
    props: {
      placeholder: {
        type: String,
        default: ''
      }
    }
  }
</script>

<style lang="scss">
  .form-search {
    position: relative;
    width: 100%;

    &__btn {
      position: absolute;
      top: 0;
      right: 0;
    }
  }
</style>
