import * as React from "react";
import Card from "@mui/material/Card";
import CardContent from "@mui/material/CardContent";
import CardMedia from "@mui/material/CardMedia";
import Typography from "@mui/material/Typography";

export default function CharacterCard({
  name,
  image,
  image_manga,
  status,
  bounty,
  appearence,
}) {
  const [isCardOpen, setIsCardOpen] = React.useState(false);

  const handleCardClick = () => {
    setIsCardOpen(!isCardOpen);
  };

  const handleCloseClick = () => {
    setIsCardOpen(false);
  };

  const imageStyle = {
    backgroundSize: "cover",
    width: "300px",
    height: "300px",
    objectFit: "cover",
    objectPosition: "top",
    transition: "transform 0.3s",
  };

  const cardStyle = {
    maxWidth: 700,
    position: "relative",
    cursor: "pointer",
    transition: "z-index 0.3s",
    zIndex: isCardOpen ? 1 : "auto",
    transform: isCardOpen ? "scale(1.2)" : "scale(1)",
    animation: isCardOpen ? "fade-in 0.3s ease-in-out" : "none",
    display: "flex",
    flexDirection: "column",
    alignItems: "center",
  };

  const cardContentStyle = {
    textAlign: "center",
    flex: "1",
    display: "flex",
    flexDirection: "column",
    justifyContent: "center",
  };

  const overlayStyle = {
    position: "fixed",
    top: 0,
    left: 0,
    width: "100%",
    height: "100%",
    backgroundColor: "rgba(0, 0, 0, 0.7)",
    display: isCardOpen ? "block" : "none",
    transition: "opacity 0.3s",
    opacity: isCardOpen ? 1 : 0,
    zIndex: 2,
    cursor: "pointer",
  };

  const modalStyle = {
    position: "fixed",
    top: "50%",
    left: "50%",
    transform: isCardOpen
      ? "translate(-50%, -50%) scale(1.2)"
      : "translate(-50%, -50%) scale(1)",
    zIndex: 3,
    display: isCardOpen ? "block" : "none",
    transition: "transform 0.3s",
  };

  return (
    <div>
      <style>
        {`
          @keyframes fade-in {
            from {
              opacity: 0;
            }
            to {
              opacity: 1;
            }
          }
        `}
      </style>

      <Card sx={cardStyle} onClick={handleCardClick}>
        <CardMedia
          component="img"
          image={image ? image : image_manga}
          style={imageStyle}
        />
        <CardContent sx={cardContentStyle}>
          <Typography gutterBottom variant="h5" component="div">
            {name}
          </Typography>
          {status && (
            <Typography variant="body2" color="text.secondary">
              {status}
            </Typography>
          )}
          {bounty && (
            <Typography variant="body2" color="text.secondary">
              Bounty: {bounty}
            </Typography>
          )}
        </CardContent>
      </Card>

      <div style={overlayStyle} onClick={handleCloseClick}></div>

      <div style={modalStyle}>
        <Card sx={{ maxWidth: 700 }}>
          <CardMedia
            component="img"
            image={image ? image : image_manga}
            style={imageStyle}
          />
          <CardContent sx={cardContentStyle}>
            <Typography gutterBottom variant="h5" component="div">
              {name}
            </Typography>
            {status && (
              <Typography variant="body2" color="text.secondary">
                {status}
              </Typography>
            )}
            {bounty && (
              <Typography variant="body2" color="text.secondary">
                Bounty: {bounty}
              </Typography>
            )}
          </CardContent>
          {isCardOpen && (
            <CardContent sx={cardContentStyle}>
              <strong>Aparência: </strong> {appearence}
              <Typography variant="body2" color="text.secondary">
                Additional information about {name}
              </Typography>
            </CardContent>
          )}
        </Card>
      </div>
    </div>
  );
}
