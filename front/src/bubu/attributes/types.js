function Types(child, type) {

    let Type = type;


    this.GetType = () => {return Type};
    this.SetType = (t) => {Type = t; return this;};
}

export default Types;