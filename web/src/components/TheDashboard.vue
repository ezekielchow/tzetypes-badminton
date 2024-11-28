<script setup lang="ts">
import type { GetRecentStatistics200Response } from '@/repositories/clients/private';
import { useGameStore } from '@/stores/game-store';
import { useUserStore } from '@/stores/user-store';
import { DateTime } from 'luxon';
import { computed, onMounted, reactive, ref } from 'vue';

const errorMessage = ref('')
const userEmail = ref('')
const recentStatistics = reactive({} as GetRecentStatistics200Response)

const userStore = useUserStore()
userStore.setBackendUrl(import.meta.env.VITE_PROXY_URL)

const gameStore = useGameStore()
gameStore.setBackendUrl(import.meta.env.VITE_PROXY_URL)

onMounted(async () => {
    await getUserEmail()
    await getRecentStatistics()
})

const submitLogout = async () => {
    const res = await userStore.logout()

    if (res instanceof Error) {
        errorMessage.value = res.message
        return
    }

    errorMessage.value = ""
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
    const res = await gameStore.getRecentStatistics()

    if (res instanceof Error) {
        errorMessage.value = res.message
        return
    }

    const data = res as GetRecentStatistics200Response

    recentStatistics.gameRecentStatistics = data.gameRecentStatistics

    console.log(recentStatistics.gameRecentStatistics);

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

</script>

<template>
    <div class="dashboard-container">
        <div class="header">
            <div>
                <h2 class="headlines"><b>Dashboard</b></h2>
                <h5>Welcome {{ userEmail }},</h5>

            </div>

            <form @submit.prevent="submitLogout">

                <div class="actions">
                    <button type="submit" class="button button-secondary">Logout</button>
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
                <button class="button button-primary">
                    New Game
                </button>
            </RouterLink>
            <div class="recent-statistics">
                <div class="statistic-header">
                    <span>{{ `MOST RECENT ${recentStatistics?.gameRecentStatistics?.gameCount ?
                        recentStatistics.gameRecentStatistics.gameCount : ""} GAME(S)` }}</span>
                    <span>{{ `LAST UPDATED: ${updatedAtDecorated}` }}</span>
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
    font-size: 0.6rem;
    font-weight: bold;
}

.mt {
    margin-top: 2rem;
}
</style>
