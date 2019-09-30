import {TYPES} from "../constants";

function Draw(config) {

    let self = this;

    this.draw = (ctx, root) => {


        let W = this.GetWidth();
        let H = this.GetHeight();
        let lineWidth = TYPES.lineWidth;
        let X = this.Coords.GetX();
        let Y = this.Coords.GetY();

        ctx.lineWidth = lineWidth;

        switch (this.GetType()) {

            case TYPES.task:


                drawRoundRect(ctx, X, Y, W, H, TYPES.cornerRadius, false, true);

                if (this.IsSelected()) {
                    ctx.strokeStyle = 'red';
                } else {
                    ctx.strokeStyle = this.GetColorDefault();
                }
                ctx.stroke();

                ctx.fillStyle = this.GetFillColorDefault();
                ctx.fill();

                addText(ctx);

                break;

            case TYPES.condition:

                let curX = X + W / 2;
                let curY = Y;


                ctx.beginPath();
                ctx.moveTo(curX, curY);
                ctx.lineTo(curX + W / 2, curY + H / 2);
                ctx.lineTo(curX, curY + H + (lineWidth / 2) - 1);
                ctx.lineTo(curX - W / 2, curY + H / 2);
                ctx.lineTo(curX + lineWidth / 2 + lineWidth / 5, curY);

                if (this.IsSelected()) {
                    ctx.strokeStyle = 'red';
                } else {
                    ctx.strokeStyle = this.GetColorDefault();
                }

                ctx.lineWidth = lineWidth * 1.3;
                ctx.stroke();
                ctx.lineWidth = lineWidth;

                ctx.fillStyle = this.GetFillColorDefault();
                ctx.fill();

                addText(ctx);

                break;

            case TYPES.divider:

                ctx.beginPath();
                ctx.moveTo(X, Y);
                ctx.lineTo(X + W, Y);
                ctx.lineTo(X + W, Y + H);
                ctx.lineTo(X, Y + H);
                ctx.lineTo(X, Y);
                ctx.fillStyle = this.GetColor();
                ctx.fill();

                break;

            case TYPES.multiselection:

                ctx.beginPath();
                ctx.moveTo(X, Y);
                ctx.lineTo(X + W, Y);
                ctx.lineTo(X + W, Y + H);
                ctx.lineTo(X, Y + H);
                ctx.lineTo(X, Y);

                ctx.strokeStyle = '#d80010';

                ctx.stroke();

            break;

            case TYPES.background:

                ctx.beginPath();
                ctx.moveTo(X, Y);
                ctx.lineTo(X + W, Y);
                ctx.lineTo(X + W, Y + H);
                ctx.lineTo(X, Y + H);
                ctx.lineTo(X, Y);

                if (this.IsSelected()) {
                    ctx.strokeStyle = 'red';
                } else {
                    ctx.strokeStyle = this.GetColorDefault();
                }

                ctx.fillStyle = this.GetColor();

                ctx.stroke();
                ctx.fill();

                break;

            case TYPES.link:

                let x2 = this.GetLinkDestinationX();
                let y2 = this.GetLinkDestinationY();
                let dest = this.GetLinkDestinationPoint();
                let src = this.GetLinkSourcePoint();

                if (dest) {
                    x2 = dest.Coords.GetX();
                    y2 = dest.Coords.GetY();
                }

                if (src && x2 && y2) {
                    // ctx.beginPath();
                    // ctx.moveTo(src.Coords.GetX(), src.Coords.GetY());
                    // ctx.lineTo(x2, y2);
                    // ctx.stroke();

                    drawArrow(ctx, src.Coords.GetX(), src.Coords.GetY(), x2, y2, TYPES.arrowTypeSimple);
                }

                break;
        }

        if (this.IsShowConnectors()) {
            drawConnectors(ctx, root);
        }


        if (this.IsShowButtons()) {
            drawButtons(ctx, root);
        }

    };

    let drawButtons = (ctx, root) => {

        let drawButton = (button) => {

            ctx.save();

            let x = button.Coords.GetX();
            let y = button.Coords.GetY();

            let radius = button.GetWidth() / 2;

            ctx.beginPath();
            ctx.arc(x, y, radius, 0, 2 * Math.PI, false);
            ctx.strokeStyle = '#0060A6';
            ctx.lineWidth = radius * 0.2;
            ctx.fill();

            drawArrow(ctx, x - radius + radius*0.2 , y, x + radius - radius*0.3, y, TYPES.arrowTypeSimple, TYPES.arrowTypeNone, radius);

            ctx.stroke();

            ctx.restore();
        };

        let buttons = this.GetButtons();

        if (buttons.length) {
            for (let i=0; i< buttons.length; i++) {
                drawButton(buttons[i]);
            }
        }

    };

    let addText = (ctx) => {

        if (this.GetText().length) {
            ctx.fillStyle = "#000000";
            ctx.font = "10px Arial";
            ctx.fillText(this.GetText(), this.Coords.GetX(), this.Coords.GetY() + this.GetHeight() / 2);
        }
    };

    let drawConnectors = (ctx, root) => {

        let drawConnector = (x, y) => {
            ctx.beginPath();
            ctx.arc(x, y, TYPES.connectionPointRadius, 0, 2 * Math.PI, false);
            ctx.strokeStyle = '#0060A6';
            ctx.fill();
            ctx.stroke();
        };

        switch (this.GetType()) {

            case TYPES.condition:
            case TYPES.task:

                ctx.fillStyle = TYPES.connectionFillStyle;
                ctx.strokeStyle = TYPES.connectionStrokeStyle;

                let points = this.GetVisibleConnectorPoints();

                for (let i=0;i<points.length;i++) {
                    if (points[i].GetVisibility()) {
                        drawConnector(points[i].Coords.GetX(), points[i].Coords.GetY());
                    }
                }
            break;
        }
    };

    /**
     * Draws a rounded rectangle using the current state of the canvas.
     * If you omit the last three params, it will draw a rectangle
     * outline with a 5 pixel border radius
     * @param {CanvasRenderingContext2D} ctx
     * @param {Number} x The top left x coordinate
     * @param {Number} y The top left y coordinate
     * @param {Number} width The width of the rectangle
     * @param {Number} height The height of the rectangle
     * @param {Number} [radius = 5] The corner radius; It can also be an object
     *                 to specify different radii for corners
     * @param {Number} [radius.tl = 0] Top left
     * @param {Number} [radius.tr = 0] Top right
     * @param {Number} [radius.br = 0] Bottom right
     * @param {Number} [radius.bl = 0] Bottom left
     * @param {Boolean} [fill = false] Whether to fill the rectangle.
     * @param {Boolean} [stroke = true] Whether to stroke the rectangle.
     */
    function drawRoundRect(ctx, x, y, width, height, radius, fill, stroke) {
        if (typeof stroke === 'undefined') {
            stroke = true;
        }
        if (typeof radius === 'undefined') {
            radius = 5;
        }
        if (typeof radius === 'number') {
            radius = {tl: radius, tr: radius, br: radius, bl: radius};
        } else {
            let defaultRadius = {tl: 0, tr: 0, br: 0, bl: 0};
            for (let side in defaultRadius) {
                radius[side] = radius[side] || defaultRadius[side];
            }
        }
        ctx.beginPath();
        ctx.moveTo(x + radius.tl, y);
        ctx.lineTo(x + width - radius.tr, y);
        ctx.quadraticCurveTo(x + width, y, x + width, y + radius.tr);
        ctx.lineTo(x + width, y + height - radius.br);
        ctx.quadraticCurveTo(x + width, y + height, x + width - radius.br, y + height);
        ctx.lineTo(x + radius.bl, y + height);
        ctx.quadraticCurveTo(x, y + height, x, y + height - radius.bl);
        ctx.lineTo(x, y + radius.tl);
        ctx.quadraticCurveTo(x, y, x + radius.tl, y);
        ctx.closePath();
        if (fill) {
            ctx.fill();
        }
        if (stroke) {
            ctx.stroke();
        }
    }

    function drawArrow(ctx, fromx, fromy, tox, toy, arrowTypeDestination, arrowTypeSource, headLen) {

        //variables to be used when creating the arrow
        if (! headLen)  {
            headLen = 20;
        }

        let angle = Math.atan2(toy-fromy,tox-fromx);

        //starting path of the arrow from the start square to the end square and drawing the stroke
        ctx.beginPath();
        ctx.moveTo(fromx, fromy);
        ctx.lineTo(tox, toy);
        ctx.stroke();

        //starting a new path from the head of the arrow to one of the sides of the point
        ctx.beginPath();
        ctx.moveTo(tox, toy);

        if (! arrowTypeSource) {
            arrowTypeSource = TYPES.arrowTypeNone;
        }

        if (! arrowTypeDestination) {
            arrowTypeDestination = TYPES.arrowTypeNone;
        }

        switch (arrowTypeDestination) {

            case TYPES.arrowTypeSimple:

                // ctx.lineTo(tox - headLen * Math.cos(angle - Math.PI/7),toy - headLen * Math.sin(angle - Math.PI/7));

                //path from the side point of the arrow, to the other side point
                ctx.lineTo(tox-headLen*Math.cos(angle+Math.PI/7),toy-headLen*Math.sin(angle+Math.PI/7));

                //path from the side point back to the tip of the arrow, and then again to the opposite side point
                ctx.lineTo(tox, toy);
                ctx.lineTo(tox-headLen*Math.cos(angle-Math.PI/7),toy-headLen*Math.sin(angle-Math.PI/7));

                break;

            // none
            default:
                break;
        }


        //draws the paths created above
        ctx.stroke();
    }
}


export default Draw;