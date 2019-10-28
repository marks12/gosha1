<template>
  <div class="categories-set">
    <div class="container">
      <ul class="categories-set__list">
        <li
          v-for="item in items"
          :key="item.Id"
          class="categories-set__item"
        >
          <nuxt-link :to="item.Url" class="categories-set__link">
            <div class="categories-set__img-wrap">
              <img :src="item.ImgSrc" alt="">
            </div>
            <div class="categories-set__title">
              {{ item.Title }}
            </div>
          </nuxt-link>
        </li>
      </ul>
    </div>
  </div>
</template>

<script>
  import data from "~/data";
  import {adminToolMixin} from "../../mixins/adminToolMixin";

  export default {
    name: 'CategoriesList',
    mixins: [adminToolMixin],
    data() {
      return {
        items: data.CategoriesSet
      }
    },
    /*props: {
      header: {
        type: String,
        default: ''
      },
      items: {
        type: Array,
        required: false
      },
    },*/
  }
</script>

<style lang="scss">

  .categories-set {
    display: flex;
    flex-direction: column;
    padding: 20px 0;
    border-bottom: 1px solid;

    &__list {
      display: flex;
      justify-content: space-between;
    }

    &__item {
      margin: 0 10px;
    }

    &__link {
      display: flex;
      flex-direction: column;
      align-items: center;
    }

    &__img-wrap {
      width: 152px;
      height: 152px;
      background: lightgray;
    }
  }

</style>
