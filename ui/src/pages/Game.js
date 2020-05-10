import React,  { Component } from 'react';
import axios from 'axios';

import CardInput from '../components/CardInput';
import GameTagHeader from '../components/GameTagHeader';
import DrawCard from '../components/DrawCard';
import GameStats from '../components/GameStats';

import fishbowl from '../assets/Fishbowl3.svg';
import './Game.css';


class GamePage extends Component {

  constructor() {
      super();
      this.state = {
          ready: false,
          round: 0,
          team_1_turn: true,
          team: "",
      }
      this.saveState = this.saveState.bind(this);
      this.startGame = this.startGame.bind(this);
      this.nextTurn = this.nextTurn.bind(this);
      this.componentSwitch = this.componentSwitch.bind(this);
  }

  saveState(data) {
    let team1 = data.teams.team_1.name
    let team2 = data.teams.team_2.name
    let currentTeam = data.team_1_turn ? team1 : team2

    this.setState({
      team_1_turn: data.team_1_turn,
      ready: data.started,
      round: data.current_round,
      team: currentTeam
    })
  }

  componentDidMount() {
    const { params: { gameId } } = this.props.match;
    axios({
        method: 'get',
        url: `/v1/api/game/${gameId}`,
        timeout: 4000,    // 4 seconds timeout
      })
    .then((response) => {
      this.saveState(response.data.result[0])
    })
    .catch(function (error) {
        console.log(error);
    });
  }

  startGame() {
    const { params: { gameId } } = this.props.match;
    axios({
        method: 'put',
        url: `/v1/api/game/${gameId}/start`,
        timeout: 4000,    // 4 seconds timeout
        data: {
            started: true,
            current_round: 1
        }
      })
    .then((response) => {
      this.saveState(response.data.result[0])
    })
    .catch(function (error) {
        console.log(error);
    });
  }

  nextTurn() {
    const { params: { gameId } } = this.props.match;
    axios({
      method: 'patch',
      url: `/v1/api/game/${gameId}`,
      timeout: 4000,    // 4 seconds timeout
      data: {
          team_1_turn: !this.state.team_1_turn,
          current_round: this.state.round,
      }
    })
    .then((response) => {
        this.saveState(response.data.result[0])
    })
    .catch(function (error) {
        console.log(error);
    });
  }

  componentSwitch(gameId) {
    switch (this.state.round) {
      // Initial game setup
      case 0:
        return <CardInput gameId={gameId} done={this.startGame}/>
      // End game play after 4 rounds
      case 5:
        return <GameStats gameId={gameId}/>
      // Default game play for rounds 1-4
      default:
        return (
          <div>
            <RoundTracker round={this.state.round}/>
            <DrawCard
              gameId={gameId}
              gameState={this.state}
              nextRound={this.startGame}
              nextTurn={this.nextTurn}
            />
          </div>
        )
    }
  }

  render() {
      const { params: { gameId } } = this.props.match;
      const gameComponent = this.componentSwitch(gameId)

      return (
        <div className="Game-page">
            <GameTagHeader gameId={gameId}/>
            <h2 className="title">This fishbowl is still under development</h2>

            <div className="row">
              <div className="col-left">
                {gameComponent}
              </div>

              <div className="col-right">
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


const RoundTracker = ({ round }) => (
  <div>
    <p>Current Round: {round}</p>
  </div>
)
  
  export default GamePage;