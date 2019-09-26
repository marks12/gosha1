function Click(config) {

    let OnDownAction = null;
    let OnUpAction = null;

    if (config && config.OnDown) {
        OnDownAction = config.OnDown;
    }

    if (config && config.OnUp) {
        OnUpAction = config.OnUp;
    }

    this.SetOnDown = (action) => {

        OnDownAction = action;
        return this;
    };

    this.GetOnDown = () => {
        return OnDownAction;
    };

    this.SetOnUp = (action) => {

        OnUpAction = action;
        return this;
    };

    this.GetOnUp = () => {
        return OnUpAction;
    };

    this.Run = (x, y) => {
        console.log('run in Move action');
    }
}

export default Click;