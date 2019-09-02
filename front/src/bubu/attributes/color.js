function Color(config) {

    let color = config && config.Color ? config.Color : '#000';

    this.GetColor = () => {return color};
    this.SetColor = (c) => {color = c; return this;};
}

export default Color;