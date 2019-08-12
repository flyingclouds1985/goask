import React, { Component } from 'react';
import frontWrapper from '../../hoc/FrontWrapper';
import swal from "sweetalert";

class Login extends Component {

    constructor(props) {
        super(props);
        this.state = {
            hasError: false,
            errorMessages: {}
        };
    }

    handleSubmit = (e) => {
        e.preventDefault();
        const username = e.target.username.value;
        const password = e.target.password.value;
        const loginValues = {
            "username": username,
            "password": password
        };


        fetch('http://localhost:9090/login', {
            method: 'POST',
            headers: {
                'Accept': 'application/json',
                'Content-Type': 'application/json',
            },
            body: JSON.stringify(loginValues)
        })
            .then(res => {
                if (res.status === 200) {
                    this.success();
                }
                return res.json()
            })
            .then(data => {
                console.log(data);
                if (data.errors !== undefined) {
                    this.errors(data)
                }
            })
    };

    success = () => {
        this.setState({hasError: false});
        swal('Let\'s go', 'You have been signed in successfully!', 'success');
        setTimeout(() => {
            window.location = '/';
        }, 2000)
    };

    errors = (data) => {
        this.setState({hasError: true, errorMessages: data.errors.message})
    };


    render () {
        return (
            <div className="col-lg-4 offset-lg-4">
                <div className="content-wrapper">
                    {
                        this.state.hasError
                            ? <div>
                                {Object.keys(this.state.errorMessages).map(field => {
                                    return (
                                        <div key={field} className="alert alert-danger" role="alert">
                                            {this.state.errorMessages[field]}
                                        </div>
                                    );
                                })
                                }
                            </div>
                            : ''
                    }
                    <div className="card border-success">
                        <div className="card-header bg-success text-white font-weight-bold">
                            Login
                        </div>
                        <div className="card-body">
                            <form onSubmit={this.handleSubmit}>
                                <div className="form-group">
                                    <label htmlFor="username">Username</label>
                                    <input name="username" type="text" className="form-control" id="username" />
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