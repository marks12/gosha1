
import {Region} from "../apiModel";
import api from "../api";
import {findItemIndex} from "../common";

let findUrl = "/api/v1/region";
let readUrl = "/api/v1/region/"; // + id
let createUrl = "/api/v1/region";
let multiCreateUrl = "/api/v1/region/list";
let updateUrl = "/api/v1/region/"; // + id
let multiUpdateUrl = "/api/v1/region/list"; // + id
let deleteUrl = "/api/v1/region/"; // + id
let multiDeleteUrl = "/api/v1/region/list"; // + id
let findOrCreateUrl = "/api/v1/region"; // + id

const region = {
    actions: {
        createRegion(context, {data, filter, header}) {

            let url = createUrl;
            if (Array.isArray && Array.isArray(data)) {
                url = multiCreateUrl
            }

            return api.create(url, data, filter, header)
                .then(function(response) {

                    context.commit("setRegion", response.Model);

                    return response;
                })
                .catch(function(err) {
                    console.error(err);
                    throw(err);
                });
        },
        deleteRegion(context, {id, header}) {

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
                    context.commit("clearRegion");
                    return response;
                })
                .catch(function(err) {
                    console.error(err);
                    throw(err);
                });
        },
        findRegion(context, {filter, header}) {

            return api.find(findUrl, filter, header)
                .then(function(response) {

                    context.commit("setRegion__List", response.List);
                    return response;
                })
                .catch(function(err) {
                    console.error(err);
                    throw(err);
                });
        },
        loadRegion(context, {id, filter, header}) {

            return api.find(readUrl + id, filter, header)
                .then(function(response) {

                    context.commit("setRegion", response.Model);
                    return response;
                })
                .catch(function(err) {
                    console.error(err);
                    throw(err);
                });
        },
        updateRegion(context, {id, data, filter, header}) {

            let url = updateUrl + id;
            if (Array.isArray && Array.isArray(data)) {
                url = multiUpdateUrl
            }

            return api.update(url, data, filter, header)
                .then(function(response) {

                    context.commit("setRegion", response.Model);
                    return response;
                })
                .catch(function(err) {
                    console.error(err);
                    throw(err);
                });
        },
        findOrCreateRegion(context, {id, data, filter, header}) {

            return api.update(findOrCreateUrl, data, filter, header)
                .then(function(response) {

                    context.commit("setRegion", response.Model);
                    return response;
                })
                .catch(function(err) {
                    console.error(err);
                    throw(err);
                });
        },
        clearListRegion(context) {
            context.commit("clearListRegion");
        },
        clearRegion(context) {
            context.commit("clearRegion");
        },
    },
    getters: {
        getRegion: (state) => {
            return state.Region;
        },
        getRegionById: state => id => {
            return state.Region__List.find(item => item.Id === id);
        },
        getListRegion: (state) => {
            return state.Region__List;
        },
    },
    mutations: {
        setRegion(state, data) {
            state.Region = data;
        },
        setRegion__List(state, data) {
            state.Region__List = data || [];
        },
        clearRegion(state) {
            state.Region = new Region();
        },
        clearListRegion(state) {
            state.Region__List = [];
        },
		updateRegionById(state, data) {
    		let index = findItemIndex(state.Region__List, function(item) {
	        	return item.Id === data.Id;
	    	});
	    
	    	if (index || index === 0) {
		        state.Region__List.splice(index, 1, data);
    		}
		},
		deleteRegionFromList(state, id) {
		    let index = findItemIndex(state.Region__List, function(item) {
		        return item.Id === id;
		    });
		    
		    if (index || index === 0) {
		        state.Region__List.splice(index, 1);
		    }
		},
		addRegionItemToList(state, item) {

			if (state.Region__List === null) {
				state.Region__List = [];
			}

		    state.Region__List.push(item);
		},
    },
    state: {
        Region: new Region(),
        Region__List: [],
    },
};

export default region;
