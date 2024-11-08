import React, { useState } from 'react';
import { loginUser } from '../api';  // Importing the API call for login
import { useNavigate } from 'react-router-dom'; // Importing useNavigate for routing
import './login.css'; // Importing CSS for styling

const Login = () => {
  // State to manage form data
  const [formData, setFormData] = useState({
    email: '',
    password: ''
  });
  
  // State to manage error and success messages
  const [error, setError] = useState('');
  const [message, setMessage] = useState('');
  
  // Hook for navigating routes
  const navigate = useNavigate(); 

  // Handles input change and updates formData state
  const handleChange = (e) => {
    setFormData({
      ...formData,
      [e.target.name]: e.target.value
    });
  };

  // Handles form submission for login
  const handleSubmit = async (e) => {
    e.preventDefault(); // Prevents the default form submission action
    setError('');
    setMessage('');

    try {
      // Making API call to login the user
      const response = await loginUser(formData);
      
      // Logging response for debugging purposes
      console.log('Login response:', response);
      
      // Save the token and role to localStorage
      localStorage.setItem('token', response.token); // Save token to local storage
      localStorage.setItem('role', response.role);   // Save role to local storage
      
      // Setting the success message
      setMessage('Login successful!');

      // Redirecting the user based on their role
      if (response.role === "User") {
        navigate('/home'); // Redirect to Home page for User role
      } else {
        navigate('/dashboard'); // Redirect to Dashboard for Admin or Courier role
      }
    } catch (error) {
      // Setting the error message if login fails
      setError('Invalid email or password');
      console.error('Login error:', error);
    }
  };

  return (
    <div className="login-container">
      <h2>Login</h2>
      {/* Displaying success or error messages */}
      {message && <p style={{ color: 'green' }}>{message}</p>}
      {error && <p style={{ color: 'red' }}>{error}</p>}
      
      <form onSubmit={handleSubmit}>
        {/* Email Input Field */}
        <div>
          <label>Email</label>
          <input
            type="email"
            name="email"
            value={formData.email}
            onChange={handleChange}
            placeholder="Enter your email"
            required
          />
        </div>
        
        {/* Password Input Field */}
        <div>
          <label>Password</label>
          <input
            type="password"
            name="password"
            value={formData.password}
            onChange={handleChange}
            placeholder="Enter your password"
            required
          />
        </div>
        
        {/* Submit Button */}
        <button type="submit" className="login-btn">Login</button>
      </form>
      
      {/* Link to navigate to Registration Page */}
      <p>
        Don't have an account?{' '}
        <span 
          onClick={() => navigate('/register')} 
          style={{ color: '#28a745', cursor: 'pointer' }}>
          Register here
        </span>
      </p>
    </div>
  );
};

export default Login;