import Task from "./elements/task";
import Condition from "./elements/condition";
import Divider from "./elements/divider";
import Background from "./elements/background";
import Clone from "./actions/clone";
import MultiSelection from "./elements/multiselection";
import Resize from "./actions/resize";
import Nothing from "./actions/nothing";

function ElementsRegister () {

    return {
        MultiSelection,
        Task,
        Condition,
        Divider,
        Background,
        Actions: {
            Clone: Clone,
            Resize: Resize,
            Nothing: Nothing,
        }
    };

}

export default ElementsRegister();