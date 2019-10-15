import Vue from 'vue'
import Vuex from 'vuex'
import projectInfo from "../../webapp/jstypes/store/ProjectInfo";
import entity from "../../webapp/jstypes/store/Entity";
import fieldType from "../../webapp/jstypes/store/FieldType";
import currentApp from "../../webapp/jstypes/store/CurrentApp";

Vue.use(Vuex);

const store = {
    modules: {
        projectInfo,
        entity,
        fieldType,
        currentApp
    },
    state: {
        version: "1.0.0", // a simple property
    },
};

export default new Vuex.Store(store);