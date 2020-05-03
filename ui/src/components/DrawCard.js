import React, { Component } from 'react';
import axios from 'axios';

import './DrawCard.css';

class DrawCard extends Component {

    constructor() {
        super();
        this.state = {
            id: "",
            card: "nothing yet...",
            showCard: false,
        }
        this.drawCard = this.drawCard.bind(this);
        this.endTurn = this.endTurn.bind(this);
    }

    endTurn() {
        this.setState({showCard: false})
    }

    drawCard() {
        axios({
            method: 'get',
            url: `/v1/api/game/${this.props.gameId}/card/random`,
            timeout: 4000,    // 4 seconds timeout
          })
        .then((response) => {
            const cards = response.data.result[0].cards;
            if (cards && cards.length) {
                this.setState(() => {
                    return { card: cards[0].value,
                        id: cards[0].id,
                        showCard: true,
                    }
                })
            }
        })
        .catch(function (error) {
            console.log(error);
        });
    }

    render() {
        return (
        <div>
            <button onClick={this.drawCard}>Draw Card</button>
            <button onClick={this.endTurn}>End Turn</button>
            { this.state.showCard ? <Card card={this.state.card}/> : null }
        </div>
        );
    }
}

class Card extends Component {
    render() {
        return (
            <div className="card">
                <div className="card-value">
                    <p>{this.props.card}</p>
                </div>
                <div className="actions">
                    <button className="done">Got it!</button>
                    <button className="skip">Skip</button>
                </div>
            </div>
        )
    }
}

export default DrawCard; 