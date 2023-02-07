import * as React from 'react';
import axios from "axios";
import Autocomplete from '@mui/material/Autocomplete';
import TextField from '@mui/material/TextField';


const HOST = process.env.REACT_APP_HOST;
const URL = `http://${HOST}:4000/autocomplete`;


export function Home() {
    const [words, setWords] = React.useState([]);
  
    return (
        <div>
        <Autocomplete
            onInputChange={async(event: any) => {
                if (event.target.value) {
                    const response = await axios.get(URL, {params: { word: event.target.value }});
                    console.log(response);
                    if (response.data.words) {
                        setWords(response.data.words);
                    }
                    console.log(words);
                } else {
                    setWords([]);
                }
            }}
            disablePortal
            id="combo-box-demo"
            options={words}
            sx={{ width: 300 }}
            renderInput={(params: any) => <TextField {...params} label="Type a word" />}
        />
        </div>
    );
  }