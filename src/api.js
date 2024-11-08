import axios from 'axios';

// Define the API URL
const API_URL = process.env.REACT_APP_API_URL || 'http://localhost:8080';

// Function to register a user
export const registerUser = async (formData) => {
  try {
    const response = await axios.post(`${API_URL}/register`, formData);
    return response.data;
  } catch (error) {
    throw error;
  }
};

export const loginUser = async (credentials) => {
  try {
    const response = await axios.post(`${API_URL}/login`, credentials, {
      withCredentials: true, // Make sure credentials are included
    });
    return response.data;
  } catch (error) {
    throw error;
  }
};

export const createOrder = async (orderData) => {
  try {
    const response = await axios.post(`${API_URL}/create-order`, orderData, {
      withCredentials: true, // Ensures the cookie containing the session ID is sent along with the request
    });
    return response.data;
  } catch (error) {
    throw error; // Throwing the error so it can be caught in the component
  }
};

// api.js
export const fetchUserOrders = async () => {
  try {
    const response = await axios.get(`${API_URL}/orders`, {
      withCredentials: true,
    });
    console.log("Full API response:", response); // Log entire response
    console.log("Response data structure:", response.data); // Log just the data part
    return response.data;
  } catch (error) {
    console.error("Error fetching user orders:", error);
    throw error;
  }
};

export const fetchOrderDetails = async (orderId) => {
  try {
    const response = await axios.get(`${API_URL}/order/${orderId}`, {
      withCredentials: true,
    });
    return response.data;
  } catch (error) {
    throw error;
  }
};

export const cancelOrder = async (orderId) => {
  try {
    const response = await axios.post(`${API_URL}/order/${orderId}/cancel`, {}, { withCredentials: true });
    return response.data;
  } catch (error) {
    throw error;
  }
};
