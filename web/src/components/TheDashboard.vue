<script setup lang="ts">
import { useUserStore } from '@/store/user-store';
import { ref } from 'vue';

const errorMessage = ref('')

const submitLogout = async () => {

    const userStore = useUserStore()
    userStore.setBackendUrl(import.meta.env.VITE_BACKEND_URL)
    const res = await userStore.logout()

    if (res instanceof Error) {
        errorMessage.value = res.message
        return
    }

    errorMessage.value = ""
}
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
