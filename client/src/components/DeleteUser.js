import React, {useState} from 'react';
import axios from 'axios';
import Cookies from 'js-cookie';

function DeleteUser({email}) {
    const [password,
        setPassword] = useState('');
    const [error,
        setError] = useState(null);
    const [isDeleted,
        setIsDeleted] = useState(false);

    const handleDelete = async() => {
        try {
            const response = await axios.delete(`http://localhost:8080/delete?email=${email}&password=${password}`);
            if (response.status === 200) {
                setIsDeleted(true);
                Cookies.remove('auth_token');
                window
                    .location
                    .replace('/login'); 
            } else {
                setError('Failed to delete the user. Please check your password.');
            }
        } catch (error) {
            setError('Failed to delete the user. Please try again later.');
        }
    };

    return (
        <div>
            {isDeleted
                ? (
                    <p>User {email}
                        has been successfully deleted.</p>
                )
                : (
                    <div>
                        <p>You are about to delete the user {email}.</p>
                        <p>Enter your password to confirm:</p>
                        <input
                            type="password"
                            value={password}
                            onChange={(e) => setPassword(e.target.value)}
                            required/> {error && <p style={{
                            color: 'red'
                        }}>{error}</p>}
                        <button onClick={handleDelete}>Delete User</button>
                    </div>
                )}
        </div>
    );
}

export default DeleteUser;
