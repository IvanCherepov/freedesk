import React from 'react';
import ReactDOM from 'react-dom';
import './index.css';
import { Router, Route, hashHistory } from 'react-router'
import App from './modules/App';
import About from './modules/About'
import Tickets from './modules/Tickets'
import Ticket from './modules/Ticket'

ReactDOM.render(
    <Router history={hashHistory}>
    <Route path="/" component={App}/>
    <Route path="/tickets" component={Tickets}/>
    <Route path="/tickets/:ticketId" component={Ticket}/>
    <Route path="/about" component={About}/>
  </Router>,
   document.getElementById('root')
);
