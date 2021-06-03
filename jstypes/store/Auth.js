
import {Auth} from "../apiModel";
import api from "../api";
import {findItemIndex} from "../common";

let findUrl = "/api/v1/auth";
let readUrl = "/api/v1/auth/"; // + id
let createUrl = "/api/v1/auth";
let multiCreateUrl = "/api/v1/auth/list";
let updateUrl = "/api/v1/auth/"; // + id
let multiUpdateUrl = "/api/v1/auth/list";
let deleteUrl = "/api/v1/auth/"; // + id
let multiDeleteUrl = "/api/v1/auth/list";
let findOrCreateUrl = "/api/v1/auth";
let updateOrCreateUrl = "/api/v1/auth";

const auth = {
    actions: {
        createAuth(context, {data, filter, header, noMutation}) {

            let url = createUrl;
            if (Array.isArray && Array.isArray(data)) {
                url = multiCreateUrl
            }

            return api.create(url, data, filter, header)
                .then(function(response) {

					if(! noMutation) {
	                    context.commit("setAuth", response.Model);
					}

                    return response;
                })
                .catch(function(err) {
                    console.error(err);
                    throw(err);
                });
        },
        deleteAuth(context, {id, header, noMutation}) {

            let url;
            let dataOrNull = null;

            if (Array.isArray && Array.isArray(id)) {
                url = multiDeleteUrl;
                dataOrNull = id.map(item => typeof item === "number" ? {Id: item} : item);
            } else {
                url = deleteUrl + id;
            }

            return api.remove(url, header, dataOrNull)
                .then(function(response) {
					if(! noMutation) {
	                    context.commit("clearAuth");
					}
                    return response;
                })
                .catch(function(err) {
                    console.error(err);
                    throw(err);
                });
        },
        findAuth(context, {filter, header, isAppend, noMutation}) {

            return api.find(findUrl, filter, header)
                .then(function(response) {

					if(! noMutation) {
						if (isAppend) {
							context.commit("appendAuth__List", response.List);
						} else {
							context.commit("setAuth__List", response.List);
						}
					}

                    return response;
                })
                .catch(function(err) {
                    console.error(err);
                    throw(err);
                });
        },
        loadAuth(context, {id, filter, header, noMutation}) {

            return api.find(readUrl + id, filter, header)
                .then(function(response) {

					if(! noMutation) {
	                    context.commit("setAuth", response.Model);
					}
                    return response;
                })
                .catch(function(err) {
                    console.error(err);
                    throw(err);
                });
        },
        updateAuth(context, {id, data, filter, header, noMutation}) {

            let url = updateUrl + id;
            if (Array.isArray && Array.isArray(data)) {
                url = multiUpdateUrl
            }

            return api.update(url, data, filter, header)
                .then(function(response) {
					if(! noMutation) {
	                    context.commit("setAuth", response.Model);
					}
                    return response;
                })
                .catch(function(err) {
                    console.error(err);
                    throw(err);
                });
        },
        findOrCreateAuth(context, {id, data, filter, header, noMutation}) {

            return api.update(findOrCreateUrl, data, filter, header)
                .then(function(response) {

					if(! noMutation) {
	                    context.commit("setAuth", response.Model);
					}
                    return response;
                })
                .catch(function(err) {
                    console.error(err);
                    throw(err);
                });
        },
        clearListAuth(context) {
            context.commit("clearListAuth");
        },
        clearAuth(context) {
            context.commit("clearAuth");
        },
    },
    getters: {
        getAuth: (state) => {
            return state.Auth;
        },
        getAuthById: state => id => {
            return state.Auth__List.find(item => item.Id === id);
        },
        getListAuth: (state) => {
            return state.Auth__List;
        },
        getRoute__Auth: state => action => {
            return state.Auth__Routes[action];
        },
        getRoutes__Auth: state => {
            return state.Auth__Routes;
        },
    },
    mutations: {
        setAuth(state, data) {
            state.Auth = data;
        },
        setAuth__List(state, data) {
            state.Auth__List = data || [];
        },
        appendAuth__List(state, data) {

            if (! state.Auth__List) {
                state.Auth__List = [];
            }

			if (data !== null) {
				state.Auth__List = state.Auth__List.concat(data);				
			}
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
        Auth__Routes: {
            find: findUrl,
            read: readUrl,
            create: createUrl,
            multiCreate: multiCreateUrl,
            update: updateUrl,
            multiUpdate: multiUpdateUrl,
            delete: deleteUrl,
            multiDelete: multiDeleteUrl,
            findOrCreate: findOrCreateUrl,
            updateOrCreate: updateOrCreateUrl,
        },
    },
};

export default auth;
