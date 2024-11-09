import React, { useEffect, useState } from 'react';
import { fetchOrdersByCourier, acceptOrDeclineOrder, updateOrderStatusCourier } from './api';

const AssignedOrdersPage = () => {
  const [orders, setOrders] = useState([]);
  const courierID = localStorage.getItem('courierID'); // Retrieve courierID

  useEffect(() => {
    if (courierID) {
      fetchAssignedOrders();
    }
  }, [courierID]);

  const fetchAssignedOrders = async () => {
    try {
      const data = await fetchOrdersByCourier(courierID); // Pass courierID to fetch orders
      setOrders(data);
    } catch (error) {
      console.error('Error fetching assigned orders:', error);
    }
  };

  const handleAcceptOrDecline = async (orderId, accept) => {
    try {
      await acceptOrDeclineOrder(orderId, accept);
      fetchAssignedOrders(); // Refresh orders list
    } catch (error) {
      console.error('Error updating order acceptance:', error);
    }
  };

  const handleStatusUpdate = async (orderId, status) => {
    try {
      await updateOrderStatusCourier(orderId, status);
      fetchAssignedOrders(); // Refresh orders list
    } catch (error) {
      console.error('Error updating order status:', error);
    }
  };

  return (
    <div>
      <h1>Assigned Orders</h1>
      <ul>
        {orders.map(order => (
          <li key={order.id}>
            <p>Order ID: {order.id}</p>
            <p>Status: {order.status}</p>
            <button onClick={() => handleAcceptOrDecline(order.id, true)}>Accept</button>
            <button onClick={() => handleAcceptOrDecline(order.id, false)}>Decline</button>
            <select onChange={(e) => handleStatusUpdate(order.id, e.target.value)}>
              <option value="">Update Status</option>
              <option value="Picked Up">Picked Up</option>
              <option value="In Transit">In Transit</option>
              <option value="Delivered">Delivered</option>
            </select>
          </li>
        ))}
      </ul>
    </div>
  );
};

export default AssignedOrdersPage;
