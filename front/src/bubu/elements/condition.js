import Names from "../attributes/names";
import Types from "../attributes/types";
import {TYPES} from "../constants";
import Coordinates from "../attributes/coordinates";

function Condition() {

    let condition = {};

    Object.assign(condition,
        new Coordinates(condition),
        new Names(condition),
        new Types(condition, TYPES.CONDITION),
    );

    return condition;
}

export default Condition;