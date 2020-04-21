
import {EntityFilter} from "../apiModel";
import api from "../api";
import {findItemIndex} from "../common";

let findUrl = "/api/v1/entityFilter";
let readUrl = "/api/v1/entityFilter/"; // + id
let createUrl = "/api/v1/entityFilter";
let multiCreateUrl = "/api/v1/entityFilter/list";
let updateUrl = "/api/v1/entityFilter/"; // + id
let multiUpdateUrl = "/api/v1/entityFilter/list"; // + id
let deleteUrl = "/api/v1/entityFilter/"; // + id
let multiDeleteUrl = "/api/v1/entityFilter/list"; // + id
let findOrCreateUrl = "/api/v1/entityFilter"; // + id

const entityFilter = {
    actions: {
        createEntityFilter(context, {data, filter, header}) {

            let url = createUrl;
            if (Array.isArray && Array.isArray(data)) {
                url = multiCreateUrl
            }

            return api.create(url, data, filter, header)
                .then(function(response) {

                    context.commit("setEntityFilter", response.Model);

                    return response;
                })
                .catch(function(err) {
                    console.error(err);
                    throw(err);
                });
        },
        deleteEntityFilter(context, {id, header}) {

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
                    context.commit("clearEntityFilter");
                    return response;
                })
                .catch(function(err) {
                    console.error(err);
                    throw(err);
                });
        },
        findEntityFilter(context, {filter, header, isAppend}) {

            return api.find(findUrl, filter, header)
                .then(function(response) {

                    if (isAppend) {
                        context.commit("appendEntityFilter__List", response.List);
                    } else {
                        context.commit("setEntityFilter__List", response.List);
                    }

                    return response;
                })
                .catch(function(err) {
                    console.error(err);
                    throw(err);
                });
        },
        loadEntityFilter(context, {id, filter, header}) {

            return api.find(readUrl + id, filter, header)
                .then(function(response) {

                    context.commit("setEntityFilter", response.Model);
                    return response;
                })
                .catch(function(err) {
                    console.error(err);
                    throw(err);
                });
        },
        updateEntityFilter(context, {id, data, filter, header}) {

            let url = updateUrl + id;
            if (Array.isArray && Array.isArray(data)) {
                url = multiUpdateUrl
            }

            return api.update(url, data, filter, header)
                .then(function(response) {

                    context.commit("setEntityFilter", response.Model);
                    return response;
                })
                .catch(function(err) {
                    console.error(err);
                    throw(err);
                });
        },
        findOrCreateEntityFilter(context, {id, data, filter, header}) {

            return api.update(findOrCreateUrl, data, filter, header)
                .then(function(response) {

                    context.commit("setEntityFilter", response.Model);
                    return response;
                })
                .catch(function(err) {
                    console.error(err);
                    throw(err);
                });
        },
        clearListEntityFilter(context) {
            context.commit("clearListEntityFilter");
        },
        clearEntityFilter(context) {
            context.commit("clearEntityFilter");
        },
    },
    getters: {
        getEntityFilter: (state) => {
            return state.EntityFilter;
        },
        getEntityFilterById: state => id => {
            return state.EntityFilter__List.find(item => item.Id === id);
        },
        getListEntityFilter: (state) => {
            return state.EntityFilter__List;
        },
        getRoute__EntityFilter: state => action => {
            return state.EntityFilter__Routes[action];
        },
        getRoutes__EntityFilter: state => {
            return state.EntityFilter__Routes;
        },
    },
    mutations: {
        setEntityFilter(state, data) {
            state.EntityFilter = data;
        },
        setEntityFilter__List(state, data) {
            state.EntityFilter__List = data || [];
        },
        appendEntityFilter__List(state, data) {

            if (! state.EntityFilter__List) {
                state.EntityFilter__List = [];
            }

            state.EntityFilter__List = state.EntityFilter__List.concat(data);
        },
        clearEntityFilter(state) {
            state.EntityFilter = new EntityFilter();
        },
        clearListEntityFilter(state) {
            state.EntityFilter__List = [];
        },
		updateEntityFilterById(state, data) {
    		let index = findItemIndex(state.EntityFilter__List, function(item) {
	        	return item.Id === data.Id;
	    	});
	    
	    	if (index || index === 0) {
		        state.EntityFilter__List.splice(index, 1, data);
    		}
		},
		deleteEntityFilterFromList(state, id) {
		    let index = findItemIndex(state.EntityFilter__List, function(item) {
		        return item.Id === id;
		    });
		    
		    if (index || index === 0) {
		        state.EntityFilter__List.splice(index, 1);
		    }
		},
		addEntityFilterItemToList(state, item) {

			if (state.EntityFilter__List === null) {
				state.EntityFilter__List = [];
			}

		    state.EntityFilter__List.push(item);
		},
    },
    state: {
        EntityFilter: new EntityFilter(),
        EntityFilter__List: [],
        EntityFilter__Routes: {
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

export default entityFilter;
