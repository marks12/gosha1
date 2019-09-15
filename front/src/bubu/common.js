
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

