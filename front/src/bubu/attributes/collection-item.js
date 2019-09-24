function CollectionItem() {

    let index = 0;

    this.GetIndex = () => {
        return index;
    };

    this.SetIndex = (i) => {
        index = i;
        return this;
    };
}

export default CollectionItem;