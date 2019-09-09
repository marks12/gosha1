import ElementsRegister from './elements-register'
import Toolbox from "./toolbox";
import WorkArea from "./work-area";
import Store from "./store";
import Renderer from "./renderer";
import SelectItem from "./actions/selectItem";
import Move from "./actions/move";
import Clone from "./actions/clone";
import Selection from "./actions/selection";
import IsRightButton from "./common";
import Mouse from "./device/mouse";
import Canvas from "./canvas";

function BuBu(canvasElementId) {

    this.canvas = new Canvas(canvasElementId);

    if (!this.canvas || !this.canvas.getContext) {
        console.error("Wrong canvas element Id. ELement not found or canvas.getContext function not exists");
        return
    }

    this.ZeroPoint = new ElementsRegister.ZeroPoint();

    Store.apply(this, arguments);
    Move.apply(this, arguments);
    Toolbox.apply(this, arguments);
    WorkArea.apply(this, arguments);
    Renderer.apply(this, arguments);
    SelectItem.apply(this, arguments);
    Clone.apply(this, arguments);
    Selection.apply(this, arguments);
    Mouse.apply(this, arguments);

    let self = this;

    this.GetCanvas = () => {
        return this.canvas;
    };


    this.SetCtx(this.canvas.getContext('2d'));

    function down(event) {

        self.Mouse.Down(event);
        self.Render();
    }

    function up(event) {

        self.Mouse.Up(event);
        self.Render();

    }

    function mover (event) {

        this.Mouse.Move(event);
        self.Render();

    }

    this.canvas.addEventListener("mousedown", down);
    this.canvas.addEventListener("mousemove", mover);
    this.canvas.addEventListener("mouseup", up);

    return {
        Add: this.AddItem,
        Elements: ElementsRegister,
        GetNames: this.GetNames,
        GetItemsByType: this.GetItemsByType,
        GetItemsByName: this.GetItemsByName,
        GetItemById: this.GetItemById,
        Render: this.Render,
    };
}

export default BuBu;
