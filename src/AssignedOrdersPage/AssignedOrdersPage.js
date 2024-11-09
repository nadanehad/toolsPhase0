import React, { useEffect, useState } from 'react';
import { fetchOrdersByCourier, acceptOrDeclineOrder, updateOrderStatusCourier } from '../api';
import { useNavigate, useLocation } from 'react-router-dom';

const AssignedOrdersPage = () => {
  const [orders, setOrders] = useState([]);
  const [loading, setLoading] = useState(true);
  const [error, setError] = useState(null);
  const courierID = localStorage.getItem('courierID');
  const navigate = useNavigate();
  const location = useLocation();

  useEffect(() => {
    if (courierID) {
      fetchAssignedOrders();
    } else {
      console.warn("Courier ID is missing in local storage.");
      navigate('/login'); 
    }
  }, [courierID, navigate]);

  const fetchAssignedOrders = async () => {
    try {
      setLoading(true);
      const data = await fetchOrdersByCourier(courierID);
      console.log("Fetched orders:", data); 
      setOrders(data);
    } catch (error) {
      console.error('Error fetching assigned orders:', error);
      setError("Failed to load assigned orders. Please try again.");
    } finally {
      setLoading(false);
    }
  };

  const handleAcceptOrDecline = (orderID, accept) => {
    if (!orderID) {
      console.error("Order ID is missing for accept/decline action.");
      setError("Unable to accept/decline order. Order ID is missing.");
      return;
    }
    console.log(`Setting state for orderID ${orderID} with accept: ${accept}`);
    navigate('/assigned-orders', { state: { orderID, accept } });
  };

  const handleStatusUpdate = (orderID, status) => {
    if (!orderID) {
      console.error("Order ID is missing for status update.");
      setError("Unable to update status. Order ID is missing.");
      return;
    }
    console.log(`Setting state for orderID ${orderID} with status: ${status}`);
    navigate('/assigned-orders', { state: { orderID, status } });
  };

  useEffect(() => {
    if (location.state && location.state.orderID) {
      const { orderID, accept, status } = location.state;

      if (accept !== undefined) {
        console.log(`Received orderID ${orderID} for action: ${accept ? 'Accept' : 'Decline'}`);
        acceptOrDeclineOrderAction(orderID, accept);
      } else if (status) {
        console.log(`Received orderID ${orderID} for status update to: ${status}`);
        updateOrderStatusAction(orderID, status);
      }
    }
  }, [location.state]);

  const acceptOrDeclineOrderAction = async (orderID, accept) => {
    try {
      console.log(`Attempting to ${accept ? 'accept' : 'decline'} order with ID ${orderID}`);
      await acceptOrDeclineOrder(orderID, accept);
      fetchAssignedOrders(); 
    } catch (error) {
      console.error('Error updating order acceptance:', error);
      setError("Failed to update order acceptance. Please try again.");
    }
  };

  const updateOrderStatusAction = async (orderID, status) => {
    try {
      console.log(`Attempting to update status for order with ID ${orderID} to ${status}`);
      await updateOrderStatusCourier(orderID, status);
      fetchAssignedOrders(); 
    } catch (error) {
      console.error('Error updating order status:', error);
      setError("Failed to update order status. Please try again.");
    }
  };

  return (
    <div>
      <h1>Assigned Orders</h1>
      {loading ? (
        <p>Loading assigned orders...</p>
      ) : error ? (
        <p style={{ color: 'red' }}>{error}</p>
      ) : orders.length === 0 ? (
        <p>No assigned orders available for this courier.</p>
      ) : (
        <ul>
          {orders.map((order) => (
            <li key={order.ID || Math.random()} style={{ marginBottom: '20px', padding: '10px', border: '1px solid #ccc' }}>
              <p><strong>Order ID:</strong> {order.ID || "ID not found"}</p>
              <p><strong>Status:</strong> {order.status}</p>
              <div style={{ marginTop: '10px' }}>
                <button onClick={() => handleAcceptOrDecline(order.ID, true)} style={{ marginRight: '5px' }}>
                  Accept
                </button>
                <button onClick={() => handleAcceptOrDecline(order.ID, false)} style={{ marginRight: '5px' }}>
                  Decline
                </button>
                <select
                  onChange={(e) => handleStatusUpdate(order.ID, e.target.value)}
                  defaultValue=""
                  style={{ marginLeft: '10px' }}
                >
                  <option value="" disabled>Update Status</option>
                  <option value="Picked Up">Picked Up</option>
                  <option value="In Transit">In Transit</option>
                  <option value="Delivered">Delivered</option>
                </select>
              </div>
            </li>
          ))}
        </ul>
      )}
    </div>
  );
};

export default AssignedOrdersPage;
