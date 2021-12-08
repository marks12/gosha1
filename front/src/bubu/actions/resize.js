function Resize(childElement) {

    this.Run = (x, y, sItem, root) => {

        if (sItem.Coords.GetX() === 0 && sItem.Coords.GetY() === 0) {
            sItem.Coords.SetX(x);
            sItem.Coords.SetY(y);
        }

        sItem.SetWidth(x - sItem.Coords.GetX());
        sItem.SetHeight(y - sItem.Coords.GetY());

    };
}

export default Resize;