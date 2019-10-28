<template>
  <div class="error-page">
    <h1 v-if="error.statusCode === 404">Страница не найдена</h1>
    <h1 v-else>Произошла ошибка</h1>
    <nuxt-link to="/">Вернуться на главную</nuxt-link>
  </div>
</template>

<script>
  export default {
    name: 'error',
    
    props: ['error']
  }
</script>

<style lang="scss">
  .error-page {
    min-height: 100vh;
    display: flex;
    flex-direction: column;
    justify-content: center;
    align-items: center;
    text-align: center;
  }
</style>
