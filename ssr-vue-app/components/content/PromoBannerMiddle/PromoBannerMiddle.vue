<template>
  <div class="promo-banner-middle">
    <nuxt-link :to="link" class="promo-banner-middle__link">
      <div class="container">
        <div class="promo-banner-middle__inner">
          <div class="promo-banner-middle__title">
            {{ title }}
          </div>
          <ul class="promo-banner-middle__list">
            <li
              v-for="(item, index) in items"
              :key="index"
              class="promo-banner-middle__item"
            >
              <div class="promo-banner-middle__img-wrap">
                <img :src="item.ImgSrc" :alt="item.ImgTitle" class="promo-banner-middle__img">
              </div>
            </li>
          </ul>
        </div>
      </div>
    </nuxt-link>
  </div>
</template>

<script>
  import data from '~/data.json';
  import {adminToolMixin} from '../../mixins/adminToolMixin';

  export default {
    name: 'PromoBannerMiddle',
    mixins: [adminToolMixin],
    data() {
      return {
        title: data.PromoBannerMiddle.Title,
        items: data.PromoBannerMiddle.Items,
        link: data.PromoBannerMiddle.Link
      }
    },
    /*props: {
      title: {
        type: String,
        default: ''
      },
      items: {
        type: Array,
        required: false
      },
      link: {
        type: [String, Object]
      }
    },*/
  }
</script>

<style lang="scss">
  .promo-banner-middle {
    padding: 20px 0;
    border: 1px solid;

    &__inner {
      display: flex;
      align-items: center;
    }

    &__list {
      display: flex;
    }

    &__item {
      margin: 0 10px;
    }

    &__img-wrap {
      width: 60px;
      height: 60px;
      background: lightgray;
    }
  }
</style>
