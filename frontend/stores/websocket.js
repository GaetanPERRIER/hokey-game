// stores/websocket.js
import { defineStore } from 'pinia'
import { io } from 'socket.io-client';

export const useWebsocketStore = defineStore('websocket' , {
    state: () => ({
        socket: null,
    }),

    getters: {},

    actions : {
        initSocket(token) {
            if (this.socket) return;

            this.socket = io("http://localhost:3000", {
                auth: { token }
            });

            this.socket.on('connect', () => {
                console.log('Socket connected:', this.socket.id);
            });
        }
    }
})