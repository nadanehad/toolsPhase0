// App.js
import React from 'react';
import { BrowserRouter as Router, Route, Routes } from 'react-router-dom';
import CreateOrder from './createOrder/CreateOrder'; 
import MyOrders from './myOrders/MyOrders'; 
import OrderDetails from './orderDetails/OrderDetails';
import AssignedOrdersPage from './AssignedOrdersPage/AssignedOrdersPage';
import ManageOrdersPage from './ManageOrdersPage/ManageOrdersPage';
import Home from './Userhome/UserHome'; 
import Login from './login/Login';      
import Register from './register/Register';
import './App.css'; 

function App() {
  return (
    <Router>
      <div className="App">
        <Routes>
          <Route path="/" element={<Login />} />
          <Route path="/register" element={<Register />} /> {/* Registration route */}
          <Route path="/home" element={<Home />} />
          <Route path="/create-order" element={<CreateOrder />} />
          <Route path="/my-orders" element={<MyOrders />} />
          <Route path="/order-details" element={<OrderDetails />} />
          <Route path="/assigned-orders" element={<AssignedOrdersPage />} /> 
          <Route path="/manage-orders" element={<ManageOrdersPage />} />
        </Routes>
      </div>
    </Router>
  );
}

export default App;
