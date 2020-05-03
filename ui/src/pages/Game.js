import React from 'react';
import CardInput from '../components/CardInput';
import GameTagHeader from '../components/GameTagHeader';
import DrawCard from '../components/DrawCard';

import fishbowl from '../assets/Fishbowl3.svg';
import './Game.css';

const GamePage = ({ match }) => {
    const { params: { gameId } } = match;

    return (
      <div className="Game-page">
          <GameTagHeader gameId={gameId}/>

          <div className="row">
          <h2 className="title">Sorry, your Fishbowl is still under development</h2>

            <div className="col-left">
              <div className="logo">
                <img src={fishbowl} className="bowl" alt="logo" />
              </div>
              <CardInput gameId={gameId}/>
            </div>

            <div className="col-right">
              <DrawCard gameId={gameId}/>
            </div>
          </div>
        </div>
    );
  }
  
  export default GamePage;