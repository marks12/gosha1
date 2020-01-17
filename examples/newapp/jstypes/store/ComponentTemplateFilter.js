
import {ComponentTemplateFilter} from "../apiModel";
import api from "../api";
import {findItemIndex} from "../common";

let findUrl = "/api/v1/componentTemplateFilter";
let readUrl = "/api/v1/componentTemplateFilter/"; // + id
let createUrl = "/api/v1/componentTemplateFilter";
let multiCreateUrl = "/api/v1/componentTemplateFilter/list";
let updateUrl = "/api/v1/componentTemplateFilter/"; // + id
let multiUpdateUrl = "/api/v1/componentTemplateFilter/list"; // + id
let deleteUrl = "/api/v1/componentTemplateFilter/"; // + id
let multiDeleteUrl = "/api/v1/componentTemplateFilter/list"; // + id
let findOrCreateUrl = "/api/v1/componentTemplateFilter"; // + id

const componentTemplateFilter = {
    actions: {
        createComponentTemplateFilter(context, {data, filter, header}) {

            let url = createUrl;
            if (Array.isArray && Array.isArray(data)) {
                url = multiCreateUrl
            }

            return api.create(url, data, filter, header)
                .then(function(response) {

                    context.commit("setComponentTemplateFilter", response.Model);

                    return response;
                })
                .catch(function(err) {
                    console.error(err);
                    throw(err);
                });
        },
        deleteComponentTemplateFilter(context, {id, header}) {

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
                    context.commit("clearComponentTemplateFilter");
                    return response;
                })
                .catch(function(err) {
                    console.error(err);
                    throw(err);
                });
        },
        findComponentTemplateFilter(context, {filter, header, isAppend}) {

            return api.find(findUrl, filter, header)
                .then(function(response) {

                    if (isAppend) {
                        context.commit("appendComponentTemplateFilter__List", response.List);
                    } else {
                        context.commit("setComponentTemplateFilter__List", response.List);
                    }

                    return response;
                })
                .catch(function(err) {
                    console.error(err);
                    throw(err);
                });
        },
        loadComponentTemplateFilter(context, {id, filter, header}) {

            return api.find(readUrl + id, filter, header)
                .then(function(response) {

                    context.commit("setComponentTemplateFilter", response.Model);
                    return response;
                })
                .catch(function(err) {
                    console.error(err);
                    throw(err);
                });
        },
        updateComponentTemplateFilter(context, {id, data, filter, header}) {

            let url = updateUrl + id;
            if (Array.isArray && Array.isArray(data)) {
                url = multiUpdateUrl
            }

            return api.update(url, data, filter, header)
                .then(function(response) {

                    context.commit("setComponentTemplateFilter", response.Model);
                    return response;
                })
                .catch(function(err) {
                    console.error(err);
                    throw(err);
                });
        },
        findOrCreateComponentTemplateFilter(context, {id, data, filter, header}) {

            return api.update(findOrCreateUrl, data, filter, header)
                .then(function(response) {

                    context.commit("setComponentTemplateFilter", response.Model);
                    return response;
                })
                .catch(function(err) {
                    console.error(err);
                    throw(err);
                });
        },
        clearListComponentTemplateFilter(context) {
            context.commit("clearListComponentTemplateFilter");
        },
        clearComponentTemplateFilter(context) {
            context.commit("clearComponentTemplateFilter");
        },
    },
    getters: {
        getComponentTemplateFilter: (state) => {
            return state.ComponentTemplateFilter;
        },
        getComponentTemplateFilterById: state => id => {
            return state.ComponentTemplateFilter__List.find(item => item.Id === id);
        },
        getListComponentTemplateFilter: (state) => {
            return state.ComponentTemplateFilter__List;
        },
    },
    mutations: {
        setComponentTemplateFilter(state, data) {
            state.ComponentTemplateFilter = data;
        },
        setComponentTemplateFilter__List(state, data) {
            state.ComponentTemplateFilter__List = data || [];
        },
        appendComponentTemplateFilter__List(state, data) {

            if (! state.ComponentTemplateFilter__List) {
                state.ComponentTemplateFilter__List = [];
            }

            state.ComponentTemplateFilter__List = state.ComponentTemplateFilter__List.concat(data);
        },
        clearComponentTemplateFilter(state) {
            state.ComponentTemplateFilter = new ComponentTemplateFilter();
        },
        clearListComponentTemplateFilter(state) {
            state.ComponentTemplateFilter__List = [];
        },
		updateComponentTemplateFilterById(state, data) {
    		let index = findItemIndex(state.ComponentTemplateFilter__List, function(item) {
	        	return item.Id === data.Id;
	    	});
	    
	    	if (index || index === 0) {
		        state.ComponentTemplateFilter__List.splice(index, 1, data);
    		}
		},
		deleteComponentTemplateFilterFromList(state, id) {
		    let index = findItemIndex(state.ComponentTemplateFilter__List, function(item) {
		        return item.Id === id;
		    });
		    
		    if (index || index === 0) {
		        state.ComponentTemplateFilter__List.splice(index, 1);
		    }
		},
		addComponentTemplateFilterItemToList(state, item) {

			if (state.ComponentTemplateFilter__List === null) {
				state.ComponentTemplateFilter__List = [];
			}

		    state.ComponentTemplateFilter__List.push(item);
		},
    },
    state: {
        ComponentTemplateFilter: new ComponentTemplateFilter(),
        ComponentTemplateFilter__List: [],
    },
};

export default componentTemplateFilter;
