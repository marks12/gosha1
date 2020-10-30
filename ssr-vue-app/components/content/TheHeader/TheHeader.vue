<template>
  <header class="header">
    <div class="container">
      <div class="header__inner">
        <div class="header-logo">
          <VLogo :img="require('@/assets/images/logo.svg')"/>
        </div>

        <div class="header-search">
          <VSearch :placeholder="search.placeholder"/>
        </div>

        <div class="header-second-menu">
          <ul class="header-second-menu__list">
            <li
              v-for="(item, index) in secondMenu"
              :key="index"
              class="header-second-menu__item"
            >
            <span
              class="icon"
              :class="[`icon-${item.icon}`]"
            ></span>
            </li>
          </ul>
        </div>
      </div>
      <div class="header-menu">
        <ul class="header-menu__list">
          <li
            v-for="category in catalogMenu"
            :key="category.Id"
            class="header-menu__item"
          >
            <nuxt-link
              :to="category.Url"
              class="header-menu__link"
            >
              {{ category.Title }}
            </nuxt-link>
          </li>
        </ul>
      </div>
    </div>
  </header>
</template>
<script>
  import data from '~/data.json'
  import {adminToolMixin} from '~/components/mixins/adminToolMixin';
  import VLogo from '../BaseComponents/VLogo';
  import VSearch from '../BaseComponents/VSearch';

  export default {
    name: 'Header',
    mixins: [adminToolMixin],
    components: {
      VLogo,
      VSearch
    },
    data() {
      return {
        catalogMenu: data.Header.CatalogMenu,
        secondMenu: [
          {
            icon: 'cart',
            url: '/'
          },
          {
            icon: 'user',
            url: '/'
          },
          {
            icon: 'lang-rus',
            url: '/'
          }
        ],
        search: {
          placeholder: 'Поиск по товарам, сериям, назначениям и микронутриентам'
        }
      }
    }

  }
</script>
<style lang="scss">
  @import "../../../assets/scss/components/header";
</style>
