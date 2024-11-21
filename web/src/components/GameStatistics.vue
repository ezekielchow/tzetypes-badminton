<script setup lang="ts">
import userImage from '@/assets/images/user.png';
import { type GetGame200Response } from "@/repositories/clients/public";
import { useGameStore } from '@/stores/game-store';
import { computed, onBeforeMount, reactive, ref } from 'vue';
import { useRoute } from 'vue-router';

const route = useRoute()
const gameData = reactive({} as GetGame200Response)
const errorMessage = ref("")
const winningTeam = ref("")
const firstLeftPlayerName = ref("")
const secondLeftPlayerName = ref("")
const firstRightPlayerName = ref("")
const secondRightPlayerName = ref("")
const matchDuration = ref("")
const consecutiveLeftWidth = ref("")
const consecutiveRightWidth = ref("")
const longestLeftWidth = ref("")
const longestRightWidth = ref("")
const shortestLeftWidth = ref("")
const shortestRightWidth = ref("")
const averagePerPointLeftWidth = ref("")
const averagePerpointRightWidth = ref("")

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
        consecutiveLeftWidth.value = gameData.statistics.consecutivePointsRatio.split(":")[0] + "%"
        consecutiveRightWidth.value = gameData.statistics.consecutivePointsRatio.split(":")[1] + "%"
        longestLeftWidth.value = gameData.statistics.longestPointRatio.split(":")[0] + "%"
        longestRightWidth.value = gameData.statistics.longestPointRatio.split(":")[1] + "%"
        shortestLeftWidth.value = gameData.statistics.shortestPointRatio.split(":")[0] + "%"
        shortestRightWidth.value = gameData.statistics.shortestPointRatio.split(":")[1] + "%"
        averagePerPointLeftWidth.value = gameData.statistics.averagePerPointRatio.split(":")[0] + "%"
        averagePerpointRightWidth.value = gameData.statistics.averagePerPointRatio.split(":")[1] + "%"
    }
}

const getStatistics = async () => {
    const res = await gameStore.getGameStatistics({
        gameId: Array.isArray(route.params.id) ? route.params.id[0] : route.params.id
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
        <div class="statistics-section">
            <div class="a-statistics-title">
                <div class="statistics-centering">
                    <div>{{ gameData.statistics ? gameData.statistics.leftConsecutivePoints : "" }}</div>
                    <div class="grow center">CONSECUTIVE POINTS</div>
                    <div>{{ gameData.statistics ? gameData.statistics.rightConsecutivePoints : "" }}</div>
                </div>

            </div>
            <div class="a-statistics-body">
                <div class="statistics-centering">
                    <div class="percentage-bar statistic-red" :style="{ width: consecutiveLeftWidth }"></div>
                    <div class="percentage-bar statistic-green" :style="{ width: consecutiveRightWidth }"></div>
                </div>
            </div>
            <div class="b-statistics-title">
                <div class="statistics-centering">
                    <div>{{ gameData.statistics ? gameData.statistics.leftLongestPoint : "" }}</div>
                    <div class="grow center">LONGEST POINT</div>
                    <div>{{ gameData.statistics ? gameData.statistics.rightLongestPoint : "" }}</div>
                </div>
            </div>
            <div class="b-statistics-body">
                <div class="statistics-centering">
                    <div class="percentage-bar statistic-red" :style="{ width: longestLeftWidth }"></div>
                    <div class="percentage-bar statistic-green" :style="{ width: longestRightWidth }"></div>
                </div>
            </div>
            <div class="c-statistics-title">
                <div class="statistics-centering">
                    <div>{{ gameData.statistics ? gameData.statistics.leftShortestPoint : "" }}</div>
                    <div class="grow center">SHORTEST POINT</div>
                    <div>{{ gameData.statistics ? gameData.statistics.rightShortestPoint : "" }}</div>
                </div>
            </div>
            <div class="c-statistics-body">
                <div class="statistics-centering">
                    <div class="percentage-bar statistic-red" :style="{ width: shortestLeftWidth }"></div>
                    <div class="percentage-bar statistic-green" :style="{ width: shortestRightWidth }"></div>
                </div>
            </div>
            <div class="d-statistics-title">
                <div class="statistics-centering">
                    <div>{{ gameData.statistics ? gameData.statistics.leftAveragePerPoint : "" }}</div>
                    <div class="grow center">AVERAGE TIME / POINT</div>
                    <div>{{ gameData.statistics ? gameData.statistics.rightAveragePerPoint : "" }}</div>
                </div>
            </div>
            <div class="d-statistics-body">
                <div class="statistics-centering">
                    <div class="percentage-bar statistic-red" :style="{ width: averagePerPointLeftWidth }"></div>
                    <div class="percentage-bar statistic-green" :style="{ width: averagePerpointRightWidth }"></div>
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
    color: statistic-red;
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

.statistics-section {
    margin: 10px;
    padding: 10px;
    background-color: lightgray;
    border-radius: 10px;
    display: grid;
    grid-template-columns: repeat(6, 1fr);
    grid-template-rows: repeat(8, 1fr);
    background-color: aqua;
    width: 95%;
    height: 35vh;
}

.a-statistics-title {
    grid-column: 1 / span 6;
    grid-row: 1 / span 1;
    background-color: lightgray;
}

.a-statistics-body {
    grid-column: 1 / span 6;
    grid-row: 2 / span 1;
    background-color: lightgray;
}

.b-statistics-title {
    grid-column: 1 / span 6;
    grid-row: 3 / span 1;
    background-color: lightgray;
}

.b-statistics-body {
    grid-column: 1 / span 6;
    grid-row: 4 / span 1;
    background-color: lightgray;
}

.c-statistics-title {
    grid-column: 1 / span 6;
    grid-row: 5 / span 1;
    background-color: lightgray;
}

.c-statistics-body {
    grid-column: 1 / span 6;
    grid-row: 6 / span 1;
    background-color: lightgray;
}

.d-statistics-title {
    grid-column: 1 / span 6;
    grid-row: 7 / span 1;
    background-color: lightgray;
}

.d-statistics-body {
    grid-column: 1 / span 6;
    grid-row: 8 / span 1;
    background-color: lightgray;
}

.statistics-centering {
    display: flex;
    align-items: center;
    justify-content: center;
    height: 100%;
    padding: 0 10px;
}

.grow {
    flex-grow: 1;
}

.center {
    display: flex;
    justify-content: center;
}

.percentage-bar {
    height: 100%;
}

.statistic-red {
    background-color: red;
}

.statistic-green {
    background-color: green;
}
</style>