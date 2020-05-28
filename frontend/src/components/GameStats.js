import React, { Component } from 'react'
import Confetti from 'react-confetti'
import './GameStats.css';

class GameStats extends Component {

    render() {
        const declareWinner = GetWinner(this.props.gameState.team1, this.props.gameState.team2)
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
    let winner = (team1pts > team2pts ? team1.name : team2.name) + " Wins!!"

    if (team1pts === team2pts) {
        winner = `${team1.name} and ${team2.name} have tied!!`
    }

    return (
        <div>
        <h4>{winner}</h4>
        <div className="scorekeeper row">
            <div className="score col-left">
                {team1.name}<br/>
                <b>{team1pts}</b> cards
            </div>
            <div className="score col-right">
                {team2.name}<br/>
                <b>{team2pts}</b> cards
            </div>
        </div>
        </div>
    )
}

function getPoints(team) {
    return (
        team.round_1_pts + team.round_2_pts + team.round_3_pts + team.round_4_pts
    )
}

export default GameStats