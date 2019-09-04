import ElementsRegister from '../elements-register'

function Selection(config) {

    let isSelected = false;
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