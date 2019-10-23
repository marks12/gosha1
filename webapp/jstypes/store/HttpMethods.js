
import {HttpMethods} from "../apiModel";
import api from "../api";
import {findItemIndex} from "../common";

let findUrl = "/api/v1/httpMethods";
let readUrl = "/api/v1/httpMethods/"; // + id
let createUrl = "/api/v1/httpMethods";
let updateUrl = "/api/v1/httpMethods/"; // + id
let deleteUrl = "/api/v1/httpMethods/"; // + id
let findOrCreateUrl = "/api/v1/httpMethods"; // + id

const httpMethods = {
    actions: {
        createHttpMethods(context, {data, filter, header}) {

            return api.create(createUrl, data, filter, header)
                .then(function(response) {

                    context.commit("setHttpMethods", response.Model);

                    return response;
                })
                .catch(function(err) {
                    console.error(err);
                    throw(err);
                });
        },
        deleteHttpMethods(context, {id, header}) {

            return api.remove(deleteUrl + id, header)
                .then(function(response) {
                    context.commit("clearHttpMethods");
                    return response;
                })
                .catch(function(err) {
                    console.error(err);
                    throw(err);
                });
        },
        findHttpMethods(context, {filter, header}) {

            return api.find(findUrl, filter, header)
                .then(function(response) {

                    context.commit("setHttpMethods__List", response.List);
                    return response;
                })
                .catch(function(err) {
                    console.error(err);
                    throw(err);
                });
        },
        loadHttpMethods(context, {id, filter, header}) {

            return api.find(readUrl + id, filter, header)
                .then(function(response) {

                    context.commit("setHttpMethods", response.Model);
                    return response;
                })
                .catch(function(err) {
                    console.error(err);
                    throw(err);
                });
        },
        updateHttpMethods(context, {id, data, filter, header}) {

            return api.update(updateUrl + id, data, filter, header)
                .then(function(response) {

                    context.commit("setHttpMethods", response.Model);
                    return response;
                })
                .catch(function(err) {
                    console.error(err);
                    throw(err);
                });
        },
        findOrCreateHttpMethods(context, {id, data, filter, header}) {

            return api.update(findOrCreateUrl, data, filter, header)
                .then(function(response) {

                    context.commit("setHttpMethods", response.Model);
                    return response;
                })
                .catch(function(err) {
                    console.error(err);
                    throw(err);
                });
        },
        clearListHttpMethods(context) {
            context.commit("clearListHttpMethods");
        },
    },
    getters: {
        getHttpMethods: (state) => {
            return state.HttpMethods;
        },
        getHttpMethodsById: state => id => {
            return state.HttpMethods__List.find(item => item.Id === id);
        },
        getListHttpMethods: (state) => {
            return state.HttpMethods__List;
        },
    },
    mutations: {
        setHttpMethods(state, data) {
            state.HttpMethods = data;
        },
        setHttpMethods__List(state, data) {
            state.HttpMethods__List = data || [];
        },
        clearHttpMethods(state) {
            state.HttpMethods = new HttpMethods();
        },
        clearListHttpMethods(state) {
            state.HttpMethods__List = [];
        },
		updateHttpMethodsById(state, data) {
    		let index = findItemIndex(state.HttpMethods__List, function(item) {
	        	return item.Id === data.Id;
	    	});
	    
	    	if (index || index === 0) {
		        state.HttpMethods__List.splice(index, 1, data);
    		}
		},
		deleteHttpMethodsFromList(state, id) {
		    let index = findItemIndex(state.HttpMethods__List, function(item) {
		        return item.Id === id;
		    });
		    
		    if (index || index === 0) {
		        state.HttpMethods__List.splice(index, 1);
		    }
		},
		addHttpMethodsItemToList(state, item) {

			if (state.HttpMethods__List === null) {
				state.HttpMethods__List = [];
			}

		    state.HttpMethods__List.push(item);
		},
    },
    state: {
        HttpMethods: new HttpMethods(),
        HttpMethods__List: [],
    },
};

export default httpMethods;
