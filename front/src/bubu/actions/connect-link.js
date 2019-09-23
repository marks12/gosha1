function ConnectLink() {

    this.Run = (newX, newY, sItem, root, event) => {

        if (! sItem) {
            return false;
        }

        if (! sItem.SetLinkDestinationCoords) {
            return false;
        }

        if (! root) {
            return false;
        }

        console.log('connect link', newX, newY);
        sItem.SetLinkDestinationCoords(root.GetCanvasX(event.pageX), root.GetCanvasY(event.pageY));
    };
}

export default ConnectLink;