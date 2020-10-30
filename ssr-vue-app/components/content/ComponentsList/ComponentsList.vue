<template>
  <div v-if="components">
    <template v-for="component in parsedComponents" :id="component.id">
      <component v-bind:is="component.componentBlock" :key="component.id"
                 :componentId="component.id"
                 :data="component.data"
                 :seo="component.seo"
                 :id="component.id"></component>
    </template>
  </div>
</template>

<script>
  export default {
    name: 'ComponentsList',
    props: {
      components: {
        type: Array,
        require: true,
        default: function () {
          return []
        }
      },
      templates: {
        type: Array,
        require: true,
        default: function () {
          return []
        }
      },
      seo: {
        type: Object,
        require: true
      },
      data: {
        type: Object
      }
    },
    computed: {
      parsedComponents() {
        return this.parseContent(this.components);
      }
    },
    methods: {
      parseContent(pageContent) {
        let result = [];
        pageContent = pageContent.sort(function (a, b) {
          return a.Position - b.Position;
        });
        for (let content of pageContent) {
          let template = this.templates.find(x => x.Id === content.ComponentTemplateId);

          if (template && content.IsActive) {
            result.push({
              id: content.Id.toString(),
              code: template.Code,
              path: template.Path,
              data: this.data,
              seo: this.seo,
              componentBlock: () => import(`~/components/content/${template.Path}`)
            });
          }
        }
        return result;
      }
    }
    ,
  }
</script>

<style scoped>

</style>
