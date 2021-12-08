import ElementsRegister from '../elements-register'
import {TYPES as constants} from "../constants";

function Buttons(config) {

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

    this.AddButton = ({positionVertical, positionHorizontal, OnDown, OnMove, OnUp, border, radius}) => {

        if (border === undefined) {
            border = false;
        } else {
            border = !! (border);
        }

        if (! radius) {
            border = constants.buttonRadius;
        } else {
            border = radius * 1;
        }

        let x = () => {console.error('wrong x getter');};
        let y = () => {console.error('wrong y getter');};

        let item = this;

        switch (positionHorizontal) {

            case constants.center:
                x = () => {return item.Coords.GetX() + item.GetWidth() / 2};
                break;

            case constants.right:
                x = () => {return item.Coords.GetX() + item.GetWidth() - constants.spaceBetween};
                break;

            default:
                x = () => {return item.Coords.GetX() + constants.spaceBetween};
                break;

        }

        switch (positionVertical) {

            case constants.center:
                y = () => {return item.Coords.GetY() + item.GetHeight() / 2};
                break;

            case constants.top:
                y = () => {return item.Coords.GetY() + constants.spaceBetween};
                break;

            default:
                y = () => {return item.Coords.GetY() + item.GetHeight() - constants.spaceBetween};
                break;

        }

        if (positionVertical === constants.top) {
        } else {
        }

        let b = new ElementsRegister.Button()
            .SetVisibility(true)
            .SetWidth(radius)
            .SetHeight(radius);

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