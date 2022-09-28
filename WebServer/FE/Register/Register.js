var _createClass = function () { function defineProperties(target, props) { for (var i = 0; i < props.length; i++) { var descriptor = props[i]; descriptor.enumerable = descriptor.enumerable || false; descriptor.configurable = true; if ("value" in descriptor) descriptor.writable = true; Object.defineProperty(target, descriptor.key, descriptor); } } return function (Constructor, protoProps, staticProps) { if (protoProps) defineProperties(Constructor.prototype, protoProps); if (staticProps) defineProperties(Constructor, staticProps); return Constructor; }; }();

function _classCallCheck(instance, Constructor) { if (!(instance instanceof Constructor)) { throw new TypeError("Cannot call a class as a function"); } }

function _possibleConstructorReturn(self, call) { if (!self) { throw new ReferenceError("this hasn't been initialised - super() hasn't been called"); } return call && (typeof call === "object" || typeof call === "function") ? call : self; }

function _inherits(subClass, superClass) { if (typeof superClass !== "function" && superClass !== null) { throw new TypeError("Super expression must either be null or a function, not " + typeof superClass); } subClass.prototype = Object.create(superClass && superClass.prototype, { constructor: { value: subClass, enumerable: false, writable: true, configurable: true } }); if (superClass) Object.setPrototypeOf ? Object.setPrototypeOf(subClass, superClass) : subClass.__proto__ = superClass; }

var Register = function (_React$Component) {
    _inherits(Register, _React$Component);

    function Register(props) {
        _classCallCheck(this, Register);

        var _this = _possibleConstructorReturn(this, (Register.__proto__ || Object.getPrototypeOf(Register)).call(this, props));

        _this.state = {};
        return _this;
    }

    _createClass(Register, [{
        key: "render",
        value: function render() {
            return React.createElement(Form, null);
        }
    }]);

    return Register;
}(React.Component);

var Form = function (_React$Component2) {
    _inherits(Form, _React$Component2);

    function Form(props) {
        _classCallCheck(this, Form);

        var _this2 = _possibleConstructorReturn(this, (Form.__proto__ || Object.getPrototypeOf(Form)).call(this, props));

        _this2.state = {
            Username: "",
            Password: ""
        };

        _this2.handleChangeUsername = _this2.handleChangeUsername.bind(_this2);
        _this2.handleChangePassword = _this2.handleChangePassword.bind(_this2);
        _this2.handleSubmit = _this2.handleSubmit.bind(_this2);
        return _this2;
    }

    _createClass(Form, [{
        key: "handleChangeUsername",
        value: function handleChangeUsername(event) {
            this.setState({ Username: event.target.value });
        }
    }, {
        key: "handleChangePassword",
        value: function handleChangePassword(event) {
            this.setState({ Password: event.target.value });
        }
    }, {
        key: "handleSubmit",
        value: function handleSubmit(event) {
            alert('A name was submitted: ' + this.state.Username + ' ' + this.state.Password);
            event.preventDefault();
            axios.post("http://localhost:9000/api/RegisterUser", "Username=" + this.state.Username + "&Password=" + this.state.Password).then(function (response) {
                return console.log(response);
            });
        }
    }, {
        key: "render",
        value: function render() {
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
    }]);

    return Form;
}(React.Component);

var register = React.createElement(Register, null);
var root = ReactDOM.createRoot(document.getElementById('root'));
root.render(register);