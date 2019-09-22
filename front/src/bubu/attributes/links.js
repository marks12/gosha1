import ElementsRegister from "../elements-register";

function Links(config) {

    this.AddLink = (x, y) => {

        let link = new ElementsRegister.Link({
            Coords: {
                X: x,
                Y: y,
            },
        });

        link.Select();

        this.AddItem(link);
        this.SetSelectedItem(link);
    };

    this.GetLinks = () => {
        return [];
    };
}

export default Links;