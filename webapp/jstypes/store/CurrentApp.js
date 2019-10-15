
import {CurrentApp} from "../apiModel";
import api from "../api";
import {findItemIndex} from "../common";

let findUrl = "/api/v1/currentApp";
let readUrl = "/api/v1/currentApp/"; // + id
let createUrl = "/api/v1/currentApp";
let updateUrl = "/api/v1/currentApp/"; // + id
let deleteUrl = "/api/v1/currentApp/"; // + id
let findOrCreateUrl = "/api/v1/currentApp"; // + id

const currentApp = {
    actions: {
        createCurrentApp(context, {data, filter, header}) {

            return api.create(createUrl, data, filter, header)
                .then(function(response) {

                    context.commit("setCurrentApp", response.Model);

                    return response;
                })
                .catch(function(err) {
                    return err;
                });
        },
        deleteCurrentApp(context, {id, header}) {

            return api.remove(deleteUrl + id, header)
                .then(function(response) {
                    context.commit("clearCurrentApp");
                    return response;
                })
                .catch(function(err) {
                    return err;
                });
        },
        findCurrentApp(context, {filter, header}) {

            return api.find(findUrl, filter, header)
                .then(function(response) {

                    context.commit("setCurrentApp__List", response.List);

                    return response;
                })
                .catch(function(err) {
                    return err;
                });
        },
        loadCurrentApp(context, {id, filter, header}) {

            return api.find(readUrl + id, filter, header)
                .then(function(response) {

                    context.commit("setCurrentApp", response.Model);

                    return response;
                })
                .catch(function(err) {
                    return err;
                });
        },
        updateCurrentApp(context, {id, data, filter, header}) {

            return api.update(updateUrl + id, data, filter, header)
                .then(function(response) {

                    context.commit("setCurrentApp", response.Model);

                    return response;
                })
                .catch(function(err) {
                    return err;
                });
        },
        findOrCreateCurrentApp(context, {id, data, filter, header}) {

            return api.update(findOrCreateUrl, data, filter, header)
                .then(function(response) {

                    context.commit("setCurrentApp", response.Model);

                    return response;
                })
                .catch(function(err) {
                    return err;
                });
        },
        clearListCurrentApp(context) {
            context.commit("clearListCurrentApp");
        },
    },
    getters: {
        getCurrentApp: (state) => {
            return state.CurrentApp;
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
