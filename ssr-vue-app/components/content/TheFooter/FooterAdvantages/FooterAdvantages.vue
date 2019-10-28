<template>
  <div class="footer-advantages">
    <ul class="footer-advantages__list">
      <li
        v-for="item in items"
        :key="item"
        class="footer-advantages__item"
      >
        {{ item }}
      </li>
    </ul>
  </div>
</template>

<script>
  import data from '../../../../data.json'
  import {adminToolMixin} from '../../../mixins/adminToolMixin';

  export default {
    name: 'FooterAdvantages',
    data() {
      return {
        items: data.Footer.Advantages
      }
    }
  }
</script>

<style lang="scss">

</style>
