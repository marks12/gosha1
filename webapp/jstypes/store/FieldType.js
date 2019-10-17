
import {FieldType} from "../apiModel";
import api from "../api";
import {findItemIndex} from "../common";

let findUrl = "/api/v1/fieldType";
let readUrl = "/api/v1/fieldType/"; // + id
let createUrl = "/api/v1/fieldType";
let updateUrl = "/api/v1/fieldType/"; // + id
let deleteUrl = "/api/v1/fieldType/"; // + id
let findOrCreateUrl = "/api/v1/fieldType"; // + id

const fieldType = {
    actions: {
        createFieldType(context, {data, filter, header}) {

            return api.create(createUrl, data, filter, header)
                .then(function(response) {

                    context.commit("setFieldType", response.Model);

                    return response;
                })
                .catch(function(err) {
                    return err;
                });
        },
        deleteFieldType(context, {id, header}) {

            return api.remove(deleteUrl + id, header)
                .then(function(response) {
                    context.commit("clearFieldType");
                    return response;
                })
                .catch(function(err) {
                    return err;
                });
        },
        findFieldType(context, {filter, header}) {

            return api.find(findUrl, filter, header)
                .then(function(response) {

                    context.commit("setFieldType__List", response.List);

                    return response;
                })
                .catch(function(err) {
                    return err;
                });
        },
        loadFieldType(context, {id, filter, header}) {

            return api.find(readUrl + id, filter, header)
                .then(function(response) {

                    context.commit("setFieldType", response.Model);

                    return response;
                })
                .catch(function(err) {
                    return err;
                });
        },
        updateFieldType(context, {id, data, filter, header}) {

            return api.update(updateUrl + id, data, filter, header)
                .then(function(response) {

                    context.commit("setFieldType", response.Model);

                    return response;
                })
                .catch(function(err) {
                    return err;
                });
        },
        findOrCreateFieldType(context, {id, data, filter, header}) {

            return api.update(findOrCreateUrl, data, filter, header)
                .then(function(response) {

                    context.commit("setFieldType", response.Model);

                    return response;
                })
                .catch(function(err) {
                    return err;
                });
        },
        clearListFieldType(context) {
            context.commit("clearListFieldType");
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
