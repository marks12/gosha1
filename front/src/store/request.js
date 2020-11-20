const request = {
    actions: {
        setRequestUrl(context, val) {
            context.commit("setRequestUrl", val);
        },
        setRouteUrl(context, val) {
            context.commit("setRouteUrl", val);
        },
        setHeaders(context, val) {
            context.commit("setHeaders", val);
        },
        setBodyModel(context, val) {
            context.commit("setBodyModel", val);
        },
        addHeader(context, {Name, Value}) {
            context.commit("addHeader", {Name: Name, Value: Value});
        },
        setFilters(context, val) {
            context.commit("setFilters", val);
        },
        addFilter(context, {Name, Value}) {
            context.commit("addFilter", {Name: Name, Value: Value});
        },
        resetRequest(context) {
            context.commit("headers", {});
            context.commit("bodyModel", {});
            context.commit("filters", {});
            context.commit("requestUrl", "");
        },
    },
    getters: {
        getRequestUrl(state) {
            return state.requestUrl;
        },
        getRouteUrl(state) {
            return state.routeUrl;
        },
        getHeaders: (state) => {
            return state.headers;
        },
        getFilters: (state) => {
            return state.filters;
        },
        getBodyModel: (state) => {
            return state.bodyModel;
        },
    },
    mutations: {
        setHeaders(state, data) {
            state.headers = data;
        },
        setRequestUrl(state, data) {
            state.requestUrl = data;
        },
        setRouteUrl(state, data) {
            state.routeUrl = data;
        },
        addHeader(state, {Name, Value}) {
            state.headers[Name] = Value;
        },
        setFilters(state, data) {
            state.filters = data;
        },
        addFilter(state, {Name, Value}) {
            state.filters[Name] = Value;
        },
        setBodyModel(state, data) {
            state.bodyModel = data;
        },
    },
    state: {
        headers: {},
        filters: {},
        bodyModel: {},
        requestUrl: "",
        routeUrl: "",
    },
};

export default request;
