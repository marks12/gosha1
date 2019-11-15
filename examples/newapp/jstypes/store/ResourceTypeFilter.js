
import {ResourceTypeFilter} from "../apiModel";
import api from "../api";
import {findItemIndex} from "../common";

let findUrl = "/api/v1/resourceTypeFilter";
let readUrl = "/api/v1/resourceTypeFilter/"; // + id
let createUrl = "/api/v1/resourceTypeFilter";
let updateUrl = "/api/v1/resourceTypeFilter/"; // + id
let deleteUrl = "/api/v1/resourceTypeFilter/"; // + id
let findOrCreateUrl = "/api/v1/resourceTypeFilter"; // + id

const resourceTypeFilter = {
    actions: {
        createResourceTypeFilter(context, {data, filter, header}) {

            return api.create(createUrl, data, filter, header)
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

            return api.remove(deleteUrl + id, header)
                .then(function(response) {
                    context.commit("clearResourceTypeFilter");
                    return response;
                })
                .catch(function(err) {
                    console.error(err);
                    throw(err);
                });
        },
        findResourceTypeFilter(context, {filter, header}) {

            return api.find(findUrl, filter, header)
                .then(function(response) {

                    context.commit("setResourceTypeFilter__List", response.List);
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

            return api.update(updateUrl + id, data, filter, header)
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
