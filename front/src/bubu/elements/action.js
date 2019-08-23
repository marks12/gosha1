import Id from "../attributes/id";
import Names from "../attributes/names";
import Types from "../attributes/types";
import {TYPES} from "../constants";
import Coordinates from "../attributes/coordinates";

function Action(config) {

    Id.apply(this, arguments);
    Coordinates.apply(this, arguments);
    Names.apply(this, arguments);
    Types.apply(this, arguments);
}

export default Action;