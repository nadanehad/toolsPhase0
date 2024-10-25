import React, { useState } from 'react';
import { registerUser } from './api'; // Import your API function

const Register = () => {
  const [formData, setFormData] = useState({
    name: '',       
    email: '',
    phone: '',      
    password: '',
    confirmPassword: ''
  });
  
  const [message, setMessage] = useState('');
  const [error, setError] = useState('');

  const { name, email, phone, password, confirmPassword } = formData;

  const handleChange = (e) => {
    setFormData({
      ...formData,
      [e.target.name]: e.target.value
    });
  };

  const handleSubmit = async (e) => {
    e.preventDefault();

    setError('');
    setMessage('');

    if (password.length < 8) {
        setError('Password must be at least 8 characters long');
        return;
    }

    if (password !== confirmPassword) {
        setError('Passwords do not match');
        return;
    }

    try {
        const response = await registerUser({
            name,
            email,
            phone,
            password,
        });
        setMessage('Registration successful!');
    } catch (error) {
        setError('Failed to register user');
        console.error('Error details:', error.response?.data || error.message);
    }
};

  

  return (
    <div className="register-container">
      <h2>Register</h2>
      {message && <p style={{ color: 'green' }}>{message}</p>}
      {error && <p style={{ color: 'red' }}>{error}</p>}
      <form onSubmit={handleSubmit}>
        <div>
          <label>Name</label>
          <input
            type="text"
            name="name"
            value={name}
            onChange={handleChange}
            placeholder="Enter your name"
            required
          />
        </div>
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
          <label>Phone</label>
          <input
            type="text"
            name="phone"
            value={phone}
            onChange={handleChange}
            placeholder="Enter your phone number"
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
        <div>
          <label>Confirm Password</label>
          <input
            type="password"
            name="confirmPassword"
            value={confirmPassword}
            onChange={handleChange}
            placeholder="Confirm your password"
            required
          />
        </div>
        <button type="submit">Register</button>
      </form>
    </div>
  );
};

export default Register;
