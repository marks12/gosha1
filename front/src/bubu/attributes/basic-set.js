import ToJson from "../actions/tojson";
import Id from "../attributes/id";
import Coordinates from "../attributes/coordinates";
import Draw from "../actions/draw";
import Size from "../attributes/size";
import Text from "../attributes/text";
import Color from "../attributes/color";
import Buttons from "../attributes/buttons";
import Visibility from "../attributes/visibility";
import Names from "../attributes/names";
import Connectors from "../attributes/connectors";
import Move from "../actions/move";
import Resize from "../actions/resize";
import Click from "../actions/click";
import Selection from "../actions/selection";

function BasicSet(config) {

    Names.apply(this, arguments);
    Size.apply(this, arguments);
    Text.apply(this, arguments);
    Color.apply(this, arguments);
    Draw.apply(this, arguments);
    Move.apply(this, arguments);
    Id.apply(this, arguments);
    Coordinates.apply(this, arguments);
    Selection.apply(this, arguments);
    Resize.apply(this, arguments);
    Connectors.apply(this, arguments);
    Visibility.apply(this, arguments);
    Buttons.apply(this, arguments);
    ToJson.apply(this, arguments);

    return this;
}

export default BasicSet;