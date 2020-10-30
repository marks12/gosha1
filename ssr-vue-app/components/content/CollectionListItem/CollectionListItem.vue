<template>
  <div class="collection__item">
    <div class="collection-preview">
      <div class="collection-preview__img-wrap">
        <img src="" class="collection-preview__img" alt="">
      </div>
    </div>
    <div class="collection-products">
      <ul class="collection-products__list">
        <li
          v-for="product in products"
          :key="product.Id"
          class="collection-products__item"
        >
          <div class="collection-products__img-wrap">
            <img :src="product.ImgSrc" alt="" class="collection-products__img">
          </div>
        </li>
      </ul>
    </div>
    <div v-if="title" class="collection__title">{{ title }}</div>
  </div>
</template>

<script>
  import {adminToolMixin} from "../../mixins/adminToolMixin";

  export default {
    name: 'CollectionListItem',
    mixins: [adminToolMixin],
    props: {
      previewImgSrc: {
        type: String,
        required: true
      },
      previewImgTitle: {
        type: String
      },
      title: {
        type: String,
        default: ''
      },
      products: {
        type: Array,
        required: true
      }
    }
  }
</script>

<style lang="scss">

  .collection {

    &__item {
      margin: 0 20px;
    }

    &__title {
      text-align: center;
    }
  }

  .collection-preview {

    &__img-wrap {
      text-align: center;
    }

    &__img {
      display: inline-block;
      width: 150px;
      height: 150px;
      background: lightgray;
    }
  }

  .collection-products {

    &__list {
      display: flex;
      justify-content: center;
    }

    &__item {
      margin: 0 10px;
    }

    &__img-wrap {
      display: inline-block;
      width: 30px;
      height: 30px;
      background: lightgray;
    }
  }

</style>
