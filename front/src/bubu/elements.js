import Action from "./elements/action";
import Condition from "./elements/condition";

function Elements () {

    this.CrAction = Action;
    this.CrCondition = Condition;

    return this;

}

export default new Elements();