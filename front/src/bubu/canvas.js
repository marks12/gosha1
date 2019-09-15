import ElementsRegister from "./elements-register";

function Canvas(canvasElementId) {

    let self = this;
    let canvas = document.getElementById(canvasElementId);
    let scale = 1;

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
        });
    };

    this.GetCanvasX = (x) => {
        return (x - this.Zero.Coords.GetX()) * (1 / scale);
    };

    this.GetCanvasY = (y) => {
        return (y - this.Zero.Coords.GetY()) * (1 / scale);
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

        if (delta > 0) {
            if (scale < 8) {
                scale = scale * 1.2;
            }
        } else {
            if (scale > 0.35) {
                scale = scale / 1.2;
            }
        }

        let ctx = self.GetCtx();
        ctx.setTransform(scale, 0, 0, scale, 0, 0);

        this.Render();
    };

}

export default Canvas;