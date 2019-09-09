function Canvas(canvasElementId) {

    let canvas = document.getElementById(canvasElementId);

    canvas.setAttribute("width", canvas.parentNode.parentElement.clientWidth);
    canvas.setAttribute("height", canvas.parentNode.parentElement.clientHeight);

    let canvasOffsetX = this.canvas.getBoundingClientRect().left;
    let canvasOffsetY = this.canvas.getBoundingClientRect().top;

    this.GetCanvasOffsetX = () => {
        return canvasOffsetX;
    };

    this.GetCanvasOffsetY = () => {
        return canvasOffsetY;
    };
}

export default Canvas;