import {TYPES as types} from "../constants";

function LinkDirection(config) {


    let destinationPoint = null;

    if (config && config.destinationPointId && this.GetItemById(config.destinationPointId)) {
        destinationPoint = this.GetItemById(config.destinationPointId);
    }

    let sourcePoint = null;

    if (config && config.sourcePointId && this.GetItemById(config.sourcePointId)) {
        destinationPoint = this.GetItemById(config.sourcePointId);
    }

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

        if (destinationPoint) {
            destinationPoint.ClearAssignLinkDestination();
        }

        destinationPoint = null;
        return this;
    };

    this.ClearLinkSourcePoint = () => {

        if (sourcePoint) {
            sourcePoint.ClearAssignLinkSource();
        }

        sourcePoint = null;
        return this;
    };

    this.SetLinkDestinationPoint = (point) => {

        if (sourcePoint && sourcePoint.GetId() === point.GetId()) {
            return this;
        }

        if (destinationPoint) {
            destinationPoint.ClearAssignLinkDestination(this.GetId());
        }

        destinationPoint = point;
        point.SetAssignLinkDestination(this.GetId());

        return this;
    };

    this.SetLinkSourcePoint = (point) => {

        if (destinationPoint && destinationPoint.GetId() === point.GetId()) {
            return this;
        }

        if (sourcePoint && sourcePoint.GetId && sourcePoint.GetId() !== point.GetId()) {
            sourcePoint.SetVisibility(false);
            sourcePoint.ClearAssignLinkSource();
        }

        point.SetAssignLinkSource(this.GetId());

        sourcePoint = point;
        return this;
    };
}

export default LinkDirection;