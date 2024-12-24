import axios from 'axios';

const API_URL = process.env.REACT_APP_API_URL ;
axios.defaults.withCredentials = true;

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
      withCredentials: true,
      headers: {
        'Content-Type': 'application/json', // Explicit headers
      },      
    });
    const { userID, role } = response.data;

    if (role) {
      localStorage.setItem('role', role);
    }

    if (role === 'Courier') {
      localStorage.setItem('courierID', userID);
    }

    return response.data;
  } catch (error) {
    throw error;
  }
};

export const createOrder = async (orderData) => {
  try {
    const response = await axios.post(`${API_URL}/create-order`, orderData, {
      withCredentials: true, 
    });
    return response.data;
  } catch (error) {
    throw error;
  }
};

export const fetchUserOrders = async () => {
  try {
    const response = await axios.get(`${API_URL}/orders`, {
      withCredentials: true,
    });
    return response.data;
  } catch (error) {
    throw error;
  }
};

// Fetch order details
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

// Cancel order
export const cancelOrder = async (orderId) => {
  try {
    const response = await axios.post(`${API_URL}/order/${orderId}/cancel`, {}, { withCredentials: true });
    return response.data;
  } catch (error) {
    throw error;
  }
};

// Courier Functions

export const fetchOrdersByCourier = async (courierId) => {
  try {
    const response = await axios.get(`${API_URL}/courier/${courierId}/orders`, {
      withCredentials: true,
    });
    return response.data;
  } catch (error) {
    throw error;
  }
};

// Accept or decline an order
export const acceptOrDeclineOrder = async (orderId, accept) => {
  try {
    const response = await axios.post(
      `${API_URL}/courier/order/${orderId}/accept`,
      { accept },
      { withCredentials: true }
    );
    return response.data;
  } catch (error) {
    throw error;
  }
};

// Update order status by courier
export const updateOrderStatusCourier = async (orderId, status) => {
  try {
    const response = await axios.put(
      `${API_URL}/courier/order/${orderId}/status`,
      { status },
      { withCredentials: true }
    );
    return response.data;
  } catch (error) {
    throw error;
  }
};

// Admin Functions

// Fetch all orders (admin)
export const fetchAllOrdersAdmin = async () => {
  try {
    const response = await axios.get(`${API_URL}/admin/orders`, {
      withCredentials: true,
    });
    return response.data;
  } catch (error) {
    throw error;
  }
};

// Assign an order to a courier (admin)
export const assignOrderToCourier = async (orderId, courierId) => {
  try {
    const payload = { order_id: orderId, courier_id: courierId };

    // Log the payload to ensure it's correct
    console.log('Sending assign order payload:', payload);

    const response = await axios.post(
      `${API_URL}/admin/assign-order`, // Correct API endpoint
      payload,  // Ensure you're sending the payload properly
      { withCredentials: true }
    );

    return response.data;
  } catch (error) {
    console.error('Error in API request:', error.response || error);  // Log any error
    throw error;
  }
};

// API call for reassigning courier
export const reassignOrderToCourier = async (orderId, courierId) => {
  try {
    const payload = { order_id: orderId, new_courier_id: courierId };  // Correct key

    console.log('Sending reassign payload:', payload);

    const response = await axios.post(
      `${API_URL}/admin/reassign-orders`, // Correct route for reassigning
      payload,  
      { withCredentials: true,
        headers: {
          'Content-Type': 'application/json',  // Ensure the correct content type
        },
       }
    );

    return response.data;
  } catch (error) {
    console.error('Error reassigning order:', error.response || error);
    throw error;
  }
};


// Update order by admin
export const updateOrderAdmin = async (orderId, orderData) => {
  try {
    const response = await axios.put(
      `${API_URL}/admin/order/${orderId}`,
      orderData,
      { withCredentials: true }
    );
    return response.data;
  } catch (error) {
    throw error;
  }
};

// Delete order (admin)
export const deleteOrderAdmin = async (orderId) => {
  try {
    const response = await axios.delete(`${API_URL}/admin/order/${orderId}`, {
      withCredentials: true,
    });
    return response.data;
  } catch (error) {
    throw error;
  }
};

// Fetch orders awaiting courier acceptance (admin)
export const fetchAwaitingCourierOrders = async () => {
  try {
    const response = await axios.get(`${API_URL}/admin/assigned-orders`, {
      withCredentials: true,
    });
    return response.data.orders;
  } catch (error) {
    throw error;
  }
};




