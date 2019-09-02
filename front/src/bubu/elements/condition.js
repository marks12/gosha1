import Names from "../attributes/names";
import Types from "../attributes/types";
import {TYPES} from "../constants";
import Coordinates from "../attributes/coordinates";
import Id from "../attributes/id";
import Draw from "../actions/draw";
import Size from "../attributes/size";
import Move from "../actions/move";

function Condition(config) {

    Id.apply(this, arguments);
    Coordinates.apply(this, arguments);
    Names.apply(this, arguments);
    Types.apply(this, arguments);
    this.SetType(TYPES.CONDITION);

    Size.apply(this, arguments);

    Draw.apply(this, arguments);
    Move.apply(this, arguments);

    return this;
}

export default Condition;