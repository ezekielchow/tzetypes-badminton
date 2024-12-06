<script setup lang="ts">
import type { GetRecentStatistics200Response } from '@/repositories/clients/private';
import { auth } from '@/services/firebase';
import { MyPrivateApi } from '@/services/requests-private';
import { useGameStore } from '@/stores/game-store';
import { useUserStore } from '@/stores/user-store';
import { signOut } from 'firebase/auth';
import { DateTime } from 'luxon';
import { computed, onMounted, reactive, ref } from 'vue';
import ButtonComponent from './ButtonComponent.vue';

const errorMessage = ref('')
const userEmail = ref('')
const recentStatistics = reactive({} as GetRecentStatistics200Response)
const winsWidth = ref(0)
const lossWidth = ref(0)
const pointsWonWidth = ref(0)
const pointsLossWidth = ref(0)
const shortestRallyWidth = ref(0)
const longestRallyWidth = ref(0)
const avgTimePerPointDecorated = ref("")
const avgTimePerPointWonDecorated = ref("")
const avgTimePerPointLostDecorated = ref("")
const avgTimePerGameDecorated = ref("")
const showRecentStatisticsHint = ref(false)
const isLoading = ref(false)

const userStore = useUserStore()
userStore.setBackendUrl(import.meta.env.VITE_PROXY_URL)

const gameStore = useGameStore()
gameStore.setBackendUrl(import.meta.env.VITE_PROXY_URL)

onMounted(async () => {
    await getUserEmail()
    await getRecentStatistics()
})

const submitLogout = async () => {
    isLoading.value = true

    try {
        await signOut(auth)
        const privApi = new MyPrivateApi(import.meta.env.VITE_PROXY_URL)
        privApi.deleteSession()
        isLoading.value = false
    } catch (error: any) {
        errorMessage.value = error.message
        isLoading.value = false
    }
}

const getUserEmail = async () => {
    const res = await userStore.getCurrentUser()
    if (res instanceof Error) {
        errorMessage.value = res.message
        return
    }

    errorMessage.value = ""
    userEmail.value = res.email
}

const getRecentStatistics = async () => {
    try {

        const res = await gameStore.getRecentStatistics()

        const data = res as GetRecentStatistics200Response
        if (data.gameRecentStatistics.wins === 0 && data.gameRecentStatistics.losses === 0) {
            showRecentStatisticsHint.value = true
            return
        }

        recentStatistics.gameRecentStatistics = data.gameRecentStatistics

        winsWidth.value = recentStatistics.gameRecentStatistics.wins / (recentStatistics.gameRecentStatistics.wins + recentStatistics.gameRecentStatistics.losses) * 100
        lossWidth.value = 100 - winsWidth.value

        pointsWonWidth.value = (recentStatistics.gameRecentStatistics.pointsWon / recentStatistics.gameRecentStatistics.totalPoints) * 100
        pointsLossWidth.value = 100 - pointsWonWidth.value

        shortestRallyWidth.value = recentStatistics.gameRecentStatistics.shortestRallySeconds / (recentStatistics.gameRecentStatistics.shortestRallySeconds + recentStatistics.gameRecentStatistics.longestRallySeconds) * 100
        longestRallyWidth.value = 100 - shortestRallyWidth.value

        avgTimePerPointDecorated.value = decorateSeconds(recentStatistics.gameRecentStatistics.averageTimePerPointSeconds)
        avgTimePerPointWonDecorated.value = decorateSeconds(recentStatistics.gameRecentStatistics.averageTimePerPointWonSeconds)
        avgTimePerPointLostDecorated.value = decorateSeconds(recentStatistics.gameRecentStatistics.averageTimePerPointLostSeconds)
        avgTimePerGameDecorated.value = decorateSecondsWithHour(recentStatistics.gameRecentStatistics.averageTimePerGameSeconds)

    } catch (error: any) {
        showRecentStatisticsHint.value = true
        return
    }
}

const updatedAtDecorated = computed(() => {
    if (recentStatistics?.gameRecentStatistics?.updatedAt) {
        const parsedDate = DateTime.fromFormat(
            recentStatistics.gameRecentStatistics.updatedAt.split(".")[0],
            'yyyy-MM-dd HH:mm:ss',
            { zone: 'utc' } // Assume it's in UTC
        );
        return parsedDate.toFormat('d MMM yyyy');
    }

    return ""
})

const shortestRallyDecorated = computed(() => {
    if (recentStatistics?.gameRecentStatistics?.shortestRallySeconds) {
        const timeText = decorateSeconds(recentStatistics.gameRecentStatistics.shortestRallySeconds)
        const wonText = recentStatistics.gameRecentStatistics.shortestRallyIsWon ? 'win' : 'lost'

        return `${timeText} (${wonText})`
    }

    return ""
})

const longestRallyDecorated = computed(() => {
    if (recentStatistics?.gameRecentStatistics?.longestRallySeconds) {
        const timeText = decorateSeconds(recentStatistics.gameRecentStatistics.longestRallySeconds)
        const wonText = recentStatistics.gameRecentStatistics.longestRallyIsWon ? 'win' : 'lost'

        return `${timeText} (${wonText})`
    }

    return ""
})

const decorateSeconds = (seconds: number) => {
    const minutes = Math.floor(seconds / 60)
    const leftoverSeconds = seconds % 60

    return `${minutes}m ${leftoverSeconds}s`
}

const decorateSecondsWithHour = (seconds: number) => {
    const hours = Math.floor(seconds / 3600)
    const minutes = Math.floor((seconds % 3600) / 60)
    const leftoverSeconds = seconds % 60

    return hours > 0 ? `${hours}h ${minutes}m ${leftoverSeconds}s` : `${minutes}m ${leftoverSeconds}s`
}

</script>

<template>
    <div class="dashboard-container">
        <div class="header">
            <div>
                <h2 class="headlines"><b>Dashboard</b></h2>
                <h5>Welcome {{ userEmail }},</h5>

            </div>

            <form>
                <div class="actions">
                    <ButtonComponent type="secondary" :isLoading="isLoading" @click.prevent="submitLogout">
                        Logout
                    </ButtonComponent>
                </div>

                <div v-if="errorMessage" class="error">
                    {{ errorMessage }}
                </div>
            </form>
        </div>

        <div class="content">
            <!-- <RouterLink to="/players">
                <button class="primary-button">
                    My Players
                </button>
            </RouterLink> -->
            <RouterLink to="/game/setup" class="mt">
                <ButtonComponent type="primary">
                    New Game
                </ButtonComponent>
            </RouterLink>
            <div class="recent-statistics" v-if="showRecentStatisticsHint">
                <h5>Recent Statistics</h5>
                <span>Bookmark Games to start!</span>
            </div>
            <div class="recent-statistics" v-if="!showRecentStatisticsHint">
                <div class="statistic-header">
                    <span>{{ `MOST RECENT ${recentStatistics?.gameRecentStatistics?.gameCount ?
                        recentStatistics.gameRecentStatistics.gameCount : ""} GAME(S)` }}</span>
                    <span>{{ `LAST UPDATED: ${updatedAtDecorated}` }}</span>
                </div>
                <div class="statistics-title">
                    <div :style="{ fontWeight: 'bold' }">{{ recentStatistics?.gameRecentStatistics?.wins ?
                        recentStatistics.gameRecentStatistics.wins :
                        "0" }}</div>
                    <div class="grow center" :style="{ fontWeight: 'bold' }">WINS / LOSSES</div>
                    <div :style="{ fontWeight: 'bold' }">{{ recentStatistics?.gameRecentStatistics?.losses ?
                        recentStatistics.gameRecentStatistics.losses : "0" }}</div>
                </div>
                <div class="statistics-body">
                    <div class="percentage-bar statistic-left" :style="{ width: `${winsWidth}%` }"></div>
                    <div class="percentage-bar statistic-right" :style="{ width: `${lossWidth}%` }"></div>
                </div>
                <div class="statistics-title">
                    <div :style="{ fontWeight: 'bold' }">{{ recentStatistics?.gameRecentStatistics?.pointsWon ?
                        recentStatistics.gameRecentStatistics.pointsWon :
                        "0" }}</div>
                    <div class="grow center" :style="{ fontWeight: 'bold' }">POINTS WON / LOST</div>
                    <div :style="{ fontWeight: 'bold' }">{{ recentStatistics?.gameRecentStatistics?.totalPoints ?
                        recentStatistics.gameRecentStatistics.totalPoints -
                        recentStatistics.gameRecentStatistics.pointsWon : "0" }}</div>
                </div>
                <div class="statistics-body">
                    <div class="percentage-bar statistic-left" :style="{ width: `${pointsWonWidth}%` }"></div>
                    <div class="percentage-bar statistic-right" :style="{ width: `${pointsLossWidth}%` }"></div>
                </div>
                <div class="statistics-title">
                    <div :style="{ fontWeight: 'bold' }">{{ shortestRallyDecorated }}</div>
                    <div class="grow center" :style="{ fontWeight: 'bold', textAlign: 'center' }">SHORTEST / LONGEST
                        RALLY</div>
                    <div :style="{ fontWeight: 'bold', textAlign: 'right' }">{{ longestRallyDecorated }}</div>
                </div>
                <div class="statistics-body">
                    <div class="percentage-bar statistic-left" :style="{ width: `${shortestRallyWidth}%` }"></div>
                    <div class="percentage-bar statistic-right" :style="{ width: `${longestRallyWidth}%` }"></div>
                </div>
                <div class="statistics-title" :style="{ justifyContent: 'space-between' }">
                    <div :style="{ fontWeight: 'bold' }">AVG TIME / GAME:</div>
                    <div :style="{ fontWeight: 'bold' }">{{ avgTimePerGameDecorated }}</div>
                </div>
                <div class="statistics-title" :style="{ justifyContent: 'space-between' }">
                    <div :style="{ fontWeight: 'bold' }">AVG TIME / POINT:</div>
                    <div :style="{ fontWeight: 'bold' }">{{ avgTimePerPointDecorated }}</div>
                </div>
                <div class="statistics-title" :style="{ justifyContent: 'space-between' }">
                    <div :style="{ fontWeight: 'bold' }">AVG TIME / POINT WON:</div>
                    <div :style="{ fontWeight: 'bold' }">{{ avgTimePerPointWonDecorated }}</div>
                </div>
                <div class="statistics-title" :style="{ justifyContent: 'space-between' }">
                    <div :style="{ fontWeight: 'bold' }">AVG TIME / POINT LOST:</div>
                    <div :style="{ fontWeight: 'bold' }">{{ avgTimePerPointLostDecorated }}</div>
                </div>
            </div>
        </div>
    </div>
</template>


<style scoped>
.dashboard-container {
    display: flex;
    flex-direction: column;
    padding: 1rem;
}

.header {
    display: flex;
    justify-content: space-between;
}

.content {
    display: flex;
    flex-direction: column;
    align-items: center;
    margin-top: 2rem;
}

.recent-statistics {
    background-color: #DCEBFF;
    display: flex;
    flex-direction: column;
    width: 100%;
    padding: 10px;
    margin-top: 40px;
}

.statistic-header {
    display: flex;
    justify-content: space-between;
    font-size: 0.7rem;
    font-weight: bold;
    text-decoration: underline;
}

.mt {
    margin-top: 2rem;
}

.comparison-container {
    display: flex;
    width: 100%;
    margin-top: 10px;
    color: #DCEBFF;
}

.comparison-bar {
    padding: 2px;
}

.statistics-title {
    display: flex;
    justify-content: space-evenly;
    margin-top: 15px;
}

.statistics-body {
    display: flex;
    min-width: 100%;
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
</style>
