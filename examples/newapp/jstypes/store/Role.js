
import {Role} from "../apiModel";
import api from "../api";
import {findItemIndex} from "../common";

let findUrl = "/api/v1/role";
let readUrl = "/api/v1/role/"; // + id
let createUrl = "/api/v1/role";
let multiCreateUrl = "/api/v1/role/list";
let updateUrl = "/api/v1/role/"; // + id
let multiUpdateUrl = "/api/v1/role/list"; // + id
let deleteUrl = "/api/v1/role/"; // + id
let multiDeleteUrl = "/api/v1/role/list"; // + id
let findOrCreateUrl = "/api/v1/role"; // + id

const role = {
    actions: {
        createRole(context, {data, filter, header}) {

            let url = createUrl;
            if (Array.isArray && Array.isArray(data)) {
                url = multiCreateUrl
            }

            return api.create(url, data, filter, header)
                .then(function(response) {

                    context.commit("setRole", response.Model);

                    return response;
                })
                .catch(function(err) {
                    console.error(err);
                    throw(err);
                });
        },
        deleteRole(context, {id, header}) {

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
                    context.commit("clearRole");
                    return response;
                })
                .catch(function(err) {
                    console.error(err);
                    throw(err);
                });
        },
        findRole(context, {filter, header}) {

            return api.find(findUrl, filter, header)
                .then(function(response) {

                    context.commit("setRole__List", response.List);
                    return response;
                })
                .catch(function(err) {
                    console.error(err);
                    throw(err);
                });
        },
        loadRole(context, {id, filter, header}) {

            return api.find(readUrl + id, filter, header)
                .then(function(response) {

                    context.commit("setRole", response.Model);
                    return response;
                })
                .catch(function(err) {
                    console.error(err);
                    throw(err);
                });
        },
        updateRole(context, {id, data, filter, header}) {

            let url = updateUrl + id;
            if (Array.isArray && Array.isArray(data)) {
                url = multiUpdateUrl
            }

            return api.update(url, data, filter, header)
                .then(function(response) {

                    context.commit("setRole", response.Model);
                    return response;
                })
                .catch(function(err) {
                    console.error(err);
                    throw(err);
                });
        },
        findOrCreateRole(context, {id, data, filter, header}) {

            return api.update(findOrCreateUrl, data, filter, header)
                .then(function(response) {

                    context.commit("setRole", response.Model);
                    return response;
                })
                .catch(function(err) {
                    console.error(err);
                    throw(err);
                });
        },
        clearListRole(context) {
            context.commit("clearListRole");
        },
        clearRole(context) {
            context.commit("clearRole");
        },
    },
    getters: {
        getRole: (state) => {
            return state.Role;
        },
        getRoleById: state => id => {
            return state.Role__List.find(item => item.Id === id);
        },
        getListRole: (state) => {
            return state.Role__List;
        },
    },
    mutations: {
        setRole(state, data) {
            state.Role = data;
        },
        setRole__List(state, data) {
            state.Role__List = data || [];
        },
        clearRole(state) {
            state.Role = new Role();
        },
        clearListRole(state) {
            state.Role__List = [];
        },
		updateRoleById(state, data) {
    		let index = findItemIndex(state.Role__List, function(item) {
	        	return item.Id === data.Id;
	    	});
	    
	    	if (index || index === 0) {
		        state.Role__List.splice(index, 1, data);
    		}
		},
		deleteRoleFromList(state, id) {
		    let index = findItemIndex(state.Role__List, function(item) {
		        return item.Id === id;
		    });
		    
		    if (index || index === 0) {
		        state.Role__List.splice(index, 1);
		    }
		},
		addRoleItemToList(state, item) {

			if (state.Role__List === null) {
				state.Role__List = [];
			}

		    state.Role__List.push(item);
		},
    },
    state: {
        Role: new Role(),
        Role__List: [],
    },
};

export default role;
