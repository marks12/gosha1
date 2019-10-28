<template>
  <div class="container">
    <div class="products-slider">
      <div class="products-slider__gallery">
        <div v-swiper:mySwiper="swiperOption">
          <div class="swiper-wrapper">
            <div
              v-for="item in items"
              :key="item.Id"
              class="swiper-slide products-slider__item"
            >
              <ProductCard
                :previewImgSrc="item.PreviewImgSrc"
                :price="item.Price"
                :points="item.Points"
                :title="item.Title"
                :description="item.Description"
              />
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
  import Vue from 'vue';
  import ProductCard from '../ProductCard/ProductCard';
  import data from '~/data';
  import 'swiper/dist/css/swiper.css';
  import {adminToolMixin} from '../../../mixins/adminToolMixin';

  if (process.browser) {
    const VueAwesomeSwiper = require('vue-awesome-swiper/dist/ssr');
    Vue.use(VueAwesomeSwiper);
  }

  export default {
    name: 'ProductsSlider',
    mixins: [adminToolMixin],
    components: {
      ProductCard,
    },
    /*props: {
      items: {
        type: Array
      }
    },*/

    data() {
      return {
        items: data.Products.slice(0, 6),
        swiperOption: {
          products: data.Products.slice(0, 6),
          header: data.ProductsBest.SectionHeader,
          title: data.ProductsBest.SectionTitle,
          slidesPerView: 'auto',
          navigation: {
            prevEl: '.gallery__control_prev',
            nextEl: '.gallery__control_next',
            disabledClass: 'gallery__control_state_disabled'
          },
          pagination: {
            el: '.swiper-pagination',
            // dynamicBullets: true,
            // dynamicMainBullets: 1,
            clickable: true
          }
        },
      }
    },
  }
</script>

<style lang="scss">
  @import "../../../../assets/scss/components/swiper";
  @import "../../../../assets/scss/components/products-slider";
</style>
