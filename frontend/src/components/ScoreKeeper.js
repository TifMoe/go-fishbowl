import React,  { Component } from 'react';
import { Row, Col } from 'react-bootstrap';

import './ScoreKeeper.css'

class ScoreKeeper extends Component {

    render() {
        let team1 = this.props.gameState.team1
        let team2 = this.props.gameState.team2

        let team1_pts = [team1.round_1_pts, team1.round_2_pts, team1.round_3_pts, team1.round_4_pts][this.props.gameState.round-1]
        let team2_pts = [team2.round_1_pts, team2.round_2_pts, team2.round_3_pts, team2.round_4_pts][this.props.gameState.round-1]
        return (
                <Row style={{ backgroundColor: "white", minHeight: "80px"}}>
                    <Col sm={4} style={{color: "rgb(242, 85, 119)"}}>
                        <div className="score">{team1.name}</div>
                    </Col>
                    <Col sm={2} style={{color: "rgb(242, 85, 119)"}}>
                        <div className="score"><b>{team1_pts}</b></div>
                    </Col>
                    <Col sm={2} style={{color: "rgb(46, 221, 204)"}}>
                        <div className="score"><b>{team2_pts}</b></div>
                    </Col>
                    <Col sm={4} style={{color: "rgb(46, 221, 204)"}}>
                        <div className="score">{team2.name}</div>
                    </Col>
                </Row>
        )
    }
}

export default ScoreKeeper;