<template>
  <transition name="modal">
    <div v-if="show" class="modal-popup">
      <div class="modal-popup__modal-mask">
        <div class="modal-popup__modal-wrapper">
          <div class="modal-popup__modal-container">
            <span class="modal-popup__close-button" @click="close()">X</span>
            <div class="modal-popup__header">
              <slot name="header"></slot>
            </div>
            <slot></slot>
            <div class="modal-popup__footer">
              <slot name="footer"></slot>
            </div>
          </div>
        </div>
      </div>
    </div>
  </transition>
</template>

<script>
  export default {
    name: 'VModalPopup',
    props: {
      show: {
        type: Boolean
      }
    },
    data() {
      return {}
    },
    methods: {
      close() {
        this.$emit('update:show', false);
      }
    }
  }
</script>

<style lang="scss">
  @import "../../../assets/scss/components/_modal-popup";
</style>
