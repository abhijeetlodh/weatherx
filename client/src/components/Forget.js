import React, {useState} from 'react';
import axios from 'axios';

function Forget() {
    const [email,
        setEmail] = useState('');
    const [password,
        setPassword] = useState('');
    const [newPassword,
        setNewPassword] = useState('');
    const [message,
        setMessage] = useState('');
    const [error,
        setError] = useState('');

    const handleResetPassword = async() => {
        try {
            const response = await axios.put(`http://localhost:8080/update?email=${email}&password=${password}&new_password=${newPassword}`);
            if (response.status === 200) {
                setError('');
                setMessage('Password reset successful');
            } else {
                setError('Failed to reset password');
                setMessage('');
            }
        } catch (error) {
            setError('Failed to reset password');
            setMessage('');
        }
    };

    return (
        <div>
            <h3>Reset Password</h3>
            <form>
                <div>
                    <label>Email:</label>
                    <input
                        type="email"
                        value={email}
                        onChange={(e) => setEmail(e.target.value)}
                        required/>
                </div>
                <div>
                    <label>Current Password:</label>
                    <input
                        type="password"
                        value={password}
                        onChange={(e) => setPassword(e.target.value)}
                        required/>
                </div>
                <div>
                    <label>New Password:</label>
                    <input
                        type="password"
                        value={newPassword}
                        onChange={(e) => setNewPassword(e.target.value)}
                        required/>
                </div>
                <button type="button" onClick={handleResetPassword}>
                    Reset Password
                </button>
            </form>
            {message && <p style={{
                color: 'green'
            }}>{message}</p>}
            {error && <p style={{
                color: 'red'
            }}>{error}</p>}
        </div>
    );
}

export default Forget;
