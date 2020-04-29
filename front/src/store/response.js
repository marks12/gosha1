const response = {
    actions: {
        setResponseUrl(context, val) {
            context.commit("setResponseUrl", val);
        },
        setResponse(context, val) {
            context.commit("setResponse", val);
        },
        resetResponse(context) {
            context.commit("response", {});
            context.commit("responseUrl", "");
        },
    },
    getters: {
        getResponseUrl(state) {
            return state.responseUrl;
        },
        getResponse: (state) => {
            return state.response;
        },
    },
    mutations: {
        setResponse(state, data) {
            state.response = data;
        },
        setResponseUrl(state, data) {
            state.responseUrl = data;
        },
    },
    state: {
        response: {},
        responseUrl: "",
    },
};

export default response;
