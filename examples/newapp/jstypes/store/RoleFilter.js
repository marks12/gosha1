
import {RoleFilter} from "../apiModel";
import api from "../api";
import {findItemIndex} from "../common";

let findUrl = "/api/v1/roleFilter";
let readUrl = "/api/v1/roleFilter/"; // + id
let createUrl = "/api/v1/roleFilter";
let multiCreateUrl = "/api/v1/roleFilter/list";
let updateUrl = "/api/v1/roleFilter/"; // + id
let multiUpdateUrl = "/api/v1/roleFilter/list"; // + id
let deleteUrl = "/api/v1/roleFilter/"; // + id
let multiDeleteUrl = "/api/v1/roleFilter/list"; // + id
let findOrCreateUrl = "/api/v1/roleFilter"; // + id

const roleFilter = {
    actions: {
        createRoleFilter(context, {data, filter, header}) {

            let url = createUrl;
            if (Array.isArray && Array.isArray(data)) {
                url = multiCreateUrl
            }

            return api.create(url, data, filter, header)
                .then(function(response) {

                    context.commit("setRoleFilter", response.Model);

                    return response;
                })
                .catch(function(err) {
                    console.error(err);
                    throw(err);
                });
        },
        deleteRoleFilter(context, {id, header}) {

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
                    context.commit("clearRoleFilter");
                    return response;
                })
                .catch(function(err) {
                    console.error(err);
                    throw(err);
                });
        },
        findRoleFilter(context, {filter, header, isAppend}) {

            return api.find(findUrl, filter, header)
                .then(function(response) {

                    if (isAppend) {
                        context.commit("appendRoleFilter__List", response.List);
                    } else {
                        context.commit("setRoleFilter__List", response.List);
                    }

                    return response;
                })
                .catch(function(err) {
                    console.error(err);
                    throw(err);
                });
        },
        loadRoleFilter(context, {id, filter, header}) {

            return api.find(readUrl + id, filter, header)
                .then(function(response) {

                    context.commit("setRoleFilter", response.Model);
                    return response;
                })
                .catch(function(err) {
                    console.error(err);
                    throw(err);
                });
        },
        updateRoleFilter(context, {id, data, filter, header}) {

            let url = updateUrl + id;
            if (Array.isArray && Array.isArray(data)) {
                url = multiUpdateUrl
            }

            return api.update(url, data, filter, header)
                .then(function(response) {

                    context.commit("setRoleFilter", response.Model);
                    return response;
                })
                .catch(function(err) {
                    console.error(err);
                    throw(err);
                });
        },
        findOrCreateRoleFilter(context, {id, data, filter, header}) {

            return api.update(findOrCreateUrl, data, filter, header)
                .then(function(response) {

                    context.commit("setRoleFilter", response.Model);
                    return response;
                })
                .catch(function(err) {
                    console.error(err);
                    throw(err);
                });
        },
        clearListRoleFilter(context) {
            context.commit("clearListRoleFilter");
        },
        clearRoleFilter(context) {
            context.commit("clearRoleFilter");
        },
    },
    getters: {
        getRoleFilter: (state) => {
            return state.RoleFilter;
        },
        getRoleFilterById: state => id => {
            return state.RoleFilter__List.find(item => item.Id === id);
        },
        getListRoleFilter: (state) => {
            return state.RoleFilter__List;
        },
    },
    mutations: {
        setRoleFilter(state, data) {
            state.RoleFilter = data;
        },
        setRoleFilter__List(state, data) {
            state.RoleFilter__List = data || [];
        },
        appendRoleFilter__List(state, data) {

            if (! state.RoleFilter__List) {
                state.RoleFilter__List = [];
            }

            state.RoleFilter__List = state.RoleFilter__List.concat(data);
        },
        clearRoleFilter(state) {
            state.RoleFilter = new RoleFilter();
        },
        clearListRoleFilter(state) {
            state.RoleFilter__List = [];
        },
		updateRoleFilterById(state, data) {
    		let index = findItemIndex(state.RoleFilter__List, function(item) {
	        	return item.Id === data.Id;
	    	});
	    
	    	if (index || index === 0) {
		        state.RoleFilter__List.splice(index, 1, data);
    		}
		},
		deleteRoleFilterFromList(state, id) {
		    let index = findItemIndex(state.RoleFilter__List, function(item) {
		        return item.Id === id;
		    });
		    
		    if (index || index === 0) {
		        state.RoleFilter__List.splice(index, 1);
		    }
		},
		addRoleFilterItemToList(state, item) {

			if (state.RoleFilter__List === null) {
				state.RoleFilter__List = [];
			}

		    state.RoleFilter__List.push(item);
		},
    },
    state: {
        RoleFilter: new RoleFilter(),
        RoleFilter__List: [],
    },
};

export default roleFilter;
