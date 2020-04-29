const defaultMaxWidth = 'col4';

const panel = {
    actions: {
        setPanelMaxWidth(context, val) {
            context.commit("setPanelMaxWidth", val);
        },
        resetPanelMaxWidth(context) {
            context.commit("setPanelMaxWidth", defaultMaxWidth);
        },
    },
    getters: {
        getPanelMaxWidth: (state) => {
            return state.maxWidth;
        },
    },
    mutations: {
        setPanelMaxWidth(state, data) {
            state.maxWidth = data;
        },
    },
    state: {
        maxWidth: defaultMaxWidth,
    },
};

export default panel;
