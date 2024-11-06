import React from 'react';

const Dashboard = () => {
  const email = localStorage.getItem('user');  // Retrieve the user data

  return (
    <div>
      <h1>Welcome, {email}</h1>
      <p>This is your dashboard. You are now logged in!</p>
    </div>
  );
};

export default Dashboard;
