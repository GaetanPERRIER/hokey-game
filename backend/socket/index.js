// backend/socket/index.js
const { Server } = require('socket.io');
const jwt = require("jsonwebtoken");

// Importer les controllers
const RoomController = require("../controllers/Room/room.controller");
const AuthController = require("../controllers/Player/player.controller");

function initSocket(httpServer) {

    // Initialiser Socket.IO avec les options CORS
    const io = new Server(httpServer, {
        cors: {
            origin: process.env.CORS_ORIGIN || '*', // Autoriser toutes les origines par défaut
            methods: ['GET', 'POST']
        }
    })

    io.use((socket, next) => {
        const token = socket.handshake.auth.token;

        if (!token) {
            return next(new Error("Authentification requise"));
        }

        try {
            socket.user = jwt.verify(token, process.env.JWT_SECRET); // on attache l'user à la socket
            next();
        } catch (err) {
            next(new Error("Token invalide"));
        }
    })

    // — Définir les événements Socket.IO —

    // Gérer les connexions/déconnexions Socket.IO
    io.on('connection', (socket) => {
        AuthController.handlePlayerConnection(socket.id, io);

        socket.on('disconnect', (reason) => {
            AuthController.handlePlayerDisconnection(socket.id, io);
        });

        // Gérer les événements de salle
        socket.on('room:join', () => {
            // Chercher une salle disponible
            RoomController.HandleJoinRoom(socket, io, { roomId: null });
        })

        // Récuperer les infos d'une salle
        socket.on('room:info', ({ roomId }) => {
            RoomController.HandleGetRoomInfo(socket, roomId);
        })
    })





    return io;
}

module.exports = initSocket;
