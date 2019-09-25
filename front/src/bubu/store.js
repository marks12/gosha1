import ElementsRegister from "./elements-register";
import {TYPES as constants} from "./constants";

function Store(config) {

    let Items = {};
    let ctx = null;
    let self = this;

    this.AddItem = (element) => {

        let id = element.GetId();
        Items[id] = element;
        return this;
    };

    let addPoints = (item) => {

        for (let i=0; i<4;i++) {
            item.AddConnectorPoint(i, i === 1);
        }

        return item;
    };

    let resetGetText = (item) => {

        let previosGetText = item.GetText;
        item.GetText = () => {
            return " x:" + item.Coords.GetX() + "\n" +
                " y:" + item.Coords.GetY();
        };
        return item;
    };

    this.DropElement = (event) => {

        event.preventDefault();

        event = this.Mouse.AssignCoordinates(event);

        let elementType = event.srcElement.getAttribute('data-bubu');

        let w = 100;
        let h = w;
        let x = this.GetCanvasX(event.pageX || event.onMoveCoords.x)  - w / 2;
        let y = this.GetCanvasY(event.pageY || event.onMoveCoords.y) - h / 2;

        let item = null;

        switch (elementType * 1) {

            case constants.task:

                item = new ElementsRegister.Task({
                    Width: w,
                    Height: h,
                    Coords: {
                        X: x,
                        Y: y,
                    },
                    Text: `Task`
                });

                this.AddItem(item);

                break;

            case constants.condition:

                item = new ElementsRegister.Condition({
                    Width: w,
                    Height: h,
                    Coords: {
                        X: x,
                        Y: y,
                    },
                    Text: `Condition`
                });

                this.AddItem(item);

                break;

            default:

                console.error('unknown drop type', elementType);
                return false;
        }

        addPoints(item);
        resetGetText(item);

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
        let els = [];

        for (let i in Items) {
            els.push(Items[i]);
        }

        return els;
    };

    this.GetSelectableItems = () => {

        let els = [];

        for (let i in Items) {

            if (! Items[i].IsSelectable()) {
                continue;
            }

            if (Items[i].GetVisibility && ! Items[i].GetVisibility()) {
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

    this.GetFirstElementByCoordinates = (x, y, spaceX, spaceY) => {

        let Items = this.GetSelectableItems();

        for (let i in Items) {

            let x1 = Items[i].Coords.GetX();
            let x2 = x1 + Items[i].GetWidth();

            let y1 = Items[i].Coords.GetY();
            let y2 = y1 + Items[i].GetHeight();

            if (spaceX > 0) {
                x1 -= spaceX;
                x2 += spaceX;
            }

            if (spaceY > 0) {
                y1 -= spaceY;
                y2 += spaceY;
            }

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

        if (! item) {
            this.ClearSelectedItem();
        }

        this.SetSelectedItem(item);
        return item;
    };

}

export default Store;