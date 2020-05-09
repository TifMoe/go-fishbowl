import React, { Component } from 'react';

import './NewGame.css';

class ScoreKeeper extends Component {
    render() {
        return (
            <div>
                <p>Current Round: {this.props.round}</p>
            </div>
        )
    }
}
export default ScoreKeeper