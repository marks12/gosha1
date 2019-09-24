function LinkDirection(config) {

    let destinationPoint = null;
    let sourcePoint = null;
    let destX = 0;
    let destY = 0;

    this.GetLinkDestinationPoint = () => {
        return destinationPoint;
    };

    this.GetLinkSourcePoint = () => {
        return sourcePoint;
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
        sourcePoint = null;
        return this;
    };

    this.SetLinkDestinationPoint = (point) => {
        destinationPoint = point;
        return this;
    };

    this.SetLinkSourcePoint = (point) => {
        sourcePoint = point;
        return this;
    };
}

export default LinkDirection;