import ElementsRegister from "./elements-register";

function SrcImage () {

    let imgCanva = document.createElement("canvas");
    let context = imgCanva.getContext('2d');

    this.GetSrcImageTask = function () {

        return getSource(new ElementsRegister.Task({
            Name: "Task",
            Description: "Move this task to work area for create new Task",
            Coords: {
                X: 1,
                Y: 1,
            },
            Width: 100,
            Height: 100 / 1.6,
        }))
    };

    let clear = () => {
        let w = imgCanva.getAttribute("width") * 1;
        let h = imgCanva.getAttribute("height") * 1;
        context.clearRect(0, 0, w, h);
    };

    let getSource = (img) => {

        imgCanva.setAttribute("width", (img.GetWidth()) + 2 + "px");
        imgCanva.setAttribute("height", (img.GetHeight() + 2) + "px");
        clear();

        img.draw(context);
        return imgCanva.toDataURL();
    }


}

export default SrcImage;