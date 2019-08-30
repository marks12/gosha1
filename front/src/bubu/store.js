function Store(config) {

    let Items = {};
    let ctx = null;

    this.AddItem = (element) => {

        let id = element.GetId();
        Items[id] = element;
        return this;
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

        for (let i = 0; i < Items.length; i++) {

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
    }

}

export default Store;