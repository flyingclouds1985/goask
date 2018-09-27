import React, { Component } from 'react';

class Question extends Component {
    constructor (props) {
        super(props)
    }
    render () {
        let question = this.props.location.state.question
        return (
            <div className="col-lg-8">
                <div className="content-wrapper">
                    <header>
                        <h5>
                            { question.title }
                        </h5>
                    </header>
                    <div className="card">
                        <div className="card-body">
                            <div className="vote">
                                    <a href="#" style={{color: "#586268"}}><span className="oi oi-caret-top"></span></a>
                                        <p>{question.vote}</p>
                                    <a href="#" style={{color: "#586268"}}><span className="oi oi-caret-bottom"></span></a>
                            </div>
                            <p className="card-text" style={{marginLeft: "30px"}}>
                                { question.body }
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
           

export default Question;