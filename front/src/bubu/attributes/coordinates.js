function Coordinates(config) {

    let X = config && config.Coords && typeof config.Coords.X === "number" ? config.Coords.X : 0;
    let Y = config && config.Coords && typeof config.Coords.Y === "number" ? config.Coords.Y : 0;

    let self = this;

    this.Coords = {
        GetX () {return X},
        GetY () {return Y},
        SetX (x) {X = x; return self},
        SetY (y) {Y = y; return self},
    };
}

export default Coordinates;