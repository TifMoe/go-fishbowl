import React, { Component } from 'react'
import Confetti from 'react-confetti'
import axios from 'axios';
import './GameStats.css';

class GameStats extends Component {

    constructor() {
        super();
        this.state = {
            team1: {
                name: "",
                round_1_pts: 0,
                round_2_pts: 0,
                round_3_pts: 0,
                round_4_pts: 0
            },
            team2: {
                name: "",
                round_1_pts: 0,
                round_2_pts: 0,
                round_3_pts: 0,
                round_4_pts: 0
            },
        }
    }

    componentDidMount() {
        axios({
            method: 'get',
            url: `/v1/api/game/${this.props.gameId}`,
            timeout: 4000,    // 4 seconds timeout
            })
        .then((response) => {
            this.setState({ 
                team1: response.data.result[0].teams.team_1,
                team2: response.data.result[0].teams.team_2 
            })
        })
        .catch(function (error) {
            console.log(error);
        });
    }

    render() {
        const declareWinner = GetWinner(this.state.team1, this.state.team2)
        return (
            <div>
                {declareWinner}
                <Confetti
                    colors={['#555F7D', '#F25577', '#2EDDCB', '#F4F7B4', '#F2F2F2', '#BF2C5B', '#F6FB96']}
                />
            </div>
            )
    }
}

function GetWinner(team1, team2) {
    let team1pts = getPoints(team1)
    let team2pts = getPoints(team2)
    let headline = "Congratulations!!"
    let winner = (team1pts > team2pts ? team1.name : team2.name) + " Wins!!"

    if (team1pts == team2pts) {
        headline = "Congrats all around!! "
        winner = `${team1.name} and ${team2.name} have tied!!`
    }

    return (
        <div>
        <h2>{headline}</h2>
        <h4>{winner}</h4>
            <ul className="score">
                <li>{team1.name}: {team1pts} cards</li>
                <li>{team2.name}: {team2pts} cards</li>
            </ul>
        </div>
    )
}

function getPoints(team) {
    return (
        team.round_1_pts + team.round_2_pts + team.round_3_pts + team.round_4_pts
    )
}

export default GameStats