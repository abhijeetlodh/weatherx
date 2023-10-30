// client\src\components\Register.js
import React, {useState, useEffect} from 'react';
import axios from 'axios';
import User from './User';
import Cookies from 'js-cookie';
function Register() {
    const [email,
        setEmail] = useState('');
    const [password,
        setPassword] = useState('');
    const [firstName,
        setFirstName] = useState('');
    const [error,
        setError] = useState('');
    const [registered,
        setRegistered] = useState(false);
    const [userId,
        setUserId] = useState(null);
        useEffect(() => {
            Cookies.remove('auth_token');
          }, []);
    const handleSubmit = async(e) => {
        e.preventDefault();

        if (!email || !password || !firstName) {
            setError('Please fill in all fields');
            return;
        }

        const user = {
            email: email,
            password: password,
            firstName: firstName
        };

        const registerURL = `http://localhost:8080/register?email=${user.email}&password=${user.password}&firstName=${user.firstName}`;

        try {
            const response = await axios.post(registerURL, user);

            if (response.status === 201) {
                console.log('User registered successfully');
                setRegistered(true);
                setUserId(response.data.id);
            } else {
                console.error('Registration failed:', response.data.error);
                setError('Registration failed: User with this email already exists');
            }
        } catch (error) {
            console.error('Registration failed:', error);
            setError('Registration failed: An error occurred');
        }
    };

    return (
        <div className='RegisterForm'>
            {registered
                ? (
                    <div>
                        <h2>Registration Successful</h2>
                        <User firstName={firstName} email={email} userId={userId}/>
                    </div>
                )
                : (
                    <div>
                        <h2>Register</h2>
                        <form onSubmit={handleSubmit}>
                            <div>
                                <label>Email:</label>
                                <input type="email" value={email} onChange={(e) => setEmail(e.target.value)}/>
                            </div>
                            <div>
                                <label>Password:</label>
                                <input
                                    type="password"
                                    value={password}
                                    onChange={(e) => setPassword(e.target.value)}/>
                            </div>
                            <div>
                                <label>First Name:</label>
                                <input
                                    type="text"
                                    value={firstName}
                                    onChange={(e) => setFirstName(e.target.value)}/>
                            </div>
                            <button type="submit">Register</button>
                        </form>
                        {error && <p>{error}</p>}
                    </div>
                )}
        </div>
    );
}

export default Register;
