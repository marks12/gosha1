import Names from "../attributes/names";
import Types from "../attributes/types";
import {TYPES} from "../constants";
import Coordinates from "../attributes/coordinates";

function Action() {

    let action = {};

    Object.assign(action,
        new Coordinates(action),
        new Names(action),
        new Types(action, TYPES.CONDITION),
    );

    return action;
}

export default Action;