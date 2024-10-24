<script setup lang="ts">
import shuttlecock from '@/assets/images/shuttlecock.png';
import { CurrentServer } from '@/enums/game';
import { useGameStore } from '@/stores/game-store';
import type { LocalGameStep } from '@/types/game';
import { DateTime } from "luxon";
import { onBeforeUnmount, onMounted, reactive, ref } from 'vue';

const isLandscape = ref(false)
const currentCourtState = reactive({
    leftEvenPlayer: "",
    leftOddPlayer: "",
    rightEvenPlayer: "",
    rightOddPlayer: "",
    currentServer: ""
})

const gameStore = useGameStore()
gameStore.setBackendUrl(import.meta.env.VITE_BACKEND_URL)

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
}

const handleScorePoint = (position: string) => {
    const lastProgress = gameStore.currentGameProgress[gameStore.currentGameProgress.length - 1]

    if (position === "left") {
        const progress: LocalGameStep = {
            isSynced: false,
            id: "",
            gameId: gameStore.currentGameSettings.id,
            teamLeftScore: lastProgress.teamLeftScore + 1,
            teamRightScore: lastProgress.teamRightScore,
            scoreAt: DateTime.now().toString(),
            stepNum: gameStore.currentGameProgress.length + 1,
            currentServer: "" as CurrentServer,
            leftEvenPlayerName: lastProgress.leftEvenPlayerName,
            leftOddPlayerName: lastProgress.leftOddPlayerName,
            rightEvenPlayerName: lastProgress.rightEvenPlayerName,
            rightOddPlayerName: lastProgress.rightOddPlayerName,
            createdAt: "",
            updatedAt: "",
        }

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

        gameStore.currentGameProgress = gameStore.currentGameProgress.concat(progress)
    } else {
        const progress = {
            isSynced: false,
            id: "",
            gameId: gameStore.currentGameSettings.id,
            teamLeftScore: lastProgress.teamLeftScore,
            teamRightScore: lastProgress.teamRightScore + 1,
            scoreAt: DateTime.now().toString(),
            stepNum: gameStore.currentGameProgress.length + 1,
            currentServer: "" as CurrentServer,
            leftEvenPlayerName: lastProgress.leftEvenPlayerName,
            leftOddPlayerName: lastProgress.leftOddPlayerName,
            rightEvenPlayerName: lastProgress.rightEvenPlayerName,
            rightOddPlayerName: lastProgress.rightOddPlayerName,
            createdAt: "",
            updatedAt: "",
        }

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

        gameStore.currentGameProgress = gameStore.currentGameProgress.concat(progress)
    }
}

</script>

<template>
    <div>
        <div v-if="!isLandscape" class="portrait-warning">
            Please rotate your device to landscape orientation.
        </div>
        <div v-else class="main-content">
            <div class="header-actions">header actions</div>
            <div class="content-section">
                <div>
                    <button class="primary-button" @click="handleScorePoint('left')">
                        Add +
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
                        {{ currentCourtState.leftOddPlayer }}
                    </div>
                    <div class="left-bottom-player squares">
                        {{ currentCourtState.leftEvenPlayer }}
                    </div>
                    <div class="right-top-player squares">
                        {{ currentCourtState.rightEvenPlayer }}
                    </div>
                    <div class="right-bottom-player squares">
                        {{ currentCourtState.rightOddPlayer }}
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
                <div><button class="primary-button" @click="handleScorePoint('right')">
                        Add +
                    </button></div>
            </div>
            <div class="footer-actions">footer actions</div>
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
        width: 70vw;
        height: 70vh;
        background-color: green;
        position: relative;
        border: 4px solid white;
        margin: 1rem;
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
    }
}
</style>