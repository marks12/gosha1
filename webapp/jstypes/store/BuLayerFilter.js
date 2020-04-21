
import {BuLayerFilter} from "../apiModel";
import api from "../api";
import {findItemIndex} from "../common";

let findUrl = "/api/v1/buLayerFilter";
let readUrl = "/api/v1/buLayerFilter/"; // + id
let createUrl = "/api/v1/buLayerFilter";
let multiCreateUrl = "/api/v1/buLayerFilter/list";
let updateUrl = "/api/v1/buLayerFilter/"; // + id
let multiUpdateUrl = "/api/v1/buLayerFilter/list"; // + id
let deleteUrl = "/api/v1/buLayerFilter/"; // + id
let multiDeleteUrl = "/api/v1/buLayerFilter/list"; // + id
let findOrCreateUrl = "/api/v1/buLayerFilter"; // + id

const buLayerFilter = {
    actions: {
        createBuLayerFilter(context, {data, filter, header}) {

            let url = createUrl;
            if (Array.isArray && Array.isArray(data)) {
                url = multiCreateUrl
            }

            return api.create(url, data, filter, header)
                .then(function(response) {

                    context.commit("setBuLayerFilter", response.Model);

                    return response;
                })
                .catch(function(err) {
                    console.error(err);
                    throw(err);
                });
        },
        deleteBuLayerFilter(context, {id, header}) {

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
                    context.commit("clearBuLayerFilter");
                    return response;
                })
                .catch(function(err) {
                    console.error(err);
                    throw(err);
                });
        },
        findBuLayerFilter(context, {filter, header, isAppend}) {

            return api.find(findUrl, filter, header)
                .then(function(response) {

                    if (isAppend) {
                        context.commit("appendBuLayerFilter__List", response.List);
                    } else {
                        context.commit("setBuLayerFilter__List", response.List);
                    }

                    return response;
                })
                .catch(function(err) {
                    console.error(err);
                    throw(err);
                });
        },
        loadBuLayerFilter(context, {id, filter, header}) {

            return api.find(readUrl + id, filter, header)
                .then(function(response) {

                    context.commit("setBuLayerFilter", response.Model);
                    return response;
                })
                .catch(function(err) {
                    console.error(err);
                    throw(err);
                });
        },
        updateBuLayerFilter(context, {id, data, filter, header}) {

            let url = updateUrl + id;
            if (Array.isArray && Array.isArray(data)) {
                url = multiUpdateUrl
            }

            return api.update(url, data, filter, header)
                .then(function(response) {

                    context.commit("setBuLayerFilter", response.Model);
                    return response;
                })
                .catch(function(err) {
                    console.error(err);
                    throw(err);
                });
        },
        findOrCreateBuLayerFilter(context, {id, data, filter, header}) {

            return api.update(findOrCreateUrl, data, filter, header)
                .then(function(response) {

                    context.commit("setBuLayerFilter", response.Model);
                    return response;
                })
                .catch(function(err) {
                    console.error(err);
                    throw(err);
                });
        },
        clearListBuLayerFilter(context) {
            context.commit("clearListBuLayerFilter");
        },
        clearBuLayerFilter(context) {
            context.commit("clearBuLayerFilter");
        },
    },
    getters: {
        getBuLayerFilter: (state) => {
            return state.BuLayerFilter;
        },
        getBuLayerFilterById: state => id => {
            return state.BuLayerFilter__List.find(item => item.Id === id);
        },
        getListBuLayerFilter: (state) => {
            return state.BuLayerFilter__List;
        },
        getRoute__BuLayerFilter: state => action => {
            return state.BuLayerFilter__Routes[action];
        },
        getRoutes__BuLayerFilter: state => {
            return state.BuLayerFilter__Routes;
        },
    },
    mutations: {
        setBuLayerFilter(state, data) {
            state.BuLayerFilter = data;
        },
        setBuLayerFilter__List(state, data) {
            state.BuLayerFilter__List = data || [];
        },
        appendBuLayerFilter__List(state, data) {

            if (! state.BuLayerFilter__List) {
                state.BuLayerFilter__List = [];
            }

            state.BuLayerFilter__List = state.BuLayerFilter__List.concat(data);
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
        BuLayerFilter__Routes: {
            find: findUrl,
            read: readUrl,
            create: createUrl,
            multiCreate: multiCreateUrl,
            update: updateUrl,
            multiUpdate: multiUpdateUrl,
            delete: deleteUrl,
            multiDelete: multiDeleteUrl,
            findOrCreate: findOrCreateUrl,
        },
    },
};

export default buLayerFilter;
