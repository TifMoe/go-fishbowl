import React,  { Component } from 'react';
import fishbowl from '../assets/funfish.svg';

import './CardCounter.css'

class CelebrationFish extends Component {
    render() {
        return (
            <div>
                <div className="logo">
                    <img src={fishbowl} alt="logo" />
                </div>
                {/* TODO: Replace with "Start New Game" button */}
                <div className="counter">
                    <div className="count">
                        Hoooray!!
                    </div>
                </div>
            </div>
        )
    }
}

export default CelebrationFish;