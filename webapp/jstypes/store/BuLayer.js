
import {BuLayer} from "../apiModel";
import api from "../api";
import {findItemIndex} from "../common";

let findUrl = "/api/v1/buLayer";
let readUrl = "/api/v1/buLayer/"; // + id
let createUrl = "/api/v1/buLayer";
let updateUrl = "/api/v1/buLayer/"; // + id
let deleteUrl = "/api/v1/buLayer/"; // + id
let findOrCreateUrl = "/api/v1/buLayer"; // + id

const buLayer = {
    actions: {
        createBuLayer(context, {data, filter, header}) {

            return api.create(createUrl, data, filter, header)
                .then(function(response) {

                    context.commit("setBuLayer", response.Model);

                    return response;
                })
                .catch(function(err) {
                    return err;
                });
        },
        deleteBuLayer(context, {id, header}) {

            return api.remove(deleteUrl + id, header)
                .then(function(response) {
                    context.commit("clearBuLayer");
                    return response;
                })
                .catch(function(err) {
                    return err;
                });
        },
        findBuLayer(context, {filter, header}) {

            return api.find(findUrl, filter, header)
                .then(function(response) {

                    context.commit("setBuLayer__List", response.List);

                    return response;
                })
                .catch(function(err) {
                    return err;
                });
        },
        loadBuLayer(context, {id, filter, header}) {

            return api.find(readUrl + id, filter, header)
                .then(function(response) {

                    context.commit("setBuLayer", response.Model);

                    return response;
                })
                .catch(function(err) {
                    return err;
                });
        },
        updateBuLayer(context, {id, data, filter, header}) {

            return api.update(updateUrl + id, data, filter, header)
                .then(function(response) {

                    context.commit("setBuLayer", response.Model);

                    return response;
                })
                .catch(function(err) {
                    return err;
                });
        },
        findOrCreateBuLayer(context, {id, data, filter, header}) {

            return api.update(findOrCreateUrl, data, filter, header)
                .then(function(response) {

                    context.commit("setBuLayer", response.Model);

                    return response;
                })
                .catch(function(err) {
                    return err;
                });
        },
        clearListBuLayer(context) {
            context.commit("clearListBuLayer");
        },
    },
    getters: {
        getBuLayer: (state) => {
            return state.BuLayer;
        },
        getListBuLayer: (state) => {
            return state.BuLayer__List;
        },
    },
    mutations: {
        setBuLayer(state, data) {
            state.BuLayer = data;
        },
        setBuLayer__List(state, data) {
            state.BuLayer__List = data || [];
        },
        clearBuLayer(state) {
            state.BuLayer = new BuLayer();
        },
        clearListBuLayer(state) {
            state.BuLayer__List = [];
        },
		updateBuLayerById(state, data) {
    		let index = findItemIndex(state.BuLayer__List, function(item) {
	        	return item.Id === data.Id;
	    	});
	    
	    	if (index || index === 0) {
		        state.BuLayer__List.splice(index, 1, data);
    		}
		},
		deleteBuLayerFromList(state, id) {
		    let index = findItemIndex(state.BuLayer__List, function(item) {
		        return item.Id === id;
		    });
		    
		    if (index || index === 0) {
		        state.BuLayer__List.splice(index, 1);
		    }
		},
		addBuLayerItemToList(state, item) {

			if (state.BuLayer__List === null) {
				state.BuLayer__List = [];
			}

		    state.BuLayer__List.push(item);
		},
    },
    state: {
        BuLayer: new BuLayer(),
        BuLayer__List: [],
    },
};

export default buLayer;
