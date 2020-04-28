
import {Pagination} from "../apiModel";
import api from "../api";
import {findItemIndex} from "../common";

let findUrl = "/api/v1/pagination";
let readUrl = "/api/v1/pagination/"; // + id
let createUrl = "/api/v1/pagination";
let multiCreateUrl = "/api/v1/pagination/list";
let updateUrl = "/api/v1/pagination/"; // + id
let multiUpdateUrl = "/api/v1/pagination/list";
let deleteUrl = "/api/v1/pagination/"; // + id
let multiDeleteUrl = "/api/v1/pagination/list";
let findOrCreateUrl = "/api/v1/pagination";
let updateOrCreateUrl = "/api/v1/pagination";

const pagination = {
    actions: {
        createPagination(context, {data, filter, header, noMutation}) {

            let url = createUrl;
            if (Array.isArray && Array.isArray(data)) {
                url = multiCreateUrl
            }

            return api.create(url, data, filter, header)
                .then(function(response) {

					if(! noMutation) {
	                    context.commit("setPagination", response.Model);
					}

                    return response;
                })
                .catch(function(err) {
                    console.error(err);
                    throw(err);
                });
        },
        deletePagination(context, {id, header, noMutation}) {

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
	                    context.commit("clearPagination");
					}
                    return response;
                })
                .catch(function(err) {
                    console.error(err);
                    throw(err);
                });
        },
        findPagination(context, {filter, header, isAppend, noMutation}) {

            return api.find(findUrl, filter, header)
                .then(function(response) {

					if(! noMutation) {
						if (isAppend) {
							context.commit("appendPagination__List", response.List);
						} else {
							context.commit("setPagination__List", response.List);
						}
					}

                    return response;
                })
                .catch(function(err) {
                    console.error(err);
                    throw(err);
                });
        },
        loadPagination(context, {id, filter, header, noMutation}) {

            return api.find(readUrl + id, filter, header)
                .then(function(response) {

					if(! noMutation) {
	                    context.commit("setPagination", response.Model);
					}
                    return response;
                })
                .catch(function(err) {
                    console.error(err);
                    throw(err);
                });
        },
        updatePagination(context, {id, data, filter, header, noMutation}) {

            let url = updateUrl + id;
            if (Array.isArray && Array.isArray(data)) {
                url = multiUpdateUrl
            }

            return api.update(url, data, filter, header)
                .then(function(response) {
					if(! noMutation) {
	                    context.commit("setPagination", response.Model);
					}
                    return response;
                })
                .catch(function(err) {
                    console.error(err);
                    throw(err);
                });
        },
        findOrCreatePagination(context, {id, data, filter, header, noMutation}) {

            return api.update(findOrCreateUrl, data, filter, header)
                .then(function(response) {

					if(! noMutation) {
	                    context.commit("setPagination", response.Model);
					}
                    return response;
                })
                .catch(function(err) {
                    console.error(err);
                    throw(err);
                });
        },
        clearListPagination(context) {
            context.commit("clearListPagination");
        },
        clearPagination(context) {
            context.commit("clearPagination");
        },
    },
    getters: {
        getPagination: (state) => {
            return state.Pagination;
        },
        getPaginationById: state => id => {
            return state.Pagination__List.find(item => item.Id === id);
        },
        getListPagination: (state) => {
            return state.Pagination__List;
        },
        getRoute__Pagination: state => action => {
            return state.Pagination__Routes[action];
        },
        getRoutes__Pagination: state => {
            return state.Pagination__Routes;
        },
    },
    mutations: {
        setPagination(state, data) {
            state.Pagination = data;
        },
        setPagination__List(state, data) {
            state.Pagination__List = data || [];
        },
        appendPagination__List(state, data) {

            if (! state.Pagination__List) {
                state.Pagination__List = [];
            }

            state.Pagination__List = state.Pagination__List.concat(data);
        },
        clearPagination(state) {
            state.Pagination = new Pagination();
        },
        clearListPagination(state) {
            state.Pagination__List = [];
        },
		updatePaginationById(state, data) {
    		let index = findItemIndex(state.Pagination__List, function(item) {
	        	return item.Id === data.Id;
	    	});
	    
	    	if (index || index === 0) {
		        state.Pagination__List.splice(index, 1, data);
    		}
		},
		deletePaginationFromList(state, id) {
		    let index = findItemIndex(state.Pagination__List, function(item) {
		        return item.Id === id;
		    });
		    
		    if (index || index === 0) {
		        state.Pagination__List.splice(index, 1);
		    }
		},
		addPaginationItemToList(state, item) {

			if (state.Pagination__List === null) {
				state.Pagination__List = [];
			}

		    state.Pagination__List.push(item);
		},
    },
    state: {
        Pagination: new Pagination(),
        Pagination__List: [],
        Pagination__Routes: {
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

export default pagination;
