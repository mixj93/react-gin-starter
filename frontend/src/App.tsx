import React from 'react'
import { BrowserRouter as Router, Switch, Route, Link } from 'react-router-dom'
import Home from './Home'
import About from './About'
import Contact from './Contact'
import logo from './logo.svg'
import './App.css'

function App() {
  return (
    <div className="App">
      <Router>
        <div>
          <header className="App-header">
            <img src={logo} className="App-logo" alt="logo" />
            <nav className="App-nav">
              <ul>
                <li>
                  <Link to="/">Home</Link>
                </li>
                <li>
                  <Link to="/about">About</Link>
                </li>
                <li>
                  <Link to="/contact">Contact</Link>
                </li>
              </ul>
            </nav>
          </header>

          <Switch>
            <Route path="/about">
              <About />
            </Route>
            <Route path="/contact">
              <Contact />
            </Route>
            <Route path="/">
              <Home />
            </Route>
          </Switch>
        </div>
      </Router>
    </div>
  )
}

export default App
