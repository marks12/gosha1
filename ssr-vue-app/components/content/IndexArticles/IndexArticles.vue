<template>
  <div class="index-articles">
    <ul class="index-articles__list clearfix">
      <li
        v-for="item in items"
        :key="item.Id"
        class="index-articles__item"
        :class="{ 'index-articles__item_small': item.View === 'small' }"
      >
        <nuxt-link :to="item.Url" class="index-articles__link">
          <span>{{ item.Label }}</span>
          {{ item.Title }}

          <template v-if="item.View === 'small'">
            <span>{{ item.Date }}</span>
          </template>
        </nuxt-link>
      </li>
    </ul>
  </div>
</template>

<script>
  import data from '~/data';
  import {adminToolMixin} from '../../mixins/adminToolMixin';

  export default {
    name: 'IndexArticles',
    data() {
      return {
        items: data.IndexArticles,
      }
    },
    mixins: [adminToolMixin],
    /*props: {
      items: {
        type: Array,
        required: true
      },
    },*/
  }
</script>

<style lang="scss">

  .index-articles {
    width: 1434px;
    margin: 0 auto;
    padding: 20px 0;
    border-bottom: 1px solid;

    &__list {
      margin: 0 -15px;
    }

    &__item {
      float: left;
      width: 458px;
      height: 733px;
      margin: 0 15px;
      background: lightgray;

      &_small {
        height: 350px;

        & + & {
          margin-top: 30px;
        }
      }
    }

    &__link {
      display: flex;
      flex-direction: column;
    }
  }

</style>
