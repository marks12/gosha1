<template>
  <div class="index-projects-slider">
    <div class="index-projects-slider__wrap">
      <div class="index-projects-slider__gallery">
        <div v-swiper:mySwiper="swiperOption">
          <div class="swiper-wrapper">
            <div
              v-for="item in items"
              :key="item.Id"
              class="swiper-slide index-projects-slider__item"
            >
              <div
                class="index-projects-slider__inner"
                :style="{'background-image': `url(${item.BgImgSrc})`}"
              >
                <div class="index-projects-slider__img-wrap">
                  <img
                    v-if="item.ImgSrc"
                    :src="item.ImgSrc"
                    class="index-projects-slider__img"
                    alt="">
                  <img
                    v-else
                    :src="require('assets/images/logo-icon.svg')"
                    class="index-projects-slider__img index-projects-slider__img_dummy"
                    alt=""
                  >
                </div>
                <div class="index-projects-slider__content">
                  <div class="index-projects-slider__header">
                    {{ item.Header }}
                  </div>
                  <div class="index-projects-slider__text" v-html="item.Content">

                  </div>
                </div>
              </div>
            </div>
          </div>
          <div class="gallery-nav">
            <span class="gallery__control gallery__control_prev">назад</span>
            <div class="swiper-pagination"></div>
            <span class="gallery__control gallery__control_next">вперед</span>
          </div>
        </div>
      </div>
    </div>
  </div>
</template>

<script>
  import Vue from 'vue'
  import data from '~/data'
  import 'swiper/dist/css/swiper.css'
  import {adminToolMixin} from '../../mixins/adminToolMixin';

  if (process.browser) {
    const VueAwesomeSwiper = require('vue-awesome-swiper/dist/ssr');
    Vue.use(VueAwesomeSwiper);
  }

  export default {
    name: 'IndexProjectsSlider',
    mixins: [adminToolMixin],
    /*props: {
      items: {
        type: Array
      }
    },*/
    data() {
      return {
        swiperOption: {
          slidesPerView: 'auto',
          grabCursor: true,
          navigation: {
            prevEl: '.gallery__control_prev',
            nextEl: '.gallery__control_next',
            disabledClass: 'gallery__control_state_disabled'
          },
          pagination: {
            el: '.swiper-pagination',
            clickable: true
          }
        },
        items: data.Projects
      }
    },
  }
</script>

<style lang="scss">
  @import "../../../assets/scss/components/swiper";

  .index-projects-slider {

    &__item {
      height: auto;
    }

    &__inner {
      height: 100%;
      padding: 25px 30px 100px 30px;
      color: #fff;
      background-color: #53696F;
      background-position: top center;
      background-repeat: no-repeat;
      background-size: cover;
    }

    &__header {
      margin-bottom: 18px;
      font-size: 26px;
      line-height: 30px;
    }

    &__img-wrap {
      height: 68px;
      margin-bottom: 20px;
      line-height: 0;
    }

    &__img {
      height: 100%;
    }

    &__text {

      & > *:first-child {
        margin-top: 0;
      }
    }

    .swiper-container {
      overflow: visible;
    }

    .gallery-nav {
      bottom: 40px;
      width: 100%;
      display: flex;
      justify-content: center;
    }

    .swiper-pagination-bullet {
      background: #fff;
    }

    @media (min-width: $screen-lg) {
      height: 100%;

      .swiper-container {
        height: 100%;
        padding-bottom: 0;
        overflow: hidden;
      }

      &__wrap,
      &__gallery {
        height: 100%;
      }

      &__inner {
        padding-top: 50px;
      }

      &__img-wrap {
        height: auto;
        max-height: 205px;
        margin-bottom: 57px;
        text-align: center;
      }

      &__img {
        height: auto;
        max-height: 100%;
      }

      &__img_dummy {
        height: 200px;
      }

      &__content {
        padding: 0 38px;
      }
    }
  }
</style>
