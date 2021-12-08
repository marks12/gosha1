import Vue from 'vue'
import Vuex from 'vuex'
import projectInfo from "../../webapp/jstypes/store/ProjectInfo";
import entity from "../../webapp/jstypes/store/Entity";
import fieldType from "../../webapp/jstypes/store/FieldType";
import currentApp from "../../webapp/jstypes/store/CurrentApp";
import panel from "./store/panel";
import request from "./store/request";
import response from "./store/response";

Vue.use(Vuex);

const gosha = {
    namespaced: true,
    modules: {
        projectInfo,
        entity,
        fieldType,
        currentApp,
        panel,
        request,
        response,
    },
};

const store = {
    modules: {
        gosha,
    },
    state: {
        version: "1.0.0", // a simple property
    },
};

export default new Vuex.Store(store);