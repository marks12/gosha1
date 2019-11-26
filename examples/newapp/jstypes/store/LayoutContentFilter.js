
import {LayoutContentFilter} from "../apiModel";
import api from "../api";
import {findItemIndex} from "../common";

let findUrl = "/api/v1/layoutContentFilter";
let readUrl = "/api/v1/layoutContentFilter/"; // + id
let createUrl = "/api/v1/layoutContentFilter";
let multiCreateUrl = "/api/v1/layoutContentFilter/list";
let updateUrl = "/api/v1/layoutContentFilter/"; // + id
let multiUpdateUrl = "/api/v1/layoutContentFilter/list"; // + id
let deleteUrl = "/api/v1/layoutContentFilter/"; // + id
let multiDeleteUrl = "/api/v1/layoutContentFilter/list"; // + id
let findOrCreateUrl = "/api/v1/layoutContentFilter"; // + id

const layoutContentFilter = {
    actions: {
        createLayoutContentFilter(context, {data, filter, header}) {

            let url = createUrl;
            if (Array.isArray && Array.isArray(data)) {
                url = multiCreateUrl
            }

            return api.create(url, data, filter, header)
                .then(function(response) {

                    context.commit("setLayoutContentFilter", response.Model);

                    return response;
                })
                .catch(function(err) {
                    console.error(err);
                    throw(err);
                });
        },
        deleteLayoutContentFilter(context, {id, header}) {

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
                    context.commit("clearLayoutContentFilter");
                    return response;
                })
                .catch(function(err) {
                    console.error(err);
                    throw(err);
                });
        },
        findLayoutContentFilter(context, {filter, header}) {

            return api.find(findUrl, filter, header)
                .then(function(response) {

                    context.commit("setLayoutContentFilter__List", response.List);
                    return response;
                })
                .catch(function(err) {
                    console.error(err);
                    throw(err);
                });
        },
        loadLayoutContentFilter(context, {id, filter, header}) {

            return api.find(readUrl + id, filter, header)
                .then(function(response) {

                    context.commit("setLayoutContentFilter", response.Model);
                    return response;
                })
                .catch(function(err) {
                    console.error(err);
                    throw(err);
                });
        },
        updateLayoutContentFilter(context, {id, data, filter, header}) {

            let url = updateUrl + id;
            if (Array.isArray && Array.isArray(data)) {
                url = multiUpdateUrl
            }

            return api.update(url, data, filter, header)
                .then(function(response) {

                    context.commit("setLayoutContentFilter", response.Model);
                    return response;
                })
                .catch(function(err) {
                    console.error(err);
                    throw(err);
                });
        },
        findOrCreateLayoutContentFilter(context, {id, data, filter, header}) {

            return api.update(findOrCreateUrl, data, filter, header)
                .then(function(response) {

                    context.commit("setLayoutContentFilter", response.Model);
                    return response;
                })
                .catch(function(err) {
                    console.error(err);
                    throw(err);
                });
        },
        clearListLayoutContentFilter(context) {
            context.commit("clearListLayoutContentFilter");
        },
        clearLayoutContentFilter(context) {
            context.commit("clearLayoutContentFilter");
        },
    },
    getters: {
        getLayoutContentFilter: (state) => {
            return state.LayoutContentFilter;
        },
        getLayoutContentFilterById: state => id => {
            return state.LayoutContentFilter__List.find(item => item.Id === id);
        },
        getListLayoutContentFilter: (state) => {
            return state.LayoutContentFilter__List;
        },
    },
    mutations: {
        setLayoutContentFilter(state, data) {
            state.LayoutContentFilter = data;
        },
        setLayoutContentFilter__List(state, data) {
            state.LayoutContentFilter__List = data || [];
        },
        clearLayoutContentFilter(state) {
            state.LayoutContentFilter = new LayoutContentFilter();
        },
        clearListLayoutContentFilter(state) {
            state.LayoutContentFilter__List = [];
        },
		updateLayoutContentFilterById(state, data) {
    		let index = findItemIndex(state.LayoutContentFilter__List, function(item) {
	        	return item.Id === data.Id;
	    	});
	    
	    	if (index || index === 0) {
		        state.LayoutContentFilter__List.splice(index, 1, data);
    		}
		},
		deleteLayoutContentFilterFromList(state, id) {
		    let index = findItemIndex(state.LayoutContentFilter__List, function(item) {
		        return item.Id === id;
		    });
		    
		    if (index || index === 0) {
		        state.LayoutContentFilter__List.splice(index, 1);
		    }
		},
		addLayoutContentFilterItemToList(state, item) {

			if (state.LayoutContentFilter__List === null) {
				state.LayoutContentFilter__List = [];
			}

		    state.LayoutContentFilter__List.push(item);
		},
    },
    state: {
        LayoutContentFilter: new LayoutContentFilter(),
        LayoutContentFilter__List: [],
    },
};

export default layoutContentFilter;
