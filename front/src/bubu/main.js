import bubuElements from './elements'

function BuBu() {

    let Elements = {};

    function AddElement(element) {

        let id = element.GetId();

        Elements[id] = element;
        return this;
    }

    function GetNames() {

        let names = [];

        for (let i in Elements) {
            names.push(Elements[i].GetName());
        }

        return names;
    }

    function GetElementsByName(name) {

        let els = [];

        for (let i in Elements) {
            if (Elements[i].GetName() === name) {
                els.push(Elements[i]);
            }
        }

        return els;
    }

    function GetElementsByType(type) {

        let els = [];

        for (let i = 0; i < Elements.length; i++) {
            if (Elements[i].GetType() === type) {
                els.push(Elements[i]);
            }
        }

        return els;
    }

    function GetElementById(id) {

        return Elements[id];
    }

    return {
        Add: AddElement,
        Elements: bubuElements,
        GetNames: GetNames,
        GetElementsByType: GetElementsByType,
        GetElementsByName: GetElementsByName,
        GetElementById: GetElementById,
    };
}

export default BuBu;
