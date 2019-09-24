function Parent(config) {

    let parentId = config && config.Parent ? config.Parent : null;

    this.SetParentId = (id) => {
        parentId = id;
        return this;
    };

    this.GetParentId = () => {
        return parentId;
    };
}

export default Parent;