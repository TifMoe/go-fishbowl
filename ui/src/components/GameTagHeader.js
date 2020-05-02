import React, { Component } from 'react';
import './GameTagHeader.css';

class GameTagHeader extends Component {
    render() {
        return (
            <div class="header" id="myHeader">
                Send this link to the other players: <a href={`https://fishbowl.rocks/game/${this.props.gameId}`} className="Game-link">
                    fishbowl.rocks/game/{this.props.gameId}
                  </a>
            </div>
        );
    }
}

export default GameTagHeader; 