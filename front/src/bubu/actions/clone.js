function Clone(childElement) {

    this.Run = (x, y, root) => {

        let el = new childElement().Coords.SetX(x).Coords.SetY(y);

        let items = root.GetSelectedItems();

        for (let i in items) {
            items[i].Blur();
        }

        root.GetSelectedItem().Blur();

        el.SetText('x:' + el.Coords.GetX() + ' y:' + el.Coords.GetY());

        el.Blur();
        root.AddItem(el);
        root.SetSelectedItem(el);
    };
}

export default Clone;