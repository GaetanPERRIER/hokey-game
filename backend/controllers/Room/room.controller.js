// controllers/room.controller.js
const RoomService = require("../../services/Room/room.service");

// Gérer la création d'une salle
function HandleJoinRoom(socket, io, { roomId }) {
    if (!roomId) {
        const availableRooms = RoomService.ListAvailableRooms();
        if (availableRooms.length > 0) {
            roomId = availableRooms[0]; // Rejoindre la première salle disponible
        } else {
            roomId = `room-${Date.now()}`; // Créer une nouvelle salle avec un ID unique
        }
    }
    RoomService.JoinRoom(socket, roomId, io);
    socket.emit("room:joined", { roomId });
}

// Gérer la sortie d'une salle
function HandleLeaveRoom(socket, io, { roomId }) {
    if (!roomId) {
        socket.emit("room:error", { message: "Room ID is required" });
        return;
    }

    RoomService.LeaveRoom(socket, roomId, io);
    socket.emit("room:left", { roomId });
}

// Gérer la déconnection d'un joueur
function HandlePlayerDisconnect() {
    // Créer une méthode dans le service qui retire un joueur de toutes les salles auxquelles il appartient
}

function HandleGetRoomInfo(socket, roomId) {
    const room = RoomService.GetRoom(roomId);
    if (!room) {
        return { error: "Room not found" };
    }

    socket.emit("room:updated", { room });
}

module.exports = { HandleJoinRoom, HandleLeaveRoom, HandlePlayerDisconnect, HandleGetRoomInfo };