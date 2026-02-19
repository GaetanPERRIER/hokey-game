import { ref, computed } from 'vue'
import { defineStore } from 'pinia'

export const useGameStore = defineStore('counter', () => {

    const puckPosition = ref({ x: 0, y: 0 })

    function updatePuckPosition(x: number, y: number) {
        puckPosition.value = { x, y }
    }


})
