<script setup lang="ts">
import homeImage from '@/assets/images/home.png';
import loadingImage from '@/assets/images/loading.png';
import shuttlecock from '@/assets/images/shuttlecock.png';
import { CurrentServer, GameTypes } from '@/enums/game';
import { GameStartRequestSchemaGameTypeEnum, GameStartRequestSchemaServingSideEnum, type GameStartRequestSchema } from '@/repositories/clients/private';
import { useGameStore } from '@/stores/game-store';
import { onBeforeUnmount, onMounted, ref } from 'vue';
import { useRouter } from 'vue-router';

const router = useRouter()
const isLandscape = ref(false)
const gameType = ref(GameTypes.GAME_TYPE_DOUBLES);
const firstServer = ref(CurrentServer.SERVER_LEFT_EVEN);
const leftOddPlayer = ref("")
const leftEvenPlayer = ref("")
const rightEvenPlayer = ref("")
const rightOddPlayer = ref("")
const formIsLoading = ref(false)
const gameStore = useGameStore()
gameStore.setBackendUrl(import.meta.env.VITE_PROXY_URL)

const errorMessage = ref("")

onMounted(() => {
    handleOrientationChange();
    window.addEventListener('resize', handleOrientationChange);
})

onBeforeUnmount(() => {
    window.removeEventListener('resize', handleOrientationChange);
});

const handleOrientationChange = () => {
    isLandscape.value = window.matchMedia("(orientation: landscape)").matches;
};

const switchSides = () => {
    const old = {
        "leftOddPlayer": leftOddPlayer.value,
        "leftEvenPlayer": leftEvenPlayer.value,
        "rightEvenPlayer": rightEvenPlayer.value,
        "rightOddPlayer": rightOddPlayer.value,
    }

    leftOddPlayer.value = old.rightOddPlayer
    leftEvenPlayer.value = old.rightEvenPlayer
    rightEvenPlayer.value = old.leftEvenPlayer
    rightOddPlayer.value = old.leftOddPlayer
}

const handleStartGame = async () => {
    formIsLoading.value = true

    const body: GameStartRequestSchema = {
        gameType: GameStartRequestSchemaGameTypeEnum.Singles as GameStartRequestSchemaGameTypeEnum,
        leftEvenPlayerName: leftEvenPlayer.value,
        rightEvenPlayerName: rightEvenPlayer.value,
        servingSide: firstServer.value === CurrentServer.SERVER_LEFT_EVEN ? GameStartRequestSchemaServingSideEnum.LeftEven : GameStartRequestSchemaServingSideEnum.RightEven
    }

    if (gameType.value === GameTypes.GAME_TYPE_DOUBLES) {
        body.gameType = GameStartRequestSchemaGameTypeEnum.Doubles
        body.leftOddPlayerName = leftOddPlayer.value
        body.rightOddPlayerName = rightOddPlayer.value
    }

    // store loaded after request right away
    const res = await gameStore.startGame({
        gameStartRequestSchema: body
    })

    if (res instanceof Error) {
        // to make sure a new game is always started
        localStorage.removeItem("game")

        formIsLoading.value = false
        errorMessage.value = res.message
        return
    }

    router.push({
        name: 'game/playing',
        params: { id: res.game.id }
    })
}
</script>

<template>
    <div>
        <div v-if="!isLandscape" class="main-content">
            <div class="home-section">
                <RouterLink to="/dashboard"> <img class="back-link" :src="homeImage" alt="home button image"
                        width="30px" height="30px">
                </RouterLink>
            </div>
            <fieldset class="setup-form-container">
                <p class="error-message" id="error-message" v-if='errorMessage !== ""'>{{ errorMessage }}</p>

                <legend>Game Setup</legend>
                <form @submit.prevent="(event) => { event.preventDefault() }">
                    <fieldset>
                        <legend>Game Type</legend>
                        <label>
                            <input type="radio" name="gameType" :value="GameTypes.GAME_TYPE_SINGLES" v-model="gameType"
                                :disabled="formIsLoading" />
                            Singles
                        </label>
                        <label>
                            <input type="radio" name="gameType" :value="GameTypes.GAME_TYPE_DOUBLES" v-model="gameType"
                                default :disabled="formIsLoading" />
                            Doubles
                        </label>
                    </fieldset>

                    <fieldset>
                        <legend>Serving Side</legend>
                        <label>
                            <input type="radio" name="servingSide" :value="CurrentServer.SERVER_LEFT_EVEN"
                                v-model="firstServer" default :disabled="formIsLoading" />
                            Left
                        </label>
                        <label>
                            <input type="radio" name="servingSide" :value="CurrentServer.SERVER_RIGHT_EVEN"
                                v-model="firstServer" :disabled="formIsLoading" />
                            Right
                        </label>
                    </fieldset>
                </form>
                <button type="button" @click="switchSides" class="button button-secondary mb-1"
                    :disabled="formIsLoading">Switch Sides</button>
            </fieldset>
            <div class="court">
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
                <div class="left-top-player squares">
                    <div v-if="gameType === GameTypes.GAME_TYPE_DOUBLES" class="form-group">
                        <input type="text" id="name" name="name" placeholder="Enter player name" v-model="leftOddPlayer"
                            required :disabled="formIsLoading">
                        <div class="loading-wrapper"><img :src="loadingImage" alt="loading left top player" width="30px"
                                height="30px"></div>
                    </div>
                </div>
                <div class="left-bottom-player squares">
                    <div :style="{ display: 'flex', flexDirection: 'column', alignItems: 'center' }">
                        <div v-if="firstServer === CurrentServer.SERVER_LEFT_EVEN" class="shuttle-wrapper">
                            <img :src="shuttlecock" width="30px" height="30px">
                        </div>

                        <div class="form-group">
                            <input type="text" id="name" name="name" placeholder="Enter player name"
                                v-model="leftEvenPlayer" required :disabled="formIsLoading">
                            <div class="loading-wrapper"><img :src="loadingImage" alt="loading left top player"
                                    width="30px" height="30px"></div>
                        </div>
                    </div>
                </div>
                <div class="right-top-player squares">
                    <div :style="{ display: 'flex', flexDirection: 'column', alignItems: 'center' }">
                        <div v-if="firstServer === CurrentServer.SERVER_RIGHT_EVEN" class="shuttle-wrapper">
                            <img :src="shuttlecock" width="30px" height="30px">
                        </div>

                        <div class="form-group">
                            <input type="text" id="name" name="name" placeholder="Enter player name"
                                v-model="rightEvenPlayer" required :disabled="formIsLoading">
                            <div class="loading-wrapper"><img :src="loadingImage" alt="loading left top player"
                                    width="30px" height="30px"></div>
                        </div>
                    </div>
                </div>
                <div class="right-bottom-player squares">
                    <div v-if="gameType === GameTypes.GAME_TYPE_DOUBLES" class="form-group">
                        <input type="text" id="name" name="name" placeholder="Enter player name"
                            v-model="rightOddPlayer" required :disabled="formIsLoading">
                        <div class="loading-wrapper"><img :src="loadingImage" alt="loading left top player" width="30px"
                                height="30px"></div>
                    </div>
                </div>
                <!-- <div class="left-top-backline squares"></div> -->
                <!-- <div class="left-bottom-backline squares">
                </div>
                <div class="right-top-backline squares">
                    <div v-if="firstServer === CurrentServer.SERVER_RIGHT_EVEN" class="shuttle-wrapper">
                        <img :src="shuttlecock" width="30px" height="30px">
                    </div>
                </div> -->
                <!-- <div class="right-bottom-backline squares"></div> -->
            </div>
            <button type="button" class="button button-primary mb-1" :style="{ marginTop: '20px', width: '95%' }"
                :disabled="formIsLoading" @click="handleStartGame">Start Game</button>

        </div>
        <div v-else class="main-content">
            <div class="home-section">
                <RouterLink to="/dashboard"> <img class="back-link" :src="homeImage" alt="home button image"
                        width="30px" height="30px">
                </RouterLink>
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
                    <div v-if="gameType === GameTypes.GAME_TYPE_DOUBLES" class="form-group">
                        <input type="text" id="name" name="name" placeholder="Enter player name" v-model="leftOddPlayer"
                            required :disabled="formIsLoading">
                        <div class="loading-wrapper"><img :src="loadingImage" alt="loading left top player" width="30px"
                                height="30px"></div>
                    </div>
                </div>
                <div class="left-bottom-player squares">
                    <div class="form-group">
                        <input type="text" id="name" name="name" placeholder="Enter player name"
                            v-model="leftEvenPlayer" required :disabled="formIsLoading">
                        <div class="loading-wrapper"><img :src="loadingImage" alt="loading left top player" width="30px"
                                height="30px"></div>
                    </div>
                </div>
                <div class="right-top-player squares">
                    <div class="form-group">
                        <input type="text" id="name" name="name" placeholder="Enter player name"
                            v-model="rightEvenPlayer" required :disabled="formIsLoading">
                        <div class="loading-wrapper"><img :src="loadingImage" alt="loading left top player" width="30px"
                                height="30px"></div>
                    </div>
                </div>
                <div class="right-bottom-player squares">
                    <div v-if="gameType === GameTypes.GAME_TYPE_DOUBLES" class="form-group">
                        <input type="text" id="name" name="name" placeholder="Enter player name"
                            v-model="rightOddPlayer" required :disabled="formIsLoading">
                        <div class="loading-wrapper"><img :src="loadingImage" alt="loading left top player" width="30px"
                                height="30px"></div>
                    </div>
                </div>
                <div class="left-top-backline squares"></div>
                <div class="left-bottom-backline squares">
                    <div v-if="firstServer === CurrentServer.SERVER_LEFT_EVEN" class="shuttle-wrapper">
                        <img :src="shuttlecock" width="30px" height="30px">
                    </div>
                </div>
                <div class="right-top-backline squares">
                    <div v-if="firstServer === CurrentServer.SERVER_RIGHT_EVEN" class="shuttle-wrapper">
                        <img :src="shuttlecock" width="30px" height="30px">
                    </div>
                </div>
                <div class="right-bottom-backline squares"></div>
            </div>
            <fieldset class="setup-form-container">
                <p class="error-message" id="error-message" v-if='errorMessage !== ""'>{{ errorMessage }}</p>

                <legend>Game Setup</legend>
                <form @submit.prevent="(event) => { event.preventDefault() }">
                    <fieldset>
                        <legend>Game Type</legend>
                        <label>
                            <input type="radio" name="gameType" :value="GameTypes.GAME_TYPE_SINGLES" v-model="gameType"
                                :disabled="formIsLoading" />
                            Singles
                        </label><br />
                        <label>
                            <input type="radio" name="gameType" :value="GameTypes.GAME_TYPE_DOUBLES" v-model="gameType"
                                default :disabled="formIsLoading" />
                            Doubles
                        </label>
                    </fieldset>

                    <fieldset>
                        <legend>Serving Side</legend>
                        <label>
                            <input type="radio" name="servingSide" :value="CurrentServer.SERVER_LEFT_EVEN"
                                v-model="firstServer" default :disabled="formIsLoading" />
                            Left
                        </label>
                        <label>
                            <input type="radio" name="servingSide" :value="CurrentServer.SERVER_RIGHT_EVEN"
                                v-model="firstServer" :disabled="formIsLoading" />
                            Right
                        </label>
                    </fieldset>
                </form>
                <button type="button" @click="switchSides" class="button button-primary mb-1"
                    :disabled="formIsLoading">Switch Sides</button>

                <button type="button" class="button button-primary mb-1" :disabled="formIsLoading"
                    @click="handleStartGame">Start Game</button>
            </fieldset>
        </div>
    </div>
</template>

<style scoped>
@media only screen and (orientation: landscape) {
    .mt-1 {
        margin-top: 1rem;
    }

    .mb-1 {
        margin-bottom: 1rem;
    }

    .main-content {
        display: flex;
        flex-direction: row;
        min-width: 100vw;
    }

    .court {
        display: grid;
        grid-template-columns: repeat(8, 1fr);
        grid-template-rows: repeat(4, 1fr);
        width: 70vw;
        height: 90vh;
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
        grid-column: 5 / span 3;
        grid-row: 1 / span 2;
    }

    .right-bottom-player {
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
        grid-column: 8 / span 1;
        grid-row: 1 / span 2;
    }

    .right-bottom-backline {
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
}

@media only screen and (orientation: portrait) {
    .mt-1 {
        margin-top: 1rem;
    }

    .mb-1 {
        margin-bottom: 1rem;
    }

    .main-content {
        display: flex;
        flex-direction: column;
        min-width: 90vw;
        min-height: 100vh;
        padding: 10px;
    }

    .court {
        display: grid;
        grid-template-columns: repeat(8, 1fr);
        grid-template-rows: repeat(2, 1fr);
        width: 98%;
        height: 40vh;
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
        grid-column: 1 / span 4;
        grid-row: 1 / span 1;
    }

    .left-bottom-player {
        grid-column: 1 / span 4;
        grid-row: 2 / span 1;
    }

    .right-top-player {
        grid-column: 5 / span 4;
        grid-row: 1 / span 1;
    }

    .right-bottom-player {
        grid-column: 5 / span 4;
        grid-row: 2 / span 1;
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
        grid-column: 8 / span 1;
        grid-row: 1 / span 2;
    }

    .right-bottom-backline {
        grid-column: 8 / span 1;
        grid-row: 3 / span 2;
    }

    .setup-form-container {
        display: flex;
        flex-direction: column;
        max-width: 98%;
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
}


.back-link {
    margin: 5px;
}
</style>