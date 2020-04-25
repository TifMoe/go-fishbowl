import React from 'react';
import fishbowl from '../assets/Fishbowl3.svg';
// import CardInput from '../components/CardInput';
import './Game.css';

const GamePage = ({ match }) => {
    const { params: { gameId } } = match;

    return (
        <header className="Game-header">
          <img src={fishbowl} className="Game-logo" alt="logo" />
          <h1>{gameId}</h1>
          {/* <CardInput gameId={gameId}/> */}
          <h2>Sorry, your Fishbowl is still under development</h2>
        </header>
    );
  }
  
  export default GamePage;