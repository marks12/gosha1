<template>
  <section class="content-index">
    <ComponentsList :components="componentsBefore" :templates="templates" :data="this.pageInfo.Data" :seo="this.pageInfo.SeoMeta"/>
    <ComponentsList :components="components" :templates="templates" :data="this.pageInfo.Data" :seo="this.pageInfo.SeoMeta"/>
    <ComponentsList :components="componentsAfter" :templates="templates" :data="this.pageInfo.Data" :seo="this.pageInfo.SeoMeta"/>
  </section>
</template>

<script>
  import {mapActions, mapGetters} from 'vuex';
  import {PageInfoFilter} from '../../jstypes/apiModel';
  import ComponentsList from '~/components/content/ComponentsList/ComponentsList';
  import VSection from '../components/content/BaseComponents/VSection';
  import ProductSheet from '../components/content/Products/ProductSheet/ProductSheet';
  import DeliverySelector from '../components/content/DeliverySelector/DeliverySelector';

  export default {
    computed: {
      ...mapGetters({
        pageInfo: 'getListPageInfo',
      }),
      components() {
        return this.pageInfo.PageContent;
      },
      componentsBefore() {
        if (!this.pageInfo || !this.pageInfo.LayoutContent) return [];
        let index = this.pageInfo.LayoutContent.findIndex(x => x.ComponentTemplateId === 18);
        return this.pageInfo.LayoutContent.slice(0, index);
      },
      componentsAfter() {
        if (!this.pageInfo || !this.pageInfo.LayoutContent) return [];
        let index = this.pageInfo.LayoutContent.findIndex(x => x.ComponentTemplateId === 18);
        return this.pageInfo.LayoutContent.slice(index + 1, this.pageInfo.LayoutContent.length);
      },
      templates() {
        return this.pageInfo.Components;
      }
    },
    components: {
      ComponentsList,
      DeliverySelector,
      ProductSheet,
      VSection
    },
    data() {
      return {
        PageInfoFilter: new PageInfoFilter(),
      }
    },
    methods: {
      ...mapActions([
        'findPageInfo',
      ])
    },
    mounted() {
      this.$socket.bindType('AdminSessionResponse', (result) => {
        if (result.success) {
          this.$store.commit('enableAdminTool');
        }
      });

      let url = new URL(window.location);
      this.PageInfoFilter.Chpu = url.pathname;
      let pageTemplateId = url.searchParams.get('page-template-id');
      if (pageTemplateId) {
        this.PageInfoFilter.PageTemplateId = pageTemplateId;
      }
      this.PageInfoFilter.CurrentPage = 1;
      this.PageInfoFilter.PerPage = 500;
      this.findPageInfo({filter: this.PageInfoFilter});
    }
  }
</script>

<style lang="scss">

</style>

