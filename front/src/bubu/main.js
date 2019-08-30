import Elements from './elements'
import Toolbox from "./toolbox";
import Store from "./store";
import Renderer from "./renderer";

function BuBu(canvasElementId) {

    let self = this;

    let canvas = document.getElementById(canvasElementId);

    if (!canvas || !canvas.getContext) {
        console.error("Wrong canvas element Id. ELement not found or canvas.getContext function not exists");
        return
    }

    Store.apply(this, arguments);

    Toolbox.apply(this, arguments);
    Renderer.apply(this, arguments);

    this.GetCanvas = () => {
        return canvas;
    };

    canvas.setAttribute("width", canvas.parentNode.parentElement.clientWidth);
    canvas.setAttribute("height", canvas.parentNode.parentElement.clientHeight);

    let isDown = false;
    let selectedItem = null;
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

        selectedItem = getFirstElementByCoordinates(x, y);
    }

    function up() {
        isDown = false;
        selectedItem = null;
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

        if (isDown && selectedItem) {

            selectedItem.Coords.SetX(event.pageX - canvasOffsetX - selectedItemOffsetX);
            selectedItem.Coords.SetY(event.pageY - canvasOffsetY - selectedItemOffsetY);

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
