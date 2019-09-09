import Coordinates from "../attributes/coordinates";
import IsRightButton from "../common";

function Mouse(config) {

    Coordinates.apply(this, arguments);

    self = this;

    this.Mouse = new this.MouseDevice();

    this.MouseDevice = function () {

        let clickCoordsX = 0;
        let clickCoordsY = 0;
        let isDown = false;

        this.IsDown = () => {
            return isDown;
        };

        this.Down = (event) => {

            if (IsRightButton(event)) {
                return false;
            }

            isDown = true;

            event = self.AssignCoordinates(event);

            let x = event.pageX;
            let y = event.pageY;

            clickCoordsX = x - canvasOffsetX;
            clickCoordsY = y - canvasOffsetY;

            let selectedItems = self.GetSelectedItems();

            self.selectedItem = self.GetFirstElementByCoordinates(x, y);

            if (self.selectedItem) {

                if (selectedItems.length === 1) {
                    self.BlurAll();
                }

                if (self.selectedItem.IsSelectable()) {
                    self.selectedItem.Select();
                }

            } else {
                self.CreateMultiSelection(x - canvasOffsetX, y - canvasOffsetY);
                self.BlurAll();
            }
        };

        this.Up = (event) => {

            isDown = false;

            self.RemoveMultiSelection();

            self.selectedItem = null;

            selectedItemOffsetX = 0;
            selectedItemOffsetY = 0;
        };

        this.Move = (event) => {

            event = this.AssignCoordinates(event);

            let newX = event.pageX - canvasOffsetX - selectedItemOffsetX;
            let newY = event.pageY - canvasOffsetY - selectedItemOffsetY;

            if (self.Mouse.IsDown() && self.GetSelectedItem()) {

                if (self.selectedItem.GetOnMove()) {

                    let m = self.selectedItem.GetOnMove();
                    m.Run(newX, newY, self, self.selectedItem);

                } else {

                    self.selectedItem.Coords.SetX(newX);
                    self.selectedItem.Coords.SetY(newY);
                }


                let items = this.GetSelectedItems();
                for (let i in items) {

                    if (items[i].GetId() === this.selectedItem.GetId()) {
                        continue;
                    }

                    let offsetX = event.pageX - clickCoordsX - canvasOffsetX;
                    let offsetY = event.pageY - clickCoordsY - canvasOffsetY;

                    if (items[i].GetOnMove()) {

                        let m = self.selectedItem.GetOnMove();
                        m.Run(newX, newY, self, items[i]);

                    } else {

                        items[i].Coords.SetX(items[i].Coords.GetPreviousX() + offsetX);
                        items[i].Coords.SetY(items[i].Coords.GetPreviousY() + offsetY);
                    }
                }

            }

        };

        this.AssignCoordinates = (e) => {

            let eventDoc, doc, body;

            event = e || window.event; // IE-ism

            if (event.pageX == null && event.clientX != null) {
                eventDoc = (event.target && event.target.ownerDocument) || document;
                doc = eventDoc.documentElement;
                body = eventDoc.body;

                event.pageX = event.clientX +
                    (doc && doc.scrollLeft || body && body.scrollLeft || 0) -
                    (doc && doc.clientLeft || body && body.clientLeft || 0);
                event.pageY = event.clientY +
                    (doc && doc.scrollTop  || body && body.scrollTop  || 0) -
                    (doc && doc.clientTop  || body && body.clientTop  || 0 );

            }

            return event;
        }


    };
}

export default Mouse;