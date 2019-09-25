import BasicSet from "../attributes/basic-set";
import Types from "../attributes/types";
import {TYPES} from "../constants";

function Button() {

    BasicSet.apply(this, arguments);

    Types.apply(this, arguments);
    this.SetType(TYPES.button);

    return this;
}

export default Button;