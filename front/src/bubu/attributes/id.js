import {generateUUID} from "../common";

function Id(config) {

    let Id = config && config.Id ? config.Id : generateUUID();

    this.GetId = () => {return Id};
}

export default Id;