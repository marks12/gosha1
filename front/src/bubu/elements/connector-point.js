import BasicSet from "../attributes/basic-set";
import LinkDirection from "../attributes/link-direction";
import Parent from "../attributes/parent";
import CollectionItem from "../attributes/collection-item";
import Types from "../attributes/types";
import Ratio from "../attributes/ratio";
import {TYPES} from "../constants";

function ConnectorPoint() {

    CollectionItem.apply(this, arguments);
    Parent.apply(this, arguments);
    BasicSet.apply(this, arguments);
    Ratio.apply(this, arguments);
    LinkDirection.apply(this, arguments);

    Types.apply(this, arguments);
    this.SetType(TYPES.connectorPoint);

}

export default ConnectorPoint;