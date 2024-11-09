import React, { useState } from 'react';
import { loginUser } from '../api'; // Importing the API call for login
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
      // Making API call to log in the user
      const { role } = await loginUser(formData);
      
      // Logging response for debugging purposes
      console.log('User role:', role);
      
      // Store the role in localStorage
      localStorage.setItem('role', role);

      // Setting the success message
      setMessage('Login successful!');

      // Redirecting the user based on their role
      if (role === "User") {
        navigate('/home'); // Redirect to Home page for User role
      } else if (role === "Admin") {
        navigate('/manage-orders'); // Redirect to Manage Orders for Admin role
      } else if (role === "Courier") {
        navigate('/assigned-orders'); // Redirect to Assigned Orders for Courier role
      } else {
        setError('Unrecognized role'); // Error if role is unknown
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
