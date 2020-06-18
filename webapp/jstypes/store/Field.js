
import {Field} from "../apiModel";
import api from "../api";
import {findItemIndex} from "../common";

let findUrl = "/api/v1/field";
let readUrl = "/api/v1/field/"; // + id
let createUrl = "/api/v1/field";
let multiCreateUrl = "/api/v1/field/list";
let updateUrl = "/api/v1/field/"; // + id
let multiUpdateUrl = "/api/v1/field/list";
let deleteUrl = "/api/v1/field/"; // + id
let multiDeleteUrl = "/api/v1/field/list";
let findOrCreateUrl = "/api/v1/field";
let updateOrCreateUrl = "/api/v1/field";

const field = {
    actions: {
        createField(context, {data, filter, header, noMutation}) {

            let url = createUrl;
            if (Array.isArray && Array.isArray(data)) {
                url = multiCreateUrl
            }

            return api.create(url, data, filter, header)
                .then(function(response) {

					if(! noMutation) {
	                    context.commit("setField", response.Model);
					}

                    return response;
                })
                .catch(function(err) {
                    console.error(err);
                    throw(err);
                });
        },
        deleteField(context, {id, header, noMutation}) {

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
	                    context.commit("clearField");
					}
                    return response;
                })
                .catch(function(err) {
                    console.error(err);
                    throw(err);
                });
        },
        findField(context, {filter, header, isAppend, noMutation}) {

            return api.find(findUrl, filter, header)
                .then(function(response) {

					if(! noMutation) {
						if (isAppend) {
							context.commit("appendField__List", response.List);
						} else {
							context.commit("setField__List", response.List);
						}
					}

                    return response;
                })
                .catch(function(err) {
                    console.error(err);
                    throw(err);
                });
        },
        loadField(context, {id, filter, header, noMutation}) {

            return api.find(readUrl + id, filter, header)
                .then(function(response) {

					if(! noMutation) {
	                    context.commit("setField", response.Model);
					}
                    return response;
                })
                .catch(function(err) {
                    console.error(err);
                    throw(err);
                });
        },
        updateField(context, {id, data, filter, header, noMutation}) {

            let url = updateUrl + id;
            if (Array.isArray && Array.isArray(data)) {
                url = multiUpdateUrl
            }

            return api.update(url, data, filter, header)
                .then(function(response) {
					if(! noMutation) {
	                    context.commit("setField", response.Model);
					}
                    return response;
                })
                .catch(function(err) {
                    console.error(err);
                    throw(err);
                });
        },
        findOrCreateField(context, {id, data, filter, header, noMutation}) {

            return api.update(findOrCreateUrl, data, filter, header)
                .then(function(response) {

					if(! noMutation) {
	                    context.commit("setField", response.Model);
					}
                    return response;
                })
                .catch(function(err) {
                    console.error(err);
                    throw(err);
                });
        },
        clearListField(context) {
            context.commit("clearListField");
        },
        clearField(context) {
            context.commit("clearField");
        },
    },
    getters: {
        getField: (state) => {
            return state.Field;
        },
        getFieldById: state => id => {
            return state.Field__List.find(item => item.Id === id);
        },
        getListField: (state) => {
            return state.Field__List;
        },
        getRoute__Field: state => action => {
            return state.Field__Routes[action];
        },
        getRoutes__Field: state => {
            return state.Field__Routes;
        },
    },
    mutations: {
        setField(state, data) {
            state.Field = data;
        },
        setField__List(state, data) {
            state.Field__List = data || [];
        },
        appendField__List(state, data) {

            if (! state.Field__List) {
                state.Field__List = [];
            }

            state.Field__List = state.Field__List.concat(data);
        },
        clearField(state) {
            state.Field = new Field();
        },
        clearListField(state) {
            state.Field__List = [];
        },
		updateFieldById(state, data) {
    		let index = findItemIndex(state.Field__List, function(item) {
	        	return item.Id === data.Id;
	    	});
	    
	    	if (index || index === 0) {
		        state.Field__List.splice(index, 1, data);
    		}
		},
		deleteFieldFromList(state, id) {
		    let index = findItemIndex(state.Field__List, function(item) {
		        return item.Id === id;
		    });
		    
		    if (index || index === 0) {
		        state.Field__List.splice(index, 1);
		    }
		},
		addFieldItemToList(state, item) {

			if (state.Field__List === null) {
				state.Field__List = [];
			}

		    state.Field__List.push(item);
		},
    },
    state: {
        Field: new Field(),
        Field__List: [],
        Field__Routes: {
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

export default field;
