export const TYPES = {
    TASK : 100,
    CONDITION : 200,
    ARROW : 300,
    DIVIDER : 400,
    MULTISELECTION : 500,
    BACKGROUND : 600,
    ZEROPOINT : 700,

    //toolbox
    stdHeight: 40,
    spaceBetween: 0,
    toolboxWidth: 0,

    //links
    linkTypeSequenceFlowA: 10100, // ___>
    linkTypeSequenceFlowB: 10200, // <____
    linkTypeSequenceFlowC: 10300, // <____>

    linkTypeConditionalFlowA: 10400, // <>________>
    linkTypeConditionalFlowB: 10500, // <________<>

    linkTypeControlFlowA: 10600, // __\__________>
    linkTypeControlFlowB: 10700, // <__________/__

    linkTypeAssocA: 10800, // .......|>
    linkTypeAssocB: 10900, // <|.......
    linkTypeAssocC: 11000, // .........

    linkTypeMessagesFlowA: 11100, // []-------|>
    linkTypeMessagesFlowB: 11200, // <|-------[]
};