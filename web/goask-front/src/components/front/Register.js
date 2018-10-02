import React, { Component } from 'react';
import frontWrapper from '../../hoc/FrontWrapper';

class Register extends Component {
    constructor (props) {
        super(props)
    }

    state = {

    }

    handleSubmit = () => {
        
    }

    render () {
        return (
            <div className="col-lg-4 offset-lg-4">
                <div className="content-wrapper">
                    <div className="card border-success">
                        <div className="card-header bg-success text-white font-weight-bold">
                            Register
                        </div>
                        <div className="card-body">
                            <form onSubmit={this.handleSubmit}>
                                <div className="form-group">
                                    <label htmlFor="username">Username</label>
                                    <input name="username" type="text" className="form-control" id="username" />
                                </div>
                                <div className="form-group">
                                    <label htmlFor="email">Email</label>
                                    <input name="email" type="email" className="form-control" id="email" />
                                </div>
                                <div className="form-group">
                                    <label htmlFor="password">Password</label>
                                    <input name="pass" type="password" className="form-control" id="password" />
                                </div>
                                <div className="form-group">
                                    <label htmlFor="confirm password">Confirm password</label>
                                    <input name="confirmPass" type="password" className="form-control" id="confirm password" />
                                </div>

                                <button type="submit" className="btn btn-primary">Sign up</button>
                            </form>
                        </div>
                    </div>
                </div>
            </div>
       );
    } 
}

export default frontWrapper(Register);