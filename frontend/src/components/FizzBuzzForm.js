import React , {useState} from "react";
import axios from 'axios' ;



const FizzBuzzForm = () => {
    const [int1 , setInt1] = useState('');
    const [int2 , setInt2] = useState('');
    const [limit , setLimit] = useState('');
    const [str1 , setStr1] = useState('');
    const [str2 , setStr2] = useState('');
    const [fizzBuzzResult, setFizzBuzzResult] = useState('')

    const handleFizzBuzzSubmit = (e) => {
        e.preventDefault();

        axios.get(`http://localhost:8000/fizzbuzz?int1=${int1}&int2=${int2}&limit=${limit}&str1=${str1}&str2=${str2}`)
            .then((response) => {
                const data = response.data;
                //mise a jour du resultat de l'endpoint fizzbuzz avec la reponse de l'Api
                setFizzBuzzResult(data.result.join(', '));
            })
            .catch((error) => {
                console.error('Erreur de lors de la requete Fizzbuzz:', error);
            });
    };

    return (
        <div>
           <h2>Requete FizzBuzz</h2>
            <form onSubmit={handleFizzBuzzSubmit}>
                <div>
                    <label>int1 :</label>
                    <input type="number" value={int1} onChange={(e) => setInt1(e.target.value)} />
                </div>
                <div>
                    <label>int2 :</label>
                    <input type="number" value={int2} onChange={(e) => setInt2(e.target.value)} />
                </div>
                <div>
                    <label>Limit :</label>
                    <input type="number" value={limit} onChange={(e) => setLimit(e.target.value)} />
                </div>
                <div>
                    <label>str1 :</label>
                    <input type="text" value={str1} onChange={(e) => setStr1(e.target.value)} />
                </div>
                <div>
                    <label>str2 :</label>
                    <input type="text" value={str2} onChange={(e) => setStr2(e.target.value)} />
                </div>
                <button type="submit">Submit</button>
            </form>


            {fizzBuzzResult && (
                <div>
                    <h2>Resultat FizzBuzz :</h2>
                    <p>{fizzBuzzResult}</p>
                </div>
            )}
        </div>
    );
};

export default FizzBuzzForm;