import React, { Component } from 'react';
import Timer from './Timer';

import axios from 'axios';

import './DrawCard.css';
import rules from './../gameRules.json';

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
            cardCount: 0,
        }
        this.drawCard = this.drawCard.bind(this);
        this.markDone = this.markDone.bind(this);
        this.endTurn = this.endTurn.bind(this);
        this.endRound = this.endRound.bind(this);
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
              team2: response.data.result[0].teams.team_2.name,
              cardCount: response.data.result[0].unused_cards
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

    getCardCount() {
        axios({
            method: 'get',
            url: '/v1/api/game/' + this.props.gameId,
            timeout: 1000,
        })
        .then((response) => {
            this.setState({cardCount: response.data.result[0].unused_cards})
        })
        .catch(error => {
            console.log(error);
        });
    }

    markDone() {
        // Mark current card as done
        console.log(this.state.id)
        axios({
            method: 'patch',
            url: `/v1/api/game/${this.props.gameId}/card/${this.state.id}/used`,
            timeout: 4000,    // 4 seconds timeout
          })
        .then((response) => {
            // Fetch new card
            this.getCardCount()
            this.drawCard();
            this.props.updateState(response.data.result[0])
        })
        .catch(function (error) {
            console.log(error);
        });
    }

    drawCard() {
        this.setState({showNextRound: false})

        axios({
            method: 'put',
            url: `/v1/api/game/${this.props.gameId}/card/random`,
            timeout: 4000,    // 4 seconds timeout
          })
        .then((response) => {
            const cards = response.data.result[0].cards;
            const cardCount = response.data.result[0].unused_cards
            this.setState({cardCount: cardCount})

            if (cards && cards.length) {
                if (cardCount === 1) {
                    this.setState({ showSkip: false })
                } else {
                    this.setState({ showSkip: true })
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

    render() {
        const team = this.props.gameState.team_1_turn ? this.state.team1 : this.state.team2;
        const color = this.props.gameState.team_1_turn ?  "rgb(242, 85, 119, .7)":  "rgb(46, 221, 204, .7)";
        const round = this.props.gameState.round

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

                    { this.state.showNextRound ?
                        <NextRoundRules
                            round={round+1}
                            rules={rules.rounds}
                        /> :
                        // Default view showing which team is active
                        <div>
                             <div className="actions">
                                <button onClick={this.drawCard}>Start Turn</button>
                                <button onClick={this.endTurn}>End Turn</button>
                            </div>
                            <TeamUp
                                team={team}
                                color={color}
                            />
                        </div>
                    }
                </div>
            }
            <div className="actions">
                <NextRound
                    active={this.state.showNextRound}
                    nextHandler={this.endRound}/>
                <EndGame
                    active={this.state.showNextRound}
                    endHandler={this.props.endGame}/>
            </div>
        </div>
        );
    }
}

const StatusIndicator = ({unusedCards}) => (
    <div className="status-indicator">
        <p>Unused Cards: {unusedCards}</p>
    </div>
)


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
    <div>
        { round <= 4 ?
            <div className="card" style={{backgroundColor: "#5F6167", color: "white"}}>  
                    <div className="card-value">
                        <h3>Round {round}: {rules[round-1].name}</h3>
                        <p>{rules[round-1].rules}</p>
                    </div>
            </div> :
            <div></div>
        }
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
        <button
            onClick={nextHandler}
            className="next-round"
            disabled={!active}
        >Next Round</button>
)

const EndGame = ({ active, endHandler }) => (
        <button
            onClick={endHandler}
            className="next-round"
            disabled={!active}
        >End Game</button>
)

export default DrawCard; 