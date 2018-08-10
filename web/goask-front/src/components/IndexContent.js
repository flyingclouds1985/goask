import React from 'react';
import Card from './Card';


class IndexContent extends React.Component {

    render() {
        return (
            <div className="col-lg-8">
                <div className="content-wrapper">
                    <header className="clearfix">
                        <p>
                        Top Questions
                        </p>
                        <div className="btn-group float-right" role="group" aria-label="toolbar">
                            <button type="button" className="btn btn-light active">Latest</button>
                            <button type="button" className="btn btn-light">Hot</button>
                            <button type="button" className="btn btn-light">Week</button>
                            <button type="button" className="btn btn-light">Month</button>
                        </div>
                    </header>
                    <Card />

                    <nav aria-label="Page navigation example">
                        <ul className="pagination justify-content-center">
                            <li className="page-item disabled">
                            <a className="page-link" href="#" tabIndex="-1">Previous</a>
                            </li>
                            <li className="page-item"><a className="page-link" href="#">1</a></li>
                            <li className="page-item"><a className="page-link" href="#">2</a></li>
                            <li className="page-item"><a className="page-link" href="#">3</a></li>
                            <li className="page-item">
                            <a className="page-link" href="#">Next</a>
                            </li>
                        </ul>
                    </nav>
                </div> 
            </div>
        );
    }

}

export default IndexContent;