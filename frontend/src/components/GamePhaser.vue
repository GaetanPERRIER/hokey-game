<template>
    <button @click="console.log(gameState)" class="temp">LOG GAMESTATE</button>
    <div id="phaser-game"></div>
</template>

<script setup>
import {onMounted, onUnmounted, ref} from "vue"
import Phaser from "phaser"


let game = null
let socket = null

// Références des entités
const status = ref("Disconnected")
const gameState = ref({})
let puck

// Map des sprites joueurs (créés uniquement côté front)
const playerSprites = {}

const gameSize = {
    width: 1280,
    height: 720
}

const config = {
    type: Phaser.AUTO,
    parent: "phaser-game",
    width: gameSize.width,
    height: gameSize.height,
    backgroundColor: "#fff", // couleur de fond Phaser (hex)
    physics: {
        default: "arcade",
        arcade: {
            debug: false
        }
    },
    scene: {
        preload() {
            this.load.image("puck", "imgs/puck.png")
        },
        create() {
            this.cameras.main.setBackgroundColor("cyan")

            // Utiliser des ellipses pour représenter les joueurs
            const w = 40 // largeur de l'ellipse
            const h = 40 // hauteur de l'ellipse

            // positions par défaut (si le serveur n'a rien encore)
            playerSprites['1'] = this.add.ellipse(160, 360, w, h, 0xffffff).setStrokeStyle(2, 0x000000)
            playerSprites['2'] = this.add.ellipse(1120, 360, w, h, 0xff0000).setStrokeStyle(2, 0x000000)
        },
        update() {
            playerSprites['1'].setPosition(gameState.value.Players?.[1]?.X || 160, gameState.value.Players?.[1]?.Y || 360)
            playerSprites['2'].setPosition(gameState.value.Players?.[2]?.X || 1120, gameState.value.Players?.[2]?.Y || 360)
        }
    }
}

onMounted(() => {
    socket = new WebSocket("ws://localhost:8080/ws")
    socket.onopen = () => {
        status.value = "Connected"
    }

    socket.onmessage = (event) => {
        gameState.value = JSON.parse(event.data)
    }

    // input state for sending continuous updates
    const inputState = { up: false, down: false, left: false, right: false }
    const validKeys = ["ArrowUp", "ArrowDown", "ArrowLeft", "ArrowRight"]

    const keyDownHandler = (event) => {
        if (status.value === "Disconnected") return
        if (!validKeys.includes(event.key)) return
        event.preventDefault()
        switch (event.key) {
            case "ArrowUp":
                inputState.up = true
                break
            case "ArrowDown":
                inputState.down = true
                break
            case "ArrowLeft":
                inputState.left = true
                break
            case "ArrowRight":
                inputState.right = true
                break
        }
    }

    const keyUpHandler = (event) => {
        if (status.value === "Disconnected") return
        if (!validKeys.includes(event.key)) return
        event.preventDefault()
        switch (event.key) {
            case "ArrowUp":
                inputState.up = false
                break
            case "ArrowDown":
                inputState.down = false
                break
            case "ArrowLeft":
                inputState.left = false
                break
            case "ArrowRight":
                inputState.right = false
                break
        }
    }

    window.addEventListener("keydown", keyDownHandler)
    window.addEventListener("keyup", keyUpHandler)

    // send input state regularly (20Hz)
    const sendIntervalMs = 50
    const intervalId = setInterval(() => {
        if (status.value === "Disconnected") return
        if (socket && socket.readyState === WebSocket.OPEN) {
            const msg = { type: "input_state", inputs: inputState }
            try {
                socket.send(JSON.stringify(msg))
            } catch (err) {
                // ignore send errors
            }
        }
    }, sendIntervalMs)

    game = new Phaser.Game(config)

    // cleanup on unmount
    onUnmounted(() => {
        clearInterval(intervalId)
        window.removeEventListener("keydown", keyDownHandler)
        window.removeEventListener("keyup", keyUpHandler)
    })
})

onUnmounted(() => {
    socket?.close()
    game?.destroy(true)
})
</script>

<style scoped>
#phaser-game {
    width: 100%;
    height: 100%;
    display: flex;
    justify-content: center;
    align-items: center;
}

.temp {
    position: absolute;
    top: 10px;
    left: 10px;
    z-index: 1000;
}
</style>
