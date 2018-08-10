import React from 'react';

class Card extends React.Component {

    render() {
        return (
            <div className="card">
                <div className="card-body">
                    <h5 className="card-title"><a href="/questions/{{ .Id }}"></a></h5>
                    <ul className="tags-list">
                        <li>
                            <button type="button" className="btn btn-outline-success btn-sm"></button>
                        </li>
                        
                    </ul>
                </div>
                <div className="card-footer text-muted">
                    <ul className="meta-list">
                        <li>
                            <small className="text-muted">votes: </small>
                        </li>
                        <li>
                            <small className="text-muted">view: 1423</small>
                        </li>
                        <li>
                            <small className="text-muted"><span className="custom-badge badge-success">answers: </span></small>
                        </li>
                        <li>
                            <small className="text-muted">asked 52 secs ago <a href="#">Striker</a></small>
                        </li>
                    </ul>
                </div>
            </div>
        );
    }

}

export default Card;