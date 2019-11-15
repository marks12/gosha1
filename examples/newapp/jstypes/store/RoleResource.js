
import {RoleResource} from "../apiModel";
import api from "../api";
import {findItemIndex} from "../common";

let findUrl = "/api/v1/roleResource";
let readUrl = "/api/v1/roleResource/"; // + id
let createUrl = "/api/v1/roleResource";
let updateUrl = "/api/v1/roleResource/"; // + id
let deleteUrl = "/api/v1/roleResource/"; // + id
let findOrCreateUrl = "/api/v1/roleResource"; // + id

const roleResource = {
    actions: {
        createRoleResource(context, {data, filter, header}) {

            return api.create(createUrl, data, filter, header)
                .then(function(response) {

                    context.commit("setRoleResource", response.Model);

                    return response;
                })
                .catch(function(err) {
                    console.error(err);
                    throw(err);
                });
        },
        deleteRoleResource(context, {id, header}) {

            return api.remove(deleteUrl + id, header)
                .then(function(response) {
                    context.commit("clearRoleResource");
                    return response;
                })
                .catch(function(err) {
                    console.error(err);
                    throw(err);
                });
        },
        findRoleResource(context, {filter, header}) {

            return api.find(findUrl, filter, header)
                .then(function(response) {

                    context.commit("setRoleResource__List", response.List);
                    return response;
                })
                .catch(function(err) {
                    console.error(err);
                    throw(err);
                });
        },
        loadRoleResource(context, {id, filter, header}) {

            return api.find(readUrl + id, filter, header)
                .then(function(response) {

                    context.commit("setRoleResource", response.Model);
                    return response;
                })
                .catch(function(err) {
                    console.error(err);
                    throw(err);
                });
        },
        updateRoleResource(context, {id, data, filter, header}) {

            return api.update(updateUrl + id, data, filter, header)
                .then(function(response) {

                    context.commit("setRoleResource", response.Model);
                    return response;
                })
                .catch(function(err) {
                    console.error(err);
                    throw(err);
                });
        },
        findOrCreateRoleResource(context, {id, data, filter, header}) {

            return api.update(findOrCreateUrl, data, filter, header)
                .then(function(response) {

                    context.commit("setRoleResource", response.Model);
                    return response;
                })
                .catch(function(err) {
                    console.error(err);
                    throw(err);
                });
        },
        clearListRoleResource(context) {
            context.commit("clearListRoleResource");
        },
    },
    getters: {
        getRoleResource: (state) => {
            return state.RoleResource;
        },
        getRoleResourceById: state => id => {
            return state.RoleResource__List.find(item => item.Id === id);
        },
        getListRoleResource: (state) => {
            return state.RoleResource__List;
        },
    },
    mutations: {
        setRoleResource(state, data) {
            state.RoleResource = data;
        },
        setRoleResource__List(state, data) {
            state.RoleResource__List = data || [];
        },
        clearRoleResource(state) {
            state.RoleResource = new RoleResource();
        },
        clearListRoleResource(state) {
            state.RoleResource__List = [];
        },
		updateRoleResourceById(state, data) {
    		let index = findItemIndex(state.RoleResource__List, function(item) {
	        	return item.Id === data.Id;
	    	});
	    
	    	if (index || index === 0) {
		        state.RoleResource__List.splice(index, 1, data);
    		}
		},
		deleteRoleResourceFromList(state, id) {
		    let index = findItemIndex(state.RoleResource__List, function(item) {
		        return item.Id === id;
		    });
		    
		    if (index || index === 0) {
		        state.RoleResource__List.splice(index, 1);
		    }
		},
		addRoleResourceItemToList(state, item) {

			if (state.RoleResource__List === null) {
				state.RoleResource__List = [];
			}

		    state.RoleResource__List.push(item);
		},
    },
    state: {
        RoleResource: new RoleResource(),
        RoleResource__List: [],
    },
};

export default roleResource;
