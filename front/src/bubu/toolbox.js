import ElementsRegister from "./elements-register";
import {TYPES as constants} from "./constants";

function Toolbox(config, toolboxElementId) {

    let self = this;

    let bubuElement = document.getElementById(toolboxElementId);

    if (bubuElement) {

        bubuElement.appendChild((
            new Img(this.GetSrcImageTask())
        ).GetNode());

        bubuElement.appendChild((
            new Img(this.GetSrcImageCondition())
        ).GetNode());

    } else {
        console.error('Root toolbox not found. Id = ' + toolboxElementId +
            '. Please check is countainer with id=\'' + toolboxElementId + '\' ' +
            'exists or create new container like: <div id="BubuToolbox">')
    }

    function Img(src) {

        let img = document.createElement('img');

        img.setAttribute('data-bubu', 'condition');
        img.setAttribute('draggable', 'true');
        img.addEventListener('dragend', self.DropElement);
        bubuElement.appendChild(img);
        img.setAttribute('src', src);

        this.GetNode = () => {
            return img;
        }
    }

}

export default Toolbox;
