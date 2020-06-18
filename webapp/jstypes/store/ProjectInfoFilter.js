
import {ProjectInfoFilter} from "../apiModel";
import api from "../api";
import {findItemIndex} from "../common";

let findUrl = "/api/v1/projectInfoFilter";
let readUrl = "/api/v1/projectInfoFilter/"; // + id
let createUrl = "/api/v1/projectInfoFilter";
let multiCreateUrl = "/api/v1/projectInfoFilter/list";
let updateUrl = "/api/v1/projectInfoFilter/"; // + id
let multiUpdateUrl = "/api/v1/projectInfoFilter/list";
let deleteUrl = "/api/v1/projectInfoFilter/"; // + id
let multiDeleteUrl = "/api/v1/projectInfoFilter/list";
let findOrCreateUrl = "/api/v1/projectInfoFilter";
let updateOrCreateUrl = "/api/v1/projectInfoFilter";

const projectInfoFilter = {
    actions: {
        createProjectInfoFilter(context, {data, filter, header, noMutation}) {

            let url = createUrl;
            if (Array.isArray && Array.isArray(data)) {
                url = multiCreateUrl
            }

            return api.create(url, data, filter, header)
                .then(function(response) {

					if(! noMutation) {
	                    context.commit("setProjectInfoFilter", response.Model);
					}

                    return response;
                })
                .catch(function(err) {
                    console.error(err);
                    throw(err);
                });
        },
        deleteProjectInfoFilter(context, {id, header, noMutation}) {

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
	                    context.commit("clearProjectInfoFilter");
					}
                    return response;
                })
                .catch(function(err) {
                    console.error(err);
                    throw(err);
                });
        },
        findProjectInfoFilter(context, {filter, header, isAppend, noMutation}) {

            return api.find(findUrl, filter, header)
                .then(function(response) {

					if(! noMutation) {
						if (isAppend) {
							context.commit("appendProjectInfoFilter__List", response.List);
						} else {
							context.commit("setProjectInfoFilter__List", response.List);
						}
					}

                    return response;
                })
                .catch(function(err) {
                    console.error(err);
                    throw(err);
                });
        },
        loadProjectInfoFilter(context, {id, filter, header, noMutation}) {

            return api.find(readUrl + id, filter, header)
                .then(function(response) {

					if(! noMutation) {
	                    context.commit("setProjectInfoFilter", response.Model);
					}
                    return response;
                })
                .catch(function(err) {
                    console.error(err);
                    throw(err);
                });
        },
        updateProjectInfoFilter(context, {id, data, filter, header, noMutation}) {

            let url = updateUrl + id;
            if (Array.isArray && Array.isArray(data)) {
                url = multiUpdateUrl
            }

            return api.update(url, data, filter, header)
                .then(function(response) {
					if(! noMutation) {
	                    context.commit("setProjectInfoFilter", response.Model);
					}
                    return response;
                })
                .catch(function(err) {
                    console.error(err);
                    throw(err);
                });
        },
        findOrCreateProjectInfoFilter(context, {id, data, filter, header, noMutation}) {

            return api.update(findOrCreateUrl, data, filter, header)
                .then(function(response) {

					if(! noMutation) {
	                    context.commit("setProjectInfoFilter", response.Model);
					}
                    return response;
                })
                .catch(function(err) {
                    console.error(err);
                    throw(err);
                });
        },
        clearListProjectInfoFilter(context) {
            context.commit("clearListProjectInfoFilter");
        },
        clearProjectInfoFilter(context) {
            context.commit("clearProjectInfoFilter");
        },
    },
    getters: {
        getProjectInfoFilter: (state) => {
            return state.ProjectInfoFilter;
        },
        getProjectInfoFilterById: state => id => {
            return state.ProjectInfoFilter__List.find(item => item.Id === id);
        },
        getListProjectInfoFilter: (state) => {
            return state.ProjectInfoFilter__List;
        },
        getRoute__ProjectInfoFilter: state => action => {
            return state.ProjectInfoFilter__Routes[action];
        },
        getRoutes__ProjectInfoFilter: state => {
            return state.ProjectInfoFilter__Routes;
        },
    },
    mutations: {
        setProjectInfoFilter(state, data) {
            state.ProjectInfoFilter = data;
        },
        setProjectInfoFilter__List(state, data) {
            state.ProjectInfoFilter__List = data || [];
        },
        appendProjectInfoFilter__List(state, data) {

            if (! state.ProjectInfoFilter__List) {
                state.ProjectInfoFilter__List = [];
            }

            state.ProjectInfoFilter__List = state.ProjectInfoFilter__List.concat(data);
        },
        clearProjectInfoFilter(state) {
            state.ProjectInfoFilter = new ProjectInfoFilter();
        },
        clearListProjectInfoFilter(state) {
            state.ProjectInfoFilter__List = [];
        },
		updateProjectInfoFilterById(state, data) {
    		let index = findItemIndex(state.ProjectInfoFilter__List, function(item) {
	        	return item.Id === data.Id;
	    	});
	    
	    	if (index || index === 0) {
		        state.ProjectInfoFilter__List.splice(index, 1, data);
    		}
		},
		deleteProjectInfoFilterFromList(state, id) {
		    let index = findItemIndex(state.ProjectInfoFilter__List, function(item) {
		        return item.Id === id;
		    });
		    
		    if (index || index === 0) {
		        state.ProjectInfoFilter__List.splice(index, 1);
		    }
		},
		addProjectInfoFilterItemToList(state, item) {

			if (state.ProjectInfoFilter__List === null) {
				state.ProjectInfoFilter__List = [];
			}

		    state.ProjectInfoFilter__List.push(item);
		},
    },
    state: {
        ProjectInfoFilter: new ProjectInfoFilter(),
        ProjectInfoFilter__List: [],
        ProjectInfoFilter__Routes: {
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

export default projectInfoFilter;
