
import {Entity} from "../apiModel";
import api from "../api";
import {findItemIndex} from "../common";

let findUrl = "/api/v1/entity";
let readUrl = "/api/v1/entity/"; // + id
let createUrl = "/api/v1/entity";
let updateUrl = "/api/v1/entity/"; // + id
let deleteUrl = "/api/v1/entity/"; // + id
let findOrCreateUrl = "/api/v1/entity"; // + id

const entity = {
    actions: {
        createEntity(context, {data, filter, header}) {

            return api.create(createUrl, data, filter, header)
                .then(function(response) {

                    context.commit("setEntity", response.Model);

                    return response;
                })
                .catch(function(err) {
                    console.error(err);
                    throw(err);
                });
        },
        deleteEntity(context, {id, header}) {

            return api.remove(deleteUrl + id, header)
                .then(function(response) {
                    context.commit("clearEntity");
                    return response;
                })
                .catch(function(err) {
                    console.error(err);
                    throw(err);
                });
        },
        findEntity(context, {filter, header}) {

            return api.find(findUrl, filter, header)
                .then(function(response) {

                    context.commit("setEntity__List", response.List);
                    return response;
                })
                .catch(function(err) {
                    console.error(err);
                    throw(err);
                });
        },
        loadEntity(context, {id, filter, header}) {

            return api.find(readUrl + id, filter, header)
                .then(function(response) {

                    context.commit("setEntity", response.Model);
                    return response;
                })
                .catch(function(err) {
                    console.error(err);
                    throw(err);
                });
        },
        updateEntity(context, {id, data, filter, header}) {

            return api.update(updateUrl + id, data, filter, header)
                .then(function(response) {

                    context.commit("setEntity", response.Model);
                    return response;
                })
                .catch(function(err) {
                    console.error(err);
                    throw(err);
                });
        },
        findOrCreateEntity(context, {id, data, filter, header}) {

            return api.update(findOrCreateUrl, data, filter, header)
                .then(function(response) {

                    context.commit("setEntity", response.Model);
                    return response;
                })
                .catch(function(err) {
                    console.error(err);
                    throw(err);
                });
        },
        clearListEntity(context) {
            context.commit("clearListEntity");
        },
    },
    getters: {
        getEntity: (state) => {
            return state.Entity;
        },
        getEntityById: state => id => {
            return state.Entity__List.find(item => item.Id === id);
        },
        getListEntity: (state) => {
            return state.Entity__List;
        },
    },
    mutations: {
        setEntity(state, data) {
            state.Entity = data;
        },
        setEntity__List(state, data) {
            state.Entity__List = data || [];
        },
        clearEntity(state) {
            state.Entity = new Entity();
        },
        clearListEntity(state) {
            state.Entity__List = [];
        },
		updateEntityById(state, data) {
    		let index = findItemIndex(state.Entity__List, function(item) {
	        	return item.Id === data.Id;
	    	});
	    
	    	if (index || index === 0) {
		        state.Entity__List.splice(index, 1, data);
    		}
		},
		deleteEntityFromList(state, id) {
		    let index = findItemIndex(state.Entity__List, function(item) {
		        return item.Id === id;
		    });
		    
		    if (index || index === 0) {
		        state.Entity__List.splice(index, 1);
		    }
		},
		addEntityItemToList(state, item) {

			if (state.Entity__List === null) {
				state.Entity__List = [];
			}

		    state.Entity__List.push(item);
		},
    },
    state: {
        Entity: new Entity(),
        Entity__List: [],
    },
};

export default entity;
