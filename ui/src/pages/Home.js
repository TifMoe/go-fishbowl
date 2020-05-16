import React from 'react';
import { useHistory } from "react-router-dom";
import { useSpring, animated } from 'react-spring'

import {ReactComponent as TitleImg} from '../assets/TitleImage.svg';
import './Home.css';
import NewGame from '../components/NewGame';

const AnimFeTurbulence = animated('feTurbulence')
const AnimFeDisplacementMap = animated('feDisplacementMap')

function HomePage() {
    const { freq, scale, transform, opacity } = useSpring({
      from: { scale: 0, opacity: 0, transform: 'scale(0.9)', freq: '0.01, 0.0' },
      to: { scale: 100, opacity: 1, transform: 'scale(1)', freq: '0.0, 0.0' },
      config: { duration: 2000 }
    })
    const history = useHistory();
    return (
        <header className="Home-header">
           <div className="TitleImg">
              <animated.svg style={{ transform, opacity }} viewBox="0 0 800 250" width="100%" height="100%">
                <defs>
                  <filter id="water">
                    <AnimFeTurbulence type="fractalNoise" baseFrequency={freq} numOctaves="1.5" result="TURB" seed="8" />
                    <AnimFeDisplacementMap xChannelSelector="R" yChannelSelector="G" in="SourceGraphic" in2="TURB" result="DISP" scale={scale} />
                  </filter>
                </defs>
                <g filter="url(#water)">
                  <TitleImg/>
                </g>
              </animated.svg>
            </div>
          <div className="intro">
            <h2>Carry on with game nights from a distance!</h2>
            <p>Play fishbowl online over video calls with friends and family</p>
            <NewGame history={history}/>
            <p style={{paddingInlineStart: "15%", paddingInlineEnd: "15%"}}>Keep scrolling to check out the game rules or enter two cool team names above to get started!</p>
          </div>
        </header>
    );
  }
  
  export default HomePage;