<script setup lang="ts">
import userImage from '@/assets/images/user.png';
import type { GetGame200Response } from "@/repositories/clients/public";
import { useGameStore } from '@/stores/game-store';
import { computed, onBeforeMount, reactive, ref } from 'vue';

const gameData = reactive({} as GetGame200Response)
const errorMessage = ref("")
const winningTeam = ref("")
const firstLeftPlayerName = ref("")
const secondLeftPlayerName = ref("")
const firstRightPlayerName = ref("")
const secondRightPlayerName = ref("")
const matchDuration = ref("")
const gameStore = useGameStore()
gameStore.setBackendUrl(import.meta.env.VITE_PROXY_URL)

onBeforeMount(async () => {
    await getStatistics()
})

const updateDisplay = () => {

    if (gameData.steps) {
        const last = gameData.steps[gameData.steps.length - 1]
        if (last.teamLeftScore > last.teamRightScore) {
            winningTeam.value = "left"
        } else {
            winningTeam.value = "right"
        }
    }

    if (gameData.game) {
        firstLeftPlayerName.value = gameData.game.leftOddPlayerName
        secondLeftPlayerName.value = gameData.game.leftEvenPlayerName
        firstRightPlayerName.value = gameData.game.rightEvenPlayerName
        secondRightPlayerName.value = gameData.game.rightOddPlayerName
    }

    if (gameData.statistics) {
        matchDuration.value = gameData.statistics.totalGameTime
    }
}

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

    updateDisplay()
}

const leftPoint = computed(() => {
    if (gameData.steps && gameData.steps.length > 0) {
        const last = gameData.steps[gameData.steps.length - 1]
        return last.teamLeftScore
    }
    return ""
})

const rightPoint = computed(() => {
    if (gameData.steps && gameData.steps.length > 0) {
        const last = gameData.steps[gameData.steps.length - 1]
        return last.teamRightScore
    }
    return ""
})

const leftPointsClass = computed(() => ({
    isRed: winningTeam.value == "left",
    pointText: true
}))

const rightPointsClass = computed(() => ({
    isRed: winningTeam.value == "right",
    pointText: true
}))

</script>

<template>
    <div class="container">
        <div class="player-section">
            <div class="player-card">
                <div class="user-container user-mb" v-if="firstLeftPlayerName != ''">
                    <div class="profile-photo-container">
                        <img :src="userImage" alt="first left player image" width="30px" height="30px">
                    </div>
                    <span class="player-name">{{ firstLeftPlayerName }}</span>
                </div>

                <div class="user-container" v-if="secondLeftPlayerName != ''">
                    <div class="profile-photo-container">
                        <img :src="userImage" alt="second left player image" width="30px" height="30px">
                    </div>
                    <span class="player-name">{{ secondLeftPlayerName }}</span>
                </div>
            </div>
            <div class="points-section">
                <div class="points-box">
                    <span :class="leftPointsClass">{{ leftPoint }}</span>
                    <span class="point-seperator"> : </span>
                    <span :class="rightPointsClass">{{ rightPoint }}</span>
                </div>
                <div class="time-container">
                    <span>{{ matchDuration }}</span>
                </div>
            </div>
            <div class="player-card">
                <div class="user-container" v-if="firstRightPlayerName != ''">
                    <div class="profile-photo-container">
                        <img :src="userImage" alt="first right player image" width="30px" height="30px">
                    </div>
                    <span class="player-name">{{ firstRightPlayerName }}</span>
                </div>
                <div class="user-container user-mt" v-if="secondRightPlayerName != ''">
                    <div class="profile-photo-container">
                        <img :src="userImage" alt="second right player image" width="30px" height="30px">
                    </div>
                    <span class="player-name">{{ secondRightPlayerName }}</span>
                </div>
            </div>
        </div>
    </div>
</template>

<style>
.container {
    display: flex;
    flex-direction: column;
    background-color: rgb(191, 243, 191);
}

.player-section {
    display: flex;
    justify-content: center;
    width: 100vw;
    height: 30vh;
}

.player-card {
    display: flex;
    flex-direction: column;
    flex-grow: 1;
    justify-content: center;
    align-items: center;
    width: 40vw;
    padding: 0 20px;
}

.time-container {
    font-weight: bold;
    font-size: 1rem
}

.points-section {
    display: flex;
    flex-direction: column;
    width: 20vw;
    align-items: center;
    justify-content: center;

}

.points-box {
    display: flex;
    font-weight: bold;
    align-items: center;
    justify-content: center;
}

.user-container {
    display: flex;
    flex-direction: column;
    align-items: center;
    max-height: 15vh;
}

.profile-photo-container {
    border: 1px solid black;
    border-radius: 30px;
    padding: 10px;
}

.isRed {
    color: red;
}

.pointText {
    font-size: 2rem;
    margin: 0 5px;
}

.point-seperator {
    font-size: 1.5rem;
}

.user-mt {
    margin-top: 15px;
}

.user-mb {
    margin-bottom: 15px;
}

.player-name {
    font-weight: bold;
    text-wrap: wrap;
}
</style>