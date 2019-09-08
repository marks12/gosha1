import Id from "../attributes/id";
import Coordinates from "../attributes/coordinates";

function ZeroPoint(config) {

    Id.apply(this, arguments);
    Coordinates.apply(this, arguments);
}

export default ZeroPoint;
