<template>
  <div class="section-text-with-buttons">
    <div class="container">
      <div class="section-text-with-buttons__inner">
        <div class="section-text-with-buttons__сontent">
          <!-- slot слот для контента -->
          <slot name="content">
            <div v-if="header" class="section-text-with-buttons__header">{{ header }}</div>
            <div v-if="title" class="section-text-with-buttons__title">{{ title }}</div>
          </slot>
        </div>
        <div class="section-text-with-buttons__action">
          <!-- slot слот для кнопок -->
          <slot name="action">
            <template
              v-for="(button, index) in buttons"
            >
              <nuxt-link
                v-if="button.IsShow"
                :key="index"
                class="button"
                :class="{ 'button_accent': button.IsAccent}"
                :to="button.Url"
              >{{ button.Text }}
              </nuxt-link>
            </template>
          </slot>
        </div>
      </div>
    </div>
  </div>
</template>

<script>
  import {adminToolMixin} from '../../mixins/adminToolMixin';
  import data from '~/data';

  export default {
    name: 'SectionTextWithButtons',
    mixins: [adminToolMixin],
    data() {
      return {
        header: data.CallToRegistration.Header,
        title: data.CallToRegistration.Title,
        buttons: data.CallToRegistration.Buttons,
      }
    }
    /*props: {
      header: {
        type: String,
        default: ''
      },
      title: {
        type: String,
        default: ''
      },
      buttons: {
        type: Array,
        default() {
          return [];
        }
      }
    },*/
  }
</script>

<style lang="scss">
  @import "../../../assets/scss/components/section-text-with-buttons";
</style>
