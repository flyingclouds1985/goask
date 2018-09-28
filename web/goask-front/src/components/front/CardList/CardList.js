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
        return this.state.questions.map( data => {
            return <Card key={data.id}  question={data} />
        })
    }
}

export default CardList;