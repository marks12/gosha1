
import {PageTemplate} from "../apiModel";
import api from "../api";
import {findItemIndex} from "../common";

let findUrl = "/api/v1/pageTemplate";
let readUrl = "/api/v1/pageTemplate/"; // + id
let createUrl = "/api/v1/pageTemplate";
let multiCreateUrl = "/api/v1/pageTemplate/list";
let updateUrl = "/api/v1/pageTemplate/"; // + id
let multiUpdateUrl = "/api/v1/pageTemplate/list"; // + id
let deleteUrl = "/api/v1/pageTemplate/"; // + id
let multiDeleteUrl = "/api/v1/pageTemplate/list"; // + id
let findOrCreateUrl = "/api/v1/pageTemplate"; // + id

const pageTemplate = {
    actions: {
        createPageTemplate(context, {data, filter, header}) {

            let url = createUrl;
            if (Array.isArray && Array.isArray(data)) {
                url = multiCreateUrl
            }

            return api.create(url, data, filter, header)
                .then(function(response) {

                    context.commit("setPageTemplate", response.Model);

                    return response;
                })
                .catch(function(err) {
                    console.error(err);
                    throw(err);
                });
        },
        deletePageTemplate(context, {id, header}) {

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
                    context.commit("clearPageTemplate");
                    return response;
                })
                .catch(function(err) {
                    console.error(err);
                    throw(err);
                });
        },
        findPageTemplate(context, {filter, header}) {

            return api.find(findUrl, filter, header)
                .then(function(response) {

                    context.commit("setPageTemplate__List", response.List);
                    return response;
                })
                .catch(function(err) {
                    console.error(err);
                    throw(err);
                });
        },
        loadPageTemplate(context, {id, filter, header}) {

            return api.find(readUrl + id, filter, header)
                .then(function(response) {

                    context.commit("setPageTemplate", response.Model);
                    return response;
                })
                .catch(function(err) {
                    console.error(err);
                    throw(err);
                });
        },
        updatePageTemplate(context, {id, data, filter, header}) {

            let url = updateUrl + id;
            if (Array.isArray && Array.isArray(data)) {
                url = multiUpdateUrl
            }

            return api.update(url, data, filter, header)
                .then(function(response) {

                    context.commit("setPageTemplate", response.Model);
                    return response;
                })
                .catch(function(err) {
                    console.error(err);
                    throw(err);
                });
        },
        findOrCreatePageTemplate(context, {id, data, filter, header}) {

            return api.update(findOrCreateUrl, data, filter, header)
                .then(function(response) {

                    context.commit("setPageTemplate", response.Model);
                    return response;
                })
                .catch(function(err) {
                    console.error(err);
                    throw(err);
                });
        },
        clearListPageTemplate(context) {
            context.commit("clearListPageTemplate");
        },
        clearPageTemplate(context) {
            context.commit("clearPageTemplate");
        },
    },
    getters: {
        getPageTemplate: (state) => {
            return state.PageTemplate;
        },
        getPageTemplateById: state => id => {
            return state.PageTemplate__List.find(item => item.Id === id);
        },
        getListPageTemplate: (state) => {
            return state.PageTemplate__List;
        },
    },
    mutations: {
        setPageTemplate(state, data) {
            state.PageTemplate = data;
        },
        setPageTemplate__List(state, data) {
            state.PageTemplate__List = data || [];
        },
        clearPageTemplate(state) {
            state.PageTemplate = new PageTemplate();
        },
        clearListPageTemplate(state) {
            state.PageTemplate__List = [];
        },
		updatePageTemplateById(state, data) {
    		let index = findItemIndex(state.PageTemplate__List, function(item) {
	        	return item.Id === data.Id;
	    	});
	    
	    	if (index || index === 0) {
		        state.PageTemplate__List.splice(index, 1, data);
    		}
		},
		deletePageTemplateFromList(state, id) {
		    let index = findItemIndex(state.PageTemplate__List, function(item) {
		        return item.Id === id;
		    });
		    
		    if (index || index === 0) {
		        state.PageTemplate__List.splice(index, 1);
		    }
		},
		addPageTemplateItemToList(state, item) {

			if (state.PageTemplate__List === null) {
				state.PageTemplate__List = [];
			}

		    state.PageTemplate__List.push(item);
		},
    },
    state: {
        PageTemplate: new PageTemplate(),
        PageTemplate__List: [],
    },
};

export default pageTemplate;
