import React, { useState, useEffect } from 'react';
import './dark-mode.css';
import FizzBuzzForm from "./components/FizzBuzzForm";
import Statistics from "./components/Statistics";

const App = () => {
    const [isDarkMode, setIsDarkMode] = useState(false);

    useEffect(() => {
        const savedDarkMode = localStorage.getItem('darkMode');
        setIsDarkMode(savedDarkMode === 'true');
    }, []);

    const toggleDarkMode = () => {
        const newDarkMode = !isDarkMode;
        setIsDarkMode(newDarkMode);

        localStorage.setItem('darkMode', newDarkMode);
    };

    return (
        <div className={`app ${isDarkMode ? 'dark-mode' : ''}`}>
            <h1>Mon application</h1>
            <button onClick={toggleDarkMode}>Activer/Désactiver le mode sombre</button>
            <FizzBuzzForm />
            <Statistics />
        </div>
    );
};

export default App;
