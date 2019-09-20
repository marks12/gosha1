import Coordinates from "../attributes/coordinates";
import {IsRightButton} from "../common";
import {IsMiddleButtton} from "../common";

function Mouse(config) {

    Coordinates.apply(this, arguments);

    let self = this;

    this.Mouse = new MouseDevice();

    let onMoveX = 0;
    let onMoveY = 0;

    function MouseDevice() {

        let clickCoordsX = 0;
        let clickCoordsY = 0;
        let isDown = false;

        this.IsDown = () => {
            return isDown;
        };

        this.Down = (event) => {

            event = this.AssignCoordinates(event);

            if (event.altKey) {
                console.log("create alt key");
                return false;
            }

            if (IsRightButton(event)) {
                return false;
            }

            if (IsMiddleButtton(event)) {
                isDown = true;
                return true;
            }

            isDown = true;

            clickCoordsX = self.GetCanvasX(event.pageX);
            clickCoordsY = self.GetCanvasY(event.pageY);

            let selectedItems = self.GetSelectedItems();

            let sItem = self.SelectItemByCoordinates(clickCoordsX, clickCoordsY);

            if (sItem) {

                if (selectedItems.length === 1) {
                    self.BlurAll();
                }

                if (sItem.IsSelectable()) {
                    sItem.Select();
                }

            } else {
                self.CreateMultiSelection(clickCoordsX, clickCoordsY);
                self.BlurAll();
            }
        };

        this.Up = (event) => {

            isDown = false;

            self.RemoveMultiSelection();

            self.ClearSelectedItem();
        };

        this.Move = (event) => {

            let newX = self.GetCanvasX(event.pageX) - self.GetSelectedItemOffsetX();
            let newY = self.GetCanvasY(event.pageY) - self.GetSelectedItemOffsetY();

            let sItem = self.GetSelectedItem();
            
            if (self.Mouse.IsDown() && (sItem || IsMiddleButtton(event))) {

                if (sItem) {

                    if (sItem.GetOnMove()) {

                        let m = sItem.GetOnMove();
                        m.Run(newX, newY, sItem, self);

                    } else {

                        sItem.Coords.SetX(newX);
                        sItem.Coords.SetY(newY);
                    }
                }

                let items = [];

                if (IsMiddleButtton(event)) {
                    items = self.GetItems();
                } else {
                    items = self.GetSelectedItems();
                }

                let updateCoords = (item) => {

                    let offsetX = self.GetCanvasX(event.pageX) - clickCoordsX;
                    let offsetY = self.GetCanvasY(event.pageY) - clickCoordsY;

                    item.Coords.SetX(item.Coords.GetPreviousX() + offsetX);
                    item.Coords.SetY(item.Coords.GetPreviousY() + offsetY);
                };

                let moveCoords = (item) => {

                    if (onMoveX === 0 || onMoveY === 0) {
                        return
                    }

                    let offsetX = self.GetCanvasX(event.pageX) - self.GetCanvasX(onMoveX);
                    let offsetY = self.GetCanvasY(event.pageY) - self.GetCanvasY(onMoveY);

                    item.Coords.SetX(item.Coords.GetX() + offsetX);
                    item.Coords.SetY(item.Coords.GetY() + offsetY);
                };

                for (let i in items) {

                    if (sItem && items[i].GetId() === sItem.GetId()) {
                        continue;
                    }

                    if (IsMiddleButtton(event)) {

                        moveCoords(items[i])

                    } else {

                        if (items[i].GetOnMove()) {

                            let m = sItem.GetOnMove();
                            m.Run(newX, newY, self, items[i]);

                        } else {
                            updateCoords(items[i])
                        }
                    }
                }
            }

            onMoveX = event.pageX;
            onMoveY = event.pageY;
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

            if (event.pageX === 0 && event.pageY === 0 && onMoveX !== 0 && onMoveY !==0) {
                event.onMoveCoords = {
                    x: onMoveX,
                    y: onMoveY
                };
            }

            return event;
        }
    };
}

export default Mouse;