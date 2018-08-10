import React from 'react';
import Sidebar from './Sidebar';
import IndexContent from './IndexContent';

class Index extends React.Component {

    componentDidMount() {
        console.log("hello");
    }

    render() {
        return(
            <div className="row">
                <IndexContent />            
                <Sidebar />
            </div>
        );
    }

}

export default Index;