import React from 'react';
import '../App.css';

const GamePage = ({ match }) => {
    const { params: { gameId } } = match;

    return (
        <header className="App-header">
          <h1>{gameId}</h1>
          <h2>Here's a new game for you! </h2>
        </header>
    );
  }
  
  export default GamePage;