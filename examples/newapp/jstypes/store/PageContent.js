
import {PageContent} from "../apiModel";
import api from "../api";
import {findItemIndex} from "../common";

let findUrl = "/api/v1/pageContent";
let readUrl = "/api/v1/pageContent/"; // + id
let createUrl = "/api/v1/pageContent";
let multiCreateUrl = "/api/v1/pageContent/list";
let updateUrl = "/api/v1/pageContent/"; // + id
let multiUpdateUrl = "/api/v1/pageContent/list"; // + id
let deleteUrl = "/api/v1/pageContent/"; // + id
let multiDeleteUrl = "/api/v1/pageContent/list"; // + id
let findOrCreateUrl = "/api/v1/pageContent"; // + id

const pageContent = {
    actions: {
        createPageContent(context, {data, filter, header}) {

            let url = createUrl;
            if (Array.isArray && Array.isArray(data)) {
                url = multiCreateUrl
            }

            return api.create(url, data, filter, header)
                .then(function(response) {

                    context.commit("setPageContent", response.Model);

                    return response;
                })
                .catch(function(err) {
                    console.error(err);
                    throw(err);
                });
        },
        deletePageContent(context, {id, header}) {

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
                    context.commit("clearPageContent");
                    return response;
                })
                .catch(function(err) {
                    console.error(err);
                    throw(err);
                });
        },
        findPageContent(context, {filter, header}) {

            return api.find(findUrl, filter, header)
                .then(function(response) {

                    context.commit("setPageContent__List", response.List);
                    return response;
                })
                .catch(function(err) {
                    console.error(err);
                    throw(err);
                });
        },
        loadPageContent(context, {id, filter, header}) {

            return api.find(readUrl + id, filter, header)
                .then(function(response) {

                    context.commit("setPageContent", response.Model);
                    return response;
                })
                .catch(function(err) {
                    console.error(err);
                    throw(err);
                });
        },
        updatePageContent(context, {id, data, filter, header}) {

            let url = updateUrl + id;
            if (Array.isArray && Array.isArray(data)) {
                url = multiUpdateUrl
            }

            return api.update(url, data, filter, header)
                .then(function(response) {

                    context.commit("setPageContent", response.Model);
                    return response;
                })
                .catch(function(err) {
                    console.error(err);
                    throw(err);
                });
        },
        findOrCreatePageContent(context, {id, data, filter, header}) {

            return api.update(findOrCreateUrl, data, filter, header)
                .then(function(response) {

                    context.commit("setPageContent", response.Model);
                    return response;
                })
                .catch(function(err) {
                    console.error(err);
                    throw(err);
                });
        },
        clearListPageContent(context) {
            context.commit("clearListPageContent");
        },
        clearPageContent(context) {
            context.commit("clearPageContent");
        },
    },
    getters: {
        getPageContent: (state) => {
            return state.PageContent;
        },
        getPageContentById: state => id => {
            return state.PageContent__List.find(item => item.Id === id);
        },
        getListPageContent: (state) => {
            return state.PageContent__List;
        },
    },
    mutations: {
        setPageContent(state, data) {
            state.PageContent = data;
        },
        setPageContent__List(state, data) {
            state.PageContent__List = data || [];
        },
        clearPageContent(state) {
            state.PageContent = new PageContent();
        },
        clearListPageContent(state) {
            state.PageContent__List = [];
        },
		updatePageContentById(state, data) {
    		let index = findItemIndex(state.PageContent__List, function(item) {
	        	return item.Id === data.Id;
	    	});
	    
	    	if (index || index === 0) {
		        state.PageContent__List.splice(index, 1, data);
    		}
		},
		deletePageContentFromList(state, id) {
		    let index = findItemIndex(state.PageContent__List, function(item) {
		        return item.Id === id;
		    });
		    
		    if (index || index === 0) {
		        state.PageContent__List.splice(index, 1);
		    }
		},
		addPageContentItemToList(state, item) {

			if (state.PageContent__List === null) {
				state.PageContent__List = [];
			}

		    state.PageContent__List.push(item);
		},
    },
    state: {
        PageContent: new PageContent(),
        PageContent__List: [],
    },
};

export default pageContent;
