// controllers/player.controller.js
const Player = require('../../models/Player');
const socket = require("../../models/Player");

function handlePlayerConnection(socketId, io) {
    console.log(`Player connected: ${socketId}`);
}

function handlePlayerDisconnection(socketId, io) {
    console.log(`Player disconnected: ${socketId}`);
}

module.exports = { handlePlayerConnection, handlePlayerDisconnection };
