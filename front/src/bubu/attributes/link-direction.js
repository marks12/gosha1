function LinkDirection(config) {

    let destinationPoint = null;
    let srcSourcePoint = null;
    let destX = 0;
    let destY = 0;

    this.GetLinkDestinationPoint = () => {
        return destinationPoint;
    };

    this.GetLinkSourcePoint = () => {
        return srcSourcePoint;
    };

    this.SetLinkDestinationCoords = (x, y) => {
        destX = x;
        destY = y;
        return this;
    };

    this.GetLinkDestinationX = () => {
        return destX;
    };

    this.GetLinkDestinationY = () => {
        return destY;
    };

    this.ClearLinkDestinationPoint = () => {
        destinationPoint = null;
        return this;
    };

    this.ClearLinkSourcePoint = () => {
        srcSourcePoint = null;
        return this;
    };

    this.SetLinkDestinationPoint = (point) => {
        destinationPoint = point;
        return this;
    };

    this.SetLinkSourcePoint = (point) => {
        srcSourcePoint = point;
        return this;
    };
}

export default LinkDirection;