
import {LayoutContent} from "../apiModel";
import api from "../api";
import {findItemIndex} from "../common";

let findUrl = "/api/v1/layoutContent";
let readUrl = "/api/v1/layoutContent/"; // + id
let createUrl = "/api/v1/layoutContent";
let multiCreateUrl = "/api/v1/layoutContent/list";
let updateUrl = "/api/v1/layoutContent/"; // + id
let multiUpdateUrl = "/api/v1/layoutContent/list"; // + id
let deleteUrl = "/api/v1/layoutContent/"; // + id
let multiDeleteUrl = "/api/v1/layoutContent/list"; // + id
let findOrCreateUrl = "/api/v1/layoutContent"; // + id

const layoutContent = {
    actions: {
        createLayoutContent(context, {data, filter, header}) {

            let url = createUrl;
            if (Array.isArray && Array.isArray(data)) {
                url = multiCreateUrl
            }

            return api.create(url, data, filter, header)
                .then(function(response) {

                    context.commit("setLayoutContent", response.Model);

                    return response;
                })
                .catch(function(err) {
                    console.error(err);
                    throw(err);
                });
        },
        deleteLayoutContent(context, {id, header}) {

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
                    context.commit("clearLayoutContent");
                    return response;
                })
                .catch(function(err) {
                    console.error(err);
                    throw(err);
                });
        },
        findLayoutContent(context, {filter, header, isAppend}) {

            return api.find(findUrl, filter, header)
                .then(function(response) {

                    if (isAppend) {
                        context.commit("appendLayoutContent__List", response.List);
                    } else {
                        context.commit("setLayoutContent__List", response.List);
                    }

                    return response;
                })
                .catch(function(err) {
                    console.error(err);
                    throw(err);
                });
        },
        loadLayoutContent(context, {id, filter, header}) {

            return api.find(readUrl + id, filter, header)
                .then(function(response) {

                    context.commit("setLayoutContent", response.Model);
                    return response;
                })
                .catch(function(err) {
                    console.error(err);
                    throw(err);
                });
        },
        updateLayoutContent(context, {id, data, filter, header}) {

            let url = updateUrl + id;
            if (Array.isArray && Array.isArray(data)) {
                url = multiUpdateUrl
            }

            return api.update(url, data, filter, header)
                .then(function(response) {

                    context.commit("setLayoutContent", response.Model);
                    return response;
                })
                .catch(function(err) {
                    console.error(err);
                    throw(err);
                });
        },
        findOrCreateLayoutContent(context, {id, data, filter, header}) {

            return api.update(findOrCreateUrl, data, filter, header)
                .then(function(response) {

                    context.commit("setLayoutContent", response.Model);
                    return response;
                })
                .catch(function(err) {
                    console.error(err);
                    throw(err);
                });
        },
        clearListLayoutContent(context) {
            context.commit("clearListLayoutContent");
        },
        clearLayoutContent(context) {
            context.commit("clearLayoutContent");
        },
    },
    getters: {
        getLayoutContent: (state) => {
            return state.LayoutContent;
        },
        getLayoutContentById: state => id => {
            return state.LayoutContent__List.find(item => item.Id === id);
        },
        getListLayoutContent: (state) => {
            return state.LayoutContent__List;
        },
    },
    mutations: {
        setLayoutContent(state, data) {
            state.LayoutContent = data;
        },
        setLayoutContent__List(state, data) {
            state.LayoutContent__List = data || [];
        },
        appendLayoutContent__List(state, data) {

            if (! state.LayoutContent__List) {
                state.LayoutContent__List = [];
            }

            state.LayoutContent__List = state.LayoutContent__List.concat(data);
        },
        clearLayoutContent(state) {
            state.LayoutContent = new LayoutContent();
        },
        clearListLayoutContent(state) {
            state.LayoutContent__List = [];
        },
		updateLayoutContentById(state, data) {
    		let index = findItemIndex(state.LayoutContent__List, function(item) {
	        	return item.Id === data.Id;
	    	});
	    
	    	if (index || index === 0) {
		        state.LayoutContent__List.splice(index, 1, data);
    		}
		},
		deleteLayoutContentFromList(state, id) {
		    let index = findItemIndex(state.LayoutContent__List, function(item) {
		        return item.Id === id;
		    });
		    
		    if (index || index === 0) {
		        state.LayoutContent__List.splice(index, 1);
		    }
		},
		addLayoutContentItemToList(state, item) {

			if (state.LayoutContent__List === null) {
				state.LayoutContent__List = [];
			}

		    state.LayoutContent__List.push(item);
		},
    },
    state: {
        LayoutContent: new LayoutContent(),
        LayoutContent__List: [],
    },
};

export default layoutContent;
