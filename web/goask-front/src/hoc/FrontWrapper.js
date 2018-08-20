import React from 'react';
import Aux from './AuxiliaryWrapper';
import Header from '../components/common/Header';
import Footer from '../components/common/Footer';

const frontWrapper = (WrappedComponent) => {
    return (props) => {
        return (
            <Aux>
                <Header />
                
                <div className="container-fluid">
                    <WrappedComponent {...props} />
                </div>

                <Footer />
            </Aux>
        );
    }
}

export default frontWrapper;