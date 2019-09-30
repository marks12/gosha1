function Color(config) {

    let defaultColor = '#6e6e6e';
    let fillColor = '#fff';

    let color = config && config.Color ? config.Color : defaultColor;

    this.GetColor = () => {return color};
    this.SetColor = (c) => {color = c; return this;};
    this.SetColorDefault = () => {color = defaultColor; return this;};
    this.GetColorDefault = () => {return defaultColor;};
    this.GetFillColorDefault = () => {return fillColor;};
}

export default Color;