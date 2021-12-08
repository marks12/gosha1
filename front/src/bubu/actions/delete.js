function Delete() {

    let onDelete = null;

    this.GetOnDelete = () => {
        return onDelete;
    };

    this.SetOnDelete = (f) => {
        onDelete = f;
        return this;
    }
}

export default Delete;