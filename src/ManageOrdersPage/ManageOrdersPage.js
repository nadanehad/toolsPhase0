import React, { useEffect, useState } from 'react';
import { fetchAllOrdersAdmin, updateOrderAdmin, deleteOrderAdmin, reassignOrdersAdmin } from './api';

const ManageOrdersPage = () => {
  const [orders, setOrders] = useState([]);
  const [newCourierId, setNewCourierId] = useState('');

  useEffect(() => {
    fetchAllOrders();
  }, []);

  const fetchAllOrders = async () => {
    try {
      const data = await fetchAllOrdersAdmin();
      setOrders(data);
    } catch (error) {
      console.error('Error fetching all orders:', error);
    }
  };

  const handleStatusUpdate = async (orderId, status) => {
    try {
      await updateOrderAdmin(orderId, { status });
      fetchAllOrders(); // Refresh the list after updating
    } catch (error) {
      console.error('Error updating order status:', error);
    }
  };

  const handleDeleteOrder = async (orderId) => {
    try {
      await deleteOrderAdmin(orderId);
      fetchAllOrders(); // Refresh the list after deleting
    } catch (error) {
      console.error('Error deleting order:', error);
    }
  };

  const handleReassignOrders = async () => {
    try {
      await reassignOrdersAdmin(newCourierId);
      fetchAllOrders(); // Refresh the list after reassigning
    } catch (error) {
      console.error('Error reassigning orders:', error);
    }
  };

  return (
    <div>
      <h1>Manage Orders</h1>
      <table>
        <thead>
          <tr>
            <th>Order ID</th>
            <th>Status</th>
            <th>Courier ID</th>
            <th>Actions</th>
          </tr>
        </thead>
        <tbody>
          {orders.map(order => (
            <tr key={order.id}>
              <td>{order.id}</td>
              <td>{order.status}</td>
              <td>{order.courier_id || 'Unassigned'}</td>
              <td>
                <select onChange={(e) => handleStatusUpdate(order.id, e.target.value)}>
                  <option value="">Update Status</option>
                  <option value="Pending">Pending</option>
                  <option value="In Progress">In Progress</option>
                  <option value="Delivered">Delivered</option>
                </select>
                <button onClick={() => handleDeleteOrder(order.id)}>Delete</button>
              </td>
            </tr>
          ))}
        </tbody>
      </table>
      <div>
        <h2>Reassign Orders</h2>
        <input
          type="text"
          placeholder="New Courier ID"
          value={newCourierId}
          onChange={(e) => setNewCourierId(e.target.value)}
        />
        <button onClick={handleReassignOrders}>Reassign Orders</button>
      </div>
    </div>
  );
};

export default ManageOrdersPage;
