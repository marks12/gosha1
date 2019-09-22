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

    this.AddConnectorPoint = (x1, y1, x2, y2) => {
        connectorPoints.push({
            x1: x1,
            y1: y1,
            x2: x2,
            y2: y2
        })
    };

    this.ClearConnectorPoints = () => {
        connectorPoints = [];
    };

    this.IsConnectorPointCoords = (x, y) => {

        for (let i = 0; i < connectorPoints.length; i++) {

            if (x >= connectorPoints[i].x1 && x <= connectorPoints[i].x2 && y >= connectorPoints[i].y1 && y <= connectorPoints[i].y2) {
                return true;
            }
        }

        return false;
    };

}

export default Connectors;