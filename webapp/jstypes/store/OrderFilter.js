
import {OrderFilter} from "../apiModel";
import api from "../api";
import {findItemIndex} from "../common";

let findUrl = "/api/v1/orderFilter";
let readUrl = "/api/v1/orderFilter/"; // + id
let createUrl = "/api/v1/orderFilter";
let updateUrl = "/api/v1/orderFilter/"; // + id
let deleteUrl = "/api/v1/orderFilter/"; // + id
let findOrCreateUrl = "/api/v1/orderFilter"; // + id

const orderFilter = {
    actions: {
        createOrderFilter(context, {data, filter, header}) {

            return api.create(createUrl, data, filter, header)
                .then(function(response) {

                    context.commit("setOrderFilter", response.Model);

                    return response;
                })
                .catch(function(err) {
                    return err;
                });
        },
        deleteOrderFilter(context, {id, header}) {

            return api.remove(deleteUrl + id, header)
                .then(function(response) {
                    context.commit("clearOrderFilter");
                    return response;
                })
                .catch(function(err) {
                    return err;
                });
        },
        findOrderFilter(context, {filter, header}) {

            return api.find(findUrl, filter, header)
                .then(function(response) {

                    context.commit("setOrderFilter__List", response.List);

                    return response;
                })
                .catch(function(err) {
                    return err;
                });
        },
        loadOrderFilter(context, {id, filter, header}) {

            return api.find(readUrl + id, filter, header)
                .then(function(response) {

                    context.commit("setOrderFilter", response.Model);

                    return response;
                })
                .catch(function(err) {
                    return err;
                });
        },
        updateOrderFilter(context, {id, data, filter, header}) {

            return api.update(updateUrl + id, data, filter, header)
                .then(function(response) {

                    context.commit("setOrderFilter", response.Model);

                    return response;
                })
                .catch(function(err) {
                    return err;
                });
        },
        findOrCreateOrderFilter(context, {id, data, filter, header}) {

            return api.update(findOrCreateUrl, data, filter, header)
                .then(function(response) {

                    context.commit("setOrderFilter", response.Model);

                    return response;
                })
                .catch(function(err) {
                    return err;
                });
        },
        clearListOrderFilter(context) {
            context.commit("clearListOrderFilter");
        },
    },
    getters: {
        getOrderFilter: (state) => {
            return state.OrderFilter;
        },
        getListOrderFilter: (state) => {
            return state.OrderFilter__List;
        },
    },
    mutations: {
        setOrderFilter(state, data) {
            state.OrderFilter = data;
        },
        setOrderFilter__List(state, data) {
            state.OrderFilter__List = data;
        },
        clearOrderFilter(state) {
            state.OrderFilter = new OrderFilter();
        },
        clearListOrderFilter(state) {
            state.OrderFilter__List = [];
        },
		updateOrderFilterById(state, data) {
    		let index = findItemIndex(state.OrderFilter__List, function(item) {
	        	return item.Id === data.Id;
	    	});
	    
	    	if (index || index === 0) {
		        state.OrderFilter__List.splice(index, 1, data);
    		}
		},
		deleteOrderFilterFromList(state, id) {
		    let index = findItemIndex(state.OrderFilter__List, function(item) {
		        return item.Id === id;
		    });
		    
		    if (index || index === 0) {
		        state.OrderFilter__List.splice(index, 1);
		    }
		},
		addOrderFilterItemToList(state, item) {

			if (state.OrderFilter__List === null) {
				state.OrderFilter__List = [];
			}

		    state.OrderFilter__List.push(item);
		},
    },
    state: {
        OrderFilter: new OrderFilter(),
        OrderFilter__List: [],
    },
};

export default orderFilter;
