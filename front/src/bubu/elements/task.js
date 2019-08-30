import Id from "../attributes/id";
import Names from "../attributes/names";
import Types from "../attributes/types";
import {TYPES} from "../constants";
import Coordinates from "../attributes/coordinates";
import Draw from "../actions/draw";
import Size from "../attributes/size";

function Task(config) {

    Id.apply(this, arguments);
    Coordinates.apply(this, arguments);
    Names.apply(this, arguments);
    Types.apply(this, arguments);
    this.SetType(TYPES.ACTION);

    Size.apply(this, arguments);
    this.SetSize(100, 100);


    Draw.apply(this, arguments);

    return this;
}

export default Task;