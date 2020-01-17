
import {PageContentFilter} from "../apiModel";
import api from "../api";
import {findItemIndex} from "../common";

let findUrl = "/api/v1/pageContentFilter";
let readUrl = "/api/v1/pageContentFilter/"; // + id
let createUrl = "/api/v1/pageContentFilter";
let multiCreateUrl = "/api/v1/pageContentFilter/list";
let updateUrl = "/api/v1/pageContentFilter/"; // + id
let multiUpdateUrl = "/api/v1/pageContentFilter/list"; // + id
let deleteUrl = "/api/v1/pageContentFilter/"; // + id
let multiDeleteUrl = "/api/v1/pageContentFilter/list"; // + id
let findOrCreateUrl = "/api/v1/pageContentFilter"; // + id

const pageContentFilter = {
    actions: {
        createPageContentFilter(context, {data, filter, header}) {

            let url = createUrl;
            if (Array.isArray && Array.isArray(data)) {
                url = multiCreateUrl
            }

            return api.create(url, data, filter, header)
                .then(function(response) {

                    context.commit("setPageContentFilter", response.Model);

                    return response;
                })
                .catch(function(err) {
                    console.error(err);
                    throw(err);
                });
        },
        deletePageContentFilter(context, {id, header}) {

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
                    context.commit("clearPageContentFilter");
                    return response;
                })
                .catch(function(err) {
                    console.error(err);
                    throw(err);
                });
        },
        findPageContentFilter(context, {filter, header, isAppend}) {

            return api.find(findUrl, filter, header)
                .then(function(response) {

                    if (isAppend) {
                        context.commit("appendPageContentFilter__List", response.List);
                    } else {
                        context.commit("setPageContentFilter__List", response.List);
                    }

                    return response;
                })
                .catch(function(err) {
                    console.error(err);
                    throw(err);
                });
        },
        loadPageContentFilter(context, {id, filter, header}) {

            return api.find(readUrl + id, filter, header)
                .then(function(response) {

                    context.commit("setPageContentFilter", response.Model);
                    return response;
                })
                .catch(function(err) {
                    console.error(err);
                    throw(err);
                });
        },
        updatePageContentFilter(context, {id, data, filter, header}) {

            let url = updateUrl + id;
            if (Array.isArray && Array.isArray(data)) {
                url = multiUpdateUrl
            }

            return api.update(url, data, filter, header)
                .then(function(response) {

                    context.commit("setPageContentFilter", response.Model);
                    return response;
                })
                .catch(function(err) {
                    console.error(err);
                    throw(err);
                });
        },
        findOrCreatePageContentFilter(context, {id, data, filter, header}) {

            return api.update(findOrCreateUrl, data, filter, header)
                .then(function(response) {

                    context.commit("setPageContentFilter", response.Model);
                    return response;
                })
                .catch(function(err) {
                    console.error(err);
                    throw(err);
                });
        },
        clearListPageContentFilter(context) {
            context.commit("clearListPageContentFilter");
        },
        clearPageContentFilter(context) {
            context.commit("clearPageContentFilter");
        },
    },
    getters: {
        getPageContentFilter: (state) => {
            return state.PageContentFilter;
        },
        getPageContentFilterById: state => id => {
            return state.PageContentFilter__List.find(item => item.Id === id);
        },
        getListPageContentFilter: (state) => {
            return state.PageContentFilter__List;
        },
    },
    mutations: {
        setPageContentFilter(state, data) {
            state.PageContentFilter = data;
        },
        setPageContentFilter__List(state, data) {
            state.PageContentFilter__List = data || [];
        },
        appendPageContentFilter__List(state, data) {

            if (! state.PageContentFilter__List) {
                state.PageContentFilter__List = [];
            }

            state.PageContentFilter__List = state.PageContentFilter__List.concat(data);
        },
        clearPageContentFilter(state) {
            state.PageContentFilter = new PageContentFilter();
        },
        clearListPageContentFilter(state) {
            state.PageContentFilter__List = [];
        },
		updatePageContentFilterById(state, data) {
    		let index = findItemIndex(state.PageContentFilter__List, function(item) {
	        	return item.Id === data.Id;
	    	});
	    
	    	if (index || index === 0) {
		        state.PageContentFilter__List.splice(index, 1, data);
    		}
		},
		deletePageContentFilterFromList(state, id) {
		    let index = findItemIndex(state.PageContentFilter__List, function(item) {
		        return item.Id === id;
		    });
		    
		    if (index || index === 0) {
		        state.PageContentFilter__List.splice(index, 1);
		    }
		},
		addPageContentFilterItemToList(state, item) {

			if (state.PageContentFilter__List === null) {
				state.PageContentFilter__List = [];
			}

		    state.PageContentFilter__List.push(item);
		},
    },
    state: {
        PageContentFilter: new PageContentFilter(),
        PageContentFilter__List: [],
    },
};

export default pageContentFilter;
