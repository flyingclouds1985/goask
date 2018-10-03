import React from 'react';
import CardList from './CardList/CardList';


class IndexContent extends React.Component {

    state = {
        'groupBtn': 'latest'
    }

    handleBtnClick = (e) => {
        this.setState({'groupBtn': e.target.name})        
    }

    activeBtn = (name) => {
        return this.state.groupBtn === name ? 'active' : '';
    }

    render() {
        return (
            <div className="col-lg-8">
                <div className="content-wrapper">
                    <header className="clearfix">
                        <p>
                        Top Questions
                        </p>
                        <div className="btn-group float-right" role="group" aria-label="toolbar">
                            <button onClick={this.handleBtnClick} name="latest" type="button" className={"btn btn-light" + this.activeBtn('latest')}>Latest</button>
                            <button onClick={this.handleBtnClick} name="hot" type="button" className={"btn btn-light" + this.activeBtn('hot')}>Hot</button>
                            <button onClick={this.handleBtnClick} name="week" type="button" className={"btn btn-light" + this.activeBtn('week')}>Week</button>
                            <button onClick={this.handleBtnClick} name="month" type="button" className={"btn btn-light" + this.activeBtn('month')}>Month</button>
                        </div>
                    </header>

                    <CardList />

                    <nav aria-label="Page navigation example">
                        <ul className="pagination justify-content-center">
                            <li className="page-item disabled">
                            <a className="page-link" href="" tabIndex="-1">Previous</a>
                            </li>
                            <li className="page-item"><a className="page-link" href="">1</a></li>
                            <li className="page-item"><a className="page-link" href="">2</a></li>
                            <li className="page-item"><a className="page-link" href="">3</a></li>
                            <li className="page-item">
                            <a className="page-link" href="">Next</a>
                            </li>
                        </ul>
                    </nav>
                </div> 
            </div>
        );
    }

}

export default IndexContent;