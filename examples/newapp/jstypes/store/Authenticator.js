
import {Authenticator} from "../apiModel";
import api from "../api";
import {findItemIndex} from "../common";

let findUrl = "/api/v1/authenticator";
let readUrl = "/api/v1/authenticator/"; // + id
let createUrl = "/api/v1/authenticator";
let multiCreateUrl = "/api/v1/authenticator/list";
let updateUrl = "/api/v1/authenticator/"; // + id
let multiUpdateUrl = "/api/v1/authenticator/list"; // + id
let deleteUrl = "/api/v1/authenticator/"; // + id
let multiDeleteUrl = "/api/v1/authenticator/list"; // + id
let findOrCreateUrl = "/api/v1/authenticator"; // + id

const authenticator = {
    actions: {
        createAuthenticator(context, {data, filter, header}) {

            let url = createUrl;
            if (Array.isArray && Array.isArray(data)) {
                url = multiCreateUrl
            }

            return api.create(url, data, filter, header)
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

            let url = updateUrl + id;
            if (Array.isArray && Array.isArray(data)) {
                url = multiUpdateUrl
            }

            return api.update(url, data, filter, header)
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
        clearAuthenticator(context) {
            context.commit("clearAuthenticator");
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
