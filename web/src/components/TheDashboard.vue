<script setup lang="ts">
import { useUserStore } from '@/store/user-store';
import { ref } from 'vue';

const errorMessage = ref('')
const userEmail = ref('')

const userStore = useUserStore()
userStore.setBackendUrl(import.meta.env.VITE_BACKEND_URL)

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
    userEmail.value = res.user.email
}

getUserEmail()
</script>

<template>
    <div class="dashboard-container">
        <div class="header">
            <div>
                <h2><b>Dashboard</b></h2>
                <h5>Welcome {{ userEmail }},</h5>

            </div>

            <form @submit.prevent="submitLogout">

                <div class="actions">
                    <button type="submit">Logout</button>
                </div>

                <div v-if="errorMessage" class="error">
                    {{ errorMessage }}
                </div>
            </form>
        </div>

        <div class="content">
            <RouterLink to="/players">
                <button class="green-button">
                    My Players
                </button>
            </RouterLink>
        </div>
    </div>
</template>


<style scoped>
.dashboard-container {
    display: flex;
    flex-direction: column;
}

.header {
    display: flex;
    justify-content: space-between;
}

.content {
    display: flex;
    justify-content: center;
    margin-top: 2rem;
}
</style>
