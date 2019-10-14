
import {ProjectInfoFilter} from "../apiModel";
import api from "../api";
import {findItemIndex} from "../common";

let findUrl = "/api/v1/projectInfoFilter";
let readUrl = "/api/v1/projectInfoFilter/"; // + id
let createUrl = "/api/v1/projectInfoFilter";
let updateUrl = "/api/v1/projectInfoFilter/"; // + id
let deleteUrl = "/api/v1/projectInfoFilter/"; // + id
let findOrCreateUrl = "/api/v1/projectInfoFilter"; // + id

const projectInfoFilter = {
    actions: {
        createProjectInfoFilter(context, {data, filter, header}) {

            return api.create(createUrl, data, filter, header)
                .then(function(response) {

                    context.commit("setProjectInfoFilter", response.Model);

                    return response;
                })
                .catch(function(err) {
                    return err;
                });
        },
        deleteProjectInfoFilter(context, {id, header}) {

            return api.remove(deleteUrl + id, header)
                .then(function(response) {
                    context.commit("clearProjectInfoFilter");
                    return response;
                })
                .catch(function(err) {
                    return err;
                });
        },
        findProjectInfoFilter(context, {filter, header}) {

            return api.find(findUrl, filter, header)
                .then(function(response) {

                    context.commit("setProjectInfoFilter__List", response.List);

                    return response;
                })
                .catch(function(err) {
                    return err;
                });
        },
        loadProjectInfoFilter(context, {id, filter, header}) {

            return api.find(readUrl + id, filter, header)
                .then(function(response) {

                    context.commit("setProjectInfoFilter", response.Model);

                    return response;
                })
                .catch(function(err) {
                    return err;
                });
        },
        updateProjectInfoFilter(context, {id, data, filter, header}) {

            return api.update(updateUrl + id, data, filter, header)
                .then(function(response) {

                    context.commit("setProjectInfoFilter", response.Model);

                    return response;
                })
                .catch(function(err) {
                    return err;
                });
        },
        findOrCreateProjectInfoFilter(context, {id, data, filter, header}) {

            return api.update(findOrCreateUrl, data, filter, header)
                .then(function(response) {

                    context.commit("setProjectInfoFilter", response.Model);

                    return response;
                })
                .catch(function(err) {
                    return err;
                });
        },
        clearListProjectInfoFilter(context) {
            context.commit("clearListProjectInfoFilter");
        },
    },
    getters: {
        getProjectInfoFilter: (state) => {
            return state.ProjectInfoFilter;
        },
        getListProjectInfoFilter: (state) => {
            return state.ProjectInfoFilter__List;
        },
    },
    mutations: {
        setProjectInfoFilter(state, data) {
            state.ProjectInfoFilter = data;
        },
        setProjectInfoFilter__List(state, data) {
            state.ProjectInfoFilter__List = data || [];
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
    },
};

export default projectInfoFilter;
