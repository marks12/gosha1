import Action from "./elements/action";
import Condition from "./elements/condition";

function Elements () {

    this.Action = Action;
    this.Condition = Condition;

    return this;

}

export default new Elements();