import ElementsRegister from "../elements-register";
import ConnectLink from "../actions/connect-link";

function Links(config) {

    this.AddLink = (sourcePoint) => {

        let link = new ElementsRegister.Link();

        link.SetLinkSourcePoint(sourcePoint);
        link.SetOnMove(new ConnectLink());
        link.Select();

        this.AddItem(link);
        this.SetSelectedItem(link);
    };

    this.GetLinks = () => {
        return [];
    };
}

export default Links;