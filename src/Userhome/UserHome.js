import React from 'react';
import { useNavigate } from 'react-router-dom'; 
import './Userhome.css'; 

const Home = () => {
  const navigate = useNavigate();

  return (
    <div className="home-container">
      <h2>Welcome to the Package Tracking System</h2>
      <button onClick={() => navigate('/create-order')} className="home-button">Create Order</button>
      <button onClick={() => navigate('/my-orders')} className="home-button">My Orders</button>
    </div>
  );
};

export default Home;