
import {CurrentUser} from "../apiModel";
import api from "../api";
import {findItemIndex} from "../common";

let findUrl = "/api/v1/currentUser";
let readUrl = "/api/v1/currentUser/"; // + id
let createUrl = "/api/v1/currentUser";
let updateUrl = "/api/v1/currentUser/"; // + id
let deleteUrl = "/api/v1/currentUser/"; // + id
let findOrCreateUrl = "/api/v1/currentUser"; // + id

const currentUser = {
    actions: {
        createCurrentUser(context, {data, filter, header}) {

            return api.create(createUrl, data, filter, header)
                .then(function(response) {

                    context.commit("setCurrentUser", response.Model);

                    return response;
                })
                .catch(function(err) {
                    console.error(err);
                    throw(err);
                });
        },
        deleteCurrentUser(context, {id, header}) {

            return api.remove(deleteUrl + id, header)
                .then(function(response) {
                    context.commit("clearCurrentUser");
                    return response;
                })
                .catch(function(err) {
                    console.error(err);
                    throw(err);
                });
        },
        findCurrentUser(context, {filter, header}) {

            return api.find(findUrl, filter, header)
                .then(function(response) {

                    context.commit("setCurrentUser__List", response.List);
                    return response;
                })
                .catch(function(err) {
                    console.error(err);
                    throw(err);
                });
        },
        loadCurrentUser(context, {id, filter, header}) {

            return api.find(readUrl + id, filter, header)
                .then(function(response) {

                    context.commit("setCurrentUser", response.Model);
                    return response;
                })
                .catch(function(err) {
                    console.error(err);
                    throw(err);
                });
        },
        updateCurrentUser(context, {id, data, filter, header}) {

            return api.update(updateUrl + id, data, filter, header)
                .then(function(response) {

                    context.commit("setCurrentUser", response.Model);
                    return response;
                })
                .catch(function(err) {
                    console.error(err);
                    throw(err);
                });
        },
        findOrCreateCurrentUser(context, {id, data, filter, header}) {

            return api.update(findOrCreateUrl, data, filter, header)
                .then(function(response) {

                    context.commit("setCurrentUser", response.Model);
                    return response;
                })
                .catch(function(err) {
                    console.error(err);
                    throw(err);
                });
        },
        clearListCurrentUser(context) {
            context.commit("clearListCurrentUser");
        },
    },
    getters: {
        getCurrentUser: (state) => {
            return state.CurrentUser;
        },
        getCurrentUserById: state => id => {
            return state.CurrentUser__List.find(item => item.Id === id);
        },
        getListCurrentUser: (state) => {
            return state.CurrentUser__List;
        },
    },
    mutations: {
        setCurrentUser(state, data) {
            state.CurrentUser = data;
        },
        setCurrentUser__List(state, data) {
            state.CurrentUser__List = data || [];
        },
        clearCurrentUser(state) {
            state.CurrentUser = new CurrentUser();
        },
        clearListCurrentUser(state) {
            state.CurrentUser__List = [];
        },
		updateCurrentUserById(state, data) {
    		let index = findItemIndex(state.CurrentUser__List, function(item) {
	        	return item.Id === data.Id;
	    	});
	    
	    	if (index || index === 0) {
		        state.CurrentUser__List.splice(index, 1, data);
    		}
		},
		deleteCurrentUserFromList(state, id) {
		    let index = findItemIndex(state.CurrentUser__List, function(item) {
		        return item.Id === id;
		    });
		    
		    if (index || index === 0) {
		        state.CurrentUser__List.splice(index, 1);
		    }
		},
		addCurrentUserItemToList(state, item) {

			if (state.CurrentUser__List === null) {
				state.CurrentUser__List = [];
			}

		    state.CurrentUser__List.push(item);
		},
    },
    state: {
        CurrentUser: new CurrentUser(),
        CurrentUser__List: [],
    },
};

export default currentUser;
