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
        return Math.abs(onLineY - y) <= delta;
    }

    let onLineX = (y - y1) * (x2 - x1) / (y2 - y1)  + x1;

    console.log('by x', onLineX, x);

    return Math.abs(onLineX - x) <= delta;
}

export function rounder(a) {
    return Math.round(a / TYPES.positionRound) * TYPES.positionRound;
}

export function generateUUID() { // Public Domain/MIT
    let d = new Date().getTime();//Timestamp
    let d2 = (performance && performance.now && (performance.now()*1000)) || 0;//Time in microseconds since page-load or 0 if unsupported
    return 'xxxxxxxx-xxxx-4xxx-yxxx-xxxxxxxxxxxx'.replace(/[xy]/g, function(c) {
        let r = Math.random() * 16;//random number between 0 and 16
        if(d > 0){//Use timestamp until depleted
            r = (d + r)%16 | 0;
            d = Math.floor(d/16);
        } else {//Use microseconds since page-load if supported
            r = (d2 + r)%16 | 0;
            d2 = Math.floor(d2/16);
        }
        return (c === 'x' ? r : (r & 0x3 | 0x8)).toString(16);
    });
}