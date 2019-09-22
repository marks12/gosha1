import Coordinates from "../attributes/coordinates";

function Keyboard(config) {

    Coordinates.apply(this, arguments);

    let self = this;

    this.Keyboard = new KeyboardDevice();

    let code = null;

    function KeyboardDevice() {

        this.Keyup = (event) => {

        };

    }
}

export default Keyboard;