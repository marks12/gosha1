function Connectors(config) {

    let selectors = [];

    let isShow = false;

    this.ShowConnectors = () => {
        isShow = true;
    };

    this.HideConnectors = () => {
        selectors = [];
        isShow = false;
    };

    this.IsShowConnectors = () => {
        return isShow;
    };

    this.AddSelector = (x1, y1, x2, y2) => {

        selectors.push({
            x1: x1,
            x2: x2,
            y1: y1,
            y2: y2,
        });

        return this;
    }
}

export default Connectors;