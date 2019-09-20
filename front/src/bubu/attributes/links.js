import Link from "../elements/Link";

function Links(config) {

    Link.apply(this, arguments);

    let links = config && config.links && config.links.length ? config.links : [];

    /**
     * @param isSource bool
     * @param lineType int
     * @param arrowType int
     */
    this.AddLink = (isSource, lineType, arrowType) => {

        switch (linkType) {

        }
    };

    this.GetLinks = () => {
        return links;
    };
}

export default Links;