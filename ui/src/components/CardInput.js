import React, { Component } from 'react';
import axios from 'axios';
import './CardInput.css';

class CardInput extends Component {

    constructor() {
        super();
        this.state = {
            card: '',
            count: 0,
        };
        this.onSubmit = this.onSubmit.bind(this);
        this.getCardCount = this.getCardCount.bind(this);
    }

    getCardCount() {
        axios({
            method: 'get',
            url: '/v1/api/game/' + this.props.gameId,
            timeout: 4000,    // 4 seconds timeout
          })
        .then((response) => {
            console.log(response);
            this.setState({count: response.data.result[0].unused_cards})
        })
        .catch(function (error) {
            console.log(error);
        });
    }

    onChange = (e) => {
        this.setState({ [e.target.name]: e.target.value });
    }

    onSubmit(e) {
        e.preventDefault();
        if (this.state.card != "") {
            axios({
                method: 'post',
                url: '/v1/api/game/' + this.props.gameId + '/card',
                timeout: 4000,    // 4 seconds timeout
                data: {
                    value: capitalize(this.state.card),
                }
              })
            .then((result) => {
                console.log(result)
                this.setState({card: ""})
                this.getCardCount();
            })
            .catch(function (error) {
                console.log(error);
            });
        }
      }

    render() {
        const { card } = this.state;
        return (
            <div className="card-input">
                <div className="explainer">
                    <p> Invite players to submit nouns below to fill up your bowl! </p>
                    <p className="small-text"> (max 50 cards total) </p>
                </div>
                <div className="card-form">
                    <form onSubmit={this.onSubmit}>
                            <input
                                type="text"
                                name="card"
                                value={card}
                                maxLength="30"
                                minLength="2"
                                onChange={this.onChange}
                            />
                            <button type="submit">Drop it in!</button>
                    </form>
                </div>
                <StartGame startHandler={this.props.done} ready={this.state.count >= 3}/>
            </div>
        );
    }
}

const StartGame = ({ startHandler, ready }) => (
    <div>
        <button
            className="start-button"
            onClick={startHandler}
            disabled={!ready}
        >Start Game</button>
    </div>
)

function capitalize(str) {
    return str
        .toLowerCase()
        .split(' ')
        .map(function(word) {
            return word[0].toUpperCase() + word.substr(1);
        })
        .join(' ');
}

export default CardInput; 