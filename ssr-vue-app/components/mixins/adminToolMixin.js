import {mapGetters} from "vuex";

export const adminToolMixin = {
  data() {
    return {
      $_adminTool_ADMIN_TOOL_CLASS_NAME: String,
      $_adminTool_ADMIN_TOOL_HOVERED_CLASS_NAME: String,
      $_adminTool_ADMIN_TOOL_SELECTED_CLASS_NAME: String,
      $_adminTool_componentInitialized: Boolean
    }
  },
  props: {
    componentId: {
      type: String,
      default: ''
    }
  },
  computed: {
    ...mapGetters([
      'adminState',
    ]),
    $_adminTool_componentName: function () {
      return this.$options.name;
    },
  },
  watch: {
    adminState:  function (newAdminState) {
      if(newAdminState) {
        this.$_adminTool_initializeAdminTools();
      }
    }
  },
  mounted() {
    this.$_adminTool_componentInitialized = false;
    this.$socket.bindType("Mirror", (result) => {
      switch (result.data.route) {
        case "PageSectionCallback":
          this.$_adminTool_pageSectionCallback(result.data.data);
          break;
      }
    });
    if (this.adminState) {
      this.$_adminTool_initializeAdminTools();
    }
  },
  methods: {
    $_adminTool_pageSectionCallback(data) {
      if (data.id.toString() === this.componentId.toString()) {
        this.$_adminTool_selectComponent();
      }
    },
    $_adminTool_initializeAdminTools() {
      if (this.componentId && !this.$_adminTool_componentInitialized) {
        this.$_adminTool_componentInitialized = true;
        this.$_adminTool_ADMIN_TOOL_CLASS_NAME = "admin-tool";
        this.$_adminTool_ADMIN_TOOL_HOVERED_CLASS_NAME = "admin-tool_hovered";
        this.$_adminTool_ADMIN_TOOL_SELECTED_CLASS_NAME = "admin-tool_selected";

        this.$el.classList.add(this.$_adminTool_ADMIN_TOOL_CLASS_NAME);
        this.$el.onmousemove = (event) => {
          this.$_adminTool_onMouseMove(event, this.$el);
        };
        this.$el.onclick = (event) => {
          this.$_adminTool_onClick(event);
        };
        document.onmouseleave = () => {
          this.$_adminTool_removeClass(this.$_adminTool_ADMIN_TOOL_HOVERED_CLASS_NAME);
        };
      }
    },
    $_adminTool_removeClass(className) {
      [].forEach.call(document.getElementsByClassName(className), (el) => {
        el.classList.remove(className);
      });
    },
    $_adminTool_onMouseMove(event, element) {
      event.stopPropagation();
      if (element.classList.contains(this.$_adminTool_ADMIN_TOOL_HOVERED_CLASS_NAME))
        return;
      this.$_adminTool_removeClass(this.$_adminTool_ADMIN_TOOL_HOVERED_CLASS_NAME);
      element.classList.add(this.$_adminTool_ADMIN_TOOL_HOVERED_CLASS_NAME);
    },
    $_adminTool_onClick(event) {
      event.stopPropagation();
      this.$_adminTool_selectComponent();
      this.$socket.send({
        type: 'Mirror',
        data: {
          route: 'PageSectionCallback',
          data: {
            id: this.componentId,
            componentName: this.$_adminTool_componentName
          }
        }
      })
    },
    $_adminTool_selectComponent() {
      if (this.$el.classList.contains(this.$_adminTool_ADMIN_TOOL_SELECTED_CLASS_NAME))
        return;
      this.$_adminTool_removeClass(this.$_adminTool_ADMIN_TOOL_SELECTED_CLASS_NAME);
      this.$el.classList.add(this.$_adminTool_ADMIN_TOOL_SELECTED_CLASS_NAME);
    }
  }
};

