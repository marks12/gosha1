
import {User} from "../apiModel";
import api from "../api";
import {findItemIndex} from "../common";

let findUrl = "/api/v1/user";
let readUrl = "/api/v1/user/"; // + id
let createUrl = "/api/v1/user";
let updateUrl = "/api/v1/user/"; // + id
let deleteUrl = "/api/v1/user/"; // + id
let findOrCreateUrl = "/api/v1/user"; // + id

const user = {
    actions: {
        createUser(context, {data, filter, header}) {

            return api.create(createUrl, data, filter, header)
                .then(function(response) {

                    context.commit("setUser", response.Model);

                    return response;
                })
                .catch(function(err) {
                    console.error(err);
                    throw(err);
                });
        },
        deleteUser(context, {id, header}) {

            return api.remove(deleteUrl + id, header)
                .then(function(response) {
                    context.commit("clearUser");
                    return response;
                })
                .catch(function(err) {
                    console.error(err);
                    throw(err);
                });
        },
        findUser(context, {filter, header}) {

            return api.find(findUrl, filter, header)
                .then(function(response) {

                    context.commit("setUser__List", response.List);
                    return response;
                })
                .catch(function(err) {
                    console.error(err);
                    throw(err);
                });
        },
        loadUser(context, {id, filter, header}) {

            return api.find(readUrl + id, filter, header)
                .then(function(response) {

                    context.commit("setUser", response.Model);
                    return response;
                })
                .catch(function(err) {
                    console.error(err);
                    throw(err);
                });
        },
        updateUser(context, {id, data, filter, header}) {

            return api.update(updateUrl + id, data, filter, header)
                .then(function(response) {

                    context.commit("setUser", response.Model);
                    return response;
                })
                .catch(function(err) {
                    console.error(err);
                    throw(err);
                });
        },
        findOrCreateUser(context, {id, data, filter, header}) {

            return api.update(findOrCreateUrl, data, filter, header)
                .then(function(response) {

                    context.commit("setUser", response.Model);
                    return response;
                })
                .catch(function(err) {
                    console.error(err);
                    throw(err);
                });
        },
        clearListUser(context) {
            context.commit("clearListUser");
        },
    },
    getters: {
        getUser: (state) => {
            return state.User;
        },
        getUserById: state => id => {
            return state.User__List.find(item => item.Id === id);
        },
        getListUser: (state) => {
            return state.User__List;
        },
    },
    mutations: {
        setUser(state, data) {
            state.User = data;
        },
        setUser__List(state, data) {
            state.User__List = data || [];
        },
        clearUser(state) {
            state.User = new User();
        },
        clearListUser(state) {
            state.User__List = [];
        },
		updateUserById(state, data) {
    		let index = findItemIndex(state.User__List, function(item) {
	        	return item.Id === data.Id;
	    	});
	    
	    	if (index || index === 0) {
		        state.User__List.splice(index, 1, data);
    		}
		},
		deleteUserFromList(state, id) {
		    let index = findItemIndex(state.User__List, function(item) {
		        return item.Id === id;
		    });
		    
		    if (index || index === 0) {
		        state.User__List.splice(index, 1);
		    }
		},
		addUserItemToList(state, item) {

			if (state.User__List === null) {
				state.User__List = [];
			}

		    state.User__List.push(item);
		},
    },
    state: {
        User: new User(),
        User__List: [],
    },
};

export default user;
