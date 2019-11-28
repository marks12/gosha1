
import {ResourceFilter} from "../apiModel";
import api from "../api";
import {findItemIndex} from "../common";

let findUrl = "/api/v1/resourceFilter";
let readUrl = "/api/v1/resourceFilter/"; // + id
let createUrl = "/api/v1/resourceFilter";
let multiCreateUrl = "/api/v1/resourceFilter/list";
let updateUrl = "/api/v1/resourceFilter/"; // + id
let multiUpdateUrl = "/api/v1/resourceFilter/list"; // + id
let deleteUrl = "/api/v1/resourceFilter/"; // + id
let multiDeleteUrl = "/api/v1/resourceFilter/list"; // + id
let findOrCreateUrl = "/api/v1/resourceFilter"; // + id

const resourceFilter = {
    actions: {
        createResourceFilter(context, {data, filter, header}) {

            let url = createUrl;
            if (Array.isArray && Array.isArray(data)) {
                url = multiCreateUrl
            }

            return api.create(url, data, filter, header)
                .then(function(response) {

                    context.commit("setResourceFilter", response.Model);

                    return response;
                })
                .catch(function(err) {
                    console.error(err);
                    throw(err);
                });
        },
        deleteResourceFilter(context, {id, header}) {

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
                    context.commit("clearResourceFilter");
                    return response;
                })
                .catch(function(err) {
                    console.error(err);
                    throw(err);
                });
        },
        findResourceFilter(context, {filter, header}) {

            return api.find(findUrl, filter, header)
                .then(function(response) {

                    context.commit("setResourceFilter__List", response.List);
                    return response;
                })
                .catch(function(err) {
                    console.error(err);
                    throw(err);
                });
        },
        loadResourceFilter(context, {id, filter, header}) {

            return api.find(readUrl + id, filter, header)
                .then(function(response) {

                    context.commit("setResourceFilter", response.Model);
                    return response;
                })
                .catch(function(err) {
                    console.error(err);
                    throw(err);
                });
        },
        updateResourceFilter(context, {id, data, filter, header}) {

            let url = updateUrl + id;
            if (Array.isArray && Array.isArray(data)) {
                url = multiUpdateUrl
            }

            return api.update(url, data, filter, header)
                .then(function(response) {

                    context.commit("setResourceFilter", response.Model);
                    return response;
                })
                .catch(function(err) {
                    console.error(err);
                    throw(err);
                });
        },
        findOrCreateResourceFilter(context, {id, data, filter, header}) {

            return api.update(findOrCreateUrl, data, filter, header)
                .then(function(response) {

                    context.commit("setResourceFilter", response.Model);
                    return response;
                })
                .catch(function(err) {
                    console.error(err);
                    throw(err);
                });
        },
        clearListResourceFilter(context) {
            context.commit("clearListResourceFilter");
        },
        clearResourceFilter(context) {
            context.commit("clearResourceFilter");
        },
    },
    getters: {
        getResourceFilter: (state) => {
            return state.ResourceFilter;
        },
        getResourceFilterById: state => id => {
            return state.ResourceFilter__List.find(item => item.Id === id);
        },
        getListResourceFilter: (state) => {
            return state.ResourceFilter__List;
        },
    },
    mutations: {
        setResourceFilter(state, data) {
            state.ResourceFilter = data;
        },
        setResourceFilter__List(state, data) {
            state.ResourceFilter__List = data || [];
        },
        clearResourceFilter(state) {
            state.ResourceFilter = new ResourceFilter();
        },
        clearListResourceFilter(state) {
            state.ResourceFilter__List = [];
        },
		updateResourceFilterById(state, data) {
    		let index = findItemIndex(state.ResourceFilter__List, function(item) {
	        	return item.Id === data.Id;
	    	});
	    
	    	if (index || index === 0) {
		        state.ResourceFilter__List.splice(index, 1, data);
    		}
		},
		deleteResourceFilterFromList(state, id) {
		    let index = findItemIndex(state.ResourceFilter__List, function(item) {
		        return item.Id === id;
		    });
		    
		    if (index || index === 0) {
		        state.ResourceFilter__List.splice(index, 1);
		    }
		},
		addResourceFilterItemToList(state, item) {

			if (state.ResourceFilter__List === null) {
				state.ResourceFilter__List = [];
			}

		    state.ResourceFilter__List.push(item);
		},
    },
    state: {
        ResourceFilter: new ResourceFilter(),
        ResourceFilter__List: [],
    },
};

export default resourceFilter;
