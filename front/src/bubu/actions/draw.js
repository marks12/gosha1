import {TYPES} from "../constants";

function Draw(config) {

    let width = this.GetWidth();
    let height = this.GetHeight();

    this.draw = (ctx) => {

        let lineWidth = 1;

        ctx.lineWidth = lineWidth;

        switch (this.GetType())  {

            case TYPES.ACTION:

                ctx.beginPath();
                ctx.moveTo(this.Coords.GetX(), this.Coords.GetY());
                ctx.lineTo(this.Coords.GetX() + width, this.Coords.GetY());
                ctx.lineTo(this.Coords.GetX() + width, this.Coords.GetY() + height);
                ctx.lineTo(this.Coords.GetX(), this.Coords.GetY() + height);
                ctx.lineTo(this.Coords.GetX(), this.Coords.GetY());
                ctx.stroke();

                break;

            case TYPES.CONDITION:

                console.log('draw condition');

                let curX = this.Coords.GetX() + width/2;
                let curY = this.Coords.GetY();

                ctx.beginPath();
                ctx.moveTo(curX, curY);
                ctx.lineTo(curX + width / 2, curY + height / 2);
                ctx.lineTo(curX, curY + height + (lineWidth / 2) - 1);
                ctx.lineTo(curX - width/2, curY + height / 2);
                ctx.lineTo(curX, curY);
                ctx.stroke();

                break;
        }
    };
}

export default Draw;