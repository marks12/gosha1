function ToJson() {

    let obj = {};

    this.toJSON = () => {

        // let btns = this.GetButtons();
        //
        // if (btns && btns.length) {
        //     obj.Buttons = btns;
        // }

        obj.IsShowButtons = this.IsShowButtons();

        obj.Color = this.GetColor();
        if (! obj.Color) {
            delete obj.Color;
        }

        obj.Name = this.GetName();
        if (! obj.Name) {
            delete obj.Name;
        }

        obj.Description = this.GetDescription();
        if (! obj.Description) {
            delete obj.Description;
        }

        obj.ConnectorPoints = this.GetConnectorPoints();
        if (! obj.ConnectorPoints.length) {
            delete obj.ConnectorPoints;
        }

        obj.Coords = {
            X: this.Coords.GetX(),
            Y: this.Coords.GetY(),
        };

        obj.Id = this.GetId();

        if (this.GetLinkArrowType) {
            obj.LinkArrowType = this.GetLinkArrowType();
        }

        if (this.GetLinkDestinationPoint) {
            let p = this.GetLinkDestinationPoint();
            obj.DestinationPointId = p ? p.GetId() : null;

            if (! obj.DestinationPointId) {
                delete obj.DestinationPointId;
            }
        }

        if (this.GetLinkSourcePoint) {
            let p = this.GetLinkSourcePoint();
            obj.SourcePointId = p ? p.GetId() : null;

            if (! obj.SourcePointId) {
                delete obj.SourcePointId;
            }
        }

        if (this.IsSource) {
            obj.LinkLineIsSource = this.IsSource();
        }

        if (this.GetAssignedLinkDestination) {
            obj.LinkDestinationId = this.GetAssignedLinkDestination();
            if (! obj.LinkDestinationId) {
                delete obj.LinkDestinationId;
            }
        }

        if (this.GetAssignedLinkSource) {
            obj.LinkSourceId = this.GetAssignedLinkSource();
            if (! obj.LinkSourceId) {
                delete obj.LinkSourceId;
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