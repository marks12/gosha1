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
        Width: this.canvas.parentNode.parentElement.clientWidth - constants.spaceBetween * 2 - constants.toolboxWidth,
        Height: this.canvas.parentNode.parentElement.clientHeight - constants.spaceBetween * 2,
        OnMove: new ElementsRegister.Actions.Nothing(),
        IsSelectable: false,
        Color: '#fff',
    });

    this.AddItem(workArea);

    this.ZeroPoint.Coords.SetX(workArea.Coords.GetX());
    this.ZeroPoint.Coords.SetY(workArea.Coords.GetY());
}

export default WorkArea;
