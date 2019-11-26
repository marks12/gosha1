
import {Layout} from "../apiModel";
import api from "../api";
import {findItemIndex} from "../common";

let findUrl = "/api/v1/layout";
let readUrl = "/api/v1/layout/"; // + id
let createUrl = "/api/v1/layout";
let multiCreateUrl = "/api/v1/layout/list";
let updateUrl = "/api/v1/layout/"; // + id
let multiUpdateUrl = "/api/v1/layout/list"; // + id
let deleteUrl = "/api/v1/layout/"; // + id
let multiDeleteUrl = "/api/v1/layout/list"; // + id
let findOrCreateUrl = "/api/v1/layout"; // + id

const layout = {
    actions: {
        createLayout(context, {data, filter, header}) {

            let url = createUrl;
            if (Array.isArray && Array.isArray(data)) {
                url = multiCreateUrl
            }

            return api.create(url, data, filter, header)
                .then(function(response) {

                    context.commit("setLayout", response.Model);

                    return response;
                })
                .catch(function(err) {
                    console.error(err);
                    throw(err);
                });
        },
        deleteLayout(context, {id, header}) {

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
                    context.commit("clearLayout");
                    return response;
                })
                .catch(function(err) {
                    console.error(err);
                    throw(err);
                });
        },
        findLayout(context, {filter, header}) {

            return api.find(findUrl, filter, header)
                .then(function(response) {

                    context.commit("setLayout__List", response.List);
                    return response;
                })
                .catch(function(err) {
                    console.error(err);
                    throw(err);
                });
        },
        loadLayout(context, {id, filter, header}) {

            return api.find(readUrl + id, filter, header)
                .then(function(response) {

                    context.commit("setLayout", response.Model);
                    return response;
                })
                .catch(function(err) {
                    console.error(err);
                    throw(err);
                });
        },
        updateLayout(context, {id, data, filter, header}) {

            let url = updateUrl + id;
            if (Array.isArray && Array.isArray(data)) {
                url = multiUpdateUrl
            }

            return api.update(url, data, filter, header)
                .then(function(response) {

                    context.commit("setLayout", response.Model);
                    return response;
                })
                .catch(function(err) {
                    console.error(err);
                    throw(err);
                });
        },
        findOrCreateLayout(context, {id, data, filter, header}) {

            return api.update(findOrCreateUrl, data, filter, header)
                .then(function(response) {

                    context.commit("setLayout", response.Model);
                    return response;
                })
                .catch(function(err) {
                    console.error(err);
                    throw(err);
                });
        },
        clearListLayout(context) {
            context.commit("clearListLayout");
        },
        clearLayout(context) {
            context.commit("clearLayout");
        },
    },
    getters: {
        getLayout: (state) => {
            return state.Layout;
        },
        getLayoutById: state => id => {
            return state.Layout__List.find(item => item.Id === id);
        },
        getListLayout: (state) => {
            return state.Layout__List;
        },
    },
    mutations: {
        setLayout(state, data) {
            state.Layout = data;
        },
        setLayout__List(state, data) {
            state.Layout__List = data || [];
        },
        clearLayout(state) {
            state.Layout = new Layout();
        },
        clearListLayout(state) {
            state.Layout__List = [];
        },
		updateLayoutById(state, data) {
    		let index = findItemIndex(state.Layout__List, function(item) {
	        	return item.Id === data.Id;
	    	});
	    
	    	if (index || index === 0) {
		        state.Layout__List.splice(index, 1, data);
    		}
		},
		deleteLayoutFromList(state, id) {
		    let index = findItemIndex(state.Layout__List, function(item) {
		        return item.Id === id;
		    });
		    
		    if (index || index === 0) {
		        state.Layout__List.splice(index, 1);
		    }
		},
		addLayoutItemToList(state, item) {

			if (state.Layout__List === null) {
				state.Layout__List = [];
			}

		    state.Layout__List.push(item);
		},
    },
    state: {
        Layout: new Layout(),
        Layout__List: [],
    },
};

export default layout;
