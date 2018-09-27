import React from 'react';
import { Switch, Route } from 'react-router-dom';
import Question from './Question';

const QuestionRoute = () => (
        <Switch>
            <Route path='/questions/:id' component={Question} />
        </Switch>
);

export default QuestionRoute;