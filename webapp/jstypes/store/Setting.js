
import {Setting} from "../apiModel";
import api from "../api";
import {findItemIndex} from "../common";

let findUrl = "/api/v1/setting";
let readUrl = "/api/v1/setting/"; // + id
let createUrl = "/api/v1/setting";
let multiCreateUrl = "/api/v1/setting/list";
let updateUrl = "/api/v1/setting/"; // + id
let multiUpdateUrl = "/api/v1/setting/list";
let deleteUrl = "/api/v1/setting/"; // + id
let multiDeleteUrl = "/api/v1/setting/list";
let findOrCreateUrl = "/api/v1/setting";
let updateOrCreateUrl = "/api/v1/setting";

const setting = {
    actions: {
        createSetting(context, {data, filter, header, noMutation}) {

            let url = createUrl;
            if (Array.isArray && Array.isArray(data)) {
                url = multiCreateUrl
            }

            return api.create(url, data, filter, header)
                .then(function(response) {

					if(! noMutation) {
	                    context.commit("setSetting", response.Model);
					}

                    return response;
                })
                .catch(function(err) {
                    console.error(err);
                    throw(err);
                });
        },
        deleteSetting(context, {id, header, noMutation}) {

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
	                    context.commit("clearSetting");
					}
                    return response;
                })
                .catch(function(err) {
                    console.error(err);
                    throw(err);
                });
        },
        findSetting(context, {filter, header, isAppend, noMutation}) {

            return api.find(findUrl, filter, header)
                .then(function(response) {

					if(! noMutation) {
						if (isAppend) {
							context.commit("appendSetting__List", response.List);
						} else {
							context.commit("setSetting__List", response.List);
						}
					}

                    return response;
                })
                .catch(function(err) {
                    console.error(err);
                    throw(err);
                });
        },
        loadSetting(context, {id, filter, header, noMutation}) {

            return api.find(readUrl + id, filter, header)
                .then(function(response) {

					if(! noMutation) {
	                    context.commit("setSetting", response.Model);
					}
                    return response;
                })
                .catch(function(err) {
                    console.error(err);
                    throw(err);
                });
        },
        updateSetting(context, {id, data, filter, header, noMutation}) {

            let url = updateUrl + id;
            if (Array.isArray && Array.isArray(data)) {
                url = multiUpdateUrl
            }

            return api.update(url, data, filter, header)
                .then(function(response) {
					if(! noMutation) {
	                    context.commit("setSetting", response.Model);
					}
                    return response;
                })
                .catch(function(err) {
                    console.error(err);
                    throw(err);
                });
        },
        findOrCreateSetting(context, {id, data, filter, header, noMutation}) {

            return api.update(findOrCreateUrl, data, filter, header)
                .then(function(response) {

					if(! noMutation) {
	                    context.commit("setSetting", response.Model);
					}
                    return response;
                })
                .catch(function(err) {
                    console.error(err);
                    throw(err);
                });
        },
        clearListSetting(context) {
            context.commit("clearListSetting");
        },
        clearSetting(context) {
            context.commit("clearSetting");
        },
    },
    getters: {
        getSetting: (state) => {
            return state.Setting;
        },
        getSettingById: state => id => {
            return state.Setting__List.find(item => item.Id === id);
        },
        getListSetting: (state) => {
            return state.Setting__List;
        },
        getRoute__Setting: state => action => {
            return state.Setting__Routes[action];
        },
        getRoutes__Setting: state => {
            return state.Setting__Routes;
        },
    },
    mutations: {
        setSetting(state, data) {
            state.Setting = data;
        },
        setSetting__List(state, data) {
            state.Setting__List = data || [];
        },
        appendSetting__List(state, data) {

            if (! state.Setting__List) {
                state.Setting__List = [];
            }

			if (data !== null) {
				state.Setting__List = state.Setting__List.concat(data);				
			}
        },
        clearSetting(state) {
            state.Setting = new Setting();
        },
        clearListSetting(state) {
            state.Setting__List = [];
        },
		updateSettingById(state, data) {
    		let index = findItemIndex(state.Setting__List, function(item) {
	        	return item.Id === data.Id;
	    	});
	    
	    	if (index || index === 0) {
		        state.Setting__List.splice(index, 1, data);
    		}
		},
		deleteSettingFromList(state, id) {
		    let index = findItemIndex(state.Setting__List, function(item) {
		        return item.Id === id;
		    });
		    
		    if (index || index === 0) {
		        state.Setting__List.splice(index, 1);
		    }
		},
		addSettingItemToList(state, item) {

			if (state.Setting__List === null) {
				state.Setting__List = [];
			}

		    state.Setting__List.push(item);
		},
    },
    state: {
        Setting: new Setting(),
        Setting__List: [],
        Setting__Routes: {
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

export default setting;
