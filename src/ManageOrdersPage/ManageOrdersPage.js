import React, { useEffect, useState } from 'react';
import {
  fetchAllOrdersAdmin,
  updateOrderAdmin,
  deleteOrderAdmin,
  assignOrderToCourier,
  reassignOrderToCourier
} from '../api';
import './ManageOrdersPage.css';

const ManageOrdersPage = () => {
  const [orders, setOrders] = useState([]);
  const [courierIdMap, setCourierIdMap] = useState({});
  const [error, setError] = useState('');

  useEffect(() => {
    fetchAllOrders();
  }, []);

  const fetchAllOrders = async () => {
    try {
      const data = await fetchAllOrdersAdmin();
      setOrders(data);
    } catch (error) {
      console.error('Error fetching all orders:', error);
      setError('Failed to load orders. Please try again later.');
    }
  };

  const handleStatusUpdate = async (orderId, newStatus) => {
    const orderToUpdate = orders.find((order) => order.ID === orderId);

    if (!orderToUpdate) {
      console.error("Order not found for status update.");
      setError("Order not found. Please refresh and try again.");
      return;
    }

    if (orderToUpdate.status !== 'Accepted by Courier') {
      setError('You can only update the status after the courier accepts the order.');
      return;
    }

    const updatedOrderData = {
      pickup_location: orderToUpdate.pickup_location,
      dropoff_location: orderToUpdate.dropoff_location,
      package_details: orderToUpdate.package_details,
      delivery_time: orderToUpdate.delivery_time,
      delivery_fee: orderToUpdate.delivery_fee,
      courier_id: orderToUpdate.courier_id,
      status: newStatus,
    };


    try {
      await updateOrderAdmin(orderId, updatedOrderData);
      fetchAllOrders(); 
    } catch (error) {
      console.error("Error updating order status:", error);
      setError("Failed to update order status. Please try again.");
    }
  };

  const handleDeleteOrder = async (orderId) => {
    try {
      await deleteOrderAdmin(orderId);
      fetchAllOrders(); 
    } catch (error) {
      console.error('Error deleting order:', error);
    }
  };

  const handleAssignNewCourier = async (orderId) => {
    let courierId = courierIdMap[orderId];  // Retrieve the courier ID from state
    
    // Ensure courierId is a valid number
    courierId = parseInt(courierId, 10);
    
    if (isNaN(courierId)) {
      setError('Please enter a valid courier ID.');
      return;
    }
  
    const payload = { order_id: orderId, courier_id: courierId };
  
    console.log('Sending assign order payload:', payload);  // Debugging the payload
  
    try {
      await assignOrderToCourier(orderId, courierId);  // Call the API to assign the courier
      fetchAllOrders();  // Refresh the orders
    } catch (error) {
      console.error('Error assigning order to new courier:', error);
      setError('Failed to assign order to new courier. Please try again.');
    }
  };

  const handleReassignNewCourier = async (orderId) => {
    let courierId = courierIdMap[orderId];
  
    // Ensure courierId is a valid number
    courierId = parseInt(courierId, 10);
  
    if (isNaN(courierId)) {
      setError('Please enter a valid courier ID.');
      return;
    }
    
  
    try {
      await reassignOrderToCourier(orderId, courierId);  // Call the API for reassigning the courier
      fetchAllOrders(); // Refresh orders after successful reassignment
    } catch (error) {
      console.error('Error reassigning order to new courier:', error);
      setError('Failed to reassign order to new courier. Please try again.');
    }
  };

  const handleCourierIdChange = (orderId, value) => {
    setCourierIdMap((prevMap) => ({
      ...prevMap,
      [orderId]: value,
    }));
  };

  return (
    <div className="manage-orders-container">
      <h1>Manage Orders</h1>
      {error && <p style={{ color: 'red' }}>{error}</p>}
      <table className="orders-table">
        <thead>
          <tr>
            <th>Order ID</th>
            <th>Status</th>
            <th>Courier ID</th>
            <th>Actions</th>
          </tr>
        </thead>
        <tbody>
          {orders.map((order) => (
            <tr key={order.ID}>
              <td>{order.ID}</td>
              <td>{order.status}</td>
              <td>{order.courier_id || 'Unassigned'}</td>
              <td>
                <select
                  onChange={(e) => handleStatusUpdate(order.ID, e.target.value)}
                  className="status-select"
                >
                  <option value="">Update Status</option>
                  <option value="Pending">Pending</option>
                  <option value="In Progress">In Progress</option>
                  <option value="Delivered">Delivered</option>
                </select>
                <button onClick={() => handleDeleteOrder(order.ID)} className="delete-button">
                  Delete
                </button>
                {/* Show "Assign New Courier" box only if status is "Pending" or "Awaiting Courier Acceptance" */}
                {['Pending', 'Awaiting Courier Acceptance'].includes(order.status) && (
                  <>
                    <input
                      type="text"
                      placeholder="New Courier ID"
                      value={courierIdMap[order.ID] || ''}
                      onChange={(e) => handleCourierIdChange(order.ID, e.target.value)}
                      className="courier-input"
                    />
                    {order.status === 'Pending' && (
                      <button
                        onClick={() => handleAssignNewCourier(order.ID)}
                        className="assign-button"
                      >
                        Assign New Courier
                      </button>
                    )}
                    {order.status === 'Awaiting Courier Acceptance' && (
                      <button
                        onClick={() => handleReassignNewCourier(order.ID)}
                        className="reassign-button"
                      >
                        Reassign Courier
                      </button>
                    )}
                  </>
                )}
              </td>
            </tr>
          ))}
        </tbody>
      </table>
    </div>
  );
};

export default ManageOrdersPage;
