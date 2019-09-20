import Types from "../attributes/types";
import {TYPES} from "../constants";
import BasicSet from "../attributes/basic-set";

function Background(config) {

    BasicSet.apply(this, arguments);

    Types.apply(this, arguments);
    this.SetType(TYPES.background);

    return this;
}

export default Background;