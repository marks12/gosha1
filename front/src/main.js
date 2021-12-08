import Vue from 'vue'
import App from './App.vue'
import router from './router'
import store from './store'
import api from '../../webapp/jstypes/api'

Vue.config.productionTip = false;

if (process && process.env && process.env.NODE_ENV === "development") {
  api.setServerUrl('http://' + window.location.hostname + ':4343');
} else {
  api.setServerUrl('');
}

new Vue({
  router,
  store,
  render: h => h(App)
}).$mount('#app');

Vue.filter('substring', function (value, limit) {
  if (value.length > limit) {
    value = value.substring(0, (limit - 3)) + '...';
  }
  return value
})