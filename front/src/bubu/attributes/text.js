function Text(config) {

    let text = config && config.Text ? config.Text : '';

    this.GetText = () => {return text};

    this.SetText = (t) => {text = t; return this;};

}

export default Text;