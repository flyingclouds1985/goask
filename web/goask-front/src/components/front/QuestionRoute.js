import React from 'react';
import { Switch, Route } from 'react-router-dom';
import Question from './Question';
import NotFound from '../../NotFound';

const QuestionRoute = () => (
        <Switch>
            <Route path='/questions/:id' component={Question} />
            <Route component={NotFound} />
        </Switch>
);

export default QuestionRoute;