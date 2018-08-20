import React, {Component} from 'react';
import ReactDOM from 'react-dom';
import registerServiceWorker from './registerServiceWorker';
import 'bootstrap';
import 'bootstrap/dist/css/bootstrap.min.css';
import './assets/css/style.css';
import Index from './components/front/Index';
import {BrowserRouter, Route, Switch} from 'react-router-dom';

class App extends Component {
    render() {
        return (
            <BrowserRouter>
                <Switch>
                    <Route exact path="/" component={Index} />
                    <Route component={NotFound} />
                </Switch>
            </BrowserRouter>
        );
    }
}

const NotFound = () => {
    return(
        <h1>Not Found</h1>
    );
}


ReactDOM.render(<App />, document.getElementById('root'));
registerServiceWorker();
