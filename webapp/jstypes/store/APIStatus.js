
import {APIStatus} from "../apiModel";
import api from "../api";
import {findItemIndex} from "../common";

let findUrl = "/api/v1/aPIStatus";
let readUrl = "/api/v1/aPIStatus/"; // + id
let createUrl = "/api/v1/aPIStatus";
let updateUrl = "/api/v1/aPIStatus/"; // + id
let deleteUrl = "/api/v1/aPIStatus/"; // + id
let findOrCreateUrl = "/api/v1/aPIStatus"; // + id

const aPIStatus = {
    actions: {
        createAPIStatus(context, {data, filter, header}) {

            return api.create(createUrl, data, filter, header)
                .then(function(response) {

                    context.commit("setAPIStatus", response.Model);

                    return response;
                })
                .catch(function(err) {
                    return err;
                });
        },
        deleteAPIStatus(context, {id, header}) {

            return api.remove(deleteUrl + id, header)
                .then(function(response) {
                    context.commit("clearAPIStatus");
                    return response;
                })
                .catch(function(err) {
                    return err;
                });
        },
        findAPIStatus(context, {filter, header}) {

            return api.find(findUrl, filter, header)
                .then(function(response) {

                    context.commit("setAPIStatusList", response.List);

                    return response;
                })
                .catch(function(err) {
                    return err;
                });
        },
        loadAPIStatus(context, {id, filter, header}) {

            return api.find(readUrl + id, filter, header)
                .then(function(response) {

                    context.commit("setAPIStatus", response.Model);

                    return response;
                })
                .catch(function(err) {
                    return err;
                });
        },
        updateAPIStatus(context, {id, data, filter, header}) {

            return api.update(updateUrl + id, data, filter, header)
                .then(function(response) {

                    context.commit("setAPIStatus", response.Model);

                    return response;
                })
                .catch(function(err) {
                    return err;
                });
        },
        findOrCreateAPIStatus(context, {id, data, filter, header}) {

            return api.update(findOrCreateUrl, data, filter, header)
                .then(function(response) {

                    context.commit("setAPIStatus", response.Model);

                    return response;
                })
                .catch(function(err) {
                    return err;
                });
        },
        clearListAPIStatus(context) {
            context.commit("clearListAPIStatus");
        },
    },
    getters: {
        getAPIStatus: (state) => {
            return state.APIStatus;
        },
        getListAPIStatus: (state) => {
            return state.APIStatusList;
        },
    },
    mutations: {
        setAPIStatus(state, data) {
            state.APIStatus = data;
        },
        setAPIStatusList(state, data) {
            state.APIStatusList = data;
        },
        clearAPIStatus(state) {
            state.APIStatus = new APIStatus();
        },
        clearListAPIStatus(state) {
            state.APIStatusList = [];
        },
		updateAPIStatusById(state, data) {
    		let index = findItemIndex(state.APIStatusList, function(item) {
	        	return item.Id === data.Id;
	    	});
	    
	    	if (index || index === 0) {
		        state.APIStatusList.splice(index, 1, data);
    		}
		},
		deleteAPIStatusFromList(state, id) {
		    let index = findItemIndex(state.APIStatusList, function(item) {
		        return item.Id === id;
		    });
		    
		    if (index || index === 0) {
		        state.APIStatusList.splice(index, 1);
		    }
		},
		addAPIStatusItemToList(state, item) {

			if (state.APIStatusList === null) {
				state.APIStatusList = [];
			}

		    state.APIStatusList.push(item);
		},
    },
    state: {
        APIStatus: new APIStatus(),
        APIStatusList: [],
    },
};

export default aPIStatus;
