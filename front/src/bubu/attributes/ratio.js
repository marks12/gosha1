function Ratio() {

    this.GetXCenter = () => {
        return this.Coords.GetX() + this.GetWidth() / 2;
    };

    this.GetYCenter = () => {
        return this.Coords.GetY() + this.GetHeight() / 2;
    }
}

export default Ratio;