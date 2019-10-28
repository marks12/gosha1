const adminTool = {
  actions: {
  },
  getters: {
    adminState: (state) => {
      return state.isAdminTool;
    },
  },
  mutations: {
    enableAdminTool (state) {
      state.isAdminTool = true;
    },
    disableAdminTool (state) {
      state.isAdminTool = false;
    }
  },
  state() {
    return {
      isAdminTool: false
    }
  },
};

export default adminTool;
