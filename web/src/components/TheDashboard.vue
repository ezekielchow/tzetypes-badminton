<script setup lang="ts">
import { MyApi } from '@/services/requests';
import { useDashboardStore } from '@/store/dashboardStore';
import { ref } from 'vue';

const errorMessage = ref('')

const submitLogout = async () => {

    const myApi = new MyApi(import.meta.env.VITE_BACKEND_URL)

    try {
        await myApi.logoutRequest()
        myApi.deleteSession()

        errorMessage.value = ''

    } catch (error: any) {

        if (error.response) {
            const errorBody = await error.response.json() // Parse the error response body as JSON
            errorMessage.value = `Error: ${errorBody.message || 'Something went wrong'}`
            return
        }

        errorMessage.value = "Network error or unexpected error occurred"
        return
    }

}


const dashboardStore = useDashboardStore()
await dashboardStore.test()
</script>


<template>
    <div>
        <h1>dashboard</h1>

        <form @submit.prevent="submitLogout">

            <div class="actions">
                <button type="submit">Logout</button>
            </div>

            <div v-if="errorMessage" class="error">
                {{ errorMessage }}
            </div>
        </form>

    </div>
</template>


<style scoped></style>
