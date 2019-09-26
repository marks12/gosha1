import BasicSet from "../attributes/basic-set";
import Types from "../attributes/types";
import {TYPES} from "../constants";
import Click from "../actions/click";

function Button() {

    BasicSet.apply(this, arguments);
    Click.apply(this, arguments);

    Types.apply(this, arguments);
    this.SetType(TYPES.button);

    return this;
}

export default Button;