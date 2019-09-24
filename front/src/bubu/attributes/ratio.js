function Ratio() {

    this.GetXCenter = () => {
        return this.GetX() + this.GetWidth() / 2;
    };

    this.GetYCenter = () => {
        return this.GetY() + this.GetHeight() / 2;
    }
}

export default Ratio;