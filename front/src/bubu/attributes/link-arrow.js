import {TYPES as types} from "../constants";

function LinkArrow() {

    let type = types.arrowTypeSimple;

    this.SetLinkArrowType = (t) => {
        type = t;

        return this;
    };

    this.GetLinkArrowType = () => {
        return type;
    };
}

export default LinkArrow;