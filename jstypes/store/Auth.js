
import {Auth} from "../apiModel";
import api from "../api";
import {findItemIndex} from "../common";

let findUrl = "/api/v1/auth";
let readUrl = "/api/v1/auth/"; // + id
let createUrl = "/api/v1/auth";
let updateUrl = "/api/v1/auth/"; // + id
let deleteUrl = "/api/v1/auth/"; // + id
let findOrCreateUrl = "/api/v1/auth"; // + id

const auth = {
    actions: {
        createAuth(context, {data, filter, header}) {

            return api.create(createUrl, data, filter, header)
                .then(function(response) {

                    context.commit("setAuth", response.Model);

                    return response;
                })
                .catch(function(err) {
                    return err;
                });
        },
        deleteAuth(context, {id, header}) {

            return api.remove(deleteUrl + id, header)
                .then(function(response) {
                    context.commit("clearAuth");
                    return response;
                })
                .catch(function(err) {
                    return err;
                });
        },
        findAuth(context, {filter, header}) {

            return api.find(findUrl, filter, header)
                .then(function(response) {

                    context.commit("setAuth__List", response.List);

                    return response;
                })
                .catch(function(err) {
                    return err;
                });
        },
        loadAuth(context, {id, filter, header}) {

            return api.find(readUrl + id, filter, header)
                .then(function(response) {

                    context.commit("setAuth", response.Model);

                    return response;
                })
                .catch(function(err) {
                    return err;
                });
        },
        updateAuth(context, {id, data, filter, header}) {

            return api.update(updateUrl + id, data, filter, header)
                .then(function(response) {

                    context.commit("setAuth", response.Model);

                    return response;
                })
                .catch(function(err) {
                    return err;
                });
        },
        findOrCreateAuth(context, {id, data, filter, header}) {

            return api.update(findOrCreateUrl, data, filter, header)
                .then(function(response) {

                    context.commit("setAuth", response.Model);

                    return response;
                })
                .catch(function(err) {
                    return err;
                });
        },
        clearListAuth(context) {
            context.commit("clearListAuth");
        },
    },
    getters: {
        getAuth: (state) => {
            return state.Auth;
        },
        getListAuth: (state) => {
            return state.Auth__List;
        },
    },
    mutations: {
        setAuth(state, data) {
            state.Auth = data;
        },
        setAuth__List(state, data) {
            state.Auth__List = data;
        },
        clearAuth(state) {
            state.Auth = new Auth();
        },
        clearListAuth(state) {
            state.Auth__List = [];
        },
		updateAuthById(state, data) {
    		let index = findItemIndex(state.Auth__List, function(item) {
	        	return item.Id === data.Id;
	    	});
	    
	    	if (index || index === 0) {
		        state.Auth__List.splice(index, 1, data);
    		}
		},
		deleteAuthFromList(state, id) {
		    let index = findItemIndex(state.Auth__List, function(item) {
		        return item.Id === id;
		    });
		    
		    if (index || index === 0) {
		        state.Auth__List.splice(index, 1);
		    }
		},
		addAuthItemToList(state, item) {

			if (state.Auth__List === null) {
				state.Auth__List = [];
			}

		    state.Auth__List.push(item);
		},
    },
    state: {
        Auth: new Auth(),
        Auth__List: [],
    },
};

export default auth;
