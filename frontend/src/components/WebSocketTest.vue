<template>
  <div>
    <h1>Hockey Game Test</h1>
    <p>Status: {{ status }}</p>
    <pre>{{ gameState }}</pre>
  </div>
</template>

<script setup>
import { ref, onMounted } from "vue"

const status = ref("Disconnected")
const gameState = ref({})

let socket

onMounted(() => {
  socket = new WebSocket("ws://localhost:8080/ws")

  socket.onopen = () => {
    status.value = "Connected"
  }

  socket.onmessage = (event) => {
    gameState.value = JSON.parse(event.data)
  }

  // Envoi des inputs clavier
  window.addEventListener("keydown", (e) => {
    const msg = { type: "input", key: e.key }
    socket.send(JSON.stringify(msg))
  })
})
</script>
