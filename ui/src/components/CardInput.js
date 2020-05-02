import React, { Component } from 'react';
import axios from 'axios';
import './CardInput.css';

class CardInput extends Component {

    constructor() {
        super();
        this.state = {
            card: '',
        };
        this.onSubmit = this.onSubmit.bind(this);
    }

    onChange = (e) => {
        this.setState({ [e.target.name]: e.target.value });
    }

    onSubmit(e) {
        e.preventDefault();

        axios({
            method: 'post',
            url: '/v1/api/game/' + this.props.gameId + '/card',
            timeout: 4000,    // 4 seconds timeout
            data: {
                value: this.state.card,
            }
          })
        .then((result) => {
            console.log(result)
            this.setState({card: ""})
        })
        .catch(function (error) {
            console.log(error);
        });
      }

    render() {
        const { card } = this.state;
        return (
            <div className="card-input">
                <div className="explainer">
                    <p> Invite players to enter nouns below to fill up your bowl! </p>
                    <p className="small-text"> (max 30 cards total) </p>
                </div>
                    <form onSubmit={this.onSubmit}>
                        <input
                            type="text"
                            name="card"
                            value={card}
                            maxlength="30"
                            minLength="2"
                            onChange={this.onChange}
                        />
                        <button type="submit">Drop it in!</button>
                    </form>
            </div>
        );
    }
}

export default CardInput; 