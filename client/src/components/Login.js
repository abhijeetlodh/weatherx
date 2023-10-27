import React, { useState } from 'react';
import axios from 'axios';
import User from './User';
import Forget from './Forget';
import { BrowserRouter as Router, Route, Link } from 'react-router-dom';

function Login() {
  const [email, setEmail] = useState('');
  const [password, setPassword] = useState('');
  const [loginError, setLoginError] = useState(null);
  const [loggedInEmail, setLoggedInEmail] = useState('');
  const [userId, setUserId] = useState(null);
  const [showLoginForm, setShowLoginForm] = useState(true); // State to control form visibility

  const handleSubmit = async (e) => {
    e.preventDefault();

    const userCredentials = {
      email: email,
      password: password,
    };
    const loginUrl = `http://localhost:8080/login?email=${userCredentials.email}&password=${userCredentials.password}`;
    try {
      const response = await axios.get(loginUrl);
      if (response.status === 200) {
        setUserId(response.data.id);
        setLoggedInEmail(email);
      } else {
        setLoginError('Please check email id or password');
      }
    } catch (error) {
      setLoginError('Login failed');
      console.error('Login failed', error);
    }
  };

  const handleForgetPasswordClick = () => {
    setShowLoginForm(false); // Hide the login form
  };

  return (
    <Router>
      <div className="LoginForm">
        {loggedInEmail ? (
          <User email={loggedInEmail} userId={userId} />
        ) : (
          <div>
            {showLoginForm ? ( // Conditionally render based on showLoginForm state
              <div>
                <h2>Login</h2>
                {loginError && <p>{loginError}</p>}
                <form onSubmit={handleSubmit}>
                  <div>
                    <label>Email:</label>
                    <input
                      type="email"
                      value={email}
                      onChange={(e) => setEmail(e.target.value)}
                      required
                    />
                  </div>
                  <div>
                    <label>Password:</label>
                    <input
                      type="password"
                      value={password}
                      onChange={(e) => setPassword(e.target.value)}
                      required
                    />
                  </div>
                  <button type="submit">Login</button>
                </form>
                <button onClick={handleForgetPasswordClick}>Update Password</button> {/* Updated text */}
              </div>
            ) : (
              <Forget />
            )}
          </div>
        )}
      </div>
    </Router>
  );
}

export default Login;
