<template>
  <div class="cart-button">
    <div class="cart-button__counter">
      <div @click="increment()" class="cart-button__button-increment">-</div>
      <div class="cart-button__value">
        {{value}}
      </div>
      <div @click="decrement()" class="cart-button__button-decrement">+</div>
    </div>
    <div class="cart-button__cart-label">
      <VLink :value="'в корзине'"></VLink>
    </div>
  </div>
</template>

<script>
  import VLink from './VLink';

  export default {
    name: 'VCartButton',
    components: {
      VLink
    },
    props: {
      value: {
        type: Number
      }
    },
    methods: {
      decrement() {
        this.$emit('update:value', this.value + 1);
      },
      increment() {
        if(this.value > 0) {
          this.$emit('update:value', this.value - 1);
        }
      }
    }
  }
</script>

<style lang="scss">
  @import "../../../assets/scss/_variables";

  .cart-button {
    white-space: nowrap;
    border-radius: $border-radius-button;
    background: $color-soft-light;
    padding: 0px 18px 0px 0;
    width: 200px;
    display: flex;
    flex-direction: row;
    align-items: center;

    &__counter {
      background: white;
      padding: 12px 18px 12px 18px;
      border-radius: $border-radius-button;
      border: 2px solid $color-selection;
      width: 60%;

      display: flex;
      flex-direction: row;
      justify-content: space-between;
      flex-grow: 2;
    }

    &__button-increment {
      cursor: pointer;
    }

    &__value {
    }

    &__button-decrement {
      cursor: pointer;
    }

    &__cart-label a {
      text-align: right;
      display: inline-block;
      padding-left: 6px;
      flex-grow: 1;
      color: $color-weak;
      text-decoration: underline;

    }

  }

</style>
