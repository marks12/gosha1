
import {PageType} from "../apiModel";
import api from "../api";
import {findItemIndex} from "../common";

let findUrl = "/api/v1/pageType";
let readUrl = "/api/v1/pageType/"; // + id
let createUrl = "/api/v1/pageType";
let multiCreateUrl = "/api/v1/pageType/list";
let updateUrl = "/api/v1/pageType/"; // + id
let multiUpdateUrl = "/api/v1/pageType/list"; // + id
let deleteUrl = "/api/v1/pageType/"; // + id
let multiDeleteUrl = "/api/v1/pageType/list"; // + id
let findOrCreateUrl = "/api/v1/pageType"; // + id

const pageType = {
    actions: {
        createPageType(context, {data, filter, header}) {

            let url = createUrl;
            if (Array.isArray && Array.isArray(data)) {
                url = multiCreateUrl
            }

            return api.create(url, data, filter, header)
                .then(function(response) {

                    context.commit("setPageType", response.Model);

                    return response;
                })
                .catch(function(err) {
                    console.error(err);
                    throw(err);
                });
        },
        deletePageType(context, {id, header}) {

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
                    context.commit("clearPageType");
                    return response;
                })
                .catch(function(err) {
                    console.error(err);
                    throw(err);
                });
        },
        findPageType(context, {filter, header}) {

            return api.find(findUrl, filter, header)
                .then(function(response) {

                    context.commit("setPageType__List", response.List);
                    return response;
                })
                .catch(function(err) {
                    console.error(err);
                    throw(err);
                });
        },
        loadPageType(context, {id, filter, header}) {

            return api.find(readUrl + id, filter, header)
                .then(function(response) {

                    context.commit("setPageType", response.Model);
                    return response;
                })
                .catch(function(err) {
                    console.error(err);
                    throw(err);
                });
        },
        updatePageType(context, {id, data, filter, header}) {

            let url = updateUrl + id;
            if (Array.isArray && Array.isArray(data)) {
                url = multiUpdateUrl
            }

            return api.update(url, data, filter, header)
                .then(function(response) {

                    context.commit("setPageType", response.Model);
                    return response;
                })
                .catch(function(err) {
                    console.error(err);
                    throw(err);
                });
        },
        findOrCreatePageType(context, {id, data, filter, header}) {

            return api.update(findOrCreateUrl, data, filter, header)
                .then(function(response) {

                    context.commit("setPageType", response.Model);
                    return response;
                })
                .catch(function(err) {
                    console.error(err);
                    throw(err);
                });
        },
        clearListPageType(context) {
            context.commit("clearListPageType");
        },
        clearPageType(context) {
            context.commit("clearPageType");
        },
    },
    getters: {
        getPageType: (state) => {
            return state.PageType;
        },
        getPageTypeById: state => id => {
            return state.PageType__List.find(item => item.Id === id);
        },
        getListPageType: (state) => {
            return state.PageType__List;
        },
    },
    mutations: {
        setPageType(state, data) {
            state.PageType = data;
        },
        setPageType__List(state, data) {
            state.PageType__List = data || [];
        },
        clearPageType(state) {
            state.PageType = new PageType();
        },
        clearListPageType(state) {
            state.PageType__List = [];
        },
		updatePageTypeById(state, data) {
    		let index = findItemIndex(state.PageType__List, function(item) {
	        	return item.Id === data.Id;
	    	});
	    
	    	if (index || index === 0) {
		        state.PageType__List.splice(index, 1, data);
    		}
		},
		deletePageTypeFromList(state, id) {
		    let index = findItemIndex(state.PageType__List, function(item) {
		        return item.Id === id;
		    });
		    
		    if (index || index === 0) {
		        state.PageType__List.splice(index, 1);
		    }
		},
		addPageTypeItemToList(state, item) {

			if (state.PageType__List === null) {
				state.PageType__List = [];
			}

		    state.PageType__List.push(item);
		},
    },
    state: {
        PageType: new PageType(),
        PageType__List: [],
    },
};

export default pageType;
