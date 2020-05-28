import React,  { Component } from 'react';
import "./RoundRules.css"
import rules from './../gameRules.json';

// TODO: The cards here could be looped to reduce repeated code
class RoundRules extends Component {
    render() {
        const rounds = rules.rounds;
        return (
            <div>
                <div className="row" style={{minHeight: "500px", width:"100%"}}>
                    <h1 style={{color: "white"}}>4 Games in one!!</h1>
                    <div className="four-col">
                        <div className="rules-card">
                            Round 1
                            <div className="round-title">
                                {rounds[0].name}
                            </div>
                            <div className="round-rules">
                                {rounds[0].rules}
                            </div>
                        </div>
                    </div>
                    <div className="four-col">
                        <div className="rules-card">
                            Round 2
                            <div className="round-title">
                                    {rounds[1].name}
                                </div>
                                <div className="round-rules">
                                    {rounds[1].rules}
                                </div>
                            </div>
                    </div>
                    <div className="four-col">
                        <div className="rules-card">
                            Round 3
                            <div className="round-title">
                                    {rounds[2].name}
                                </div>
                                <div className="round-rules">
                                    {rounds[2].rules}
                                </div>
                            </div>
                    </div>
                    <div className="four-col">
                        <div className="rules-card">
                            Round 4
                            <div className="round-title">
                                    {rounds[3].name}
                                </div>
                                <div className="round-rules">
                                    {rounds[3].rules}
                                </div>
                            </div>
                    </div>
            </div>
            <h2 style={{paddingTop: "20px", color:"white"}}>Now head back up to enter your team names and get started!</h2>
        </div>
        )
    }
}

export default RoundRules;