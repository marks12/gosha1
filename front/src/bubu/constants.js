let connectionPointRadius = 5;
let spaceBetween = 20;

let egyptianTheme = {
    primaryFill: '#fef7dc',
    secondaryFill: '#e4ba70',
    primaryStroke: '#080909',
    secondaryStroke: '#0c1f26',
    primaryContrast: '#dd000d',
    secondaryContrast: '#00569B',
};

let defaultTheme = {
    primaryFill: '#eee',
    secondaryFill: '#fefefe',
    primaryStroke: '#eee',
    secondaryStroke: '#999',
    primaryContrast: '#f00',
    secondaryContrast: '#00f',
};

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
    link : 900,
    connectorPoint : 910,
    button : 1000,

    //toolbox
    stdHeight: 40,
    lineWidth: 1.5,
    spaceBetween: spaceBetween,
    toolboxWidth: 0,

    // button
    buttonRadius: 10,

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

    elementWidth: 120,
    connectionPointRadius: connectionPointRadius,
    cornerRadius: 5,
    activeSpaceAround: connectionPointRadius * 4,
    roundCoords: 1,

    connectionFillStyle: "#eee",
    connectionStrokeStyle: "#999",

    top: "top",
    bottom: "bottom",
    center: "center",
    left: "left",
    right: "right",

    // font

    fontColor: "#00569B",
    fontStyle: "14px Arial",

    currentTheme: {
        primaryFill: egyptianTheme.primaryFill,
        secondaryFill: egyptianTheme.secondaryFill,
        primaryStroke: egyptianTheme.primaryStroke,
        secondaryStroke: egyptianTheme.secondaryStroke,
        primaryContrast: egyptianTheme.primaryContrast,
        secondaryContrast: egyptianTheme.secondaryContrast,
    },

};