
import {Resource} from "../apiModel";
import api from "../api";
import {findItemIndex} from "../common";

let findUrl = "/api/v1/resource";
let readUrl = "/api/v1/resource/"; // + id
let createUrl = "/api/v1/resource";
let multiCreateUrl = "/api/v1/resource/list";
let updateUrl = "/api/v1/resource/"; // + id
let multiUpdateUrl = "/api/v1/resource/list"; // + id
let deleteUrl = "/api/v1/resource/"; // + id
let multiDeleteUrl = "/api/v1/resource/list"; // + id
let findOrCreateUrl = "/api/v1/resource"; // + id

const resource = {
    actions: {
        createResource(context, {data, filter, header}) {

            let url = createUrl;
            if (Array.isArray && Array.isArray(data)) {
                url = multiCreateUrl
            }

            return api.create(url, data, filter, header)
                .then(function(response) {

                    context.commit("setResource", response.Model);

                    return response;
                })
                .catch(function(err) {
                    console.error(err);
                    throw(err);
                });
        },
        deleteResource(context, {id, header}) {

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
                    context.commit("clearResource");
                    return response;
                })
                .catch(function(err) {
                    console.error(err);
                    throw(err);
                });
        },
        findResource(context, {filter, header, isAppend}) {

            return api.find(findUrl, filter, header)
                .then(function(response) {

                    if (isAppend) {
                        context.commit("appendResource__List", response.List);
                    } else {
                        context.commit("setResource__List", response.List);
                    }

                    return response;
                })
                .catch(function(err) {
                    console.error(err);
                    throw(err);
                });
        },
        loadResource(context, {id, filter, header}) {

            return api.find(readUrl + id, filter, header)
                .then(function(response) {

                    context.commit("setResource", response.Model);
                    return response;
                })
                .catch(function(err) {
                    console.error(err);
                    throw(err);
                });
        },
        updateResource(context, {id, data, filter, header}) {

            let url = updateUrl + id;
            if (Array.isArray && Array.isArray(data)) {
                url = multiUpdateUrl
            }

            return api.update(url, data, filter, header)
                .then(function(response) {

                    context.commit("setResource", response.Model);
                    return response;
                })
                .catch(function(err) {
                    console.error(err);
                    throw(err);
                });
        },
        findOrCreateResource(context, {id, data, filter, header}) {

            return api.update(findOrCreateUrl, data, filter, header)
                .then(function(response) {

                    context.commit("setResource", response.Model);
                    return response;
                })
                .catch(function(err) {
                    console.error(err);
                    throw(err);
                });
        },
        clearListResource(context) {
            context.commit("clearListResource");
        },
        clearResource(context) {
            context.commit("clearResource");
        },
    },
    getters: {
        getResource: (state) => {
            return state.Resource;
        },
        getResourceById: state => id => {
            return state.Resource__List.find(item => item.Id === id);
        },
        getListResource: (state) => {
            return state.Resource__List;
        },
    },
    mutations: {
        setResource(state, data) {
            state.Resource = data;
        },
        setResource__List(state, data) {
            state.Resource__List = data || [];
        },
        appendResource__List(state, data) {

            if (! state.Resource__List) {
                state.Resource__List = [];
            }

            state.Resource__List = state.Resource__List.concat(data);
        },
        clearResource(state) {
            state.Resource = new Resource();
        },
        clearListResource(state) {
            state.Resource__List = [];
        },
		updateResourceById(state, data) {
    		let index = findItemIndex(state.Resource__List, function(item) {
	        	return item.Id === data.Id;
	    	});
	    
	    	if (index || index === 0) {
		        state.Resource__List.splice(index, 1, data);
    		}
		},
		deleteResourceFromList(state, id) {
		    let index = findItemIndex(state.Resource__List, function(item) {
		        return item.Id === id;
		    });
		    
		    if (index || index === 0) {
		        state.Resource__List.splice(index, 1);
		    }
		},
		addResourceItemToList(state, item) {

			if (state.Resource__List === null) {
				state.Resource__List = [];
			}

		    state.Resource__List.push(item);
		},
    },
    state: {
        Resource: new Resource(),
        Resource__List: [],
    },
};

export default resource;
