import React, { Component } from 'react';
import { Link } from 'react-router';

class App extends Component {
  render() {
    return (
      <div className="App">
        <h1>Free Desk</h1>
        <ul role="nav">
          <li><Link to="/about">About</Link></li>
          <li><Link to="/tickets">Tickets</Link></li>

	  {this.props.children}

        </ul>
      </div>
    );
  }
}

export default App;
