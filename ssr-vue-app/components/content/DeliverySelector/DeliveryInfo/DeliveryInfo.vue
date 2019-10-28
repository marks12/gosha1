<template>
  <div class="delivery-info" @click="click()" v-bind:class="{ 'delivery-info_selected': selected }">
    <div class="delivery-info__date">
      <div class="delivery-info__day">{{day}}</div>
      <div class="delivery-info__day-of-month">{{dateOfMonth}}</div>
    </div>
    <div class="delivery-info__place">
      <div class="delivery-info__address">{{address}}</div>
      <div class="delivery-info__count-of-product">{{productAvailable}}</div>
    </div>
    <div class="delivery-info__price">{{deliveryPrice}}</div>
  </div>
</template>

<script>
  import 'moment/locale/ru';
  import 'moment/min/locales.min';

  export default {
    name: 'DeliveryInfo',
    props: {
      id: {
        type: Number,
        require: true
      },
      date: {
        type: Date,
        require: true
      },
      address: {
        type: String,
        require: true
      },
      countOfProduct: {
        type: Number
      },
      dateOfMonth: {
        type: String,
        require: true
      },
      price: {
        type: Number
      },
      selected: {
        type: Boolean
      }
    },
    computed: {
      day() {
        let currentDate = this.$moment(this.$moment().format('LL'));
        let date = this.$moment(this.date);
        let duration = this.$moment.duration(date.diff(currentDate));
        let dayOfDuration = duration.get('days');
        if (dayOfDuration < 0) {
          return 'Не доставляется';
        }
        if (dayOfDuration === 0) {
          return 'Сегодня';
        } else if (dayOfDuration === 1) {
          return 'Завтра';
        }
        return `Через ${dayOfDuration} ${this.declOfNum(dayOfDuration)}`;
      },
      productAvailable() {
        let result = `${this.countOfProduct} в наличии`;
        if (!this.countOfProduct) {
          result = 'Под заказ'
        }
        return result;
      },
      deliveryPrice() {
        let result = `${this.price} р.`;
        if (!this.price) {
          result = 'Бесплатно'
        }
        return result;
      }
    },
    methods: {
      declOfNum(number) {
        let text_forms = ['день', 'дня', 'дней'];
        number = Math.abs(number) % 100;
        let decimal = number % 10;
        if (number > 10 && number < 20) { return text_forms[2]; }
        if (decimal > 1 && decimal < 5) { return text_forms[1]; }
        if (decimal === 1) { return text_forms[0]; }
        return text_forms[2];
      },
      click() {
        this.$emit('click', this.id);
      }
    }
  }
</script>

<style lang="scss">
  @import "../../../../assets/scss/_variables";

  .delivery-info {
    display: flex;
    justify-content: space-between;
    padding: 14px;

    &:hover {
      cursor: pointer;
    }

    &_selected {
      border-radius: $border-radius-base;
      box-shadow: inset 0 0 0 1px $color-weak-light;
      cursor: pointer;
    }

    &__date {
      width: 30%;
    }

    &__day {
      margin-bottom: $margin;
    }

    &__day-of-month {
      color: $color-weak;
    }

    &__place {
      width: 50%;
      flex-shrink: 0;
    }

    &__address {
      margin-bottom: $margin;
    }

    &__count-of-product {
      color: $color-weak;
    }

    &__price {
      text-align: right;
      width: 20%;
      font-weight: 800;
    }
  }

</style>
