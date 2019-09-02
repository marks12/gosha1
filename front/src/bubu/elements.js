import Task from "./elements/task";
import Condition from "./elements/condition";
import Divider from "./elements/divider";
import Clone from "./actions/clone";

function ElementsRegister () {

    return {
        Task,
        Condition,
        Divider,
        Actions: {
            Clone: Clone,
        }
    };

}

export default ElementsRegister();