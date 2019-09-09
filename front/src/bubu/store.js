function Store(config) {

    let Items = {};
    let ctx = null;
    let selectedItem = null;
    let selectedItemOffsetX = 0;
    let selectedItemOffsetY = 0;

    this.GetSelectedItem = () => {
        return selectedItem;
    };

    this.SetSelectedItem = (item) => {
        selectedItem = item;
        return this;
    };

    this.AddItem = (element) => {

        let id = element.GetId();
        Items[id] = element;
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

            if (x - canvasOffsetX >= x1 && x - canvasOffsetX <= x2 && y - canvasOffsetY >= y1 && y - canvasOffsetY <= y2) {

                selectedItemOffsetX = x - canvasOffsetX - x1;
                selectedItemOffsetY = y - canvasOffsetY - y1;

                return Items[i];
            }
        }

        return null;
    };

}

export default Store;