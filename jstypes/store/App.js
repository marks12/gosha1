
import {App} from "../apiModel";
import api from "../api";
import {findItemIndex} from "../common";

let findUrl = "/api/v1/app";
let readUrl = "/api/v1/app/"; // + id
let createUrl = "/api/v1/app";
let multiCreateUrl = "/api/v1/app/list";
let updateUrl = "/api/v1/app/"; // + id
let multiUpdateUrl = "/api/v1/app/list";
let deleteUrl = "/api/v1/app/"; // + id
let multiDeleteUrl = "/api/v1/app/list";
let findOrCreateUrl = "/api/v1/app";
let updateOrCreateUrl = "/api/v1/app";

const app = {
    actions: {
        createApp(context, {data, filter, header, noMutation}) {

            let url = createUrl;
            if (Array.isArray && Array.isArray(data)) {
                url = multiCreateUrl
            }

            return api.create(url, data, filter, header)
                .then(function(response) {

					if(! noMutation) {
	                    context.commit("setApp", response.Model);
					}

                    return response;
                })
                .catch(function(err) {
                    console.error(err);
                    throw(err);
                });
        },
        deleteApp(context, {id, header, noMutation}) {

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
	                    context.commit("clearApp");
					}
                    return response;
                })
                .catch(function(err) {
                    console.error(err);
                    throw(err);
                });
        },
        findApp(context, {filter, header, isAppend, noMutation}) {

            return api.find(findUrl, filter, header)
                .then(function(response) {

					if(! noMutation) {
						if (isAppend) {
							context.commit("appendApp__List", response.List);
						} else {
							context.commit("setApp__List", response.List);
						}
					}

                    return response;
                })
                .catch(function(err) {
                    console.error(err);
                    throw(err);
                });
        },
        loadApp(context, {id, filter, header, noMutation}) {

            return api.find(readUrl + id, filter, header)
                .then(function(response) {

					if(! noMutation) {
	                    context.commit("setApp", response.Model);
					}
                    return response;
                })
                .catch(function(err) {
                    console.error(err);
                    throw(err);
                });
        },
        updateApp(context, {id, data, filter, header, noMutation}) {

            let url = updateUrl + id;
            if (Array.isArray && Array.isArray(data)) {
                url = multiUpdateUrl
            }

            return api.update(url, data, filter, header)
                .then(function(response) {
					if(! noMutation) {
	                    context.commit("setApp", response.Model);
					}
                    return response;
                })
                .catch(function(err) {
                    console.error(err);
                    throw(err);
                });
        },
        findOrCreateApp(context, {id, data, filter, header, noMutation}) {

            return api.update(findOrCreateUrl, data, filter, header)
                .then(function(response) {

					if(! noMutation) {
	                    context.commit("setApp", response.Model);
					}
                    return response;
                })
                .catch(function(err) {
                    console.error(err);
                    throw(err);
                });
        },
        clearListApp(context) {
            context.commit("clearListApp");
        },
        clearApp(context) {
            context.commit("clearApp");
        },
    },
    getters: {
        getApp: (state) => {
            return state.App;
        },
        getAppById: state => id => {
            return state.App__List.find(item => item.Id === id);
        },
        getListApp: (state) => {
            return state.App__List;
        },
        getRoute__App: state => action => {
            return state.App__Routes[action];
        },
        getRoutes__App: state => {
            return state.App__Routes;
        },
    },
    mutations: {
        setApp(state, data) {
            state.App = data;
        },
        setApp__List(state, data) {
            state.App__List = data || [];
        },
        appendApp__List(state, data) {

            if (! state.App__List) {
                state.App__List = [];
            }

			if (data !== null) {
				state.App__List = state.App__List.concat(data);				
			}
        },
        clearApp(state) {
            state.App = new App();
        },
        clearListApp(state) {
            state.App__List = [];
        },
		updateAppById(state, data) {
    		let index = findItemIndex(state.App__List, function(item) {
	        	return item.Id === data.Id;
	    	});
	    
	    	if (index || index === 0) {
		        state.App__List.splice(index, 1, data);
    		}
		},
		deleteAppFromList(state, id) {
		    let index = findItemIndex(state.App__List, function(item) {
		        return item.Id === id;
		    });
		    
		    if (index || index === 0) {
		        state.App__List.splice(index, 1);
		    }
		},
		addAppItemToList(state, item) {

			if (state.App__List === null) {
				state.App__List = [];
			}

		    state.App__List.push(item);
		},
    },
    state: {
        App: new App(),
        App__List: [],
        App__Routes: {
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

export default app;
