import React, { useEffect, useState } from 'react';
import { fetchUserOrders } from '../api'; // Function to fetch user's orders
import { useNavigate } from 'react-router-dom'; // Import useNavigate for navigation
import './myOrders.css'; // CSS for the My Orders page

const MyOrders = () => {
  const [orders, setOrders] = useState([]); // State to hold orders
  const [error, setError] = useState(''); // State to hold error messages
  const navigate = useNavigate(); // Initialize useNavigate

  useEffect(() => {
    const getOrders = async () => {
      try {
        const response = await fetchUserOrders();
        console.log("Fetched orders response:", response); // Log the fetched orders response
        setOrders(response);
      } catch (error) {
        setError('Failed to fetch orders. Please try again.');
        console.error(error);
      }
    };
    getOrders();
  }, []); // Fetch orders when component mounts

  return (
    <div className="my-orders-container">
      <h2>My Orders</h2>
      {error && <p style={{ color: 'red' }}>{error}</p>}
      <ul>
        {orders.map((order, index) => (
          <li key={order.ID || index} className="order-item">
            {order.ID ? (
              <>
                <div className="order-id">Order ID: {order.ID}</div>
                <p className="order-details">Pickup Location: {order.pickup_location}</p>
                <p className="order-details">Dropoff Location: {order.dropoff_location}</p>
                <p className="order-details">Status: {order.status}</p>
                <button 
                  onClick={() => navigate('/order-details', { state: { orderId: order.ID } })}
                  className="order-button"
                >
                  View Details
                </button>
              </>
            ) : (
              <p style={{ color: 'red' }}>Order ID is missing</p>
            )}
          </li>
        ))}
      </ul>
    </div>
  );
};

export default MyOrders;