import * as React from 'react';
import Card from '@mui/material/Card';
import CardContent from '@mui/material/CardContent';
import CardMedia from '@mui/material/CardMedia';
import Typography from '@mui/material/Typography';

export default function CharacterCard({name, image, status, bounty}) {
  return (
    <Card sx={{ maxWidth: 700 }}>
      <CardMedia
        component="img"
        image={image}
      />
      <CardContent>
        <Typography gutterBottom variant="h5" component="div">
          {name}
        </Typography>
        {status && <Typography variant="body2" color="text.secondary">
          Status: {status}          
        </Typography>}
        {bounty && <Typography variant="body2" color="text.secondary">
        Bounty: {bounty}       
        </Typography>}
      
      </CardContent>
      
    </Card>
  );
}
