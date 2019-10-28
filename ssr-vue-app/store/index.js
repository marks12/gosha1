import Vue from 'vue'
import Vuex from 'vuex'
import createCache from 'vuex-cache'
import api from '../../jstypes/api'
import adminTool from "./modules/adminTool";
import layoutContent from "../../jstypes/store/LayoutContent";
import pageContent from "../../jstypes/store/PageContent";
import componentTemplate from "../../jstypes/store/ComponentTemplate";
import pageInfo from "../../jstypes/store/PageInfo";
import * as myModule from "../../jstypes/apiModel";

// Этот блок убирает бесконечноый спам предупреждения "Cannot stringify arbitrary non-POJOs <ClassName>"
for(module of Object.keys(myModule)) {
  myModule[module].prototype.toJSON = function () {
    return 'function'
  };
}

api.setServerUrl('http://127.0.0.1:7005');
Vue.use(Vuex);

const store = () =>  new Vuex.Store({
  plugins: [createCache()],
  modules: {
    adminTool,
    pageContent,
    layoutContent,
    componentTemplate,
    pageInfo
  },
  state() {
    return {}
  },
  mutations: {
  },
})

export default store
