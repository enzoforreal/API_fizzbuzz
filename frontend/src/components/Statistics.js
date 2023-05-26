import React, { useEffect, useState } from 'react';
import axios from 'axios';
import { Bar } from 'react-chartjs-2';


const Statistics = () => {
    const [mostUsedRequest, setMostUsedRequest] = useState({});
    const [requestCounts, setRequestCounts] = useState([]);

    useEffect(() => {
        axios
            .get('http://localhost:8000/statistics')
            .then((response) => {
                const data = response.data;
                setMostUsedRequest(data.MostUsedRequest); // Utiliser data.MostUsedRequest au lieu de data.mostUsedRequest
                setRequestCounts(data.RequestCount); // Utiliser data.RequestCount au lieu de data.requestCount
            })
            .catch((error) => {
                console.error('Erreur lors de la requête de statistiques :', error);
            });
    }, []);

    const getRequestLabels = () => {
        return requestCounts?.map((pair, index) => `Requête ${index + 1}`);
    };

    const getRequestCounts = () => {
        return requestCounts?.map((pair) => pair?.count);
    };

    return (
        <div>
            <h2>Statistiques</h2>
            {Object.keys(mostUsedRequest).length > 0 && (
                <div>
                    <h3>Requête la plus utilisée :</h3>
                    <p>int1: {mostUsedRequest.int1}</p>
                    <p>int2: {mostUsedRequest.int2}</p>
                    <p>limit: {mostUsedRequest.limit}</p>
                    <p>str1: {mostUsedRequest.str1}</p>
                    <p>str2: {mostUsedRequest.str2}</p>
                </div>
            )}

            {requestCounts.length > 0 && (
                <div>
                    <h3>Décompte des requêtes :</h3>
                    <Bar
                        data={{
                            labels: getRequestLabels(),
                            datasets: [
                                {
                                    label: 'Count',
                                    data: getRequestCounts(),
                                    backgroundColor: 'rgba(75, 192, 192, 0.2)',
                                    borderColor: 'rgba(75, 192, 192, 1)',
                                    borderWidth: 1,
                                },
                            ],
                        }}
                        options={{
                            scales: {
                                y: {
                                    beginAtZero: true,
                                    stepSize: 1,
                                },
                            },
                        }}
                    />
                </div>
            )}
        </div>
    );
};

export default Statistics;
