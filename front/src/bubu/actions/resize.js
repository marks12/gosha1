function Resize(childElement) {

    this.Run = (x, y, root) => {

        if (root.selectedItem.Coords.GetX() === 0 && root.selectedItem.Coords.GetY() === 0) {
            root.selectedItem.Coords.SetX(x);
            root.selectedItem.Coords.SetY(y);
        }

        root.selectedItem.SetWidth(x - root.selectedItem.Coords.GetX());
        root.selectedItem.SetHeight(y - root.selectedItem.Coords.GetY());

        console.log('resize', x, y);
    };
}

export default Resize;