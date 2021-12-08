const request = {
    actions: {
        setRequestUrl(context, val) {
            context.commit("setRequestUrl", val);
        },
        setRouteUrl(context, val) {
            context.commit("setRouteUrl", val);
        },
        setServerUrl(context, val) {
            context.commit("setServerUrl", val);
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
        saveServerUrl(context) {
            return context.commit("saveServerUrl");
        },
    },
    getters: {
        getRequestUrl(state) {
            return state.requestUrl;
        },
        getRouteUrl(state) {
            return state.routeUrl;
        },
        getServerUrl(state) {
            return state.serverUrl;
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
        setServerUrl(state, data) {
            state.serverUrl = data;
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
        saveServerUrl(state) {
            let storeServers = localStorage.getItem("servers");
            let servers = JSON.parse(storeServers) || [];

            for (let i = 0; i < servers.length; i++) {

                if (servers[i] === state.serverUrl) {
                    return false;
                }
            }

            servers.push(state.serverUrl);
            localStorage.setItem("servers", JSON.stringify(servers));
        },
    },
    state: {
        headers: {},
        filters: {},
        bodyModel: {},
        requestUrl: "",
        routeUrl: "",
        serverUrl: "",
    },
};

export default request;
