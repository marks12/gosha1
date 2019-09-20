import Types from "../attributes/types";
import {TYPES} from "../constants";
import BasicSet from "../attributes/basic-set";

function Condition(config) {

    BasicSet.apply(this, arguments);

    Types.apply(this, arguments);
    this.SetType(TYPES.condition);

    return this;
}

export default Condition;