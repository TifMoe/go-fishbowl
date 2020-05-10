import React, { Component } from 'react';
import axios from 'axios';

import './NewGame.css';

class NewGame extends Component {

    constructor() {
        super();
        this.state = {
            gameName: 'pending',
            team1: "",
            team2: "",
        }
        this.onSubmit = this.onSubmit.bind(this);
        this.routeChange = this.routeChange.bind(this);
    }

    onChange = (e) => {
        this.setState({ [e.target.name]: e.target.value });
    }

    routeChange = () => {
        let path = `/game/${this.state.gameName}`; 
        this.props.history.push(path)
    }


    onSubmit(e) {
        e.preventDefault();
        // Save team names and fetch new game namespace
        axios({
            method: 'post',
            url: 'v1/api/game',
            timeout: 4000,    // 4 seconds timeout
            data: {
                team_1: capitalize(this.state.team1),
                team_2: capitalize(this.state.team2),
            }
          })
        .then((response) => {
            this.setState(() => {
                return { gameName: response.data.message }
            });
            this.routeChange();
        })
        .catch(function (error) {
            console.log(error);
        });
      }

    render() {
        const { team1 } = this.state;
        const { team2 } = this.state;

        return (
        <div>
            <div className="card-form">
                <form onSubmit={this.onSubmit}>
                    <input
                        className="team-input pink"
                        type="text"
                        name="team1"
                        placeholder="Team 1 Name"
                        value={team1}
                        maxLength="30"
                        minLength="2"
                        onChange={this.onChange}
                    />
                    <input
                        className="team-input blue"
                        type="text"
                        name="team2"
                        placeholder="Team 2 Name"
                        value={team2}
                        maxLength="30"
                        minLength="2"
                        onChange={this.onChange}
                    />
                    <button type="submit"> Start a new game! </button>
                </form>
            </div>

        </div>
        );
    }
}

function capitalize(str) {
    return str
        .toLowerCase()
        .split(' ')
        .map(function(word) {
            return word[0].toUpperCase() + word.substr(1);
        })
        .join(' ');
}

export default NewGame; 