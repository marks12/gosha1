
import {CurrentApp} from "../apiModel";
import api from "../api";
import {findItemIndex} from "../common";

let findUrl = "/api/v1/currentApp";
let readUrl = "/api/v1/currentApp/"; // + id
let createUrl = "/api/v1/currentApp";
let multiCreateUrl = "/api/v1/currentApp/list";
let updateUrl = "/api/v1/currentApp/"; // + id
let multiUpdateUrl = "/api/v1/currentApp/list"; // + id
let deleteUrl = "/api/v1/currentApp/"; // + id
let multiDeleteUrl = "/api/v1/currentApp/list"; // + id
let findOrCreateUrl = "/api/v1/currentApp"; // + id

const currentApp = {
    actions: {
        createCurrentApp(context, {data, filter, header}) {

            let url = createUrl;
            if (Array.isArray && Array.isArray(data)) {
                url = multiCreateUrl
            }

            return api.create(url, data, filter, header)
                .then(function(response) {

                    context.commit("setCurrentApp", response.Model);

                    return response;
                })
                .catch(function(err) {
                    console.error(err);
                    throw(err);
                });
        },
        deleteCurrentApp(context, {id, header}) {

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
                    context.commit("clearCurrentApp");
                    return response;
                })
                .catch(function(err) {
                    console.error(err);
                    throw(err);
                });
        },
        findCurrentApp(context, {filter, header, isAppend}) {

            return api.find(findUrl, filter, header)
                .then(function(response) {

                    if (isAppend) {
                        context.commit("appendCurrentApp__List", response.List);
                    } else {
                        context.commit("setCurrentApp__List", response.List);
                    }

                    return response;
                })
                .catch(function(err) {
                    console.error(err);
                    throw(err);
                });
        },
        loadCurrentApp(context, {id, filter, header}) {

            return api.find(readUrl + id, filter, header)
                .then(function(response) {

                    context.commit("setCurrentApp", response.Model);
                    return response;
                })
                .catch(function(err) {
                    console.error(err);
                    throw(err);
                });
        },
        updateCurrentApp(context, {id, data, filter, header}) {

            let url = updateUrl + id;
            if (Array.isArray && Array.isArray(data)) {
                url = multiUpdateUrl
            }

            return api.update(url, data, filter, header)
                .then(function(response) {

                    context.commit("setCurrentApp", response.Model);
                    return response;
                })
                .catch(function(err) {
                    console.error(err);
                    throw(err);
                });
        },
        findOrCreateCurrentApp(context, {id, data, filter, header}) {

            return api.update(findOrCreateUrl, data, filter, header)
                .then(function(response) {

                    context.commit("setCurrentApp", response.Model);
                    return response;
                })
                .catch(function(err) {
                    console.error(err);
                    throw(err);
                });
        },
        clearListCurrentApp(context) {
            context.commit("clearListCurrentApp");
        },
        clearCurrentApp(context) {
            context.commit("clearCurrentApp");
        },
    },
    getters: {
        getCurrentApp: (state) => {
            return state.CurrentApp;
        },
        getCurrentAppById: state => id => {
            return state.CurrentApp__List.find(item => item.Id === id);
        },
        getListCurrentApp: (state) => {
            return state.CurrentApp__List;
        },
    },
    mutations: {
        setCurrentApp(state, data) {
            state.CurrentApp = data;
        },
        setCurrentApp__List(state, data) {
            state.CurrentApp__List = data || [];
        },
        appendCurrentApp__List(state, data) {

            if (! state.CurrentApp__List) {
                state.CurrentApp__List = [];
            }

            state.CurrentApp__List = state.CurrentApp__List.concat(data);
        },
        clearCurrentApp(state) {
            state.CurrentApp = new CurrentApp();
        },
        clearListCurrentApp(state) {
            state.CurrentApp__List = [];
        },
		updateCurrentAppById(state, data) {
    		let index = findItemIndex(state.CurrentApp__List, function(item) {
	        	return item.Id === data.Id;
	    	});
	    
	    	if (index || index === 0) {
		        state.CurrentApp__List.splice(index, 1, data);
    		}
		},
		deleteCurrentAppFromList(state, id) {
		    let index = findItemIndex(state.CurrentApp__List, function(item) {
		        return item.Id === id;
		    });
		    
		    if (index || index === 0) {
		        state.CurrentApp__List.splice(index, 1);
		    }
		},
		addCurrentAppItemToList(state, item) {

			if (state.CurrentApp__List === null) {
				state.CurrentApp__List = [];
			}

		    state.CurrentApp__List.push(item);
		},
    },
    state: {
        CurrentApp: new CurrentApp(),
        CurrentApp__List: [],
    },
};

export default currentApp;
