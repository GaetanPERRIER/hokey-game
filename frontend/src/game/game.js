// javascript
import Phaser from 'phaser'
import {useGameStore} from "@/stores/game.ts";

class GameScene extends Phaser.Scene {
    constructor() {
        super({ key: 'GameScene' })
    }

    preload() {
        // charger assets ici

    }

    create() {
        this.add.text(100, 100, 'Hello World', { color: '#0f0' })
    }

    update() {
        // bouger le text


    }
}

export const config = {
    type: Phaser.AUTO,
    width: window.innerWidth,
    height: window.innerHeight,
    scale: {
        parent: 'phaser-game',
        mode: Phaser.Scale.RESIZE,
        autoCenter: Phaser.Scale.CENTER_BOTH
    },
    scene: GameScene,
}
