import React from 'react';
import fishbowl from '../assets/Fishbowl3.svg';
import CardInput from '../components/CardInput';
import GameTagHeader from '../components/GameTagHeader';

import './Game.css';

const GamePage = ({ match }) => {
    const { params: { gameId } } = match;

    return (
      <div className="Game-page">
          <GameTagHeader gameId={gameId}/>
            <h2>Sorry, your Fishbowl is still under development</h2>

          <div className="row">
            <div className="column">
              <img src={fishbowl} className="Game-logo" alt="logo" />
              <CardInput gameId={gameId}/>
            </div>

            <div className="column">
              <p>
                Draw Card component coming soon!
              </p>
            </div>
          </div>
        </div>
    );
  }
  
  export default GamePage;