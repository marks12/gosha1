import Task from "./elements/task";
import Condition from "./elements/condition";

function Elements () {

    this.Task = Task;
    this.Condition = Condition;

    return this;

}

export default new Elements();