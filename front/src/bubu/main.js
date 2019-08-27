import bubuElements from './elements'

function BuBu(canvasElementId) {

    let canvas = document.getElementById(canvasElementId);

    if (!canvas || !canvas.getContext) {
        console.error("Wrong canvas element Id. ELement not found or canvas.getContext function not exists");
        return
    }

    canvas.setAttribute("width", canvas.parentNode.parentElement.clientWidth);
    canvas.setAttribute("height", canvas.parentNode.parentElement.clientHeight);

    let isDown = false;
    let selectedElement = null;
    let canvasOffsetX = canvas.getBoundingClientRect().left;
    let canvasOffsetY = canvas.getBoundingClientRect().top;
    let selectedElementOffsetX = 0;
    let selectedElementOffsetY = 0;
    let ctx = canvas.getContext('2d');
    let Elements = {};

    function getFirstElementByCoordinates(x, y) {

        for (let i in Elements) {

            let x1 = Elements[i].Coords.GetX();
            let x2 = x1 + Elements[i].GetWidth();

            let y1 = Elements[i].Coords.GetY();
            let y2 = y1 + Elements[i].GetHeight();

            if (x - canvasOffsetX >= x1 && x - canvasOffsetX <= x2 && y - canvasOffsetY >= y1 && y - canvasOffsetY <= y2) {

                selectedElementOffsetX = x - canvasOffsetX - x1;
                selectedElementOffsetY = y - canvasOffsetY - y1;

                return Elements[i];
            }
        }

        return null;
    }


    function down(event) {

        isDown = true;
        event = assignCoordinates(event);

        let x = event.pageX;
        let y = event.pageY;

        selectedElement = getFirstElementByCoordinates(x, y);
    }

    function up() {
        isDown = false;
        selectedElement = null;
        selectedElementOffsetX = 0;
        selectedElementOffsetY = 0;
    }

    function assignCoordinates(event) {

        let eventDoc, doc, body;

        event = event || window.event; // IE-ism

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

        return event;
    }

    function mover(event) {

        event = assignCoordinates(event);

        if (isDown && selectedElement) {

            selectedElement.Coords.SetX(event.pageX - canvasOffsetX - selectedElementOffsetX);
            selectedElement.Coords.SetY(event.pageY - canvasOffsetY - selectedElementOffsetY);

            Render()
        }
    }

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

    canvas.addEventListener("mousedown", down);
    canvas.addEventListener("mousemove", mover);
    canvas.addEventListener("mouseup", up);

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
