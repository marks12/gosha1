import ElementsRegister from '../elements-register'

function Selection(config) {

    let isSelected = false;

    let isSelectable = config && config.IsSelectable !== undefined ? !!(config.IsSelectable) : true;

    let selection = null;

    this.Select = () => {
        isSelected = true;
        return this;
    };

    this.Blur = () => {
        isSelected = false;
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

    this.CreaetMultiSelection = () => {

        selection = new ElementsRegister.MultiSelection();
        selection.SetOnMove(new ElementsRegister.Actions.Resize());

        this.AddItem(selection);
        this.setSelectedItem(selection);
    };

    this.RemoveMultiSelection = () => {

        let items = this.GetSelectableItems();
        let ctx = this.GetCtx();

        if(selection) {
            console.log('selection.Coords',
                selection.Coords.GetX(),
                selection.Coords.GetY(),
                selection.GetWidth(),
                selection.GetHeight(),
            );

        }

        for (let i in items) {

            if (selection &&
                (
                    (
                        selection.Coords.GetX() <= items[i].Coords.GetX() &&
                        selection.Coords.GetY() <= items[i].Coords.GetY() &&
                        selection.GetWidth() > 0 && selection.Coords.GetX() + selection.GetWidth() >= items[i].Coords.GetX() + items[i].GetWidth() &&
                        selection.GetHeight() > 0 && selection.Coords.GetY() + selection.GetHeight() >= items[i].Coords.GetY() + items[i].GetHeight()
                    )
                        ||
                    (
                        selection.Coords.GetX() >= items[i].Coords.GetX() &&
                        selection.Coords.GetY() <= items[i].Coords.GetY() &&
                        selection.GetWidth() < 0 && selection.Coords.GetX() + selection.GetWidth() <= items[i].Coords.GetX() + items[i].GetWidth() &&
                        selection.GetHeight() > 0 && selection.Coords.GetY() + selection.GetHeight() >= items[i].Coords.GetY() + items[i].GetHeight()
                    )
                        ||
                    (
                        selection.Coords.GetX() >= items[i].Coords.GetX() &&
                        selection.Coords.GetY() >= items[i].Coords.GetY() &&
                        selection.GetWidth() < 0 && selection.Coords.GetX() + selection.GetWidth() <= items[i].Coords.GetX() + items[i].GetWidth() &&
                        selection.GetHeight() < 0 && selection.Coords.GetY() + selection.GetHeight() <= items[i].Coords.GetY() + items[i].GetHeight()
                    )
                        ||
                    (
                        selection.Coords.GetX() <= items[i].Coords.GetX() &&
                        selection.Coords.GetY() >= items[i].Coords.GetY() &&
                        selection.GetWidth() > 0 && selection.Coords.GetX() + selection.GetWidth() >= items[i].Coords.GetX() + items[i].GetWidth() &&
                        selection.GetHeight() < 0 && selection.Coords.GetY() + selection.GetHeight() <= items[i].Coords.GetY() + items[i].GetHeight()
                    )
                )
            ) {
                items[i].Select();
            } else {
                items[i].Blur();
            }

        }

        if (selection) {
            this.RemoveItem(selection)
            selection = null;
        }
    };

}

export default Selection;