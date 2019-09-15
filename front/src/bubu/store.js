import ElementsRegister from "./elements-register";
import {TYPES as constants} from "./constants";

function Store(config) {

    let Items = {};
    let ctx = null;

    this.AddItem = (element) => {

        let id = element.GetId();
        Items[id] = element;
        return this;
    };

    this.DropElement = (event) => {

        let elementType = event.srcElement.getAttribute('data-bubu');

        let w = 100;
        let h = w;
        let x = this.GetCanvasX(event.pageX) - w / 2;
        let y = this.GetCanvasX(event.pageY) + h / 2;

        switch (elementType * 1) {

            case constants.TASK:

                this.AddItem((new ElementsRegister.Task({
                    Width: w,
                    Height: h,
                    Coords: {
                        X: x,
                        Y: y,
                    },
                    Text: `x: ${x} y: ${y}`
                })));

                break;

            case constants.CONDITION:

                this.AddItem((new ElementsRegister.Condition({
                    Width: w,
                    Height: h,
                    Coords: {
                        X: x,
                        Y: y,
                    },
                    Text: `x: ${x} y: ${y}`
                })));

                break;

            default:

                console.error('unknown drop type', elementType);

                break;
        }

        this.Render();

        return this;
    };

    this.RemoveItem = (element) => {

        for (let i in Items) {
            if (Items[i].GetId() === element.GetId()) {
                delete Items[i];
                return;
            }
        }

    };

    this.GetNames = () => {

        let names = [];

        for (let i in Items) {
            names.push(Items[i].GetName());
        }

        return names;
    };

    this.GetItemsByName = (name) => {

        let els = [];

        for (let i in Items) {

            if (Items[i].GetName() === name) {
                els.push(Items[i]);
            }
        }
        return els;
    };

    this.GetItemsByType = (type) => {

        let els = [];

        for (let i in Items) {

            if (Items[i].GetType() === type) {
                els.push(Items[i]);
            }
        }

        return els;
    };

    this.GetItemById = (id) => {

        return Items[id];
    };

    this.SetCtx = (c) => {
        ctx = c;
        return this;
    };

    this.GetCtx = () => {
        return ctx;
    };

    this.GetItems = () => {
        return Items;
    };

    this.GetSelectableItems = () => {

        let els = [];

        for (let i in Items) {

            if (! Items[i].IsSelectable()) {
                continue;
            }

            els.push(Items[i]);
        }

        return els;
    };

    this.GetSelectedItems = () => {

        let els = [];

        for (let i in Items) {

            if (! Items[i].IsSelected()) {
                continue;
            }

            els.push(Items[i]);
        }

        return els;
    };

    this.GetFirstElementByCoordinates = (x, y) => {

        let Items = this.GetSelectableItems();

        for (let i in Items) {

            let x1 = Items[i].Coords.GetX();
            let x2 = x1 + Items[i].GetWidth();

            let y1 = Items[i].Coords.GetY();
            let y2 = y1 + Items[i].GetHeight();

            if (x >= x1 && x <= x2 && y >= y1 && y <= y2) {

                this.SetSelectedItemOffsetX(x - x1);
                this.SetSelectedItemOffsetY(y - y1);

                return Items[i];
            }
        }

        return null;
    };

    this.SelectItemByCoordinates = (x, y) => {

        let item = this.GetFirstElementByCoordinates(x, y);
        this.SetSelectedItem(item);
        return item;
    };

}

export default Store;