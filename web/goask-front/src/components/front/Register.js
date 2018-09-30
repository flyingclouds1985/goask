import React, { Component } from 'react';
import frontWrapper from '../../hoc/FrontWrapper';

class Register extends Component {
   render () {
        return (
            <div className="col-lg-4 offset-lg-4">
                <div className="content-wrapper">
                    <div className="card border-success">
                        <div className="card-header bg-success text-white font-weight-bold">
                            Register
                        </div>
                        <div className="card-body">
                            <form>
                                <div className="form-group">
                                    <label htmlFor="username">Username</label>
                                    <input type="text" className="form-control" id="username" />
                                </div>
                                <div className="form-group">
                                    <label htmlFor="email">email</label>
                                    <input type="email" className="form-control" id="email" />
                                </div>
                                <div className="form-group">
                                    <label htmlFor="password">password</label>
                                    <input type="text" className="form-control" id="password" />
                                </div>
                                <div className="form-group">
                                    <label htmlFor="confirm password">confirm password</label>
                                    <input type="text" className="form-control" id="confirm password" />
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