import React, {useState, useEffect} from "react";
import Navbar from "../components/Navbar";
import CharacterCard from "../components/CharactersCard";
import { Grid } from "@mui/material";
import { Container } from "@mui/system";
import axios from "axios";
export const Home = () => {
    const [characters, setCharacters] = useState([])
    useEffect(() => {
        getCharacter()
        
    },[])

    const getCharacter = () => {        
        axios.get("http://localhost:8000/api/v1/personagens/")
        .then((res) => setCharacters(res.data))
    }
    const filteredCharacters = (name) => {
        const filtered = []
        if(name === "" ){
            getCharacter()
            window.location.reload();
        }
        for(const i in characters) {
            if(characters[i].Name && characters[i].Name.includes(name)){
                filtered.push(characters[i])
            }
        }
        setCharacters(filtered)
    }
    return (
        <div>
           <Navbar filteredCharacters={filteredCharacters}/>
           <Container>    
            <Grid container>
                {characters.map((data, key) =>{
                const splitedImage = data.Image.split(".png")
                    return (<Grid item xs={3} key={key} padding="10px">
                        <CharacterCard name={data.Name} image={`${splitedImage[0]}.png`} status={data.Status} bounty={data.Bounty}/>
                    </Grid>)
                })}
            </Grid>
           </Container>
        </div>
    )
}