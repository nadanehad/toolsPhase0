import React, { useEffect, useState } from 'react';
import { fetchOrderDetails, cancelOrder } from '../api'; 
import { useLocation, useNavigate } from 'react-router-dom';
import './orderDetails.css';

const OrderDetails = () => {
  const location = useLocation();
  const navigate = useNavigate();
  const orderId = location.state?.orderId; 

  const [order, setOrder] = useState(null); 
  const [error, setError] = useState('');   
  const [cancelMessage, setCancelMessage] = useState(''); 

  useEffect(() => {
    if (orderId) {
      const getOrderDetails = async () => {
        try {
          const response = await fetchOrderDetails(orderId);
          setOrder(response);
        } catch (error) {
          setError('Failed to fetch order details. Please try again.');
          console.error(error);
        }
      };
      getOrderDetails();
    } else {
      setError('Order ID is missing. Please navigate from the My Orders page.');
      setTimeout(() => navigate('/my-orders'), 3000); 
    }
  }, [orderId, navigate]);

  const handleCancelOrder = async () => {
    try {
      if (order.status === 'Pending') { 
        await cancelOrder(orderId); 
        setCancelMessage('Order has been canceled successfully.');
        setOrder({ ...order, status: 'Canceled' }); // Update order status locally
      } else {
        setCancelMessage('Only pending orders can be canceled.');
      }
    } catch (error) {
      setCancelMessage('Failed to cancel the order. Please try again.');
      console.error(error);
    }
  };

  return (
    <div className="order-details-container">
      {error && <p style={{ color: 'red' }}>{error}</p>}
      {order ? (
        <>
          <h2>Order ID: {order.id}</h2>
          <p>Pickup Location: {order.pickup_location}</p>
          <p>Dropoff Location: {order.dropoff_location}</p>
          <p>Package Details: {order.package_details}</p>
          <p>Delivery Time: {order.delivery_time}</p>
          <p>Status: {order.status}</p>

          {/* Display cancel message if present */}
          {cancelMessage && <p style={{ color: order.status === 'Canceled' ? 'green' : 'red' }}>{cancelMessage}</p>}

          {/* Render the Cancel Order button if the order status is Pending */}
          {order.status === 'Pending' && (
            <button onClick={handleCancelOrder} className="cancel-order-button">
              Cancel Order
            </button>
          )}
        </>
      ) : (
        !error && <p>Loading order details...</p>
      )}
    </div>
  );
};

export default OrderDetails;