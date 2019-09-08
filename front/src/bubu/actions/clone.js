function Clone(childElement) {

    this.Run = (x, y, root) => {

        let el = new childElement().Coords.SetX(x).Coords.SetY(y);

        let items = root.GetSelectedItems();

        for (let i in items) {
            items[i].Blur();
        }

        root.selectedItem.Blur();

        el.Blur();
        root.AddItem(el);
        root.setSelectedItem(el);
    };
}

export default Clone;