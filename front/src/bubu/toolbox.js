import ElementsRegister from "./elements-register";
import {TYPES as constants} from "./constants";

function Toolbox(config) {

    this.AddItem(new ElementsRegister.Background({
        Name: "toolbox-background",
        Description: "Подложка для тулбокса",
        Coords: {
            X: 1,
            Y: constants.spaceBetween,
        },
        Width: constants.toolboxWidth,
        Height: this.canvas.parentNode.parentElement.clientHeight - 2 * constants.spaceBetween,
        OnMove: new ElementsRegister.Actions.Nothing(),
        IsSelectable: false,
        Color: '#fff',
    }));

    this.AddItem(new ElementsRegister.Condition({
        Name: "Condition",
        Description: "Move this condition to work area for create new Condition",
        Coords: {
            X: constants.spaceBetween,
            Y: constants.spaceBetween * 2,
        },
        Width: constants.stdHeight,
        Height: constants.stdHeight,
        OnMove: new ElementsRegister.Actions.Clone(ElementsRegister.Condition),
    }));

    this.AddItem(new ElementsRegister.Task({
        Name: "Task",
        Description: "Move this task to work area for create new Task",
        Coords: {
            X: constants.spaceBetween,
            Y: 3 * constants.spaceBetween + constants.stdHeight,
        },
        Width: constants.stdHeight,
        Height: constants.stdHeight / 1.6,
        OnMove: new ElementsRegister.Actions.Clone(ElementsRegister.Task),
    }));
}

export default Toolbox;
