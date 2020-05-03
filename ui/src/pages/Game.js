import React,  { Component } from 'react';
import axios from 'axios';

import CardInput from '../components/CardInput';
import GameTagHeader from '../components/GameTagHeader';
import DrawCard from '../components/DrawCard';

import fishbowl from '../assets/Fishbowl3.svg';
import './Game.css';


class GamePage extends Component {

  constructor() {
      super();
      this.state = {
          ready: false,
          round: 0
      }
      this.startGame = this.startGame.bind(this);
  }

  componentWillMount() {
    const { params: { gameId } } = this.props.match;
    axios({
        method: 'get',
        url: `/v1/api/game/${gameId}`,
        timeout: 4000,    // 4 seconds timeout
      })
    .then((response) => {
        console.log(response)
        this.setState({ready: response.data.result[0].started})
    })
    .catch(function (error) {
        console.log(error);
    });
  }

  startGame() {
    const { params: { gameId } } = this.props.match;
    axios({
        method: 'patch',
        url: `/v1/api/game/${gameId}`,
        timeout: 4000,    // 4 seconds timeout
        data: {
            started: true,
            current_round: 1
        }
      })
    .then((response) => {
        console.log(response)
        this.setState({ ready: true });
    })
    .catch(function (error) {
        console.log(error);
    });
  }

  render() {
      const { params: { gameId } } = this.props.match;
      return (
        <div className="Game-page">
            <GameTagHeader gameId={gameId}/>
            <h2 className="title">Fair warning... this fishbowl is still under development</h2>

            <div className="row">
              <div className="col-left">
                { this.state.ready ?
                  <DrawCard gameId={gameId}/> :
                  <CardInput gameId={gameId} done={this.startGame}/>
                }
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
  
  export default GamePage;