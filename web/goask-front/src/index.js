import React, {Component} from 'react';
import ReactDOM from 'react-dom';
import registerServiceWorker from './registerServiceWorker';
import 'bootstrap';
import 'bootstrap/dist/css/bootstrap.min.css';
import './assets/css/style.css';
import Index from './components/front/Index';
import {BrowserRouter, Route, Switch} from 'react-router-dom';
import QuestionRoute from './components/front/QuestionRoute';
import NotFound from './NotFound';

class App extends Component {
    render() {
        return (
            <BrowserRouter>
                <Switch>
                    <Route exact path="/" component={Index} />
                    <Route path="/questions/" component={QuestionRoute} />
                    <Route path="/404" component={NotFound} />  
                    <Route component={NotFound} />
                </Switch>
            </BrowserRouter>
        );
    }
}




ReactDOM.render(<App />, document.getElementById('root'));
registerServiceWorker();
