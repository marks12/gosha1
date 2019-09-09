import ElementsRegister from './elements-register'
import Toolbox from "./toolbox";
import WorkArea from "./work-area";
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

    this.ZeroPoint = new ElementsRegister.ZeroPoint();

    Store.apply(this, arguments);

    Move.apply(this, arguments);
    Toolbox.apply(this, arguments);
    WorkArea.apply(this, arguments);
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
    let clickCoordsX = 0;
    let clickCoordsY = 0;

    this.SetCtx(this.canvas.getContext('2d'));

    function isRightButton(e) {

        let isRightMB;
        e = e || window.event;

        if ("which" in e)  // Gecko (Firefox), WebKit (Safari/Chrome) & Opera
            isRightMB = e.which === 3;
        else if ("button" in e)  // IE, Opera
            isRightMB = e.button === 2;

        return isRightMB;
    }

    function down(event) {

        if (isRightButton(event)) {
            return false;
        }

        isDown = true;

        event = self.AssignCoordinates(event);

        let x = event.pageX;
        let y = event.pageY;

        clickCoordsX = x - canvasOffsetX;
        clickCoordsY = y - canvasOffsetY;

        let selectedItems = self.GetSelectedItems();

        self.selectedItem = getFirstElementByCoordinates(x, y);

        if (self.selectedItem) {

            if (selectedItems.length === 1) {
                self.BlurAll();
            }

            if (self.selectedItem.IsSelectable()) {
                self.selectedItem.Select();
            }

        } else {
            self.CreateMultiSelection(x - canvasOffsetX, y - canvasOffsetY);
            self.BlurAll();
        }

        self.Render();
    }

    function up() {

        isDown = false;

        self.RemoveMultiSelection();

        self.selectedItem = null;

        selectedItemOffsetX = 0;
        selectedItemOffsetY = 0;

        self.Render();

    }

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


            let items = this.GetSelectedItems();
            for (let i in items) {

                if (items[i].GetId() === this.selectedItem.GetId()) {
                    continue;
                }

                let offsetX = event.pageX - clickCoordsX - canvasOffsetX;
                let offsetY = event.pageY - clickCoordsY - canvasOffsetY;

                if (items[i].GetOnMove()) {

                    let m = self.selectedItem.GetOnMove();
                    m.Run(newX, newY, self, items[i]);

                } else {

                    items[i].Coords.SetX(items[i].Coords.GetPreviousX() + offsetX);
                    items[i].Coords.SetY(items[i].Coords.GetPreviousY() + offsetY);
                }
            }

            this.Render();
        }
    };


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
