<template>
  <footer class="footer">
    <div class="footer-main">
      <div class="container">
        <div class="footer-main__inner">
          <div class="footer-main__left">
            <FooterAdvantages/>
            <div class="footer-menu">
              <div class="footer-menu__inner">
                <div
                  v-for="section in footerMenu"
                  :key="section.SectionTitle"
                  class="footer-menu__section"
                >
                  <div class="footer-menu__section-title">{{ section.SectionTitle }}</div>
                  <div
                    v-for="(category, index) in section.Groups"
                    :key="index"
                    class="footer-menu__category">
                    <ul class="footer-menu__list">
                      <li
                        v-for="(item, index) in category.Items"
                        :key="index"
                        class="footer-menu__item"
                      >
                        <nuxt-link
                          :to="item.Url"
                          class="footer-menu__link"
                        >
                          {{ item.Title }}
                        </nuxt-link>
                      </li>
                    </ul>
                  </div>
                </div>
              </div>
            </div>
          </div>
          <div class="footer-main__right">
            <FooterProjects :items="projects"/>
          </div>
        </div>
      </div>
    </div>

    <div class="footer-bottom">
      <div class="container">
        <div class="footer-info">
          <div class="footer-info__left">
            <div class="footer-info-menu">
              <ul class="footer-info-menu__list">
                <li
                  v-for="item in infoMenu"
                  :key="item.Id"
                  class="footer-info-menu__item">
                  <nuxt-link :to="item.Link" class="footer-info-menu__link">
                    {{ item.Title }}
                  </nuxt-link>
                </li>
              </ul>
            </div>
          </div>
          <div class="footer-info__right">
            <div class="footer-social">
              <SocialBlock/>
            </div>
          </div>
        </div>
        <div class="footer-bottom__main">
          <div class="footer-bottom__main-left">
            <div class="footer-contacts">
              <div class="footer-contacts__logo">
                <VLogo :img="require('@/assets/images/logo-icon.svg')"/>
              </div>
              <div class="footer-contacts__content">
                <a href="tel:88007558755" class="footer-contacts__phone">
                  {{ contacts.phone }}
                </a>
                <div class="footer-contacts__title">
                  {{ contacts.title }}
                </div>
              </div>
            </div>
          </div>
          <div class="footer-bottom__main-right">
            <div class="footer-copyright" v-html="copyright"></div>
          </div>
        </div>
      </div>
    </div>
  </footer>
</template>

<script>
  import data from '../../../data.json';
  import FooterProjects from './FooterProjects/FooterProjects';
  import FooterAdvantages from './FooterAdvantages/FooterAdvantages';
  import SocialBlock from '../SocialBlock/SocialBlock'
  import VLogo from '../BaseComponents/VLogo';
  import {adminToolMixin} from '../../mixins/adminToolMixin';

  export default {
    name: 'Footer',
    mixins: [adminToolMixin],
    components: {
      SocialBlock,
      FooterAdvantages,
      FooterProjects,
      VLogo
    },
    data() {
      return {
        advantages: data.Footer.Advantages,
        copyright: data.Footer.Copyright,
        contacts: {
          phone: data.Footer.Contacts.Phone,
          title: data.Footer.Contacts.Title
        },
        projects: data.Footer.Projects,
        footerMenu: data.Footer.MainMenu,
        infoMenu: data.Footer.InfoMenu
      }
    }
  }
</script>

<style lang="scss">
  @import "../../../assets/scss/components/footer";
</style>
