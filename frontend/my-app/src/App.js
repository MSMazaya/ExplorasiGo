import './index.css'
import './App.css'
import React,{useEffect} from 'react'
import Navbar from './components/Navbar'
import Phone from './components/Phone'
import Button from './components/Button';
import {
  BrowserRouter as Router,
  Switch,
  Route,
} from "react-router-dom";

function App() {
  
  return (
    <div>
      <Router>
        <Switch>
          <Route path="/phone">
            <Phone/>
          </Route>
          <Route path="/">            
            <Navbar/>
            <Button/> 
          </Route>
        </Switch>
      </Router>
    </div>
  );
}

export default App;
