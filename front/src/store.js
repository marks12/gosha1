import Vue from 'vue'
import Vuex from 'vuex'
import projectInfo from "../../webapp/jstypes/store/ProjectInfo";
import entity from "../../webapp/jstypes/store/Entity";

Vue.use(Vuex);

const store = {
    modules: {
        projectInfo,
        entity,
    },
    state: {
        version: "1.0.0", // a simple property
    },
};

export default new Vuex.Store(store);