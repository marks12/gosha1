
import {RegionType} from "../apiModel";
import api from "../api";
import {findItemIndex} from "../common";

let findUrl = "/api/v1/regionType";
let readUrl = "/api/v1/regionType/"; // + id
let createUrl = "/api/v1/regionType";
let multiCreateUrl = "/api/v1/regionType/list";
let updateUrl = "/api/v1/regionType/"; // + id
let multiUpdateUrl = "/api/v1/regionType/list"; // + id
let deleteUrl = "/api/v1/regionType/"; // + id
let multiDeleteUrl = "/api/v1/regionType/list"; // + id
let findOrCreateUrl = "/api/v1/regionType"; // + id

const regionType = {
    actions: {
        createRegionType(context, {data, filter, header}) {

            let url = createUrl;
            if (Array.isArray && Array.isArray(data)) {
                url = multiCreateUrl
            }

            return api.create(url, data, filter, header)
                .then(function(response) {

                    context.commit("setRegionType", response.Model);

                    return response;
                })
                .catch(function(err) {
                    console.error(err);
                    throw(err);
                });
        },
        deleteRegionType(context, {id, header}) {

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
                    context.commit("clearRegionType");
                    return response;
                })
                .catch(function(err) {
                    console.error(err);
                    throw(err);
                });
        },
        findRegionType(context, {filter, header, isAppend}) {

            return api.find(findUrl, filter, header)
                .then(function(response) {

                    if (isAppend) {
                        context.commit("appendRegionType__List", response.List);
                    } else {
                        context.commit("setRegionType__List", response.List);
                    }

                    return response;
                })
                .catch(function(err) {
                    console.error(err);
                    throw(err);
                });
        },
        loadRegionType(context, {id, filter, header}) {

            return api.find(readUrl + id, filter, header)
                .then(function(response) {

                    context.commit("setRegionType", response.Model);
                    return response;
                })
                .catch(function(err) {
                    console.error(err);
                    throw(err);
                });
        },
        updateRegionType(context, {id, data, filter, header}) {

            let url = updateUrl + id;
            if (Array.isArray && Array.isArray(data)) {
                url = multiUpdateUrl
            }

            return api.update(url, data, filter, header)
                .then(function(response) {

                    context.commit("setRegionType", response.Model);
                    return response;
                })
                .catch(function(err) {
                    console.error(err);
                    throw(err);
                });
        },
        findOrCreateRegionType(context, {id, data, filter, header}) {

            return api.update(findOrCreateUrl, data, filter, header)
                .then(function(response) {

                    context.commit("setRegionType", response.Model);
                    return response;
                })
                .catch(function(err) {
                    console.error(err);
                    throw(err);
                });
        },
        clearListRegionType(context) {
            context.commit("clearListRegionType");
        },
        clearRegionType(context) {
            context.commit("clearRegionType");
        },
    },
    getters: {
        getRegionType: (state) => {
            return state.RegionType;
        },
        getRegionTypeById: state => id => {
            return state.RegionType__List.find(item => item.Id === id);
        },
        getListRegionType: (state) => {
            return state.RegionType__List;
        },
    },
    mutations: {
        setRegionType(state, data) {
            state.RegionType = data;
        },
        setRegionType__List(state, data) {
            state.RegionType__List = data || [];
        },
        appendRegionType__List(state, data) {

            if (! state.RegionType__List) {
                state.RegionType__List = [];
            }

            state.RegionType__List = state.RegionType__List.concat(data);
        },
        clearRegionType(state) {
            state.RegionType = new RegionType();
        },
        clearListRegionType(state) {
            state.RegionType__List = [];
        },
		updateRegionTypeById(state, data) {
    		let index = findItemIndex(state.RegionType__List, function(item) {
	        	return item.Id === data.Id;
	    	});
	    
	    	if (index || index === 0) {
		        state.RegionType__List.splice(index, 1, data);
    		}
		},
		deleteRegionTypeFromList(state, id) {
		    let index = findItemIndex(state.RegionType__List, function(item) {
		        return item.Id === id;
		    });
		    
		    if (index || index === 0) {
		        state.RegionType__List.splice(index, 1);
		    }
		},
		addRegionTypeItemToList(state, item) {

			if (state.RegionType__List === null) {
				state.RegionType__List = [];
			}

		    state.RegionType__List.push(item);
		},
    },
    state: {
        RegionType: new RegionType(),
        RegionType__List: [],
    },
};

export default regionType;
