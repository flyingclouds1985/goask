import React, { Component } from 'react';
import { Route } from 'react-router-dom';
import frontWrapper from '../../hoc/FrontWrapper';
import NotFound from '../../NotFound';

class Question extends Component {
    constructor (props) {
        super(props)
        this.id = props.match.params.id
    }

    state = {
        question: {},
        isNotFound: false,
        hasTag: false,
        hasComment: false,
    }

    checkQuestionProperty = (question) => {
        this.setState({question})
        if (question.tags !== null) 
            this.setState({hasTag: true})
        if (question.comments !== null)
            this.setState({hasComment: true})
    }

    componentDidMount = () => {
        if (this.props.location.state === undefined) {
            fetch('http://localhost:9090/questions/' + this.id)
            .then(res => {
                if (res.status === 404)
                    throw new Error("NotFound")
                return res.json()
            })
            .then(data => this.checkQuestionProperty(data))
            .catch(err => {
                if (err.message === 'NotFound')
                    this.setState({isNotFound: true})
            })
        } else {
            this.checkQuestionProperty(this.props.location.state.question)
        }
    }


    render () {
        return (
            this.state.isNotFound 
            ? <Route component={NotFound} />
            : <div className="col-lg-8">
                <div className="content-wrapper">
                    <header>
                        <h5>
                            { this.state.question.title }
                        </h5>
                    </header>
                    <div className="card">
                        <div className="card-body">
                            <div className="vote">
                                    <a href="" style={{color: "#586268"}}><span className="oi oi-caret-top"></span></a>
                                        <p>{this.state.question.vote}</p>
                                    <a href="" style={{color: "#586268"}}><span className="oi oi-caret-bottom"></span></a>
                            </div>
                            <p className="card-text" style={{marginLeft: "30px"}}>
                                { this.state.question.body }
                            </p>
                            <ul className="tags-list" style={{marginLeft: "30px", marginTop: "20px"}}>
                                {
                                    this.state.hasTag
                                    ? this.state.question.tags.map(tag => {
                                        return (
                                            <li key={ tag.id }>
                                                <button type="button" className="btn btn-outline-success btn-sm">{ tag.name }</button>
                                            </li>
                                        );
                                    })
                                    : ''
                                }
                                <li className="float-right">
                                        <button type="button" className="btn btn-primary btn-sm">edit</button>
                                </li>
                            </ul>
                        </div>
                        <div className="card-footer text-muted">
                            <ul className="comments">
                                {
                                    this.state.hasComment
                                    ? this.state.question.comments.map(comment => {
                                        return (
                                            <li key={ comment.id }>
                                                <small className="comment-text">{ comment.body }<span className="oi oi-person"></span> <a href="">apokryfos</a> <span className="text-muted">Jul 3 at 16:03</span><a href="" style={{marginLeft: 5}}>edit</a></small>
                                            </li>
                                        );
                                    })
                                    : ''
                                }
                            </ul>
                            <a href="">add comment</a>
                        </div>
                    </div>
                </div> 
            </div>
        );
    }
}
           

export default frontWrapper(Question);