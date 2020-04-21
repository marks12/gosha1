
import {EntityFieldFilter} from "../apiModel";
import api from "../api";
import {findItemIndex} from "../common";

let findUrl = "/api/v1/entityFieldFilter";
let readUrl = "/api/v1/entityFieldFilter/"; // + id
let createUrl = "/api/v1/entityFieldFilter";
let multiCreateUrl = "/api/v1/entityFieldFilter/list";
let updateUrl = "/api/v1/entityFieldFilter/"; // + id
let multiUpdateUrl = "/api/v1/entityFieldFilter/list"; // + id
let deleteUrl = "/api/v1/entityFieldFilter/"; // + id
let multiDeleteUrl = "/api/v1/entityFieldFilter/list"; // + id
let findOrCreateUrl = "/api/v1/entityFieldFilter"; // + id

const entityFieldFilter = {
    actions: {
        createEntityFieldFilter(context, {data, filter, header}) {

            let url = createUrl;
            if (Array.isArray && Array.isArray(data)) {
                url = multiCreateUrl
            }

            return api.create(url, data, filter, header)
                .then(function(response) {

                    context.commit("setEntityFieldFilter", response.Model);

                    return response;
                })
                .catch(function(err) {
                    console.error(err);
                    throw(err);
                });
        },
        deleteEntityFieldFilter(context, {id, header}) {

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
                    context.commit("clearEntityFieldFilter");
                    return response;
                })
                .catch(function(err) {
                    console.error(err);
                    throw(err);
                });
        },
        findEntityFieldFilter(context, {filter, header, isAppend}) {

            return api.find(findUrl, filter, header)
                .then(function(response) {

                    if (isAppend) {
                        context.commit("appendEntityFieldFilter__List", response.List);
                    } else {
                        context.commit("setEntityFieldFilter__List", response.List);
                    }

                    return response;
                })
                .catch(function(err) {
                    console.error(err);
                    throw(err);
                });
        },
        loadEntityFieldFilter(context, {id, filter, header}) {

            return api.find(readUrl + id, filter, header)
                .then(function(response) {

                    context.commit("setEntityFieldFilter", response.Model);
                    return response;
                })
                .catch(function(err) {
                    console.error(err);
                    throw(err);
                });
        },
        updateEntityFieldFilter(context, {id, data, filter, header}) {

            let url = updateUrl + id;
            if (Array.isArray && Array.isArray(data)) {
                url = multiUpdateUrl
            }

            return api.update(url, data, filter, header)
                .then(function(response) {

                    context.commit("setEntityFieldFilter", response.Model);
                    return response;
                })
                .catch(function(err) {
                    console.error(err);
                    throw(err);
                });
        },
        findOrCreateEntityFieldFilter(context, {id, data, filter, header}) {

            return api.update(findOrCreateUrl, data, filter, header)
                .then(function(response) {

                    context.commit("setEntityFieldFilter", response.Model);
                    return response;
                })
                .catch(function(err) {
                    console.error(err);
                    throw(err);
                });
        },
        clearListEntityFieldFilter(context) {
            context.commit("clearListEntityFieldFilter");
        },
        clearEntityFieldFilter(context) {
            context.commit("clearEntityFieldFilter");
        },
    },
    getters: {
        getEntityFieldFilter: (state) => {
            return state.EntityFieldFilter;
        },
        getEntityFieldFilterById: state => id => {
            return state.EntityFieldFilter__List.find(item => item.Id === id);
        },
        getListEntityFieldFilter: (state) => {
            return state.EntityFieldFilter__List;
        },
        getRoute__EntityFieldFilter: state => action => {
            return state.EntityFieldFilter__Routes[action];
        },
        getRoutes__EntityFieldFilter: state => {
            return state.EntityFieldFilter__Routes;
        },
    },
    mutations: {
        setEntityFieldFilter(state, data) {
            state.EntityFieldFilter = data;
        },
        setEntityFieldFilter__List(state, data) {
            state.EntityFieldFilter__List = data || [];
        },
        appendEntityFieldFilter__List(state, data) {

            if (! state.EntityFieldFilter__List) {
                state.EntityFieldFilter__List = [];
            }

            state.EntityFieldFilter__List = state.EntityFieldFilter__List.concat(data);
        },
        clearEntityFieldFilter(state) {
            state.EntityFieldFilter = new EntityFieldFilter();
        },
        clearListEntityFieldFilter(state) {
            state.EntityFieldFilter__List = [];
        },
		updateEntityFieldFilterById(state, data) {
    		let index = findItemIndex(state.EntityFieldFilter__List, function(item) {
	        	return item.Id === data.Id;
	    	});
	    
	    	if (index || index === 0) {
		        state.EntityFieldFilter__List.splice(index, 1, data);
    		}
		},
		deleteEntityFieldFilterFromList(state, id) {
		    let index = findItemIndex(state.EntityFieldFilter__List, function(item) {
		        return item.Id === id;
		    });
		    
		    if (index || index === 0) {
		        state.EntityFieldFilter__List.splice(index, 1);
		    }
		},
		addEntityFieldFilterItemToList(state, item) {

			if (state.EntityFieldFilter__List === null) {
				state.EntityFieldFilter__List = [];
			}

		    state.EntityFieldFilter__List.push(item);
		},
    },
    state: {
        EntityFieldFilter: new EntityFieldFilter(),
        EntityFieldFilter__List: [],
        EntityFieldFilter__Routes: {
            find: findUrl,
            read: readUrl,
            create: createUrl,
            multiCreate: multiCreateUrl,
            update: updateUrl,
            multiUpdate: multiUpdateUrl,
            delete: deleteUrl,
            multiDelete: multiDeleteUrl,
            findOrCreate: findOrCreateUrl,
        },
    },
};

export default entityFieldFilter;
