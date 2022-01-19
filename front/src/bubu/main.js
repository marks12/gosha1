import ElementsRegister from './elements-register'
import Toolbox from "./toolbox";
import WorkArea from "./work-area";
import Store from "./store";
import Renderer from "./renderer";
import SelectItem from "./actions/selectItem";
import Move from "./actions/move";
import Clone from "./actions/clone";
import Selection from "./actions/selection";
import Keyboard from "./device/keyboard";
import Mouse from "./device/mouse";
import Canvas from "./canvas";
import SrcImage from "./src-image";
import Connectors from "./attributes/connectors";
import Links from "./attributes/links";
import Buttons from "./attributes/buttons";

function BuBu(canvasElementId, ToolboxElementId) {

    Store.apply(this, arguments);
    Mouse.apply(this, arguments);
    Keyboard.apply(this, arguments);
    Canvas.apply(this, arguments);
    SrcImage.apply(this, arguments);
    Connectors.apply(this, arguments);
    Links.apply(this, arguments);
    Buttons.apply(this, arguments);

    if (!this.GetCanvas() || !this.GetContext) {
        console.error("Wrong canvas element Id. Element not found or canvas.getContext function not exists");
        return
    }

    Move.apply(this, arguments);
    Toolbox.apply(this, arguments);
    WorkArea.apply(this, arguments);
    Renderer.apply(this, arguments);
    SelectItem.apply(this, arguments);
    Clone.apply(this, arguments);
    Selection.apply(this, arguments);

    return {
        Add: this.AddItem,
        Elements: ElementsRegister,
        GetNames: this.GetNames,
        GetItemsByType: this.GetItemsByType,
        GetItemsByName: this.GetItemsByName,
        GetItemById: this.GetItemById,
        Render: this.Render,
        UpdateToolbox: this.UpdateToolbox,
        ResetCanvas: this.ResetCanvas,
        GetSrcImageTask: this.GetSrcImageTask,
        GetSrcImageCondition: this.GetSrcImageCondition,
        UpdateCanvas: this.UpdateCanvas,
        GetData: this.GetStore,
    };
}

export default BuBu;
