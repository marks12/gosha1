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
            Width: 50,
            Height: 50 / 1.6,
        }));
    };

    this.GetSrcImageCondition = function () {

        return getSource(new ElementsRegister.Condition({
            Name: "Condition",
            Description: "Move this condition to work area for create new Condition",
            Coords: {
                X: 25,
                Y: 1,
            },
            Width: 50,
            Height: 50,
        }));
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