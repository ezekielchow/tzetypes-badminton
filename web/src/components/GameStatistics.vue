<script setup lang="ts">
import homeImage from '@/assets/images/home.png';
import trophyImage from '@/assets/images/trophy.png';
import userImage from '@/assets/images/user.png';
import ShareButton from '@/components/ShareButton.vue';
import type { CreateOrUpdateGameHistoryRequestSchemaPlayerPositionEnum } from '@/repositories/clients/private';
import { type GetGame200Response } from "@/repositories/clients/public";
import router from '@/router';
import { useGameStore } from '@/stores/game-store';
import { useSessionStore } from '@/stores/session-store';
import { useUserStore } from '@/stores/user-store';
import type { Auth } from 'firebase/auth';
import Swal from 'sweetalert2';
import { computed, inject, onBeforeMount, reactive, ref } from 'vue';
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
const playerIdentityVal = ref("")

const auth = inject<Auth>("auth");

const shareURL = `${window.location.href}`
const gameStore = useGameStore()
gameStore.setBackendUrl(import.meta.env.VITE_PROXY_URL)

const sessionStore = useSessionStore()
const userStore = useUserStore()

onBeforeMount(async () => {
    await getStatistics()
})

let isLoggedIn = false
const user = userStore.firebaseUser
if (user) {
    isLoggedIn = true
}

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

    const user = userStore.firebaseUser
    if (user) {
        await getGameHistory(gameData.game.id)
    }
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
            rightPlayers = `${rightPlayers},${gameData.game.rightOddPlayerName}`
        }

        let leftWinner = "", rightWinner = ""
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
    pointText: true
}))

const rightPointsClass = computed(() => ({
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

const handlePlayerIdentify = async (value: string) => {

    const user = userStore.firebaseUser
    if (!user) {
        document.getElementsByName("player_identify").forEach((radio) => {
            (radio as HTMLInputElement).checked = false
        });

        return Swal.fire({
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

            return
        })
    }

    const res = await gameStore.createOrUpdateGameHistory({
        gameId: gameData.game.id,
        createOrUpdateGameHistoryRequestSchema: {
            playerPosition: value as CreateOrUpdateGameHistoryRequestSchemaPlayerPositionEnum
        }
    })

    if (res instanceof Error) {
        errorMessage.value = res.message
        document.getElementsByName("player_identify").forEach((radio) => {
            (radio as HTMLInputElement).checked = false
        });
        return
    }
}

const getGameHistory = async (gameId: string) => {
    const res = await gameStore.getGameHistory({
        gameId: gameId
    })

    if (res instanceof Error) {
        return
    }

    playerIdentityVal.value = res.gameHistory.playerPosition
}
</script>

<template>
    <div class="container">
        <div class="header-section">
            <div>
                <RouterLink v-if="isLoggedIn" to="/dashboard"> <img class="back-link" :src="homeImage"
                        alt="home button image" width="30px" height="30px">
                </RouterLink>
            </div>
            <div>
                <ShareButton :title="'🏸 Badminton Game Results 🏸'" :text="playerScoreText" :url="shareURL" />
            </div>
        </div>
        <p class="error-message" id="error-message" v-if='errorMessage !== ""'>{{ errorMessage }}</p>
        <div class="player-section">
            <div class="player-card">
                <div class="user-container-box user-mb" v-if="leftOddPlayerName != ''">
                    <div class="identify-box">
                        <input type="radio" @change="handlePlayerIdentify('left_odd_player_identify')"
                            name="player_identify" value="left_odd_player_identify" v-model="playerIdentityVal">
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
                            name="player_identify" value="left_even_player_identify" v-model="playerIdentityVal">
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
                    <div v-if="winningTeam == 'right'" :style="{ width: '30px', height: '30px' }"></div>
                    <img v-if="winningTeam == 'left'" :src="trophyImage" width="30" height="30">
                    <span :class="leftPointsClass">{{ leftPoint }}</span>
                    <span class="point-seperator"> : </span>
                    <span :class="rightPointsClass">{{ rightPoint }}</span>
                    <img v-if="winningTeam == 'right'" :src="trophyImage" width="30" height="30" />
                    <div v-if="winningTeam == 'left'" :style="{ width: '30px', height: '30px' }"></div>
                </div>
                <div class="time-container">
                    <span>{{ matchDuration }}</span>
                </div>
            </div>
            <div class="player-card">
                <div class="user-container-box" v-if="rightEvenPlayerName != ''">

                    <div class="identify-box">
                        <input type="radio" @change="handlePlayerIdentify('right_even_player_identify')"
                            name="player_identify" value="right_even_player_identify" v-model="playerIdentityVal">
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
                            name="player_identify" value="right_odd_player_identify" v-model="playerIdentityVal">
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
            <div class="statistics-title">
                <div>{{ gameData.statistics ? gameData.statistics.leftConsecutivePoints : "" }}</div>
                <div class="grow center" :style="{ fontWeight: 'bold' }">CONSECUTIVE POINTS</div>
                <div>{{ gameData.statistics ? gameData.statistics.rightConsecutivePoints : "" }}</div>
            </div>
            <div class="statistics-body">
                <div class="percentage-bar statistic-left" :style="{ width: consecutiveLeftWidth }"></div>
                <div class="percentage-bar statistic-right" :style="{ width: consecutiveRightWidth }"></div>
            </div>
            <div class="statistics-title">
                <div>{{ gameData.statistics ? gameData.statistics.leftLongestPoint : "" }}</div>
                <div class="grow center" :style="{ fontWeight: 'bold' }">LONGEST POINT</div>
                <div>{{ gameData.statistics ? gameData.statistics.rightLongestPoint : "" }}</div>
            </div>
            <div class="statistics-body">
                <div class="percentage-bar statistic-left" :style="{ width: longestLeftWidth }"></div>
                <div class="percentage-bar statistic-right" :style="{ width: longestRightWidth }"></div>
            </div>
            <div class="statistics-title">
                <div>{{ gameData.statistics ? gameData.statistics.leftShortestPoint : "" }}</div>
                <div class="grow center" :style="{ fontWeight: 'bold' }">SHORTEST POINT</div>
                <div>{{ gameData.statistics ? gameData.statistics.rightShortestPoint : "" }}</div>
            </div>
            <div class="statistics-body">
                <div class="percentage-bar statistic-left" :style="{ width: shortestLeftWidth }"></div>
                <div class="percentage-bar statistic-right" :style="{ width: shortestRightWidth }"></div>
            </div>
            <div class="statistics-title">
                <div>{{ gameData.statistics ? gameData.statistics.leftAveragePerPoint : "" }}</div>
                <div class="grow center" :style="{ fontWeight: 'bold' }">AVERAGE TIME / POINT</div>
                <div>{{ gameData.statistics ? gameData.statistics.rightAveragePerPoint : "" }}</div>
            </div>
            <div class="statistics-body">
                <div class="percentage-bar statistic-left" :style="{ width: averagePerPointLeftWidth }"></div>
                <div class="percentage-bar statistic-right" :style="{ width: averagePerpointRightWidth }"></div>
            </div>
        </div>
    </div>
</template>

<style scoped>
.container {
    display: flex;
    flex-direction: column;
    min-height: 100vh;
}

.player-section {
    display: flex;
    justify-content: center;
    width: 100vw;
    min-height: 30vh;
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

.isWinner {
    background-image: url('@/assets/images/trophy.png');
    background-size: cover;
    background-position: center;
    background-repeat: no-repeat;
}

.pointText {
    font-size: 2rem;
    margin: 0 5px;
}

.point-seperator {
    font-size: 1.5rem;
}

.user-mt {
    margin-top: 30px;
}

.user-mb {
    margin-bottom: 30px;
}

.player-name {
    font-weight: bold;
    text-wrap: wrap;
}

.statistics-section {
    margin-top: 30px;
    padding: 10px;
    background-color: #DCEBFF;
    border-radius: 10px;
    display: flex;
    flex-direction: column;
    width: 100%;
    min-height: 35vh;
}

.grow {
    flex-grow: 1;
}

.center {
    display: flex;
    justify-content: center;
}

.percentage-bar {
    min-height: 20px;
}

.statistic-left {
    background-color: #4A90E2;
}

.statistic-right {
    background-color: #27AE60;
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

.statistics-title {
    display: flex;
    justify-content: space-evenly;
}

.statistics-body {
    display: flex;
    min-width: 100%;
    margin-bottom: 20px;
}
</style>