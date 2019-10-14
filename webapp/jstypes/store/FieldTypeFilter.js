
import {FieldTypeFilter} from "../apiModel";
import api from "../api";
import {findItemIndex} from "../common";

let findUrl = "/api/v1/fieldTypeFilter";
let readUrl = "/api/v1/fieldTypeFilter/"; // + id
let createUrl = "/api/v1/fieldTypeFilter";
let updateUrl = "/api/v1/fieldTypeFilter/"; // + id
let deleteUrl = "/api/v1/fieldTypeFilter/"; // + id
let findOrCreateUrl = "/api/v1/fieldTypeFilter"; // + id

const fieldTypeFilter = {
    actions: {
        createFieldTypeFilter(context, {data, filter, header}) {

            return api.create(createUrl, data, filter, header)
                .then(function(response) {

                    context.commit("setFieldTypeFilter", response.Model);

                    return response;
                })
                .catch(function(err) {
                    return err;
                });
        },
        deleteFieldTypeFilter(context, {id, header}) {

            return api.remove(deleteUrl + id, header)
                .then(function(response) {
                    context.commit("clearFieldTypeFilter");
                    return response;
                })
                .catch(function(err) {
                    return err;
                });
        },
        findFieldTypeFilter(context, {filter, header}) {

            return api.find(findUrl, filter, header)
                .then(function(response) {

                    context.commit("setFieldTypeFilter__List", response.List);

                    return response;
                })
                .catch(function(err) {
                    return err;
                });
        },
        loadFieldTypeFilter(context, {id, filter, header}) {

            return api.find(readUrl + id, filter, header)
                .then(function(response) {

                    context.commit("setFieldTypeFilter", response.Model);

                    return response;
                })
                .catch(function(err) {
                    return err;
                });
        },
        updateFieldTypeFilter(context, {id, data, filter, header}) {

            return api.update(updateUrl + id, data, filter, header)
                .then(function(response) {

                    context.commit("setFieldTypeFilter", response.Model);

                    return response;
                })
                .catch(function(err) {
                    return err;
                });
        },
        findOrCreateFieldTypeFilter(context, {id, data, filter, header}) {

            return api.update(findOrCreateUrl, data, filter, header)
                .then(function(response) {

                    context.commit("setFieldTypeFilter", response.Model);

                    return response;
                })
                .catch(function(err) {
                    return err;
                });
        },
        clearListFieldTypeFilter(context) {
            context.commit("clearListFieldTypeFilter");
        },
    },
    getters: {
        getFieldTypeFilter: (state) => {
            return state.FieldTypeFilter;
        },
        getListFieldTypeFilter: (state) => {
            return state.FieldTypeFilter__List;
        },
    },
    mutations: {
        setFieldTypeFilter(state, data) {
            state.FieldTypeFilter = data;
        },
        setFieldTypeFilter__List(state, data) {
            state.FieldTypeFilter__List = data || [];
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
    },
};

export default fieldTypeFilter;
