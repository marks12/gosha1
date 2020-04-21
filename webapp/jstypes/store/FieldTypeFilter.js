
import {FieldTypeFilter} from "../apiModel";
import api from "../api";
import {findItemIndex} from "../common";

let findUrl = "/api/v1/fieldTypeFilter";
let readUrl = "/api/v1/fieldTypeFilter/"; // + id
let createUrl = "/api/v1/fieldTypeFilter";
let multiCreateUrl = "/api/v1/fieldTypeFilter/list";
let updateUrl = "/api/v1/fieldTypeFilter/"; // + id
let multiUpdateUrl = "/api/v1/fieldTypeFilter/list"; // + id
let deleteUrl = "/api/v1/fieldTypeFilter/"; // + id
let multiDeleteUrl = "/api/v1/fieldTypeFilter/list"; // + id
let findOrCreateUrl = "/api/v1/fieldTypeFilter"; // + id

const fieldTypeFilter = {
    actions: {
        createFieldTypeFilter(context, {data, filter, header}) {

            let url = createUrl;
            if (Array.isArray && Array.isArray(data)) {
                url = multiCreateUrl
            }

            return api.create(url, data, filter, header)
                .then(function(response) {

                    context.commit("setFieldTypeFilter", response.Model);

                    return response;
                })
                .catch(function(err) {
                    console.error(err);
                    throw(err);
                });
        },
        deleteFieldTypeFilter(context, {id, header}) {

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
                    context.commit("clearFieldTypeFilter");
                    return response;
                })
                .catch(function(err) {
                    console.error(err);
                    throw(err);
                });
        },
        findFieldTypeFilter(context, {filter, header, isAppend}) {

            return api.find(findUrl, filter, header)
                .then(function(response) {

                    if (isAppend) {
                        context.commit("appendFieldTypeFilter__List", response.List);
                    } else {
                        context.commit("setFieldTypeFilter__List", response.List);
                    }

                    return response;
                })
                .catch(function(err) {
                    console.error(err);
                    throw(err);
                });
        },
        loadFieldTypeFilter(context, {id, filter, header}) {

            return api.find(readUrl + id, filter, header)
                .then(function(response) {

                    context.commit("setFieldTypeFilter", response.Model);
                    return response;
                })
                .catch(function(err) {
                    console.error(err);
                    throw(err);
                });
        },
        updateFieldTypeFilter(context, {id, data, filter, header}) {

            let url = updateUrl + id;
            if (Array.isArray && Array.isArray(data)) {
                url = multiUpdateUrl
            }

            return api.update(url, data, filter, header)
                .then(function(response) {

                    context.commit("setFieldTypeFilter", response.Model);
                    return response;
                })
                .catch(function(err) {
                    console.error(err);
                    throw(err);
                });
        },
        findOrCreateFieldTypeFilter(context, {id, data, filter, header}) {

            return api.update(findOrCreateUrl, data, filter, header)
                .then(function(response) {

                    context.commit("setFieldTypeFilter", response.Model);
                    return response;
                })
                .catch(function(err) {
                    console.error(err);
                    throw(err);
                });
        },
        clearListFieldTypeFilter(context) {
            context.commit("clearListFieldTypeFilter");
        },
        clearFieldTypeFilter(context) {
            context.commit("clearFieldTypeFilter");
        },
    },
    getters: {
        getFieldTypeFilter: (state) => {
            return state.FieldTypeFilter;
        },
        getFieldTypeFilterById: state => id => {
            return state.FieldTypeFilter__List.find(item => item.Id === id);
        },
        getListFieldTypeFilter: (state) => {
            return state.FieldTypeFilter__List;
        },
        getRoute__FieldTypeFilter: state => action => {
            return state.FieldTypeFilter__Routes[action];
        },
        getRoutes__FieldTypeFilter: state => {
            return state.FieldTypeFilter__Routes;
        },
    },
    mutations: {
        setFieldTypeFilter(state, data) {
            state.FieldTypeFilter = data;
        },
        setFieldTypeFilter__List(state, data) {
            state.FieldTypeFilter__List = data || [];
        },
        appendFieldTypeFilter__List(state, data) {

            if (! state.FieldTypeFilter__List) {
                state.FieldTypeFilter__List = [];
            }

            state.FieldTypeFilter__List = state.FieldTypeFilter__List.concat(data);
        },
        clearFieldTypeFilter(state) {
            state.FieldTypeFilter = new FieldTypeFilter();
        },
        clearListFieldTypeFilter(state) {
            state.FieldTypeFilter__List = [];
        },
		updateFieldTypeFilterById(state, data) {
    		let index = findItemIndex(state.FieldTypeFilter__List, function(item) {
	        	return item.Id === data.Id;
	    	});
	    
	    	if (index || index === 0) {
		        state.FieldTypeFilter__List.splice(index, 1, data);
    		}
		},
		deleteFieldTypeFilterFromList(state, id) {
		    let index = findItemIndex(state.FieldTypeFilter__List, function(item) {
		        return item.Id === id;
		    });
		    
		    if (index || index === 0) {
		        state.FieldTypeFilter__List.splice(index, 1);
		    }
		},
		addFieldTypeFilterItemToList(state, item) {

			if (state.FieldTypeFilter__List === null) {
				state.FieldTypeFilter__List = [];
			}

		    state.FieldTypeFilter__List.push(item);
		},
    },
    state: {
        FieldTypeFilter: new FieldTypeFilter(),
        FieldTypeFilter__List: [],
        FieldTypeFilter__Routes: {
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

export default fieldTypeFilter;
