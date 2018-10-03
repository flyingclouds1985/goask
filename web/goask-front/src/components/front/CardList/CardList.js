import React, {Component} from 'react';
import Card from './Card';

class CardList extends Component {
    state = {
        questions: [],
    }

    componentDidMount() {
        fetch('http://localhost:9090/questions/')
            .then(response => response.json())
            .then(data => this.setState({questions: data}));
    }

    render() {
        return (
            this.state.questions !== null
            ? this.state.questions.map( data => {
                return <Card key={data.id}  question={data} />
            })
            : <div className="alert alert-warning" role="alert">
                There is no question right now!
            </div>
        );
    }
}

export default CardList;