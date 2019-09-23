function Connectors(config) {

    let connectorPoints = [];

    let isShow = false;

    this.ShowConnectors = () => {
        isShow = true;
    };

    this.HideConnectors = () => {
        connectorPoints = [];
        isShow = false;
    };

    this.IsShowConnectors = () => {
        return isShow;
    };

    this.AddConnectorPoint = (x1, y1, x2, y2, elementId) => {
        connectorPoints.push({
            x1: x1,
            y1: y1,
            x2: x2,
            y2: y2,
            xCenter: (x1 + x2) / 2,
            yCenter: (y1 + y2) / 2,
            elementId: elementId,
            index: connectorPoints.length,
            move: (x, y) => {
                self.x1 += x;
                this.x2 += x;
                this.y1 += y;
                this.y2 += y;
                this.xCenter += x;
                this.yCenter += y;
            },
        })
    };

    this.ClearConnectorPoints = () => {
        connectorPoints = [];
    };

    this.GetConnectorPointByIndex = (index) => {

        if (connectorPoints && connectorPoints[index]) {
            return connectorPoints[index];
        }
        return null;
    };

    this.IsConnectorPointCoords = (x, y) => {
        for (let i = 0; i < connectorPoints.length; i++) {
            if (x >= connectorPoints[i].x1 && x <= connectorPoints[i].x2 && y >= connectorPoints[i].y1 && y <= connectorPoints[i].y2) {
                return true;
            }
        }
        return false;
    };

    this.GetConnectorPoint = (x, y) => {
        for (let i = 0; i < connectorPoints.length; i++) {
            if (x >= connectorPoints[i].x1 && x <= connectorPoints[i].x2 && y >= connectorPoints[i].y1 && y <= connectorPoints[i].y2) {
                return connectorPoints[i];
            }
        }
        return false;
    };
}

export default Connectors;