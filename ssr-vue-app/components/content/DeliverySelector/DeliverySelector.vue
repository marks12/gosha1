<template>
  <div class="delivery-selector">
    <div class="delivery-selector__wrapper">
      <div class="delivery-selector__left">
        <div class="delivery-selector__courier-delivery">
          <div class="delivery-list">
            <div class="delivery-list__wrapper">
              <div class="delivery-list__title">Курьером</div>
              <div class="delivery-list__description">При заказе от 1500р - бесплатно</div>
            </div>
            <DeliveryInfo
              @click="selectDelivery(0)"
              :id="0"
              :selected="selectedDeliveryId === 0"
              :date="new Date(2019, 8, 14)"
              :dateOfMonth="'14 сентября'"
              :price="500"
              :address="'В любую точку Новосибирска'"
              :count-of-product="5"/>
          </div>
        </div>
        <div class="delivery-selector__pickup">
          <div class="delivery-list">
            <div class="delivery-list__wrapper">
              <div class="delivery-list__title">Самовывоз</div>
              <div class="delivery-list__description"></div>
            </div>
            <DeliveryInfo v-for="point in pointsOfSale" :key="point.id"
                          @click="selectDelivery(point.id)"
                          :id="point.id"
                          :selected="selectedDeliveryId === point.id"
                          :date="point.date"
                          :dateOfMonth="point.dateOfMonth"
                          :address="point.address"
                          :price="point.price"
                          :count-of-product="point.countOfProduct"/>
          </div>
        </div>
      </div>
      <div class="delivery-selector__right">
        <div class="map">
          Тут будет карта
        </div>
      </div>
    </div>
  </div>
</template>

<script>
  import DeliveryInfo from './DeliveryInfo/DeliveryInfo';
  import VButton from '../BaseComponents/VButton';

  export default {
    name: 'DeliverySelector',
    components: {
      DeliveryInfo,
      VButton
    },
    data() {
      return {
        selectedDeliveryId: 0,
        pointsOfSale: [
          {
            id: 1,
            date: new Date(2019, 8, 12),
            dateOfMonth: '12 сентября',
            address: 'Ленина 21',
            countOfProduct: 5
          },
          {
            id: 2,
            date: new Date(2019, 8, 13),
            dateOfMonth: '13 сентября',
            price: 500,
            address: 'Комсомольского проспект 22',
            countOfProduct: 5
          },
          {
            id: 3,
            date: new Date(2019, 8, 14),
            dateOfMonth: '14 сентября',
            address:'Проспект Дзержинского',
            countOfProduct: 5
          },
          {
            id: 4,
            date: new Date(2019, 8, 15),
            dateOfMonth: '15 сентября',
            address: 'Ленина 23',
            countOfProduct: 5
          },
          {
            id: 5,
            date: new Date(2019, 8, 15),
            dateOfMonth: '15 сентября',
            address: 'Ленинградская слобода 265',
            countOfProduct: 5
          },
        ]
      }
    },
    methods: {
      selectDelivery(selectedId) {
        this.selectedDeliveryId = selectedId;
      }
    }
  }
</script>

<style lang="scss">
  @import "../../../assets/scss/components/_delivery-selector";
</style>
