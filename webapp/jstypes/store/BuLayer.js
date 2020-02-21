
import {BuLayer} from "../apiModel";
import api from "../api";
import {findItemIndex} from "../common";

let findUrl = "/api/v1/buLayer";
let readUrl = "/api/v1/buLayer/"; // + id
let createUrl = "/api/v1/buLayer";
let multiCreateUrl = "/api/v1/buLayer/list";
let updateUrl = "/api/v1/buLayer/"; // + id
let multiUpdateUrl = "/api/v1/buLayer/list"; // + id
let deleteUrl = "/api/v1/buLayer/"; // + id
let multiDeleteUrl = "/api/v1/buLayer/list"; // + id
let findOrCreateUrl = "/api/v1/buLayer"; // + id

const buLayer = {
    actions: {
        createBuLayer(context, {data, filter, header}) {

            let url = createUrl;
            if (Array.isArray && Array.isArray(data)) {
                url = multiCreateUrl
            }

            return api.create(url, data, filter, header)
                .then(function(response) {

                    context.commit("setBuLayer", response.Model);

                    return response;
                })
                .catch(function(err) {
                    console.error(err);
                    throw(err);
                });
        },
        deleteBuLayer(context, {id, header}) {

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
                    context.commit("clearBuLayer");
                    return response;
                })
                .catch(function(err) {
                    console.error(err);
                    throw(err);
                });
        },
        findBuLayer(context, {filter, header, isAppend}) {

            return api.find(findUrl, filter, header)
                .then(function(response) {

                    if (isAppend) {
                        context.commit("appendBuLayer__List", response.List);
                    } else {
                        context.commit("setBuLayer__List", response.List);
                    }

                    return response;
                })
                .catch(function(err) {
                    console.error(err);
                    throw(err);
                });
        },
        loadBuLayer(context, {id, filter, header}) {

            return api.find(readUrl + id, filter, header)
                .then(function(response) {

                    context.commit("setBuLayer", response.Model);
                    return response;
                })
                .catch(function(err) {
                    console.error(err);
                    throw(err);
                });
        },
        updateBuLayer(context, {id, data, filter, header}) {

            let url = updateUrl + id;
            if (Array.isArray && Array.isArray(data)) {
                url = multiUpdateUrl
            }

            return api.update(url, data, filter, header)
                .then(function(response) {

                    context.commit("setBuLayer", response.Model);
                    return response;
                })
                .catch(function(err) {
                    console.error(err);
                    throw(err);
                });
        },
        findOrCreateBuLayer(context, {id, data, filter, header}) {

            return api.update(findOrCreateUrl, data, filter, header)
                .then(function(response) {

                    context.commit("setBuLayer", response.Model);
                    return response;
                })
                .catch(function(err) {
                    console.error(err);
                    throw(err);
                });
        },
        clearListBuLayer(context) {
            context.commit("clearListBuLayer");
        },
        clearBuLayer(context) {
            context.commit("clearBuLayer");
        },
    },
    getters: {
        getBuLayer: (state) => {
            return state.BuLayer;
        },
        getBuLayerById: state => id => {
            return state.BuLayer__List.find(item => item.Id === id);
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
        appendBuLayer__List(state, data) {

            if (! state.BuLayer__List) {
                state.BuLayer__List = [];
            }

            state.BuLayer__List = state.BuLayer__List.concat(data);
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
