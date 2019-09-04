function Color(config) {

    let defaultColor = '#6e6e6e';

    let color = config && config.Color ? config.Color : defaultColor;

    this.GetColor = () => {return color};
    this.SetColor = (c) => {color = c; return this;};
    this.SetColorDefault = () => {color = defaultColor; return this;};
    this.GetColorDefault = () => {return defaultColor;};
}

export default Color;