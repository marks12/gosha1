function ConnectLink() {

    this.Run = (newX, newY, sItem, root) => {

        if (! sItem) {
            return false;
        }

        console.log('connect link', newX, newY);
        sItem.SetLinkDestinationCoords(newX, newY);
    };
}

export default ConnectLink;