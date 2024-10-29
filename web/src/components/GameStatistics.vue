<script setup lang="ts">
import shuttlecock from '@/assets/images/shuttlecock.png';
import type { GameStep, GetGame200Response } from "@/repositories/clients/public";
import { useGameStore } from '@/stores/game-store';
import { onBeforeMount, reactive, ref } from 'vue';

const gameData = reactive({} as GetGame200Response)
const errorMessage = ref("")
const noOfSteps = ref(0)
const gameStore = useGameStore()
gameStore.setBackendUrl(import.meta.env.VITE_BACKEND_URL)

onBeforeMount(async () => {
    await getStatistics()
})

const getStatistics = async () => {
    const res = await gameStore.getGameStatistics({
        gameId: gameStore.currentGameSettings.id
    })

    if (res instanceof Error) {
        errorMessage.value = res.message
        return
    }

    const data = res as GetGame200Response

    gameData.game = data.game
    gameData.steps = data.steps.slice(1)
    gameData.statistics = data.statistics
    noOfSteps.value = data.steps.length

    console.log(gameData);

}

const generateShuttleStyles = (index: number, current: GameStep) => {
    const greenStyles = {
        gridColumn: `${index + (index * 2)} / span 1`,
        gridRow: "1 / span 1",
    }

    const redStyles = {
        gridColumn: `${index + (index * 2)} / span 1`,
        gridRow: "3 / span 1",
        marginTop: "5px"
    }

    if (index == 0) {
        if (current.teamLeftScore > 0) {
            return greenStyles
        }

        return redStyles
    }

    const previous = gameData.steps[index - 1]

    if (current.teamLeftScore > previous.teamLeftScore) {
        return greenStyles
    }

    return redStyles
}

const generateMiddleStyles = (index: number) => {
    return {
        gridColumn: `${index + (index * 2)} / span 1`,
        gridRow: "2 / span 1",
    }
}

</script>

<template>
    <div class="container">
        <div class="actions-section">
        </div>
        <div class="points-section">
            <div class="player-circle-container">
                <div class="player-circle green"></div>
                <div class="player-circle red"></div>
            </div>
            <div class="points-display-container">
                <div v-for="(step, index) in gameData.steps" :key="step.id" :style="generateShuttleStyles(index, step)">
                    <img :src="shuttlecock" width="20px" height="20px">
                </div>
                <div v-for="(step, index) in gameData.steps" :key="`middle_line_${step.id}`"
                    :style="generateMiddleStyles(index)">
                    <div class="middle-line-background"></div>
                </div>
            </div>
        </div>
        <div class="actions-section">

        </div>
    </div>
</template>

<style>
.container {
    display: flex;
    flex-direction: column;
}

.actions-section {
    display: flex;
    flex-direction: column;
    height: 15vh;
    padding: 1rem;
}

.points-section {
    display: flex;
    height: 70vh;
    padding: 1rem;
    align-items: center;
}

.player-circle-container {
    display: flex;
    flex-direction: column;
}

.player-circle {
    height: 90px;
    width: 90px;
    border-radius: 45px;
    margin-right: 1rem;
}

.player-circle.green {
    background-color: green;
    margin-bottom: 1rem;
}

.player-circle.red {
    background-color: #D32F2F;
    margin-top: 1rem;
}

.points-display-container {
    display: grid;
    grid-template-columns: repeat(v-bind("noOfSteps"), 1fr 0.2fr);
    grid-template-rows: 2fr 0.2fr 2fr;
    width: 75vw;
    height: 50px;
}

.middle-line-background {
    height: 3px;
    background-color: darkgray;
    margin-left: 1px;
    margin-right: 1px;
}
</style>