import React, { Component } from 'react';
import Timer from './Timer';

import axios from 'axios';

import './DrawCard.css';
import rules from './../assets/roundRules.json';

class DrawCard extends Component {

    constructor() {
        super();
        this.state = {
            id: "",
            card: "",
            team1: "Team 1",
            team2: "Team 2",
            showCard: false,
            showNextRound: false,
            showSkip: true,
        }
        this.drawCard = this.drawCard.bind(this);
        this.markDone = this.markDone.bind(this);
        this.endTurn = this.endTurn.bind(this);
        this.endRound = this.endRound.bind(this);
        this.nextRoundTransition = this.nextRoundTransition.bind(this);
    }

    componentDidMount() {
        axios({
            method: 'get',
            url: `/v1/api/game/${this.props.gameId}`,
            timeout: 4000,    // 4 seconds timeout
          })
        .then((response) => {
          // On page load find current team in play
          this.setState({
              team1: response.data.result[0].teams.team_1.name,
              team2: response.data.result[0].teams.team_2.name
            })
        })
        .catch(function (error) {
            console.log(error);
        });
      }

    endTurn() {
        this.setState({showCard: false});
        this.props.nextTurn();
    }

    endRound() {
        this.setState({showNextRound: false})
        this.props.nextRound();
    }

    markDone() {
        // Mark current card as done
        axios({
            method: 'patch',
            url: `/v1/api/game/${this.props.gameId}/card/${this.state.id}/used`,
            timeout: 4000,    // 4 seconds timeout
          })
        .then((response) => {
            // Fetch new card
            this.drawCard();
        })
        .catch(function (error) {
            console.log(error);
        });
    }

    drawCard() {
        this.setState({showNextRound: false})

        axios({
            method: 'get',
            url: `/v1/api/game/${this.props.gameId}/card/random`,
            timeout: 4000,    // 4 seconds timeout
          })
        .then((response) => {
            const cards = response.data.result[0].cards;
            const cardCount = response.data.result[0].unused_cards

            if (cards && cards.length) {
                if (cardCount == 1) {
                    this.setState({ showSkip: false })
                }

                // Show new card
                this.setState(() => {
                    return { card: cards[0].value,
                        id: cards[0].id,
                        showCard: true,
                    }
                })
            } else {
                // If no cards, end turn and show next round button
                this.endTurn();
                this.setState({showNextRound: true})
            }
        })
        .catch(function (error) {
            console.log(error);
        });
    }

    nextRoundTransition() {
        const team = this.props.gameState.team_1_turn ? this.state.team1 : this.state.team2;
        const color = this.props.gameState.team_1_turn ?  "rgb(242, 85, 119, .7)":  "rgb(46, 221, 204, .7)";
        const round = this.props.gameState.round

        if (this.state.showNextRound) {
            if (round >= 4) {
                // End game play after 4 rounds
                return <p>Game over</p>
            }
            // Show next round rules in transition to next round
            return (
                <NextRoundRules
                    round={round+1}
                    rules={rules.rounds}
                />
            )
        }
         // Default view showing active team for inactive players
         return (
            <TeamUp
                team={team}
                color={color}
            />
         )
    }

    render() {
        return (
        <div className="draw-card">

            { this.state.showCard ?
                // Player actively drawing new cards
                <div>
                    <div className="actions">
                        <Timer timesUpHandler={this.endTurn}/>
                    </div>
                    <Card
                        card={this.state.card}
                        showSkip={this.state.showSkip}
                        doneHandler={this.markDone}
                        drawHandler={this.drawCard}
                    />
                </div>:

                // Card values hidden
                <div>
                    <div className="actions">
                        <button onClick={this.drawCard}>Start Turn</button>
                        <button onClick={this.endTurn}>End Turn</button>
                    </div>
                    {this.nextRoundTransition()}
                </div>
            }
            <NextRound
                active={this.state.showNextRound}
                nextHandler={this.endRound}/>
        </div>
        );
    }
}

const Card = ({ card, showSkip, doneHandler, drawHandler }) => (
    <div className="card">
        <div className="card-value">
            <p>{card}</p>
        </div>
        <div className="actions">
            <button onClick={doneHandler} className="done">Got it!</button>
            <button onClick={drawHandler}  disabled={!showSkip} className="skip">Skip</button>
        </div>
    </div>
)

const NextRoundRules = ({ round, rules }) => (
    <div className="card" style={{backgroundColor: "#5F6167", color: "white"}}>
        <div className="card-value">
            <h3>Round {round}: {rules[round-1].name}</h3>
            <p>{rules[round-1].rules}</p>
        </div>
    </div>
)

const TeamUp = ({ team, color }) => (
    <div className="card" style={{backgroundColor: color, color: "white"}}>
        <div className="card-value">
            <p>{team}'s Turn!</p>
        </div>
    </div>
)

const NextRound = ({ active, nextHandler }) => (
    <div>
        <button
            onClick={nextHandler}
            className="next-round"
            disabled={!active}
        >Start Round</button>
    </div>
)

export default DrawCard; 