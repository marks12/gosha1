import Elements from "./elements";

function Toolbox(config) {

    this.AddItem(new Elements.Divider({
        Coords: {
            X: 200,
            Y: 20,
        },
        Width: 1,
        Height: 200,
        Color: '#ccc'
    }));

    this.AddItem(new Elements.Condition({
        Name: "Condition",
        Description: "Move this condition to work area for create new Condition",
        Coords: {
            X: 0,
            Y: 20,
        },
        Width: 60,
        Height: 60,
        OnMove: new Elements.Actions.Clone(Elements.Condition),
    }));


    this.AddItem(new Elements.Task({
        Name: "Task",
        Description: "Move this task to work area for create new Task",
        Coords: {
            X: 80,
            Y: 20,
        },
        Width: 100,
        Height: 60,
        OnMove: new Elements.Actions.Clone(Elements.Task),
    }));
}

export default Toolbox;
