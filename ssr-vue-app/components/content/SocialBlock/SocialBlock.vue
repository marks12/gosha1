<template>
  <div class="social">
    <ul class="social__list">
      <li
        v-for="item in items"
        :key="item"
        class="social__item"
      >
        <a href="#" class="social__link">
          <VIcon :name="item"/>
        </a>
      </li>
    </ul>
  </div>
</template>

<script>

  import VIcon from '../BaseComponents/VIcon';
  import {adminToolMixin} from '../../mixins/adminToolMixin';

  export default {
    name: 'SocialBlock',
    mixins: [adminToolMixin],
    components: {VIcon},
    data() {
      return {
        items: [
          'vk',
          'whatsapp',
          'instagram',
          'viber'
        ]
      }
    }
  }
</script>

<style lang="scss">

  .social {

    &__list {
      display: flex;
      flex-wrap: wrap;
    }

    &__item {
      flex-shrink: 0;
      margin-right: 16px;
    }

    &__link {
      display: inline-block;
      width: 20px;
      height: 20px;
      color: #fff;
      background: $color-selection;
    }
  }

</style>
