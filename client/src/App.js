import React, {useState, useEffect} from 'react';
import './App.css';
import Register from './components/Register';
import Login from './components/Login';
import User from './components/User';
import Cookies from 'js-cookie';

function App() {
    const [isRegistering,
        setIsRegistering] = useState(true);
    const [isLoggedIn,
        setIsLoggedIn] = useState(false);
    const [userId,
        setUserId] = useState('');

    const toggleForm = () => {
        setIsRegistering(!isRegistering);
    };

    const handleLogin = (userId) => {
        setIsLoggedIn(true);
        setUserId(userId);
    };
    
    useEffect(() => {
        const authToken = Cookies.get('auth_token');
        if (authToken) {
          setIsLoggedIn(true);
        }
      }, []);

    const handleLogout = () => {
        setIsLoggedIn(false);
        setUserId('');
    };
    
    return (
        <div className="App">
            <header className="App-header">
                {isLoggedIn
                    ? (<User userId={userId}/>)
                    : (
                        <div>
                            {userId && <p>User ID: {userId}</p>}
                            {/* Display user ID if available */}
                            {isRegistering
                                ? <Register onLogin={handleLogin}/>
                                : <Login onLogin={handleLogin}/>}
                            {!isLoggedIn && (
                                <button onClick={toggleForm}>
                                    {isRegistering
                                        ? 'Already have an account? Login'
                                        : "Don't have an account? Register"}
                                </button>
                            )}
                        </div>
                    )}
                {isLoggedIn && (
                    <button onClick={handleLogout}>Logout</button>
                )}
            </header>
        </div>
    );
}

export default App;
