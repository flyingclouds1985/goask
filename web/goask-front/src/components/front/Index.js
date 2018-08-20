import React from 'react';
import Sidebar from './Sidebar';
import IndexContent from './IndexContent';
import frontWrapper from '../../hoc/FrontWrapper';

class Index extends React.Component {

    render() {
        return(
            <div className="row">
                <IndexContent />            
                <Sidebar />
            </div>
        );
    }

}

export default frontWrapper(Index);