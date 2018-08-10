import React from 'react';

class Sidebar extends React.Component {
    render() {
        return (
            <div className="col-lg-4">
                <div className="sidebar-wrapper">
                    <div className="card">
                        <div className="card-header">Top Tags</div>
                        <div className="card-body">
                            <p className="card-text"><a href="#">Javascript</a> <span className="badge badge-light">12312229</span></p>
                            <p className="card-text"><a href="#">Golang</a> <span className="badge badge-light">12229</span></p>
                            <p className="card-text"><a href="#">Java</a> <span className="badge badge-light">32143</span></p>
                        </div>
                    </div>
                </div>
            </div>
        );
    }
}

export default Sidebar;