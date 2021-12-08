const response = {
    actions: {
        setResponseUrl(context, val) {
            context.commit("setResponseUrl", val);
        },
        setResponse(context, val) {
            context.commit("setResponse", val);
        },
        setResponseCode(context, val) {
            context.commit("setResponseCode", val);
        },
        setResponseStatusText(context, val) {
            context.commit("setResponseStatusText", val);
        },
        resetResponse(context) {
            context.commit("setResponse", "");
            context.commit("setResponseUrl", "");
            context.commit("setResponseCode", null);
            context.commit("setResponseStatusText", "");
        },
    },
    getters: {
        getResponseUrl(state) {
            return state.responseUrl;
        },
        getResponse: (state) => {
            return state.response;
        },
        getResponseCode: (state) => {
            return state.responseCode;
        },
        getResponseStatusText: (state) => {
            return state.responseStatusText;
        },
    },
    mutations: {
        setResponse(state, data) {
            state.response = data;
        },
        setResponseCode(state, data) {
            state.responseCode = data;
        },
        setResponseStatusText(state, data) {
            state.responseStatusText = data;
        },
        setResponseUrl(state, data) {
            state.responseUrl = data;
        },
    },
    state: {
        responseCode: null,
        responseStatusText: "",
        response: "",
        responseUrl: "",
    },
};

export default response;
