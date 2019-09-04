import Task from "./elements/task";
import Condition from "./elements/condition";
import Divider from "./elements/divider";
import Clone from "./actions/clone";
import MultiSelection from "./elements/multiselection";
import Resize from "./actions/resize";

function ElementsRegister () {

    return {
        MultiSelection,
        Task,
        Condition,
        Divider,
        Actions: {
            Clone: Clone,
            Resize: Resize,
        }
    };

}

export default ElementsRegister();