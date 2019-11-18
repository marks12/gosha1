import {TYPES} from "./constants";

export function IsRightButton(e) {

    let isRightMB;
    e = e || window.event;

    if ("which" in e)  // Gecko (Firefox), WebKit (Safari/Chrome) & Opera
        isRightMB = e.which === 3;
    else if ("button" in e)  // IE, Opera
        isRightMB = e.button === 2;

    return isRightMB;
}

/**
 * @return {boolean}
 */
export function IsMiddleButtton(e) {
    return !!(e && (e.which === 2 || e.button === 4));
}

/**
 * @return {boolean}
 */
export function IsPointOnLine(ax, ay, bx, by, x, y, delta) {
    /**
     *       (y2-y1)
     * y = ----------- (x - x1) + y1
     *       (x2-x1)
     */

    /**
     *       y(x2 - x1) - y1(x2 + y1)
     * x = --------------------------- + x1
     *              y2 - y1
     */

    let x1, x2, y1, y2;

    x1 = ax;
    x2 = bx;

    y1 = ay;
    y2 = by;


    if (Math.abs(Math.abs(x2) - Math.abs(x1)) < delta) {

        let checkY = false;

        if (y1 > y2) {
            checkY = Math.abs(y) < Math.abs(y1) && Math.abs(y) > Math.abs(y2)
        } else {
            checkY = Math.abs(y) > Math.abs(y1) && Math.abs(y) < Math.abs(y2)
        }

        return checkY && Math.abs(Math.abs(x) - Math.abs(x1)) < delta / 2;
    }

    if (Math.abs(Math.abs(y2) - Math.abs(y1)) < delta) {

        let checkX = false;

        if (x1 < x2) {
            checkX = Math.abs(x) > Math.abs(x1) && Math.abs(x) < Math.abs(x2)
        } else {
            checkX = Math.abs(x) < Math.abs(x1) && Math.abs(x) > Math.abs(x2)
        }

        return checkX && Math.abs(Math.abs(y) - Math.abs(y1)) < delta / 2;
    }

    if (Math.abs(Math.abs(x2) - Math.abs(x1)) > delta) {
        let onLineY = ((y2 - y1) / (x2 - x1)) * (x - x1) + y1;

        console.log('by y', onLineY, y);

        return Math.abs(onLineY - y) <= delta;
    }

    let onLineX = (y - y1) * (x2 - x1) / (y2 - y1)  + x1;

    console.log('by x', onLineX, x);

    return Math.abs(onLineX - x) <= delta;
}

export function rounder(a) {
    return Math.round(a / TYPES.positionRound) * TYPES.positionRound;
}