import {TYPES as types} from "../constants";

function LinkArrow(config) {

    let type = config && config.LinkArrowType ? config.LinkArrowType : types.arrowTypeSimple;

    this.SetLinkArrowType = (t) => {
        type = t;

        return this;
    };

    this.GetLinkArrowType = () => {
        return type;
    };
}

export default LinkArrow;