import React from 'react';
import fishbowl from '../assets/Fishbowl3.svg';
// import CardInput from '../components/CardInput';
import './Game.css';

const GamePage = ({ match }) => {
    const { params: { gameId } } = match;

    return (
        <header className="Game-header">
          <h2>Sorry, your Fishbowl is still under development</h2>
          <img src={fishbowl} className="Game-logo" alt="logo" />
          <p>
            Send this link to your friends to join
            : <a href={`https://fishbowl.rocks/game/${gameId}`} className="Game-link">
                fishbowl.rocks/game/{gameId}
              </a>
          </p>
          {/* <CardInput gameId={gameId}/> */}
        </header>
    );
  }
  
  export default GamePage;