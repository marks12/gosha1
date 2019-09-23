function ConnectLink() {

    this.Run = (newX, newY, sItem, root, event) => {

        if (!sItem) {
            return false;
        }

        if (!sItem.SetLinkDestinationCoords) {
            return false;
        }

        if (!root) {
            return false;
        }


        let clickCoordsX = root.GetCanvasX(event.pageX);
        let clickCoordsY = root.GetCanvasY(event.pageY);

        sItem.SetLinkDestinationCoords(clickCoordsX, clickCoordsY);

        let point = root.GetConnectorPoint(clickCoordsX, clickCoordsY);

        if (point && sItem.GetId() !== point.elementId) {
            sItem.SetLinkDestinationPoint(point);
        } else {
            sItem.ClearLinkDestinationPoint();
        }

        root.ShowDstConnectors(clickCoordsX, clickCoordsY);
    };
}

export default ConnectLink;