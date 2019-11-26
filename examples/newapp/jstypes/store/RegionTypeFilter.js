
import {RegionTypeFilter} from "../apiModel";
import api from "../api";
import {findItemIndex} from "../common";

let findUrl = "/api/v1/regionTypeFilter";
let readUrl = "/api/v1/regionTypeFilter/"; // + id
let createUrl = "/api/v1/regionTypeFilter";
let multiCreateUrl = "/api/v1/regionTypeFilter/list";
let updateUrl = "/api/v1/regionTypeFilter/"; // + id
let multiUpdateUrl = "/api/v1/regionTypeFilter/list"; // + id
let deleteUrl = "/api/v1/regionTypeFilter/"; // + id
let multiDeleteUrl = "/api/v1/regionTypeFilter/list"; // + id
let findOrCreateUrl = "/api/v1/regionTypeFilter"; // + id

const regionTypeFilter = {
    actions: {
        createRegionTypeFilter(context, {data, filter, header}) {

            let url = createUrl;
            if (Array.isArray && Array.isArray(data)) {
                url = multiCreateUrl
            }

            return api.create(url, data, filter, header)
                .then(function(response) {

                    context.commit("setRegionTypeFilter", response.Model);

                    return response;
                })
                .catch(function(err) {
                    console.error(err);
                    throw(err);
                });
        },
        deleteRegionTypeFilter(context, {id, header}) {

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
                    context.commit("clearRegionTypeFilter");
                    return response;
                })
                .catch(function(err) {
                    console.error(err);
                    throw(err);
                });
        },
        findRegionTypeFilter(context, {filter, header}) {

            return api.find(findUrl, filter, header)
                .then(function(response) {

                    context.commit("setRegionTypeFilter__List", response.List);
                    return response;
                })
                .catch(function(err) {
                    console.error(err);
                    throw(err);
                });
        },
        loadRegionTypeFilter(context, {id, filter, header}) {

            return api.find(readUrl + id, filter, header)
                .then(function(response) {

                    context.commit("setRegionTypeFilter", response.Model);
                    return response;
                })
                .catch(function(err) {
                    console.error(err);
                    throw(err);
                });
        },
        updateRegionTypeFilter(context, {id, data, filter, header}) {

            let url = updateUrl + id;
            if (Array.isArray && Array.isArray(data)) {
                url = multiUpdateUrl
            }

            return api.update(url, data, filter, header)
                .then(function(response) {

                    context.commit("setRegionTypeFilter", response.Model);
                    return response;
                })
                .catch(function(err) {
                    console.error(err);
                    throw(err);
                });
        },
        findOrCreateRegionTypeFilter(context, {id, data, filter, header}) {

            return api.update(findOrCreateUrl, data, filter, header)
                .then(function(response) {

                    context.commit("setRegionTypeFilter", response.Model);
                    return response;
                })
                .catch(function(err) {
                    console.error(err);
                    throw(err);
                });
        },
        clearListRegionTypeFilter(context) {
            context.commit("clearListRegionTypeFilter");
        },
        clearRegionTypeFilter(context) {
            context.commit("clearRegionTypeFilter");
        },
    },
    getters: {
        getRegionTypeFilter: (state) => {
            return state.RegionTypeFilter;
        },
        getRegionTypeFilterById: state => id => {
            return state.RegionTypeFilter__List.find(item => item.Id === id);
        },
        getListRegionTypeFilter: (state) => {
            return state.RegionTypeFilter__List;
        },
    },
    mutations: {
        setRegionTypeFilter(state, data) {
            state.RegionTypeFilter = data;
        },
        setRegionTypeFilter__List(state, data) {
            state.RegionTypeFilter__List = data || [];
        },
        clearRegionTypeFilter(state) {
            state.RegionTypeFilter = new RegionTypeFilter();
        },
        clearListRegionTypeFilter(state) {
            state.RegionTypeFilter__List = [];
        },
		updateRegionTypeFilterById(state, data) {
    		let index = findItemIndex(state.RegionTypeFilter__List, function(item) {
	        	return item.Id === data.Id;
	    	});
	    
	    	if (index || index === 0) {
		        state.RegionTypeFilter__List.splice(index, 1, data);
    		}
		},
		deleteRegionTypeFilterFromList(state, id) {
		    let index = findItemIndex(state.RegionTypeFilter__List, function(item) {
		        return item.Id === id;
		    });
		    
		    if (index || index === 0) {
		        state.RegionTypeFilter__List.splice(index, 1);
		    }
		},
		addRegionTypeFilterItemToList(state, item) {

			if (state.RegionTypeFilter__List === null) {
				state.RegionTypeFilter__List = [];
			}

		    state.RegionTypeFilter__List.push(item);
		},
    },
    state: {
        RegionTypeFilter: new RegionTypeFilter(),
        RegionTypeFilter__List: [],
    },
};

export default regionTypeFilter;
