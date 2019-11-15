
import {Authenticator} from "../apiModel";
import api from "../api";
import {findItemIndex} from "../common";

let findUrl = "/api/v1/authenticator";
let readUrl = "/api/v1/authenticator/"; // + id
let createUrl = "/api/v1/authenticator";
let updateUrl = "/api/v1/authenticator/"; // + id
let deleteUrl = "/api/v1/authenticator/"; // + id
let findOrCreateUrl = "/api/v1/authenticator"; // + id

const authenticator = {
    actions: {
        createAuthenticator(context, {data, filter, header}) {

            return api.create(createUrl, data, filter, header)
                .then(function(response) {

                    context.commit("setAuthenticator", response.Model);

                    return response;
                })
                .catch(function(err) {
                    console.error(err);
                    throw(err);
                });
        },
        deleteAuthenticator(context, {id, header}) {

            return api.remove(deleteUrl + id, header)
                .then(function(response) {
                    context.commit("clearAuthenticator");
                    return response;
                })
                .catch(function(err) {
                    console.error(err);
                    throw(err);
                });
        },
        findAuthenticator(context, {filter, header}) {

            return api.find(findUrl, filter, header)
                .then(function(response) {

                    context.commit("setAuthenticator__List", response.List);
                    return response;
                })
                .catch(function(err) {
                    console.error(err);
                    throw(err);
                });
        },
        loadAuthenticator(context, {id, filter, header}) {

            return api.find(readUrl + id, filter, header)
                .then(function(response) {

                    context.commit("setAuthenticator", response.Model);
                    return response;
                })
                .catch(function(err) {
                    console.error(err);
                    throw(err);
                });
        },
        updateAuthenticator(context, {id, data, filter, header}) {

            return api.update(updateUrl + id, data, filter, header)
                .then(function(response) {

                    context.commit("setAuthenticator", response.Model);
                    return response;
                })
                .catch(function(err) {
                    console.error(err);
                    throw(err);
                });
        },
        findOrCreateAuthenticator(context, {id, data, filter, header}) {

            return api.update(findOrCreateUrl, data, filter, header)
                .then(function(response) {

                    context.commit("setAuthenticator", response.Model);
                    return response;
                })
                .catch(function(err) {
                    console.error(err);
                    throw(err);
                });
        },
        clearListAuthenticator(context) {
            context.commit("clearListAuthenticator");
        },
    },
    getters: {
        getAuthenticator: (state) => {
            return state.Authenticator;
        },
        getAuthenticatorById: state => id => {
            return state.Authenticator__List.find(item => item.Id === id);
        },
        getListAuthenticator: (state) => {
            return state.Authenticator__List;
        },
    },
    mutations: {
        setAuthenticator(state, data) {
            state.Authenticator = data;
        },
        setAuthenticator__List(state, data) {
            state.Authenticator__List = data || [];
        },
        clearAuthenticator(state) {
            state.Authenticator = new Authenticator();
        },
        clearListAuthenticator(state) {
            state.Authenticator__List = [];
        },
		updateAuthenticatorById(state, data) {
    		let index = findItemIndex(state.Authenticator__List, function(item) {
	        	return item.Id === data.Id;
	    	});
	    
	    	if (index || index === 0) {
		        state.Authenticator__List.splice(index, 1, data);
    		}
		},
		deleteAuthenticatorFromList(state, id) {
		    let index = findItemIndex(state.Authenticator__List, function(item) {
		        return item.Id === id;
		    });
		    
		    if (index || index === 0) {
		        state.Authenticator__List.splice(index, 1);
		    }
		},
		addAuthenticatorItemToList(state, item) {

			if (state.Authenticator__List === null) {
				state.Authenticator__List = [];
			}

		    state.Authenticator__List.push(item);
		},
    },
    state: {
        Authenticator: new Authenticator(),
        Authenticator__List: [],
    },
};

export default authenticator;
