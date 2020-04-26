import React from 'react';
import title from '../assets/TitleImg.svg';
import './Home.css';
import Picker from '../components/NamePicker';

const HomePage = () => {
    return (
        <header className="Home-header">
          <img src={title} className="Home-logo" alt="logo" />
          <h2>Coming Soon:</h2>
          <p>
            A React + Go web app for playing Fishbowl over video calls!
          </p>
          <Picker />
        </header>
    );
  }
  
  export default HomePage;