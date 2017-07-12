import React from 'react'
import Client from './Client'
import { Link } from 'react-router'

// Switched to using native JavaScript classes
// (supported because this code gets transpiled into legacy javascript)
export default class Tickets extends React.Component {
  constructor(props) {
    super(props);
    // initialize state
    this.state = {
	data: []
    };
  }

  // When the component mounts, perform our network request
  componentDidMount() {
    Client.listIndex()
      .then((data) => {
        this.setState({
          data: data 
        });
      })
      .catch((error) => {
        alert('An error ocurred while trying to fetch list of tickets from Zendesk');
        throw error;
      });
  }

  render() {
    return <div>
      <ul>
      {this.state.data.map(function(ticket, i){
        //return <li key={i}>{ticket.subject}</li>
	
        return <li><Link to={`/tickets/${ticket.id}`}>{ticket.subject}</Link></li>	
        }
     )}
      </ul>
    </div>
  }
} 
