
import React,  { Component } from 'react';
import './ScoreKeeper.css'

class ScoreKeeper extends Component {
    render() {
        let team1 = this.props.team1
        let team2 = this.props.team2

        let team1_pts = [team1.round_1_pts, team1.round_2_pts, team1.round_3_pts, team1.round_4_pts].reduce((a, b) => a + b, 0)
        let team2_pts = [team2.round_1_pts, team2.round_2_pts, team2.round_3_pts, team2.round_4_pts].reduce((a, b) => a + b, 0)
        return (
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
        )
    }
}

export default ScoreKeeper;