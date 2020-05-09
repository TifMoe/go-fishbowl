import React from 'react';
import { useHistory } from "react-router-dom";
import title from '../assets/TitleImg.svg';
import './Home.css';
import NewGame from '../components/NewGame';

const HomePage = () => {
    const history = useHistory();
    return (
        <header className="Home-header">
          <img src={title} className="Home-logo" alt="logo" />
          <h2>Coming Soon:</h2>
          <p>
            A React + Go web app for playing Fishbowl over video calls!
          </p>
          <NewGame history={history}/>
        </header>
    );
  }
  
  export default HomePage;