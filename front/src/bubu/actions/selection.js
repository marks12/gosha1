import ElementsRegister from '../elements-register'

function Selection(config) {

    let isSelected = false;

    let isSelectable = config && config.IsSelectable !== undefined ? !!(config.IsSelectable) : true;

    let selection = null;

    this.Select = () => {
        isSelected = true;
        return this;
    };

    this.GetSelection = () => {
        return selection;
    };

    this.Blur = () => {
        isSelected = false;
        return this;
    };

    this.BlurAll = () => {
        let items = this.GetSelectableItems();

        for (let i in items) {
            items[i].Blur();
        }
        return this;
    };

    this.SelectSwitch = () => {
        isSelected = ! isSelected;
        return this;
    };

    this.IsSelected = () => {
        return isSelected;
    };

    this.SetSelectable = (b) => {
        isSelectable = b;
        return this;
    };

    this.IsSelectable = () => {
        return isSelectable;
    };

    this.CreateMultiSelection = (x, y) => {

        selection = new ElementsRegister.MultiSelection();
        selection.SetOnMove(new ElementsRegister.Actions.Resize());
        selection.Coords.SetX(x);
        selection.Coords.SetY(y);
        selection.SetWidth(0);
        selection.SetHeight(0);

        this.AddItem(selection);
        this.SetSelectedItem(selection);
    };

    this.SavePreviousAll = () => {

        let items = this.GetSelectableItems();
        for (let i in items) {
            items[i].Coords.SavePrevious();
        }
    };

    this.RemoveMultiSelection = () => {

        let items = this.GetSelectableItems();

        console.log('selection', selection);

        if (selection) {

            for (let i in items) {

                if (
                    (
                        selection.Coords.GetX() <= items[i].Coords.GetX() &&
                        selection.Coords.GetY() <= items[i].Coords.GetY() &&
                        selection.GetWidth() > 0 &&
                        selection.GetHeight() > 0 &&
                        selection.Coords.GetX() + selection.GetWidth() >= items[i].Coords.GetX() + items[i].GetWidth() &&
                        selection.Coords.GetY() + selection.GetHeight() >= items[i].Coords.GetY() + items[i].GetHeight()
                    )
                    ||
                    (
                        selection.Coords.GetX() >= items[i].Coords.GetX() &&
                        selection.Coords.GetY() <= items[i].Coords.GetY() &&
                        selection.GetWidth() < 0 &&
                        selection.GetHeight() > 0 &&
                        selection.Coords.GetX() + selection.GetWidth() <= items[i].Coords.GetX() &&
                        selection.Coords.GetY() + selection.GetHeight() >= items[i].Coords.GetY() + items[i].GetHeight()
                    )
                    ||
                    (
                        selection.Coords.GetX() >= items[i].Coords.GetX() &&
                        selection.Coords.GetY() >= items[i].Coords.GetY() &&
                        selection.GetWidth() < 0 &&
                        selection.GetHeight() < 0 &&
                        selection.Coords.GetX() + selection.GetWidth() <= items[i].Coords.GetX() &&
                        selection.Coords.GetY() + selection.GetHeight() <= items[i].Coords.GetY()
                    )
                    ||
                    (
                        selection.Coords.GetX() <= items[i].Coords.GetX() &&
                        selection.Coords.GetY() >= items[i].Coords.GetY() &&
                        selection.GetWidth() > 0 &&
                        selection.GetHeight() < 0 &&
                        selection.Coords.GetX() + selection.GetWidth() >= items[i].Coords.GetX() + items[i].GetWidth() &&
                        selection.Coords.GetY() + selection.GetHeight() <= items[i].Coords.GetY() + items[i].GetHeight()
                    )
                ) {
                    items[i].Select();
                } else {
                    items[i].Blur();
                }

                items[i].Coords.SavePrevious();
            }

            this.RemoveItem(selection);
            selection = null;

        }

        this.SavePreviousAll();
    };


}

export default Selection;