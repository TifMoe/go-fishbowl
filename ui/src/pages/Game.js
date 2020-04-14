import React from 'react';
import fishbowl from '../assets/Fishbowl4.svg';
import './Game.css';

const GamePage = ({ match }) => {
    const { params: { gameId } } = match;

    return (
        <header className="Game-header">
          <img src={fishbowl} className="Game-logo" alt="logo" />
          <h1>{gameId}</h1>
          <h2>Sorry, Fishbowl still under development</h2>
        </header>
    );
  }
  
  export default GamePage;