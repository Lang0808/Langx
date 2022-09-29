class Login extends React.Component {
    constructor(props) {
        super(props);
        this.state = {};
    }

    render() {
        return React.createElement(Form, null);
    }
}

class Form extends React.Component {
    constructor(props) {
        super(props);
        this.state = {
            Username: "",
            Password: ""
        };

        this.handleChangeUsername = this.handleChangeUsername.bind(this);
        this.handleChangePassword = this.handleChangePassword.bind(this);
        this.handleSubmit = this.handleSubmit.bind(this);
    }

    handleChangeUsername(event) {
        this.setState({ Username: event.target.value });
    }

    handleChangePassword(event) {
        this.setState({ Password: event.target.value });
    }

    handleSubmit(event) {
        alert('A name was submitted: ' + this.state.Username + ' ' + this.state.Password);
        event.preventDefault();
        axios.post("http://localhost:9000/api/LoginUser", `Username=${this.state.Username}&Password=${this.state.Password}`).then(response => console.log(response));
    }

    render() {
        return React.createElement(
            "form",
            { onSubmit: this.handleSubmit },
            React.createElement(
                "label",
                null,
                "Username:",
                React.createElement("input", { type: "text", value: this.state.Username, onChange: this.handleChangeUsername })
            ),
            React.createElement(
                "label",
                null,
                "Password:",
                React.createElement("input", { type: "password", value: this.state.Password, onChange: this.handleChangePassword })
            ),
            React.createElement("input", { type: "submit", value: "Submit" })
        );
    }

}

const register = React.createElement(Login, null);
const root = ReactDOM.createRoot(document.getElementById('root'));
root.render(register);