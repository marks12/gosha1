import LinkDirection from "../attributes/link-direction";
import LinkArrow from "../attributes/link-arrow";
import LinkLine from "../attributes/link-line";
import BasicSet from "../attributes/basic-set";
import Types from "../attributes/types";
import {TYPES} from "../constants";

function Link(config) {

    BasicSet.apply(this, arguments);

    LinkArrow.apply(this, arguments);
    LinkLine.apply(this, arguments);
    LinkDirection.apply(this, arguments);

    Types.apply(this, arguments);
    this.SetType(TYPES.link);

}

export default Link;