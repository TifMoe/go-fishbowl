import React from 'react';
import { useHistory } from "react-router-dom";

import './Home.css';
import NewGame from '../components/NewGame';
import TitleImage from '../components/TitleImage';

function HomePage() {
    const history = useHistory();
    return (
        <header className="Home-header">
          <h2>Game nights can carry on from a distance with</h2>
          <TitleImage />
          <div className="intro">
            <p>Play fishbowl online over video calls with friends and family</p>
            <NewGame history={history}/>
            <p className="small-text">Keep scrolling to check out the game rules or enter two cool team names above to get started!</p>
          </div>
        </header>
    );
  }
  
  export default HomePage;