
import {AuthFilter} from "../apiModel";
import api from "../api";
import {findItemIndex} from "../common";

let findUrl = "/api/v1/authFilter";
let readUrl = "/api/v1/authFilter/"; // + id
let createUrl = "/api/v1/authFilter";
let multiCreateUrl = "/api/v1/authFilter/list";
let updateUrl = "/api/v1/authFilter/"; // + id
let multiUpdateUrl = "/api/v1/authFilter/list"; // + id
let deleteUrl = "/api/v1/authFilter/"; // + id
let multiDeleteUrl = "/api/v1/authFilter/list"; // + id
let findOrCreateUrl = "/api/v1/authFilter"; // + id

const authFilter = {
    actions: {
        createAuthFilter(context, {data, filter, header}) {

            let url = createUrl;
            if (Array.isArray && Array.isArray(data)) {
                url = multiCreateUrl
            }

            return api.create(url, data, filter, header)
                .then(function(response) {

                    context.commit("setAuthFilter", response.Model);

                    return response;
                })
                .catch(function(err) {
                    console.error(err);
                    throw(err);
                });
        },
        deleteAuthFilter(context, {id, header}) {

            let url;
            let dataOrNull = null;

            if (Array.isArray && Array.isArray(id)) {
                url = multiDeleteUrl;
                dataOrNull = id;
            } else {
                url = deleteUrl + id;
            }

            return api.remove(url, header, dataOrNull)
                .then(function(response) {
                    context.commit("clearAuthFilter");
                    return response;
                })
                .catch(function(err) {
                    console.error(err);
                    throw(err);
                });
        },
        findAuthFilter(context, {filter, header, isAppend}) {

            return api.find(findUrl, filter, header)
                .then(function(response) {

                    if (isAppend) {
                        context.commit("appendAuthFilter__List", response.List);
                    } else {
                        context.commit("setAuthFilter__List", response.List);
                    }

                    return response;
                })
                .catch(function(err) {
                    console.error(err);
                    throw(err);
                });
        },
        loadAuthFilter(context, {id, filter, header}) {

            return api.find(readUrl + id, filter, header)
                .then(function(response) {

                    context.commit("setAuthFilter", response.Model);
                    return response;
                })
                .catch(function(err) {
                    console.error(err);
                    throw(err);
                });
        },
        updateAuthFilter(context, {id, data, filter, header}) {

            let url = updateUrl + id;
            if (Array.isArray && Array.isArray(data)) {
                url = multiUpdateUrl
            }

            return api.update(url, data, filter, header)
                .then(function(response) {

                    context.commit("setAuthFilter", response.Model);
                    return response;
                })
                .catch(function(err) {
                    console.error(err);
                    throw(err);
                });
        },
        findOrCreateAuthFilter(context, {id, data, filter, header}) {

            return api.update(findOrCreateUrl, data, filter, header)
                .then(function(response) {

                    context.commit("setAuthFilter", response.Model);
                    return response;
                })
                .catch(function(err) {
                    console.error(err);
                    throw(err);
                });
        },
        clearListAuthFilter(context) {
            context.commit("clearListAuthFilter");
        },
        clearAuthFilter(context) {
            context.commit("clearAuthFilter");
        },
    },
    getters: {
        getAuthFilter: (state) => {
            return state.AuthFilter;
        },
        getAuthFilterById: state => id => {
            return state.AuthFilter__List.find(item => item.Id === id);
        },
        getListAuthFilter: (state) => {
            return state.AuthFilter__List;
        },
    },
    mutations: {
        setAuthFilter(state, data) {
            state.AuthFilter = data;
        },
        setAuthFilter__List(state, data) {
            state.AuthFilter__List = data || [];
        },
        appendAuthFilter__List(state, data) {

            if (! state.AuthFilter__List) {
                state.AuthFilter__List = [];
            }

            state.AuthFilter__List = state.AuthFilter__List.concat(data);
        },
        clearAuthFilter(state) {
            state.AuthFilter = new AuthFilter();
        },
        clearListAuthFilter(state) {
            state.AuthFilter__List = [];
        },
		updateAuthFilterById(state, data) {
    		let index = findItemIndex(state.AuthFilter__List, function(item) {
	        	return item.Id === data.Id;
	    	});
	    
	    	if (index || index === 0) {
		        state.AuthFilter__List.splice(index, 1, data);
    		}
		},
		deleteAuthFilterFromList(state, id) {
		    let index = findItemIndex(state.AuthFilter__List, function(item) {
		        return item.Id === id;
		    });
		    
		    if (index || index === 0) {
		        state.AuthFilter__List.splice(index, 1);
		    }
		},
		addAuthFilterItemToList(state, item) {

			if (state.AuthFilter__List === null) {
				state.AuthFilter__List = [];
			}

		    state.AuthFilter__List.push(item);
		},
    },
    state: {
        AuthFilter: new AuthFilter(),
        AuthFilter__List: [],
    },
};

export default authFilter;
