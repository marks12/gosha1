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

        if (selection) {
            this.RemoveItem(selection)
        }
    };

}

export default Selection;