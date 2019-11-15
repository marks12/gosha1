
import {CurrentUserFilter} from "../apiModel";
import api from "../api";
import {findItemIndex} from "../common";

let findUrl = "/api/v1/currentUserFilter";
let readUrl = "/api/v1/currentUserFilter/"; // + id
let createUrl = "/api/v1/currentUserFilter";
let updateUrl = "/api/v1/currentUserFilter/"; // + id
let deleteUrl = "/api/v1/currentUserFilter/"; // + id
let findOrCreateUrl = "/api/v1/currentUserFilter"; // + id

const currentUserFilter = {
    actions: {
        createCurrentUserFilter(context, {data, filter, header}) {

            return api.create(createUrl, data, filter, header)
                .then(function(response) {

                    context.commit("setCurrentUserFilter", response.Model);

                    return response;
                })
                .catch(function(err) {
                    console.error(err);
                    throw(err);
                });
        },
        deleteCurrentUserFilter(context, {id, header}) {

            return api.remove(deleteUrl + id, header)
                .then(function(response) {
                    context.commit("clearCurrentUserFilter");
                    return response;
                })
                .catch(function(err) {
                    console.error(err);
                    throw(err);
                });
        },
        findCurrentUserFilter(context, {filter, header}) {

            return api.find(findUrl, filter, header)
                .then(function(response) {

                    context.commit("setCurrentUserFilter__List", response.List);
                    return response;
                })
                .catch(function(err) {
                    console.error(err);
                    throw(err);
                });
        },
        loadCurrentUserFilter(context, {id, filter, header}) {

            return api.find(readUrl + id, filter, header)
                .then(function(response) {

                    context.commit("setCurrentUserFilter", response.Model);
                    return response;
                })
                .catch(function(err) {
                    console.error(err);
                    throw(err);
                });
        },
        updateCurrentUserFilter(context, {id, data, filter, header}) {

            return api.update(updateUrl + id, data, filter, header)
                .then(function(response) {

                    context.commit("setCurrentUserFilter", response.Model);
                    return response;
                })
                .catch(function(err) {
                    console.error(err);
                    throw(err);
                });
        },
        findOrCreateCurrentUserFilter(context, {id, data, filter, header}) {

            return api.update(findOrCreateUrl, data, filter, header)
                .then(function(response) {

                    context.commit("setCurrentUserFilter", response.Model);
                    return response;
                })
                .catch(function(err) {
                    console.error(err);
                    throw(err);
                });
        },
        clearListCurrentUserFilter(context) {
            context.commit("clearListCurrentUserFilter");
        },
    },
    getters: {
        getCurrentUserFilter: (state) => {
            return state.CurrentUserFilter;
        },
        getCurrentUserFilterById: state => id => {
            return state.CurrentUserFilter__List.find(item => item.Id === id);
        },
        getListCurrentUserFilter: (state) => {
            return state.CurrentUserFilter__List;
        },
    },
    mutations: {
        setCurrentUserFilter(state, data) {
            state.CurrentUserFilter = data;
        },
        setCurrentUserFilter__List(state, data) {
            state.CurrentUserFilter__List = data || [];
        },
        clearCurrentUserFilter(state) {
            state.CurrentUserFilter = new CurrentUserFilter();
        },
        clearListCurrentUserFilter(state) {
            state.CurrentUserFilter__List = [];
        },
		updateCurrentUserFilterById(state, data) {
    		let index = findItemIndex(state.CurrentUserFilter__List, function(item) {
	        	return item.Id === data.Id;
	    	});
	    
	    	if (index || index === 0) {
		        state.CurrentUserFilter__List.splice(index, 1, data);
    		}
		},
		deleteCurrentUserFilterFromList(state, id) {
		    let index = findItemIndex(state.CurrentUserFilter__List, function(item) {
		        return item.Id === id;
		    });
		    
		    if (index || index === 0) {
		        state.CurrentUserFilter__List.splice(index, 1);
		    }
		},
		addCurrentUserFilterItemToList(state, item) {

			if (state.CurrentUserFilter__List === null) {
				state.CurrentUserFilter__List = [];
			}

		    state.CurrentUserFilter__List.push(item);
		},
    },
    state: {
        CurrentUserFilter: new CurrentUserFilter(),
        CurrentUserFilter__List: [],
    },
};

export default currentUserFilter;
