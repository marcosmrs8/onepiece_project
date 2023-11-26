import React, { useState, useEffect } from "react";
import Navbar from "../components/Navbar";
import CharacterCard from "../components/CharactersCard";
import { Grid } from "@mui/material";
import { Container } from "@mui/system";
import axios from "axios";

export const Home = () => {
  const [characters, setCharacters] = useState([]);
  const [filteredCharacters, setFilteredCharacters] = React.useState([]);

  React.useEffect(() => {
    getCharacter();
  }, []);

  const getCharacter = () => {
    axios.get("http://localhost:8000/api/v1/personagens/").then((res) => {
      setCharacters(res.data);
      setFilteredCharacters(res.data);
    });
  };

  const filteredCharactersByName = (name) => {
    if (name === "") {
      setFilteredCharacters(characters);
    } else {
      const filtered = characters.filter((character) =>
        character.Name.includes(name)
      );
      setFilteredCharacters(filtered);
    }
  };

  return (
    <div>
      <Navbar filteredCharacters={filteredCharactersByName} />
      <Container>
        <Grid container>
          {filteredCharacters.map((data, key) => {
            const splitedImage = data.Image.split(".png");
            return (
              <Grid item xs={3} key={key} padding="10px">
                <CharacterCard
                  name={data.Name}
                  image={`${splitedImage[0]}.png`}
                  image_manga={`${splitedImage[1]}.png`}
                  status={data.Status}
                  bounty={data.Bounty}
                  appearence={data.Appearance}
                />
              </Grid>
            );
          })}
        </Grid>
      </Container>
    </div>
  );
};
