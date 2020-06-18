
import {APIStatus} from "../apiModel";
import api from "../api";
import {findItemIndex} from "../common";

let findUrl = "/api/v1/aPIStatus";
let readUrl = "/api/v1/aPIStatus/"; // + id
let createUrl = "/api/v1/aPIStatus";
let multiCreateUrl = "/api/v1/aPIStatus/list";
let updateUrl = "/api/v1/aPIStatus/"; // + id
let multiUpdateUrl = "/api/v1/aPIStatus/list";
let deleteUrl = "/api/v1/aPIStatus/"; // + id
let multiDeleteUrl = "/api/v1/aPIStatus/list";
let findOrCreateUrl = "/api/v1/aPIStatus";
let updateOrCreateUrl = "/api/v1/aPIStatus";

const aPIStatus = {
    actions: {
        createAPIStatus(context, {data, filter, header, noMutation}) {

            let url = createUrl;
            if (Array.isArray && Array.isArray(data)) {
                url = multiCreateUrl
            }

            return api.create(url, data, filter, header)
                .then(function(response) {

					if(! noMutation) {
	                    context.commit("setAPIStatus", response.Model);
					}

                    return response;
                })
                .catch(function(err) {
                    console.error(err);
                    throw(err);
                });
        },
        deleteAPIStatus(context, {id, header, noMutation}) {

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
	                    context.commit("clearAPIStatus");
					}
                    return response;
                })
                .catch(function(err) {
                    console.error(err);
                    throw(err);
                });
        },
        findAPIStatus(context, {filter, header, isAppend, noMutation}) {

            return api.find(findUrl, filter, header)
                .then(function(response) {

					if(! noMutation) {
						if (isAppend) {
							context.commit("appendAPIStatus__List", response.List);
						} else {
							context.commit("setAPIStatus__List", response.List);
						}
					}

                    return response;
                })
                .catch(function(err) {
                    console.error(err);
                    throw(err);
                });
        },
        loadAPIStatus(context, {id, filter, header, noMutation}) {

            return api.find(readUrl + id, filter, header)
                .then(function(response) {

					if(! noMutation) {
	                    context.commit("setAPIStatus", response.Model);
					}
                    return response;
                })
                .catch(function(err) {
                    console.error(err);
                    throw(err);
                });
        },
        updateAPIStatus(context, {id, data, filter, header, noMutation}) {

            let url = updateUrl + id;
            if (Array.isArray && Array.isArray(data)) {
                url = multiUpdateUrl
            }

            return api.update(url, data, filter, header)
                .then(function(response) {
					if(! noMutation) {
	                    context.commit("setAPIStatus", response.Model);
					}
                    return response;
                })
                .catch(function(err) {
                    console.error(err);
                    throw(err);
                });
        },
        findOrCreateAPIStatus(context, {id, data, filter, header, noMutation}) {

            return api.update(findOrCreateUrl, data, filter, header)
                .then(function(response) {

					if(! noMutation) {
	                    context.commit("setAPIStatus", response.Model);
					}
                    return response;
                })
                .catch(function(err) {
                    console.error(err);
                    throw(err);
                });
        },
        clearListAPIStatus(context) {
            context.commit("clearListAPIStatus");
        },
        clearAPIStatus(context) {
            context.commit("clearAPIStatus");
        },
    },
    getters: {
        getAPIStatus: (state) => {
            return state.APIStatus;
        },
        getAPIStatusById: state => id => {
            return state.APIStatus__List.find(item => item.Id === id);
        },
        getListAPIStatus: (state) => {
            return state.APIStatus__List;
        },
        getRoute__APIStatus: state => action => {
            return state.APIStatus__Routes[action];
        },
        getRoutes__APIStatus: state => {
            return state.APIStatus__Routes;
        },
    },
    mutations: {
        setAPIStatus(state, data) {
            state.APIStatus = data;
        },
        setAPIStatus__List(state, data) {
            state.APIStatus__List = data || [];
        },
        appendAPIStatus__List(state, data) {

            if (! state.APIStatus__List) {
                state.APIStatus__List = [];
            }

            state.APIStatus__List = state.APIStatus__List.concat(data);
        },
        clearAPIStatus(state) {
            state.APIStatus = new APIStatus();
        },
        clearListAPIStatus(state) {
            state.APIStatus__List = [];
        },
		updateAPIStatusById(state, data) {
    		let index = findItemIndex(state.APIStatus__List, function(item) {
	        	return item.Id === data.Id;
	    	});
	    
	    	if (index || index === 0) {
		        state.APIStatus__List.splice(index, 1, data);
    		}
		},
		deleteAPIStatusFromList(state, id) {
		    let index = findItemIndex(state.APIStatus__List, function(item) {
		        return item.Id === id;
		    });
		    
		    if (index || index === 0) {
		        state.APIStatus__List.splice(index, 1);
		    }
		},
		addAPIStatusItemToList(state, item) {

			if (state.APIStatus__List === null) {
				state.APIStatus__List = [];
			}

		    state.APIStatus__List.push(item);
		},
    },
    state: {
        APIStatus: new APIStatus(),
        APIStatus__List: [],
        APIStatus__Routes: {
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

export default aPIStatus;
