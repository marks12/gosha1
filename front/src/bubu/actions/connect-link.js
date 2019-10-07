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

        // меняем source на более удобный
        // берем все соурсы и определяем ближайший, переключаемся на него

        setNearSourcePoint(root, sItem, clickCoordsX, clickCoordsY);

        sItem.SetLinkDestinationCoords(clickCoordsX, clickCoordsY);

        let point = root.GetNearConnectorPoint(clickCoordsX, clickCoordsY, sItem.GetLinkSourcePoint());

        if (point && sItem.GetId() !== point.GetParentId() && point.GetId() !== sItem.GetLinkSourcePoint().GetId()) {

            sItem.SetLinkDestinationPoint(point);
        } else {
            sItem.ClearLinkDestinationPoint();
        }

        root.ShowDstConnectors(clickCoordsX, clickCoordsY);
    };

    function getNearPoint(clickCoordsX, clickCoordsY, points) {

        let minDx = 2200000000;
        let nearpoint = null;

        for (let i in points) {

            let dx = points[i].Coords.GetX() - clickCoordsX;
            let dy = points[i].Coords.GetY() - clickCoordsY;

            let vectorLen = Math.sqrt((dx*dx) + (dy*dy));

            if (vectorLen < minDx) {
                minDx = vectorLen;
                nearpoint = points[i];
            }
        }

        return nearpoint;
    }

    function setNearSourcePoint(root, sItem, clickCoordsX, clickCoordsY) {

        let source = sItem.GetLinkSourcePoint();
        let sourceBlock = root.GetItemById(source.GetParentId());
        let points = sourceBlock.GetConnectorPoints();

        let p = getNearPoint(clickCoordsX, clickCoordsY, points);

        if (! p) {
            return false;
        }

        if (p.GetId() !== source.GetId()) {
            sItem.SetLinkSourcePoint(p);
            p.SetVisibility(true);
        }
    }
    //
    // function setNearDestinationPoint(root, sItem, destinationBlock, sourcePointX, sourcePointY) {
    //
    //     let points = destinationBlock.GetConnectorPoints();
    //     let dstPoint = destinationBlock.GetLinkDestinationPoint();
    //
    //     let p = getNearPoint(sourcePointX, sourcePointY, points);
    //
    //     if (! p) {
    //         return false;
    //     }
    //
    //     if (! dstPoint || p.GetId() !== dstPoint.GetId()) {
    //         sItem.SetLinkDestinationPoint(p);
    //         p.SetVisibility(true);
    //     }
    // }
}

export default ConnectLink;