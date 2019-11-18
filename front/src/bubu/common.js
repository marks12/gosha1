
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
export function IsPointOnLine(x1, y1, x2, y2, x, y, delta) {
    /**
     *       (y2-y1)
     * y = ----------- (x - x1) + y1
     *       (x2-x1)
     */

    let onLineY = ((y2 - y1) / (x2 - x1)) * (x - x1) + y1;

    return (Math.abs(onLineY) - y) <= delta;
}