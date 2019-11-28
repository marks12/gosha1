
import {ResourceType} from "../apiModel";
import api from "../api";
import {findItemIndex} from "../common";

let findUrl = "/api/v1/resourceType";
let readUrl = "/api/v1/resourceType/"; // + id
let createUrl = "/api/v1/resourceType";
let multiCreateUrl = "/api/v1/resourceType/list";
let updateUrl = "/api/v1/resourceType/"; // + id
let multiUpdateUrl = "/api/v1/resourceType/list"; // + id
let deleteUrl = "/api/v1/resourceType/"; // + id
let multiDeleteUrl = "/api/v1/resourceType/list"; // + id
let findOrCreateUrl = "/api/v1/resourceType"; // + id

const resourceType = {
    actions: {
        createResourceType(context, {data, filter, header}) {

            let url = createUrl;
            if (Array.isArray && Array.isArray(data)) {
                url = multiCreateUrl
            }

            return api.create(url, data, filter, header)
                .then(function(response) {

                    context.commit("setResourceType", response.Model);

                    return response;
                })
                .catch(function(err) {
                    console.error(err);
                    throw(err);
                });
        },
        deleteResourceType(context, {id, header}) {

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
                    context.commit("clearResourceType");
                    return response;
                })
                .catch(function(err) {
                    console.error(err);
                    throw(err);
                });
        },
        findResourceType(context, {filter, header}) {

            return api.find(findUrl, filter, header)
                .then(function(response) {

                    context.commit("setResourceType__List", response.List);
                    return response;
                })
                .catch(function(err) {
                    console.error(err);
                    throw(err);
                });
        },
        loadResourceType(context, {id, filter, header}) {

            return api.find(readUrl + id, filter, header)
                .then(function(response) {

                    context.commit("setResourceType", response.Model);
                    return response;
                })
                .catch(function(err) {
                    console.error(err);
                    throw(err);
                });
        },
        updateResourceType(context, {id, data, filter, header}) {

            let url = updateUrl + id;
            if (Array.isArray && Array.isArray(data)) {
                url = multiUpdateUrl
            }

            return api.update(url, data, filter, header)
                .then(function(response) {

                    context.commit("setResourceType", response.Model);
                    return response;
                })
                .catch(function(err) {
                    console.error(err);
                    throw(err);
                });
        },
        findOrCreateResourceType(context, {id, data, filter, header}) {

            return api.update(findOrCreateUrl, data, filter, header)
                .then(function(response) {

                    context.commit("setResourceType", response.Model);
                    return response;
                })
                .catch(function(err) {
                    console.error(err);
                    throw(err);
                });
        },
        clearListResourceType(context) {
            context.commit("clearListResourceType");
        },
        clearResourceType(context) {
            context.commit("clearResourceType");
        },
    },
    getters: {
        getResourceType: (state) => {
            return state.ResourceType;
        },
        getResourceTypeById: state => id => {
            return state.ResourceType__List.find(item => item.Id === id);
        },
        getListResourceType: (state) => {
            return state.ResourceType__List;
        },
    },
    mutations: {
        setResourceType(state, data) {
            state.ResourceType = data;
        },
        setResourceType__List(state, data) {
            state.ResourceType__List = data || [];
        },
        clearResourceType(state) {
            state.ResourceType = new ResourceType();
        },
        clearListResourceType(state) {
            state.ResourceType__List = [];
        },
		updateResourceTypeById(state, data) {
    		let index = findItemIndex(state.ResourceType__List, function(item) {
	        	return item.Id === data.Id;
	    	});
	    
	    	if (index || index === 0) {
		        state.ResourceType__List.splice(index, 1, data);
    		}
		},
		deleteResourceTypeFromList(state, id) {
		    let index = findItemIndex(state.ResourceType__List, function(item) {
		        return item.Id === id;
		    });
		    
		    if (index || index === 0) {
		        state.ResourceType__List.splice(index, 1);
		    }
		},
		addResourceTypeItemToList(state, item) {

			if (state.ResourceType__List === null) {
				state.ResourceType__List = [];
			}

		    state.ResourceType__List.push(item);
		},
    },
    state: {
        ResourceType: new ResourceType(),
        ResourceType__List: [],
    },
};

export default resourceType;
