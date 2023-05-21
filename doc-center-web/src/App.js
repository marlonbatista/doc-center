import React, { useState, useEffect } from "react";
import { Route, Link, Routes } from "react-router-dom";
import "bootstrap/dist/css/bootstrap.min.css";
import "./App.css";
import "@fortawesome/fontawesome-free/css/all.css";
import "@fortawesome/fontawesome-free/js/all.js";
import { AppBar, Toolbar, Typography } from '@mui/material';
import IconButton from '@mui/material/IconButton';

//import AddDocument from "./components/AddDocument";
//import Document from "./components/Document";
//import DocumentsList from "./components/DocumentList";

import AuthService from "./services/auth.service";

import Login from "./components/Login";
import Register from "./components/Register";
import Home from "./components/Home";
import Profile from "./components/Profile";
import BoardUser from "./components/BoardUser";
import BoardModerator from "./components/BoardModerator";
import BoardAdmin from "./components/BoardAdmin";
import EventBus from "./common/EventBus";

import File from './pages/File'


const App = () => {
  const [showModeratorBoard, setShowModeratorBoard] = useState(false);
  const [showAdminBoard, setShowAdminBoard] = useState(false);
  const [currentUser, setCurrentUser] = useState(undefined);

  useEffect(() => {
    const user = AuthService.getCurrentUser();

    if (user) {
      setCurrentUser(user);
      setShowModeratorBoard(user.roles.includes("ROLE_MODERATOR"));
      setShowAdminBoard(user.roles.includes("ROLE_ADMIN"));
    }

    EventBus.on("logout", () => {
      logOut();
    });

    return () => {
      EventBus.remove("logout");
    };
  }, []);

  const logOut = () => {
    AuthService.logout();
    setShowModeratorBoard(false);
    setShowAdminBoard(false);
    setCurrentUser(undefined);
  };

  return (
    <div>
       <AppBar position="static">
       <Toolbar>
          <IconButton
            size="large"
            edge="start"
            sx={{ mr: 2 }}
          >

          </IconButton>

          <Typography variant="h6" component="div" sx={{ flexGrow: 1, color: "inherit", textDecoration: "none" }}>
          <Link to={"/home"} style={{ color: "white" }}>
            Doc Center
          </Link>
          </Typography>
       
          <Typography variant="h6" component="div" sx={{ color: "inherit", display: "flex", justifyContent: "space-between" }}>
          <Link to={"/login"}  style={{ color: "white" }}>
            Login
          </Link>
          <Link to={"/register"}  style={{ color: "white", marginLeft: "16px"}}>
            SignUp
          </Link>
          </Typography>
       
        </Toolbar>
      </AppBar>
      
      
   
      <div className="container mt-3">
        <Routes>
          <Route path="/" element={<Home/>} />
          <Route path="/home" element={<Home/>} />
          <Route path="/login" element={<Login/>} />
          <Route path="/register" element={<Register/>} />
          <Route path="/profile" element={<Profile/>} />
          <Route path="/user" element={<BoardUser/>} />
          <Route path="/mod" element={<BoardModerator/>} />
          <Route path="/admin" element={<BoardAdmin/>} />
          <Route path="/editFile" element={<File/>} />
         {/* <Route path="/add" element={<AddDocument/>} />          
          <Route path="/tutorials" element={<DocumentsList/>} />
          <Route path="/tutorials/:id" element={<Document/>} /> */}
        </Routes>

      </div>
    </div>
  );
};

export default App;
