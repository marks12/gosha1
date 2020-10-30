<template>
  <div class="promo-line">
    <nuxt-link :to="url" class="promo-line__link">{{ text }}</nuxt-link>
  </div>
</template>

<script>
  import data from '../../../data.json';
  import {adminToolMixin} from '../../mixins/adminToolMixin';

  export default {
    name: 'PromoLineTop',
    mixins: [adminToolMixin],
    data() {
      return {
        text: data.PromoLine.Text,
        url: data.PromoLine.Url
      }
    },
  }
</script>

<style lang="scss">

  .promo-line {
    text-align: center;
    background: lightgray;
  }

</style>
