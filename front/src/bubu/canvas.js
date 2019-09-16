import ElementsRegister from "./elements-register";

function Canvas(canvasElementId) {

    let self = this;
    let canvas = document.getElementById(canvasElementId);
    let scale = 1;
    let zoomCount = 0;

    this.Zero = new ElementsRegister.ZeroPoint();

    canvas.setAttribute("width", canvas.parentNode.parentElement.clientWidth);
    canvas.setAttribute("height", canvas.parentNode.parentElement.clientHeight);

    this.Zero.Coords.SetX(canvas.getBoundingClientRect().left);
    this.Zero.Coords.SetY(canvas.getBoundingClientRect().top);

    this.UpdateCanvas = () => {

        setTimeout(() => {

            canvas.setAttribute("width", document.getElementById(canvasElementId).parentNode.parentElement.clientWidth);
            canvas.setAttribute("height", document.getElementById(canvasElementId).parentNode.parentElement.clientHeight);

            this.Zero.Coords.SetX(canvas.getBoundingClientRect().left);
            this.Zero.Coords.SetY(canvas.getBoundingClientRect().top);

            this.CanvasScale({Delta: 0});
        });
    };

    this.GetCanvasX = (x) => {
        return (x - this.Zero.Coords.GetX()) * (1 / scale);
    };

    this.GetCanvasY = (y) => {
        return (y - this.Zero.Coords.GetY()) * (1 / scale);
    };

    this.GetZero = () => {
        return this.Zero;
    };

    this.UpdateCanvas();

    window.addEventListener("resize", this.UpdateCanvas);

    this.SetCtx(canvas.getContext('2d'));

    canvas.addEventListener("mousedown", down);
    canvas.addEventListener("mousemove", mover);
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

        let delta = event.deltaY || event.detail || event.wheelDelta;

        let zoommer = 1.06;

        let stopZoom = true;

        if (delta > 0) {
            if (scale < 8) {
                scale = scale * zoommer;
                zoomCount++;
                stopZoom = false;
            }
        } else {
            if (scale > 0.35) {
                scale = scale / zoommer;
                zoomCount--;
                stopZoom = false;
            }
        }

        if (stopZoom) {
            return true;
        }

        let items = this.GetSelectableItems();

        if (! items.length) {
            return true;
        }

        let distanceX = [];
        let distanceY = [];

        for (let i in items) {
            distanceX[i] = this.GetCanvasX(event.pageX) - items[i].Coords.GetX();
            distanceY[i] = this.GetCanvasY(event.pageY) - items[i].Coords.GetY();
        }

        let ctx = self.GetCtx();
        ctx.setTransform(scale, 0, 0, scale, 0, 0);

        for (let i in items) {
            items[i].Coords.SetX( this.GetCanvasX(event.pageX) + distanceX[i]);
            items[i].Coords.SetY( this.GetCanvasY(event.pageY) + distanceY[i]);

            console.log(distanceX[i]);
            console.log(distanceY[i]);

        }

        this.Render();

    };
}

export default Canvas;