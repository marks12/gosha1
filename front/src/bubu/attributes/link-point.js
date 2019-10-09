import {TYPES as types} from "../constants";

function LinkPoint(config) {

    let linkDestinationId = config && config.LinkDestinationId ? config.LinkDestinationId : null;
    let linkSourceId = config && config.LinkSourceId ? config.LinkSourceId : null;

    if (config && config.LinkSourceId) {
        linkSourceId = config.LinkSourceId;
    }

    if (config && config.LinkDestinationId) {
        linkDestinationId = config.LinkDestinationId;
    }

    this.SetAssignLinkDestination = (linkId) => {

        if (linkDestinationId !== null && linkDestinationId !== linkId) {
            console.log('need create dest point');
        }

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

        if (linkSourceId !== null && linkSourceId !== linkId) {
            console.log('need create source point');
        }

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