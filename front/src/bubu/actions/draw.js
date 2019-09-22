import {TYPES} from "../constants";
import clone from "../common/objects";

function Draw(config) {

    this.draw = (ctx, root) => {

        let W = this.GetWidth();
        let H = this.GetHeight();
        let lineWidth = 1;
        let X = this.Coords.GetX();
        let Y = this.Coords.GetY();

        ctx.lineWidth = lineWidth;

        switch (this.GetType()) {

            case TYPES.task:

                ctx.beginPath();
                ctx.moveTo(X, Y);
                ctx.lineTo(X + W, Y);
                ctx.lineTo(X + W, Y + H);
                ctx.lineTo(X, Y + H);
                ctx.lineTo(X, Y);
                // ctx.shadowBlur = 10;
                // ctx.shadowOffsetX = 20;
                // ctx.shadowColor = "black";
                // ctx.fillStyle = "red";

                addText(ctx);

                if (this.IsSelected()) {
                    ctx.strokeStyle = 'red';
                } else {
                    ctx.strokeStyle = this.GetColorDefault();
                }

                ctx.stroke();

                break;

            case TYPES.condition:

                let curX = X + W / 2;
                let curY = Y;

                ctx.beginPath();
                ctx.moveTo(curX, curY);
                ctx.lineTo(curX + W / 2, curY + H / 2);
                ctx.lineTo(curX, curY + H + (lineWidth / 2) - 1);
                ctx.lineTo(curX - W / 2, curY + H / 2);
                ctx.lineTo(curX, curY);

                addText(ctx);

                if (this.IsSelected()) {
                    ctx.strokeStyle = 'red';
                } else {
                    ctx.strokeStyle = this.GetColorDefault();
                }

                ctx.stroke();

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

                ctx.strokeStyle = 'red';

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
        }

        if (this.IsShowConnectors()) {
            drawConnectors(ctx);
        }

        addLinks(ctx);


    };

    let addText = (ctx) => {

        if (this.GetText().length) {
            ctx.fillStyle = "#000000";
            ctx.font = "10px Arial";
            ctx.fillText(this.GetText(), this.Coords.GetX(), this.Coords.GetY() + this.GetHeight() / 2);
        }
    };

    let addLinks = (ctx) => {

        let links = this.GetLinks();

        for (let i=0; i < links.length; i++) {

            let link = links[i];
            console.log('link', link);
        }
    };

    let drawConnectors = (ctx) => {

        let connectionX1 = 0;
        let connectionX2 = 0;
        let connectionY1 = 0;
        let connectionY2 = 0;

        let connectionCenter =  0;

        let drawConnector = (x, y) => {
            ctx.beginPath();
            ctx.arc(x, y, TYPES.connectionPointRadius, 0, 2 * Math.PI, false);
            ctx.fill();
            ctx.stroke();
        };

        switch (this.GetType()) {

            case TYPES.condition:
            case TYPES.task:

                let x = this.Coords.GetX();
                let y = this.Coords.GetY();

                let h = this.GetHeight();
                let w = this.GetWidth();

                ctx.fillStyle = TYPES.connectionFillStyle;
                ctx.strokeStyle = TYPES.connectionStrokeStyle;

                drawConnector(x, y + h / 2);
                drawConnector(x + w, y + h / 2);
                drawConnector(x + w / 2, y);
                drawConnector(x + w / 2, y + h);

                break;

        }


    };
}


export default Draw;