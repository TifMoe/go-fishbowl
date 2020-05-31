import React from 'react';
import { BrowserRouter, Route, Switch } from 'react-router-dom';

import './App.css';
import HomePage from './pages/Home';
import GamePage from './pages/Game';

function App() {
  return (
    <BrowserRouter>
    <div className="App">
        <Switch>
         <Route path="/" component={HomePage} exact/>
         <Route path="/game/:gameId" component={GamePage}/>
       </Switch>
    </div>
  </BrowserRouter>
  );
}

export default App;