import Editor from "./Editor.js";
import Display from "./Display.js";

class CreatePost extends React.Component {
    constructor(props) {
        super(props);
        this.state = {
            content: "",
            fontWeight: "normal",
            fontSize: "13",
            color: "black"
        };

        this.handleChange = this.handleChange.bind(this);
        this.handleSubmit = this.handleSubmit.bind(this);
    }

    handleChange(val) {
        console.log(val);
        this.setState({
            content: val
        });
    }

    handleSubmit() {
        console.log("CLICKED");
        axios.post("http://localhost:9000/api/CreatePost", `data=${this.state.content}`).then(response => {
            console.log(response);
        }).catch(error => {
            console.log(error);
        });
    }

    changeFontWeight(event) {
        console.log("change font weight");
        this.setState({
            fontWeight: event.target.value
        });
    }

    changeFontSize(event) {
        console.log("change font size");
        this.setState({
            fontSize: event.target.value
        });
    }

    changeColor(event) {
        console.log("change color");
        this.setState({
            color: event.target.value
        });
    }

    createDiv() {
        console.log("create div");
    }

    render() {
        return React.createElement(
            "div",
            { id: "wrap" },
            React.createElement(
                "div",
                { id: "container" },
                React.createElement(
                    "div",
                    { className: "sub-container" },
                    React.createElement(Editor, { onChange: this.handleChange, content: this.state.content })
                ),
                React.createElement(
                    "div",
                    { className: "sub-container" },
                    React.createElement(Display, { text: this.state.content })
                )
            ),
            React.createElement(
                "div",
                { id: "function" },
                React.createElement(
                    "div",
                    { "class": "sub-function" },
                    React.createElement(
                        "div",
                        { "class": "container-function" },
                        React.createElement(
                            "label",
                            null,
                            "Text-color"
                        ),
                        React.createElement(
                            "select",
                            { id: "color", value: this.state.color, onChange: this.changeColor },
                            React.createElement(
                                "option",
                                { value: "red" },
                                "Red"
                            ),
                            React.createElement(
                                "option",
                                { value: "blue" },
                                "Blue"
                            ),
                            React.createElement(
                                "option",
                                { value: "green" },
                                "Green"
                            )
                        )
                    ),
                    React.createElement(
                        "div",
                        { "class": "container-function" },
                        React.createElement(
                            "label",
                            null,
                            "Font-weight"
                        ),
                        React.createElement("input", { id: "font-weight", type: "text", value: this.state.fontWeight, onChange: this.changeFontWeight })
                    ),
                    React.createElement(
                        "div",
                        { "class": "container-function" },
                        React.createElement(
                            "label",
                            null,
                            "Font-size"
                        ),
                        React.createElement("input", { id: "font-size", type: "text", value: this.state.fontSize, onChange: this.changeFontSize })
                    )
                ),
                React.createElement(
                    "button",
                    { onClick: this.handleSubmit },
                    "SUBMIT"
                )
            )
        );
    }
}

const register = React.createElement(CreatePost, null);
const root = ReactDOM.createRoot(document.getElementById('root'));
root.render(register);