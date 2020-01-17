
import {PageInfo} from "../apiModel";
import api from "../api";
import {findItemIndex} from "../common";

let findUrl = "/api/v1/pageInfo";
let readUrl = "/api/v1/pageInfo/"; // + id
let createUrl = "/api/v1/pageInfo";
let multiCreateUrl = "/api/v1/pageInfo/list";
let updateUrl = "/api/v1/pageInfo/"; // + id
let multiUpdateUrl = "/api/v1/pageInfo/list"; // + id
let deleteUrl = "/api/v1/pageInfo/"; // + id
let multiDeleteUrl = "/api/v1/pageInfo/list"; // + id
let findOrCreateUrl = "/api/v1/pageInfo"; // + id

const pageInfo = {
    actions: {
        createPageInfo(context, {data, filter, header}) {

            let url = createUrl;
            if (Array.isArray && Array.isArray(data)) {
                url = multiCreateUrl
            }

            return api.create(url, data, filter, header)
                .then(function(response) {

                    context.commit("setPageInfo", response.Model);

                    return response;
                })
                .catch(function(err) {
                    console.error(err);
                    throw(err);
                });
        },
        deletePageInfo(context, {id, header}) {

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
                    context.commit("clearPageInfo");
                    return response;
                })
                .catch(function(err) {
                    console.error(err);
                    throw(err);
                });
        },
        findPageInfo(context, {filter, header, isAppend}) {

            return api.find(findUrl, filter, header)
                .then(function(response) {

                    if (isAppend) {
                        context.commit("appendPageInfo__List", response.List);
                    } else {
                        context.commit("setPageInfo__List", response.List);
                    }

                    return response;
                })
                .catch(function(err) {
                    console.error(err);
                    throw(err);
                });
        },
        loadPageInfo(context, {id, filter, header}) {

            return api.find(readUrl + id, filter, header)
                .then(function(response) {

                    context.commit("setPageInfo", response.Model);
                    return response;
                })
                .catch(function(err) {
                    console.error(err);
                    throw(err);
                });
        },
        updatePageInfo(context, {id, data, filter, header}) {

            let url = updateUrl + id;
            if (Array.isArray && Array.isArray(data)) {
                url = multiUpdateUrl
            }

            return api.update(url, data, filter, header)
                .then(function(response) {

                    context.commit("setPageInfo", response.Model);
                    return response;
                })
                .catch(function(err) {
                    console.error(err);
                    throw(err);
                });
        },
        findOrCreatePageInfo(context, {id, data, filter, header}) {

            return api.update(findOrCreateUrl, data, filter, header)
                .then(function(response) {

                    context.commit("setPageInfo", response.Model);
                    return response;
                })
                .catch(function(err) {
                    console.error(err);
                    throw(err);
                });
        },
        clearListPageInfo(context) {
            context.commit("clearListPageInfo");
        },
        clearPageInfo(context) {
            context.commit("clearPageInfo");
        },
    },
    getters: {
        getPageInfo: (state) => {
            return state.PageInfo;
        },
        getPageInfoById: state => id => {
            return state.PageInfo__List.find(item => item.Id === id);
        },
        getListPageInfo: (state) => {
            return state.PageInfo__List;
        },
    },
    mutations: {
        setPageInfo(state, data) {
            state.PageInfo = data;
        },
        setPageInfo__List(state, data) {
            state.PageInfo__List = data || [];
        },
        appendPageInfo__List(state, data) {

            if (! state.PageInfo__List) {
                state.PageInfo__List = [];
            }

            state.PageInfo__List = state.PageInfo__List.concat(data);
        },
        clearPageInfo(state) {
            state.PageInfo = new PageInfo();
        },
        clearListPageInfo(state) {
            state.PageInfo__List = [];
        },
		updatePageInfoById(state, data) {
    		let index = findItemIndex(state.PageInfo__List, function(item) {
	        	return item.Id === data.Id;
	    	});
	    
	    	if (index || index === 0) {
		        state.PageInfo__List.splice(index, 1, data);
    		}
		},
		deletePageInfoFromList(state, id) {
		    let index = findItemIndex(state.PageInfo__List, function(item) {
		        return item.Id === id;
		    });
		    
		    if (index || index === 0) {
		        state.PageInfo__List.splice(index, 1);
		    }
		},
		addPageInfoItemToList(state, item) {

			if (state.PageInfo__List === null) {
				state.PageInfo__List = [];
			}

		    state.PageInfo__List.push(item);
		},
    },
    state: {
        PageInfo: new PageInfo(),
        PageInfo__List: [],
    },
};

export default pageInfo;
