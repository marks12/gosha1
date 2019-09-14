function Canvas(canvasElementId) {

    let self = this;

    let canvas = document.getElementById(canvasElementId);

    canvas.setAttribute("width", canvas.parentNode.parentElement.clientWidth);
    canvas.setAttribute("height", canvas.parentNode.parentElement.clientHeight);

    let canvasOffsetX = canvas.getBoundingClientRect().left;
    let canvasOffsetY = canvas.getBoundingClientRect().top;

    this.UpdateCanvas = () => {

        setTimeout(() => {

            canvas.setAttribute("width", document.getElementById(canvasElementId).parentNode.parentElement.clientWidth);
            canvas.setAttribute("height", document.getElementById(canvasElementId).parentNode.parentElement.clientHeight);

            canvasOffsetX = canvas.getBoundingClientRect().left;
            canvasOffsetY = canvas.getBoundingClientRect().top;
        });
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
    canvas.addEventListener("ondrop", importElement);

    canvas.addEventListener("onwheel", wheel);
    canvas.addEventListener("mousewheel", wheel);
    canvas.addEventListener("MozMousePixelScroll", wheel);

    function importElement(event) {
        console.log('import element event', event);
    }

    this.GetCanvas = () => {
        return this.canvas;
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

    this.GetCanvasOffsetX = () => {
        return canvasOffsetX;
    };

    this.GetCanvasOffsetY = () => {
        return canvasOffsetY;
    };

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

        let ctx = self.GetCtx();

        if (delta > 0) {
            ctx.scale(1.1, 1.1);
        } else {
            ctx.scale(0.9, 0.9);
        }

    };

}

export default Canvas;