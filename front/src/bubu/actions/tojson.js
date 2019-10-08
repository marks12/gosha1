function ToJson() {

    let obj = {};

    this.toJSON = () => {

        let btns = this.GetButtons();

        if (btns && btns.length) {
            obj.buttons = btns;
        }

        obj.isShowButtons = this.IsShowButtons();

        obj.color = this.GetColor();
        if (! obj.color) {
            delete obj.color;
        }

        obj.name = this.GetName();
        if (! obj.name) {
            delete obj.name;
        }

        obj.description = this.GetDescription();
        if (! obj.description) {
            delete obj.description;
        }

        obj.connectorPoints = this.GetConnectorPoints();
        if (! obj.connectorPoints.length) {
            delete obj.connectorPoints;
        }

        obj.Coords = {
            X: this.Coords.GetX(),
            Y: this.Coords.GetY(),
        };

        obj.id = this.GetId();

        if (this.GetLinkArrowType) {
            obj.LinkArrowType = this.GetLinkArrowType();
        }

        if (this.GetLinkDestinationPoint) {
            let p = this.GetLinkDestinationPoint();
            obj.destinationPointId = p ? p : null;

            if (! obj.destinationPointId) {
                delete obj.destinationPointId;
            }
        }

        if (this.GetLinkSourcePoint) {
            let p = this.GetLinkSourcePoint();
            obj.sourcePointId = p ? p : null;

            if (! obj.sourcePointId) {
                delete obj.sourcePointId;
            }
        }

        if (this.IsSource) {
            obj.linkLineIsSource = this.IsSource();
        }

        if (this.GetAssignedLinkDestination) {
            obj.linkDestinationId = this.GetAssignedLinkDestination();
            if (! obj.linkDestinationId) {
                delete obj.linkDestinationId;
            }
        }

        if (this.GetAssignedLinkSource) {
            obj.linkSourceId = this.GetAssignedLinkSource();
            if (! obj.linkSourceId) {
                delete obj.linkSourceId;
            }
        }

        if (this.GetParentId) {
            obj.Parent = this.GetParentId();
            if (! obj.Parent) {
                delete obj.Parent;
            }
        }

        obj.Width = this.GetWidth();
        obj.Height = this.GetHeight();

        obj.Text = this.GetText();
        if (! obj.Text) {
            delete obj.Text;
        }

        obj.Type = this.GetType();
        obj.Visibility = this.GetVisibility();

        return obj;
    };
}

export default ToJson;