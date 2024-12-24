import React from 'react';
import ReactDOM from 'react-dom/client';
import './index.css';            // Global styles
import App from './App';     // Ensure the path is correct to App.js
import reportWebVitals from './reportWebVitals';

const root = ReactDOM.createRoot(document.getElementById('root'));
root.render(
  <React.StrictMode>
    <App />
  </React.StrictMode>
);

reportWebVitals();
