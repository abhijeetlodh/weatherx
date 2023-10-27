// client\src\components\User.js
import React, {useState} from 'react';
import axios from 'axios';
import Cookies from 'js-cookie';
import {useEffect} from 'react';
import {DateTime} from 'luxon';
import DeleteUser from './DeleteUser';
import {Link} from 'react-router-dom';

function User({email, firstName, userId}) {
    const [cityName,
        setCityName] = useState('');
    const [metric,
        setMetric] = useState('K');
    const [weatherData,
        setWeatherData] = useState(null);
    const [error,
        setError] = useState('');
    const [weatherHistory,
        setWeatherHistory] = useState(null);
    const [isDeleteUserClicked,
        setIsDeleteUserClicked] = useState(false);
    const fetchWeatherHistory = async(userId) => {
        return new Promise(async(resolve, reject) => {
            try {
                const url = `http://localhost:8080/fetchhistorytable?user_id=${userId}`;
                const response = await axios.get(url);

                if (response.status === 200) {
                    resolve(response.data);
                } else {
                    reject('Failed to retrieve weather history');
                }
            } catch (error) {
                reject('Failed to retrieve weather history');
            }
        });
    };

    const saveSearchData = async(userId, temperature, cityName, metric) => {
        try {
            console.log('userId:', userId);
            const response = await axios.post('http://localhost:8080/saveSearchData', {
                "user_id": userId,
                "temp_f": temperature.toString(),
                "location": cityName,
                "metric": metric
            });

            if (response.status === 200) {
                console.log('Search data saved successfully');
            } else {
                console.error('Failed to save search data:', response.data.error);
            }
        } catch (error) {
            console.error('Failed to save search data:', error);
        }
    };

    const handleSubmit = async(e) => {
        e.preventDefault();

        if (!['K', 'C', 'F'].includes(metric)) {
            setError('Invalid metric value. Please use K, C, or F.');
            return;
        }
        setError('');
        try {
            const response = await axios.get(`http://localhost:8080/weather?cityName=${cityName}&metric=${metric}`);
            if (response.status === 200) {
                setWeatherData(response.data);
                saveSearchData(userId, response.data.Temperature, cityName, metric);
            } else {
                setError('Weather request failed. Please check your inputs.');
            }
        } catch (error) {
            setError('');
        }
        fetchWeatherHistory(userId).then((data) => {
            setWeatherHistory(data);
        }).catch((error) => {
            console.error(error);
        });
    };

    const handleLogout = async() => {
        try {
            const response = await axios.post('http://localhost:8080/logout');
            if (response.status === 200) {
                console.log('Logout successful');
                const authToken = Cookies.get('auth_token');
                console.log('Cookie:', authToken);
                Cookies.remove('auth_token');
                window
                    .location
                    .replace('/login'); 
            }
        } catch (error) {
            console.error('Logout request failed:', error);
        }
    };

    const handleDeleteUserClick = () => {
        setIsDeleteUserClicked(true);
    };

    return (
        <div>
            {isDeleteUserClicked
                ? (<DeleteUser email={email}/>)
                : (
                    <div>
                        <h2>Welcome, {firstName
                                ? firstName
                                : email}</h2>
                        <p>User ID: {userId}</p>
                        <button onClick={handleLogout}>Logout</button>
                        <h3>Weather Information</h3>
                        <form onSubmit={handleSubmit}>
                            <div>
                                <label>City Name:</label>
                                <input
                                    type="text"
                                    value={cityName}
                                    onChange={(e) => setCityName(e.target.value)}
                                    required/>
                            </div>
                            <div>
                                <label>Metric:</label>
                                <select value={metric} onChange={(e) => setMetric(e.target.value)} required>
                                    <option value="K">Kelvin (K)</option>
                                    <option value="C">Celsius (C)</option>
                                    <option value="F">Fahrenheit (F)</option>
                                </select>
                            </div>
                            <button type="submit">Get Weather</button>
                        </form>

                        {error && <p style={{
                            color: 'red'
                        }}>{error}</p>}

                        {weatherData && (
                            <div>
                                <div>
                                <h4>Weather Information for {cityName}</h4>
                                  <table id='result-table'>
                                    <tbody>
                                      <tr>
                                        <th>Metric</th>
                                        <th>Temperature</th>
                                      </tr>
                                      <tr>
                                        <td>{metric}</td>
                                        <td>{weatherData.Temperature}</td>
                                      </tr>
                                    </tbody>
                                  </table>
                                </div>
                                <h3>Weather History</h3>
                                {weatherHistory
                                    ? (
                                        <table>
                                            <thead>
                                                <tr>
                                                    <th>Number</th>
                                                    <th>User ID</th>
                                                    <th>City Name</th>
                                                    <th>Temperature</th>
                                                    <th>Time</th>
                                                </tr>
                                            </thead>
                                            <tbody>
                                                {weatherHistory.map((data, index) => (
                                                    <tr key={index}>
                                                        <td>{index + 1}</td>
                                                        <td>{userId}</td>
                                                        <td>{data.CityName}</td>
                                                        <td>{`${data.Temperature}${data['Metric']}`}</td>
                                                        <td>{data['Time(ist)']}</td>
                                                        {console.log(data['Metric'])}
                                                    </tr>
                                                ))}
                                            </tbody>
                                        </table>
                                    )
                                    : (
                                        <p>Loading weather history...</p>
                                    )}
                            </div>
                        )}
                        <button onClick={handleDeleteUserClick}>Delete Account</button>
                    </div>
                )}
        </div>
    );

}

export default User;
