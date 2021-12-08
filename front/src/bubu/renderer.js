function Renderer(config) {

    this.clearAll = () => {

        let ctx = this.GetCtx();
        let w = this.GetCanvasWidth();
        let h = this.GetCanvasHeight();
        let scale = this.GetScale();

        if (scale < 1) {
            scale = 1;
        }

        ctx.clearRect(-w * scale, -h * scale, w * scale * 100, h * scale * 100);
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