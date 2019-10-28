<template>
  <div class="product">
    <nuxt-link
      :to="'/'"
      class="product__link"
    >
      <img
        class="product__img"
        :src="previewImgSrc"
        :alt="title"
      >
      <div class="product__info">
        <div class="product__coast">
          <div class="product__price">
            <span class="product__price-value">{{ formattedPrice }}</span> р.
          </div>
          <div class="product__points">{{ formattedPoints }} б.</div>
        </div>
        <div class="product__title">{{ title }}</div>
        <div class="product__description">{{ description }}</div>
      </div>
    </nuxt-link>
  </div>
</template>

<script>
  const formatter = new Intl.NumberFormat('ru');

  export default {
    name: 'ProductCard',
    props: {
      previewImgSrc: {
        type: String,
      },
      price: {
        type: Number,
      },
      points: {
        type: Number,
      },
      title: {
        type: String,
        required: true
      },
      description: {
        type: String,
      },
      discount: {
        type: Number
      }
    },

    computed: {
      formattedPrice() {
        return formatter.format(this.price);
      },

      formattedPoints() {
        return formatter.format(this.points);
      }
    }
  }
</script>

<style lang="scss">
  @import "../../../../assets/scss/components/product-card";
</style>
