import React, { Component } from 'react';
import axios from 'axios';

import './NextRound.css';

class NextRound extends Component {

    constructor() {
        super();
        this.state = {
            round: 1,
        }
        this.nextRound = this.nextRound.bind(this);
    }

    nextRound() {
        // set all cards to un-used state to start next round
        axios({
            method: 'put',
            url: `/v1/api/game/${this.props.gameId}/start`,
            timeout: 4000,    // 4 seconds timeout
          })
        .then((response) => {
            this.setState({round: response.data.result[0].current_round});
        })
        .catch(function (error) {
            console.log(error);
        });
    }

    render() {
        return (
        <div>
            <button 
                onClick={this.nextRound}
                className="next-round"
                disabled={!this.props.active}
            >Next Round</button>
        </div>
        );
    }
}

export default NextRound; 