import ElementsRegister from './elements-register'
import Toolbox from "./toolbox";
import Store from "./store";
import Renderer from "./renderer";
import SelectItem from "./actions/selectItem";
import Move from "./actions/move";
import Clone from "./actions/clone";
import Selection from "./actions/selection";

function BuBu(canvasElementId) {

    this.canvas = document.getElementById(canvasElementId);

    if (!this.canvas || !this.canvas.getContext) {
        console.error("Wrong canvas element Id. ELement not found or canvas.getContext function not exists");
        return
    }

    Store.apply(this, arguments);

    Move.apply(this, arguments);
    Toolbox.apply(this, arguments);
    Renderer.apply(this, arguments);
    SelectItem.apply(this, arguments);
    Clone.apply(this, arguments);
    Selection.apply(this, arguments);

    let self = this;

    this.GetCanvas = () => {
        return this.canvas;
    };

    this.canvas.setAttribute("width", this.canvas.parentNode.parentElement.clientWidth);
    this.canvas.setAttribute("height", this.canvas.parentNode.parentElement.clientHeight);

    let isDown = false;
    this.selectedItem = null;
    let canvasOffsetX = this.canvas.getBoundingClientRect().left;
    let canvasOffsetY = this.canvas.getBoundingClientRect().top;
    let selectedItemOffsetX = 0;
    let selectedItemOffsetY = 0;

    this.SetCtx(this.canvas.getContext('2d'));

    function down(event) {

        isDown = true;

        event = self.AssignCoordinates(event);

        let x = event.pageX;
        let y = event.pageY;

        self.selectedItem = getFirstElementByCoordinates(x, y);

        if (self.selectedItem) {
            self.selectedItem.SelectSwitch();
            self.Render();
        } else {

            self.CreaetMultiSelection();

        }

    }

    function up() {
        isDown = false;

        if (self.selectedItem) {
            self.selectedItem.Blur();

            self.RemoveMultiSelection();

            self.Render();

        }

        self.selectedItem = null;
        selectedItemOffsetX = 0;
        selectedItemOffsetY = 0;
    }

    let getFirstElementByCoordinates = (x, y) => {

        let Items = this.GetSelectableItems();

        for (let i in Items) {

            let x1 = Items[i].Coords.GetX();
            let x2 = x1 + Items[i].GetWidth();

            let y1 = Items[i].Coords.GetY();
            let y2 = y1 + Items[i].GetHeight();

            if (x - canvasOffsetX >= x1 && x - canvasOffsetX <= x2 && y - canvasOffsetY >= y1 && y - canvasOffsetY <= y2) {

                selectedItemOffsetX = x - canvasOffsetX - x1;
                selectedItemOffsetY = y - canvasOffsetY - y1;

                return Items[i];
            }
        }

        return null;
    };


    this.mover = (event) => {

        event = this.AssignCoordinates(event);

        let newX = event.pageX - canvasOffsetX - selectedItemOffsetX;
        let newY = event.pageY - canvasOffsetY - selectedItemOffsetY;

        if (isDown && self.selectedItem) {

            if (self.selectedItem.GetOnMove()) {

                let m = self.selectedItem.GetOnMove();
                m.Run(newX, newY, self, self.selectedItem);

            } else {

                self.selectedItem.Coords.SetX(newX);
                self.selectedItem.Coords.SetY(newY);
            }

            this.Render();
        }



    };

    this.canvas.addEventListener("mousedown", down);
    this.canvas.addEventListener("mousemove", this.mover);
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
