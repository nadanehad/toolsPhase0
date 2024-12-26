import React, { useState } from 'react';
import { createOrder } from '../api'; 
import './createOrder.css'; 

const CreateOrder = () => {
  const [formData, setFormData] = useState({
    pickup_location: '',  
    dropoff_location: '',
    package_details: '',
    delivery_time: '', 
    delivery_fee: 0 
  });
  const [message, setMessage] = useState('');
  const [error, setError] = useState('');

  const handleChange = (e) => {
    const { name, value, type } = e.target;
    setFormData((prevData) => ({
      ...prevData,
      [name]: type === "number" ? parseFloat(value) || '' : value 
    }));
  };

  const handleSubmit = async (e) => {
    e.preventDefault();
    setError('');
    setMessage('');
    const formattedData = {
      ...formData,
      delivery_time: new Date(formData.delivery_time).toISOString() // Convert to ISO 8601
    };

    console.log("Submitting order with data:", formData); 

    try {
      const response = await createOrder(formData);
      console.log("Order created response:", response); 
      setMessage('Order created successfully!');
      setFormData({ pickup_location: '', dropoff_location: '', package_details: '', delivery_time: '', delivery_fee: 0 });
    } catch (error) {
      setError('Failed to create order. Please try again.');
      console.error("Error creating order:", error); 
    }
  };

  return (
    <div className="create-order-container">
      <h2>Create Order</h2>
      {message && <p style={{ color: 'green' }}>{message}</p>}
      {error && <p style={{ color: 'red' }}>{error}</p>}
      <form onSubmit={handleSubmit}>
        <div>
          <label>Pickup Location</label>
          <input
            type="text"
            name="pickup_location"
            value={formData.pickup_location}
            onChange={handleChange}
            required
          />
        </div>
        <div>
          <label>Dropoff Location</label>
          <input
            type="text"
            name="dropoff_location"
            value={formData.dropoff_location}
            onChange={handleChange}
            required
          />
        </div>
        <div>
          <label>Package Details</label>
          <textarea
            name="package_details"
            value={formData.package_details}
            onChange={handleChange}
            required
          />
        </div>
        <div>
          <label>Delivery Time</label>
          <input
            type="datetime-local"
            name="delivery_time"
            value={formData.delivery_time}
            onChange={handleChange}
            required
          />
        </div>
        <div>
          <label>Delivery Fee</label>
          <input
            type="number"
            name="delivery_fee"
            value={formData.delivery_fee}
            onChange={handleChange}
            required
          />
        </div>
        <button type="submit">Create Order</button>
      </form>
    </div>
  );
};

export default CreateOrder;
