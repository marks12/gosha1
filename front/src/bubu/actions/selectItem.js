import {TYPES} from "../constants";

function SelectItem() {

    let offsetX = 0;
    let offsetY = 0;
    let selectedItem = null;

    this.SetSelectedItem = (item) => {

        selectedItem = item;
        return this;
    };

    this.GetSelectedItem = () => {
        return selectedItem;
    };

    this.GetSelectedItemOffsetX = () => {
        return offsetX;
    };

    this.SetSelectedItemOffsetX = (x) => {
        offsetX = x;
        return this;
    };

    this.GetSelectedItemOffsetY = () => {
        return offsetY;
    };

    this.SetSelectedItemOffsetY = (y) => {
        offsetY = y;
        return this;
    };

    this.ClearSelectedItem = () => {
        offsetX = 0;
        offsetY = 0;
        selectedItem = null;
    };

    this.ClearConnectors = () => {

        let items = this.GetSelectableItems();

        for (let i in items) {
            items[i].HideConnectors();
        }
    };

    this.ClearMusltiSelection = () => {

        let items = this.GetItems();

        for (let i in items) {

            if (items[i].GetType() === TYPES.multiselection) {
                this.RemoveItem(items[i]);
            }
        }
    };
}

export default SelectItem