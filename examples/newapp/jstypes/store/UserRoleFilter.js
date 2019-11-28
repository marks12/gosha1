
import {UserRoleFilter} from "../apiModel";
import api from "../api";
import {findItemIndex} from "../common";

let findUrl = "/api/v1/userRoleFilter";
let readUrl = "/api/v1/userRoleFilter/"; // + id
let createUrl = "/api/v1/userRoleFilter";
let multiCreateUrl = "/api/v1/userRoleFilter/list";
let updateUrl = "/api/v1/userRoleFilter/"; // + id
let multiUpdateUrl = "/api/v1/userRoleFilter/list"; // + id
let deleteUrl = "/api/v1/userRoleFilter/"; // + id
let multiDeleteUrl = "/api/v1/userRoleFilter/list"; // + id
let findOrCreateUrl = "/api/v1/userRoleFilter"; // + id

const userRoleFilter = {
    actions: {
        createUserRoleFilter(context, {data, filter, header}) {

            let url = createUrl;
            if (Array.isArray && Array.isArray(data)) {
                url = multiCreateUrl
            }

            return api.create(url, data, filter, header)
                .then(function(response) {

                    context.commit("setUserRoleFilter", response.Model);

                    return response;
                })
                .catch(function(err) {
                    console.error(err);
                    throw(err);
                });
        },
        deleteUserRoleFilter(context, {id, header}) {

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
                    context.commit("clearUserRoleFilter");
                    return response;
                })
                .catch(function(err) {
                    console.error(err);
                    throw(err);
                });
        },
        findUserRoleFilter(context, {filter, header}) {

            return api.find(findUrl, filter, header)
                .then(function(response) {

                    context.commit("setUserRoleFilter__List", response.List);
                    return response;
                })
                .catch(function(err) {
                    console.error(err);
                    throw(err);
                });
        },
        loadUserRoleFilter(context, {id, filter, header}) {

            return api.find(readUrl + id, filter, header)
                .then(function(response) {

                    context.commit("setUserRoleFilter", response.Model);
                    return response;
                })
                .catch(function(err) {
                    console.error(err);
                    throw(err);
                });
        },
        updateUserRoleFilter(context, {id, data, filter, header}) {

            let url = updateUrl + id;
            if (Array.isArray && Array.isArray(data)) {
                url = multiUpdateUrl
            }

            return api.update(url, data, filter, header)
                .then(function(response) {

                    context.commit("setUserRoleFilter", response.Model);
                    return response;
                })
                .catch(function(err) {
                    console.error(err);
                    throw(err);
                });
        },
        findOrCreateUserRoleFilter(context, {id, data, filter, header}) {

            return api.update(findOrCreateUrl, data, filter, header)
                .then(function(response) {

                    context.commit("setUserRoleFilter", response.Model);
                    return response;
                })
                .catch(function(err) {
                    console.error(err);
                    throw(err);
                });
        },
        clearListUserRoleFilter(context) {
            context.commit("clearListUserRoleFilter");
        },
        clearUserRoleFilter(context) {
            context.commit("clearUserRoleFilter");
        },
    },
    getters: {
        getUserRoleFilter: (state) => {
            return state.UserRoleFilter;
        },
        getUserRoleFilterById: state => id => {
            return state.UserRoleFilter__List.find(item => item.Id === id);
        },
        getListUserRoleFilter: (state) => {
            return state.UserRoleFilter__List;
        },
    },
    mutations: {
        setUserRoleFilter(state, data) {
            state.UserRoleFilter = data;
        },
        setUserRoleFilter__List(state, data) {
            state.UserRoleFilter__List = data || [];
        },
        clearUserRoleFilter(state) {
            state.UserRoleFilter = new UserRoleFilter();
        },
        clearListUserRoleFilter(state) {
            state.UserRoleFilter__List = [];
        },
		updateUserRoleFilterById(state, data) {
    		let index = findItemIndex(state.UserRoleFilter__List, function(item) {
	        	return item.Id === data.Id;
	    	});
	    
	    	if (index || index === 0) {
		        state.UserRoleFilter__List.splice(index, 1, data);
    		}
		},
		deleteUserRoleFilterFromList(state, id) {
		    let index = findItemIndex(state.UserRoleFilter__List, function(item) {
		        return item.Id === id;
		    });
		    
		    if (index || index === 0) {
		        state.UserRoleFilter__List.splice(index, 1);
		    }
		},
		addUserRoleFilterItemToList(state, item) {

			if (state.UserRoleFilter__List === null) {
				state.UserRoleFilter__List = [];
			}

		    state.UserRoleFilter__List.push(item);
		},
    },
    state: {
        UserRoleFilter: new UserRoleFilter(),
        UserRoleFilter__List: [],
    },
};

export default userRoleFilter;
