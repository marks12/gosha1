
import {PageInfoFilter} from "../apiModel";
import api from "../api";
import {findItemIndex} from "../common";

let findUrl = "/api/v1/pageInfoFilter";
let readUrl = "/api/v1/pageInfoFilter/"; // + id
let createUrl = "/api/v1/pageInfoFilter";
let multiCreateUrl = "/api/v1/pageInfoFilter/list";
let updateUrl = "/api/v1/pageInfoFilter/"; // + id
let multiUpdateUrl = "/api/v1/pageInfoFilter/list"; // + id
let deleteUrl = "/api/v1/pageInfoFilter/"; // + id
let multiDeleteUrl = "/api/v1/pageInfoFilter/list"; // + id
let findOrCreateUrl = "/api/v1/pageInfoFilter"; // + id

const pageInfoFilter = {
    actions: {
        createPageInfoFilter(context, {data, filter, header}) {

            let url = createUrl;
            if (Array.isArray && Array.isArray(data)) {
                url = multiCreateUrl
            }

            return api.create(url, data, filter, header)
                .then(function(response) {

                    context.commit("setPageInfoFilter", response.Model);

                    return response;
                })
                .catch(function(err) {
                    console.error(err);
                    throw(err);
                });
        },
        deletePageInfoFilter(context, {id, header}) {

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
                    context.commit("clearPageInfoFilter");
                    return response;
                })
                .catch(function(err) {
                    console.error(err);
                    throw(err);
                });
        },
        findPageInfoFilter(context, {filter, header, isAppend}) {

            return api.find(findUrl, filter, header)
                .then(function(response) {

                    if (isAppend) {
                        context.commit("appendPageInfoFilter__List", response.List);
                    } else {
                        context.commit("setPageInfoFilter__List", response.List);
                    }

                    return response;
                })
                .catch(function(err) {
                    console.error(err);
                    throw(err);
                });
        },
        loadPageInfoFilter(context, {id, filter, header}) {

            return api.find(readUrl + id, filter, header)
                .then(function(response) {

                    context.commit("setPageInfoFilter", response.Model);
                    return response;
                })
                .catch(function(err) {
                    console.error(err);
                    throw(err);
                });
        },
        updatePageInfoFilter(context, {id, data, filter, header}) {

            let url = updateUrl + id;
            if (Array.isArray && Array.isArray(data)) {
                url = multiUpdateUrl
            }

            return api.update(url, data, filter, header)
                .then(function(response) {

                    context.commit("setPageInfoFilter", response.Model);
                    return response;
                })
                .catch(function(err) {
                    console.error(err);
                    throw(err);
                });
        },
        findOrCreatePageInfoFilter(context, {id, data, filter, header}) {

            return api.update(findOrCreateUrl, data, filter, header)
                .then(function(response) {

                    context.commit("setPageInfoFilter", response.Model);
                    return response;
                })
                .catch(function(err) {
                    console.error(err);
                    throw(err);
                });
        },
        clearListPageInfoFilter(context) {
            context.commit("clearListPageInfoFilter");
        },
        clearPageInfoFilter(context) {
            context.commit("clearPageInfoFilter");
        },
    },
    getters: {
        getPageInfoFilter: (state) => {
            return state.PageInfoFilter;
        },
        getPageInfoFilterById: state => id => {
            return state.PageInfoFilter__List.find(item => item.Id === id);
        },
        getListPageInfoFilter: (state) => {
            return state.PageInfoFilter__List;
        },
    },
    mutations: {
        setPageInfoFilter(state, data) {
            state.PageInfoFilter = data;
        },
        setPageInfoFilter__List(state, data) {
            state.PageInfoFilter__List = data || [];
        },
        appendPageInfoFilter__List(state, data) {

            if (! state.PageInfoFilter__List) {
                state.PageInfoFilter__List = [];
            }

            state.PageInfoFilter__List = state.PageInfoFilter__List.concat(data);
        },
        clearPageInfoFilter(state) {
            state.PageInfoFilter = new PageInfoFilter();
        },
        clearListPageInfoFilter(state) {
            state.PageInfoFilter__List = [];
        },
		updatePageInfoFilterById(state, data) {
    		let index = findItemIndex(state.PageInfoFilter__List, function(item) {
	        	return item.Id === data.Id;
	    	});
	    
	    	if (index || index === 0) {
		        state.PageInfoFilter__List.splice(index, 1, data);
    		}
		},
		deletePageInfoFilterFromList(state, id) {
		    let index = findItemIndex(state.PageInfoFilter__List, function(item) {
		        return item.Id === id;
		    });
		    
		    if (index || index === 0) {
		        state.PageInfoFilter__List.splice(index, 1);
		    }
		},
		addPageInfoFilterItemToList(state, item) {

			if (state.PageInfoFilter__List === null) {
				state.PageInfoFilter__List = [];
			}

		    state.PageInfoFilter__List.push(item);
		},
    },
    state: {
        PageInfoFilter: new PageInfoFilter(),
        PageInfoFilter__List: [],
    },
};

export default pageInfoFilter;
