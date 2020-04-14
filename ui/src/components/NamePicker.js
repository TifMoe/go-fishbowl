import React, { Component } from 'react';
import { Link } from "react-router-dom";
import axios from 'axios';

import './NamePicker.css';

class Picker extends Component {

    constructor() {
        super();
        this.state = {
            msg: 'pending'
        }
    }

    componentWillMount() {
        axios.get('v1/api/random/name')
            .then((response) => {
                this.setState(() => {
                    return { msg: response.data.message }
                })
            })
            .catch(function (error) {
                console.log(error);
            });

    }

    render() {
        return (
        <div>
            <Link to={`/game/${this.state.msg}`} >
                <button>
                    Start a new game!
                </button>
            </Link>
            <p> /{this.state.msg} </p>
        </div>
        );
    }
}

export default Picker; 