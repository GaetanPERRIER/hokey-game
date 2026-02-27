// stores/websocket.js
import { defineStore } from 'pinia'
import { io } from 'socket.io-client';

export const useRoomStore = defineStore('room' , {
    state: () => ({
        room: null,
    }),

    getters: {},

    actions : {
        setRoom(room) {
            this.room = room;
        }
    }
})