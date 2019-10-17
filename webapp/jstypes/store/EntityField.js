
import {EntityField} from "../apiModel";
import api from "../api";
import {findItemIndex} from "../common";

let findUrl = "/api/v1/entityField";
let readUrl = "/api/v1/entityField/"; // + id
let createUrl = "/api/v1/entityField";
let updateUrl = "/api/v1/entityField/"; // + id
let deleteUrl = "/api/v1/entityField/"; // + id
let findOrCreateUrl = "/api/v1/entityField"; // + id

const entityField = {
    actions: {
        createEntityField(context, {data, filter, header}) {

            return api.create(createUrl, data, filter, header)
                .then(function(response) {

                    context.commit("setEntityField", response.Model);

                    return response;
                })
                .catch(function(err) {
                    return err;
                });
        },
        deleteEntityField(context, {id, header}) {

            return api.remove(deleteUrl + id, header)
                .then(function(response) {
                    context.commit("clearEntityField");
                    return response;
                })
                .catch(function(err) {
                    return err;
                });
        },
        findEntityField(context, {filter, header}) {

            return api.find(findUrl, filter, header)
                .then(function(response) {

                    context.commit("setEntityField__List", response.List);

                    return response;
                })
                .catch(function(err) {
                    return err;
                });
        },
        loadEntityField(context, {id, filter, header}) {

            return api.find(readUrl + id, filter, header)
                .then(function(response) {

                    context.commit("setEntityField", response.Model);

                    return response;
                })
                .catch(function(err) {
                    return err;
                });
        },
        updateEntityField(context, {id, data, filter, header}) {

            return api.update(updateUrl + id, data, filter, header)
                .then(function(response) {

                    context.commit("setEntityField", response.Model);

                    return response;
                })
                .catch(function(err) {
                    return err;
                });
        },
        findOrCreateEntityField(context, {id, data, filter, header}) {

            return api.update(findOrCreateUrl, data, filter, header)
                .then(function(response) {

                    context.commit("setEntityField", response.Model);

                    return response;
                })
                .catch(function(err) {
                    return err;
                });
        },
        clearListEntityField(context) {
            context.commit("clearListEntityField");
        },
    },
    getters: {
        getEntityField: (state) => {
            return state.EntityField;
        },
        getEntityFieldById: state => id => {
            return state.EntityField__List.find(item => item.Id === id);
        },
        getListEntityField: (state) => {
            return state.EntityField__List;
        },
    },
    mutations: {
        setEntityField(state, data) {
            state.EntityField = data;
        },
        setEntityField__List(state, data) {
            state.EntityField__List = data || [];
        },
        clearEntityField(state) {
            state.EntityField = new EntityField();
        },
        clearListEntityField(state) {
            state.EntityField__List = [];
        },
		updateEntityFieldById(state, data) {
    		let index = findItemIndex(state.EntityField__List, function(item) {
	        	return item.Id === data.Id;
	    	});
	    
	    	if (index || index === 0) {
		        state.EntityField__List.splice(index, 1, data);
    		}
		},
		deleteEntityFieldFromList(state, id) {
		    let index = findItemIndex(state.EntityField__List, function(item) {
		        return item.Id === id;
		    });
		    
		    if (index || index === 0) {
		        state.EntityField__List.splice(index, 1);
		    }
		},
		addEntityFieldItemToList(state, item) {

			if (state.EntityField__List === null) {
				state.EntityField__List = [];
			}

		    state.EntityField__List.push(item);
		},
    },
    state: {
        EntityField: new EntityField(),
        EntityField__List: [],
    },
};

export default entityField;
