import React from 'react';
import logo from '../logo.svg';
import '../App.css';
import Picker from '../components/NamePicker';

const HomePage = () => {
    return (
        <header className="App-header">
          <img src={logo} className="App-logo" alt="logo" />
          <h2>Coming Soon:</h2>
          <p>
            A React + Go web app for playing Fishbowl over video calls!
          </p>
          <Picker />
        </header>
    );
  }
  
  export default HomePage;