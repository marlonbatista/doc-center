import React from "react";
import { Switch, Route, Link } from "react-router-dom";
import "bootstrap/dist/css/bootstrap.min.css";
import "@fortawesome/fontawesome-free/css/all.css";
import "@fortawesome/fontawesome-free/js/all.js";
import "./App.css";
import Footer from "./components/Footer";
import AddDocument from "./components/AddDocument";
import Document from "./components/Document";
import DocumentsList from "./components/DocumentList";


function App() {
  return (
    <div>
      <nav className="navbar navbar-expand navbar-dark bg-dark">
        <a href="/tutorials" className="navbar-brand">
          doc-Center
        </a>
        <div className="navbar-nav mr-auto">
          <li className="nav-item">
            <Link to={"/tutorials"} className="nav-link">
              Documents
            </Link>
          </li>
          <li className="nav-item">
            <Link to={"/add"} className="nav-link">
              Add
            </Link>
          </li>
        </div>
      </nav>
      <footer className = 'na'>
        
      </footer>
      <div className="container mt-3">
        <Switch>
          <Route exact path={["/", "/tutorials"]} component={DocumentsList} />
          <Route exact path="/add" component={AddDocument} />
          <Route path="/tutorials/:id" component={Document} />
        </Switch>
      </div>

    <Footer/>      
    </div>
  );


}


export default App;
