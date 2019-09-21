let connectionPointRadius = 10;
let spaceBetween = 20;

export const TYPES = {
    task : 100,
    condition : 200,
    arrow : 300,
    divider : 400,
    multiselection : 500,
    background : 600,
    zeroPoint : 700,
    eventStart : 800,
    eventIntermediate : 810,
    eventEnd : 820,

    //toolbox
    stdHeight: 40,
    spaceBetween: spaceBetween,
    toolboxWidth: 0,

    //links
    linkTypeLine: 10100, // _________________
    linkTypeDashes: 10110, // ---------------
    linkTypeDots: 10120, // .................

    arrowTypeRhombus: 11100, // <>------
    arrowTypeNone: 11110, // ---------
    arrowTypeSimple: 11120, // -------->
    arrowTypeFilled: 11130, // --------|>
    arrowTypeCircle: 11140, // --------o                 /
    arrowTypeMany: 11150, //                     -----------
                               //                        \

    arrowTypeOne: 11160, // --------¹
    arrowTypeRoot: 11170, // -/------

    connectionPointRadius: connectionPointRadius,
    activeSpaceAround: connectionPointRadius + spaceBetween,

    connectionFillStyle: "#eee",
    connectionStrokeStyle: "#999",

};