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

    this.AddConnectorPoint = (side, isVisible) => {

        let x = () => {console.error('wrong x getter');};
        let y = () => {console.error('wrong y getter');};

        let item = this;

        switch (side) {
            case 'top':
                x = () => {return item.Coords.GetX() + item.GetWidth() / 2};
                y = () => {return item.Coords.GetY()};
                break;
            case 'right':
                x = () => {return item.Coords.GetX() + item.GetWidth()};
                y = () => {return item.Coords.GetY() + item.GetHeight() / 2};
                break;
            case 'bottom':
                x = () => {return item.Coords.GetX() + item.GetWidth() / 2};
                y = () => {return item.Coords.GetY() + item.GetHeight()};
                break;
            case 'left':
                x = () => {return item.Coords.GetX()};
                y = () => {return item.Coords.GetY() + item.GetHeight() / 2};
                break;
        }

        let cp = new ElementsRegister.ConnectorPoint()
            .SetIndex(side)
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

    this.GetNearConnectorPoint = (x, y, sourceConnectorPoint, dstItem) => {

        let item;

        if (dstItem) {
            item = dstItem;
        } else {
            item = this.FindFirstElementByCoordinates(x, y);
        }

        if (! item || ! item.GetConnectorPoints) {
            return false;
        }

        let connectorPoints = item.GetConnectorPoints();

        let minDx = 2200000000;
        let nearpoint = null;

        for (let i = 0; i < connectorPoints.length; i++) {

            let dx = connectorPoints[i].Coords.GetX() - sourceConnectorPoint.Coords.GetX();
            let dy = connectorPoints[i].Coords.GetY() - sourceConnectorPoint.Coords.GetY();

            let vectorLen = Math.sqrt((dx*dx) + (dy*dy));

            if (vectorLen < minDx) {
                minDx = vectorLen;
                nearpoint = connectorPoints[i];
            }
        }

        return nearpoint;
    };

    this.ReConnectPoints = (root, isNotCallLinked) => {

        let cp = this.GetConnectorPoints();

        let x = this.Coords.GetX();
        let y = this.Coords.GetY();

        for (let i in cp) {

            let sp = cp[i].GetAssignedLinkSource();
            let dp = cp[i].GetAssignedLinkDestination();

            let resign = (id) => {

                let link = root.GetItemById(id);

                if (! link) {
                    return;
                }

                let dpForSp = link.GetLinkDestinationPoint();
                let spCp = link.GetLinkSourcePoint();
                let dstBlock = root.GetItemById(dpForSp.GetParentId());
                let srcBlock = root.GetItemById(spCp.GetParentId());


                let newDstPoint = root.GetNearConnectorPoint(x, y, spCp, dstBlock);
                link.SetLinkDestinationPoint(newDstPoint);

                let newSrcPoint = root.GetNearConnectorPoint(x, y, newDstPoint, srcBlock);
                link.SetLinkSourcePoint(newSrcPoint);

                if (! isNotCallLinked) {
                    dstBlock.ReConnectPoints(root, true);
                    srcBlock.ReConnectPoints(root, true);
                }
            };

            if (sp) {
                resign(sp);
            }

            if (dp) {
                resign(dp);
            }

        }

    };
}

export default Connectors;