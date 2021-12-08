import {TYPES as constants} from '../constants';

function Coordinates(config) {

    let X = config && config.Coords && typeof config.Coords.X === "number" ? config.Coords.X : 1;
    let Y = config && config.Coords && typeof config.Coords.Y === "number" ? config.Coords.Y : 1;

    let PreviousX = 1;
    let PreviousY = 1;

    let self = this;

    let roundCoord = (x) => {

        if (self.GetType) {
            switch (self.GetType()) {

                case constants.task:
                case constants.condition:
                    return Math.round(x / constants.roundCoords) * constants.roundCoords;
                    break;
            }
        }


        return x;
    };

    this.Coords = {
        /**
         * @return {number}
         */
        GetX () {return X},

        /**
         * @return {number}
         */
        GetY () {return Y},

        /**
         * @return {number}
         */
        GetPreviousX () {return PreviousX},

        /**
         * @return {number}
         */
        GetPreviousY () {return PreviousY},

        SetX (x) {
            X = roundCoord(x); return self
        },
        SetY (y) {
            Y = roundCoord(y); return self
        },
        AddX (x) {X = PreviousX + x; return self},
        AddY (y) {Y = PreviousY + y; return self},
        SavePrevious () {
            PreviousX = this.GetX();
            PreviousY = this.GetY();
        },
    };

}

export default Coordinates;