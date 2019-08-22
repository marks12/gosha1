function Names(child) {

    let Name = "";
    let Description = "";

    return {
        GetName: () => {return Name},
        SetName: function (name) { Name = name; return child},
        GetDescription: () => {return Description},
        SetDescription: function (desc) { Description = desc; return child},
    };
}

export default Names;