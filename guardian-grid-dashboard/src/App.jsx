import React, { useState } from 'react';
import { BrowserRouter as Router } from 'react-router-dom';
import Sidebar from './components/Sidebar';
import AppRoutes from './routing/AppRoutes';
import Stage1 from './auth-pages/Stage1';
import "./App.css"
import Stage10 from './auth-pages/Stage10';
import Stage11 from './auth-pages/Stage11';

function App() {
  const [isAuthenticated, setIsAuthenticated] = useState(
    // localStorage.getItem("auth") === "true"
    true
  );

  const handleAuthSuccess = () => {
    localStorage.setItem("auth", "true");
    setIsAuthenticated(true);
  };

  return (
    <Router>
      <AppRoutes/>
    </Router>
  );
}

export default App;