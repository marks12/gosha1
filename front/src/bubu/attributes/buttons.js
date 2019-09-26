import ElementsRegister from '../elements-register'
import {TYPES as constants, TYPES as Const} from "../constants";

function Buttons() {

    let buttons = [];

    let isShowButtons = false;

    this.IsShowButtons = () => {
        return isShowButtons;
    };

    this.SetShowButtons = (bool) => {
        isShowButtons = bool;
        return this;
    };

    this.ShowElementButtons = (x, y) => {

        this.HideAllButtons();

        let item = this.GetFirstElementByCoordinates(
            x,
            y
        );

        if (item) {
            item.ShowButtons();
        } else {
            this.HideAllButtons();
        }

    };

    this.HideAllButtons = () => {

        let items = this.GetSelectableItems();

        for (let i in items) {
            items[i].HideButtons();
        }
    };

    this.AddButton = ({positionVertical, positionHorizontal, OnDown, OnMove, OnUp}) => {

        let x = () => {console.error('wrong x getter');};
        let y = () => {console.error('wrong y getter');};

        let item = this;

        if (positionHorizontal === constants.right) {
            x = () => {return item.Coords.GetX() + item.GetWidth() - constants.spaceBetween};
        } else {
            x = () => {return item.Coords.GetX() + constants.spaceBetween};
        }

        if (positionVertical === constants.top) {
            y = () => {return item.Coords.GetY() + constants.spaceBetween};
        } else {
            y = () => {return item.Coords.GetY() + item.GetHeight() - constants.spaceBetween};
        }

        let b = new ElementsRegister.Button()
            .SetVisibility(true)
            .SetWidth(constants.connectionPointRadius)
            .SetHeight(constants.connectionPointRadius);

        if (OnDown) {
            b.SetOnDown(OnDown);
        }

        if (OnMove) {
            b.SetOnMove(OnMove);
        }

        if (OnUp) {
            b.SetOnUp(OnUp);
        }

        b.Coords.GetX = x;
        b.Coords.GetY = y;


        buttons.push(b);

        return b;
    };

    this.ClearButtons = () => {
        buttons = [];
    };

    this.FindButton = (x, y) => {

        let items = this.GetSelectableItems();

        for (let index in items) {

            let btns = items[index].GetButtons();

            for (let i in btns) {
                if (
                    btns[i].Coords.GetX() - constants.activeSpaceAround <= x &&
                    btns[i].Coords.GetY() - constants.activeSpaceAround <= y &&
                    btns[i].Coords.GetX() + btns[i].GetWidth() + constants.activeSpaceAround >= x &&
                    btns[i].Coords.GetY() + btns[i].GetWidth() + constants.activeSpaceAround >= y) {

                    return btns[i];
                }
            }

        }

        return null;
    };


    this.GetButtons = () => {
        return buttons;
    };

    this.HideButtons = () => {

        isShowButtons = false;

        for (let i=0; i<buttons; i++) {
            buttons[i].SetVisibility(false);
        }
    };

    this.ShowButtons = () => {

        isShowButtons = true;

        for (let i=0; i<buttons.length; i++) {
            buttons[i].SetVisibility(true);
        }
    };
}

export default Buttons;