<template>
  <div class="product__list">
    <ProductCard
      v-for="item in items"
      :key="item.Id"
      :previewImgSrc="item.PreviewImgSrc"
      :price="item.Price"
      :points="item.Points"
      :title="item.Title"
      :description="item.Description"
    />
  </div>
</template>

<script>
  import data from '~/data';
  import ProductCard from '../ProductCard/ProductCard';
  import {adminToolMixin} from '../../../mixins/adminToolMixin';

  export default {
    name: 'ProductList',
    data() {
      return {
        items: data.Products.slice(0, 4)
      }
    },
    mixins: [adminToolMixin],
    components: {ProductCard},
    /*props: {
      items: {
        type: Array,
        required: true
      }
    }*/
  }
</script>

<style lang="scss">
  .product {

    &__list {
      display: flex;
    }
  }
</style>
