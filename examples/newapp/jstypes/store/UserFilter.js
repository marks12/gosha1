
import {UserFilter} from "../apiModel";
import api from "../api";
import {findItemIndex} from "../common";

let findUrl = "/api/v1/userFilter";
let readUrl = "/api/v1/userFilter/"; // + id
let createUrl = "/api/v1/userFilter";
let updateUrl = "/api/v1/userFilter/"; // + id
let deleteUrl = "/api/v1/userFilter/"; // + id
let findOrCreateUrl = "/api/v1/userFilter"; // + id

const userFilter = {
    actions: {
        createUserFilter(context, {data, filter, header}) {

            return api.create(createUrl, data, filter, header)
                .then(function(response) {

                    context.commit("setUserFilter", response.Model);

                    return response;
                })
                .catch(function(err) {
                    console.error(err);
                    throw(err);
                });
        },
        deleteUserFilter(context, {id, header}) {

            return api.remove(deleteUrl + id, header)
                .then(function(response) {
                    context.commit("clearUserFilter");
                    return response;
                })
                .catch(function(err) {
                    console.error(err);
                    throw(err);
                });
        },
        findUserFilter(context, {filter, header}) {

            return api.find(findUrl, filter, header)
                .then(function(response) {

                    context.commit("setUserFilter__List", response.List);
                    return response;
                })
                .catch(function(err) {
                    console.error(err);
                    throw(err);
                });
        },
        loadUserFilter(context, {id, filter, header}) {

            return api.find(readUrl + id, filter, header)
                .then(function(response) {

                    context.commit("setUserFilter", response.Model);
                    return response;
                })
                .catch(function(err) {
                    console.error(err);
                    throw(err);
                });
        },
        updateUserFilter(context, {id, data, filter, header}) {

            return api.update(updateUrl + id, data, filter, header)
                .then(function(response) {

                    context.commit("setUserFilter", response.Model);
                    return response;
                })
                .catch(function(err) {
                    console.error(err);
                    throw(err);
                });
        },
        findOrCreateUserFilter(context, {id, data, filter, header}) {

            return api.update(findOrCreateUrl, data, filter, header)
                .then(function(response) {

                    context.commit("setUserFilter", response.Model);
                    return response;
                })
                .catch(function(err) {
                    console.error(err);
                    throw(err);
                });
        },
        clearListUserFilter(context) {
            context.commit("clearListUserFilter");
        },
    },
    getters: {
        getUserFilter: (state) => {
            return state.UserFilter;
        },
        getUserFilterById: state => id => {
            return state.UserFilter__List.find(item => item.Id === id);
        },
        getListUserFilter: (state) => {
            return state.UserFilter__List;
        },
    },
    mutations: {
        setUserFilter(state, data) {
            state.UserFilter = data;
        },
        setUserFilter__List(state, data) {
            state.UserFilter__List = data || [];
        },
        clearUserFilter(state) {
            state.UserFilter = new UserFilter();
        },
        clearListUserFilter(state) {
            state.UserFilter__List = [];
        },
		updateUserFilterById(state, data) {
    		let index = findItemIndex(state.UserFilter__List, function(item) {
	        	return item.Id === data.Id;
	    	});
	    
	    	if (index || index === 0) {
		        state.UserFilter__List.splice(index, 1, data);
    		}
		},
		deleteUserFilterFromList(state, id) {
		    let index = findItemIndex(state.UserFilter__List, function(item) {
		        return item.Id === id;
		    });
		    
		    if (index || index === 0) {
		        state.UserFilter__List.splice(index, 1);
		    }
		},
		addUserFilterItemToList(state, item) {

			if (state.UserFilter__List === null) {
				state.UserFilter__List = [];
			}

		    state.UserFilter__List.push(item);
		},
    },
    state: {
        UserFilter: new UserFilter(),
        UserFilter__List: [],
    },
};

export default userFilter;
