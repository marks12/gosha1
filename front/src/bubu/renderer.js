function Renderer(config) {

    this.clearAll = () => {

        let ctx = this.GetCtx();
        let canvas = this.GetCanvas();

        ctx.clearRect(0, 0, canvas.width, canvas.height);
    };

    this.Render = () => {

        this.clearAll();

        let ctx = this.GetCtx();
        let Items = this.GetItems();

        for (let i in Items) {
            Items[i].draw(ctx);
        }
    };

    this.AssignCoordinates = (e) => {

        let eventDoc, doc, body;

        event = e || window.event; // IE-ism

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
}

export default Renderer;