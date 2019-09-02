function Move(config) {

    let OnMoveAction = null;

    if (config && config.OnMove) {
        OnMoveAction = config.OnMove;
    }

    let self = this;

    this.setOnMove = (action) => {
        self.OnMoveAction = action;
        return self;
    };

    this.getOnMove = () => {
        return OnMoveAction;
    };

    this.Run = (x, y) => {
        console.log('run in Move action');
    }
}

export default Move;