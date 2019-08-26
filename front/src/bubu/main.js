import bubuElements from './elements'

function BuBu(canvasElementId) {

    let canvas = document.getElementById(canvasElementId);

    if (!canvas || !canvas.getContext) {
        console.error("Wrong canvas element Id. ELement not found or canvas.getContext function not exists");
        return
    }

    let isDown = false;

    function down() {
        isDown = true;
    }
    function up() {
        isDown = false;
    }

    function mover(event) {

        let eventDoc, doc, body;

        event = event || window.event; // IE-ism

        // If pageX/Y aren't available and clientX/Y are,
        // calculate pageX/Y - logic taken from jQuery.
        // (This is to support old IE)
        if (event.pageX == null && event.clientX != null) {
            eventDoc = (event.target && event.target.ownerDocument) || document;
            doc = eventDoc.documentElement;
            body = eventDoc.body;

            event.pageX = event.clientX +
                (doc && doc.scrollLeft || body && body.scrollLeft || 0) -
                (doc && doc.clientLeft || body && body.clientLeft || 0);
            event.pageY = event.clientY +
                (doc && doc.scrollTop  || body && body.scrollTop  || 0) -
                (doc && doc.clientTop  || body && body.clientTop  || 0 );
        }


        if (isDown) {
            for (let i in Elements) {
                Elements[i].Coords.SetX(event.pageX - 48 - 20);
                Elements[i].Coords.SetY(event.pageY - 72 - 30);
                break;
            }

            Render()
        }

    }



    canvas.addEventListener("mousedown", down);
    canvas.addEventListener("mousemove", mover);
    canvas.addEventListener("mouseup", up);

    let ctx = canvas.getContext('2d');
    let Elements = {};

    function AddElement(element) {

        let id = element.GetId();
        Elements[id] = element;
        return this;
    }

    function GetNames() {

        let names = [];

        for (let i in Elements) {
            names.push(Elements[i].GetName());
        }

        return names;
    }

    function GetElementsByName(name) {

        let els = [];

        for (let i in Elements) {

            if (Elements[i].GetName() === name) {
                els.push(Elements[i]);
            }
        }
        return els;
    }

    function GetElementsByType(type) {

        let els = [];

        for (let i = 0; i < Elements.length; i++) {

            if (Elements[i].GetType() === type) {
                els.push(Elements[i]);
            }
        }
        return els;
    }

    function GetElementById(id) {

        return Elements[id];
    }

    function clearAll() {
        ctx.clearRect(0, 0, canvas.width, canvas.height);
    }

    function Render() {

        clearAll();

        for (let i in Elements) {
            Elements[i].draw(ctx);
        }
    }


    return {
        Add: AddElement,
        Elements: bubuElements,
        GetNames: GetNames,
        GetElementsByType: GetElementsByType,
        GetElementsByName: GetElementsByName,
        GetElementById: GetElementById,
        Render: Render,
    };
}

export default BuBu;
