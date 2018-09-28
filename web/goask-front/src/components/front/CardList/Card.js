import React from 'react';
import { Link } from 'react-router-dom';

class Card extends React.Component {

    render() {
        return (
                <div className="card">
                    <div className="card-body">
                        <h5 className="card-title">
                            <Link  to={{pathname: "/questions/" + this.props.question.id, state: {question:this.props.question} }}>
                                {this.props.question.title}
                            </Link>
                        </h5>
                        <ul className="tags-list">
                            {
                                this.props.question.tags
                                ?this.props.question.tags.map( tag => {
                                    return (
                                        <li key={tag.id}>
                                            <button type="button" className="btn btn-outline-success btn-sm">{tag.name}</button>
                                        </li>
                            
                                    );
                                })
                                :""
                            }
                            
                        </ul>
                    </div>
                    <div className="card-footer text-muted">
                        <ul className="meta-list">
                            <li>
                                <small className="text-muted">votes: {this.props.question.vote}</small>
                            </li>
                            <li>
                                <small className="text-muted">view: 1423</small>
                            </li>
                            <li>
                                <small className="text-muted"><span className="custom-badge badge-success">answers: </span></small>
                            </li>
                            <li>
                                <small className="text-muted">asked 52 secs ago <a href="">Striker</a></small>
                            </li>
                        </ul>
                    </div>
                </div>
        );
    }

}

export default Card;