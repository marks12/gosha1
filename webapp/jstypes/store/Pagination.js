
import {Pagination} from "../apiModel";
import api from "../api";
import {findItemIndex} from "../common";

let findUrl = "/api/v1/pagination";
let readUrl = "/api/v1/pagination/"; // + id
let createUrl = "/api/v1/pagination";
let updateUrl = "/api/v1/pagination/"; // + id
let deleteUrl = "/api/v1/pagination/"; // + id
let findOrCreateUrl = "/api/v1/pagination"; // + id

const pagination = {
    actions: {
        createPagination(context, {data, filter, header}) {

            return api.create(createUrl, data, filter, header)
                .then(function(response) {

                    context.commit("setPagination", response.Model);

                    return response;
                })
                .catch(function(err) {
                    return err;
                });
        },
        deletePagination(context, {id, header}) {

            return api.remove(deleteUrl + id, header)
                .then(function(response) {
                    context.commit("clearPagination");
                    return response;
                })
                .catch(function(err) {
                    return err;
                });
        },
        findPagination(context, {filter, header}) {

            return api.find(findUrl, filter, header)
                .then(function(response) {

                    context.commit("setPagination__List", response.List);

                    return response;
                })
                .catch(function(err) {
                    return err;
                });
        },
        loadPagination(context, {id, filter, header}) {

            return api.find(readUrl + id, filter, header)
                .then(function(response) {

                    context.commit("setPagination", response.Model);

                    return response;
                })
                .catch(function(err) {
                    return err;
                });
        },
        updatePagination(context, {id, data, filter, header}) {

            return api.update(updateUrl + id, data, filter, header)
                .then(function(response) {

                    context.commit("setPagination", response.Model);

                    return response;
                })
                .catch(function(err) {
                    return err;
                });
        },
        findOrCreatePagination(context, {id, data, filter, header}) {

            return api.update(findOrCreateUrl, data, filter, header)
                .then(function(response) {

                    context.commit("setPagination", response.Model);

                    return response;
                })
                .catch(function(err) {
                    return err;
                });
        },
        clearListPagination(context) {
            context.commit("clearListPagination");
        },
    },
    getters: {
        getPagination: (state) => {
            return state.Pagination;
        },
        getListPagination: (state) => {
            return state.Pagination__List;
        },
    },
    mutations: {
        setPagination(state, data) {
            state.Pagination = data;
        },
        setPagination__List(state, data) {
            state.Pagination__List = data || [];
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
    },
};

export default pagination;
