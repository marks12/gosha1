
import {RegionFilter} from "../apiModel";
import api from "../api";
import {findItemIndex} from "../common";

let findUrl = "/api/v1/regionFilter";
let readUrl = "/api/v1/regionFilter/"; // + id
let createUrl = "/api/v1/regionFilter";
let multiCreateUrl = "/api/v1/regionFilter/list";
let updateUrl = "/api/v1/regionFilter/"; // + id
let multiUpdateUrl = "/api/v1/regionFilter/list"; // + id
let deleteUrl = "/api/v1/regionFilter/"; // + id
let multiDeleteUrl = "/api/v1/regionFilter/list"; // + id
let findOrCreateUrl = "/api/v1/regionFilter"; // + id

const regionFilter = {
    actions: {
        createRegionFilter(context, {data, filter, header}) {

            let url = createUrl;
            if (Array.isArray && Array.isArray(data)) {
                url = multiCreateUrl
            }

            return api.create(url, data, filter, header)
                .then(function(response) {

                    context.commit("setRegionFilter", response.Model);

                    return response;
                })
                .catch(function(err) {
                    console.error(err);
                    throw(err);
                });
        },
        deleteRegionFilter(context, {id, header}) {

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
                    context.commit("clearRegionFilter");
                    return response;
                })
                .catch(function(err) {
                    console.error(err);
                    throw(err);
                });
        },
        findRegionFilter(context, {filter, header, isAppend}) {

            return api.find(findUrl, filter, header)
                .then(function(response) {

                    if (isAppend) {
                        context.commit("appendRegionFilter__List", response.List);
                    } else {
                        context.commit("setRegionFilter__List", response.List);
                    }

                    return response;
                })
                .catch(function(err) {
                    console.error(err);
                    throw(err);
                });
        },
        loadRegionFilter(context, {id, filter, header}) {

            return api.find(readUrl + id, filter, header)
                .then(function(response) {

                    context.commit("setRegionFilter", response.Model);
                    return response;
                })
                .catch(function(err) {
                    console.error(err);
                    throw(err);
                });
        },
        updateRegionFilter(context, {id, data, filter, header}) {

            let url = updateUrl + id;
            if (Array.isArray && Array.isArray(data)) {
                url = multiUpdateUrl
            }

            return api.update(url, data, filter, header)
                .then(function(response) {

                    context.commit("setRegionFilter", response.Model);
                    return response;
                })
                .catch(function(err) {
                    console.error(err);
                    throw(err);
                });
        },
        findOrCreateRegionFilter(context, {id, data, filter, header}) {

            return api.update(findOrCreateUrl, data, filter, header)
                .then(function(response) {

                    context.commit("setRegionFilter", response.Model);
                    return response;
                })
                .catch(function(err) {
                    console.error(err);
                    throw(err);
                });
        },
        clearListRegionFilter(context) {
            context.commit("clearListRegionFilter");
        },
        clearRegionFilter(context) {
            context.commit("clearRegionFilter");
        },
    },
    getters: {
        getRegionFilter: (state) => {
            return state.RegionFilter;
        },
        getRegionFilterById: state => id => {
            return state.RegionFilter__List.find(item => item.Id === id);
        },
        getListRegionFilter: (state) => {
            return state.RegionFilter__List;
        },
    },
    mutations: {
        setRegionFilter(state, data) {
            state.RegionFilter = data;
        },
        setRegionFilter__List(state, data) {
            state.RegionFilter__List = data || [];
        },
        appendRegionFilter__List(state, data) {

            if (! state.RegionFilter__List) {
                state.RegionFilter__List = [];
            }

            state.RegionFilter__List = state.RegionFilter__List.concat(data);
        },
        clearRegionFilter(state) {
            state.RegionFilter = new RegionFilter();
        },
        clearListRegionFilter(state) {
            state.RegionFilter__List = [];
        },
		updateRegionFilterById(state, data) {
    		let index = findItemIndex(state.RegionFilter__List, function(item) {
	        	return item.Id === data.Id;
	    	});
	    
	    	if (index || index === 0) {
		        state.RegionFilter__List.splice(index, 1, data);
    		}
		},
		deleteRegionFilterFromList(state, id) {
		    let index = findItemIndex(state.RegionFilter__List, function(item) {
		        return item.Id === id;
		    });
		    
		    if (index || index === 0) {
		        state.RegionFilter__List.splice(index, 1);
		    }
		},
		addRegionFilterItemToList(state, item) {

			if (state.RegionFilter__List === null) {
				state.RegionFilter__List = [];
			}

		    state.RegionFilter__List.push(item);
		},
    },
    state: {
        RegionFilter: new RegionFilter(),
        RegionFilter__List: [],
    },
};

export default regionFilter;
