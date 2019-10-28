<template>
  <div class="container">
    <div class="collection__list">
      <CollectionListItem
        v-for="collection in collections"
        :key="collection.Id"
        :preview-img-src="collection.PreviewImg.ImgSrc"
        :preview-img-title="collection.PreviewImg.ImgTitle"
        :title="collection.Title"
        :products="collection.Products"
      />
    </div>
  </div>
</template>

<script>
  import data from '~/data';
  import CollectionListItem from '../CollectionListItem/CollectionListItem';
  import {adminToolMixin} from "../../mixins/adminToolMixin";

  export default {
    name: 'CollectionList',
    data() {
      return {
        collections: data.Collections.Items
      }
    },
    mixins: [adminToolMixin],
    components: {CollectionListItem},
    /*props: {
      collections: {
        type: Array,
        required: true
      },
    },*/
  }
</script>

<style lang="scss">

  .collection__list {
    display: flex;
  }

</style>
