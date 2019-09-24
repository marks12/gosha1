import {TYPES} from "../constants";
import clone from "../common/objects";

function Draw(config) {

    let self = this;

    this.draw = (ctx, root) => {

        let W = this.GetWidth();
        let H = this.GetHeight();
        let lineWidth = 2;
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
                ctx.lineTo(X, Y - lineWidth / 2);
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
                ctx.lineTo(curX + lineWidth / 2 + lineWidth / 5, curY);

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
                    ctx.beginPath();
                    ctx.moveTo(src.Coords.GetX(), src.Coords.GetY());
                    ctx.lineTo(x2, y2);
                    ctx.stroke();
                }

                break;
        }

        if (this.IsShowConnectors()) {
            drawConnectors(ctx, root);
        }

        if (root) {
            addLinks(ctx, root);
        }

    };

    let addText = (ctx) => {

        if (this.GetText().length) {
            ctx.fillStyle = "#000000";
            ctx.font = "10px Arial";
            ctx.fillText(this.GetText(), this.Coords.GetX(), this.Coords.GetY() + this.GetHeight() / 2);
        }
    };

    let addLinks = (ctx, root) => {

        let links = root.GetLinks();

        for (let i=0; i < links.length; i++) {

            let link = links[i];
            console.log('link', link);
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

                let points = this.GetConnectorPoints();

                for (let i=0;i<points.length;i++) {
                    drawConnector(points[i].Coords.GetX(), points[i].Coords.GetY());
                }

                break;

        }


    };
}


export default Draw;