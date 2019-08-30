import Elements from "./elements";

function Toolbox(config) {

    this.AddItem(new Elements.Condition({
        Name: "Condition",
        Description: "Move this condition to work area for create new Condition",
    }));

    this.AddItem(new Elements.Task({
        Name: "Task",
        Description: "Move this task to work area for create new Task",
    }));
}

export default Toolbox;
