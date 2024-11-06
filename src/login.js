import React, { useState } from 'react';
import { loginUser } from './api'; // Import the login API call

const Login = () => {
  const [formData, setFormData] = useState({
    email: '',
    password: ''
  });
  const [error, setError] = useState('');
  const [message, setMessage] = useState('');

  const { email, password } = formData;

  const handleChange = (e) => {
    setFormData({
      ...formData,
      [e.target.name]: e.target.value
    });
  };

  const handleSubmit = async (e) => {
    e.preventDefault();
    setError(''); // Clear any previous error message
    setMessage(''); // Clear any previous success message

    try {
      const response = await loginUser({ email, password });
      setMessage('Login successful!'); // Success message
      console.log('Login successful:', response);
    } catch (error) {
      setError('Invalid email or password'); // Error message
      console.error(error);
    }
  };

  return (
    <div className="login-container">
      <h2>Login</h2>
      {message && <p style={{ color: 'green' }}>{message}</p>} {/* Success message */}
      {error && <p style={{ color: 'red' }}>{error}</p>} {/* Error message */}
      <form onSubmit={handleSubmit}>
        <div>
          <label>Email</label>
          <input
            type="email"
            name="email"
            value={email}
            onChange={handleChange}
            placeholder="Enter your email"
            required
          />
        </div>
        <div>
          <label>Password</label>
          <input
            type="password"
            name="password"
            value={password}
            onChange={handleChange}
            placeholder="Enter your password"
            required
          />
        </div>
        <button type="submit">Login</button>
      </form>
    </div>
  );
};

export default Login;
