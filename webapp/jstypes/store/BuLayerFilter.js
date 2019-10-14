
import {BuLayerFilter} from "../apiModel";
import api from "../api";
import {findItemIndex} from "../common";

let findUrl = "/api/v1/buLayerFilter";
let readUrl = "/api/v1/buLayerFilter/"; // + id
let createUrl = "/api/v1/buLayerFilter";
let updateUrl = "/api/v1/buLayerFilter/"; // + id
let deleteUrl = "/api/v1/buLayerFilter/"; // + id
let findOrCreateUrl = "/api/v1/buLayerFilter"; // + id

const buLayerFilter = {
    actions: {
        createBuLayerFilter(context, {data, filter, header}) {

            return api.create(createUrl, data, filter, header)
                .then(function(response) {

                    context.commit("setBuLayerFilter", response.Model);

                    return response;
                })
                .catch(function(err) {
                    return err;
                });
        },
        deleteBuLayerFilter(context, {id, header}) {

            return api.remove(deleteUrl + id, header)
                .then(function(response) {
                    context.commit("clearBuLayerFilter");
                    return response;
                })
                .catch(function(err) {
                    return err;
                });
        },
        findBuLayerFilter(context, {filter, header}) {

            return api.find(findUrl, filter, header)
                .then(function(response) {

                    context.commit("setBuLayerFilter__List", response.List);

                    return response;
                })
                .catch(function(err) {
                    return err;
                });
        },
        loadBuLayerFilter(context, {id, filter, header}) {

            return api.find(readUrl + id, filter, header)
                .then(function(response) {

                    context.commit("setBuLayerFilter", response.Model);

                    return response;
                })
                .catch(function(err) {
                    return err;
                });
        },
        updateBuLayerFilter(context, {id, data, filter, header}) {

            return api.update(updateUrl + id, data, filter, header)
                .then(function(response) {

                    context.commit("setBuLayerFilter", response.Model);

                    return response;
                })
                .catch(function(err) {
                    return err;
                });
        },
        findOrCreateBuLayerFilter(context, {id, data, filter, header}) {

            return api.update(findOrCreateUrl, data, filter, header)
                .then(function(response) {

                    context.commit("setBuLayerFilter", response.Model);

                    return response;
                })
                .catch(function(err) {
                    return err;
                });
        },
        clearListBuLayerFilter(context) {
            context.commit("clearListBuLayerFilter");
        },
    },
    getters: {
        getBuLayerFilter: (state) => {
            return state.BuLayerFilter;
        },
        getListBuLayerFilter: (state) => {
            return state.BuLayerFilter__List;
        },
    },
    mutations: {
        setBuLayerFilter(state, data) {
            state.BuLayerFilter = data;
        },
        setBuLayerFilter__List(state, data) {
            state.BuLayerFilter__List = data || [];
        },
        clearBuLayerFilter(state) {
            state.BuLayerFilter = new BuLayerFilter();
        },
        clearListBuLayerFilter(state) {
            state.BuLayerFilter__List = [];
        },
		updateBuLayerFilterById(state, data) {
    		let index = findItemIndex(state.BuLayerFilter__List, function(item) {
	        	return item.Id === data.Id;
	    	});
	    
	    	if (index || index === 0) {
		        state.BuLayerFilter__List.splice(index, 1, data);
    		}
		},
		deleteBuLayerFilterFromList(state, id) {
		    let index = findItemIndex(state.BuLayerFilter__List, function(item) {
		        return item.Id === id;
		    });
		    
		    if (index || index === 0) {
		        state.BuLayerFilter__List.splice(index, 1);
		    }
		},
		addBuLayerFilterItemToList(state, item) {

			if (state.BuLayerFilter__List === null) {
				state.BuLayerFilter__List = [];
			}

		    state.BuLayerFilter__List.push(item);
		},
    },
    state: {
        BuLayerFilter: new BuLayerFilter(),
        BuLayerFilter__List: [],
    },
};

export default buLayerFilter;
