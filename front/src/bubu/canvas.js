function Canvas(canvasElementId) {

    let self = this;

    let canvas = document.getElementById(canvasElementId);

    for (let i in canvas.parentNode.parentElement) {
        console.log(i, canvas.parentNode.parentElement[i]);
    }

    canvas.setAttribute("width", canvas.parentNode.parentElement.clientWidth);
    canvas.setAttribute("height", canvas.parentNode.parentElement.clientHeight);

    let canvasOffsetX = canvas.getBoundingClientRect().left;
    let canvasOffsetY = canvas.getBoundingClientRect().top;

    this.SetCtx(canvas.getContext('2d'));

    canvas.addEventListener("mousedown", down);
    canvas.addEventListener("mousemove", mover);
    canvas.addEventListener("mouseup", up);
    canvas.addEventListener("ondrop", importElement);

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


}

export default Canvas;