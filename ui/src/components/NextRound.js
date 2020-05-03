import React, { Component } from 'react';
import axios from 'axios';

import './NextRound.css';

class NextRound extends Component {

    constructor() {
        super();
        this.state = {
            round: 1,
            disabled: true
        }
        this.nextRound = this.nextRound.bind(this);
    }

    nextRound() {
        // set all cards to un-used state to start next round
        // TODO: This should also set round number in backend
        axios({
            method: 'patch',
            url: `/v1/api/game/${this.props.gameId}/start`,
            timeout: 4000,    // 4 seconds timeout
          })
        .then((response) => {
            console.log(response);
            this.setState({round: this.state.round + 1});
            this.setState({disabled: true})
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
                disabled={!this.props.ready && this.state.disabled}
            >Next Round</button>
        </div>
        );
    }
}

export default NextRound; 