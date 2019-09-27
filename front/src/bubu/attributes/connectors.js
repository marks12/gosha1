import ElementsRegister from '../elements-register';
import {TYPES as constants} from "../constants";

function Connectors(config) {

    let connectorPoints = [];

    let isShow = false;

    this.ShowConnectors = () => {
        isShow = true;
    };

    this.HideConnectors = () => {
        isShow = false;
    };

    this.IsShowConnectors = () => {
        return isShow;
    };

    this.AddConnectorPoint = (index, isVisible) => {

        let x = () => {console.error('wrong x getter');};
        let y = () => {console.error('wrong y getter');};

        let item = this;

        switch (index) {
            case 0:
                x = () => {return item.Coords.GetX() + item.GetWidth() / 2};
                y = () => {return item.Coords.GetY()};
                break;
            case 1:
                x = () => {return item.Coords.GetX() + item.GetWidth()};
                y = () => {return item.Coords.GetY() + item.GetHeight() / 2};
                break;
            case 2:
                x = () => {return item.Coords.GetX() + item.GetWidth() / 2};
                y = () => {return item.Coords.GetY() + item.GetHeight()};
                break;
            case 3:
                x = () => {return item.Coords.GetX()};
                y = () => {return item.Coords.GetY() + item.GetHeight() / 2};
                break;
        }

        let cp = new ElementsRegister.ConnectorPoint()
            .SetIndex(index)
            .SetVisibility(isVisible)
            .SetWidth(constants.connectionPointRadius)
            .SetHeight(constants.connectionPointRadius)
            .SetParentId(this.GetId());

        cp.Coords.GetX = x;
        cp.Coords.GetY = y;

        connectorPoints.push(cp);
    };

    this.ClearConnectorPoints = () => {
        connectorPoints = [];
    };

    this.GetConnectorPoints = () => {

        let cps = [];

        for (let i = 0; i < connectorPoints.length; i++) {
            cps.push(connectorPoints[i]);
        }

        return cps;
    };

    this.GetVisibleConnectorPoints = () => {

        let cps = [];

        for (let i = 0; i < connectorPoints.length; i++) {
            if (connectorPoints[i].GetVisibility()) {
                cps.push(connectorPoints[i]);
            }
        }

        return cps;
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

        let items = this.GetSelectableItems();

        for (let index in items) {

            if (! items[index].GetConnectorPoints) {
                continue;
            }

            let connectorPoints = items[index].GetConnectorPoints();

            for (let i = 0; i < connectorPoints.length; i++) {

                let itemX1 = connectorPoints[i].Coords.GetX() - (connectorPoints[i].GetWidth() + constants.activeSpaceAround);
                let itemX2 = connectorPoints[i].Coords.GetX() + (connectorPoints[i].GetWidth() + constants.activeSpaceAround);
                let itemY1 = connectorPoints[i].Coords.GetY() - (connectorPoints[i].GetHeight() + constants.activeSpaceAround);
                let itemY2 = connectorPoints[i].Coords.GetY() + (connectorPoints[i].GetHeight() + constants.activeSpaceAround);

                if (x >= itemX1 && x <= itemX2 && y >= itemY1 && y <= itemY2) {

                    return connectorPoints[i];
                }
            }
        }

        return false;
    };
}

export default Connectors;