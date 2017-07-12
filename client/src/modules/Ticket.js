import React from 'react'
import Client from './Client'

// Switched to using native JavaScript classes
// (supported because this code gets transpiled into legacy javascript)
export default class Comments extends React.Component {
  constructor(props) {
    super(props);
    // initialize state
    this.state = {
      data: []
    };
  }

  // When the component mounts, perform our network request
  componentDidMount() {
    Client.search(this.props.params.ticketId)
      .then((data) => {
        this.setState({
          data: data
        });
      })
      .catch((error) => {
        alert('An error ocurred while trying to fetch the ticket from Zendesk');
        throw error;
      });
  }

  render() {
    let dataToDisplay;
    if (!this.state.data) {
      dataToDisplay = 'Loading...';
    } else {
      dataToDisplay = JSON.stringify(this.state.data);
    }

    return <div>
       <ul>
       {this.state.data.map(function(comment,i){
	return <li key={i}>{comment.body}</li>
	}
	
	)}
       </ul>
      </div>
  }
}
