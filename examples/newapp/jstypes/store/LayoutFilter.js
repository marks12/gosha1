
import {LayoutFilter} from "../apiModel";
import api from "../api";
import {findItemIndex} from "../common";

let findUrl = "/api/v1/layoutFilter";
let readUrl = "/api/v1/layoutFilter/"; // + id
let createUrl = "/api/v1/layoutFilter";
let multiCreateUrl = "/api/v1/layoutFilter/list";
let updateUrl = "/api/v1/layoutFilter/"; // + id
let multiUpdateUrl = "/api/v1/layoutFilter/list"; // + id
let deleteUrl = "/api/v1/layoutFilter/"; // + id
let multiDeleteUrl = "/api/v1/layoutFilter/list"; // + id
let findOrCreateUrl = "/api/v1/layoutFilter"; // + id

const layoutFilter = {
    actions: {
        createLayoutFilter(context, {data, filter, header}) {

            let url = createUrl;
            if (Array.isArray && Array.isArray(data)) {
                url = multiCreateUrl
            }

            return api.create(url, data, filter, header)
                .then(function(response) {

                    context.commit("setLayoutFilter", response.Model);

                    return response;
                })
                .catch(function(err) {
                    console.error(err);
                    throw(err);
                });
        },
        deleteLayoutFilter(context, {id, header}) {

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
                    context.commit("clearLayoutFilter");
                    return response;
                })
                .catch(function(err) {
                    console.error(err);
                    throw(err);
                });
        },
        findLayoutFilter(context, {filter, header, isAppend}) {

            return api.find(findUrl, filter, header)
                .then(function(response) {

                    if (isAppend) {
                        context.commit("appendLayoutFilter__List", response.List);
                    } else {
                        context.commit("setLayoutFilter__List", response.List);
                    }

                    return response;
                })
                .catch(function(err) {
                    console.error(err);
                    throw(err);
                });
        },
        loadLayoutFilter(context, {id, filter, header}) {

            return api.find(readUrl + id, filter, header)
                .then(function(response) {

                    context.commit("setLayoutFilter", response.Model);
                    return response;
                })
                .catch(function(err) {
                    console.error(err);
                    throw(err);
                });
        },
        updateLayoutFilter(context, {id, data, filter, header}) {

            let url = updateUrl + id;
            if (Array.isArray && Array.isArray(data)) {
                url = multiUpdateUrl
            }

            return api.update(url, data, filter, header)
                .then(function(response) {

                    context.commit("setLayoutFilter", response.Model);
                    return response;
                })
                .catch(function(err) {
                    console.error(err);
                    throw(err);
                });
        },
        findOrCreateLayoutFilter(context, {id, data, filter, header}) {

            return api.update(findOrCreateUrl, data, filter, header)
                .then(function(response) {

                    context.commit("setLayoutFilter", response.Model);
                    return response;
                })
                .catch(function(err) {
                    console.error(err);
                    throw(err);
                });
        },
        clearListLayoutFilter(context) {
            context.commit("clearListLayoutFilter");
        },
        clearLayoutFilter(context) {
            context.commit("clearLayoutFilter");
        },
    },
    getters: {
        getLayoutFilter: (state) => {
            return state.LayoutFilter;
        },
        getLayoutFilterById: state => id => {
            return state.LayoutFilter__List.find(item => item.Id === id);
        },
        getListLayoutFilter: (state) => {
            return state.LayoutFilter__List;
        },
    },
    mutations: {
        setLayoutFilter(state, data) {
            state.LayoutFilter = data;
        },
        setLayoutFilter__List(state, data) {
            state.LayoutFilter__List = data || [];
        },
        appendLayoutFilter__List(state, data) {

            if (! state.LayoutFilter__List) {
                state.LayoutFilter__List = [];
            }

            state.LayoutFilter__List = state.LayoutFilter__List.concat(data);
        },
        clearLayoutFilter(state) {
            state.LayoutFilter = new LayoutFilter();
        },
        clearListLayoutFilter(state) {
            state.LayoutFilter__List = [];
        },
		updateLayoutFilterById(state, data) {
    		let index = findItemIndex(state.LayoutFilter__List, function(item) {
	        	return item.Id === data.Id;
	    	});
	    
	    	if (index || index === 0) {
		        state.LayoutFilter__List.splice(index, 1, data);
    		}
		},
		deleteLayoutFilterFromList(state, id) {
		    let index = findItemIndex(state.LayoutFilter__List, function(item) {
		        return item.Id === id;
		    });
		    
		    if (index || index === 0) {
		        state.LayoutFilter__List.splice(index, 1);
		    }
		},
		addLayoutFilterItemToList(state, item) {

			if (state.LayoutFilter__List === null) {
				state.LayoutFilter__List = [];
			}

		    state.LayoutFilter__List.push(item);
		},
    },
    state: {
        LayoutFilter: new LayoutFilter(),
        LayoutFilter__List: [],
    },
};

export default layoutFilter;
