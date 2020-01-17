
import {ComponentTemplate} from "../apiModel";
import api from "../api";
import {findItemIndex} from "../common";

let findUrl = "/api/v1/componentTemplate";
let readUrl = "/api/v1/componentTemplate/"; // + id
let createUrl = "/api/v1/componentTemplate";
let multiCreateUrl = "/api/v1/componentTemplate/list";
let updateUrl = "/api/v1/componentTemplate/"; // + id
let multiUpdateUrl = "/api/v1/componentTemplate/list"; // + id
let deleteUrl = "/api/v1/componentTemplate/"; // + id
let multiDeleteUrl = "/api/v1/componentTemplate/list"; // + id
let findOrCreateUrl = "/api/v1/componentTemplate"; // + id

const componentTemplate = {
    actions: {
        createComponentTemplate(context, {data, filter, header}) {

            let url = createUrl;
            if (Array.isArray && Array.isArray(data)) {
                url = multiCreateUrl
            }

            return api.create(url, data, filter, header)
                .then(function(response) {

                    context.commit("setComponentTemplate", response.Model);

                    return response;
                })
                .catch(function(err) {
                    console.error(err);
                    throw(err);
                });
        },
        deleteComponentTemplate(context, {id, header}) {

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
                    context.commit("clearComponentTemplate");
                    return response;
                })
                .catch(function(err) {
                    console.error(err);
                    throw(err);
                });
        },
        findComponentTemplate(context, {filter, header, isAppend}) {

            return api.find(findUrl, filter, header)
                .then(function(response) {

                    if (isAppend) {
                        context.commit("appendComponentTemplate__List", response.List);
                    } else {
                        context.commit("setComponentTemplate__List", response.List);
                    }

                    return response;
                })
                .catch(function(err) {
                    console.error(err);
                    throw(err);
                });
        },
        loadComponentTemplate(context, {id, filter, header}) {

            return api.find(readUrl + id, filter, header)
                .then(function(response) {

                    context.commit("setComponentTemplate", response.Model);
                    return response;
                })
                .catch(function(err) {
                    console.error(err);
                    throw(err);
                });
        },
        updateComponentTemplate(context, {id, data, filter, header}) {

            let url = updateUrl + id;
            if (Array.isArray && Array.isArray(data)) {
                url = multiUpdateUrl
            }

            return api.update(url, data, filter, header)
                .then(function(response) {

                    context.commit("setComponentTemplate", response.Model);
                    return response;
                })
                .catch(function(err) {
                    console.error(err);
                    throw(err);
                });
        },
        findOrCreateComponentTemplate(context, {id, data, filter, header}) {

            return api.update(findOrCreateUrl, data, filter, header)
                .then(function(response) {

                    context.commit("setComponentTemplate", response.Model);
                    return response;
                })
                .catch(function(err) {
                    console.error(err);
                    throw(err);
                });
        },
        clearListComponentTemplate(context) {
            context.commit("clearListComponentTemplate");
        },
        clearComponentTemplate(context) {
            context.commit("clearComponentTemplate");
        },
    },
    getters: {
        getComponentTemplate: (state) => {
            return state.ComponentTemplate;
        },
        getComponentTemplateById: state => id => {
            return state.ComponentTemplate__List.find(item => item.Id === id);
        },
        getListComponentTemplate: (state) => {
            return state.ComponentTemplate__List;
        },
    },
    mutations: {
        setComponentTemplate(state, data) {
            state.ComponentTemplate = data;
        },
        setComponentTemplate__List(state, data) {
            state.ComponentTemplate__List = data || [];
        },
        appendComponentTemplate__List(state, data) {

            if (! state.ComponentTemplate__List) {
                state.ComponentTemplate__List = [];
            }

            state.ComponentTemplate__List = state.ComponentTemplate__List.concat(data);
        },
        clearComponentTemplate(state) {
            state.ComponentTemplate = new ComponentTemplate();
        },
        clearListComponentTemplate(state) {
            state.ComponentTemplate__List = [];
        },
		updateComponentTemplateById(state, data) {
    		let index = findItemIndex(state.ComponentTemplate__List, function(item) {
	        	return item.Id === data.Id;
	    	});
	    
	    	if (index || index === 0) {
		        state.ComponentTemplate__List.splice(index, 1, data);
    		}
		},
		deleteComponentTemplateFromList(state, id) {
		    let index = findItemIndex(state.ComponentTemplate__List, function(item) {
		        return item.Id === id;
		    });
		    
		    if (index || index === 0) {
		        state.ComponentTemplate__List.splice(index, 1);
		    }
		},
		addComponentTemplateItemToList(state, item) {

			if (state.ComponentTemplate__List === null) {
				state.ComponentTemplate__List = [];
			}

		    state.ComponentTemplate__List.push(item);
		},
    },
    state: {
        ComponentTemplate: new ComponentTemplate(),
        ComponentTemplate__List: [],
    },
};

export default componentTemplate;
