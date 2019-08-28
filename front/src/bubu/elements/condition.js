import Names from "../attributes/names";
import Types from "../attributes/types";
import {TYPES} from "../constants";
import Coordinates from "../attributes/coordinates";
import Id from "../attributes/id";
import Draw from "../actions/draw";
import Size from "../attributes/size";

function Condition(config) {

    Id.apply(this, arguments);
    Coordinates.apply(this, arguments);
    Names.apply(this, arguments);
    Types.apply(this, arguments);
    this.SetType(TYPES.CONDITION);

    Size.apply(this, arguments);
    this.SetSize(200, 200);

    Draw.apply(this, arguments);

}

export default Condition;