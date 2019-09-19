function Links(config) {

    let links = config && config.links && config.links.length ? config.links : [];

    /**
     * \
     * @param linkedId
     * @param linkType default
     * @constructor
     */
    this.AddLink = (linkedId, linkType) => {

        switch (linkType) {

        }

    };

    this.GetLinks = () => {
        return links;
    };

}

export default Links;