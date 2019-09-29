function LinkPoint(config) {

    let linkDestinationId = null;
    let linkSourceId = null;

    if (config && config.LinkSourceId) {
        linkSourceId = config.LinkSourceId;
    }

    if (config && config.LinkDestinationId) {
        linkDestinationId = config.LinkDestinationId;
    }

    this.SetAssignLinkDestination = (linkId) => {
        linkDestinationId = linkId;
        return this;
    };

    this.ClearAssignLinkDestination = () => {
        linkDestinationId = null;
    };

    this.GetAssignedLinkDestination = () => {
        return linkDestinationId;
    };

    this.SetAssignLinkSource = (linkId) => {
        linkSourceId = linkId;
        return this;
    };

    this.GetAssignedLinkSource = () => {
        return linkSourceId;
    };

    this.ClearAssignLinkSource = () => {
        linkSourceId = null;
    };
}

export default LinkPoint;