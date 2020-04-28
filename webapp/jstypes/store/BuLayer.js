
import {BuLayer} from "../apiModel";
import api from "../api";
import {findItemIndex} from "../common";

let findUrl = "/api/v1/buLayer";
let readUrl = "/api/v1/buLayer/"; // + id
let createUrl = "/api/v1/buLayer";
let multiCreateUrl = "/api/v1/buLayer/list";
let updateUrl = "/api/v1/buLayer/"; // + id
let multiUpdateUrl = "/api/v1/buLayer/list";
let deleteUrl = "/api/v1/buLayer/"; // + id
let multiDeleteUrl = "/api/v1/buLayer/list";
let findOrCreateUrl = "/api/v1/buLayer";
let updateOrCreateUrl = "/api/v1/buLayer";

const buLayer = {
    actions: {
        createBuLayer(context, {data, filter, header, noMutation}) {

            let url = createUrl;
            if (Array.isArray && Array.isArray(data)) {
                url = multiCreateUrl
            }

            return api.create(url, data, filter, header)
                .then(function(response) {

					if(! noMutation) {
	                    context.commit("setBuLayer", response.Model);
					}

                    return response;
                })
                .catch(function(err) {
                    console.error(err);
                    throw(err);
                });
        },
        deleteBuLayer(context, {id, header, noMutation}) {

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
					if(! noMutation) {
	                    context.commit("clearBuLayer");
					}
                    return response;
                })
                .catch(function(err) {
                    console.error(err);
                    throw(err);
                });
        },
        findBuLayer(context, {filter, header, isAppend, noMutation}) {

            return api.find(findUrl, filter, header)
                .then(function(response) {

					if(! noMutation) {
						if (isAppend) {
							context.commit("appendBuLayer__List", response.List);
						} else {
							context.commit("setBuLayer__List", response.List);
						}
					}

                    return response;
                })
                .catch(function(err) {
                    console.error(err);
                    throw(err);
                });
        },
        loadBuLayer(context, {id, filter, header, noMutation}) {

            return api.find(readUrl + id, filter, header)
                .then(function(response) {

					if(! noMutation) {
	                    context.commit("setBuLayer", response.Model);
					}
                    return response;
                })
                .catch(function(err) {
                    console.error(err);
                    throw(err);
                });
        },
        updateBuLayer(context, {id, data, filter, header, noMutation}) {

            let url = updateUrl + id;
            if (Array.isArray && Array.isArray(data)) {
                url = multiUpdateUrl
            }

            return api.update(url, data, filter, header)
                .then(function(response) {
					if(! noMutation) {
	                    context.commit("setBuLayer", response.Model);
					}
                    return response;
                })
                .catch(function(err) {
                    console.error(err);
                    throw(err);
                });
        },
        findOrCreateBuLayer(context, {id, data, filter, header, noMutation}) {

            return api.update(findOrCreateUrl, data, filter, header)
                .then(function(response) {

					if(! noMutation) {
	                    context.commit("setBuLayer", response.Model);
					}
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
        getRoute__BuLayer: state => action => {
            return state.BuLayer__Routes[action];
        },
        getRoutes__BuLayer: state => {
            return state.BuLayer__Routes;
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
        BuLayer__Routes: {
            find: findUrl,
            read: readUrl,
            create: createUrl,
            multiCreate: multiCreateUrl,
            update: updateUrl,
            multiUpdate: multiUpdateUrl,
            delete: deleteUrl,
            multiDelete: multiDeleteUrl,
            findOrCreate: findOrCreateUrl,
            updateOrCreate: updateOrCreateUrl,
        },
    },
};

export default buLayer;
