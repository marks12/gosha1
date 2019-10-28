<template>
  <div class="section">
    <h2 v-if="header" class="section__header">{{ header }}</h2>
    <div v-if="title" class="section__title">{{ title }}</div>
    <slot></slot>
  </div>
</template>

<script>
  export default {
    name: 'VSection',
    props: {
      header: {
        type: String,
        default: ''
      },

      title: {
        type: String,
        default: ''
      }
    }
  }
</script>

<style lang="scss">

  .section {
    display: flex;
    flex-direction: column;
    align-items: center;
    border: 1px solid;
  }

</style>
