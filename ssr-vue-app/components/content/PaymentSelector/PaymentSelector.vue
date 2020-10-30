<template>
  <div class="payment-selector">
    <div class="payment-selector__payment-type-list">
      <div @click="clickOnPayment('visa')" class="payment-type payment-type_visa">
        <img src="/images/payment-systems/visa.png">
        <div class="payment-type__description_visa">
          Банковской картой
        </div>
      </div>
      <div @click="clickOnPayment('qiwi')" class="payment-type">
        <img src="/images/payment-systems/qiwi.png">
      </div>
      <div @click="clickOnPayment('yandex_money')" class="payment-type">
        <img src="/images/payment-systems/yandex_money.png">
      </div>
      <div @click="clickOnPayment('prom_bank')" class="payment-type">
        <img src="/images/payment-systems/prom_bank.png">
      </div>
      <div @click="clickOnPayment('russkii_standart')" class="payment-type">
        <img src="/images/payment-systems/russkii_standart.png">
      </div>
      <div @click="clickOnPayment('pay_pal')" class="payment-type">
        <img src="/images/payment-systems/pay_pal.png">
      </div>
      <div @click="clickOnPayment('web_money')" class="payment-type">
        <img src="/images/payment-systems/web_money.png">
      </div>
      <div @click="clickOnPayment('cash')" class="payment-type">
        <img src="/images/payment-systems/vtb24.png">
      </div>
      <div @click="clickOnPayment('alfa_bank')" class="payment-type">
        <img src="/images/payment-systems/alfa_bank.png">
      </div>
      <div @click="clickOnPayment('cash')" class="payment-type payment-type_cash">
        <img src="/images/payment-systems/cash.png">
        <div class="payment-type__description_cash">
          Наличными при получении
        </div>
      </div>
    </div>
    <hr class="payment-selector__line">
    <div class="payment-selector__payment-additional">
      <div class="payment-selector__row-wrapper">
        <div class="payment-selector__input_label">Оплата бонусами</div>
        <div class="payment-selector__input">
          <input type="number" value="0">
          <div class="payment-selector__input-description">Доступно: 205 б</div>
        </div>
        <div class="payment-selector__description">Можно оплатить до 100% корзины (1500р.)</div>
      </div>
      <div class="payment-selector__row-wrapper">
        <div class="payment-selector__input_label">Купить в рассорчку</div>
        <div class="payment-selector__input">
          <input type="number" value="0">
          <div class="payment-selector__input-description">Доступно: 1875р</div>
        </div>
        <div class="payment-selector__description">В конце периода будет списано с вашего чека</div>
      </div>
    </div>
  </div>
</template>

<script>
  export default {
    name: 'PaymentSelector',
    methods: {
      clickOnPayment(paymentType) {
        this.$emit('clickOnPayment', paymentType);
      }
    }
  }
</script>

<style lang="scss">
  @import "../../../assets/scss/components/_payment-selector";
</style>
