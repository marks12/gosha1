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

            case TYPES.TASK:

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

            case TYPES.CONDITION:

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

            case TYPES.DIVIDER:

                ctx.beginPath();
                ctx.moveTo(X, Y);
                ctx.lineTo(X + W, Y);
                ctx.lineTo(X + W, Y + H);
                ctx.lineTo(X, Y + H);
                ctx.lineTo(X, Y);
                ctx.fillStyle = this.GetColor();
                ctx.fill();

                break;

            case TYPES.MULTISELECTION:

                ctx.beginPath();
                ctx.moveTo(X, Y);
                ctx.lineTo(X + W, Y);
                ctx.lineTo(X + W, Y + H);
                ctx.lineTo(X, Y + H);
                ctx.lineTo(X, Y);

                ctx.strokeStyle = 'red';

                ctx.stroke();

                break;

            case TYPES.BACKGROUND:

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
    };

    let addText = (ctx) => {

        if (this.GetText().length) {
            ctx.fillStyle = "#000000";
            ctx.font = "10px Arial";
            ctx.fillText(this.GetText(), this.Coords.GetX(), this.Coords.GetY() + this.GetHeight() / 2);
        }
    }
}


export default Draw;