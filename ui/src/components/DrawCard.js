import React, { Component } from 'react';
import axios from 'axios';

import './DrawCard.css';

class DrawCard extends Component {

    constructor() {
        super();
        this.state = {
            id: "",
            card: "",
            showCard: false,
            teamA: true,
        }
        this.drawCard = this.drawCard.bind(this);
        this.endTurn = this.endTurn.bind(this);
        this.markDone = this.markDone.bind(this);
    }

    endTurn() {
        this.setState({showCard: false})
        this.setState({teamA: !this.state.teamA})
    }

    markDone() {
        // Mark current card as done
        axios({
            method: 'patch',
            url: `/v1/api/game/${this.props.gameId}/card/${this.state.id}/used`,
            timeout: 4000,    // 4 seconds timeout
          })
        .then((response) => {
            console.log(response);
            // Fetch new card
            this.drawCard();
        })
        .catch(function (error) {
            console.log(error);
        });
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
                // Show new card
                this.setState(() => {
                    return { card: cards[0].value,
                        id: cards[0].id,
                        showCard: true,
                    }
                })
            } else {
                // If no cards, end turn
                this.endTurn();
            }
        })
        .catch(function (error) {
            console.log(error);
        });
    }

    render() {
        const team = this.state.teamA ? "Pink": "Blue";
        const color = this.state.teamA ?  "rgb(242, 85, 119, .7)":  "rgb(46, 221, 204, .7)";
        return (
        <div>
            <button onClick={this.drawCard}>Start Turn</button>
            <button onClick={this.endTurn}>End Turn</button>
            { this.state.showCard ?
                <Card
                    card={this.state.card}
                    doneHandler={this.markDone}
                    drawHandler={this.drawCard}
                /> :
                <PlaceHolder
                    team={team}
                    color={color}
                />
            }
        </div>
        );
    }
}

const Card = ({ card, doneHandler, drawHandler }) => (
    <div className="card">
        <div className="card-value">
            <p>{card}</p>
        </div>
        <div className="actions">
            <button onClick={doneHandler} className="done">Got it!</button>
            <button onClick={drawHandler} className="skip">Skip</button>
        </div>
    </div>
)

const PlaceHolder = ({ team, color }) => (
    <div className="card" style={{backgroundColor: color, color: "white"}}>
        <div className="card-value">
            <p>{team} Team's Turn!</p>
        </div>
    </div>
)

export default DrawCard; 