import Id from "../attributes/id";
import Types from "../attributes/types";
import {TYPES} from "../constants";
import Coordinates from "../attributes/coordinates";
import Draw from "../actions/draw";
import Size from "../attributes/size";
import Color from "../attributes/color";
import Move from "../actions/move";
import BasicSet from "../attributes/basic-set";

function Divider(config) {


    BasicSet.apply(this, arguments);
    Types.apply(this, arguments);
    this.SetType(TYPES.divider);

    return this;
}

export default Divider;