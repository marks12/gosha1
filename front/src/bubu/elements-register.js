import Task from "./elements/task";
import Link from "./elements/link";
import Condition from "./elements/condition";
import Divider from "./elements/divider";
import Background from "./elements/background";
import ZeroPoint from "./elements/zero-point";
import Clone from "./actions/clone";
import MultiSelection from "./elements/multiselection";
import Resize from "./actions/resize";
import Nothing from "./actions/nothing";

function ElementsRegister () {

    return {
        MultiSelection,
        Task,
        Link,
        Condition,
        Divider,
        Background,
        ZeroPoint,
        Actions: {
            Clone: Clone,
            Resize: Resize,
            Nothing: Nothing,
        }
    };

}

export default ElementsRegister();