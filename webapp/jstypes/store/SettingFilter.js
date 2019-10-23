
import {SettingFilter} from "../apiModel";
import api from "../api";
import {findItemIndex} from "../common";

let findUrl = "/api/v1/settingFilter";
let readUrl = "/api/v1/settingFilter/"; // + id
let createUrl = "/api/v1/settingFilter";
let updateUrl = "/api/v1/settingFilter/"; // + id
let deleteUrl = "/api/v1/settingFilter/"; // + id
let findOrCreateUrl = "/api/v1/settingFilter"; // + id

const settingFilter = {
    actions: {
        createSettingFilter(context, {data, filter, header}) {

            return api.create(createUrl, data, filter, header)
                .then(function(response) {

                    context.commit("setSettingFilter", response.Model);

                    return response;
                })
                .catch(function(err) {
                    console.error(err);
                    throw(err);
                });
        },
        deleteSettingFilter(context, {id, header}) {

            return api.remove(deleteUrl + id, header)
                .then(function(response) {
                    context.commit("clearSettingFilter");
                    return response;
                })
                .catch(function(err) {
                    console.error(err);
                    throw(err);
                });
        },
        findSettingFilter(context, {filter, header}) {

            return api.find(findUrl, filter, header)
                .then(function(response) {

                    context.commit("setSettingFilter__List", response.List);
                    return response;
                })
                .catch(function(err) {
                    console.error(err);
                    throw(err);
                });
        },
        loadSettingFilter(context, {id, filter, header}) {

            return api.find(readUrl + id, filter, header)
                .then(function(response) {

                    context.commit("setSettingFilter", response.Model);
                    return response;
                })
                .catch(function(err) {
                    console.error(err);
                    throw(err);
                });
        },
        updateSettingFilter(context, {id, data, filter, header}) {

            return api.update(updateUrl + id, data, filter, header)
                .then(function(response) {

                    context.commit("setSettingFilter", response.Model);
                    return response;
                })
                .catch(function(err) {
                    console.error(err);
                    throw(err);
                });
        },
        findOrCreateSettingFilter(context, {id, data, filter, header}) {

            return api.update(findOrCreateUrl, data, filter, header)
                .then(function(response) {

                    context.commit("setSettingFilter", response.Model);
                    return response;
                })
                .catch(function(err) {
                    console.error(err);
                    throw(err);
                });
        },
        clearListSettingFilter(context) {
            context.commit("clearListSettingFilter");
        },
    },
    getters: {
        getSettingFilter: (state) => {
            return state.SettingFilter;
        },
        getSettingFilterById: state => id => {
            return state.SettingFilter__List.find(item => item.Id === id);
        },
        getListSettingFilter: (state) => {
            return state.SettingFilter__List;
        },
    },
    mutations: {
        setSettingFilter(state, data) {
            state.SettingFilter = data;
        },
        setSettingFilter__List(state, data) {
            state.SettingFilter__List = data || [];
        },
        clearSettingFilter(state) {
            state.SettingFilter = new SettingFilter();
        },
        clearListSettingFilter(state) {
            state.SettingFilter__List = [];
        },
		updateSettingFilterById(state, data) {
    		let index = findItemIndex(state.SettingFilter__List, function(item) {
	        	return item.Id === data.Id;
	    	});
	    
	    	if (index || index === 0) {
		        state.SettingFilter__List.splice(index, 1, data);
    		}
		},
		deleteSettingFilterFromList(state, id) {
		    let index = findItemIndex(state.SettingFilter__List, function(item) {
		        return item.Id === id;
		    });
		    
		    if (index || index === 0) {
		        state.SettingFilter__List.splice(index, 1);
		    }
		},
		addSettingFilterItemToList(state, item) {

			if (state.SettingFilter__List === null) {
				state.SettingFilter__List = [];
			}

		    state.SettingFilter__List.push(item);
		},
    },
    state: {
        SettingFilter: new SettingFilter(),
        SettingFilter__List: [],
    },
};

export default settingFilter;
