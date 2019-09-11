import Coordinates from "../attributes/coordinates";
import IsRightButton from "../common";

function Mouse(config) {

    Coordinates.apply(this, arguments);

    let self = this;

    this.Mouse = new MouseDevice();

    function MouseDevice() {

        let clickCoordsX = 0;
        let clickCoordsY = 0;
        let isDown = false;

        this.IsDown = () => {
            return isDown;
        };

        this.Down = (event) => {

            event = this.AssignCoordinates(event);

            if (IsRightButton(event)) {
                return false;
            }

            isDown = true;

            let x = event.pageX;
            let y = event.pageY;

            clickCoordsX = x - self.GetCanvasOffsetX();
            clickCoordsY = y - self.GetCanvasOffsetY();

            let selectedItems = self.GetSelectedItems();

            let sItem = self.SelectItemByCoordinates(x, y);

            if (sItem) {

                if (selectedItems.length === 1) {
                    self.BlurAll();
                }

                if (sItem.IsSelectable()) {
                    sItem.Select();
                }

            } else {
                self.CreateMultiSelection(x - self.GetCanvasOffsetX(), y - self.GetCanvasOffsetY());
                self.BlurAll();
            }
        };

        this.Up = (event) => {

            isDown = false;

            self.RemoveMultiSelection();

            self.ClearSelectedItem();
        };

        this.Move = (event) => {

            event = this.AssignCoordinates(event);


            let newX = event.pageX - self.GetCanvasOffsetX() - self.GetSelectedItemOffsetX();
            let newY = event.pageY - self.GetCanvasOffsetY() - self.GetSelectedItemOffsetY();

            // console.log('event.pageX', event.pageX, 'newX', newX, 'self.GetCanvasOffsetX()', self.GetSelectedItemOffsetX());

            let sItem = self.GetSelectedItem();
            
            if (self.Mouse.IsDown() && sItem) {

                if (sItem.GetOnMove()) {

                    let m = sItem.GetOnMove();
                    m.Run(newX, newY, sItem, self);

                } else {

                    sItem.Coords.SetX(newX);
                    sItem.Coords.SetY(newY);
                }

                let items = self.GetSelectedItems();
                for (let i in items) {

                    if (items[i].GetId() === sItem.GetId()) {
                        continue;
                    }

                    let offsetX = event.pageX - clickCoordsX - self.GetCanvasOffsetX();
                    let offsetY = event.pageY - clickCoordsY - self.GetCanvasOffsetY();

                    if (items[i].GetOnMove()) {

                        let m = sItem.GetOnMove();
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