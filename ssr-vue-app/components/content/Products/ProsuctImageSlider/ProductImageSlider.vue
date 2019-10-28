<template>
  <div class="product-images">
    <div class="product-images__main-image">
      <img :src="imageSrc">
    </div>
    <div class="product-images__images-slider">
      <div class="product-images__swiper-button_prev"></div>
      <div class="product-images__swiper">
        <div v-swiper:mySwiper="swiperOption">
          <div class="swiper-wrapper">
            <div @click="clickOnImage(item)" class="swiper-slide" v-for="(item, index) in imageSet" :key="index">
              <img :src="item" alt="index">
            </div>
          </div>
        </div>
      </div>
      <div class="product-images__swiper-button_next"></div>
    </div>
  </div>
</template>

<script>
  import Vue from 'vue';
  import 'swiper/dist/css/swiper.css';

  if (process.browser) {
    const VueAwesomeSwiper = require('vue-awesome-swiper/dist/ssr');
    Vue.use(VueAwesomeSwiper);
  }

  export default {
    name: 'ProductImageSlider',
    data() {
      return {
        imageSrc: '',
        swiperOption: {
          slidesPerView: 5,
          spaceBetween: 14,
          centerInsufficientSlides: true,
          navigation: {
            prevEl: '.product-images__swiper-button_prev',
            nextEl: '.product-images__swiper-button_next',
            disabledClass: 'gallery__control_state_disabled'
          },
        },
      }
    },
    props: {
      imageSet: {
        type: Array,
        require: true
      },
    },
    mounted() {
      this.imageSrc = this.imageSet[0];
    },
    computed: {
      discountPrice() {
        return this.price / 100 * 200;
      }
    },
    methods: {
      clickOnImage(src) {
        this.imageSrc = src;
      },
      addToCart() {
        this.countOfOrder++;
      }
    }
  }
</script>

<style lang="scss">
  @import "../../../../assets/scss/components/swiper";
  @import "../../../../assets/scss/_variables";

  $arrow-size: 10px;

  .product-images {

    &__main-image {
      align-items: center;
      justify-content: center;
      flex-basis: auto;
      height: 300px;
      text-align: center;
      background: $color-soft-light;
      margin: $margin;
    }

    &__images-slider {
      display: flex;
      align-items: center;
      justify-content: center;
      flex-basis: auto;
    }

    &__swiper {
      width: calc(100% - #{$margin * 2});

      .swiper-slide {
        background: $color-soft-light;
      }
    }

    &__swiper-button_prev, &__swiper-button_next {
      border: solid $color-base-light;
      border-width: 0 $border-radius-small $border-radius-small 0;
      display: inline-block;
      padding: $arrow-size;
      outline: none;
      cursor: pointer;
      z-index: 10;
    }

    &__swiper-button_prev {
      transform: rotate(135deg);
      -webkit-transform: rotate(135deg);
    }

    &__swiper-button_next {
      transform: rotate(-45deg);
      -webkit-transform: rotate(-45deg);
    }
  }
</style>
