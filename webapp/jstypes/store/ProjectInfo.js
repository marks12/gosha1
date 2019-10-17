
import {ProjectInfo} from "../apiModel";
import api from "../api";
import {findItemIndex} from "../common";

let findUrl = "/api/v1/projectInfo";
let readUrl = "/api/v1/projectInfo/"; // + id
let createUrl = "/api/v1/projectInfo";
let updateUrl = "/api/v1/projectInfo/"; // + id
let deleteUrl = "/api/v1/projectInfo/"; // + id
let findOrCreateUrl = "/api/v1/projectInfo"; // + id

const projectInfo = {
    actions: {
        createProjectInfo(context, {data, filter, header}) {

            return api.create(createUrl, data, filter, header)
                .then(function(response) {

                    context.commit("setProjectInfo", response.Model);

                    return response;
                })
                .catch(function(err) {
                    return err;
                });
        },
        deleteProjectInfo(context, {id, header}) {

            return api.remove(deleteUrl + id, header)
                .then(function(response) {
                    context.commit("clearProjectInfo");
                    return response;
                })
                .catch(function(err) {
                    return err;
                });
        },
        findProjectInfo(context, {filter, header}) {

            return api.find(findUrl, filter, header)
                .then(function(response) {

                    context.commit("setProjectInfo__List", response.List);

                    return response;
                })
                .catch(function(err) {
                    return err;
                });
        },
        loadProjectInfo(context, {id, filter, header}) {

            return api.find(readUrl + id, filter, header)
                .then(function(response) {

                    context.commit("setProjectInfo", response.Model);

                    return response;
                })
                .catch(function(err) {
                    return err;
                });
        },
        updateProjectInfo(context, {id, data, filter, header}) {

            return api.update(updateUrl + id, data, filter, header)
                .then(function(response) {

                    context.commit("setProjectInfo", response.Model);

                    return response;
                })
                .catch(function(err) {
                    return err;
                });
        },
        findOrCreateProjectInfo(context, {id, data, filter, header}) {

            return api.update(findOrCreateUrl, data, filter, header)
                .then(function(response) {

                    context.commit("setProjectInfo", response.Model);

                    return response;
                })
                .catch(function(err) {
                    return err;
                });
        },
        clearListProjectInfo(context) {
            context.commit("clearListProjectInfo");
        },
    },
    getters: {
        getProjectInfo: (state) => {
            return state.ProjectInfo;
        },
        getProjectInfoById: state => id => {
            return state.ProjectInfo__List.find(item => item.Id === id);
        },
        getListProjectInfo: (state) => {
            return state.ProjectInfo__List;
        },
    },
    mutations: {
        setProjectInfo(state, data) {
            state.ProjectInfo = data;
        },
        setProjectInfo__List(state, data) {
            state.ProjectInfo__List = data || [];
        },
        clearProjectInfo(state) {
            state.ProjectInfo = new ProjectInfo();
        },
        clearListProjectInfo(state) {
            state.ProjectInfo__List = [];
        },
		updateProjectInfoById(state, data) {
    		let index = findItemIndex(state.ProjectInfo__List, function(item) {
	        	return item.Id === data.Id;
	    	});
	    
	    	if (index || index === 0) {
		        state.ProjectInfo__List.splice(index, 1, data);
    		}
		},
		deleteProjectInfoFromList(state, id) {
		    let index = findItemIndex(state.ProjectInfo__List, function(item) {
		        return item.Id === id;
		    });
		    
		    if (index || index === 0) {
		        state.ProjectInfo__List.splice(index, 1);
		    }
		},
		addProjectInfoItemToList(state, item) {

			if (state.ProjectInfo__List === null) {
				state.ProjectInfo__List = [];
			}

		    state.ProjectInfo__List.push(item);
		},
    },
    state: {
        ProjectInfo: new ProjectInfo(),
        ProjectInfo__List: [],
    },
};

export default projectInfo;
