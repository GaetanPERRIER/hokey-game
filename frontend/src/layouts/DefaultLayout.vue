<script setup>

import { io } from 'socket.io-client'
import {computed, onMounted, onUnmounted} from "vue";
import { useWebsocketStore} from "../../stores/websocket.js";

// Stores
const websocketStore = useWebsocketStore();


</script>

<template>
    <main>
        <slot />
    </main>
</template>

<style scoped lang="scss">

</style>