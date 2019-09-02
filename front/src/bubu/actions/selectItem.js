function SelectItem() {

    this.setSelectedItem = (item) => {

        this.selectedItem = item;
        return this;
    };

    this.getSelectedItem = () => {
        return selectedItem;
    };
}

export default SelectItem