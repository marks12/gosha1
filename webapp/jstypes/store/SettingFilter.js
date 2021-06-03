
import {SettingFilter} from "../apiModel";
import api from "../api";
import {findItemIndex} from "../common";

let findUrl = "/api/v1/settingFilter";
let readUrl = "/api/v1/settingFilter/"; // + id
let createUrl = "/api/v1/settingFilter";
let multiCreateUrl = "/api/v1/settingFilter/list";
let updateUrl = "/api/v1/settingFilter/"; // + id
let multiUpdateUrl = "/api/v1/settingFilter/list";
let deleteUrl = "/api/v1/settingFilter/"; // + id
let multiDeleteUrl = "/api/v1/settingFilter/list";
let findOrCreateUrl = "/api/v1/settingFilter";
let updateOrCreateUrl = "/api/v1/settingFilter";

const settingFilter = {
    actions: {
        createSettingFilter(context, {data, filter, header, noMutation}) {

            let url = createUrl;
            if (Array.isArray && Array.isArray(data)) {
                url = multiCreateUrl
            }

            return api.create(url, data, filter, header)
                .then(function(response) {

					if(! noMutation) {
	                    context.commit("setSettingFilter", response.Model);
					}

                    return response;
                })
                .catch(function(err) {
                    console.error(err);
                    throw(err);
                });
        },
        deleteSettingFilter(context, {id, header, noMutation}) {

            let url;
            let dataOrNull = null;

            if (Array.isArray && Array.isArray(id)) {
                url = multiDeleteUrl;
                dataOrNull = id.map(item => typeof item === "number" ? {Id: item} : item);
            } else {
                url = deleteUrl + id;
            }

            return api.remove(url, header, dataOrNull)
                .then(function(response) {
					if(! noMutation) {
	                    context.commit("clearSettingFilter");
					}
                    return response;
                })
                .catch(function(err) {
                    console.error(err);
                    throw(err);
                });
        },
        findSettingFilter(context, {filter, header, isAppend, noMutation}) {

            return api.find(findUrl, filter, header)
                .then(function(response) {

					if(! noMutation) {
						if (isAppend) {
							context.commit("appendSettingFilter__List", response.List);
						} else {
							context.commit("setSettingFilter__List", response.List);
						}
					}

                    return response;
                })
                .catch(function(err) {
                    console.error(err);
                    throw(err);
                });
        },
        loadSettingFilter(context, {id, filter, header, noMutation}) {

            return api.find(readUrl + id, filter, header)
                .then(function(response) {

					if(! noMutation) {
	                    context.commit("setSettingFilter", response.Model);
					}
                    return response;
                })
                .catch(function(err) {
                    console.error(err);
                    throw(err);
                });
        },
        updateSettingFilter(context, {id, data, filter, header, noMutation}) {

            let url = updateUrl + id;
            if (Array.isArray && Array.isArray(data)) {
                url = multiUpdateUrl
            }

            return api.update(url, data, filter, header)
                .then(function(response) {
					if(! noMutation) {
	                    context.commit("setSettingFilter", response.Model);
					}
                    return response;
                })
                .catch(function(err) {
                    console.error(err);
                    throw(err);
                });
        },
        findOrCreateSettingFilter(context, {id, data, filter, header, noMutation}) {

            return api.update(findOrCreateUrl, data, filter, header)
                .then(function(response) {

					if(! noMutation) {
	                    context.commit("setSettingFilter", response.Model);
					}
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
        clearSettingFilter(context) {
            context.commit("clearSettingFilter");
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
        getRoute__SettingFilter: state => action => {
            return state.SettingFilter__Routes[action];
        },
        getRoutes__SettingFilter: state => {
            return state.SettingFilter__Routes;
        },
    },
    mutations: {
        setSettingFilter(state, data) {
            state.SettingFilter = data;
        },
        setSettingFilter__List(state, data) {
            state.SettingFilter__List = data || [];
        },
        appendSettingFilter__List(state, data) {

            if (! state.SettingFilter__List) {
                state.SettingFilter__List = [];
            }

			if (data !== null) {
				state.SettingFilter__List = state.SettingFilter__List.concat(data);				
			}
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
        SettingFilter__Routes: {
            find: findUrl,
            read: readUrl,
            create: createUrl,
            multiCreate: multiCreateUrl,
            update: updateUrl,
            multiUpdate: multiUpdateUrl,
            delete: deleteUrl,
            multiDelete: multiDeleteUrl,
            findOrCreate: findOrCreateUrl,
            updateOrCreate: updateOrCreateUrl,
        },
    },
};

export default settingFilter;
