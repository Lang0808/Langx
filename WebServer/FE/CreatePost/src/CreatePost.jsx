import Editor from "./Editor.js";
import Display from "./Display.js";

class CreatePost extends React.Component{
    constructor(props) {
        super(props);
        this.state={
            content: "",
            fontWeight: "normal",
            fontSize: "13",
            color: "black",
        };

        this.handleChange=this.handleChange.bind(this);
        this.handleSubmit=this.handleSubmit.bind(this);
    }

    handleChange(val){
        console.log(val);
        this.setState({
            content: val
        });
    }

    handleSubmit(){
        console.log("CLICKED");
        axios.post("http://localhost:9000/api/CreatePost", `data=${this.state.content}`)
            .then((response)=>{console.log(response)})
            .catch((error)=>{console.log(error)});
    }

    changeFontWeight(event){
        console.log("change font weight")
        this.setState({
            fontWeight: event.target.value
        });
    }

    changeFontSize(event){
        console.log("change font size")
        this.setState({
            fontSize: event.target.value
        });
    }

    changeColor(event){
        console.log("change color")
        this.setState({
            color: event.target.value
        });
    }

    createDiv(){
        console.log("create div")

    }

    render(){
        return (
            <div id="wrap">
                <div id="container">
                    <div className="sub-container">
                        <Editor onChange={this.handleChange} content={this.state.content}></Editor>
                    </div>
                    <div className="sub-container">
                        <Display text={this.state.content}></Display>
                    </div>
                </div>
                <div id="function">
                    <div class="sub-function">
                        <div class="container-function">
                            <label>Text-color</label>
                            <select id="color" value={this.state.color} onChange={this.changeColor}>
                                <option value="red">Red</option>
                                <option value="blue">Blue</option>
                                <option value="green">Green</option>
                            </select>
                        </div>
                        <div class="container-function">
                            <label>Font-weight</label>
                            <input id="font-weight" type="text" value={this.state.fontWeight} onChange={this.changeFontWeight}/>
                        </div>
                        <div class="container-function">
                            <label>Font-size</label>
                            <input id="font-size" type="text" value={this.state.fontSize} onChange={this.changeFontSize}/>
                        </div>
                    </div>
                    <button onClick={this.handleSubmit}>SUBMIT</button>
                </div>
            </div>
        )
    }
}

const register=<CreatePost/>;
const root = ReactDOM.createRoot(document.getElementById('root'));
root.render(register);