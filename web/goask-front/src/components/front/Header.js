import React from 'react';
import { Link } from 'react-router-dom';

const header = (props) => {
    const { pathname } = props.location
    const home = pathname === '/' ? 'active' : '';
    const register = pathname === '/register' ? 'active' : '';
    const login = pathname === '/login' ? 'active' : '';

    return (
        <nav className="navbar navbar-expand-lg navbar-light bg-light">
            <button className="navbar-toggler" type="button" data-toggle="collapse" data-target="#navbarTogglerDemo01" aria-controls="navbarTogglerDemo01" aria-expanded="false" aria-label="Toggle navigation">
                <span className="navbar-toggler-icon"></span>
            </button>
            <div className="collapse navbar-collapse" id="navbarTogglerDemo01">
                <a className="navbar-brand" href="/">GOASK</a>
                <ul className="navbar-nav mr-auto mt-2 mt-lg-0">
                <li className={"nav-item " + home}>
                    <a className="nav-link" href="/">Home</a>
                </li>
                <li className={"nav-item " + register}>
                    <Link className="nav-link" to='/register'>Register</Link>
                </li>
                <li className={"nav-item " + login}>
                    <Link className="nav-link" to='/login'>Login</Link>
                </li>
                </ul>
                <form className="form-inline my-2 my-lg-0">
                    <input className="form-control mr-sm-2" type="search" placeholder="Search" aria-label="Search" />
                    <button className="btn btn-outline-success my-2 my-sm-0" type="submit">Search</button>
                </form>
            </div>
        </nav>
    );
}

export default header;