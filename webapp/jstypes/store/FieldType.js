
import {FieldType} from "../apiModel";
import api from "../api";
import {findItemIndex} from "../common";

let findUrl = "/api/v1/fieldType";
let readUrl = "/api/v1/fieldType/"; // + id
let createUrl = "/api/v1/fieldType";
let multiCreateUrl = "/api/v1/fieldType/list";
let updateUrl = "/api/v1/fieldType/"; // + id
let multiUpdateUrl = "/api/v1/fieldType/list"; // + id
let deleteUrl = "/api/v1/fieldType/"; // + id
let multiDeleteUrl = "/api/v1/fieldType/list"; // + id
let findOrCreateUrl = "/api/v1/fieldType"; // + id

const fieldType = {
    actions: {
        createFieldType(context, {data, filter, header}) {

            let url = createUrl;
            if (Array.isArray && Array.isArray(data)) {
                url = multiCreateUrl
            }

            return api.create(url, data, filter, header)
                .then(function(response) {

                    context.commit("setFieldType", response.Model);

                    return response;
                })
                .catch(function(err) {
                    console.error(err);
                    throw(err);
                });
        },
        deleteFieldType(context, {id, header}) {

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
                    context.commit("clearFieldType");
                    return response;
                })
                .catch(function(err) {
                    console.error(err);
                    throw(err);
                });
        },
        findFieldType(context, {filter, header, isAppend}) {

            return api.find(findUrl, filter, header)
                .then(function(response) {

                    if (isAppend) {
                        context.commit("appendFieldType__List", response.List);
                    } else {
                        context.commit("setFieldType__List", response.List);
                    }

                    return response;
                })
                .catch(function(err) {
                    console.error(err);
                    throw(err);
                });
        },
        loadFieldType(context, {id, filter, header}) {

            return api.find(readUrl + id, filter, header)
                .then(function(response) {

                    context.commit("setFieldType", response.Model);
                    return response;
                })
                .catch(function(err) {
                    console.error(err);
                    throw(err);
                });
        },
        updateFieldType(context, {id, data, filter, header}) {

            let url = updateUrl + id;
            if (Array.isArray && Array.isArray(data)) {
                url = multiUpdateUrl
            }

            return api.update(url, data, filter, header)
                .then(function(response) {

                    context.commit("setFieldType", response.Model);
                    return response;
                })
                .catch(function(err) {
                    console.error(err);
                    throw(err);
                });
        },
        findOrCreateFieldType(context, {id, data, filter, header}) {

            return api.update(findOrCreateUrl, data, filter, header)
                .then(function(response) {

                    context.commit("setFieldType", response.Model);
                    return response;
                })
                .catch(function(err) {
                    console.error(err);
                    throw(err);
                });
        },
        clearListFieldType(context) {
            context.commit("clearListFieldType");
        },
        clearFieldType(context) {
            context.commit("clearFieldType");
        },
    },
    getters: {
        getFieldType: (state) => {
            return state.FieldType;
        },
        getFieldTypeById: state => id => {
            return state.FieldType__List.find(item => item.Id === id);
        },
        getListFieldType: (state) => {
            return state.FieldType__List;
        },
    },
    mutations: {
        setFieldType(state, data) {
            state.FieldType = data;
        },
        setFieldType__List(state, data) {
            state.FieldType__List = data || [];
        },
        appendFieldType__List(state, data) {

            if (! state.FieldType__List) {
                state.FieldType__List = [];
            }

            state.FieldType__List = state.FieldType__List.concat(data);
        },
        clearFieldType(state) {
            state.FieldType = new FieldType();
        },
        clearListFieldType(state) {
            state.FieldType__List = [];
        },
		updateFieldTypeById(state, data) {
    		let index = findItemIndex(state.FieldType__List, function(item) {
	        	return item.Id === data.Id;
	    	});
	    
	    	if (index || index === 0) {
		        state.FieldType__List.splice(index, 1, data);
    		}
		},
		deleteFieldTypeFromList(state, id) {
		    let index = findItemIndex(state.FieldType__List, function(item) {
		        return item.Id === id;
		    });
		    
		    if (index || index === 0) {
		        state.FieldType__List.splice(index, 1);
		    }
		},
		addFieldTypeItemToList(state, item) {

			if (state.FieldType__List === null) {
				state.FieldType__List = [];
			}

		    state.FieldType__List.push(item);
		},
    },
    state: {
        FieldType: new FieldType(),
        FieldType__List: [],
    },
};

export default fieldType;
