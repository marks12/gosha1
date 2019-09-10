import ElementsRegister from "./elements-register";
import {TYPES as constants} from "./constants";

function WorkArea(config) {

    let workArea = new ElementsRegister.Background({
        Name: "work-area-background",
        Description: "Подложка для рабочей области",
        Coords: {
            X: constants.toolboxWidth + constants.spaceBetween,
            Y: constants.spaceBetween,
        },
        Width: this.GetCanvasWidth() - constants.spaceBetween * 2 - constants.toolboxWidth,
        Height: this.GetCanvasHeight() - constants.spaceBetween * 2,
        OnMove: new ElementsRegister.Actions.Nothing(),
        IsSelectable: false,
        Color: '#fff',
    });

    this.AddItem(workArea);

    this.Zero.Coords.SetX(0);
    this.Zero.Coords.SetY(0);
}

export default WorkArea;
