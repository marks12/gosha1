import Types from "../attributes/types";
import {TYPES} from "../constants";
import BasicSet from "../attributes/basic-set";

function Multiselection(config) {

    BasicSet.apply(this, arguments);

    Types.apply(this, arguments);
    this.SetType(TYPES.MULTISELECTION);

    return this;
}

export default Multiselection;