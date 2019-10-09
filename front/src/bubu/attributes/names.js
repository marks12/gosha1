function Names(config) {

    let Name = config && config.Name && config.Name.length ? config.Name : "";
    let Description = config && config.Description && config.Description.length ? config.Description : "";

    this.GetName = () => {return Name};
    this.SetName = (name) => {Name = name; return this;};

    this.GetDescription = () => {return Description};
    this.SetDescription = (desc) => {Description = desc; return this;};
}

export default Names;