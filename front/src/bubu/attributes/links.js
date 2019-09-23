import ElementsRegister from "../elements-register";
import ConnectLink from "../actions/connect-link";

function Links(config) {

    this.AddLink = (x, y) => {

        let link = new ElementsRegister.Link({
            Coords: {
                X: x,
                Y: y,
            },
        });

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