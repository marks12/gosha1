<template>
  <div class="product-sheet">
    <div class="container">
      <div class="product-sheet__inner">
        <div class="product-sheet__left">
          <ProductImageSlider :imageSet="imageSet"/>
        </div>
        <div class="product-sheet__right">
          <div class="product-info">
            <div class="product-info__title">
              <h1>{{title}}</h1>
            </div>
            <div class="product-info__description">{{description}}</div>
            <div class="product-info__price-block">
              <div class="product-info__old-price">{{discountPrice}}р.</div>
              <div class="product-info__price">
                {{price}}р.
                <div class="product-info__points">{{points}} баллов</div>
              </div>
            </div>
            <div class="product-info__button-block">
              <div class="product-info__top_row">
                <div class="product-info__add-to-cart-button">
                  <VButton v-if="countOfOrder === 0" @click="addToCart()" :value="'Добавить в корзину'"></VButton>
                  <VCartButton v-else :value.sync="countOfOrder"></VCartButton>
                </div>
                <div class="product-info__pickup">
                  <p>Самовывоз сегодня {{countOfOrder}} шт.</p>
                  <p>
                    <VLink @click="deliverySelectorShow = true" :value="'с Ленина 21'"/>
                  </p>
                </div>
                <div class="product-info__delivery">
                  <p>Доставка завтра</p>
                  <p>
                    <VLink @click="deliverySelectorShow = true" :value="'по Новосибирску'"/>
                  </p>
                </div>
              </div>
              <div class="product-info__bottom_row">
                <div class="product-info__quick-order-button">
                  <VButton :disabled="countOfOrder === 0" :value="'Быстрый заказ'"></VButton>
                </div>
                <div class="product-info__payment">
                  <p>Оплата онлайн банковской картой</p>
                  <p>
                    <VLink @click="paymentSelectorShow = true" :value="'Другие варианты'"/>
                  </p>
                </div>
              </div>
            </div>
          </div>
        </div>
      </div>
    </div>

    <VModalPopup :show.sync="deliverySelectorShow">
      <template v-slot:header>Выберите способ получения</template>
      <template>
        <DeliverySelector/>
      </template>
      <template v-slot:footer>
        <VButton @click="deliverySelectorShow = false" :value="'Выбрать'"/>
      </template>
    </VModalPopup>

    <VModalPopup :show.sync="paymentSelectorShow">
      <template v-slot:header>Выберите способ оплаты</template>
      <template>
        <PaymentSelector/>
      </template>
      <template v-slot:footer>
        <VButton @click="paymentSelectorShow = false" :value="'Выбрать'"/>
      </template>
    </VModalPopup>
  </div>
</template>

<script>
  import data from '~/data';
  import VButton from '../../BaseComponents/VButton';
  import VLink from '../../BaseComponents/VLink';
  import VCartButton from '../../BaseComponents/VCartButton';
  import ProductImageSlider from '../ProsuctImageSlider/ProductImageSlider';
  import VModalPopup from '../../BaseComponents/VModalPopup';
  import DeliverySelector from '../../DeliverySelector/DeliverySelector';
  import PaymentSelector from '../../PaymentSelector/PaymentSelector';
  import {adminToolMixin} from '../../../mixins/adminToolMixin';

  export default {
    name: 'ProductSheet',
    mixins: [adminToolMixin],
    props: {
      seo: {
        type: Object,
        require: true
      },
      data: {
        type: Object,
        default: function () {
          return {};
        }
      }
    },
    components: {
      ProductImageSlider,
      VButton,
      VLink,
      VCartButton,
      VModalPopup,
      DeliverySelector,
      PaymentSelector
    },
    data() {
      return {
        imageSet: data.Products.map(x => x.PreviewImgSrc),
        price: data.Products[0].Price,
        points: data.Products[0].Points,
        deliverySelectorShow: false,
        paymentSelectorShow: false,
        countOfOrder: 0,
      }
    },
    mounted() {
    },
    computed: {
      discountPrice() {
        return this.price / 100 * 200;
      },
      title() {
        return this.data && this.data.Name ? this.data.Name : "Тестовый заголовок";
      },
      description() {
        return this.data && this.data.Description ? this.data.Description : "Тестовое описание" ;
      }
    },
    methods: {
      addToCart() {
        this.countOfOrder++;
      }
    }
  }
</script>

<style lang="scss">
  @import "../../../../assets/scss/components/_product-sheet";
</style>
