import ElementsRegister from "./elements-register";

function Toolbox(config) {

    let stdHeight = 40;
    let spaceBeetween = 20;

    this.AddItem(new ElementsRegister.Background({
        Name: "toolbox-background",
        Description: "Подложка для тулбокса",
        Coords: {
            X: 0,
            Y: 0,
        },
        Width: 100,
        Height: this.canvas.parentNode.parentElement.clientHeight,
        OnMove: new ElementsRegister.Actions.Nothing(),
        IsSelectable: false,
        Color: '#fff',
    }));

    this.AddItem(new ElementsRegister.Divider({
        Coords: {
            X: 100,
            Y: spaceBeetween,
        },
        Width: 1,
        Height: this.canvas.parentNode.parentElement.clientHeight,
        Color: '#ccc'
    }));

    this.AddItem(new ElementsRegister.Condition({
        Name: "Condition",
        Description: "Move this condition to work area for create new Condition",
        Coords: {
            X: spaceBeetween,
            Y: spaceBeetween,
        },
        Width: stdHeight,
        Height: stdHeight,
        OnMove: new ElementsRegister.Actions.Clone(ElementsRegister.Condition),
    }));

    this.AddItem(new ElementsRegister.Task({
        Name: "Task",
        Description: "Move this task to work area for create new Task",
        Coords: {
            X: spaceBeetween,
            Y: 2 * spaceBeetween + stdHeight,
        },
        Width: stdHeight,
        Height: stdHeight / 1.6,
        OnMove: new ElementsRegister.Actions.Clone(ElementsRegister.Task),
    }));
}

export default Toolbox;
