import React from 'react';
import Aux from './AuxiliaryWrapper';
import Header from '../components/front/Header';
import Footer from '../components/front/Footer';

const frontWrapper = (WrappedComponent) => {
    return (props) => {
        return (
            <Aux>
                <Header {...props} />
                
                <div className="container-fluid">
                    <WrappedComponent {...props} />
                </div>

                <Footer {...props} />
            </Aux>
        );
    }
}

export default frontWrapper;