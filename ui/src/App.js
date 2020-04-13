import React from 'react';
import logo from './logo.svg';
import './App.css';
import PingComponent from './components/Ping';

function App() {
  return (
    <div className="App">
      <header className="App-header">
        <img src={logo} className="App-logo" alt="logo" />
        <p>
          <b>Coming soon:</b> A React + Go web app deployed to Digital Ocean using Docker
        </p>
        <PingComponent />
      </header>
    </div>
  );
}

export default App;
