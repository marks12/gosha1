package cmd

const storeTemplate = `
import {{Entity}} from "../apiModel";
import api from "../api";

let findUrl = "/api/v1/{entity}";
let readUrl = "/api/v1/{entity}/"; // + id
let createUrl = "/api/v1/{entity}";
let updateUrl = "/api/v1/{entity}/"; // + id
let deleteUrl = "/api/v1/{entity}/"; // + id

const {entity} = {
    actions: {
        create{Entity}(context, {data, filter, header}) {

            return api.create(createUrl, data, filter, header)
                .then(function(response) {

                    context.commit("set{Entity}", response.Model);

                    return response;
                })
                .catch(function(err) {
                    return err;
                });
        },
        delete{Entity}(context, {id, header}) {

            return api.remove(deleteUrl + id, header)
                .then(function(response) {
                    context.commit("clear{Entity}");
                    return response;
                })
                .catch(function(err) {
                    return err;
                });
        },
        find{Entity}(context, {filter, header}) {

            return api.find(findUrl, filter, header)
                .then(function(response) {

                    context.commit("set{Entity}List", response.List);

                    return response;
                })
                .catch(function(err) {
                    return err;
                });
        },
        load{Entity}(context, {id, filter, header}) {

            return api.find(readUrl + id, filter, header)
                .then(function(response) {

                    context.commit("set{Entity}", response.Model);

                    return response;
                })
                .catch(function(err) {
                    return err;
                });
        },
        update{Entity}(context, {id, data, filter, header}) {

            return api.update(updateUrl + id, data, filter, header)
                .then(function(response) {

                    context.commit("set{Entity}", response.Model);

                    return response;
                })
                .catch(function(err) {
                    return err;
                });
        },
    },
    getters: {
        get{Entity}: (state) => {
            return state.{Entity};
        },
        get{Entity}List: (state) => {
            return state.{Entity}List;
        },
    },
    mutations: {
        set{Entity}(state, data) {
            state.{Entity} = data;
        },
        set{Entity}List(state, data) {
            state.{Entity}List = data;
        },
        clear{Entity}(state) {
            state.{Entity} = new {Entity}();
        },
    },
    state: {
        {Entity}: new {Entity}(),
        {Entity}List: [],
    },
};

export default {entity};
`
