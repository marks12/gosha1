
import {PageTypeFilter} from "../apiModel";
import api from "../api";
import {findItemIndex} from "../common";

let findUrl = "/api/v1/pageTypeFilter";
let readUrl = "/api/v1/pageTypeFilter/"; // + id
let createUrl = "/api/v1/pageTypeFilter";
let multiCreateUrl = "/api/v1/pageTypeFilter/list";
let updateUrl = "/api/v1/pageTypeFilter/"; // + id
let multiUpdateUrl = "/api/v1/pageTypeFilter/list"; // + id
let deleteUrl = "/api/v1/pageTypeFilter/"; // + id
let multiDeleteUrl = "/api/v1/pageTypeFilter/list"; // + id
let findOrCreateUrl = "/api/v1/pageTypeFilter"; // + id

const pageTypeFilter = {
    actions: {
        createPageTypeFilter(context, {data, filter, header}) {

            let url = createUrl;
            if (Array.isArray && Array.isArray(data)) {
                url = multiCreateUrl
            }

            return api.create(url, data, filter, header)
                .then(function(response) {

                    context.commit("setPageTypeFilter", response.Model);

                    return response;
                })
                .catch(function(err) {
                    console.error(err);
                    throw(err);
                });
        },
        deletePageTypeFilter(context, {id, header}) {

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
                    context.commit("clearPageTypeFilter");
                    return response;
                })
                .catch(function(err) {
                    console.error(err);
                    throw(err);
                });
        },
        findPageTypeFilter(context, {filter, header, isAppend}) {

            return api.find(findUrl, filter, header)
                .then(function(response) {

                    if (isAppend) {
                        context.commit("appendPageTypeFilter__List", response.List);
                    } else {
                        context.commit("setPageTypeFilter__List", response.List);
                    }

                    return response;
                })
                .catch(function(err) {
                    console.error(err);
                    throw(err);
                });
        },
        loadPageTypeFilter(context, {id, filter, header}) {

            return api.find(readUrl + id, filter, header)
                .then(function(response) {

                    context.commit("setPageTypeFilter", response.Model);
                    return response;
                })
                .catch(function(err) {
                    console.error(err);
                    throw(err);
                });
        },
        updatePageTypeFilter(context, {id, data, filter, header}) {

            let url = updateUrl + id;
            if (Array.isArray && Array.isArray(data)) {
                url = multiUpdateUrl
            }

            return api.update(url, data, filter, header)
                .then(function(response) {

                    context.commit("setPageTypeFilter", response.Model);
                    return response;
                })
                .catch(function(err) {
                    console.error(err);
                    throw(err);
                });
        },
        findOrCreatePageTypeFilter(context, {id, data, filter, header}) {

            return api.update(findOrCreateUrl, data, filter, header)
                .then(function(response) {

                    context.commit("setPageTypeFilter", response.Model);
                    return response;
                })
                .catch(function(err) {
                    console.error(err);
                    throw(err);
                });
        },
        clearListPageTypeFilter(context) {
            context.commit("clearListPageTypeFilter");
        },
        clearPageTypeFilter(context) {
            context.commit("clearPageTypeFilter");
        },
    },
    getters: {
        getPageTypeFilter: (state) => {
            return state.PageTypeFilter;
        },
        getPageTypeFilterById: state => id => {
            return state.PageTypeFilter__List.find(item => item.Id === id);
        },
        getListPageTypeFilter: (state) => {
            return state.PageTypeFilter__List;
        },
    },
    mutations: {
        setPageTypeFilter(state, data) {
            state.PageTypeFilter = data;
        },
        setPageTypeFilter__List(state, data) {
            state.PageTypeFilter__List = data || [];
        },
        appendPageTypeFilter__List(state, data) {

            if (! state.PageTypeFilter__List) {
                state.PageTypeFilter__List = [];
            }

            state.PageTypeFilter__List = state.PageTypeFilter__List.concat(data);
        },
        clearPageTypeFilter(state) {
            state.PageTypeFilter = new PageTypeFilter();
        },
        clearListPageTypeFilter(state) {
            state.PageTypeFilter__List = [];
        },
		updatePageTypeFilterById(state, data) {
    		let index = findItemIndex(state.PageTypeFilter__List, function(item) {
	        	return item.Id === data.Id;
	    	});
	    
	    	if (index || index === 0) {
		        state.PageTypeFilter__List.splice(index, 1, data);
    		}
		},
		deletePageTypeFilterFromList(state, id) {
		    let index = findItemIndex(state.PageTypeFilter__List, function(item) {
		        return item.Id === id;
		    });
		    
		    if (index || index === 0) {
		        state.PageTypeFilter__List.splice(index, 1);
		    }
		},
		addPageTypeFilterItemToList(state, item) {

			if (state.PageTypeFilter__List === null) {
				state.PageTypeFilter__List = [];
			}

		    state.PageTypeFilter__List.push(item);
		},
    },
    state: {
        PageTypeFilter: new PageTypeFilter(),
        PageTypeFilter__List: [],
    },
};

export default pageTypeFilter;
