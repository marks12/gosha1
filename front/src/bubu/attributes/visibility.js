function Visibility (config) {

    let isVisibile = config && typeof config.Visibility === "boolean" ? config.Visibility : true;

    /**
     * @param bool
     * @returns {Visibility}
     * @constructor
     */
    this.SetVisibility = (bool) => {
        isVisibile = bool;
        return this;
    };

    /**
     * @returns {boolean}
     * @constructor
     */
    this.GetVisibility = () => {
        return isVisibile;
    };
}

export default Visibility;