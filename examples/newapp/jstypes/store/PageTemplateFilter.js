
import {PageTemplateFilter} from "../apiModel";
import api from "../api";
import {findItemIndex} from "../common";

let findUrl = "/api/v1/pageTemplateFilter";
let readUrl = "/api/v1/pageTemplateFilter/"; // + id
let createUrl = "/api/v1/pageTemplateFilter";
let multiCreateUrl = "/api/v1/pageTemplateFilter/list";
let updateUrl = "/api/v1/pageTemplateFilter/"; // + id
let multiUpdateUrl = "/api/v1/pageTemplateFilter/list"; // + id
let deleteUrl = "/api/v1/pageTemplateFilter/"; // + id
let multiDeleteUrl = "/api/v1/pageTemplateFilter/list"; // + id
let findOrCreateUrl = "/api/v1/pageTemplateFilter"; // + id

const pageTemplateFilter = {
    actions: {
        createPageTemplateFilter(context, {data, filter, header}) {

            let url = createUrl;
            if (Array.isArray && Array.isArray(data)) {
                url = multiCreateUrl
            }

            return api.create(url, data, filter, header)
                .then(function(response) {

                    context.commit("setPageTemplateFilter", response.Model);

                    return response;
                })
                .catch(function(err) {
                    console.error(err);
                    throw(err);
                });
        },
        deletePageTemplateFilter(context, {id, header}) {

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
                    context.commit("clearPageTemplateFilter");
                    return response;
                })
                .catch(function(err) {
                    console.error(err);
                    throw(err);
                });
        },
        findPageTemplateFilter(context, {filter, header}) {

            return api.find(findUrl, filter, header)
                .then(function(response) {

                    context.commit("setPageTemplateFilter__List", response.List);
                    return response;
                })
                .catch(function(err) {
                    console.error(err);
                    throw(err);
                });
        },
        loadPageTemplateFilter(context, {id, filter, header}) {

            return api.find(readUrl + id, filter, header)
                .then(function(response) {

                    context.commit("setPageTemplateFilter", response.Model);
                    return response;
                })
                .catch(function(err) {
                    console.error(err);
                    throw(err);
                });
        },
        updatePageTemplateFilter(context, {id, data, filter, header}) {

            let url = updateUrl + id;
            if (Array.isArray && Array.isArray(data)) {
                url = multiUpdateUrl
            }

            return api.update(url, data, filter, header)
                .then(function(response) {

                    context.commit("setPageTemplateFilter", response.Model);
                    return response;
                })
                .catch(function(err) {
                    console.error(err);
                    throw(err);
                });
        },
        findOrCreatePageTemplateFilter(context, {id, data, filter, header}) {

            return api.update(findOrCreateUrl, data, filter, header)
                .then(function(response) {

                    context.commit("setPageTemplateFilter", response.Model);
                    return response;
                })
                .catch(function(err) {
                    console.error(err);
                    throw(err);
                });
        },
        clearListPageTemplateFilter(context) {
            context.commit("clearListPageTemplateFilter");
        },
        clearPageTemplateFilter(context) {
            context.commit("clearPageTemplateFilter");
        },
    },
    getters: {
        getPageTemplateFilter: (state) => {
            return state.PageTemplateFilter;
        },
        getPageTemplateFilterById: state => id => {
            return state.PageTemplateFilter__List.find(item => item.Id === id);
        },
        getListPageTemplateFilter: (state) => {
            return state.PageTemplateFilter__List;
        },
    },
    mutations: {
        setPageTemplateFilter(state, data) {
            state.PageTemplateFilter = data;
        },
        setPageTemplateFilter__List(state, data) {
            state.PageTemplateFilter__List = data || [];
        },
        clearPageTemplateFilter(state) {
            state.PageTemplateFilter = new PageTemplateFilter();
        },
        clearListPageTemplateFilter(state) {
            state.PageTemplateFilter__List = [];
        },
		updatePageTemplateFilterById(state, data) {
    		let index = findItemIndex(state.PageTemplateFilter__List, function(item) {
	        	return item.Id === data.Id;
	    	});
	    
	    	if (index || index === 0) {
		        state.PageTemplateFilter__List.splice(index, 1, data);
    		}
		},
		deletePageTemplateFilterFromList(state, id) {
		    let index = findItemIndex(state.PageTemplateFilter__List, function(item) {
		        return item.Id === id;
		    });
		    
		    if (index || index === 0) {
		        state.PageTemplateFilter__List.splice(index, 1);
		    }
		},
		addPageTemplateFilterItemToList(state, item) {

			if (state.PageTemplateFilter__List === null) {
				state.PageTemplateFilter__List = [];
			}

		    state.PageTemplateFilter__List.push(item);
		},
    },
    state: {
        PageTemplateFilter: new PageTemplateFilter(),
        PageTemplateFilter__List: [],
    },
};

export default pageTemplateFilter;
