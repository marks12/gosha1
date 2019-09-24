import ElementsRegister from '../elements-register';

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

    this.AddConnectorPoint = (x1, y1, x2, y2, parentId) => {

        connectorPoints.push((
            new ElementsRegister.ConnectorPoint()
                .Coords.SetX(x1)
                .Coords.SetY(y1)
                .SetWidth(x2 - x1)
                .SetHeight(y2 - y1)
                .SetParentId(parentId)
                .SetIndex(connectorPoints.length)
        ));
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