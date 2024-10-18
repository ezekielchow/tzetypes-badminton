<script setup lang="ts">
import router from '@/router';
import { usePlayerStore } from '@/stores/player-store';
import { ref } from 'vue';

const errorMessage = ref("")
const formIsLoading = ref(false)
const playerName = ref("")

const playerStore = usePlayerStore()
playerStore.setBackendUrl(import.meta.env.VITE_BACKEND_URL)

const addPlayer = async () => {
    formIsLoading.value = true

    const res = await playerStore.addPlayer({
        addPlayerRequestSchema: {
            name: playerName.value
        }
    })

    if (res instanceof Error) {
        formIsLoading.value = false
        errorMessage.value = res.message
        return
    }

    formIsLoading.value = false
    router.push('/players')
}

</script>

<template>
    <div class="container">
        <h2><b>Add Player</b></h2>

        <div class="navigation-wrapper">
            <nav aria-label="Breadcrumb">
                <ol class="breadcrumb">
                    <RouterLink to="/dashboard">Dashboard</RouterLink> /
                    <RouterLink to="/players">Players</RouterLink> /
                    <RouterLink to="/players/add">Add Player</RouterLink>
                </ol>
            </nav>
        </div>

        <div class="form-container">
            <p class="error-message" id="error-message" v-if='errorMessage !== ""'>{{ errorMessage }}</p>
            <form @submit.prevent="addPlayer">
                <div class="form-group">
                    <label for="name">Name</label>
                    <input type="text" id="name" name="name" placeholder="Enter player name" v-model="playerName"
                        required>
                </div>
                <button type="submit" class="submit-btn" :disabled="formIsLoading">Submit</button>
            </form>
        </div>
    </div>
</template>


<style scoped>
.container {
    display: flex;
    flex-direction: column;
}

.navigation-wrapper {
    display: flex;
    justify-content: space-between;
    align-items: center;
    border-radius: 0.25rem;
    padding: 0.5rem;
    border: 0.5px solid silver;
}

.form-container {
    margin-top: 1rem;
    align-self: center;
}
</style>
