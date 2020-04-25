import React, { Component } from 'react';
import axios from 'axios';
import './NamePicker.css';

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
            <form onSubmit={this.onSubmit}>
                <input
                type="text"
                name="card"
                value={card}
                onChange={this.onChange}
                />
                <button type="submit">Drop in bowl!</button>
            </form>
        );
    }
}

export default CardInput; 