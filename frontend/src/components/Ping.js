import React, { Component } from 'react'
import axios from 'axios';
class PingComponent extends Component {

    constructor() {
        super();
        this.state = {
            msg: 'waiting'
        }
    }

    componentWillMount() {
        axios.get('api/ping')
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
            <h3>Testing the go server... </h3>
            <h1>" {this.state.msg} "</h1>
        </div>
        );
    }
}

export default PingComponent; 