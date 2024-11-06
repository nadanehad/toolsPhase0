import React from 'react';
import { BrowserRouter as Router, Route, Routes, Link } from 'react-router-dom';
import Login from './login';    // Login component
import Register from './Register';  // Register component
import Dashboard from './Dashboard';  // Dashboard component (protected)
import './App.css';  // Your styles

function App() {
  const isAuthenticated = !!localStorage.getItem('user');  // Check if user is logged in

  return (
    <Router>
      <div className="App">
        <nav>
          <Link to="/">Login</Link> | <Link to="/register">Register</Link> | {isAuthenticated && <Link to="/dashboard">Dashboard</Link>}
        </nav>
        <Routes>
          <Route path="/" element={<Login />} />
          <Route path="/register" element={<Register />} />
          {/* Protect the dashboard route */}
          <Route path="/dashboard" element={isAuthenticated ? <Dashboard /> : <Login />} />
        </Routes>
      </div>
    </Router>
  );
}

export default App;
