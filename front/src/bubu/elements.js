import Task from "./elements/task";
import Condition from "./elements/condition";
import Clone from "./actions/clone";

function ElementsRegister () {

    return {
        Task,
        Condition,
        Actions: {
            Clone: Clone,
        }
    };

}

export default ElementsRegister();