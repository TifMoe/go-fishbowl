import React, { Component } from 'react'
import { Container, Row, Col } from 'react-bootstrap';

import './GameStats.css';

class GameStats extends Component {

    // newGame is event emitter to tell backend to start new game in namespace
    newGame = () => {
        let data = JSON.stringify({
            gameID: this.props.gameId,
        });
        console.log('Starting new game: ', this.props.gameId);
        this.props.socket.emit('resetGame', data);
    }

    render() {
        const team1 = this.props.gameState.team1
        const team2 = this.props.gameState.team2
        let team1pts = getPoints(team1)
        let team2pts = getPoints(team2)

        let winner = (team1pts > team2pts ? team1.name : team2.name) + " Wins!!"
        if (team1pts === team2pts) {
            winner = `${team1.name} and ${team2.name} have tied!!`
        }

        return (
            <div>
                <Container style={{marginLeft: "10%", marginRight: "10%"}}>
                    <Row style={{marginBottom: "50px", marginTop: "50px"}}>
                        <Col>
                            <h2>{winner}</h2>
                        </Col>
                    </Row>
                    <Row style={{ 
                        backgroundColor: "white", 
                        minHeight: "120px",
                        borderRadius: "5px",
                    }}>
                        <Col sm={6} style={{color: "rgb(242, 85, 119)"}}>
                            <div className="score">
                                {team1.name}<br/>
                                <b>{team1pts}</b> cards
                            </div>
                        </Col>
                        <Col sm={6} style={{color: "rgb(46, 221, 204)"}}>
                            <div className="score">
                                {team2.name}<br/>
                                <b>{team2pts}</b> cards
                            </div>
                        </Col>
                    </Row>
                    <Row>
                        <Col>
                            <button onClick={this.newGame} className="next-round">
                                Start New Game
                            </button>
                        </Col>
                    </Row>
             </Container>
            </div>
            )
    }
}


function getPoints(team) {
    return (
        team.round_1_pts + team.round_2_pts + team.round_3_pts + team.round_4_pts
    )
}

export default GameStats