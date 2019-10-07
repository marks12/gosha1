import Types from "../attributes/types";
import {TYPES} from "../constants";
import BasicSet from "../attributes/basic-set";
import Delete from "../actions/delete";

function Condition(config) {

    BasicSet.apply(this, arguments);
    Delete.apply(this, arguments);

    Types.apply(this, arguments);
    this.SetType(TYPES.condition);

    return this;
}

export default Condition;