const express = require("express");
const axios = require("axios");
const jwt = require("jsonwebtoken");
const app = express();
const PORT = 4000;

app.use(express.json());

var token = "";

const refetchAccessToken = () => {
  var data =
    '{\r\n    "companyName": "sahiljena-travels.pvt.ltd",\r\n    "clientID": "32f015c8-fb45-41c9-8ecd-9aa1de12ea89",\r\n    "clientSecret": "MaKhenPFRwUDDMup"\r\n}';

  var config = {
    method: "post",
    url: "http://localhost:3000/auth",
    headers: {
      "Content-Type": "text/plain",
    },
    data: data,
  };
  axios(config)
    .then(function (response) {
      token = response.data.access_token;
      //console.log(token);
      return;
    })
    .catch(function (error) {
      console.log(error);
    });
};

const handleAuthExpiry = (req, res, next) => {
  if (token === "") {
    console.log("Setting Token");
    refetchAccessToken();
  } else {
    const decoded = jwt.decode(token);
    const expirationTime = decoded.exp;
    const currentTime = Math.floor(Date.now() / 1000);
    if (currentTime > expirationTime) {
      refetchAccessToken();
      console.log("Access Token Refereshed ");
    }
  }

  next();
};

app.post("/", (req, res) => {
  const { name } = req.body;
  res.status(201);
  res.json({ name: name });
});

app.get("/trains", handleAuthExpiry, (req, res) => {
  var config = {
    method: "get",
    url: "http://localhost:3000/trains",
    headers: {
      Authorization: "Bearer " + token,
    },
  };
  //console.log(config);
  axios(config)
    .then(function (response) {
      let trains = response.data;
      const sortedTrains = trains.sort((a, b) => {
        if (a.price.sleeper !== b.price.sleeper) {
          return a.price.sleeper - b.price.sleeper;
        }
        if (a.price.AC !== b.price.AC) {
          return a.price.AC - b.price.AC;
        }
        if (a.departureTime.Hours !== b.departureTime.Hours) {
          return a.departureTime.Hours - b.departureTime.Hours;
        }
        if (a.departureTime.Minutes !== b.departureTime.Minutes) {
          return a.departureTime.Minutes - b.departureTime.Minutes;
        }
        return a.delayedBy - b.delayedBy;
      });
      console.log(sortedTrains);
      res.status(200).json(sortedTrains);
    })
    .catch(function (error) {
      console.log(error);
      res.status(500).json({ error: "Token Expired Refersh again!!" });
    });
});

app.listen(PORT, (error) => {
  if (!error) {
    console.log("Server is running on port : " + PORT);
  } else {
    console.log("Error Occured");
  }
});
