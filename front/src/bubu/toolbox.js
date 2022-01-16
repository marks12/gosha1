import {TYPES as constants} from "./constants";

function Toolbox(config, toolboxElementId) {

    let self = this;

    let renderToolBox = () => {

        let parentItem = document.getElementById(toolboxElementId);

        if (parentItem) {

            parentItem.innerHTML = '';

            parentItem.appendChild((
                new Img(this.GetSrcImageCondition(), constants.condition)
            ).GetNode());

            parentItem.appendChild((
                new Img(this.GetSrcImageTask(), constants.task)
            ).GetNode());
        }
    }

    if (document.getElementById(toolboxElementId)) {
        renderToolBox()
    } else {

        console.error('Root toolbox not found. Id = ' + toolboxElementId +
            '. Please check is countainer with id=\'' + toolboxElementId + '\' ' +
            'exists or create new container like: <div id="BubuToolbox"></div>')
    }

    this.UpdateToolbox = () => {
        renderToolBox();
    }

    function Img(src, type) {

        let img = document.createElement('img');

        img.setAttribute('data-bubu', type);
        img.setAttribute('draggable', 'true');

        img.addEventListener('dragstart', (ev)=>{
            return true;
        });

        img.addEventListener('dragenter', (event)=>{
            event.preventDefault();
            return true;
        });

        img.addEventListener('dragover', (event)=>{
            event.preventDefault();
        });

        img.addEventListener('dragend', self.DropElement);
        img.addEventListener('drop', self.DropElement);

        img.setAttribute('src', src);

        this.GetNode = () => {
            return img;
        }
    }

}

export default Toolbox;
