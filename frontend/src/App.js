import React from 'react';
import Main from './Pages/Main'
import Game from './Pages/Game'

import {
  BrowserRouter as Router,
  Switch,
  Route,
  //Redirect,
  //Link
} from "react-router-dom";

class App extends React.Component {
  render() {

    return (
      <div>
        <Routes/>
      </div>
    );
  }
}

class Routes extends React.Component {
  constructor(props) {
    super(props);
    this.state = {

    };
  }

  render() {

    return (
      <Router>
        <Switch>
          <Route exact path="/" >
            <Main />
          </Route>
          <Route exact path="/Game" >
            <Game />
          </Route>
        </Switch>
      </Router>
    );
  }
}

export default App;