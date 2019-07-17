package cmd

const storeTemplate = `
import {{Entity}} from "../apiModel";
import api from "../api";
import {findItemIndex} from "../common";

let findUrl = "/api/v1/{entity}";
let readUrl = "/api/v1/{entity}/"; // + id
let createUrl = "/api/v1/{entity}";
let updateUrl = "/api/v1/{entity}/"; // + id
let deleteUrl = "/api/v1/{entity}/"; // + id
let findOrCreateUrl = "/api/v1/{entity}"; // + id

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
        findOrCreate{Entity}(context, {id, data, filter, header}) {

            return api.update(findOrCreateUrl, data, filter, header)
                .then(function(response) {

                    context.commit("set{Entity}", response.Model);

                    return response;
                })
                .catch(function(err) {
                    return err;
                });
        },
        clearList{Entity}(context) {
            context.commit("clearList{Entity}");
        },
    },
    getters: {
        get{Entity}: (state) => {
            return state.{Entity};
        },
        getList{Entity}: (state) => {
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
        clearList{Entity}(state) {
            state.{Entity}List = [];
        },
		update{Entity}ById(state, data) {
    		let index = findItemIndex(state.{Entity}List, function(item) {
	        	return item.Id === data.Id;
	    	});
	    
	    	if (index && index !== 0) {
		        state.{Entity}List.splice(index, 1, data);
    		}
		},
		delete{Entity}FromList(state, id) {
		    let index = findItemIndex(state.{Entity}List, function(item) {
		        return item.Id === id;
		    });
		    
		    if (index && index !== 0) {
		        state.{Entity}List.splice(index, 1);
		    }
		},
		add{Entity}ItemToList(state, item) {
		    state.{Entity}List.push(item);
		},
    },
    state: {
        {Entity}: new {Entity}(),
        {Entity}List: [],
    },
};

export default {entity};
`
