import React, { Component } from 'react';
import Timer from './Timer';

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
        }
        this.drawCard = this.drawCard.bind(this);
        this.markDone = this.markDone.bind(this);
        this.endTurn = this.endTurn.bind(this);
        this.endRound = this.endRound.bind(this);
    }

    componentDidMount() {
        if (this.props.gameState.unused_cards === 0) {
            this.setState({ showNextRound: true })
        }
        // EVENT LISTENERS //
        this.props.socket.on('randomCard', this.randomCard);
        this.props.socket.on('gameState', this.listenRoundChange);
      }

    componentWillUnmount() {
        this.props.socket.off('randomCard', this.randomCard);
        this.props.socket.off('gameState', this.listenRoundChange);
    }

    // Event listener to show new Round Rules when turn changes
    listenRoundChange = (data) => {
        let response = JSON.parse(data)
        const count = response.unused_cards;

        if (count === 0) {
            this.setState({showNextRound: true})
        } else {
            this.setState({showNextRound: false})
        }

    }

    // Event listener for new random card draw
    randomCard = (data) => {
        let response = JSON.parse(data)

        const cards = response.cards;
        const cardCount = response.unused_cards

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
    }

    drawCard() {
        this.setState({showNextRound: false})
        this.getRandomCard()
    }

    endTurn() {
        this.setState({showCard: false});
        this.nextTurn();
    }

    endRound() {
        this.setState({showNextRound: false})
        this.startNextRound();
    }

    // EVENT EMMITTERS //
    // startNextRound is an event emitter that tells the backend to start a new round
    startNextRound = () => {
        let data = JSON.stringify({
            gameID: this.props.gameId
        });
        console.log('Starting round ', this.props.gameId);
        this.props.socket.emit('startRound', data);
    }

    // getRandomCard is an event emitter to request a random card from the backend
    getRandomCard = () => {
        let data = JSON.stringify({
            gameID: this.props.gameId
        });
        console.log('Sequesting new random card...');
        this.props.socket.emit('getRandomCard', data);
    }

    // markDone is an event emitter to mark a card as used in the current round
    markDone() {
        // Mark current card as done
        let data = JSON.stringify({
            gameID: this.props.gameId,
            cardID: this.state.id
        });
        console.log('Marking as used: ', this.state.id);
        this.props.socket.emit('usedCard', data);

        // Fetch new card
        this.drawCard();
    }

    // nextTurn is event emitter to tell backend about an update to the game
    nextTurn = () => {
        let data = JSON.stringify({
            gameID: this.props.gameId,
            team_1_turn: !this.props.gameState.team_1_turn,
            current_round: this.props.gameState.round
        });
        console.log('Updating for next turn: Game ', this.props.gameId);
        this.props.socket.emit('updateGame', data);
    }

    // endGame is event emitter to tell backend to force end the game
    endGame = () => {
        let data = JSON.stringify({
            gameID: this.props.gameId,
            current_round: 5
        });
        console.log('Force ending game ', this.props.gameId);
        this.props.socket.emit('updateGame', data);
    }

    render() {
        const team = this.props.gameState.team_1_turn ? this.props.gameState.team1.name : this.props.gameState.team2.name;
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
                    endHandler={this.endGame}/>
            </div>
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