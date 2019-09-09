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
            Items[i].draw(ctx, this);
        }
    };


}

export default Renderer;