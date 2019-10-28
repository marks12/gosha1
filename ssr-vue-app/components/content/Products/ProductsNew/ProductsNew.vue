<template>
  <div class="products-new">
    <div class="products-new__header">{{ header }}</div>
    <div class="products-new__list">
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
  </div>
</template>

<script>
  import data from '~/data';
  import ProductCard from '../ProductCard/ProductCard';
  import {adminToolMixin} from '../../../mixins/adminToolMixin';

  export default {
    name: 'ProductsNew',
    components: {ProductCard},
    data() {
      return {
        header: data.ProductsNew.Header,
        items: data.Products.slice(0, 5)
      }
    },
    mixins: [adminToolMixin],
    /*props: {
      header: {
        type: String,
        default: ''
      },
      items: {
        type: Array,
        required: false
      }
    }*/
  }
</script>

<style lang="scss">
  .products-new {
    display: flex;
    flex-direction: column;
    align-items: center;
    border: 1px solid;

    &__list {
      display: flex;
    }
  }
</style>
