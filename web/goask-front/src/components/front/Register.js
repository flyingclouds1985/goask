import React, { Component } from 'react';
import frontWrapper from '../../hoc/FrontWrapper';
import swal from 'sweetalert';

class Register extends Component {
    state = {
        hasError: false,
        errorMessages: {},
    };

    handleSubmit = (e) => {
        e.preventDefault();
        const t = e.target;
        const username = t.username.value;
        const email = t.email.value;
        const password = t.password.value;
        const confirmPassword = t.confirmPassword.value;
        const user = {username, email, password, confirmPassword};

        fetch('http://localhost:9090/users/', {
            method: 'POST',
            headers: {
                'Accept': 'application/json',
                'Content-Type': 'application/json',
            },
            body: JSON.stringify(user)
        })
        .then(res => {
            if (res.status === 200) {
                this.success();
            } else {
                return res.json()
            }
            return res.json()
        })
        .then(data => {
            if (data.errors !== undefined) {
                this.errors(data)
            }            
        }).catch(err => {
            console.log("err: ", err)
        })
    };

    success = () => {
        this.setState({hasError: false});
        swal('Let\'s go', 'You have been signed up successfully!', 'success');
        setTimeout(() => {
                window.location = '/login';
        }, 2000)
    };

    errors = (data) => {
        if(data.errors.status === 500) {
            this.setState({hasError: true, errorMessages: data.errors.error})
        }
        if(data.errors.status === 400) {
            this.setState({hasError: true, errorMessages: data.errors.message})
        }

    };

    render () {
        return (
            <div className="col-lg-6 offset-lg-3">
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
                                    <input name="password" type="password" className="form-control" id="password" />
                                </div>
                                <div className="form-group">
                                    <label htmlFor="confirm password">Confirm password</label>
                                    <input name="confirmPassword" type="password" className="form-control" id="confirm password" />
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