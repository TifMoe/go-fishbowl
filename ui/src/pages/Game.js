import React,  { Component } from 'react';
import Socket from '../Socket';

import CardInput from '../components/CardInput';
import GameTagHeader from '../components/GameTagHeader';
import DrawCard from '../components/DrawCard';
import GameStats from '../components/GameStats';
import ScoreKeeper from '../components/ScoreKeeper';

import fishbowl from '../assets/Fishbowl3.svg';
import rules from './../gameRules.json';
import './Game.css';


class GamePage extends Component {

  constructor() {
      super();
      this.socket = new Socket();
      this.state = {
          connected: false,
          ready: false,
          round: 0,
          team_1_turn: true,
          unused_cards: 0,
          team1: {
            name: "Team 1",
            round_1_pts: 0,
            round_2_pts: 0,
            round_3_pts: 0,
            round_4_pts: 0,
          },
          team2: {
            name: "Team 2",
            round_1_pts: 0,
            round_2_pts: 0,
            round_3_pts: 0,
            round_4_pts: 0,
          }
      }
  }

  saveState = (data) => {
    this.setState({
      team1: data.teams.team_1,
      team2: data.teams.team_2,
      team_1_turn: data.team_1_turn,
      ready: data.started,
      round: data.current_round,
      unused_cards: data.unused_cards,
    })
  }

  componentDidMount = () => {
    console.log("Setting up component")
    const { params: { gameId } } = this.props.match
    this.gameId = gameId

    // handle connect and disconnect events.
    this.socket.on('connect', this.onConnect);
    this.socket.on('disconnect', this.onDisconnect); 

    /* EVENT LISTENERS */
    this.socket.on('cardCount', this.cardCount);
    this.socket.on('gameState', this.gameState)
  }

  // onConnect sets the state to true indicating the socket has connected successfully
  onConnect = () => {
      console.log("Connected!")
      this.setState({connected: true});
      this.getGame()
  }

  // onDisconnect sets the state to false indicating the socket has been disconnected
  onDisconnect = () => {
      this.setState({connected: false});
  }

  // cardCount is an event listener for updates to the card count
  cardCount = (data) => {
      this.setState({unused_cards: data})
  }

  // gameState is an event listener for updates to the game state
  gameState = (data) => {
      this.saveState(JSON.parse(data))
  }


  // EVENT EMITTER //
  // getGame is an event emitter to request game state on socket connection
  getGame = () => {
    let data = JSON.stringify({
        gameID: this.gameId
    });
    console.log('Fetching state for ', this.gameId);
    this.socket.emit('getGame', data);
  }

  componentSwitch = (socket) => {
    var title;
    var leftComponent;
    var rightComponent;

    switch (this.state.round) {
      // Initial game setup
      case 0:
        title = <h2>Enter nouns below to get started!</h2>;
        leftComponent = <CardInput gameId={this.gameId} socket={socket}/>;
        break
      case 5: // Force end of game
        title = <h2>Congratulations!!</h2>;
        leftComponent = <GameStats gameId={this.gameId} gameState={this.state}/>;
        break
      // Transition to Game Stats page at the end of round 4
      case 4:
        if (this.state.unused_cards === 0) { // Natural end of game
          title = <h2>Congratulations!!</h2>;
          leftComponent = <GameStats gameId={this.gameId} gameState={this.state}/>;
          break
        }
        // fallthrough
      default:
          title = <RoundTracker round={this.state.round} team1={this.state.team1} team2={this.state.team2}/>;
          leftComponent = <div>
              <DrawCard
                gameId={this.gameId}
                socket={socket}
                gameState={this.state}
              />
            </div>;
          rightComponent = <ScoreKeeper team1={this.state.team1} team2={this.state.team2}/>;
          break
    }
    return {
      "title": title,
      "leftComponent": leftComponent,
      "rightComponent": rightComponent,
    }
  }

  render() {
      const element  = this.componentSwitch(this.socket)

      return (
        <div className="Game-page">
          <GameTagHeader gameId={this.gameId}/>

          <div className="row">
            <div className="title">{element.title}</div>
          </div>

            <div className="row">
              <div className="col-left">
                {element.leftComponent}
              </div>

              <div className="col-right">
                  {element.rightComponent}
                  <div className="logo">
                    <img src={fishbowl} className="bowl" alt="logo" />
                  </div>
                  {/* TODO: Add counter component here */}
              </div>
            </div>

          </div>
      );
    }
  }


function RoundTracker({ round }) {
  return (
    <div className="round-name">
        <h2><b>{rules.rounds[round-1].name}</b></h2>
    </div>
  )
}
  
  export default GamePage;