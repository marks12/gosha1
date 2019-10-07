import Coordinates from "../attributes/coordinates";

function Keyboard(config) {

    Coordinates.apply(this, arguments);

    let self = this;

    this.Keyboard = new KeyboardDevice();

    let code = null;

    function KeyboardDevice() {

        this.Keypress = (event) => {

            if (event.key === "Delete") {
                let items = self.GetSelectedItems();
                for (let i in items) {

                    let del = items[i].GetOnDelete();

                    if (del) {
                        del(event);
                    }

                    self.RemoveItem(items[i]);
                }

                self.Render();

                console.log(self.GetItems());
            }

            event.preventDefault();
        };

    }
}

export default Keyboard;