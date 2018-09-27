import React, { Component } from 'react';
import frontWrapper from '../../hoc/FrontWrapper';

class Question extends Component {
    constructor (props) {
        super(props)
        this.id = props.match.params.id
    }

    state = {
        question: {}
    }

    componentDidMount = () => {
        if (this.props.location.state === undefined) {
            fetch('http://localhost:9090/questions/' + this.id)
            .then(response => {
                if (response.status === 404)
                    throw new Error("NotFound")
                return response.json()
            })
            .then(data => {this.setState({question: data})})
            .catch(err => {
                if (err.message === 'NotFound') {
                    this.props.history.push('/404')
                }
            })
        } else {
            this.setState({question: this.props.location.state.question})
        }
    }


    render () {
        return (
            <div className="col-lg-8">
                <div className="content-wrapper">
                    <header>
                        <h5>
                            { this.state.question.title }
                        </h5>
                    </header>
                    <div className="card">
                        <div className="card-body">
                            <div className="vote">
                                    <a href="#" style={{color: "#586268"}}><span className="oi oi-caret-top"></span></a>
                                        <p>{this.state.question.vote}</p>
                                    <a href="#" style={{color: "#586268"}}><span className="oi oi-caret-bottom"></span></a>
                            </div>
                            <p className="card-text" style={{marginLeft: "30px"}}>
                                { this.state.question.body }
                            </p>
                            <ul className="tags-list" style={{marginLeft: "30px", marginTop: "20px"}}>
                                {/* {{ range .Tags }}
                                <li>
                                        <button type="button" className="btn btn-outline-success btn-sm">{{ .Name }}</button>
                                </li>
                                {{ end }} */}
                                <li className="float-right">
                                        <button type="button" className="btn btn-primary btn-sm">edit</button>
                                </li>
                            </ul>
                        </div>
                        <div className="card-footer text-muted">
                            <ul className="comments">
                                {/* {{ range .Comments }}
                                    <li>
                                        <small className="comment-text">{{ .Body }}<span className="oi oi-person"></span> <a href="#">apokryfos</a> <span className="text-muted">Jul 3 at 16:03</span><a href="#" style="margin-left: 5px;">edit</a></small>
                                    </li>
                                {{ end }} */}
                            </ul>
                            <a href="#">add comment</a>
                        </div>
                    </div>
                </div> 
            </div>
        );
    }
}
           

export default frontWrapper(Question);