package cmd

const FindUrl = "/api/v1/{entity}"
const ReadUrl = "/api/v1/{entity}/" // + id
const CreateUrl = "/api/v1/{entity}"
const MultiCreateUrl = "/api/v1/{entity}/list"
const UpdateUrl = "/api/v1/{entity}/" // + id
const MultiUpdateUrl = "/api/v1/{entity}/list"
const DeleteUrl = "/api/v1/{entity}/" // + id
const MultiDeleteUrl = "/api/v1/{entity}/list"
const FindOrCreateUrl = "/api/v1/{entity}"
const UpdateOrCreateUrl = "/api/v1/{entity}"

const storeTemplate = `
import {{Entity}} from "../apiModel";
import api from "../api";
import {findItemIndex} from "../common";

let findUrl = "` + FindUrl + `";
let readUrl = "` + ReadUrl + `"; // + id
let createUrl = "` + CreateUrl + `";
let multiCreateUrl = "` + MultiCreateUrl + `";
let updateUrl = "` + UpdateUrl + `"; // + id
let multiUpdateUrl = "` + MultiUpdateUrl + `";
let deleteUrl = "` + DeleteUrl + `"; // + id
let multiDeleteUrl = "` + MultiDeleteUrl + `";
let findOrCreateUrl = "` + FindOrCreateUrl + `";
let findOrCreateUrl = "` + UpdateOrCreateUrl + `";

const {entity} = {
    actions: {
        create{Entity}(context, {data, filter, header}) {

            let url = createUrl;
            if (Array.isArray && Array.isArray(data)) {
                url = multiCreateUrl
            }

            return api.create(url, data, filter, header)
                .then(function(response) {

                    context.commit("set{Entity}", response.Model);

                    return response;
                })
                .catch(function(err) {
                    console.error(err);
                    throw(err);
                });
        },
        delete{Entity}(context, {id, header}) {

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
                    context.commit("clear{Entity}");
                    return response;
                })
                .catch(function(err) {
                    console.error(err);
                    throw(err);
                });
        },
        find{Entity}(context, {filter, header, isAppend}) {

            return api.find(findUrl, filter, header)
                .then(function(response) {

                    if (isAppend) {
                        context.commit("append{Entity}__List", response.List);
                    } else {
                        context.commit("set{Entity}__List", response.List);
                    }

                    return response;
                })
                .catch(function(err) {
                    console.error(err);
                    throw(err);
                });
        },
        load{Entity}(context, {id, filter, header}) {

            return api.find(readUrl + id, filter, header)
                .then(function(response) {

                    context.commit("set{Entity}", response.Model);
                    return response;
                })
                .catch(function(err) {
                    console.error(err);
                    throw(err);
                });
        },
        update{Entity}(context, {id, data, filter, header}) {

            let url = updateUrl + id;
            if (Array.isArray && Array.isArray(data)) {
                url = multiUpdateUrl
            }

            return api.update(url, data, filter, header)
                .then(function(response) {

                    context.commit("set{Entity}", response.Model);
                    return response;
                })
                .catch(function(err) {
                    console.error(err);
                    throw(err);
                });
        },
        findOrCreate{Entity}(context, {id, data, filter, header}) {

            return api.update(findOrCreateUrl, data, filter, header)
                .then(function(response) {

                    context.commit("set{Entity}", response.Model);
                    return response;
                })
                .catch(function(err) {
                    console.error(err);
                    throw(err);
                });
        },
        clearList{Entity}(context) {
            context.commit("clearList{Entity}");
        },
        clear{Entity}(context) {
            context.commit("clear{Entity}");
        },
    },
    getters: {
        get{Entity}: (state) => {
            return state.{Entity};
        },
        get{Entity}ById: state => id => {
            return state.{Entity}__List.find(item => item.Id === id);
        },
        getList{Entity}: (state) => {
            return state.{Entity}__List;
        },
        getRoute__{Entity}: state => action => {
            return state.{Entity}__Routes[action];
        },
        getRoutes__{Entity}: state => {
            return state.{Entity}__Routes;
        },
    },
    mutations: {
        set{Entity}(state, data) {
            state.{Entity} = data;
        },
        set{Entity}__List(state, data) {
            state.{Entity}__List = data || [];
        },
        append{Entity}__List(state, data) {

            if (! state.{Entity}__List) {
                state.{Entity}__List = [];
            }

            state.{Entity}__List = state.{Entity}__List.concat(data);
        },
        clear{Entity}(state) {
            state.{Entity} = new {Entity}();
        },
        clearList{Entity}(state) {
            state.{Entity}__List = [];
        },
		update{Entity}ById(state, data) {
    		let index = findItemIndex(state.{Entity}__List, function(item) {
	        	return item.Id === data.Id;
	    	});
	    
	    	if (index || index === 0) {
		        state.{Entity}__List.splice(index, 1, data);
    		}
		},
		delete{Entity}FromList(state, id) {
		    let index = findItemIndex(state.{Entity}__List, function(item) {
		        return item.Id === id;
		    });
		    
		    if (index || index === 0) {
		        state.{Entity}__List.splice(index, 1);
		    }
		},
		add{Entity}ItemToList(state, item) {

			if (state.{Entity}__List === null) {
				state.{Entity}__List = [];
			}

		    state.{Entity}__List.push(item);
		},
    },
    state: {
        {Entity}: new {Entity}(),
        {Entity}__List: [],
        {Entity}__Routes: {
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

export default {entity};
`
