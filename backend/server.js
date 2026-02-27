// Importer les modules nécessaires
require('dotenv').config();
const express = require('express');
const cors = require('cors');

const connectDB = require("./config/db")
const authRoutes = require("./routes/auth.routes")

const http = require('http');
const config = require('./config/default');

const initRoles = require("./config/initRoles")
const { handlePlayerConnection, handlePlayerDisconnection } = require('./controllers/Player/player.controller');
const initSocket = require('./socket');

// Créer une instance d'Express
const app = express();

// Connecter à la base de données et initialiser les rôles
connectDB().then(() => {
    initRoles()
})

// Healthcheck simple
app.get('/health', (req, res) => {
    res.json({ status: 'ok', uptime: process.uptime() });
});

// Créer un serveur HTTP
const server = http.createServer(app);
const io = initSocket(server)

// Démarrer le serveur
app.use(cors())
app.use(express.json())
app.use("/api/auth", authRoutes)

const PORT = process.env.PORT || 3000;
server.listen(PORT, () => {
    console.log(`Server listening on port ${PORT}`);
});
