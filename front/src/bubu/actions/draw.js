import {TYPES} from "../constants";
import clone from "../common/objects";

function Draw(config) {

    this.draw = (ctx) => {

        let width = this.GetWidth();
        let height = this.GetHeight();

        let lineWidth = 1;

        ctx.lineWidth = lineWidth;

        switch (this.GetType()) {

            case TYPES.ACTION:

                ctx.beginPath();
                ctx.moveTo(this.Coords.GetX(), this.Coords.GetY());
                ctx.lineTo(this.Coords.GetX() + width, this.Coords.GetY());
                ctx.lineTo(this.Coords.GetX() + width, this.Coords.GetY() + height);
                ctx.lineTo(this.Coords.GetX(), this.Coords.GetY() + height);
                ctx.lineTo(this.Coords.GetX(), this.Coords.GetY());
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

                let curX = this.Coords.GetX() + width / 2;
                let curY = this.Coords.GetY();

                ctx.beginPath();
                ctx.moveTo(curX, curY);
                ctx.lineTo(curX + width / 2, curY + height / 2);
                ctx.lineTo(curX, curY + height + (lineWidth / 2) - 1);
                ctx.lineTo(curX - width / 2, curY + height / 2);
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
                ctx.moveTo(this.Coords.GetX(), this.Coords.GetY());
                ctx.lineTo(this.Coords.GetX() + width, this.Coords.GetY());
                ctx.lineTo(this.Coords.GetX() + width, this.Coords.GetY() + height);
                ctx.lineTo(this.Coords.GetX(), this.Coords.GetY() + height);
                ctx.lineTo(this.Coords.GetX(), this.Coords.GetY());
                ctx.fillStyle = this.GetColor();
                ctx.fill();

                break;

            case TYPES.MULTISELECTION:

                ctx.beginPath();
                ctx.moveTo(this.Coords.GetX(), this.Coords.GetY());
                ctx.lineTo(this.Coords.GetX() + width, this.Coords.GetY());
                ctx.lineTo(this.Coords.GetX() + width, this.Coords.GetY() + height);
                ctx.lineTo(this.Coords.GetX(), this.Coords.GetY() + height);
                ctx.lineTo(this.Coords.GetX(), this.Coords.GetY());

                ctx.strokeStyle = 'red';

                ctx.stroke();

                break;

            case TYPES.BACKGROUND:

                ctx.beginPath();
                ctx.moveTo(this.Coords.GetX(), this.Coords.GetY());
                ctx.lineTo(this.Coords.GetX() + width, this.Coords.GetY());
                ctx.lineTo(this.Coords.GetX() + width, this.Coords.GetY() + height);
                ctx.lineTo(this.Coords.GetX(), this.Coords.GetY() + height);
                ctx.lineTo(this.Coords.GetX(), this.Coords.GetY());

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