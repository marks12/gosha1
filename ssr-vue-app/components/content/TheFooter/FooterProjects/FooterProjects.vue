<template>
  <div class="footer-projects">
    <h2 class="footer-projects__header">Проекты</h2>
    <ul class="footer-projects__list">
      <li
        v-for="item in items"
        :key="item.Id"
        class="footer-projects__item"
      >
        <nuxt-link
          :to="item.Id"
          class="footer-projects__link"
        >
          <img
            :src="item.Img"
            :alt="item.Title">
        </nuxt-link>
      </li>
    </ul>
  </div>
</template>

<script>
  export default {
    name: 'FooterProjects',
    props: {
      items: {
        type: Array,
      }
    }
  }
</script>

<style lang="scss">

</style>
