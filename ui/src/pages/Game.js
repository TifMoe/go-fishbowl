import React,  { Component } from 'react';
import axios from 'axios';

import CardInput from '../components/CardInput';
import GameTagHeader from '../components/GameTagHeader';
import DrawCard from '../components/DrawCard';
import GameStats from '../components/GameStats';

import fishbowl from '../assets/Fishbowl3.svg';
import rules from './../gameRules.json';
import './Game.css';


class GamePage extends Component {

  constructor() {
      super();
      this.state = {
          ready: false,
          round: 0,
          team_1_turn: true,
          unused_cards: 0,
          team: "",
          team1: {
            round_1_pts: 0,
            round_2_pts: 0,
            round_3_pts: 0,
            round_4_pts: 0,
          },
          team2: {
            round_1_pts: 0,
            round_2_pts: 0,
            round_3_pts: 0,
            round_4_pts: 0,
          }
      }
      this.saveState = this.saveState.bind(this);
      this.startGame = this.startGame.bind(this);
      this.endGame = this.endGame.bind(this);
      this.nextTurn = this.nextTurn.bind(this);
      this.componentSwitch = this.componentSwitch.bind(this);
  }

  saveState(data) {
    let team1 = data.teams.team_1.name
    let team2 = data.teams.team_2.name
    let currentTeam = data.team_1_turn ? team1 : team2

    this.setState({
      team1: data.teams.team_1,
      team2: data.teams.team_2,
      team_1_turn: data.team_1_turn,
      ready: data.started,
      round: data.current_round,
      unused_cards: data.unused_cards,
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
      })
    .then((response) => {
      this.saveState(response.data.result[0])
    })
    .catch(function (error) {
        console.log(error);
    });
  }

  endGame() {
    const { params: { gameId } } = this.props.match;
    axios({
        method: 'patch',
        url: `/v1/api/game/${gameId}`,
        timeout: 4000,    // 4 seconds timeout
        data: {
            current_round: 5 // Forces end of game
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
    var title;
    var component;

    switch (this.state.round) {
      // Initial game setup
      case 0:
        title = <h2>Enter nouns below to get started!</h2>;
        component = <CardInput gameId={gameId} done={this.startGame}/>;
        break
      case 5: // Force end of game
        title = <h2>Congratulations!!</h2>;
        component = <GameStats gameId={gameId}/>;
        break
      // Transition to Game Stats page at the end of round 4
      case 4:
        if (this.state.unused_cards === 0) { // Natural end of game
          title = <h2>Congratulations!!</h2>;
          component = <GameStats gameId={gameId}/>;
          break
        }
        // fallthrough
      default:
          title = <RoundTracker round={this.state.round} team1={this.state.team1} team2={this.state.team2}/>;
          component = <div>
            <DrawCard
              gameId={gameId}
              gameState={this.state}
              updateState={this.saveState}
              nextRound={this.startGame}
              nextTurn={this.nextTurn}
              endGame={this.endGame}
            />
          </div>;
          break
    }
    return {
      "title": title,
      "component": component
    }
  }

  render() {
      const { params: { gameId } } = this.props.match;
      const element  = this.componentSwitch(gameId)

      return (
        <div className="Game-page">
            <GameTagHeader gameId={gameId}/>
            <div className="title">{element.title}</div>

            <div className="row">
              <div className="col-left">
                {element.component}
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


function RoundTracker({ round, team1, team2 }) {
  let team1_pts = [team1.round_1_pts, team1.round_2_pts, team1.round_3_pts, team1.round_4_pts].reduce((a, b) => a + b, 0)
  let team2_pts = [team2.round_1_pts, team2.round_2_pts, team2.round_3_pts, team2.round_4_pts].reduce((a, b) => a + b, 0)

  return (
    <div className="scorekeeper row">
      <h2><b>{rules.rounds[round-1].name}</b></h2>
      <div>
        <div className="score col-left">
            {team1.name}<br/>
            <b>{team1_pts}</b>
          </div>
        <div className="score col-right">
            {team2.name}<br/>
            <b>{team2_pts}</b>
        </div>
      </div>
    </div>
  )
}
  
  export default GamePage;