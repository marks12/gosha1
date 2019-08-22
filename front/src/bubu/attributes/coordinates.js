function Coordinates(child) {

    let X = 0;
    let Y = 0;

    return {
        GetX: () => {return X},
        GetY: () => {return Y},
        SetX: (x) => {X = x; return this},
        SetY: (y) => {Y = y; return this},
    };
}

export default Coordinates;