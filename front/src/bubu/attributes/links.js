import ElementsRegister from "../elements-register";
import ConnectLink from "../actions/connect-link";
import {TYPES} from "../constants";

function Links(config) {

    this.AddLink = (sourcePoint) => {

        let link = new ElementsRegister.Link();

        link.SetLinkSourcePoint(sourcePoint);
        link.SetOnMove(new ConnectLink());
        link.Select();

        this.AddItem(link);
        this.SetSelectedItem(link);
    };

    this.RemoveBrokenLinks = () => {

        let items = this.GetItems();

        for (let i in items) {

            if (items[i].GetType() === TYPES.link) {

                let ps = items[i].GetLinkSourcePoint();
                let pd = items[i].GetLinkDestinationPoint();

                if (ps && ! this.GetItemById(ps.GetParentId())) {
                    this.RemoveItem(items[i]);
                    continue;
                }

                if (pd && ! this.GetItemById(pd.GetParentId())) {
                    this.RemoveItem(items[i]);
                }
            }
        }
    };
}

export default Links;