function CollectionItem() {

    let side = 'left';

    this.GetSide = () => {
        return side;
    };

    this.SetIndex = (s) => {
        side = s;
        return this;
    };
}

export default CollectionItem;