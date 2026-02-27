<script setup>

import { useWebsocketStore} from "../../stores/websocket.js";
import { useRoomStore} from "../../stores/room.js";
import { computed, onMounted} from "vue";

// Stores
const websocketStore = useWebsocketStore();
const roomStore = useRoomStore();

// Variables calculées
const socket = computed(() => websocketStore.socket);
const room = computed(() => roomStore.room);

onMounted(() => {
    // Écouter les événements de socket.io

    socket.value.on('room:updated', (data) => {
        console.log('Mise à jour de la salle :', data);
        roomStore.setRoom(data);

    });

    // Tenter de rejoindre une salle de jeu
    socket.value.emit('room:join');
})

</script>

<template>
    <div class="game">
        <h1>Game</h1>
            <div v-if="room">
                <p>{{ room.players.size }}</p>
            </div>
    </div>
</template>

<style scoped lang="scss">

</style>