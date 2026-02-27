// services/room.service.js
const rooms = new Map() // roomId -> Set(socket.id)

// Créer une nouvelle salle ou récupérer une salle existante
function CreateRoom(roomId, meta = {}) {
    if (!rooms.has(roomId)) {
        rooms.set(roomId, {players: new Set, meta});
        console.log(`Room created: ${roomId}`);
    }
    return rooms.get(roomId)
}

// Rejoindre une salle
function JoinRoom(socket, roomId, io) {
    const room = CreateRoom(roomId);
    room.players.add(socket.id);
    socket.join(roomId);
    io.to(roomId).emit('room:updated', {
        // convertir le Set en tableau pour être sérialisable
        players: Array.from(room.players),
        meta: room.meta
    });    return room;
}

// Quitter une salle
function LeaveRoom(socket, roomId, io) {
    const room = rooms.get(roomId);
    if (!room)
        return;

    room.players.delete(socket.id);
    socket.leave(roomId);
    io.to(roomId).emit('room:updated', { room });
    if (room.players.size === 0) {
        rooms.delete(roomId);
    }
}

// Lister toutes les salles
function ListRooms() {
    return Array.from(rooms.keys());
}

// Lister les salles disponibles (Une salle contient 2 joueurs max)
function ListAvailableRooms() {
    return Array.from(rooms.entries())
        .filter(([_, room]) => room.players.size < 2)
        .map(([roomId, _]) => roomId);
}

function GetRoom(roomId) {
    return rooms.get(roomId);
}

module.exports = { CreateRoom, JoinRoom, LeaveRoom, ListRooms, ListAvailableRooms, GetRoom };

