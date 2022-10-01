class Editor extends React.Component{
    constructor(props) {
        super(props);
        this.state = {
            selectionStart: 0,
            selectionEnd: 0
        }

        /*this.handleKeyDown=this.handleKeyDown.bind(this);
        this.handleClick=this.handleClick.bind(this);*/

        this.handleChange=this.handleChange.bind(this);
        this.handleKeyDown=this.handleKeyDown.bind(this);
    }

/*    handleKeyDown(event){

        var code=event.code;
        if(code==="Tab"){
            console.log("Tab");
            var content=this.insertAtCursor("\t");
            this.props.onChange(content);
        }
        else{
            var content=this.insertAtCursor(event.key);
            this.props.onChange(content);
        }
        this.setState({
            selectionStart: this.state.selectionStart+1,
            selectionEnd: this.state.selectionEnd+1,
        })
        this.setCaretToPos(document.getElementById("textarea"), this.state.selectionStart);
    }
cd
    handleClick(event){
        var start=document.getElementById("textarea").selectionStart;
        var end=document.getElementById("textarea").selectionEnd;
        console.log(start+" "+end)
        this.setState({
            selectionStart: start,
            selectionEnd: end,
        })
        this.setCaretToPos(document.getElementById("textarea"), this.state.selectionStart);
    }

    function setSelectionRange(input, selectionStart, selectionEnd) {
        if (input.setSelectionRange) {
            input.focus();
            input.setSelectionRange(selectionStart, selectionEnd);
        }
        else if (input.createTextRange) {
            var range = input.createTextRange();
            range.collapse(true);
            range.moveEnd('character', selectionEnd);
            range.moveStart('character', selectionStart);
            range.select();
        }
    }

    function setCaretToPos (input, pos) {
        setSelectionRange(input, pos, pos);
    }

    insertAtCursor(myValue) {
        var start=this.state.selectionStart;
        var end=this.state.selectionEnd;
        var content=this.props.content;
        console.log(start+" "+end+" "+content);
        if (start || start === 0) {
            var value = content.substring(0, start)
                + myValue
                + content.substring(end, content.length);
            return value;
        } else {
            content += myValue;
            return content;
        }
    }*/

    handleKeyDown(e){
        console.log(e.key);
        if (e.key === 'Tab') {
            e.preventDefault();
            var textarea=document.getElementById("textarea");
            var start = textarea.selectionStart;
            var end = textarea.selectionEnd;



            // set textarea value to: text before caret + tab + text after caret
            textarea.value = textarea.value.substring(0, start) +
                "\t" + textarea.value.substring(end);

            // put caret at right position again
            this.selectionStart =
                this.selectionEnd = start + 1;
        }
    }


    handleChange(event){
        console.log("handleChange");
        this.props.onChange(event.target.value);
    }

    render(){
        return (
            <form id="form-container">
                <textarea id="textarea" onChange={this.handleChange}
                          onKeyDown={this.handleKeyDown}
                         /* onKeyDown={this.handleKeyDown}*/ value={this.props.content}
                          /*onClick={this.handleClick}*/>

                </textarea>
            </form>
        )
    }

}

export default Editor;