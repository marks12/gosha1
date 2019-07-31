
import {EntityFilter} from "../apiModel";
import api from "../api";
import {findItemIndex} from "../common";

let findUrl = "/api/v1/entityFilter";
let readUrl = "/api/v1/entityFilter/"; // + id
let createUrl = "/api/v1/entityFilter";
let updateUrl = "/api/v1/entityFilter/"; // + id
let deleteUrl = "/api/v1/entityFilter/"; // + id
let findOrCreateUrl = "/api/v1/entityFilter"; // + id

const entityFilter = {
    actions: {
        createEntityFilter(context, {data, filter, header}) {

            return api.create(createUrl, data, filter, header)
                .then(function(response) {

                    context.commit("setEntityFilter", response.Model);

                    return response;
                })
                .catch(function(err) {
                    return err;
                });
        },
        deleteEntityFilter(context, {id, header}) {

            return api.remove(deleteUrl + id, header)
                .then(function(response) {
                    context.commit("clearEntityFilter");
                    return response;
                })
                .catch(function(err) {
                    return err;
                });
        },
        findEntityFilter(context, {filter, header}) {

            return api.find(findUrl, filter, header)
                .then(function(response) {

                    context.commit("setEntityFilterList", response.List);

                    return response;
                })
                .catch(function(err) {
                    return err;
                });
        },
        loadEntityFilter(context, {id, filter, header}) {

            return api.find(readUrl + id, filter, header)
                .then(function(response) {

                    context.commit("setEntityFilter", response.Model);

                    return response;
                })
                .catch(function(err) {
                    return err;
                });
        },
        updateEntityFilter(context, {id, data, filter, header}) {

            return api.update(updateUrl + id, data, filter, header)
                .then(function(response) {

                    context.commit("setEntityFilter", response.Model);

                    return response;
                })
                .catch(function(err) {
                    return err;
                });
        },
        findOrCreateEntityFilter(context, {id, data, filter, header}) {

            return api.update(findOrCreateUrl, data, filter, header)
                .then(function(response) {

                    context.commit("setEntityFilter", response.Model);

                    return response;
                })
                .catch(function(err) {
                    return err;
                });
        },
        clearListEntityFilter(context) {
            context.commit("clearListEntityFilter");
        },
    },
    getters: {
        getEntityFilter: (state) => {
            return state.EntityFilter;
        },
        getListEntityFilter: (state) => {
            return state.EntityFilterList;
        },
    },
    mutations: {
        setEntityFilter(state, data) {
            state.EntityFilter = data;
        },
        setEntityFilterList(state, data) {
            state.EntityFilterList = data;
        },
        clearEntityFilter(state) {
            state.EntityFilter = new EntityFilter();
        },
        clearListEntityFilter(state) {
            state.EntityFilterList = [];
        },
		updateEntityFilterById(state, data) {
    		let index = findItemIndex(state.EntityFilterList, function(item) {
	        	return item.Id === data.Id;
	    	});
	    
	    	if (index || index === 0) {
		        state.EntityFilterList.splice(index, 1, data);
    		}
		},
		deleteEntityFilterFromList(state, id) {
		    let index = findItemIndex(state.EntityFilterList, function(item) {
		        return item.Id === id;
		    });
		    
		    if (index || index === 0) {
		        state.EntityFilterList.splice(index, 1);
		    }
		},
		addEntityFilterItemToList(state, item) {

			if (state.EntityFilterList === null) {
				state.EntityFilterList = [];
			}

		    state.EntityFilterList.push(item);
		},
    },
    state: {
        EntityFilter: new EntityFilter(),
        EntityFilterList: [],
    },
};

export default entityFilter;
