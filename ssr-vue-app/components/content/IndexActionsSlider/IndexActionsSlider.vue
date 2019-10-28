<template>
  <div class="index-actions-slider">
    <div class="index-actions-slider__wrap">
      <div class="index-actions-slider__gallery">
        <div v-swiper:mySwiper="swiperOption">
          <div class="swiper-wrapper">
            <div
              v-for="item in items"
              :key="item.Id"
              class="swiper-slide index-actions-slider__item"
            >
              <div class="index-actions-slider__inner">
                <div class="index-actions-slider__header">
                  {{ item.Header }}
                </div>
                <div class="index-actions-slider__price">
                  <div class="index-actions-slider__price-old">
                    <span class="index-actions-slider__price-old-value">{{ item.PriceOld }} р.</span>
                    <span class="index-actions-slider__price-label">{{ item.Label }}</span>
                  </div>
                  <div class="index-actions-slider__price-value">
                    {{ item.Price }} р.
                  </div>
                </div>
                <div class="index-actions-slider__img-wrap">
                  <img :src="item.ImgSrc" alt="">
                </div>
                <div class="index-actions-slider__action">
                  <button class="button button_accent">{{ item.BtnText }}</button>
                  <a :href="item.Link">{{ item.LinkText }}</a>
                </div>
                <div class="index-actions-slider__title" v-html="item.Title"></div>
              </div>
            </div>
          </div>
          <div class="swiper-pagination"></div>
        </div>
      </div>
    </div>
  </div>
</template>

<script>
  import Vue from 'vue'
  import 'swiper/dist/css/swiper.css'
  import {adminToolMixin} from '../../mixins/adminToolMixin';

  if (process.browser) {
    const VueAwesomeSwiper = require('vue-awesome-swiper/dist/ssr');
    Vue.use(VueAwesomeSwiper);
  }

  export default {
    name: 'IndexActionsSlider',
    mixins: [adminToolMixin],
    props: {
      items: {
        type: Array,
        required: true
      }
    },
    data() {
      return {
        swiperOption: {
          slidesPerView: 'auto',
          loop: true,
          effect: 'fade',
          grabCursor: true,
          pagination: {
            el: '.swiper-pagination',
            clickable: true
          }
        },
      }
    },
  }
</script>

<style lang="scss">
  @import "../../../assets/scss/components/swiper";

  .index-actions-slider {
    position: relative;

    .swiper-container {
      padding-bottom: 63px;
      overflow: visible;
    }

    .swiper-slide {
      opacity: 0 !important;
    }

    .swiper-slide-active {
      opacity: 1 !important;
    }

    .swiper-pagination {
      bottom: 28px;
    }

    &__inner {
      position: relative;
    }

    &__header {
      margin-bottom: 8px;
      font-size: 32px;
      line-height: 37px;
      font-family: $font-header;
      color: $color-base-light;
    }

    &__price-old {
      display: flex;
      flex-wrap: wrap;
      align-items: center;
    }

    &__price-old-value {
      margin-right: 6px;
      font-size: 20px;
      line-height: 23px;
      color: $color-weak-light;
      text-decoration-line: line-through;
    }

    &__price-label {
      padding: 5px 16px 6px 16px;
      font-size: $font-size-small;
      line-height: $line-height-small;
      color: #fff;
      white-space: nowrap;
      background-color: $color-attention-light;
    }

    &__price-value {
      margin-top: 3px;
      font-size: 40px;
      line-height: 47px;
      font-family: $font-header;
      color: #000;
    }

    &__img-wrap {
      margin: -20px -25px 0;

      img {
        width: 100%;
      }
    }

    &__action {
      margin-top: -35px;

      .button {
        height: 40px;
        min-width: 120px;
        font-size: $font-size-base;
        line-height: 36px;
        border-radius: 20px;
      }
    }

    &__title {
      margin-top: 16px;
      font-size: $font-size-small;
      line-height: $line-height-small;
    }

    @media (min-width: $screen-lg) {
      flex-grow: 1;
      padding-top: 61px;

      &__header {
        font-size: $font-size-h0;
        line-height: $line-height-h0;
      }

      &__price-old-value {
        font-size: 26px;
        line-height: 30px;
      }

      &__price-value {
        font-size: $font-size-h0;
        line-height: $line-height-h0;
      }

      &__img-wrap {
        position: absolute;
        top: 26px;
        right: 5px;
        width: 559px;
        margin: 0;
      }

      &__header {
        margin-bottom: 38px;
      }

      &__price-label {
        padding-top: 6px;
        padding-bottom: 7px;
        font-weight: 700;
        font-size: $font-size-base;
        line-height: $line-height-base;
      }

      &__action {
        margin-top: 40px;

        .button {
          min-width: 152px;
          height: 56px;
          line-height: 52px;
          border-radius: 90px;
        }
      }

      &__title {
        margin-top: 36px;
      }

      &__wrap,
      &__gallery {
        height: 100%;
      }

      .swiper-container {
        height: 100%;
        padding-bottom: 100px;
      }

      .swiper-pagination {
        bottom: 45px;
      }
    }
  }
</style>
