import BuBu from "../bubu/main";

class Documentor {

    bubuStore = {};

    createInstance(canvasId, toolboxId) {
        return new BuBu(canvasId, toolboxId);
    }

    GetInstance(canvasId, toolboxId) {

        let key = this.getStoreKey(canvasId,toolboxId);

        if (! this.bubuStore.hasOwnProperty(key)) {
            this.bubuStore[key] = this.createInstance(canvasId, toolboxId);
        }

        this.bubuStore[key].UpdateToolbox();
        this.bubuStore[key].ResetCanvas(canvasId);
        return this.bubuStore[key];
    }

    getStoreKey(canvasId, toolboxId) {
        return 'bubu_' + canvasId.toString() + '_' + toolboxId.toString();
    }
}

let documentor = new Documentor();

export default documentor;