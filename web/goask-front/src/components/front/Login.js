import React, { Component } from 'react';
import frontWrapper from '../../hoc/FrontWrapper';

class Login extends Component {
    render () {
        return (
            <div className="col-lg-4 offset-lg-4">
                <div className="content-wrapper">
                    <div className="card border-success">
                        <div className="card-header bg-success text-white font-weight-bold">
                            Login
                        </div>
                        <div className="card-body">
                            <form onSubmit={this.handleSubmit}>
                                <div className="form-group">
                                    <label htmlFor="email">Email</label>
                                    <input name="email" type="email" className="form-control" id="email" />
                                </div>
                                <div className="form-group">
                                    <label htmlFor="password">Password</label>
                                    <input name="pass" type="password" className="form-control" id="password" />
                                </div>
                                <button type="submit" className="btn btn-primary">Sign in</button>
                            </form>
                        </div>
                    </div>
                </div>
            </div>
        );
    }
}

export default frontWrapper(Login);