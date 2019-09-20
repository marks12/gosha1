import LinkDirection from "../attributes/link-direction";
import LinkArrow from "../attributes/link-arrow";
import LinkLine from "../attributes/link-line";

function Link(config) {

    LinkArrow.apply(this, arguments);
    LinkLine.apply(this, arguments);
    LinkDirection.apply(this, arguments);

}

export default Link;