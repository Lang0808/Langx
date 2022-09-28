class Register extends React.Component{
    constructor(props) {
        super(props);
        this.state={
        };
    }

    render(){
        return (
            <Form></Form>
        );
    }
}

class Form extends React.Component{
    constructor(props) {
        super(props);
        this.state = {
            Username: "",
            Password: "",
        }

        this.handleChangeUsername=this.handleChangeUsername.bind(this)
        this.handleChangePassword=this.handleChangePassword.bind(this)
        this.handleSubmit=this.handleSubmit.bind(this)
    }

    handleChangeUsername(event){
        this.setState(
            {Username: event.target.value}
        );
    }

    handleChangePassword(event){
        this.setState(
            {Password: event.target.value}
        );
    }

    handleSubmit(event) {
        alert('A name was submitted: ' + this.state.Username+' '+this.state.Password);
        event.preventDefault();
        axios.post("http://localhost:9000/api/RegisterUser",
            `Username=${this.state.Username}&Password=${this.state.Password}`)
            .then((response) => console.log(response));
    }

    render(){
        return (
            <form onSubmit={this.handleSubmit}>
                <label>
                    Username:
                    <input type="text" value={this.state.Username} onChange={this.handleChangeUsername}/>
                </label>
                <label>
                    Password:
                    <input type="password" value={this.state.Password} onChange={this.handleChangePassword}/>
                </label>
                <input type="submit" value="Submit" />
            </form>
        )
    }

}

const register=<Register/>;
const root = ReactDOM.createRoot(document.getElementById('root'));
root.render(register);