function Names(config) {

    let Name = config && config.name && config.name.length ? config.name : "";
    let Description = config && config.description && config.description.length ? config.description : "";

    this.GetName = () => {return Name};
    this.SetName = (name) => {Name = name; return this;};

    this.GetDescription = () => {return Description};
    this.SetDescription = (desc) => {Description = desc; return this;};
}

export default Names;