<template>
  <div class="index-promo">
    <div class="container">
      <div class="index-promo__inner">
        <Header/>
        <IndexActionsSlider :items="actions"/>

        <!--<div class="index-promo__left">

        </div>
        <div class="index-promo__right">
          <IndexProjectsSlider :items="projects" />
        </div>-->
      </div>
    </div>
  </div>
</template>

<script>
  import data from '../../../data.json';
  import IndexActionsSlider from '../IndexActionsSlider/IndexActionsSlider';
  import IndexProjectsSlider from '../IndexProjectsSlider/IndexProjectsSlider';
  import Header from '../TheHeader/TheHeader';
  import {adminToolMixin} from '../../mixins/adminToolMixin';

  export default {
    name: 'IndexPromo',
    mixins: [adminToolMixin],
    components: {
      IndexProjectsSlider,
      IndexActionsSlider,
      Header
    },
    data() {
      return {
        actions: data.Actions,
        projects: data.Projects
      }
    }
  }
</script>

<style lang="scss">
  @import "../../../assets/scss/components/index-promo";
</style>
