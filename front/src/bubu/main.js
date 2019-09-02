import Elements from './elements'
import Toolbox from "./toolbox";
import Store from "./store";
import Renderer from "./renderer";
import SelectItem from "./actions/selectItem";
import Move from "./actions/move";
import Clone from "./actions/clone";

function BuBu(canvasElementId) {

    let canvas = document.getElementById(canvasElementId);

    if (!canvas || !canvas.getContext) {
        console.error("Wrong canvas element Id. ELement not found or canvas.getContext function not exists");
        return
    }

    Store.apply(this, arguments);

    Move.apply(this, arguments);
    Toolbox.apply(this, arguments);
    Renderer.apply(this, arguments);
    SelectItem.apply(this, arguments);
    Clone.apply(this, arguments);

    let self = this;

    this.GetCanvas = () => {
        return canvas;
    };

    canvas.setAttribute("width", canvas.parentNode.parentElement.clientWidth);
    canvas.setAttribute("height", canvas.parentNode.parentElement.clientHeight);

    let isDown = false;
    this.selectedItem = null;
    let canvasOffsetX = canvas.getBoundingClientRect().left;
    let canvasOffsetY = canvas.getBoundingClientRect().top;
    let selectedItemOffsetX = 0;
    let selectedItemOffsetY = 0;

    this.SetCtx(canvas.getContext('2d'));

    function down(event) {

        isDown = true;

        event = self.AssignCoordinates(event);

        let x = event.pageX;
        let y = event.pageY;

        self.selectedItem = getFirstElementByCoordinates(x, y);
    }

    function up() {
        isDown = false;
        self.selectedItem = null;
        selectedItemOffsetX = 0;
        selectedItemOffsetY = 0;
    }

    let getFirstElementByCoordinates = (x, y) => {

        let Items = this.GetItems();

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

            if (self.selectedItem.getOnMove()) {

                let m = self.selectedItem.getOnMove();
                console.log(m);
                m.Run(newX, newY, self);


            } else {

                self.selectedItem.Coords.SetX(newX);
                self.selectedItem.Coords.SetY(newY);
            }

            this.Render();
        }
    };

    canvas.addEventListener("mousedown", down);
    canvas.addEventListener("mousemove", this.mover);
    canvas.addEventListener("mouseup", up);

    return {
        Add: this.AddItem,
        Elements: Elements,
        GetNames: this.GetNames,
        GetItemsByType: this.GetItemsByType,
        GetItemsByName: this.GetItemsByName,
        GetItemById: this.GetItemById,
        Render: this.Render,
    };
}

export default BuBu;
