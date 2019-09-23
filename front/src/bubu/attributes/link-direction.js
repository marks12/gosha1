function LinkDirection(config) {

    let destination = null;
    let destX = 0;
    let destY = 0;

    this.GetLinkDestination = () => {
        return destination;
    };

    this.SetLinkDestinationCoords = (x, y) => {
        destX = x;
        destY = y;
    };

    this.GetLinkDestinationX = () => {
        return destX;
    };

    this.GetLinkDestinationY = () => {
        return destY;
    };

    this.ClearLinkDestination = () => {
        destination = null;
    };
}

export default LinkDirection;