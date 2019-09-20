import Types from "../attributes/types";
import {TYPES} from "../constants";
import BasicSet from "../attributes/basic-set";

function Task(config) {

    BasicSet.apply(this, arguments);

    Types.apply(this, arguments);
    this.SetType(TYPES.task);

    return this;
}

export default Task;