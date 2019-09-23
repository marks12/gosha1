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

    this.ClearLinkDestination = () => {
        destination = null;
    };
}

export default LinkDirection;