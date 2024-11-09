import React, { useEffect, useState } from 'react';
import {
  fetchAllOrdersAdmin,
  updateOrderAdmin,
  deleteOrderAdmin,
  reassignOrdersAdmin,
} from '../api';
import './ManageOrdersPage.css'; 
const ManageOrdersPage = () => {
  const [orders, setOrders] = useState([]);
  const [newCourierId, setNewCourierId] = useState('');
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

  const handleReassignOrder = async (orderId) => {
    try {
      await reassignOrdersAdmin({ order_id: orderId, new_courier_id: newCourierId });
      fetchAllOrders(); 
    } catch (error) {
      console.error('Error reassigning order:', error);
    }
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
                <select onChange={(e) => handleStatusUpdate(order.ID, e.target.value)} className="status-select">
                  <option value="">Update Status</option>
                  <option value="Pending">Pending</option>
                  <option value="In Progress">In Progress</option>
                  <option value="Delivered">Delivered</option>
                </select>
                <button
                  onClick={() => handleDeleteOrder(order.ID)}
                  className="delete-button"
                >
                  Delete
                </button>
                {['Pending', 'Awaiting Courier Acceptance'].includes(order.status) && (
                  <>
                    <input
                      type="text"
                      placeholder="New Courier ID"
                      value={newCourierId}
                      onChange={(e) => setNewCourierId(e.target.value)}
                      className="courier-input"
                    />
                    <button
                      onClick={() => handleReassignOrder(order.ID)}
                      className="reassign-button"
                    >
                      Reassign Order
                    </button>
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
