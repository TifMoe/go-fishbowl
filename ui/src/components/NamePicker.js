import React, { Component } from 'react'
import axios from 'axios';
class Picker extends Component {

    constructor() {
        super();
        this.state = {
            msg: 'waiting'
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
            <h1>Testing the Go server: </h1>
            <p>
                Please give me a new random name when I refresh ... 
            </p>
            <h3>" {this.state.msg} "</h3>
        </div>
        );
    }
}

export default Picker; 