<script setup lang="ts">
import shuttlecock from '@/assets/images/shuttlecock.png';
import { CurrentServer, GameTypes } from '@/enums/game';
import { useGameStore } from '@/stores/game-store';
import type { LocalGameStep } from '@/types/game';
import { DateTime } from "luxon";
import Swal from 'sweetalert2';
import { v4 as uuidv4 } from 'uuid';
import { onBeforeUnmount, onMounted, reactive, ref } from 'vue';
import { useRouter } from 'vue-router';

const router = useRouter()
const errorMessage = ref("")
const isLoading = ref(false)
const isLandscape = ref(false)
const pointsOrientation = ref("equal")
const currentCourtState = reactive({
    leftEvenPlayer: "",
    leftOddPlayer: "",
    rightEvenPlayer: "",
    rightOddPlayer: "",
    currentServer: "",
    teamLeftScore: 0,
    teamRightScore: 0
})

const gameStore = useGameStore()
gameStore.setBackendUrl(import.meta.env.VITE_PROXY_URL)

gameStore.$subscribe(() => {
    updateCourtState()
})


onMounted(() => {
    handleOrientationChange();
    window.addEventListener('resize', handleOrientationChange);

    updateCourtState()
})

onBeforeUnmount(() => {
    window.removeEventListener('resize', handleOrientationChange);
});

const handleOrientationChange = () => {
    isLandscape.value = window.matchMedia("(orientation: landscape)").matches;
};

const updateCourtState = () => {
    const lastProgress = gameStore.currentGameProgress[gameStore.currentGameProgress.length - 1]

    currentCourtState.leftEvenPlayer = lastProgress.leftEvenPlayerName
    currentCourtState.leftOddPlayer = lastProgress.leftOddPlayerName
    currentCourtState.rightEvenPlayer = lastProgress.rightEvenPlayerName
    currentCourtState.rightOddPlayer = lastProgress.rightOddPlayerName
    currentCourtState.currentServer = lastProgress.currentServer
    currentCourtState.teamLeftScore = lastProgress.teamLeftScore
    currentCourtState.teamRightScore = lastProgress.teamRightScore
}

const handleScorePoint = (position: string) => {

    if (gameStore.currentGameSettings.gameType == GameTypes.GAME_TYPE_DOUBLES) {
        handleDoubleScoring(position)
        return
    }

    handleSingleScoring(position)
}

const handleSingleScoring = (position: string) => {

    const lastProgress = gameStore.currentGameProgress[gameStore.currentGameProgress.length - 1]
    const leftPlayerName = lastProgress.leftEvenPlayerName == "" ? lastProgress.leftOddPlayerName : lastProgress.leftEvenPlayerName
    const rightPlayerName = lastProgress.rightEvenPlayerName == "" ? lastProgress.rightOddPlayerName : lastProgress.rightEvenPlayerName

    const progress: LocalGameStep = {
        isSynced: false,
        id: "",
        gameId: gameStore.currentGameSettings.id,
        scoreAt: DateTime.utc().toString(),
        stepNum: gameStore.currentGameProgress.length + 1,
        currentServer: "" as CurrentServer,
        leftEvenPlayerName: lastProgress.leftEvenPlayerName,
        leftOddPlayerName: lastProgress.leftOddPlayerName,
        rightEvenPlayerName: lastProgress.rightEvenPlayerName,
        rightOddPlayerName: lastProgress.rightOddPlayerName,
        teamLeftScore: lastProgress.teamLeftScore,
        teamRightScore: lastProgress.teamRightScore,
        syncId: uuidv4(),
        createdAt: "",
        updatedAt: "",
    }

    let isEvenLeader: boolean = false

    if (position === "left") {
        progress.teamLeftScore = lastProgress.teamLeftScore + 1
        isEvenLeader = progress.teamLeftScore % 2 == 0

        if (isEvenLeader) {
            progress.currentServer = CurrentServer.SERVER_LEFT_EVEN
        } else {
            progress.currentServer = CurrentServer.SERVER_LEFT_ODD
        }
    } else {
        progress.teamRightScore = lastProgress.teamRightScore + 1
        isEvenLeader = progress.teamRightScore % 2 == 0

        if (isEvenLeader) {
            progress.currentServer = CurrentServer.SERVER_RIGHT_EVEN
        } else {
            progress.currentServer = CurrentServer.SERVER_RIGHT_ODD
        }
    }

    if (isEvenLeader) {
        progress.leftEvenPlayerName = leftPlayerName
        progress.leftOddPlayerName = ""
        progress.rightEvenPlayerName = rightPlayerName
        progress.rightOddPlayerName = ""
    } else {
        progress.leftEvenPlayerName = ""
        progress.leftOddPlayerName = leftPlayerName
        progress.rightEvenPlayerName = ""
        progress.rightOddPlayerName = rightPlayerName
    }

    gameStore.currentGameProgress = gameStore.currentGameProgress.concat(progress)
}

const handleDoubleScoring = (position: string) => {
    const lastProgress = gameStore.currentGameProgress[gameStore.currentGameProgress.length - 1]

    const progress: LocalGameStep = {
        isSynced: false,
        id: "",
        gameId: gameStore.currentGameSettings.id,
        teamLeftScore: lastProgress.teamLeftScore,
        teamRightScore: lastProgress.teamRightScore,
        scoreAt: DateTime.utc().toString(),
        stepNum: gameStore.currentGameProgress.length + 1,
        currentServer: "" as CurrentServer,
        leftEvenPlayerName: lastProgress.leftEvenPlayerName,
        leftOddPlayerName: lastProgress.leftOddPlayerName,
        rightEvenPlayerName: lastProgress.rightEvenPlayerName,
        rightOddPlayerName: lastProgress.rightOddPlayerName,
        syncId: uuidv4(),
        createdAt: "",
        updatedAt: "",
    }

    if (position === "left") {
        progress.teamLeftScore = lastProgress.teamLeftScore + 1

        if (lastProgress.currentServer == CurrentServer.SERVER_LEFT_EVEN || lastProgress.currentServer == CurrentServer.SERVER_LEFT_ODD) {
            progress.currentServer = lastProgress.currentServer == CurrentServer.SERVER_LEFT_EVEN ? CurrentServer.SERVER_LEFT_ODD : CurrentServer.SERVER_LEFT_EVEN

            const evenPlayer = progress.leftEvenPlayerName
            progress.leftEvenPlayerName = progress.leftOddPlayerName
            progress.leftOddPlayerName = evenPlayer
        } else {
            if (progress.teamLeftScore % 2 != 0) {
                progress.currentServer = CurrentServer.SERVER_LEFT_ODD
            } else {
                progress.currentServer = CurrentServer.SERVER_LEFT_EVEN
            }
        }

    } else {
        progress.teamRightScore = lastProgress.teamRightScore + 1

        if (lastProgress.currentServer == CurrentServer.SERVER_RIGHT_EVEN || lastProgress.currentServer == CurrentServer.SERVER_RIGHT_ODD) {
            progress.currentServer = lastProgress.currentServer == CurrentServer.SERVER_RIGHT_EVEN ? CurrentServer.SERVER_RIGHT_ODD : CurrentServer.SERVER_RIGHT_EVEN

            const evenPlayer = progress.rightEvenPlayerName
            progress.rightEvenPlayerName = progress.rightOddPlayerName
            progress.rightOddPlayerName = evenPlayer
        } else {
            if (progress.teamRightScore % 2 != 0) {
                progress.currentServer = CurrentServer.SERVER_RIGHT_ODD
            } else {
                progress.currentServer = CurrentServer.SERVER_RIGHT_EVEN
            }
        }
    }

    gameStore.currentGameProgress = gameStore.currentGameProgress.concat(progress)
}

const handlePointsOrientation = (orientation: string) => {
    pointsOrientation.value = orientation
}

const handleUndo = () => {
    Swal.fire({
        title: 'Confirm undo?',
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
            let lastProgress = gameStore.currentGameProgress[gameStore.currentGameProgress.length - 1]

            while (lastProgress.id == "") {
                await delaySeconds(0.5)
                lastProgress = gameStore.currentGameProgress[gameStore.currentGameProgress.length - 1]
            }

            const toRemove = gameStore.currentGameProgress.splice(gameStore.currentGameProgress.length - 1, 1)

            if (toRemove.length > 0) {
                gameStore.stepsToRemove = gameStore.stepsToRemove.concat(toRemove[0].id)
            }
        }
    })
}

const isNeedsSyncing = () => {
    const toAdd = gameStore.currentGameProgress.filter((step) => {
        return !step.isSynced
    })

    return toAdd.length > 0 || gameStore.stepsToRemove.length > 0
}

const delaySeconds = async (milliseconds: number) => {
    return new Promise(resolve => setTimeout(resolve, milliseconds));
}

const handleEndGame = async () => {

    isLoading.value = true

    Swal.fire({
        title: 'Confirm end game?',
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

            let needsSyncing = isNeedsSyncing()
            while (needsSyncing) {
                await delaySeconds(1)
                needsSyncing = isNeedsSyncing()
            }

            const res = await gameStore.endGame({
                gameId: gameStore.currentGameSettings.id,
                endGameRequest: {
                    isEnded: true
                }
            })

            if (res instanceof Error) {
                isLoading.value = false
                errorMessage.value = res.message
                return
            }

            isLoading.value = false

            const gameId = gameStore.currentGameSettings.id
            localStorage.removeItem("game")

            router.push({
                name: "game/statistics",
                params: {
                    id: gameId
                }
            })
        }
    })

    isLoading.value = false
}

</script>

<template>
    <div>
        <div v-if="!isLandscape" class="portrait-warning">
            Please rotate your device to landscape orientation.
        </div>
        <div v-else class="main-content">
            <div class="header-actions">
                <div class="points-control">
                    <button class="button-primary points-control-button" @click="handlePointsOrientation('equal')">
                        {{ "O : O" }}
                    </button>
                    <button class="button-primary points-control-button" @click="handlePointsOrientation('left')">
                        {{ "O : []" }}
                    </button>
                    <button class="button-primary points-control-button" @click="handlePointsOrientation('right')">
                        {{ "[] : O" }}
                    </button>
                </div>
                <div class="points">{{ `${currentCourtState.teamLeftScore} :
                    ${currentCourtState.teamRightScore}` }}
                </div>
                <div></div>
            </div>
            <div class="content-section">
                <button v-if="pointsOrientation == 'equal'" class="add-button equal" @click="handleScorePoint('left')">
                    + 1
                </button>
                <div v-if="pointsOrientation == 'left'" class="sides-add-button-wrapper">
                    <button class="add-button sides" @click="handleScorePoint('left')">
                        + 1
                    </button>
                    <button class="add-button sides red" @click="handleScorePoint('right')">
                        + 1
                    </button>
                </div>
                <div class="court">
                    <div class="sideline sideline-left squares"></div>
                    <div class="top-court squares">
                        <div class="net">
                            <div class="v-line"></div>
                        </div>
                    </div>
                    <div class="bottom-court squares">
                        <div class="net">
                            <div class="v-line"></div>
                        </div>
                    </div>
                    <div class="sideline sideline-right squares"></div>
                    <div class="left-top-player squares">
                        <div class="player-names">{{ currentCourtState.leftOddPlayer }}</div>
                    </div>
                    <div class="left-bottom-player squares">
                        <div class="player-names">{{ currentCourtState.leftEvenPlayer }}</div>
                    </div>
                    <div class="right-top-player squares">
                        <div class="player-names">{{ currentCourtState.rightEvenPlayer }}</div>
                    </div>
                    <div class="right-bottom-player squares">
                        <div class="player-names">{{ currentCourtState.rightOddPlayer }}</div>
                    </div>
                    <div class="left-top-backline squares">
                        <div v-if="currentCourtState.currentServer === CurrentServer.SERVER_LEFT_ODD"
                            class="shuttle-wrapper">
                            <img :src="shuttlecock" width="30px" height="30px">
                        </div>
                    </div>
                    <div class="left-bottom-backline squares">
                        <div v-if="currentCourtState.currentServer === CurrentServer.SERVER_LEFT_EVEN"
                            class="shuttle-wrapper">
                            <img :src="shuttlecock" width="30px" height="30px">
                        </div>
                    </div>
                    <div class="right-top-backline squares">
                        <div v-if="currentCourtState.currentServer === CurrentServer.SERVER_RIGHT_EVEN"
                            class="shuttle-wrapper">
                            <img :src="shuttlecock" width="30px" height="30px">
                        </div>
                    </div>
                    <div class="right-bottom-backline squares">
                        <div v-if="currentCourtState.currentServer === CurrentServer.SERVER_RIGHT_ODD"
                            class="shuttle-wrapper">
                            <img :src="shuttlecock" width="30px" height="30px">
                        </div>
                    </div>
                </div>
                <button v-if="pointsOrientation == 'equal'" class="add-button equal red"
                    @click="handleScorePoint('right')">
                    + 1
                </button>
                <div v-if="pointsOrientation == 'right'" class="sides-add-button-wrapper">
                    <button class="add-button sides" @click="handleScorePoint('left')">
                        + 1
                    </button>
                    <button class="add-button sides red" @click="handleScorePoint('right')">
                        + 1
                    </button>
                </div>
            </div>
            <div class="footer-actions">
                <div class="push-end">
                    <p class="error-message" id="error-message" v-if='errorMessage !== ""'>{{ errorMessage }}</p>
                    <div>
                        <button class="button-secondary footer-buttons" @click="handleEndGame()"
                            :disabled="isLoading">{{ isLoading ?
                                "Loading.." : "End Game"
                            }}
                        </button>
                    </div>
                    <button class="button-secondary undo-button footer-buttons" @click="handleUndo()"
                        :disabled="isLoading">Undo</button>
                </div>
            </div>
        </div>
    </div>
</template>

<style scoped>
.portrait-warning {
    display: block;
    text-align: center;
    font-size: 20px;
    color: #333;
}

@media only screen and (orientation: landscape) {
    .mt-1 {
        margin-top: 1rem;
    }

    .mb-1 {
        margin-bottom: 1rem;
    }

    .main-content {
        display: flex;
        flex-direction: column;
        min-width: 100vw;
    }

    .court {
        display: grid;
        grid-template-columns: repeat(8, 1fr);
        grid-template-rows: repeat(4, 1fr);
        width: 75vw;
        height: 66vh;
        background-color: green;
        position: relative;
        border: 4px solid white;
    }

    .squares {
        border: 1px solid white;
        display: flex;
        justify-content: center;
        align-items: center;
        color: white;
        font-size: 0.8rem;
    }

    .net {
        height: 100%;
        align-items: center;
    }

    .v-line {
        height: 100%;
        border-left: 1px solid white;
    }

    .service-top {
        grid-row: 1 / span 2;
    }

    .service-bottom {
        grid-row: 3 / span 2;
    }

    .top-court,
    .bottom-court {
        grid-column: 1 / span 8;
    }

    .top-court {
        grid-row: 1 / span 2;
    }

    .bottom-court {
        grid-row: 3 / span 2;
    }

    .sideline {
        grid-row: 1 / span 4;
    }

    .sideline-left {
        grid-column: 1 / span 1;
    }

    .sideline-right {
        grid-column: 8 / span 1;
    }

    .left-top-player {
        grid-column: 2 / span 3;
        grid-row: 1 / span 2;
    }

    .left-bottom-player {
        grid-column: 2 / span 3;
        grid-row: 3 / span 2;
    }

    .right-top-player {
        background-color: #D32F2F;
        grid-column: 5 / span 3;
        grid-row: 1 / span 2;
    }

    .right-bottom-player {
        background-color: #D32F2F;
        grid-column: 5 / span 3;
        grid-row: 3 / span 2;
    }

    .left-top-backline {
        grid-column: 1 / span 1;
        grid-row: 1 / span 2;
    }

    .left-bottom-backline {
        grid-column: 1 / span 1;
        grid-row: 3 / span 2;
    }

    .right-top-backline {
        background-color: #D32F2F;
        grid-column: 8 / span 1;
        grid-row: 1 / span 2;
    }

    .right-bottom-backline {
        background-color: #D32F2F;
        grid-column: 8 / span 1;
        grid-row: 3 / span 2;
    }

    .setup-form-container {
        display: flex;
        flex-direction: column;
        margin-top: 1rem;
    }

    form {
        padding: 0.5rem;
        flex-grow: 1;
    }

    .form-group {
        display: flex;
        padding: 0 1rem;
    }

    fieldset {
        margin-bottom: 10px;
        border: 1px solid #ccc;
        padding: 10px;
    }

    legend {
        font-weight: bold;
    }

    .loading-wrapper {
        display: none;
        padding: 0.5rem;
    }

    .shuttle-wrapper {
        padding: 0.5rem;
    }

    .content-section {
        display: flex;
        align-items: center;
        justify-content: center;
    }

    .add-button {
        background-color: green;
        color: white;
        border: none;
        border-radius: 4px;
        cursor: pointer;
        font-weight: bold;
        font-size: 1.5rem;
    }

    .add-button.equal {
        margin-bottom: 1rem;
        flex-grow: 1;
        height: 200px;
    }

    .add-button.sides {
        margin-bottom: 1rem;
        flex-grow: 1;
        height: 100px;
    }

    .add-button.red {
        background-color: #D32F2F;
    }

    .player-names {
        font-weight: bold;
        font-size: 1.5rem;
        padding: 0.5rem;
    }

    .header-actions {
        display: grid;
        grid-template-columns: repeat(3, 1fr);
        margin-top: 0.5rem;
        font-weight: bold;
        font-size: 1.5rem;
    }

    .footer-actions {
        display: flex;
        justify-content: end;
        margin-top: 0.5rem;
        margin-right: 0.5rem;
    }

    .footer-buttons {
        border: none;
        border-radius: 4px;
        cursor: pointer;
        font-weight: bold;
        font-size: 1.5rem;
        padding: 1rem;
    }

    .push-end {
        display: flex;
        margin-right: 2rem;
    }

    .undo-button {
        margin-left: 1rem;
    }

    .points-control {
        display: flex;
    }

    .points-control-button {
        margin-left: 0.5rem;
        border: none;
        border-radius: 4px;
        cursor: pointer;
        font-weight: bold;
        font-size: 1rem;
        padding: 0.5rem;
    }

    .sides-add-button-wrapper {
        display: flex;
        flex-direction: column;
        width: 100px;
    }

    .points {
        display: flex;
        justify-content: center;
        font-size: 2rem;
    }
}
</style>