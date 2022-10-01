class Display extends React.Component {
    constructor(props) {
        super(props);
    }

    render() {
        return React.createElement("div", { dangerouslySetInnerHTML: { __html: `${this.props.text}` } });
    }
}

export default Display;