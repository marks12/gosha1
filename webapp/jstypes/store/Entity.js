
import {Entity} from "../apiModel";
import api from "../api";
import {findItemIndex} from "../common";

let findUrl = "/api/v1/entity";
let readUrl = "/api/v1/entity/"; // + id
let createUrl = "/api/v1/entity";
let multiCreateUrl = "/api/v1/entity/list";
let updateUrl = "/api/v1/entity/"; // + id
let multiUpdateUrl = "/api/v1/entity/list";
let deleteUrl = "/api/v1/entity/"; // + id
let multiDeleteUrl = "/api/v1/entity/list";
let findOrCreateUrl = "/api/v1/entity";
let updateOrCreateUrl = "/api/v1/entity";

const entity = {
    actions: {
        createEntity(context, {data, filter, header, noMutation}) {

            let url = createUrl;
            if (Array.isArray && Array.isArray(data)) {
                url = multiCreateUrl
            }

            return api.create(url, data, filter, header)
                .then(function(response) {

					if(! noMutation) {
	                    context.commit("setEntity", response.Model);
					}

                    return response;
                })
                .catch(function(err) {
                    console.error(err);
                    throw(err);
                });
        },
        deleteEntity(context, {id, header, noMutation}) {

            let url;
            let dataOrNull = null;

            if (Array.isArray && Array.isArray(id)) {
                url = multiDeleteUrl;
                dataOrNull = id.map(item => typeof item === "number" ? {Id: item} : item);
            } else {
                url = deleteUrl + id;
            }

            return api.remove(url, header, dataOrNull)
                .then(function(response) {
					if(! noMutation) {
	                    context.commit("clearEntity");
					}
                    return response;
                })
                .catch(function(err) {
                    console.error(err);
                    throw(err);
                });
        },
        findEntity(context, {filter, header, isAppend, noMutation}) {

            return api.find(findUrl, filter, header)
                .then(function(response) {

					if(! noMutation) {
						if (isAppend) {
							context.commit("appendEntity__List", response.List);
						} else {
							context.commit("setEntity__List", response.List);
						}
					}

                    return response;
                })
                .catch(function(err) {
                    console.error(err);
                    throw(err);
                });
        },
        loadEntity(context, {id, filter, header, noMutation}) {

            return api.find(readUrl + id, filter, header)
                .then(function(response) {

					if(! noMutation) {
	                    context.commit("setEntity", response.Model);
					}
                    return response;
                })
                .catch(function(err) {
                    console.error(err);
                    throw(err);
                });
        },
        updateEntity(context, {id, data, filter, header, noMutation}) {

            let url = updateUrl + id;
            if (Array.isArray && Array.isArray(data)) {
                url = multiUpdateUrl
            }

            return api.update(url, data, filter, header)
                .then(function(response) {
					if(! noMutation) {
	                    context.commit("setEntity", response.Model);
					}
                    return response;
                })
                .catch(function(err) {
                    console.error(err);
                    throw(err);
                });
        },
        findOrCreateEntity(context, {id, data, filter, header, noMutation}) {

            return api.update(findOrCreateUrl, data, filter, header)
                .then(function(response) {

					if(! noMutation) {
	                    context.commit("setEntity", response.Model);
					}
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
        clearEntity(context) {
            context.commit("clearEntity");
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
        getRoute__Entity: state => action => {
            return state.Entity__Routes[action];
        },
        getRoutes__Entity: state => {
            return state.Entity__Routes;
        },
    },
    mutations: {
        setEntity(state, data) {
            state.Entity = data;
        },
        setEntity__List(state, data) {
            state.Entity__List = data || [];
        },
        appendEntity__List(state, data) {

            if (! state.Entity__List) {
                state.Entity__List = [];
            }

			if (data !== null) {
				state.Entity__List = state.Entity__List.concat(data);				
			}
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
        Entity__Routes: {
            find: findUrl,
            read: readUrl,
            create: createUrl,
            multiCreate: multiCreateUrl,
            update: updateUrl,
            multiUpdate: multiUpdateUrl,
            delete: deleteUrl,
            multiDelete: multiDeleteUrl,
            findOrCreate: findOrCreateUrl,
            updateOrCreate: updateOrCreateUrl,
        },
    },
};

export default entity;
