import React,  { Component } from 'react';
import fishbowl from '../assets/Fishbowl3.svg';

import './CardCounter.css'

class CardCounter extends Component {

    constructor() {
        super();
        this.state = {
            cardCount: 0,
        }
    }

    componentDidMount() {
        // EVENT LISTENERS //
        this.props.socket.on('cardCount', this.updateCount);
        this.props.socket.on('gameState', this.gameState);
      }
    
    componentWillUnmount() {
        this.props.socket.off('cardCount', this.updateCount);
        this.props.socket.off('gameState', this.gameState);
    }

    updateCount = (data) => {
        this.setState({ cardCount: data })
    }

    // gameState is an event listener for updates to the game state
    gameState = (data) => {
        let response = JSON.parse(data)
        this.setState({ cardCount: response.unused_cards })
    }

    // getGame is an event emitter to request game state on socket connection
    getGame = () => {
        let data = JSON.stringify({
            gameID: this.gameId
        });
        console.log('Fetching state for ', this.gameId);
        this.socket.emit('getGame', data);
      }

    render() {

        return (
            <div>
                <div className="logo">
                    <img src={fishbowl} alt="logo" />
                </div>
                <div className="counter">
                    <div className="count">
                        <b>{this.state.cardCount}</b><br/>
                        cards
                    </div>
                </div>
            </div>
        )
    }
}

export default CardCounter;