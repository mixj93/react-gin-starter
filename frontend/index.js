import * as React from 'react';
import * as ReactDOM from 'react-dom';
import * as apis from './apis';

export default class App extends React.Component {
  constructor(props) {
    super(props);
    this.state = {
      response: "",
      loading: false
    };
  }

  getPing() {
    this.setState({loading: true});
    apis.getPing().then((data) => {
      console.log('data', data)
      this.setState({
        response: data,
        loading: false
      });
    }).catch((err) => {
      console.log('err', err)
      this.setState({loading: false});
    })
  }

  componentDidMount() {
    this.getPing()
  }

  render() {
    return (
      <div>
        <h1>Hello World!</h1>
        <p>Server responses: {this.state.loading ? '... Loading ...' : this.state.response}</p>
      </div>
    )
  }
}

ReactDOM.render(<App />, document.getElementById('root'));
