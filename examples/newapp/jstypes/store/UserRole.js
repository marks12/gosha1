
import {UserRole} from "../apiModel";
import api from "../api";
import {findItemIndex} from "../common";

let findUrl = "/api/v1/userRole";
let readUrl = "/api/v1/userRole/"; // + id
let createUrl = "/api/v1/userRole";
let updateUrl = "/api/v1/userRole/"; // + id
let deleteUrl = "/api/v1/userRole/"; // + id
let findOrCreateUrl = "/api/v1/userRole"; // + id

const userRole = {
    actions: {
        createUserRole(context, {data, filter, header}) {

            return api.create(createUrl, data, filter, header)
                .then(function(response) {

                    context.commit("setUserRole", response.Model);

                    return response;
                })
                .catch(function(err) {
                    console.error(err);
                    throw(err);
                });
        },
        deleteUserRole(context, {id, header}) {

            return api.remove(deleteUrl + id, header)
                .then(function(response) {
                    context.commit("clearUserRole");
                    return response;
                })
                .catch(function(err) {
                    console.error(err);
                    throw(err);
                });
        },
        findUserRole(context, {filter, header}) {

            return api.find(findUrl, filter, header)
                .then(function(response) {

                    context.commit("setUserRole__List", response.List);
                    return response;
                })
                .catch(function(err) {
                    console.error(err);
                    throw(err);
                });
        },
        loadUserRole(context, {id, filter, header}) {

            return api.find(readUrl + id, filter, header)
                .then(function(response) {

                    context.commit("setUserRole", response.Model);
                    return response;
                })
                .catch(function(err) {
                    console.error(err);
                    throw(err);
                });
        },
        updateUserRole(context, {id, data, filter, header}) {

            return api.update(updateUrl + id, data, filter, header)
                .then(function(response) {

                    context.commit("setUserRole", response.Model);
                    return response;
                })
                .catch(function(err) {
                    console.error(err);
                    throw(err);
                });
        },
        findOrCreateUserRole(context, {id, data, filter, header}) {

            return api.update(findOrCreateUrl, data, filter, header)
                .then(function(response) {

                    context.commit("setUserRole", response.Model);
                    return response;
                })
                .catch(function(err) {
                    console.error(err);
                    throw(err);
                });
        },
        clearListUserRole(context) {
            context.commit("clearListUserRole");
        },
    },
    getters: {
        getUserRole: (state) => {
            return state.UserRole;
        },
        getUserRoleById: state => id => {
            return state.UserRole__List.find(item => item.Id === id);
        },
        getListUserRole: (state) => {
            return state.UserRole__List;
        },
    },
    mutations: {
        setUserRole(state, data) {
            state.UserRole = data;
        },
        setUserRole__List(state, data) {
            state.UserRole__List = data || [];
        },
        clearUserRole(state) {
            state.UserRole = new UserRole();
        },
        clearListUserRole(state) {
            state.UserRole__List = [];
        },
		updateUserRoleById(state, data) {
    		let index = findItemIndex(state.UserRole__List, function(item) {
	        	return item.Id === data.Id;
	    	});
	    
	    	if (index || index === 0) {
		        state.UserRole__List.splice(index, 1, data);
    		}
		},
		deleteUserRoleFromList(state, id) {
		    let index = findItemIndex(state.UserRole__List, function(item) {
		        return item.Id === id;
		    });
		    
		    if (index || index === 0) {
		        state.UserRole__List.splice(index, 1);
		    }
		},
		addUserRoleItemToList(state, item) {

			if (state.UserRole__List === null) {
				state.UserRole__List = [];
			}

		    state.UserRole__List.push(item);
		},
    },
    state: {
        UserRole: new UserRole(),
        UserRole__List: [],
    },
};

export default userRole;
