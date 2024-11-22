<script setup lang="ts">
import userImage from '@/assets/images/user.png';
import ShareButton from '@/components/ShareButton.vue';
import { type GetGame200Response } from "@/repositories/clients/public";
import router from '@/router';
import { useGameStore } from '@/stores/game-store';
import { useSessionStore } from '@/stores/session-store';
import Swal from 'sweetalert2';
import { computed, onBeforeMount, reactive, ref } from 'vue';
import { useRoute } from 'vue-router';

const route = useRoute()
const gameData = reactive({} as GetGame200Response)
const errorMessage = ref("")
const winningTeam = ref("")
const leftOddPlayerName = ref("")
const leftEvenPlayerName = ref("")
const rightEvenPlayerName = ref("")
const rightOddPlayerName = ref("")
const matchDuration = ref("")
const consecutiveLeftWidth = ref("")
const consecutiveRightWidth = ref("")
const longestLeftWidth = ref("")
const longestRightWidth = ref("")
const shortestLeftWidth = ref("")
const shortestRightWidth = ref("")
const averagePerPointLeftWidth = ref("")
const averagePerpointRightWidth = ref("")

const shareURL = `${window.location.href}`
const gameStore = useGameStore()
gameStore.setBackendUrl(import.meta.env.VITE_PROXY_URL)

const sessionStore = useSessionStore()

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
        leftOddPlayerName.value = gameData.game.leftOddPlayerName
        leftEvenPlayerName.value = gameData.game.leftEvenPlayerName
        rightEvenPlayerName.value = gameData.game.rightEvenPlayerName
        rightOddPlayerName.value = gameData.game.rightOddPlayerName
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
    gameData.steps = data.steps
    gameData.statistics = data.statistics

    updateDisplay()
    updateMetaTags()
}

const playerScoreText = computed(() => {
    if (gameData.game && gameData.steps && gameData.steps.length > 0) {
        const last = gameData.steps[gameData.steps.length - 1]

        let leftPlayers = gameData.game.leftEvenPlayerName
        if (gameData.game.leftOddPlayerName !== "") {
            leftPlayers = `${leftPlayers},${gameData.game.leftOddPlayerName}`
        }

        let rightPlayers = gameData.game.rightEvenPlayerName
        if (gameData.game.rightOddPlayerName !== "") {
            rightPlayers = `${leftPlayers},${gameData.game.rightOddPlayerName}`
        }

        let leftWinner, rightWinner = ""
        if (last.teamLeftScore > last.teamRightScore) {
            leftWinner = "\uD83C\uDFC6"
        } else {
            rightWinner = "\uD83C\uDFC6"
        }

        return `${leftPlayers} ${leftWinner} ${last.teamLeftScore}:${last.teamRightScore} ${rightWinner} ${rightPlayers}\n`
    }
    return ""
})

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

const updateMetaTags = () => {
    document.title = "🏸 Badminton Game Results 🏸";
    const metaTags = [
        { property: "og:title", content: "🏸 Badminton Game Results 🏸" },
        { property: "og:description", content: playerScoreText.value },
        { property: "og:image", content: "/images/badminton-graph-image.jpg" },
        { property: "og:url", content: window.location.href },
        { property: "og:type", content: "website" },
    ];

    metaTags.forEach(tag => {
        const meta = document.createElement("meta");
        meta.setAttribute("property", tag.property);
        meta.setAttribute("content", tag.content);
        document.head.appendChild(meta);
    });
}

const handlePlayerIdentify = (value: string) => {

    const token = sessionStorage.getItem('session_token')
    if (!token || token === "") {
        document.getElementsByName("player_identify").forEach((radio) => {
            (radio as HTMLInputElement).checked = false
        });

        Swal.fire({
            title: 'Sign up now to unlock this feature, ready to join?',
            showCancelButton: true,
            confirmButtonText: 'Yes',
            customClass: {
                actions: 'my-actions',
                cancelButton: 'order-1 right-gap',
                confirmButton: 'order-2',
                denyButton: 'order-3',
            },
        }).then(async (result) => {
            if (result.isConfirmed) {
                sessionStore.toRedirectToUrl = shareURL
                router.push("/signup-player")
            }
        })
    }

    // call api bind user to player

}
</script>

<template>
    <div class="container">
        <div class="header-section">
            <div></div>
            <div>
                <ShareButton :title="'🏸 Badminton Game Results 🏸'" :text="playerScoreText" :url="shareURL" />
            </div>
        </div>
        <div class="player-section">
            <div class="player-card">
                <div class="user-container-box user-mb" v-if="leftOddPlayerName != ''">
                    <div class="identify-box">
                        <input type="radio" @change="handlePlayerIdentify('left_odd_player_identify')"
                            name="player_identify" value="left_odd_player_identify">
                        <label for="player_identify"> This is me</label>
                    </div>
                    <div class="user-container">
                        <div class="profile-photo-container">
                            <img :src="userImage" alt="left odd player image" width="30px" height="30px">
                        </div>
                        <span class="player-name">{{ leftOddPlayerName }}</span>
                    </div>
                </div>

                <div class="user-container-box" v-if="leftEvenPlayerName != ''">
                    <div class="identify-box">
                        <input type="radio" @change="handlePlayerIdentify('left_even_player_identify')"
                            name="player_identify" value="left_even_player_identify">
                        <label for="player_identify"> This is me</label>
                    </div>
                    <div class="user-container">
                        <div class="profile-photo-container">
                            <img :src="userImage" alt="left even player image" width="30px" height="30px">
                        </div>
                        <span class="player-name">{{ leftEvenPlayerName }}</span>
                    </div>
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
                <div class="user-container-box" v-if="rightEvenPlayerName != ''">

                    <div class="identify-box">
                        <input type="radio" @change="handlePlayerIdentify('right_even_player_identify')"
                            name="player_identify" value="right_even_player_identify">
                        <label for="player_identify"> This is me</label>
                    </div>
                    <div class="user-container">
                        <div class="profile-photo-container">
                            <img :src="userImage" alt="right even player image" width="30px" height="30px">
                        </div>
                        <span class="player-name">{{ rightEvenPlayerName }}</span>
                    </div>
                </div>
                <div class="user-container-box user-mt" v-if="rightOddPlayerName != ''">
                    <div class="identify-box">
                        <input type="radio" @change="handlePlayerIdentify('right_odd_player_identify')"
                            name="player_identify" value="right_odd_player_identify">
                        <label for="player_identify"> This is me</label>
                    </div>
                    <div class="user-container">
                        <div class="profile-photo-container">
                            <img :src="userImage" alt="right odd player image" width="30px" height="30px">
                        </div>
                        <span class="player-name">{{ rightOddPlayerName }}</span>
                    </div>
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

.header-section {
    display: flex;
    justify-content: space-between;
    padding: 15px 15px;
}

.whatsapp-box {
    display: flex;
    justify-content: center;
    align-items: center;
}

.user-container-box {
    display: flex;
    flex-direction: column;
    justify-content: center;
    align-items: center;
}

.identify-box {
    margin-bottom: 10px;
}
</style>