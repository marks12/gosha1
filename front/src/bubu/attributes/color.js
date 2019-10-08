import {TYPES} from "../constants";

function Color(config) {

    let defaultColor = TYPES.currentTheme.primaryStroke;
    let fillColor = TYPES.currentTheme.primaryFill;
    let color = config && config.Color ? config.Color : defaultColor;

    this.GetColor = () => {return color};
    this.SetColor = (c) => {color = c; return this;};
    this.SetColorDefault = () => {color = defaultColor; return this;};
    this.GetColorDefault = () => {return defaultColor;};
    this.GetFillColorDefault = () => {return fillColor;};
}

export default Color;