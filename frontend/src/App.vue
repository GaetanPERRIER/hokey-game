<script setup>
import DefaultLayout from "@/layouts/DefaultLayout.vue";
import { useWebsocketStore} from "../stores/websocket.js";

const websocketStore = useWebsocketStore();

// Récupération du token depuis le localStorage
const token = localStorage.getItem("token");

if (token) {
    websocketStore.initSocket(token);
}

</script>

<template>
    <div id="app">
        <DefaultLayout>
            <router-view />
        </DefaultLayout>
    </div>
</template>

<style>

#app { font-family: Avenir, Helvetica, Arial, sans-serif; }

</style>

