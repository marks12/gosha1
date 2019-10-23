
import {Field} from "../apiModel";
import api from "../api";
import {findItemIndex} from "../common";

let findUrl = "/api/v1/field";
let readUrl = "/api/v1/field/"; // + id
let createUrl = "/api/v1/field";
let updateUrl = "/api/v1/field/"; // + id
let deleteUrl = "/api/v1/field/"; // + id
let findOrCreateUrl = "/api/v1/field"; // + id

const field = {
    actions: {
        createField(context, {data, filter, header}) {

            return api.create(createUrl, data, filter, header)
                .then(function(response) {

                    context.commit("setField", response.Model);

                    return response;
                })
                .catch(function(err) {
                    console.error(err);
                    throw(err);
                });
        },
        deleteField(context, {id, header}) {

            return api.remove(deleteUrl + id, header)
                .then(function(response) {
                    context.commit("clearField");
                    return response;
                })
                .catch(function(err) {
                    console.error(err);
                    throw(err);
                });
        },
        findField(context, {filter, header}) {

            return api.find(findUrl, filter, header)
                .then(function(response) {

                    context.commit("setField__List", response.List);
                    return response;
                })
                .catch(function(err) {
                    console.error(err);
                    throw(err);
                });
        },
        loadField(context, {id, filter, header}) {

            return api.find(readUrl + id, filter, header)
                .then(function(response) {

                    context.commit("setField", response.Model);
                    return response;
                })
                .catch(function(err) {
                    console.error(err);
                    throw(err);
                });
        },
        updateField(context, {id, data, filter, header}) {

            return api.update(updateUrl + id, data, filter, header)
                .then(function(response) {

                    context.commit("setField", response.Model);
                    return response;
                })
                .catch(function(err) {
                    console.error(err);
                    throw(err);
                });
        },
        findOrCreateField(context, {id, data, filter, header}) {

            return api.update(findOrCreateUrl, data, filter, header)
                .then(function(response) {

                    context.commit("setField", response.Model);
                    return response;
                })
                .catch(function(err) {
                    console.error(err);
                    throw(err);
                });
        },
        clearListField(context) {
            context.commit("clearListField");
        },
    },
    getters: {
        getField: (state) => {
            return state.Field;
        },
        getFieldById: state => id => {
            return state.Field__List.find(item => item.Id === id);
        },
        getListField: (state) => {
            return state.Field__List;
        },
    },
    mutations: {
        setField(state, data) {
            state.Field = data;
        },
        setField__List(state, data) {
            state.Field__List = data || [];
        },
        clearField(state) {
            state.Field = new Field();
        },
        clearListField(state) {
            state.Field__List = [];
        },
		updateFieldById(state, data) {
    		let index = findItemIndex(state.Field__List, function(item) {
	        	return item.Id === data.Id;
	    	});
	    
	    	if (index || index === 0) {
		        state.Field__List.splice(index, 1, data);
    		}
		},
		deleteFieldFromList(state, id) {
		    let index = findItemIndex(state.Field__List, function(item) {
		        return item.Id === id;
		    });
		    
		    if (index || index === 0) {
		        state.Field__List.splice(index, 1);
		    }
		},
		addFieldItemToList(state, item) {

			if (state.Field__List === null) {
				state.Field__List = [];
			}

		    state.Field__List.push(item);
		},
    },
    state: {
        Field: new Field(),
        Field__List: [],
    },
};

export default field;
