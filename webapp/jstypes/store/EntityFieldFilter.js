
import {EntityFieldFilter} from "../apiModel";
import api from "../api";
import {findItemIndex} from "../common";

let findUrl = "/api/v1/entityFieldFilter";
let readUrl = "/api/v1/entityFieldFilter/"; // + id
let createUrl = "/api/v1/entityFieldFilter";
let updateUrl = "/api/v1/entityFieldFilter/"; // + id
let deleteUrl = "/api/v1/entityFieldFilter/"; // + id
let findOrCreateUrl = "/api/v1/entityFieldFilter"; // + id

const entityFieldFilter = {
    actions: {
        createEntityFieldFilter(context, {data, filter, header}) {

            return api.create(createUrl, data, filter, header)
                .then(function(response) {

                    context.commit("setEntityFieldFilter", response.Model);

                    return response;
                })
                .catch(function(err) {
                    return err;
                });
        },
        deleteEntityFieldFilter(context, {id, header}) {

            return api.remove(deleteUrl + id, header)
                .then(function(response) {
                    context.commit("clearEntityFieldFilter");
                    return response;
                })
                .catch(function(err) {
                    return err;
                });
        },
        findEntityFieldFilter(context, {filter, header}) {

            return api.find(findUrl, filter, header)
                .then(function(response) {

                    context.commit("setEntityFieldFilter__List", response.List);

                    return response;
                })
                .catch(function(err) {
                    return err;
                });
        },
        loadEntityFieldFilter(context, {id, filter, header}) {

            return api.find(readUrl + id, filter, header)
                .then(function(response) {

                    context.commit("setEntityFieldFilter", response.Model);

                    return response;
                })
                .catch(function(err) {
                    return err;
                });
        },
        updateEntityFieldFilter(context, {id, data, filter, header}) {

            return api.update(updateUrl + id, data, filter, header)
                .then(function(response) {

                    context.commit("setEntityFieldFilter", response.Model);

                    return response;
                })
                .catch(function(err) {
                    return err;
                });
        },
        findOrCreateEntityFieldFilter(context, {id, data, filter, header}) {

            return api.update(findOrCreateUrl, data, filter, header)
                .then(function(response) {

                    context.commit("setEntityFieldFilter", response.Model);

                    return response;
                })
                .catch(function(err) {
                    return err;
                });
        },
        clearListEntityFieldFilter(context) {
            context.commit("clearListEntityFieldFilter");
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
    },
    mutations: {
        setEntityFieldFilter(state, data) {
            state.EntityFieldFilter = data;
        },
        setEntityFieldFilter__List(state, data) {
            state.EntityFieldFilter__List = data || [];
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
    },
};

export default entityFieldFilter;
