<template>
  <span class="logo">
    <nuxt-link to="/" class="logo__link">
      <img
        class="logo__img"
        :src="img"
        alt="Siberian Wellness">
    </nuxt-link>
  </span>
</template>

<script>
  export default {
    name: 'VLogo',
    props: {
      img: {
        type: String,
        required: true
      }
    }
  }
</script>

<style lang="scss">
  .logo {
    display: inline-block;
    vertical-align: middle;

  }
</style>

