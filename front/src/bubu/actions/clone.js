function Clone(childElement) {

    this.Run = (x, y, root) => {

        let el = new childElement().Coords.SetX(x).Coords.SetY(y);
        root.AddItem(el);
        root.setSelectedItem(el);
    };
}

export default Clone;