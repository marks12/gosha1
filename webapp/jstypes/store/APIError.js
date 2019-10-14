
import {APIError} from "../apiModel";
import api from "../api";
import {findItemIndex} from "../common";

let findUrl = "/api/v1/aPIError";
let readUrl = "/api/v1/aPIError/"; // + id
let createUrl = "/api/v1/aPIError";
let updateUrl = "/api/v1/aPIError/"; // + id
let deleteUrl = "/api/v1/aPIError/"; // + id
let findOrCreateUrl = "/api/v1/aPIError"; // + id

const aPIError = {
    actions: {
        createAPIError(context, {data, filter, header}) {

            return api.create(createUrl, data, filter, header)
                .then(function(response) {

                    context.commit("setAPIError", response.Model);

                    return response;
                })
                .catch(function(err) {
                    return err;
                });
        },
        deleteAPIError(context, {id, header}) {

            return api.remove(deleteUrl + id, header)
                .then(function(response) {
                    context.commit("clearAPIError");
                    return response;
                })
                .catch(function(err) {
                    return err;
                });
        },
        findAPIError(context, {filter, header}) {

            return api.find(findUrl, filter, header)
                .then(function(response) {

                    context.commit("setAPIError__List", response.List);

                    return response;
                })
                .catch(function(err) {
                    return err;
                });
        },
        loadAPIError(context, {id, filter, header}) {

            return api.find(readUrl + id, filter, header)
                .then(function(response) {

                    context.commit("setAPIError", response.Model);

                    return response;
                })
                .catch(function(err) {
                    return err;
                });
        },
        updateAPIError(context, {id, data, filter, header}) {

            return api.update(updateUrl + id, data, filter, header)
                .then(function(response) {

                    context.commit("setAPIError", response.Model);

                    return response;
                })
                .catch(function(err) {
                    return err;
                });
        },
        findOrCreateAPIError(context, {id, data, filter, header}) {

            return api.update(findOrCreateUrl, data, filter, header)
                .then(function(response) {

                    context.commit("setAPIError", response.Model);

                    return response;
                })
                .catch(function(err) {
                    return err;
                });
        },
        clearListAPIError(context) {
            context.commit("clearListAPIError");
        },
    },
    getters: {
        getAPIError: (state) => {
            return state.APIError;
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
