import Task from "./elements/task";
import Link from "./elements/link";
import ConnectorPoint from "./elements/connector-point";
import Condition from "./elements/condition";
import Divider from "./elements/divider";
import Background from "./elements/background";
import ZeroPoint from "./elements/zero-point";
import Clone from "./actions/clone";
import MultiSelection from "./elements/multiselection";
import Resize from "./actions/resize";
import Nothing from "./actions/nothing";
import Button from "./elements/button";

function ElementsRegister () {

    return {
        MultiSelection,
        Task,
        Link,
        ConnectorPoint,
        Condition,
        Divider,
        Background,
        ZeroPoint,
        Button,
        Actions: {
            Clone: Clone,
            Resize: Resize,
            Nothing: Nothing,
        }
    };

}

export default ElementsRegister();