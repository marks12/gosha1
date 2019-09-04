import ElementsRegister from "./elements-register";

function Toolbox(config) {

    let stdHeight = 40;
    let spaceBeetween = 20;

    this.AddItem(new ElementsRegister.Divider({
        Coords: {
            X: 200,
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
            X: 0,
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
            X: 80,
            Y: spaceBeetween,
        },
        Width: stdHeight * 1.6,
        Height: stdHeight,
        OnMove: new ElementsRegister.Actions.Clone(ElementsRegister.Task),
    }));
}

export default Toolbox;
