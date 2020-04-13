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
            <h1>Testing the go server... </h1>
            <h3>" {this.state.msg} "</h3>
        </div>
        );
    }
}

export default PingComponent; 