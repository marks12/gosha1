
import {ComponentGroup} from "../apiModel";
import api from "../api";
import {findItemIndex} from "../common";

let findUrl = "/api/v1/componentGroup";
let readUrl = "/api/v1/componentGroup/"; // + id
let createUrl = "/api/v1/componentGroup";
let multiCreateUrl = "/api/v1/componentGroup/list";
let updateUrl = "/api/v1/componentGroup/"; // + id
let multiUpdateUrl = "/api/v1/componentGroup/list"; // + id
let deleteUrl = "/api/v1/componentGroup/"; // + id
let multiDeleteUrl = "/api/v1/componentGroup/list"; // + id
let findOrCreateUrl = "/api/v1/componentGroup"; // + id

const componentGroup = {
    actions: {
        createComponentGroup(context, {data, filter, header}) {

            let url = createUrl;
            if (Array.isArray && Array.isArray(data)) {
                url = multiCreateUrl
            }

            return api.create(url, data, filter, header)
                .then(function(response) {

                    context.commit("setComponentGroup", response.Model);

                    return response;
                })
                .catch(function(err) {
                    console.error(err);
                    throw(err);
                });
        },
        deleteComponentGroup(context, {id, header}) {

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
                    context.commit("clearComponentGroup");
                    return response;
                })
                .catch(function(err) {
                    console.error(err);
                    throw(err);
                });
        },
        findComponentGroup(context, {filter, header, isAppend}) {

            return api.find(findUrl, filter, header)
                .then(function(response) {

                    if (isAppend) {
                        context.commit("appendComponentGroup__List", response.List);
                    } else {
                        context.commit("setComponentGroup__List", response.List);
                    }

                    return response;
                })
                .catch(function(err) {
                    console.error(err);
                    throw(err);
                });
        },
        loadComponentGroup(context, {id, filter, header}) {

            return api.find(readUrl + id, filter, header)
                .then(function(response) {

                    context.commit("setComponentGroup", response.Model);
                    return response;
                })
                .catch(function(err) {
                    console.error(err);
                    throw(err);
                });
        },
        updateComponentGroup(context, {id, data, filter, header}) {

            let url = updateUrl + id;
            if (Array.isArray && Array.isArray(data)) {
                url = multiUpdateUrl
            }

            return api.update(url, data, filter, header)
                .then(function(response) {

                    context.commit("setComponentGroup", response.Model);
                    return response;
                })
                .catch(function(err) {
                    console.error(err);
                    throw(err);
                });
        },
        findOrCreateComponentGroup(context, {id, data, filter, header}) {

            return api.update(findOrCreateUrl, data, filter, header)
                .then(function(response) {

                    context.commit("setComponentGroup", response.Model);
                    return response;
                })
                .catch(function(err) {
                    console.error(err);
                    throw(err);
                });
        },
        clearListComponentGroup(context) {
            context.commit("clearListComponentGroup");
        },
        clearComponentGroup(context) {
            context.commit("clearComponentGroup");
        },
    },
    getters: {
        getComponentGroup: (state) => {
            return state.ComponentGroup;
        },
        getComponentGroupById: state => id => {
            return state.ComponentGroup__List.find(item => item.Id === id);
        },
        getListComponentGroup: (state) => {
            return state.ComponentGroup__List;
        },
    },
    mutations: {
        setComponentGroup(state, data) {
            state.ComponentGroup = data;
        },
        setComponentGroup__List(state, data) {
            state.ComponentGroup__List = data || [];
        },
        appendComponentGroup__List(state, data) {

            if (! state.ComponentGroup__List) {
                state.ComponentGroup__List = [];
            }

            state.ComponentGroup__List = state.ComponentGroup__List.concat(data);
        },
        clearComponentGroup(state) {
            state.ComponentGroup = new ComponentGroup();
        },
        clearListComponentGroup(state) {
            state.ComponentGroup__List = [];
        },
		updateComponentGroupById(state, data) {
    		let index = findItemIndex(state.ComponentGroup__List, function(item) {
	        	return item.Id === data.Id;
	    	});
	    
	    	if (index || index === 0) {
		        state.ComponentGroup__List.splice(index, 1, data);
    		}
		},
		deleteComponentGroupFromList(state, id) {
		    let index = findItemIndex(state.ComponentGroup__List, function(item) {
		        return item.Id === id;
		    });
		    
		    if (index || index === 0) {
		        state.ComponentGroup__List.splice(index, 1);
		    }
		},
		addComponentGroupItemToList(state, item) {

			if (state.ComponentGroup__List === null) {
				state.ComponentGroup__List = [];
			}

		    state.ComponentGroup__List.push(item);
		},
    },
    state: {
        ComponentGroup: new ComponentGroup(),
        ComponentGroup__List: [],
    },
};

export default componentGroup;
