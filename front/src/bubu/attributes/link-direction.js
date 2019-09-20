function LinkDirection(config) {

    let isSource = true;
    let isDestination = !isSource;

    this.MarkAsSource = () => {
        isSource = true;
        isDestination = !isSource;

        return this;
    };

    this.MarkAsDestination = () => {
        isSource = false;
        isDestination = !isSource;

        return this;
    };

    this.IsSource = () => {
        return isSource;
    };

    this.IsDestination = () => {
        return isDestination;
    };
}

export default LinkDirection;