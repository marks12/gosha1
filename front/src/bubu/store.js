function Store(config) {

    let Items = {};
    let ctx = null;

    this.AddItem = (element) => {

        let id = element.GetId();
        Items[id] = element;
        return this;
    };

    this.DropElement = (event) => {

        console.log('event', event);

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

        let canvasOffsetX = this.GetCanvasOffsetX();
        let canvasOffsetY = this.GetCanvasOffsetY();

        for (let i in Items) {

            let x1 = Items[i].Coords.GetX();
            let x2 = x1 + Items[i].GetWidth();

            let y1 = Items[i].Coords.GetY();
            let y2 = y1 + Items[i].GetHeight();

            if (x - canvasOffsetX >= x1 && x - canvasOffsetX <= x2 && y - canvasOffsetY >= y1 && y - canvasOffsetY <= y2) {

                this.SetSelectedItemOffsetX(x - canvasOffsetX - x1);
                this.SetSelectedItemOffsetY(y - canvasOffsetY - y1);

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