import ElementsRegister from "./elements-register";
import {TYPES as Const} from "./constants";

function Canvas(canvasElementId) {

    let self = this;
    let canvas = document.getElementById(canvasElementId);
    let scale = 1;
    let zoomCount = 0;
    let mode = 1;

    this.Zero = new ElementsRegister.ZeroPoint();

    canvas.setAttribute("width", canvas.parentNode.parentElement.clientWidth);
    canvas.setAttribute("height", canvas.parentNode.parentElement.clientHeight);

    this.Zero.Coords.SetX(canvas.getBoundingClientRect().left);
    this.Zero.Coords.SetY(canvas.getBoundingClientRect().top);

    this.UpdateCanvas = () => {

        canvas.setAttribute("width", document.getElementById(canvasElementId).parentNode.parentElement.clientWidth);
        canvas.setAttribute("height", document.getElementById(canvasElementId).parentNode.parentElement.clientHeight);

        this.Zero.Coords.SetX(canvas.getBoundingClientRect().left);
        this.Zero.Coords.SetY(canvas.getBoundingClientRect().top);

        this.Render();
    };

    setTimeout(() => {
        this.UpdateCanvas();
    });


    this.GetCanvasX = (x) => {
        let a = (x - this.Zero.Coords.GetX());
        return a * (1 / scale);
    };

    this.GetCanvasY = (y) => {
        let a = (y - this.Zero.Coords.GetY());
        return  a * (1 / scale);
    };

    this.GetZero = () => {
        return this.Zero;
    };

    window.addEventListener("resize", this.UpdateCanvas);

    this.SetCtx(canvas.getContext('2d'));

    canvas.addEventListener("mousedown", down);
    canvas.addEventListener("mousemove", mover);
    canvas.addEventListener("dragover", mover);
    canvas.addEventListener("mouseup", up);
    canvas.addEventListener("mouseout", (event) => {
        this.Mouse.Up(event);
        this.Render();
    });

    canvas.addEventListener("onwheel", wheel);
    canvas.addEventListener("mousewheel", wheel);
    canvas.addEventListener("MozMousePixelScroll", wheel);

    this.GetCanvas = () => {
        return this.canvas;
    };

    this.GetScale = () => {
        return scale;
    };

    function down(event) {

        self.Mouse.Down(event);
        self.Render();
    }

    function up(event) {

        self.Mouse.Up(event);
        self.Render();

    }

    function wheel(event) {
        self.CanvasScale(event);
    }

    function mover (event) {

        self.Mouse.Move(event);
        self.Render();

    }

    this.GetCanvas = () => {
        return canvas;
    };

    this.GetContext = () => {
        return canvas.getContext();
    };

    this.GetCanvasHeight = () => {
        return canvas.getAttribute("height") * 1;
    };

    this.GetCanvasWidth = () => {
        return canvas.getAttribute("width") * 1;
    };

    this.CanvasScale = (event) => {

        mode *= -1;

        let maxScale = 8;
        let minScale = 0.35;
        let delta = event.deltaY || event.detail || event.wheelDelta;

        let zoommer = 1.07;

        let previosScale = (()=>{return scale})();

        if (delta >= 0) {

            if (scale < maxScale) {
                scale = scale * zoommer;
                zoomCount++;
            }

        } else {
            if (scale >= minScale) {
                scale = scale / zoommer;
                zoomCount--;
            }
        }

        let items = this.GetItems();

        if (! items.length) {
            return true;
        }

        let distanceX = [];
        let distanceY = [];

        this.UpdateCanvas();

        for (let i in items) {
            distanceX[i] = this.GetCanvasX(event.pageX) - items[i].Coords.GetX();
            distanceY[i] = this.GetCanvasY(event.pageY) - items[i].Coords.GetY();
        }

        self.GetCtx().scale(scale, scale);

        for (let i in items) {

            let newX = this.GetCanvasX(event.pageX) + distanceX[i];
            let newY = this.GetCanvasY(event.pageY) + distanceY[i];

            if (mode < 0 && scale < maxScale && scale >= minScale) {
                newX += 3;
                newY += 3;
            }

            items[i].Coords.SetX(newX);
            items[i].Coords.SetY(newY);
        }

        if (mode >= 0) {
                this.Render();
        } else {
            setTimeout(()=>{
                this.CanvasScale(event);
            });
        }
    };

    this.ShowElementConnectors = (x, y) => {

        let item = this.GetFirstElementByCoordinates(
            x,
            y,
            Const.connectionPointRadius + Const.spaceBetween,
            Const.connectionPointRadius + Const.spaceBetween
        );

        if (item) {
            console.log('element',item.Coords.GetX());
            item.ShowConnectors();
        } else {
            self.ClearConnectors();
            console.log('no element');
        }
    };
}

export default Canvas;