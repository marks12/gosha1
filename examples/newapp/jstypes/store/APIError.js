
import {APIError} from "../apiModel";
import api from "../api";
import {findItemIndex} from "../common";

let findUrl = "/api/v1/aPIError";
let readUrl = "/api/v1/aPIError/"; // + id
let createUrl = "/api/v1/aPIError";
let multiCreateUrl = "/api/v1/aPIError/list";
let updateUrl = "/api/v1/aPIError/"; // + id
let multiUpdateUrl = "/api/v1/aPIError/list"; // + id
let deleteUrl = "/api/v1/aPIError/"; // + id
let multiDeleteUrl = "/api/v1/aPIError/list"; // + id
let findOrCreateUrl = "/api/v1/aPIError"; // + id

const aPIError = {
    actions: {
        createAPIError(context, {data, filter, header}) {

            let url = createUrl;
            if (Array.isArray && Array.isArray(data)) {
                url = multiCreateUrl
            }

            return api.create(url, data, filter, header)
                .then(function(response) {

                    context.commit("setAPIError", response.Model);

                    return response;
                })
                .catch(function(err) {
                    console.error(err);
                    throw(err);
                });
        },
        deleteAPIError(context, {id, header}) {

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
                    context.commit("clearAPIError");
                    return response;
                })
                .catch(function(err) {
                    console.error(err);
                    throw(err);
                });
        },
        findAPIError(context, {filter, header}) {

            return api.find(findUrl, filter, header)
                .then(function(response) {

                    context.commit("setAPIError__List", response.List);
                    return response;
                })
                .catch(function(err) {
                    console.error(err);
                    throw(err);
                });
        },
        loadAPIError(context, {id, filter, header}) {

            return api.find(readUrl + id, filter, header)
                .then(function(response) {

                    context.commit("setAPIError", response.Model);
                    return response;
                })
                .catch(function(err) {
                    console.error(err);
                    throw(err);
                });
        },
        updateAPIError(context, {id, data, filter, header}) {

            let url = updateUrl + id;
            if (Array.isArray && Array.isArray(data)) {
                url = multiUpdateUrl
            }

            return api.update(url, data, filter, header)
                .then(function(response) {

                    context.commit("setAPIError", response.Model);
                    return response;
                })
                .catch(function(err) {
                    console.error(err);
                    throw(err);
                });
        },
        findOrCreateAPIError(context, {id, data, filter, header}) {

            return api.update(findOrCreateUrl, data, filter, header)
                .then(function(response) {

                    context.commit("setAPIError", response.Model);
                    return response;
                })
                .catch(function(err) {
                    console.error(err);
                    throw(err);
                });
        },
        clearListAPIError(context) {
            context.commit("clearListAPIError");
        },
        clearAPIError(context) {
            context.commit("clearAPIError");
        },
    },
    getters: {
        getAPIError: (state) => {
            return state.APIError;
        },
        getAPIErrorById: state => id => {
            return state.APIError__List.find(item => item.Id === id);
        },
        getListAPIError: (state) => {
            return state.APIError__List;
        },
    },
    mutations: {
        setAPIError(state, data) {
            state.APIError = data;
        },
        setAPIError__List(state, data) {
            state.APIError__List = data || [];
        },
        clearAPIError(state) {
            state.APIError = new APIError();
        },
        clearListAPIError(state) {
            state.APIError__List = [];
        },
		updateAPIErrorById(state, data) {
    		let index = findItemIndex(state.APIError__List, function(item) {
	        	return item.Id === data.Id;
	    	});
	    
	    	if (index || index === 0) {
		        state.APIError__List.splice(index, 1, data);
    		}
		},
		deleteAPIErrorFromList(state, id) {
		    let index = findItemIndex(state.APIError__List, function(item) {
		        return item.Id === id;
		    });
		    
		    if (index || index === 0) {
		        state.APIError__List.splice(index, 1);
		    }
		},
		addAPIErrorItemToList(state, item) {

			if (state.APIError__List === null) {
				state.APIError__List = [];
			}

		    state.APIError__List.push(item);
		},
    },
    state: {
        APIError: new APIError(),
        APIError__List: [],
    },
};

export default aPIError;
