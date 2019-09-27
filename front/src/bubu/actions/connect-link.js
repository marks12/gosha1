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

        let point = root.GetConnectorPoint(clickCoordsX, clickCoordsY);

        if (point && sItem.GetId() !== point.GetParentId()) {
            sItem.SetLinkDestinationPoint(point);
        } else {
            sItem.ClearLinkDestinationPoint();
        }

        root.ShowDstConnectors(clickCoordsX, clickCoordsY);
    };

    function getNearPoint(clickCoordsX, clickCoordsY, points) {
        return undefined;
    }

    function setNearSourcePoint(root, sItem, clickCoordsX, clickCoordsY) {

        let source = sItem.GetLinkSourcePoint();
        let sourceBlock = root.GetItemById(source.GetParentId())
        let points = sourceBlock.GetConnectorPoints();

        console.log('points', points);
        console.log('----------');

        let p = getNearPoint(clickCoordsX, clickCoordsY, points);

        if (p.GetId() !== source.GetId()) {
            sItem.SetLinkSourcePoint(p);
            source.SetVisibility(false);
        }

        console.log('----------');
    }
}

export default ConnectLink;