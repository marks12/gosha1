
import {ComponentGroupFilter} from "../apiModel";
import api from "../api";
import {findItemIndex} from "../common";

let findUrl = "/api/v1/componentGroupFilter";
let readUrl = "/api/v1/componentGroupFilter/"; // + id
let createUrl = "/api/v1/componentGroupFilter";
let multiCreateUrl = "/api/v1/componentGroupFilter/list";
let updateUrl = "/api/v1/componentGroupFilter/"; // + id
let multiUpdateUrl = "/api/v1/componentGroupFilter/list"; // + id
let deleteUrl = "/api/v1/componentGroupFilter/"; // + id
let multiDeleteUrl = "/api/v1/componentGroupFilter/list"; // + id
let findOrCreateUrl = "/api/v1/componentGroupFilter"; // + id

const componentGroupFilter = {
    actions: {
        createComponentGroupFilter(context, {data, filter, header}) {

            let url = createUrl;
            if (Array.isArray && Array.isArray(data)) {
                url = multiCreateUrl
            }

            return api.create(url, data, filter, header)
                .then(function(response) {

                    context.commit("setComponentGroupFilter", response.Model);

                    return response;
                })
                .catch(function(err) {
                    console.error(err);
                    throw(err);
                });
        },
        deleteComponentGroupFilter(context, {id, header}) {

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
                    context.commit("clearComponentGroupFilter");
                    return response;
                })
                .catch(function(err) {
                    console.error(err);
                    throw(err);
                });
        },
        findComponentGroupFilter(context, {filter, header, isAppend}) {

            return api.find(findUrl, filter, header)
                .then(function(response) {

                    if (isAppend) {
                        context.commit("appendComponentGroupFilter__List", response.List);
                    } else {
                        context.commit("setComponentGroupFilter__List", response.List);
                    }

                    return response;
                })
                .catch(function(err) {
                    console.error(err);
                    throw(err);
                });
        },
        loadComponentGroupFilter(context, {id, filter, header}) {

            return api.find(readUrl + id, filter, header)
                .then(function(response) {

                    context.commit("setComponentGroupFilter", response.Model);
                    return response;
                })
                .catch(function(err) {
                    console.error(err);
                    throw(err);
                });
        },
        updateComponentGroupFilter(context, {id, data, filter, header}) {

            let url = updateUrl + id;
            if (Array.isArray && Array.isArray(data)) {
                url = multiUpdateUrl
            }

            return api.update(url, data, filter, header)
                .then(function(response) {

                    context.commit("setComponentGroupFilter", response.Model);
                    return response;
                })
                .catch(function(err) {
                    console.error(err);
                    throw(err);
                });
        },
        findOrCreateComponentGroupFilter(context, {id, data, filter, header}) {

            return api.update(findOrCreateUrl, data, filter, header)
                .then(function(response) {

                    context.commit("setComponentGroupFilter", response.Model);
                    return response;
                })
                .catch(function(err) {
                    console.error(err);
                    throw(err);
                });
        },
        clearListComponentGroupFilter(context) {
            context.commit("clearListComponentGroupFilter");
        },
        clearComponentGroupFilter(context) {
            context.commit("clearComponentGroupFilter");
        },
    },
    getters: {
        getComponentGroupFilter: (state) => {
            return state.ComponentGroupFilter;
        },
        getComponentGroupFilterById: state => id => {
            return state.ComponentGroupFilter__List.find(item => item.Id === id);
        },
        getListComponentGroupFilter: (state) => {
            return state.ComponentGroupFilter__List;
        },
    },
    mutations: {
        setComponentGroupFilter(state, data) {
            state.ComponentGroupFilter = data;
        },
        setComponentGroupFilter__List(state, data) {
            state.ComponentGroupFilter__List = data || [];
        },
        appendComponentGroupFilter__List(state, data) {

            if (! state.ComponentGroupFilter__List) {
                state.ComponentGroupFilter__List = [];
            }

            state.ComponentGroupFilter__List = state.ComponentGroupFilter__List.concat(data);
        },
        clearComponentGroupFilter(state) {
            state.ComponentGroupFilter = new ComponentGroupFilter();
        },
        clearListComponentGroupFilter(state) {
            state.ComponentGroupFilter__List = [];
        },
		updateComponentGroupFilterById(state, data) {
    		let index = findItemIndex(state.ComponentGroupFilter__List, function(item) {
	        	return item.Id === data.Id;
	    	});
	    
	    	if (index || index === 0) {
		        state.ComponentGroupFilter__List.splice(index, 1, data);
    		}
		},
		deleteComponentGroupFilterFromList(state, id) {
		    let index = findItemIndex(state.ComponentGroupFilter__List, function(item) {
		        return item.Id === id;
		    });
		    
		    if (index || index === 0) {
		        state.ComponentGroupFilter__List.splice(index, 1);
		    }
		},
		addComponentGroupFilterItemToList(state, item) {

			if (state.ComponentGroupFilter__List === null) {
				state.ComponentGroupFilter__List = [];
			}

		    state.ComponentGroupFilter__List.push(item);
		},
    },
    state: {
        ComponentGroupFilter: new ComponentGroupFilter(),
        ComponentGroupFilter__List: [],
    },
};

export default componentGroupFilter;
