class Display extends React.Component{
    constructor(props) {
        super(props);
    }

    render(){
        return (
            <div dangerouslySetInnerHTML={{ __html: `${this.props.text}` }}>
            </div>
        )
    }
}

export default Display;