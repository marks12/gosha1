
import {ResourceTypeFilter} from "../apiModel";
import api from "../api";
import {findItemIndex} from "../common";

let findUrl = "/api/v1/resourceTypeFilter";
let readUrl = "/api/v1/resourceTypeFilter/"; // + id
let createUrl = "/api/v1/resourceTypeFilter";
let multiCreateUrl = "/api/v1/resourceTypeFilter/list";
let updateUrl = "/api/v1/resourceTypeFilter/"; // + id
let multiUpdateUrl = "/api/v1/resourceTypeFilter/list"; // + id
let deleteUrl = "/api/v1/resourceTypeFilter/"; // + id
let multiDeleteUrl = "/api/v1/resourceTypeFilter/list"; // + id
let findOrCreateUrl = "/api/v1/resourceTypeFilter"; // + id

const resourceTypeFilter = {
    actions: {
        createResourceTypeFilter(context, {data, filter, header}) {

            let url = createUrl;
            if (Array.isArray && Array.isArray(data)) {
                url = multiCreateUrl
            }

            return api.create(url, data, filter, header)
                .then(function(response) {

                    context.commit("setResourceTypeFilter", response.Model);

                    return response;
                })
                .catch(function(err) {
                    console.error(err);
                    throw(err);
                });
        },
        deleteResourceTypeFilter(context, {id, header}) {

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
                    context.commit("clearResourceTypeFilter");
                    return response;
                })
                .catch(function(err) {
                    console.error(err);
                    throw(err);
                });
        },
        findResourceTypeFilter(context, {filter, header, isAppend}) {

            return api.find(findUrl, filter, header)
                .then(function(response) {

                    if (isAppend) {
                        context.commit("appendResourceTypeFilter__List", response.List);
                    } else {
                        context.commit("setResourceTypeFilter__List", response.List);
                    }

                    return response;
                })
                .catch(function(err) {
                    console.error(err);
                    throw(err);
                });
        },
        loadResourceTypeFilter(context, {id, filter, header}) {

            return api.find(readUrl + id, filter, header)
                .then(function(response) {

                    context.commit("setResourceTypeFilter", response.Model);
                    return response;
                })
                .catch(function(err) {
                    console.error(err);
                    throw(err);
                });
        },
        updateResourceTypeFilter(context, {id, data, filter, header}) {

            let url = updateUrl + id;
            if (Array.isArray && Array.isArray(data)) {
                url = multiUpdateUrl
            }

            return api.update(url, data, filter, header)
                .then(function(response) {

                    context.commit("setResourceTypeFilter", response.Model);
                    return response;
                })
                .catch(function(err) {
                    console.error(err);
                    throw(err);
                });
        },
        findOrCreateResourceTypeFilter(context, {id, data, filter, header}) {

            return api.update(findOrCreateUrl, data, filter, header)
                .then(function(response) {

                    context.commit("setResourceTypeFilter", response.Model);
                    return response;
                })
                .catch(function(err) {
                    console.error(err);
                    throw(err);
                });
        },
        clearListResourceTypeFilter(context) {
            context.commit("clearListResourceTypeFilter");
        },
        clearResourceTypeFilter(context) {
            context.commit("clearResourceTypeFilter");
        },
    },
    getters: {
        getResourceTypeFilter: (state) => {
            return state.ResourceTypeFilter;
        },
        getResourceTypeFilterById: state => id => {
            return state.ResourceTypeFilter__List.find(item => item.Id === id);
        },
        getListResourceTypeFilter: (state) => {
            return state.ResourceTypeFilter__List;
        },
    },
    mutations: {
        setResourceTypeFilter(state, data) {
            state.ResourceTypeFilter = data;
        },
        setResourceTypeFilter__List(state, data) {
            state.ResourceTypeFilter__List = data || [];
        },
        appendResourceTypeFilter__List(state, data) {

            if (! state.ResourceTypeFilter__List) {
                state.ResourceTypeFilter__List = [];
            }

            state.ResourceTypeFilter__List = state.ResourceTypeFilter__List.concat(data);
        },
        clearResourceTypeFilter(state) {
            state.ResourceTypeFilter = new ResourceTypeFilter();
        },
        clearListResourceTypeFilter(state) {
            state.ResourceTypeFilter__List = [];
        },
		updateResourceTypeFilterById(state, data) {
    		let index = findItemIndex(state.ResourceTypeFilter__List, function(item) {
	        	return item.Id === data.Id;
	    	});
	    
	    	if (index || index === 0) {
		        state.ResourceTypeFilter__List.splice(index, 1, data);
    		}
		},
		deleteResourceTypeFilterFromList(state, id) {
		    let index = findItemIndex(state.ResourceTypeFilter__List, function(item) {
		        return item.Id === id;
		    });
		    
		    if (index || index === 0) {
		        state.ResourceTypeFilter__List.splice(index, 1);
		    }
		},
		addResourceTypeFilterItemToList(state, item) {

			if (state.ResourceTypeFilter__List === null) {
				state.ResourceTypeFilter__List = [];
			}

		    state.ResourceTypeFilter__List.push(item);
		},
    },
    state: {
        ResourceTypeFilter: new ResourceTypeFilter(),
        ResourceTypeFilter__List: [],
    },
};

export default resourceTypeFilter;
