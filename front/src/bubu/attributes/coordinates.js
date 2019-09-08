function Coordinates(config) {

    let X = config && config.Coords && typeof config.Coords.X === "number" ? config.Coords.X : 0;
    let Y = config && config.Coords && typeof config.Coords.Y === "number" ? config.Coords.Y : 0;

    let PreviousX = 0;
    let PreviousY = 0;

    let self = this;

    this.Coords = {
        GetX () {return X},
        GetY () {return Y},
        GetPreviousX () {return PreviousX},
        GetPreviousY () {return PreviousY},
        SetX (x) {X = x; return self},
        SetY (y) {Y = y; return self},
        AddX (x) {X = PreviousX + x; return self},
        AddY (y) {Y = PreviousY + y; return self},
        SavePrevious () {
            PreviousX = this.GetX();
            PreviousY = this.GetY();
        },
    };
}

export default Coordinates;